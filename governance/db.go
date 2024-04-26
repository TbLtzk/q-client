package governance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/pkg/errors"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/ethdb"
	"gitlab.com/q-dev/q-client/ethdb/leveldb"
	"gitlab.com/q-dev/q-client/log"
)

// database wraps level db
// for convenient access to root/exclusion lists.
type database struct {
	store ethdb.KeyValueStore
}

var (
	activeRootKey   = []byte("current-root-set")
	desiredRootKey  = []byte("desired-root-set")
	proposedRootKey = []byte("proposed-root-set")
)

var (
	activeExclusionKey      = []byte("current-exclusion-set")
	desiredExclusionKey     = []byte("desired-exclusion-set")
	proposedExclusionKey    = []byte("proposed-exclusion-set")
	quarantinedExclusionKey = []byte("quarantined-exclusion-set")
)

var (
	approvalPrefix          = []byte("wl-cosign-")
	approvalLastBlockPrefix = []byte("wl-block")
)

var (
	constitutionStorageKey     = []byte("constitution-storage")
	constitutionFileRequestKey = []byte("constitution-file-request")
	knownConstitutionFilesKey  = []byte("constitution-known-files")
)

var (
	rootQuotaKey      = []byte("root-quota")
	exclusionQuotaKey = []byte("exclusion-quota")
)

func newDatabase(path string) (*database, error) {
	db, err := leveldb.New(path, 0, 0, "gov", false)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init leveldb")
	}

	return &database{store: db}, nil
}

func (db *database) getActiveRootSet() (*rootSet, error) {
	return db.getRootSet(activeRootKey)
}

func (db *database) saveActiveRootSet(set *rootSet) {
	if err := db.saveRootSet(set, activeRootKey); err != nil {
		log.Crit("failed to save active root set", "err", err)
	}
}

func (db *database) getDesiredRootSet() (*rootSet, error) {
	return db.getRootSet(desiredRootKey)
}

func (db *database) saveDesiredRootSet(set *rootSet) {
	if err := db.saveRootSet(set, desiredRootKey); err != nil {
		log.Crit("failed to save desired root set", err, err)
	}
}

func (db *database) deleteDesiredRootSet() {
	if err := db.store.Delete(desiredRootKey); err != nil {
		log.Crit("failed to delete desired root set", "err", err)
	}
}

func (db *database) getProposedRootSet() (*rootSet, error) {
	return db.getRootSet(proposedRootKey)
}

func (db *database) saveProposedRootSet(set *rootSet) {
	if err := db.saveRootSet(set, proposedRootKey); err != nil {
		log.Crit("failed to save proposed root set", "err", err)
	}
}

func (db *database) deleteProposedRootSet() {
	if err := db.store.Delete(proposedRootKey); err != nil {
		log.Crit("failed to delete proposed root set", "err", err)
	}
}

func (db *database) getActiveExclusionSet() *exclusionSet {
	set, err := db.getExclusionSet(activeExclusionKey)
	if err != nil {
		log.Crit("failed to get active exclusion set", "err", err)
	}

	return set
}

func (db *database) saveActiveExclusionSet(set *exclusionSet) {
	if err := db.saveExclusionSet(set, activeExclusionKey); err != nil {
		log.Crit("failed to save active exclusion set")
	}
}

func (db *database) getDesiredExclusionSet() *exclusionSet {
	set, err := db.getExclusionSet(desiredExclusionKey)
	if err != nil {
		log.Crit("failed to get desired exclusion set", "err", err)
	}

	return set
}

func (db *database) saveDesiredExclusionSet(set *exclusionSet) {
	if err := db.saveExclusionSet(set, desiredExclusionKey); err != nil {
		log.Crit("failed to save desired exclusion set", "err", err)
	}
}

func (db *database) deleteDesiredExclusionSet() {
	if err := db.store.Delete(desiredExclusionKey); err != nil {
		log.Crit("failed to delete desired exclusion list")
	}
}

func (db *database) getProposedExclusionSet() *exclusionSet {
	set, err := db.getExclusionSet(proposedExclusionKey)
	if err != nil {
		log.Crit("failed to get proposed exclusion set", "err", err)
	}

	return set
}

