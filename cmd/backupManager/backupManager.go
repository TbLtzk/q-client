package backupManager

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/pkg/errors"
	"github.com/prometheus/tsdb/fileutil"
	"github.com/robfig/cron/v3"
	"gitlab.com/q-dev/q-client/log"
	"golang.org/x/sys/unix"
)

const defaultRotationPeriod = 3 // days

type Config struct {
	Datadir        string
	BackupDir      string // If not set - temp dir will be used instead
	Incremental    bool
	S3Export       bool
	S3Bucket       string
	S3KeyPrefix    string
	AWSRegion      string
	AWSAccessKey   string
	AWSSecretKey   string
	AWSToken       string
	CronSpec       *string
	RotationPeriod int
}

type BackupInfo struct {
	Name          string
	Timestamp     time.Time
	Parent        string // In case of incremental backup
	Hash          string
	ParentHash    string
	BackupSize    int64
	DirectorySize int64
	Files         []ArchivedFile
}

type ArchivedFile struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

type BackupManager struct {
	config           *Config
	prepareCallback  func() error
	onJobEndCallback func(*BackupInfo) error
	s3Client         *s3.Client
	lock             fileutil.Releaser
	stop             chan struct{}
	isOnSchedule     bool
	cron             *cron.Cron
}

func NewBackupManager(conf *Config, prepareCallback func() error, onJobEndCallback func(*BackupInfo) error) (*BackupManager, error) {
	log.Info("Backup manager. Initializing")

	if conf.Datadir == "" {
		return nil, errors.New("Initialization failed: datadir is empty.")
	}
	if conf.S3Export && conf.S3Bucket == "" {
		return nil, errors.New("Initialization failed: S3 bucket is empty.")
	}
	if conf.S3Export && conf.AWSRegion == "" {
		return nil, errors.New("Initialization failed: AWS region is empty.")
	}
	if conf.S3Export && conf.AWSAccessKey == "" {
		return nil, errors.New("Initialization failed: AWS access key is empty.")
	}
	if conf.S3Export && conf.AWSSecretKey == "" {
		return nil, errors.New("Initialization failed: AWS secret key is empty.")
	}

	mgr := &BackupManager{
		config:           conf,
		prepareCallback:  prepareCallback,
		onJobEndCallback: onJobEndCallback,
		stop:             make(chan struct{}),
	}

	if conf.S3Export {
		log.Info("Backup manager. S3 config has been found. Initializing S3 client")

		s3Config, err := config.LoadDefaultConfig(
			context.Background(),
			config.WithRegion(conf.AWSRegion),
		)
		if err != nil {
			return nil, errors.Wrap(err, "Cannot initialize AWS config")
		}
		mgr.s3Client = s3.NewFromConfig(s3Config)
	}

	return mgr, nil
}

// Wait blocks until the node is closed.
func (mgr *BackupManager) Wait() {
	<-mgr.stop
}

func (mgr *BackupManager) Stop() {
	if mgr.cron != nil {
		mgr.cron.Stop()
	}
	close(mgr.stop)
}

func (mgr *BackupManager) ScheduleExport() error {
	log.Info("Backup manager. Scheduling backup job")

	if mgr.config.CronSpec != nil {
		mgr.cron = cron.New()

		// var job cron.Job{func() {}}
		// job.Run()

		var err error
		var entryID cron.EntryID

		entryID, err = mgr.cron.AddFunc(*mgr.config.CronSpec, func() {
			defer mgr.cron.Remove(entryID)
			if err := mgr.CreateBackup(); err != nil {
				log.Error("Scheduled backup failed", "error", err)
			}
		})
		if err != nil {
			return errors.Wrap(err, "Failed to create cron job")
		}
		mgr.isOnSchedule = true
		mgr.cron.Start()
	} else {
		return errors.New("Cannot schedule backup job: cron spec is empty.")
	}
	return nil
}

