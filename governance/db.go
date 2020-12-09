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
		return nil, errors.Wrap(err, "corrupted data")
	}

	set, err := newRootSet(&rootList)
	if err != nil {
		panic(errors.Wrap(err, "malformed root set"))
	}

	return set, nil
}

var (
	currentRootKey  = []byte("current-root-set")
	desiredRootKey  = []byte("desired-root-set")
	proposedRootKey = []byte("proposed-root-set")
)
