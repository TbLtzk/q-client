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
)

// RootManager stores root and exclusion lists.
// todo: shouldn't be exported
type RootManager struct {
	keystore Keystore

	db *database

	rootLock        sync.Mutex
	desiredRootFeed *event.Feed

	active   *rootSet
	desired  *rootSet
	proposed *rootSet

	exLock        sync.Mutex
	desiredExFeed *event.Feed

	activeExSet   *exclusionSet
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

	activeRootSet, err := db.getActiveRootSet()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get active root set")
	}

	if activeRootSet != nil && activeRootSet.timestamp >= defaultRootSet.timestamp {
		log.Info("Using saved root set", "hash", activeRootSet.hash.Hex())
	} else {
		activeRootSet = defaultRootSet
		db.saveActiveRootSet(activeRootSet)
		log.Info("Using predefined root set", "hash", activeRootSet.hash.Hex())
	}

	desiredRootSet, _ := db.getDesiredRootSet()
	proposedRootSet, _ := db.getProposedRootSet()

	manager := &RootManager{
		db:       db,
		keystore: ks,

		desiredRootFeed: &event.Feed{},
		active:          activeRootSet,
		desired:         desiredRootSet,
		proposed:        proposedRootSet,

		desiredExFeed: &event.Feed{},
		activeExSet:   db.getActiveExclusionSet(),
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
	if s.activeExSet == nil {
		return set
	}

	for addr, block := range s.activeExSet.addrToBlock {
		set[addr] = block
	}

	return set
}

func (s *RootManager) isRootNode() bool {
	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	return s.isMember(s.active.rootAddresses)
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
	for _, addr := range s.getActiveRootSet().rootAddresses {
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
	s.activeExSet = set
	s.db.saveActiveExclusionSet(set)

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
	s.active = set
	s.db.saveActiveRootSet(s.active)

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

func (s *RootManager) getActiveRootSet() *rootSet {
	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	return s.active.copy()
}

func (s *RootManager) getDesiredRootSet() *rootSet {
	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	return s.desired.copy()
}

func (s *RootManager) getProposedRootSet() *rootSet {
	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	return s.proposed.copy()
}

func (s *RootManager) isAcceptableExclusionSet(set *exclusionSet) bool {
	if s.activeExSet != nil && set.timestamp <= s.activeExSet.timestamp {
		return false
	}

	return s.getActiveRootSet().isEnoughExSetSignatures(set)
}

func (s *RootManager) getActiveExclusionSet() *exclusionSet {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	return s.activeExSet.copy()
}

func (s *RootManager) getDesiredExclusionSet() *exclusionSet {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	return s.desiredExSet.copy()
}

func (s *RootManager) getProposedExclusionSet() *exclusionSet {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	return s.proposedExSet.copy()
}