func (mgr *BackupManager) prepareToBackup() error {
	log.Info("Backup manager. Preparing for a job")

	_, err := os.Stat(mgr.config.Datadir)
	if err != nil {
		return errors.Wrap(err, "Preparation failed: data directory does not exist.")
	}

	lock, _, errLock := fileutil.Flock(filepath.Join(mgr.config.Datadir, "LOCK"))
	if err != nil {
		return errors.Wrap(errLock, "Preparation failed: . Failed to acquire database lock")
	}
	mgr.lock = lock

	return nil
}

func (mgr *BackupManager) prepareToRestore() error {
	log.Info("Backup manager. Preparing for a job")

	_, err := os.Stat(mgr.config.Datadir)
	if err == nil {
		size, errSize := mgr.getDirectorySize(mgr.config.Datadir)
		if errSize != nil {
			return errors.Wrap(errSize, "\tPreparation failed: failed to get directory size.")
		}
		if size > 0 {
			return errors.New("Preparation failed: data directory is not empty.")
		}
	} else {
		log.Info("\tCreating data directory")
		if err := os.MkdirAll(mgr.config.Datadir, 0755); err != nil {
			return errors.Wrap(err, "\tPreparation failed: failed to create data directory")
		}
	}
	return nil
}

func (mgr *BackupManager) postExecute() error {
	// log.Info("Backup manager. Running post-execute actions")
	if mgr.lock != nil {
		return mgr.lock.Release()
	}
	return nil
}

func (mgr *BackupManager) RestoreBackup(backupName *string) error {
	startTime := time.Now()
	// This one will be run in any case
	err := mgr.prepareToRestore()
	if err != nil {
		return err
	}
	defer mgr.postExecute() // TODO handle error?

	log.Info("Backup manager. Backup restoration job started")

	if backupName == nil {
		latestBackupName, errName := mgr.getLastFullBackupName()

		if latestBackupName == "" {
			return errors.Wrap(errName, "Failed to restore backup: no backups found")
		}
		backupName = &latestBackupName
	} else {
		// Check if backup exists. Skip this check if S3 export is enabled
		if !mgr.config.S3Export {
			backupPath := *backupName
			if !strings.Contains(*backupName, mgr.config.BackupDir) {
				backupPath = filepath.Join(mgr.config.BackupDir, *backupName)
			}

			if _, errStat := os.Stat(backupPath); err != nil {
				return errors.Wrap(errStat, "Failed to restore backup: backup does not exist")
			}
		}
	}

	// File exists. Restore it
	if errRestore := mgr.restoreBackup(*backupName); errRestore != nil {
		return errors.Wrap(errRestore, "Failed to restore backup. No backups found")
	}

	log.Info("Backup manager. Backup restoration job finished successfully.", "Time spent", fmt.Sprint(time.Since(startTime).String()))

	return nil
}

