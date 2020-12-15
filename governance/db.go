package governance

import (
	"encoding/json"

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

func newDatabase(path string) (*database, error) {
	db, err := leveldb.New(path, 0, 0, "gov")
	if err != nil {
		return nil, errors.Wrap(err, "failed to init leveldb")
	}

	return &database{store: db}, nil
}

func (db *database) getCurrentRootSet() (*rootSet, error) {
	return db.getRootSet(currentRootKey)
}

func (db *database) saveCurrentRootSet(set *rootSet) {
	if err := db.saveRootSet(set, currentRootKey); err != nil {
		log.Crit("failed to save current root set", "err", err)
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

func (db *database) getCurrentExclusionSet() *exclusionSet {
	set, err := db.getExclusionSet(currentExclusionSetKey)
	if err != nil {
		log.Crit("failed to get current exclusion set", "err", err)
	}

	return set
}

func (db *database) saveCurrentExclusionSet(set *exclusionSet) {
	if err := db.saveExclusionSet(set, currentExclusionSetKey); err != nil {
		log.Crit("failed to save current exclusion set")
	}
}

func (db *database) getDesiredExclusionSet() *exclusionSet {
	set, err := db.getExclusionSet(desiredExclusionSetKey)
	if err != nil {
		log.Crit("failed to get desired exclusion set", "err", err)
	}

	return set
}

func (db *database) saveDesiredExclusionSet(set *exclusionSet) {
	if err := db.saveExclusionSet(set, desiredExclusionSetKey); err != nil {
		log.Crit("failed to save desired exclusion set", "err", err)
	}
}

func (db *database) deleteDesiredExclusionSet() {
	if err := db.store.Delete(desiredExclusionSetKey); err != nil {
		log.Crit("failed to delete desired exclusion list")
	}
}

func (db *database) getProposedExclusionSet() *exclusionSet {
	set, err := db.getExclusionSet(proposedExclusionSetKey)
	if err != nil {
		log.Crit("failed to get proposed exclusion set", "err", err)
	}

	return set
}

func (db *database) saveProposedExclusionSet(set *exclusionSet) {
	if err := db.saveExclusionSet(set, proposedExclusionSetKey); err != nil {
		log.Crit("failed to store proposed exclusion set", "err", err)
	}
}

func (db *database) deleteProposedExclusionSet() {
	if err := db.store.Delete(proposedExclusionSetKey); err != nil {
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

var (
	currentRootKey  = []byte("current-root-set")
	desiredRootKey  = []byte("desired-root-set")
	proposedRootKey = []byte("proposed-root-set")
)

var (
	currentExclusionSetKey  = []byte("current-exclusion-set")
	desiredExclusionSetKey  = []byte("desired-exclusion-set")
	proposedExclusionSetKey = []byte("proposed-exclusion-set")
)