func (db *database) saveProposedExclusionSet(set *exclusionSet) {
	if err := db.saveExclusionSet(set, proposedExclusionKey); err != nil {
		log.Crit("failed to store proposed exclusion set", "err", err)
	}
}

func (db *database) deleteProposedExclusionSet() {
	if err := db.store.Delete(proposedExclusionKey); err != nil {
		log.Crit("failed to delete proposed exclusion set", "err", err)
	}
}

func (db *database) saveRootSet(set *rootSet, key []byte) error {
	value, err := json.Marshal(set.makeList())
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal root list"))
	}

	return db.store.Put(key, value)
}

func (db *database) getRootSet(key []byte) (*rootSet, error) {
	ok, err := db.store.Has(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if key exists")
	}

	if !ok {
		return nil, nil
	}

	raw, err := db.store.Get(key)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get from db")
	}

	if raw == nil {
		return nil, nil
	}

	var rootList common.RootList
	if err := json.Unmarshal(raw, &rootList); err != nil {
		panic(errors.Wrap(err, "failed to unmarshal root list"))
	}

	set, err := newRootSet(&rootList)
	if err != nil {
		panic(errors.Wrap(err, "malformed root set"))
	}

	return set, nil
}

func (db *database) saveExclusionSet(set *exclusionSet, key []byte) error {
	raw, err := json.Marshal(set.makeList())
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal exclusion set"))
	}

	return db.store.Put(key, raw)
}

func (db *database) getExclusionSet(key []byte) (*exclusionSet, error) {
	ok, err := db.store.Has(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if key exists")
	}

	if !ok {
		return nil, nil
	}

	raw, err := db.store.Get(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get from db")
	}

	var exclusionList common.ValidatorExclusionList
	if err := json.Unmarshal(raw, &exclusionList); err != nil {
		panic(errors.Wrap(err, "failed to unmarshal exclusion list"))
	}

	set, err := newExclusionSet(&exclusionList)
	if err != nil {
		panic(errors.Wrap(err, "malformed exclusion set"))
	}

	return set, nil
}

func (db *database) getLastApprovals() *common.RootNodeApprovalList {
	bn, err := db.getApprovalLastBlockNumber()
	if err != nil {
		log.Crit("failed to get last root node approval block number", "err", err)
	}
	if bn == nil || bn.Uint64() == 0 {
		return nil
	}
	wl, errWl := db.getApprovalRecordsByBlockNumber(bn)
	if errWl != nil {
		log.Crit("failed to get root node approvals by block number", "err", errWl)
	}
	res := common.RootNodeApprovalList{
		BlockNumber: bn,
		Approvals:   wl,
	}
	return &res
}

func (db *database) updateApprovalBlockNumber(blockNumber *big.Int) error {
	dbBN, err := db.getApprovalLastBlockNumber()
	if err != nil {
		return err
	}
	if dbBN != nil {
		raw, errGet := db.store.Get(approvalLastBlockPrefix)
		if errGet != nil {
			return errors.Wrap(errGet, "failed to get from db")
		}
		dbBN.SetBytes(raw)
	} else {
		dbBN = new(big.Int)
	}

	if blockNumber.Int64() > dbBN.Int64() {
		return db.store.Put(approvalLastBlockPrefix, blockNumber.Bytes())
	}

	return nil
}

func (db *database) getApprovalLastBlockNumber() (*big.Int, error) {
	blockNumber := new(big.Int)

	has, err := db.store.Has(approvalLastBlockPrefix)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check last approval block number")
	}
	if !has {
		db.store.Put(approvalLastBlockPrefix, blockNumber.Bytes())
	}
	raw, errGet := db.store.Get(approvalLastBlockPrefix)
	if errGet != nil {
		return nil, errors.Wrap(errGet, "failed to get from db")
	}
	blockNumber.SetBytes(raw)
	return blockNumber, nil
}

func (db *database) saveApprovalRecord(approval common.RootNodeApproval) error {
	var resApprovals []common.RootNodeApproval

	if exRecords, errEx := db.getApprovalRecordsByBlockNumber(approval.BlockNumber); errEx == nil && exRecords != nil {
		for _, exApproval := range exRecords {
			if bytes.Equal(exApproval.Signature, approval.Signature) {
				continue
			}
			resApprovals = append(resApprovals, exApproval)
		}
	}
	resApprovals = append(resApprovals, approval)

	if len(resApprovals) == 0 {
		return nil
	}

	value, err := json.Marshal(resApprovals)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal root node approvals list"))
	}

	if errSave := db.store.Put(approval.GetApprovalDbKey(approvalPrefix), value); errSave != nil {
		return errSave
	}

	return db.updateApprovalBlockNumber(approval.BlockNumber)
}