func (mgr *BackupManager) CreateBackup() error {
	startTime := time.Now()

	// This one will be run in any case
	if err := mgr.prepareToBackup(); err != nil {
		return err
	}

	// This one will be run only if prepareCallback is set
	if mgr.prepareCallback != nil {
		if errPrepare := mgr.prepareCallback(); errPrepare != nil {
			return errors.Wrap(errPrepare, "Failed to prepare for the mgr job")
		}
	}

	var results *BackupInfo
	defer func() {
		if mgr.onJobEndCallback != nil {
			// We need to call callback in any case. Even if job failed.
			// Results are ignored as callback isn't the part of the job.
			mgr.onJobEndCallback(results)
		}
	}()

	defer mgr.postExecute() // TODO handle error?

	log.Info("Backup manager. Backup job started")

	timestamp := startTime.UTC()
	backupDir := os.TempDir()

	if mgr.config.BackupDir != "" {
		// If backup dir doesn't exist - create one
		if _, errE := os.OpenFile(mgr.config.BackupDir, os.O_RDONLY, 0); errE != nil {
			log.Info("\tBackup directory doesn't exist. Creating")
			if errMk := os.MkdirAll(mgr.config.BackupDir, 0755); errMk != nil && !os.IsNotExist(errMk) {
				return errors.Wrap(errMk, "Backup directory initialization error")
			}
		}
		backupDir = mgr.config.BackupDir
	} else {
		// Backup dir is not set. Check if S3 export is enabled.
		if !mgr.config.S3Export {
			return errors.New("Cannot create archive: backup dir is not set and S3 export is disabled. Please set backup dir or enable S3 export.")
		}
	}

	freeSpace, err := mgr.getDiskFreeSpace(mgr.config.Datadir)
	if err != nil {
		return errors.Wrap(err, "Failed to calculate free disk space for the job")
	}

	datadirSize, err := mgr.getDirectorySize(mgr.config.Datadir)
	if err != nil {
		return errors.Wrap(err, "Failed to calculate directory size")
	}

	if freeSpace < datadirSize {
		return errors.New("Not enough free space left.")
	}

	var backupFiles map[string]string
	backupFilename := fmt.Sprintf("%d", timestamp.Unix()) // edgecase and default value if no backups found
	var lastBackupInfo *BackupInfo

	backupFiles, lastBackupInfo, err = mgr.getIncrementalBackupFilesWithHashes("")
	if err != nil || backupFiles == nil {
		// Ignore error. It's not critical. Otherwise, full backup will be created
		backupFiles = make(map[string]string)
	}
	// If incremental is enabled - let's name backup file according to the part name
	if mgr.config.Incremental && lastBackupInfo != nil {
		fileName := strings.TrimSuffix(lastBackupInfo.Name, filepath.Ext(lastBackupInfo.Name))
		if len(fileName) == 10 {
			// It is a full backup
			backupFilename = fileName + "_1"
		} else if strings.Contains(fileName, "_") {
			// It is an incremental backup. Let's increase the number
			preefix := strings.Split(fileName, "_")[0]
			suffix := strings.Split(fileName, "_")[1]
			if suffix != "" {
				if lastNumber, err := strconv.ParseUint(suffix, 10, 64); err == nil {
					backupFilename = preefix + "_" + strconv.FormatUint(lastNumber+1, 10)
				}
			}
		}
	}

	// Create the archive file.
	backupFileFullName := filepath.Join(backupDir, backupFilename) + ".tar.gz"
	backupFile, err := os.Create(backupFileFullName)
	if err != nil {
		return errors.Wrap(err, "Failed to create the archive file")
	}

	// Create the archive file.
	logFilename := filepath.Join(backupDir, backupFilename) + ".json"
	logFile, err := os.Create(logFilename)
	if err != nil {
		return errors.Wrap(err, "Failed to create the archive log file")
	}

	// We're using temp directory. Need to clean up after yourself
	if mgr.config.BackupDir == "" {
		defer func() {
			if err := os.Remove(backupFileFullName); err != nil {
				log.Warn("\tFailed to remove archive temp file", err)
			}
			if err := os.Remove(logFilename); err != nil {
				log.Warn("\tFailed to remove archive temp file", err)
			}
		}()
	}

	gzipWriter := gzip.NewWriter(backupFile)
	tarWriter := tar.NewWriter(gzipWriter)

	var files []ArchivedFile
	i := 0

	log.Info("\tScanning data directory")
	err = filepath.Walk(mgr.config.Datadir, func(path string, info os.FileInfo, someErr error) error {
		if info.IsDir() {
			return nil
		}

		if i > 20 {
			return nil
		}

		// Skip keystore files. It's a bad practice to store them in the same directory as the data.
		if matched, _ := regexp.MatchString("^.*(keystore|pwd|pass|password).*$", path+info.Name()); matched {
			log.Info("\tSkipping keystore/password file: " + info.Name())
			return nil
		}
		// Special case for the keystore file. It's a json file with 491 bytes size.
		if info.Size() == 491 {
			dat, err := os.ReadFile(path)
			if err == nil {
				// This is simplified keystore file. Skip it
				var kf struct {
					Address string `json:"address"`
					Crypto  struct {
					} `json:"crypto"`
					Id string `json:"id"`
				}
				if err := json.Unmarshal(dat, &kf); err == nil {
					return nil
				}
			}
		}

		hash := mgr.getHashOfFile(path)
		shortName := strings.TrimPrefix(path, mgr.config.Datadir)

		// Check if file was changed since last backup. If incremental backup is enabled. Otherwise file will be added to the archive anyway
		if hashStr, ok := backupFiles[shortName]; ok {
			if hashStr == hash {
				log.Info("\tSkipping unchanged file: " + path)
				return nil
			}
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		header.Name = shortName

		// Add the file to the tar archive.
		err = tarWriter.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tarWriter, file)
		if err != nil {
			return err
		}

		files = append(files, ArchivedFile{
			Name: header.Name,
			Hash: hash,
		})

		i++

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "Failed to create the archive")
	}
	tarWriter.Close()
	gzipWriter.Close()

	resultFile, _ := backupFile.Stat()
	backupFile.Close()

	var parentHash string
	var parentName string
	if lastBackupInfo != nil {
		parentHash = lastBackupInfo.Hash
		parentName = lastBackupInfo.Name
	}

	jr, err := json.Marshal(BackupInfo{
		Name:          backupFilename,
		Timestamp:     timestamp,
		BackupSize:    resultFile.Size(),
		DirectorySize: datadirSize,
		Hash:          mgr.getHashOfFile(backupFileFullName),
		Parent:        parentName,
		ParentHash:    parentHash,
		Files:         files,
	})

	_, errLog := logFile.Write(jr)
	if errLog != nil {
		return errors.Wrap(err, "Failed to write log file")
	}

	logFile.Close()

	log.Info("\tUploading files to S3")
	if mgr.config.S3Export {
		if errPut := mgr.uploadFilesToS3(backupFile, logFile); err != nil {
			return errors.Wrap(errPut, "Failed to upload archive to the S3 storage")
		}
	}

	log.Info("\tPerforming rotation")
	if err := mgr.performRotation(); err != nil {
		log.Error("Scheduled backup rotation failed", "error", err)
	}

	log.Info("Backup manager. Job finished successfully:\n" +
		"\tTime spent: " + time.Since(startTime).String() + "\n" +
		"\tBackup file: " + backupFileFullName + "\n" +
		"\tLog file: " + logFilename + "\n")

	return nil
}

