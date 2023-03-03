package governance

import (
	"bytes"
	"crypto/sha256"
	"gitlab.com/q-dev/q-client/event"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/contracts"
	"gitlab.com/q-dev/q-client/log"
)

var (
	errConstitutionFileRequestIsEmpty = errors.New("constitution file request is empty")
)

const (
	draftsDir              = "drafts"
	constitutionFilePrefix = `Q-Constitution-AB`
	filePattern            = constitutionFilePrefix + `_0x[A-Fa-f0-9]{6}\.adoc`
)

type ConstitutionManager struct {
	baseDir        string
	db             *database
	reg            *contracts.Registry
	requiredHashes []common.Hash

	constitutionFeed *event.Feed
	storageLock      sync.Mutex
}

func NewConstitutionManager(datadir string, db *database, rm *RootManager) (*ConstitutionManager, error) {
	if rm == nil {
		err := errors.New("Constitution storage initialization error: rootManager missing")
		log.Error("Constitution storage initialization error", "error", err)
		return nil, err
	}
	manager := &ConstitutionManager{
		baseDir:          filepath.Join(datadir, "constitution-storage"),
		db:               db,
		constitutionFeed: &event.Feed{},
		reg:              rm.reg,
	}

	if errE := manager.fileExists(manager.baseDir); errE != nil {
		if err := os.Mkdir(manager.baseDir, 0755); err != nil && !os.IsNotExist(err) {
			log.Error("Constitution storage initialization error", "error", err)
			return nil, err
		}
	}
	draftsPath := filepath.Join(manager.baseDir, draftsDir)
	if errE := manager.fileExists(draftsPath); errE != nil {
		if err := os.Mkdir(draftsPath, 0755); err != nil && !os.IsNotExist(err) {
			log.Error("Constitution storage initialization error", "error", err)
			return nil, err
		}
	}
	manager.validateStorage()

	requiredHashes, _ := db.getConstitutionFileRequests()
	manager.requiredHashes = requiredHashes

	return manager, nil
}

// TODO call init
func (cm *ConstitutionManager) InitRegistry(reg *contracts.Registry) {
	cm.reg = reg
}

func (cm *ConstitutionManager) validateStorage() {
	log.Info("Validating constitution storage")
	cm.storageLock.Lock()
	defer cm.storageLock.Unlock()

	fsFiles, errF := cm.validateStorageDir()
	if errF != nil {
		log.Error("Constitution storage validation failed", "error", errF)
		return
	}
	dbFiles, errDb := cm.validateConstitutionDatabase()
	if errDb != nil {
		log.Error("Constitution storage validation failed", "error", errDb)
		return
	}

	fileRequests, errDb := cm.validateConstitutionFileRequestDatabase()
	if errDb != nil {
		log.Error("Constitution storage validation failed", "error", errDb)
		return
	}

	type removeCandidate struct {
		path             string
		errorDescription string
		delete           bool
	}

	var filesToRemove []*removeCandidate

	for _, fsFile := range fsFiles {
		if fsFile == draftsDir {
			continue
		}
		path := filepath.Join(cm.baseDir, fsFile)
		fsFileExistsInDB := false
		var fsFileCheckError error
		for _, dbFile := range dbFiles {
			if dbFile.Name == fsFile {
				fsFileExistsInDB = true

				hash := cm.getHashOfFile(path)
				if hash != dbFile.Hash {
					fsFileCheckError = errors.New("Wrong file hash") //TODO other checks?
				}

				break
			}
		}
		if !fsFileExistsInDB { //But it can be renamed, will check later
			filesToRemove = append(filesToRemove, &removeCandidate{
				path:             path,
				errorDescription: "File doesn't exist in database",
				delete:           true,
			})
			continue
		}
		if fsFileCheckError != nil {
			filesToRemove = append(filesToRemove, &removeCandidate{
				path:             path,
				errorDescription: fsFileCheckError.Error(),
				delete:           true,
			})
		}
	}

	//Files were validated early, just check if they're existing
	var resDbFiles []common.ConstitutionFile
	for _, dbFile := range dbFiles {
		path := filepath.Join(cm.baseDir, dbFile.Name)

		dbFileExistsInStorage := false
		for _, fsFile := range fsFiles {
			if dbFile.Name == fsFile {
				if errE := cm.fileExists(filepath.Join(cm.baseDir, fsFile)); errE == nil {
					dbFileExistsInStorage = true
					break
				}
			}
		}

		if !dbFileExistsInStorage {
			isRenamed := false
			//Check candidates for removal, file can be renamed. If so - return it's correct name
			for _, candidate := range filesToRemove {
				hash := cm.getHashOfFile(candidate.path)
				if hash == dbFile.Hash {
					//It means that file was renamed
					os.Rename(path, dbFile.Name)
					isRenamed = true
					candidate.delete = false
				}
			}

			//We can't find file on disk anyway
			if !isRenamed {
				log.Error("Removing file %s record from database. File doesn't exist in the storage", "path", path)
				continue
			}
		}
		resDbFiles = append(resDbFiles, dbFile)
	}

	//Remove wrong records from db
	if len(resDbFiles) != len(dbFiles) {
		log.Warn("Updating constitution storage in the database")
		if err := cm.db.saveConstitutionStorage(&resDbFiles); err != nil {
			log.Error("Constitution storage update error", "err", err)
		}

	}

	//Remove all incorrect files
	if len(filesToRemove) > 0 {
		for _, candidate := range filesToRemove {
			if candidate.delete {
				log.Error("Removing file from constitution storage directory. File doesn't exist in the database", "err", candidate.path)
				if err := os.Remove(candidate.path); err != nil {
					log.Error("Failed to delete file", "err", err, "path", candidate.path)
				}
			}
		}
	}

	//Check requests and remove the fulfilled ones
	var resRequests []common.Hash
	for _, request := range fileRequests {
		isFulfilled := false
		for _, file := range resDbFiles {
			if file.Hash == request {
				isFulfilled = true
				break
			}
		}
		if !isFulfilled {
			resRequests = append(resRequests, request)
		}
	}
	if len(resRequests) != len(fileRequests) {
		log.Warn("Updating constitution file request list in the database")
		if err := cm.db.saveConstitutionFileRequests(&resRequests); err != nil {
			log.Error("Constitution file request storage update error", "err", err)
		}
	}

}