func (db *database) getApprovalRecordsByBlockNumber(blockNumber *big.Int) ([]common.RootNodeApproval, error) {
	key := append(approvalPrefix, blockNumber.Bytes()...)

	ok, err := db.store.Has(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if key exists")
	}

	if !ok {
		return nil, nil
	}

	raw, err := db.store.Get(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get from db")
	}

	var signs []common.RootNodeApproval
	if err := json.Unmarshal(raw, &signs); err != nil {
		panic(errors.Wrap(err, "failed to unmarshal root node approvals"))
	}
	return signs, nil
}

func (db *database) getConstitutionFiles() ([]common.ConstitutionFile, error) {
	has, err := db.store.Has(constitutionStorageKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if constitution storage exists")
	}
	if !has {
		if errC := db.saveConstitutionStorage([]common.ConstitutionFile{}); errC != nil {
			return nil, errors.Wrap(err, "failed to initialize constitution storage")
		}
	}
	raw, errGet := db.store.Get(constitutionStorageKey)
	if errGet != nil {
		return nil, errors.Wrap(errGet, "failed to get constitution storage from db")
	}
	var files []common.ConstitutionFile
	if err := json.Unmarshal(raw, &files); err != nil {
		panic(errors.Wrap(err, "failed to unmarshal constitution storage"))
	}

	return files, nil
}

func (db *database) getConstitutionFileRequests() ([]common.Hash, error) {
	has, err := db.store.Has(constitutionFileRequestKey)

	if err != nil {
		return nil, errors.Wrap(err, "failed to check if constitution request storage exists")
	}
	if !has {
		if errC := db.saveConstitutionFileRequests([]common.Hash{}); errC != nil {
			return nil, errors.Wrap(err, "failed to initialize constitution request storage")
		}
	}
	raw, errGet := db.store.Get(constitutionFileRequestKey)
	if errGet != nil {
		return nil, errors.Wrap(errGet, "failed to get constitution file requests from db")
	}
	var hashes []common.Hash
	if err := json.Unmarshal(raw, &hashes); err != nil {
		panic(errors.Wrap(err, "failed to unmarshal constitution file requests"))
	}

	return hashes, nil
}

func (db *database) getKnownConstitutionFiles() ([]common.Hash, error) {
	has, err := db.store.Has(knownConstitutionFilesKey)

	if err != nil {
		return nil, errors.Wrap(err, "failed to check if known constitution files storage exists")
	}
	if !has {
		if errC := db.saveKnownConstitutionFiles([]common.Hash{}); errC != nil {
			return nil, errors.Wrap(err, "failed to initialize known constitution files storage")
		}
	}
	raw, errGet := db.store.Get(knownConstitutionFilesKey)
	if errGet != nil {
		return nil, errors.Wrap(errGet, "failed to get known constitution files from db")
	}
	var hashes []common.Hash
	if err := json.Unmarshal(raw, &hashes); err != nil {
		panic(errors.Wrap(err, "failed to unmarshal known constitution files"))
	}

	return hashes, nil
}

func (db *database) saveConstitutionStorage(storage []common.ConstitutionFile) error {
	raw, err := json.Marshal(storage)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal constitution storage"))
	}

	return db.store.Put(constitutionStorageKey, raw)
}

func (db *database) saveConstitutionFileRequests(storage []common.Hash) error {
	raw, err := json.Marshal(storage)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal constitution file requests"))
	}

	return db.store.Put(constitutionFileRequestKey, raw)
}

func (db *database) saveKnownConstitutionFiles(storage []common.Hash) error {
	raw, err := json.Marshal(storage)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal known constitution files"))
	}
	return db.store.Put(knownConstitutionFilesKey, raw)
}