func (mgr *BackupManager) uploadFilesToS3(files ...*os.File) error {
	for _, f := range files {
		file, err := os.Open(f.Name())
		if err != nil {
			return err
		}

		name := strings.TrimPrefix(file.Name(), mgr.config.BackupDir+"/")

		largeBuffer := bufio.NewReader(file)
		var partMiBs int64 = 10
		uploader := manager.NewUploader(mgr.s3Client, func(u *manager.Uploader) {
			u.PartSize = partMiBs * 1024 * 1024
		})
		_, errUpload := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(mgr.config.S3Bucket),
			Key:    aws.String(name),
			Body:   largeBuffer,
		})
		if errUpload != nil {
			return errors.Wrap(errUpload, "Failed to upload file to S3")
		}
	}
	return nil
}

func (mgr *BackupManager) getDiskFreeSpace(path string) (int64, error) {
	var stat unix.Statfs_t
	if err := unix.Statfs(path, &stat); err != nil {
		return 0, fmt.Errorf("failed to call Statfs: %v", err)
	}

	var bavail = stat.Bavail
	// nolint:staticcheck
	if stat.Bavail < 0 {
		bavail = 0
	}
	//nolint:unconvert
	return int64(bavail) * int64(stat.Bsize), nil
}

func (mgr *BackupManager) getDirectorySize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && f.Name()[0] != '.' {
			size += f.Size()
		}
		return nil
	})
	return size, err
}