func (cm *ConstitutionManager) validateConstitutionDatabase() ([]common.ConstitutionFile, error) {
	files, err := cm.db.getConstitutionFiles()
	if err != nil {
		return nil, errors.Wrap(err, "Error during constitution storage validation")
	}
	return files, nil
}

func (cm *ConstitutionManager) validateConstitutionFileRequestDatabase() ([]common.Hash, error) {
	requests, err := cm.db.getConstitutionFileRequests()
	if err != nil {
		return nil, errors.Wrap(err, "Error during constitution file requests validation")
	}
	return requests, nil
}

func (cm *ConstitutionManager) validateStorageDir() ([]string, error) {
	var validFileNames []string

	files, err := ioutil.ReadDir(cm.baseDir)
	if err != nil {
		return nil, errors.Wrap(err, "Error processing constitution storage path")
	}

	for _, file := range files {
		if !cm.fileHasValidName(file) {
			continue
		}
		validFileNames = append(validFileNames, file.Name())
	}
	return validFileNames, nil
}

// gov.addConstitutionFile("a.adoc","0x5c5f8818734daeee729e3ff22cec58843d24b11b0addb7b8a6b0f194f6af81c6")
func (cm *ConstitutionManager) addConstitutionFile(filename string) error {
	cm.validateStorage()

	//cm.storageLock.Lock()
	//defer cm.storageLock.Unlock()

	candidatePath := filepath.Join(cm.baseDir, draftsDir, filename)

	if errE := cm.fileExists(candidatePath); errE != nil {
		return errors.Wrap(errE, "Cannot open constitution file")
	}

	hash := cm.getHashOfFile(candidatePath)

	ok, errV := cm.isHashValid(hash)
	if errV != nil {
		return errV
	}
	if !ok {
		return errors.New("Provided file's hash is invalid")
	}

	cFile := common.ConstitutionFile{
		Name:      cm.filenameFromHash(hash),
		Hash:      hash,
		CreatedAt: time.Now().Unix(),
	}

	contents, errC := cm.getFileContents(candidatePath)
	if errC != nil {
		return errC
	}
	return cm.storeConstitutionFile(contents, cFile, true)
}

func (cm *ConstitutionManager) isHashValid(hash common.Hash) (bool, error) {
	wrapped := "Cannot check hash validity"
	if cm.reg == nil {
		err := errors.New("Contract registry not initialized")
		log.Error(wrapped, "error", err)
		return false, err
	}

	cv := cm.reg.ConstitutionVoting()
	cvE, err := cv.ConstitutionVotingFilterer.FilterProposalExecuted(nil, nil)
	defer cvE.Close()

	if err != nil {
		return false, errors.Wrap(err, wrapped)
	}
	ok := cvE.Next()
	for ok {
		//log.Error(common.BytesToHash(cvE.Event.ConstitutionHash[:]).String())
		if bytes.Equal(cvE.Event.ConstitutionHash[:], hash.Bytes()) {
			return true, nil
		}
		ok = cvE.Next()
	}

	return false, nil
}

func (cm *ConstitutionManager) filenameFromHash(constitutionHash common.Hash) string {
	return constitutionFilePrefix + "_" + constitutionHash.String()[:8] + ".adoc"
}

