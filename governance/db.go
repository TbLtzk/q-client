package governance

import (
	"bytes"
	"encoding/json"
	"gitlab.com/q-dev/q-client/common"
	"math/big"

	"github.com/pkg/errors"
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
	activeExclusionKey   = []byte("current-exclusion-set")
	desiredExclusionKey  = []byte("desired-exclusion-set")
	proposedExclusionKey = []byte("proposed-exclusion-set")
)

var (
	approvalPrefix          = []byte("wl-cosign-")
	approvalLastBlockPrefix = []byte("wl-block")
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

func (db *database) saveApprovalRecord(approval common.RootNodeApproval, updateNum bool) error {

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
	if updateNum {
		return db.updateApprovalBlockNumber(approval.BlockNumber)
	}

	return nil
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
