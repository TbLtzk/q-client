package governance

import (
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/accounts"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/event"
	"gitlab.com/q-dev/q-client/log"
)

var (
	errInvalidSignature     = errors.New("list contains invalid signature")
	errHashMismatch         = errors.New("hash mismatch")
	errInvalidExclusionList = errors.New("invalid exclusion list")
	errInvalidRootList      = errors.New("invalid root list")
)

// RootManager stores root and exclusion lists.
// todo: shouldn't be exported
type RootManager struct {
	keystore Keystore

	db *database

	rootLock        sync.Mutex
	desiredRootFeed *event.Feed

	current  *rootSet
	desired  *rootSet
	proposed *rootSet

	exLock        sync.Mutex
	desiredExFeed *event.Feed

	currentExSet  *exclusionSet
	desiredExSet  *exclusionSet
	proposedExSet *exclusionSet
}

type Keystore interface {
	IsUnlocked(addr common.Address) bool
	SignHash(a accounts.Account, hash []byte) ([]byte, error)
}

func newRootManager(ks Keystore, datadir string, cfg *Config) (*RootManager, error) {
	db, err := newDatabase(filepath.Join(datadir, "gov"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to init gov database")
	}

	defaultRootSet, err := newRootSet(&cfg.RootList)
	if err != nil {
		return nil, errors.Wrap(err, "malformed default root list")
	}

	currentRootSet, err := db.getCurrentRootSet()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current root set")
	}

	if currentRootSet != nil && currentRootSet.timestamp >= defaultRootSet.timestamp {
		log.Info("Using saved root set", "hash", currentRootSet.hash.Hex())
	} else {
		currentRootSet = defaultRootSet
		db.saveCurrentRootSet(currentRootSet)
		log.Info("Using predefined root set", "hash", currentRootSet.hash.Hex())
	}

	desiredRootSet, _ := db.getDesiredRootSet()
	proposedRootSet, _ := db.getProposedRootSet()

	manager := &RootManager{
		db:       db,
		keystore: ks,

		desiredRootFeed: &event.Feed{},
		current:         currentRootSet,
		desired:         desiredRootSet,
		proposed:        proposedRootSet,

		desiredExFeed: &event.Feed{},
		currentExSet:  db.getCurrentExclusionSet(),
		desiredExSet:  db.getDesiredExclusionSet(),
		proposedExSet: db.getProposedExclusionSet(),
	}

	return manager, nil
}

// ExclusionSet returns set of excluded validators addresses.
func (s *RootManager) ExclusionSetValidators() map[common.Address]uint64 {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	set := make(map[common.Address]uint64)
	if s.currentExSet == nil {
		return set
	}

	for addr, block := range s.currentExSet.addrToBlock {
		set[addr] = block
	}

	return set
}

func (s *RootManager) isRootNode() bool {
	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	return s.isMember(s.current.rootAddresses)
}

func (s *RootManager) isMember(set []common.Address) bool {
	for _, addr := range set {
		if s.keystore.IsUnlocked(addr) {
			return true
		}
	}

	return false
}

func (s *RootManager) signRootSet(set *rootSet) bool {
	var isMember bool
	for _, addr := range set.rootAddresses {
		if !s.keystore.IsUnlocked(addr) {
			continue
		}

		isMember = true
		signature, err := s.keystore.SignHash(accounts.Account{Address: addr}, set.hash.Bytes())
		if err != nil {
			log.Error("failed to sign root set", "err", err)
			continue
		}

		if set.addSignature(addr, signature) {
			log.Info("signed root list", "hash", set.hash.Hex(), "signer", addr.Hex())
		}
	}

	return isMember
}

func (s *RootManager) signExclusionSet(set *exclusionSet) bool {
	var isSigned bool
	for _, addr := range s.currentRootSet().rootAddresses {
		if !s.keystore.IsUnlocked(addr) {
			continue
		}

		signature, err := s.keystore.SignHash(accounts.Account{Address: addr}, set.hash.Bytes())
		if err != nil {
			log.Error("failed to sign exclusion list", "err", err)
			continue
		}

		set.addSignature(addr, signature)
		isSigned = true
		log.Info("Signed exclusion list", "hash", set.hash.Hex(), "signer", addr.Hex())
	}

	return isSigned
}

// unsafe for concurrent usage
// lock exLock externally first
func (s *RootManager) upgradeExclusionSet(set *exclusionSet) {
	s.currentExSet = set
	s.db.saveCurrentExclusionSet(set)

	log.Info("Upgraded exclusion list", "hash", set.hash.Hex(), "timestamp", set.timestamp)

	if s.desiredExSet != nil && s.desiredExSet.timestamp <= set.timestamp {
		if s.desiredExSet.hash != set.hash {
			log.Info("Dropping obsolete desired exclusion set", "timestamp", set.timestamp)
		}

		s.desiredExSet = nil
		s.db.deleteDesiredExclusionSet()
	}

	if s.proposedExSet != nil && s.proposedExSet.timestamp <= set.timestamp {
		log.Info("Dropping obsolete proposed exclusion set", "timestamp", set.timestamp)

		s.proposedExSet = nil
		s.db.deleteProposedExclusionSet()
	}
}

// unsafe for concurrent calls.
// must be locked before calling
func (s *RootManager) upgradeRootSet(set *rootSet) {
	s.current = set
	s.db.saveCurrentRootSet(s.current)

	log.Info("Upgraded root list", "hash", set.hash.Hex(), "timestamp", set.timestamp)

	if s.desired == nil || s.desired.timestamp > set.timestamp {
		return
	}

	if s.desired.hash != set.hash {
		log.Info("Dropping outdated desired root list",
			"hash", s.desired.hash.Hex(),
			"timestamp", s.desired.timestamp)
	}

	s.desired = nil
	s.db.deleteDesiredRootSet()
}

func (s *RootManager) currentRootSet() *rootSet {
	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	return s.current.copy()
}

func (s *RootManager) desiredRootSet() *rootSet {
	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	return s.desired.copy()
}

func (s *RootManager) proposedRootSet() *rootSet {
	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	return s.proposed.copy()
}

func (s *RootManager) isAcceptableExclusionSet(set *exclusionSet) bool {
	if s.currentExSet != nil && set.timestamp <= s.currentExSet.timestamp {
		return false
	}

	return s.currentRootSet().isEnoughExSetSignatures(set)
}

func (s *RootManager) currentExclusionSet() *exclusionSet {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	return s.currentExSet.copy()
}

func (s *RootManager) desiredExclusionSet() *exclusionSet {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	return s.desiredExSet.copy()
}

func (s *RootManager) proposedExclusionSet() *exclusionSet {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	return s.proposedExSet.copy()
}
