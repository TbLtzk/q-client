package governance

import (
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
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/log"
)

const (
	draftsDir              = "drafts"
	constitutionFilePrefix = "Q-Constitution-AB"
	filePattern            = constitutionFilePrefix + `\d{1,}_[0-9a-fA-F]{8}.adoc`
)

type ConstitutionFile struct {
	Name             string       `json:"name"`
	ConstitutionHash *common.Hash `json:"constitutionHash"`
	FileHash         *common.Hash `json:"fileHash"`
	CreatedAt        int64
}

type ConstitutionManager struct {
	baseDir     string
	db          *database
	reg         *contracts.Registry
	rm          *RootManager
	storageLock sync.Mutex
}

func NewConstitutionManager(datadir string, db *database, rm *RootManager) (*ConstitutionManager, error) {
	if rm == nil {
		err := errors.New("Constitution storage initialization error: rootManager missing")
		log.Error("Constitution storage initialization error", err)
		return nil, err
	}
	manager := &ConstitutionManager{
		baseDir: filepath.Join(datadir, "constitution-storage"),
		db:      db,
		rm:      rm,
	}

	if errE := manager.fileExists(manager.baseDir); errE != nil {
		if err := os.Mkdir(manager.baseDir, 0755); err != nil && !os.IsNotExist(err) {
			log.Error("Constitution storage initialization error", err)
			return nil, err
		}
	}
	draftsPath := filepath.Join(manager.baseDir, draftsDir)
	if errE := manager.fileExists(draftsPath); errE != nil {
		if err := os.Mkdir(draftsPath, 0755); err != nil && !os.IsNotExist(err) {
			log.Error("Constitution storage initialization error", err)
			return nil, err
		}
	}
	manager.validateStorage()

	return manager, nil
}

//TODO call init
func (cm *ConstitutionManager) InitRegistry(reg *contracts.Registry) {
	cm.reg = reg
}

func (cm *ConstitutionManager) validateStorage() {
	log.Info("Validating constitution storage")
	cm.storageLock.Lock()
	defer cm.storageLock.Unlock()

	fsFiles, errF := cm.validateStorageDir()
	if errF != nil {
		log.Error("Constitution storage validation failed", errF)
		return
	}
	dbFiles, errDb := cm.validateDatabase()
	if errDb != nil {
		log.Error("Constitution storage validation failed", errDb)
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

				hash := cm.getFileHash(path)
				if &hash != dbFile.FileHash {
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
	var resDbFiles []ConstitutionFile
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
				hash := cm.getFileHash(candidate.path)
				if &hash == dbFile.FileHash {
					//It means that file was renamed
					os.Rename(path, dbFile.Name)
					isRenamed = true
					candidate.delete = false
				}
			}

			//We can't find file on disk anyway
			if !isRenamed {
				log.Error("Removing file record from database. File doesn't exist in the storage", path)
				continue
			}
		}
		resDbFiles = append(resDbFiles, dbFile)
	}

	//Remove wrong records from db
	if len(resDbFiles) != len(dbFiles) {
		log.Warn("Updating constitution storage in the database")
		if err := cm.db.saveConstitutionStorage(&resDbFiles); err != nil {
			log.Error("Constitution storage update error", err)
		}

	}

	//Remove all incorrect files
	if len(filesToRemove) > 0 {
		for _, candidate := range filesToRemove {
			if candidate.delete {
				log.Error("Removing file from constitution storage directory. File doesn't exist in the database", candidate.path)
				if err := os.Remove(candidate.path); err != nil {
					log.Error("Failed to delete file", candidate.path, err)
				}
			}
		}
	}
}

func (cm *ConstitutionManager) validateDatabase() ([]ConstitutionFile, error) {
	//cm.storageLock.Lock()
	//defer cm.storageLock.Unlock() //storage locked in func validateStorage()

	files, err := cm.db.getConstitutionFiles()
	if err != nil {
		return nil, errors.Wrap(err, "Error during constitution storage validation")
	}
	return files, nil
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

func (cm *ConstitutionManager) addConstitutionFile(filename string, constitutionHash *common.Hash) error {
	if !cm.rm.isRootNode() {
		return errors.New("Cannot add constitution file. Not a root node")
	}

	cm.validateStorage()

	cm.storageLock.Lock()
	defer cm.storageLock.Unlock()

	if constitutionHash == nil || constitutionHash == &(common.Hash{}) {
		return errors.New("Constitution hash cannot be empty")
	}

	newFileName := constitutionFilePrefix + "_" + constitutionHash.String()[:6] + ".adoc"

	dbFiles, err := cm.db.getConstitutionFiles()
	if err != nil {
		return errors.New("Failed to load constitution storage from database")
	}

	candidatePath := filepath.Join(cm.baseDir, draftsDir, filename)
	hash := cm.getFileHash(candidatePath)

	//No need to update same file
	for _, dbFile := range dbFiles {
		if dbFile.ConstitutionHash == constitutionHash && dbFile.FileHash == &hash {
			return errors.New("Cannot add constitution file. Same file already exists")
		}
	}

	if errE := cm.fileExists(candidatePath); errE != nil {
		log.Error("Cannot open constitution file", errE)
		return err
	}

	//Check constitution change?
	cFile := ConstitutionFile{
		Name:             newFileName,
		ConstitutionHash: constitutionHash,
		FileHash:         &hash,
		CreatedAt:        time.Now().Unix(),
	}

	newFilePath := filepath.Join(cm.baseDir, newFileName)
	if errRn := os.Rename(candidatePath, newFilePath); errRn != nil {
		log.Error("Cannot save new constitution file to file storage", newFilePath)
		return errRn
	}

	dbFiles = append(dbFiles, cFile)
	if errSave := cm.db.saveConstitutionStorage(&dbFiles); errSave != nil {
		log.Error("Cannot save new constitution file to database", errSave)
		return errSave
	}

	return nil
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

func (cm *ConstitutionManager) getFileHash(filePath string) common.Hash {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error("Cannot get file hash", err)
	}
	return cm.getHashByFileContent(bytes)
}

func (cm *ConstitutionManager) getHashByFileContent(bytes []byte) common.Hash {
	h := crypto.Keccak256(bytes)
	var value common.Hash
	value.SetBytes(h)
	return value
}

//TODO
func (cm *ConstitutionManager) fileExists(path string) error {
	_, errF := os.OpenFile(path, os.O_RDONLY, 0)
	return errF
}