func (mgr *BackupManager) restoreBackup(s string) error {
	var err error

	infos, err := mgr.getIncrementalBackupList(s)
	if err != nil {
		return errors.Wrap(err, "Failed to get incremental backup list")
	}
	for i, info := range infos {
		name := info.Name + ".tar.gz"
		if mgr.s3Client != nil {
			log.Info("\tDownloading and restoring backup from S3", "part", i+1, "of", len(infos))
			err = mgr.restoreBackupFromS3(name)
		} else {
			log.Info("Restoring backup from local storage", "part", i+1, "of", len(infos))
			err = mgr.restoreBackupFromLocal(name)
		}
		if err != nil {
			return errors.Wrap(err, "Failed to restore backup from S3")
		}
	}
	return nil
}

func (mgr *BackupManager) restoreBackupFromS3(filename string) error {
	// First, check if we have enough space to restore the backup
	backupSize, err := mgr.getBackupSizeFromS3(filename)
	if err != nil {
		return errors.Wrap(err, "Failed to get backup file size")
	}
	freeSpace, err := mgr.getDiskFreeSpace(mgr.config.Datadir)
	if err != nil {
		return errors.Wrap(err, "Failed to get free disk space")
	}

	// Check if we have enough space to restore the backup
	// We need at least 2x the size of the backup to be sure we have enough space
	if backupSize*2 > freeSpace {
		return errors.New("Not enough disk space to restore the backup")
	}

	result, err := mgr.s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(mgr.config.S3Bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return errors.Wrap(err, "Failed to get backup file")
	}
	defer result.Body.Close()

	// Unpack the backup file
	if err := mgr.unpackBackup(result.Body); err != nil {
		return errors.Wrap(err, "Failed to unpack backup file")
	}

	return nil
}

func (mgr *BackupManager) restoreBackupFromLocal(s string) error {
	// Unpack the backup file
	backupFile, err := os.Open(mgr.config.BackupDir + "/" + s)
	if err != nil {
		return errors.Wrap(err, "Failed to open backup file")
	}

	defer backupFile.Close()

	if err := mgr.unpackBackup(backupFile); err != nil {
		return errors.Wrap(err, "Failed to unpack backup file")
	}

	return nil
}

func (mgr *BackupManager) unpackBackup(file io.Reader) error {
	targetDir := mgr.config.Datadir

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tarReader := tar.NewReader(gzr)

	// Iterate through the files in the archive.
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			// End of archive
			break
		}
		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			// Create a directory
			filePath, err := sanitizeFilePath(targetDir, header.Name)
			if err != nil {
				return err
			}
			if err := os.MkdirAll(filePath, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			// Create a file
			filePath, err := sanitizeFilePath(targetDir, filepath.Dir(header.Name))
			if err != nil {
				return err
			}
			if err := os.MkdirAll(filePath, 0755); err != nil {
				return err
			}

			filePath, err = sanitizeFilePath(targetDir, header.Name)
			if err != nil {
				return err
			}
			outFile, err := os.Create(filePath)
			if err != nil {
				return err
			}
			defer outFile.Close()

			// Copy the file data
			totalRead := int64(0)
			for {
				n, err := io.CopyN(outFile, tarReader, 1024)
				totalRead += n
				// print bytes read followed by a carriage return
				fmt.Printf("Bytes read: %d\r", totalRead)
				if err != nil {
					if err == io.EOF {
						fmt.Println()
						break
					}
					return err
				}
			}
		}
	}

	return nil
}

func sanitizeFilePath(d, t string) (v string, err error) {
	v = filepath.Join(d, t)
	if strings.HasPrefix(v, filepath.Clean(d)) {
		return v, nil
	}

	return "", fmt.Errorf("%s: %s", "content filepath is tainted", t)
}

