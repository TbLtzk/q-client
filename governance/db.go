package governance

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/ethdb"
	"gitlab.com/q-dev/q-client/ethdb/leveldb"
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

func (db *database) saveCurrentRootSet(set *rootSet) error {
	return db.saveRootSet(set, currentRootKey)
}

func (db *database) saveRootSet(set *rootSet, key []byte) error {
	value, err := json.Marshal(set.makeList())
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal root list"))
	}

	return db.store.Put(key, value)
}

func (db *database) getCurrentRootSet() (*rootSet, error) {
	return db.getRootSet(currentRootKey)
}

func (db *database) getRootSet(key []byte) (*rootSet, error) {
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
	currentRootKey = []byte("current-root-set")
)