func (db *database) addExclusionSetToQuarantine(set *exclusionSet) error {
	if set == nil {
		return nil
	}
	log.Info("set != nil (addExclusionSetToQuarantine)")

	var resSets []common.ValidatorExclusionList

	exRecords, err := db.getExclusionSetsFromQuarantine()
	if err != nil {
		log.Error("Failed to get exclusion sets from quarantine", "err", err)
	}
	log.Info(fmt.Sprintf("exRecords: %v (addExclusionSetToQuarantine)", exRecords))
	for _, exSet := range exRecords {
		if exSet.hash == set.hash {
			return nil //It's already there
		}
		resSets = append(resSets, exSet.makeList())
	}
	resSets = append(resSets, set.makeList())

	value, err := json.Marshal(resSets)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal quarantine exclusion sets"))
	}
	log.Info(fmt.Sprintf("value: %v (addExclusionSetToQuarantine)", value))

	if err := db.store.Put(quarantinedExclusionKey, value); err != nil {
		panic(errors.Wrap(err, "failed save quarantined exclusion sets"))
	}

	return nil
}

func (db *database) removeExclusionSetFromQuarantine(set *exclusionSet) (*[]exclusionSet, error) {
	if set == nil {
		return nil, nil
	}

	var resSets []exclusionSet

	exRecords, err := db.getExclusionSetsFromQuarantine()
	if err != nil {
		log.Error("Failed to get exclusion sets from quarantine", "err", err)
	}
	for _, exSet := range exRecords {
		if exSet.hash == set.hash {
			continue
		}
		resSets = append(resSets, exSet)
	}

	value, err := json.Marshal(resSets)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal quarantine exclusion sets"))
	}

	if err := db.store.Put(quarantinedExclusionKey, value); err != nil {
		panic(errors.Wrap(err, "failed save quarantined exclusion sets"))
	}

	return &resSets, nil
}

// It is most unlikely, but potentially we can have multiple exclusion sets in the quarantine
func (db *database) getExclusionSetsFromQuarantine() ([]exclusionSet, error) {
	ok, err := db.store.Has(quarantinedExclusionKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if key exists")
	}

	if !ok {
		return nil, nil
	}

	raw, err := db.store.Get(quarantinedExclusionKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get quarantined  from db")
	}

	var exclusionLists []common.ValidatorExclusionList
	if err := json.Unmarshal(raw, &exclusionLists); err != nil {
		panic(errors.Wrap(err, "failed to unmarshal quarantined exclusion lists"))
	}

	var sets []exclusionSet
	for i := range exclusionLists {
		set, err := newExclusionSet(&exclusionLists[i])
		if err != nil {
			//don't return error, just skip this list
			continue
		}
		sets = append(sets, *set)
	}

	return sets, nil
}

func (db *database) getQuarantinedExclusionSetByHash(hash *common.Hash) (*exclusionSet, error) {
	sets, err := db.getExclusionSetsFromQuarantine()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get exclusion sets from quarantine")
	}
	for _, set := range sets {
		if set.hash == *hash {
			return &set, nil
		}
	}
	return nil, nil
}

func (db *database) saveQuotaEntries(entries map[common.Address][]common.ListQuotaEntry, quotaKey []byte) error {
	var res []common.ListQuotaEntry
	for _, quotaEntries := range entries {
		res = append(res, quotaEntries...)
	}
	raw, err := json.Marshal(res)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal quota entries"))
	}
	return db.store.Put(quotaKey, raw)
}

func (db *database) getQuotaEntries(prefix []byte) (map[common.Address][]common.ListQuotaEntry, error) {
	has, err := db.store.Has(prefix)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if quota entries storage exists")
	}
	if !has {
		if errC := db.saveQuotaEntries(map[common.Address][]common.ListQuotaEntry{}, prefix); errC != nil {
			return nil, errors.Wrap(err, "failed to initialize quota entry storage")
		}
	}
	raw, errGet := db.store.Get(prefix)
	if errGet != nil {
		return nil, errors.Wrap(errGet, "failed to get quota entries from db")
	}
	var entries []common.ListQuotaEntry
	if err = json.Unmarshal(raw, &entries); err != nil {
		panic(errors.Wrap(err, "failed to unmarshal quota entries"))
	}
	res := make(map[common.Address][]common.ListQuotaEntry)
	for _, entry := range entries {
		res[entry.Author] = append(res[entry.Author], entry)
	}
	return res, nil
}