func (mgr *BackupManager) getBackupSizeFromS3(filename string) (int64, error) {
	// We can download and unmarshall the json file to get the size of the backup
	// However, this is not very efficient, so we will just list the objects in the bucket
	// and get the size of the backup file

	result, err := mgr.s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(mgr.config.S3Bucket),
		Prefix: aws.String(filename),
	})
	if err != nil {
		return 0, err
	} else {
		for _, content := range result.Contents {
			return *content.Size, nil
		}
	}
	return 0, err
}

// TODO preserve incremental backups
func (mgr *BackupManager) performRotation() error {
	period := int64(mgr.config.RotationPeriod)
	if period == 0 {
		period = defaultRotationPeriod
	}

	// Get the list of backups
	backupList, err := mgr.getBackupNames("")
	if err != nil {
		return errors.Wrap(err, "Failed to get backup list")
	}

	// If the number of backups is less than the rotation period, we don't need to do anything
	if len(backupList) == 0 || len(backupList) <= mgr.config.RotationPeriod {
		return nil // Nothing to do
	}

	candidates := []string{}

	// Get the list of backups that are older than the rotation period
	for _, s := range backupList {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return errors.Wrap(err, "Failed to parse backup date")
		}
		tm := time.Unix(i, 0)

		desiredPeriod := time.Hour * 24 * time.Duration(period)
		timePassed := time.Since(tm)
		// Check if the backup is older than the rotation period
		if timePassed > desiredPeriod {
			candidates = append(candidates, s)
		}
	}

	// Delete the oldest backups
	for _, candidate := range candidates {
		if err := mgr.deleteBackup(candidate); err != nil {
			return errors.Wrap(err, "Failed to delete backup")
		}
	}
	return nil
}

func (mgr *BackupManager) getIncrementalBackupFilesWithHashes(backupName string) (map[string]string, *BackupInfo, error) {
	if !mgr.config.Incremental {
		return nil, nil, nil
	}
	infos, err := mgr.getIncrementalBackupList(backupName)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to get incremental backup list")
	}
	files := make(map[string]string)
	var lastInfo *BackupInfo
	// infos are sorted by date ascending. So, if file has been changed - the last one will be the correct one
	for i := range infos {
		for _, file := range infos[i].Files {
			files[file.Name] = file.Hash
		}
		lastInfo = &infos[i]
	}

	// As a result we will have a map of files with their hashes.
	return files, lastInfo, nil
}

// Get the list of backups with the last full backup as the prefix (full included). Sorted by date descending
func (mgr *BackupManager) getIncrementalBackupList(fullBackup string) ([]BackupInfo, error) {
	log.Info("\tGetting incremental backups")
	backupName := fullBackup
	var err error

	if fullBackup == "" {
		backupName, err = mgr.getLastFullBackupName()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to get last full backup")
		}
	}
	if backupName == "" {
		return nil, nil
	}
	// Get the list of backups with the last full backup as the prefix (full included). Sorted by date descending
	backupInfos, err := mgr.getBackupInfos(backupName)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get backup list")
	}

	log.Info("\tValidating incremental backups")
	backupInfos, err = mgr.validateIncrementalBackups(backupInfos)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to validate incremental backups")
	}
	return backupInfos, nil
}

