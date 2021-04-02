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
	errInvalidSignature = errors.New("list contains invalid signature")
	errHashMismatch     = errors.New("hash mismatch")
)

// RootManager stores root and exclusion lists.
// todo: shouldn't be exported
type RootManager struct {
	keystore Keystore

	db *database

	lock           sync.Mutex
	targetRootFeed *event.Feed
	current        *rootSet
	target         *rootSet
	proposed       *rootSet

	exLock           sync.Mutex
	targetExListFeed *event.Feed
	exclusionSet     *exclusionSet
	targetExSet      *exclusionSet
	proposedExSet    *exclusionSet
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

		targetRootFeed: &event.Feed{},
		current:        currentRootSet,
		target:         desiredRootSet,
		proposed:       proposedRootSet,

		targetExListFeed: &event.Feed{},
		exclusionSet:     db.getCurrentExclusionSet(),
		targetExSet:      db.getDesiredExclusionSet(),
		proposedExSet:    db.getProposedExclusionSet(),
	}

	return manager, nil
}

// ExclusionSet returns set of excluded validators addresses.
func (s *RootManager) ExclusionSet() map[common.Address]uint64 {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	set := make(map[common.Address]uint64)
	if s.exclusionSet == nil {
		return set
	}

	for addr, block := range s.exclusionSet.addrToBlock {
		set[addr] = block
	}

	return set
}

func (s *RootManager) isRootNode() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

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

func (s *RootManager) findUnlocked(set *rootSet) []common.Address {
	var unlocked []common.Address
	for _, addr := range set.rootAddresses {
		if s.keystore.IsUnlocked(addr) {
			unlocked = append(unlocked, addr)
		}
	}

	return unlocked
}

// unsafe for concurrent usage
// lock exLock externally first
func (s *RootManager) upgradeExclusionSet(set *exclusionSet) {
	s.exclusionSet = set
	s.db.saveCurrentExclusionSet(set)

	log.Info("Upgraded exclusion list", "hash", set.hash.Hex(), "timestamp", set.timestamp)

	if s.targetExSet != nil && s.targetExSet.timestamp <= set.timestamp {
		if s.targetExSet.hash != set.hash {
			log.Info("Dropping obsolete desired exclusion set", "timestamp", set.timestamp)
		}

		s.targetExSet = nil
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

	if s.target == nil || s.target.timestamp > set.timestamp {
		return
	}

	if s.target.hash != set.hash {
		log.Info("Dropping outdated desired root list",
			"hash", s.target.hash.Hex(),
			"timestamp", s.target.timestamp)
	}

	s.target = nil
	s.db.deleteDesiredRootSet()
}

func (s *RootManager) currentRootSet() *rootSet {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.current.copy()
}

func (s *RootManager) desiredRootSet() *rootSet {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.target.copy()
}

func (s *RootManager) proposedRootSet() *rootSet {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.proposed.copy()
}

func (s *RootManager) isAcceptableExclusionSet(set *exclusionSet) bool {
	if s.exclusionSet != nil && set.timestamp <= s.exclusionSet.timestamp {
		return false
	}

	return s.currentRootSet().isEnoughExSetSignatures(set)
}

func (s *RootManager) currentExclusionSet() *exclusionSet {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	return s.exclusionSet.copy()
}

func (s *RootManager) desiredExclusionSet() *exclusionSet {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	return s.targetExSet.copy()
}

func (s *RootManager) proposedExclusionSet() *exclusionSet {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	return s.proposedExSet.copy()
}