func (cm *ConstitutionManager) storeConstitutionFile(contents []byte, cFile common.ConstitutionFile, regular bool) error {
	dbFiles, err := cm.db.getConstitutionFiles()
	if err != nil {
		return errors.New("Failed to load constitution storage from the database")
	}

	resDir := cm.baseDir

	if regular {
		//No need to update same file
		for _, dbFile := range dbFiles {
			if dbFile.Hash == cFile.Hash {
				//return errors.New("Cannot add constitution file. Same file already exists")
				return nil //exists, skip
			}
		}
	} else {
		resDir = filepath.Join(resDir, draftsDir)
	}

	newFilePath := filepath.Join(resDir, cFile.Name)

	if errRn := os.WriteFile(newFilePath, contents, 0644); errRn != nil {
		log.Error("Cannot save new constitution file to file storage", "error", newFilePath)
		return errRn
	}

	//We only need add a record to the DB in case if file is not a draft
	if regular {
		dbFiles = append(dbFiles, cFile)
		if errSave := cm.db.saveConstitutionStorage(&dbFiles); errSave != nil {
			log.Error("Cannot save new constitution file to database", "error", errSave)
			return errSave
		}
	}
	log.Info("Constitution file with hash" + cFile.Hash.String() + " added successfully")
	return nil
}

func (cm *ConstitutionManager) addConstitutionFileRequest(requiredHash *common.Hash) (*common.Hash, error) {
	cm.validateStorage()

	cm.storageLock.Lock()
	defer cm.storageLock.Unlock()

	if requiredHash == nil || requiredHash.String() == "" {
		return nil, errors.New("Hash cannot be empty")
	}

	hashes, err := cm.db.getConstitutionFileRequests()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load constitution file request storage from the database")
	}
	for _, dbHash := range hashes {
		if dbHash == *requiredHash {
			return nil, errors.New("File with the requested hash already exists in the request list")
		}
	}

	dbFiles, errF := cm.db.getConstitutionFiles()
	if errF != nil {
		return nil, errors.Wrap(err, "Failed to load constitution file storage from the database")
	}
	for _, file := range dbFiles {
		if file.Hash == *requiredHash {
			return nil, errors.New("File with the requested hash already exists in the storage")
		}
	}

	hashes = append(hashes, *requiredHash)
	if errSave := cm.db.saveConstitutionFileRequests(&hashes); errSave != nil {
		return nil, errors.Wrap(errSave, "Failed to save constitution file requests to the database")
	}

	cm.requiredHashes = hashes

	return requiredHash, nil
}

func (cm *ConstitutionManager) fileHasValidName(file fs.FileInfo) bool {
	match, _ := regexp.MatchString(filePattern, file.Name())

	if (file.IsDir() && file.Name() != draftsDir) || (!file.IsDir() && !match) {
		log.Warn("Constitution storage error", errors.New("filename doesn't match requirements"))
		os.Remove(filepath.Join(cm.baseDir, file.Name()))
		return false
	}
	return true
}

func (cm *ConstitutionManager) getHashOfFile(filePath string) common.Hash {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error("Cannot get file hash", "error", err)
	}
	bytes = cm.preformatFileContents(bytes)
	return cm.getHashByFileContent(bytes)
}

func (cm *ConstitutionManager) preformatFileContents(bytes []byte) []byte {
	contents := string(bytes)
	//TODO
	//re := regexp.MustCompile(`\r?\n`)
	//contents = re.ReplaceAllString(contents, "\r")
	////What else?
	return []byte(contents)
}

func (cm *ConstitutionManager) getHashByFileContent(bytes []byte) common.Hash {
	bytes = cm.preformatFileContents(bytes)

	hasher := sha256.New()
	hasher.Write(bytes)

	var value common.Hash
	value.SetBytes(hasher.Sum(nil)[:])
	return value
}

// TODO
func (cm *ConstitutionManager) fileExists(path string) error {
	_, errF := os.OpenFile(path, os.O_RDONLY, 0)
	return errF
}

func (cm *ConstitutionManager) populateConstitutionFileRequest() common.ConstitutionFilesRequest {
	var res common.ConstitutionFilesRequest
	if cm.requiredHashes != nil {
		res.Hashes = cm.requiredHashes
	}
	return res
}

// fileName should be just name without basedir
func (cm *ConstitutionManager) getFileContents(fileName string) ([]byte, error) {
	contents, errC := os.ReadFile(fileName)
	if errC != nil {
		return nil, errC
	}

	return cm.preformatFileContents(contents), nil
}

func (cm *ConstitutionManager) getDraftFiles() ([]common.ConstitutionFile, error) {
	var resfiles []common.ConstitutionFile

	draftsPath := filepath.Join(cm.baseDir, draftsDir)

	files, err := ioutil.ReadDir(draftsPath)
	if err != nil {
		return nil, errors.Wrap(err, "Error processing constitution storage drafts path")
	}

	for _, file := range files {
		path := filepath.Join(draftsPath, file.Name())

		cont, errC := cm.getFileContents(path)
		if errC != nil {
			return nil, errC
		}

		resfiles = append(resfiles, common.ConstitutionFile{
			Name: file.Name(),
			Hash: cm.getHashByFileContent(cont),
		})
	}

	return resfiles, nil
}