func (mgr *BackupManager) validateIncrementalBackups(incrementalBackups []BackupInfo) ([]BackupInfo, error) {
	log.Info("\t\tValidating incremental backups")

	if len(incrementalBackups) == 0 {
		return nil, nil
	}
	// if there's only one file - it's a full backup, no need to validate
	if len(incrementalBackups) == 1 {
		return incrementalBackups, nil
	}

	var resList []BackupInfo // This guy contains the list of incremental backups that are valid. Includes the last full backup

	// backups are sorted by date descending. Revert the order and follow the chain of incremental backups
	// first backup is the last full backup
	sort.Slice(incrementalBackups, func(i, j int) bool {
		return incrementalBackups[i].Timestamp.Unix() < incrementalBackups[j].Timestamp.Unix()
	})

	var previousBackup BackupInfo
	// resList = append(resList, previousBackup) //add the last full backup to the list
	validationFailed := false

	// Iterate through the list and check if the incremental backups are valid (consistency check)
	for _, backup := range incrementalBackups {
		// skip previous backup. Also skip all backups that are older than the previous backup (we don't need to check them)
		if backup.Hash == previousBackup.Hash || backup.Timestamp.Unix() < previousBackup.Timestamp.Unix() {
			continue
		}
		log.Info("\t\tValidating incremental backup: " + backup.Timestamp.String())

		// check if the backup is valid
		if !validationFailed {
			if err := mgr.validateIncrementalBackup(previousBackup, backup); err != nil {
				validationFailed = true
			}
		}

		if validationFailed {
			log.Info("\t\t\tValidation failed. Deleting backup: " + backup.Timestamp.String())
			// delete the backup - it's not valid
			if err := mgr.deleteBackup(backup.Timestamp.String()); err != nil {
				return resList, errors.Wrap(err, "Failed to delete invalid backup")
			}
		} else {
			resList = append(resList, backup)
			previousBackup = backup
		}
	}
	// as a result we'll have a list of incremental backups that are valid on top of the last full backup
	return resList, nil
}

// Compare the hash of the previous backup with the parent hash of the current backup
func (mgr *BackupManager) validateIncrementalBackup(parent BackupInfo, current BackupInfo) error {
	// for now, assume that backup files themselves are valid, so we only need to check the backup logs
	// TODO validate backup files

	if current.ParentHash != parent.Hash && current.Hash != parent.Hash {
		return errors.New("Backup is not incremental")
	}
	return nil
}

func (mgr *BackupManager) getLastFullBackupName() (string, error) {
	// Get the list of backups
	backupList, err := mgr.getBackupNames("")
	if err != nil {
		return "", errors.Wrap(err, "Failed to get last full backup")
	}
	if len(backupList) > 0 {
		for _, name := range backupList {
			if !strings.Contains(name, "_") {
				return name, nil
			}
		}
	}
	return "", nil
}

func (mgr *BackupManager) getBackupInfos(prefix string) ([]BackupInfo, error) {
	var backupInfos []BackupInfo
	var err error

	if mgr.s3Client != nil {
		backupInfos, err = mgr.getBackupInfosFromS3(prefix)
	} else {
		backupInfos, err = mgr.getBackupInfosFromLocal(prefix)
	}
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get backup list")
	}
	// Sort the backup list by date
	sort.Slice(backupInfos, func(i, j int) bool {
		return backupInfos[i].Timestamp.Unix() > backupInfos[j].Timestamp.Unix()
	})
	return backupInfos, nil
}

func (mgr *BackupManager) getBackupNames(prefix string) ([]string, error) {
	var backupList, resList []string
	var err error
	if mgr.s3Client != nil {
		backupList, err = mgr.getBackupNamesFromS3()
	} else {
		backupList, err = mgr.getBackupNamesFromLocal()
	}
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get backup list")
	}
	// Sort the backup list by date
	sort.Slice(backupList, func(i, j int) bool {
		return backupList[i] > backupList[j]
	})
	// If a prefix is specified, filter the list
	if prefix != "" {
		for _, s := range backupList {
			if strings.HasPrefix(s, prefix) {
				resList = append(resList, s)
			}
		}
		return resList, nil
	}
	return backupList, nil
}

func (mgr *BackupManager) getBackupNamesFromS3() ([]string, error) {
	var res []string
	result, err := mgr.s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(mgr.config.S3Bucket),
	})
	if err != nil {
		return nil, err
	} else {
		for _, content := range result.Contents {
			if strings.HasSuffix(*content.Key, ".tar.gz") {
				res = append(res, strings.TrimSuffix(*content.Key, ".tar.gz"))
			}
		}
		return res, nil
	}
}

func (mgr *BackupManager) getBackupNamesFromLocal() ([]string, error) {
	var res []string
	files, err := os.ReadDir(mgr.config.BackupDir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".tar.gz") {
			res = append(res, strings.TrimSuffix(f.Name(), ".tar.gz"))
		}
	}
	return res, nil
}

func (mgr *BackupManager) deleteBackup(candidate string) error {
	if mgr.s3Client != nil {
		err := mgr.deleteBackupFromS3(candidate)
		if err != nil {
			return err
		}
	}
	return mgr.deleteBackupFromLocal(candidate)
}

func (mgr *BackupManager) deleteBackupFromS3(candidate string) error {
	_, err1 := mgr.s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(mgr.config.S3Bucket),
		Key:    aws.String(candidate + ".tar.gz"),
	})
	if err1 != nil {
		return err1
	}
	_, err1 = mgr.s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(mgr.config.S3Bucket),
		Key:    aws.String(candidate + ".json"),
	})
	return err1
}

func (mgr *BackupManager) deleteBackupFromLocal(candidate string) error {
	err := os.Remove(filepath.Join(mgr.config.BackupDir, candidate+".tar.gz"))
	if err != nil {
		return err
	}
	return os.Remove(filepath.Join(mgr.config.BackupDir, candidate+".json"))
}

func (mgr *BackupManager) IsOnSchedule() bool {
	return mgr.isOnSchedule
}

func (mgr *BackupManager) getBackupInfoFromS3(backup string) (*BackupInfo, error) {
	result, err := mgr.s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(mgr.config.S3Bucket),
		Key:    aws.String(backup),
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get backup file")
	}
	defer result.Body.Close()
	var backupInfos *BackupInfo

	decoder := json.NewDecoder(result.Body)
	if err := decoder.Decode(&backupInfos); err != nil {
		return nil, errors.Wrap(err, "Failed to decode backup file")
	}

	return backupInfos, nil
}

func (mgr *BackupManager) getBackupInfoFromLocal(backup string) (*BackupInfo, error) {
	var backupInfos *BackupInfo
	file, err := os.Open(filepath.Join(mgr.config.BackupDir, backup+".json"))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open backup file")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&backupInfos); err != nil {
		return nil, errors.Wrap(err, "Failed to decode backup file")
	}
	return backupInfos, nil
}

func (mgr *BackupManager) getBackupInfosFromS3(prefix string) ([]BackupInfo, error) {
	var backupInfos []BackupInfo
	backupNames, err := mgr.getBackupNamesFromS3()
	if err != nil {
		return nil, err
	}
	for _, backup := range backupNames {
		if strings.HasPrefix(backup, prefix) {
			backupInfo, err := mgr.getBackupInfoFromS3(backup + ".json")
			if err != nil {
				return nil, err
			}
			if backupInfo.Name == "" {
				backupInfo.Name = backup
			}

			backupInfos = append(backupInfos, *backupInfo)
		}
	}
	return backupInfos, nil
}

func (mgr *BackupManager) getBackupInfosFromLocal(prefix string) ([]BackupInfo, error) {
	var backupInfos []BackupInfo
	backupNames, err := mgr.getBackupNamesFromLocal()
	if err != nil {
		return nil, err
	}
	for _, backup := range backupNames {
		if strings.HasPrefix(backup, prefix) {
			backupInfo, err := mgr.getBackupInfoFromLocal(backup + ".json")
			if err != nil {
				return nil, err
			}
			if backupInfo.Name == "" {
				backupInfo.Name = backup
			}

			backupInfos = append(backupInfos, *backupInfo)
		}
	}
	return backupInfos, nil
}

func (mgr *BackupManager) getHashOfFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()

	if file == nil {
		return ""
	}
	buf := make([]byte, 1024*1024)
	h := sha256.New()
	if _, err := io.CopyBuffer(h, file, buf); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
