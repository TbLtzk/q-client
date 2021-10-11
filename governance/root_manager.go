package governance

import (
	"fmt"
	"math"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/accounts"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/contracts"
	"gitlab.com/q-dev/q-client/core"
	"gitlab.com/q-dev/q-client/event"
	"gitlab.com/q-dev/q-client/log"
)

var (
	errNotRootNode      = errors.New("not a root node")
	errInvalidSignature = errors.New("list contains invalid signature")
	errHashMismatch     = errors.New("hash mismatch")

	errInvalidExclusionList          = errors.New("invalid exclusion list")
	errProposedExclusionListObsolete = errors.New("proposed exclusion list is obsolete")
	errProposedExclusionListEmpty    = errors.New("proposed exclusion list is empty")

	errProposedRootListObsolete = errors.New("proposed root list is obsolete")
	errProposedRootListEmpty    = errors.New("proposed root list is empty")
)

// RootManager stores root and exclusion lists.
// todo: shouldn't be exported
type RootManager struct {
	keystore  Keystore
	networkId uint64

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

	// initialized externaly
	bc  *core.BlockChain
	reg *contracts.Registry
}

// Config of root manager
type Config struct {
	RootList common.RootList `toml:"-"`
}

type DiffEntry struct {
	Name string
	Diff []common.Address
}

type Keystore interface {
	IsUnlocked(addr common.Address) bool
	SignHash(a accounts.Account, hash []byte) ([]byte, error)
}

func NewRootManager(ks Keystore, networkId uint64, datadir string, cfg *Config) (*RootManager, error) {
	dbName := fmt.Sprintf("gov-%d", networkId)
	db, err := newDatabase(filepath.Join(datadir, dbName))
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
		keystore:  ks,
		networkId: networkId,

		db: db,

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

func (s *RootManager) InitBlockChain(bc *core.BlockChain) {
	s.bc = bc
}

func (s *RootManager) InitRegistry(reg *contracts.Registry) {
	s.reg = reg
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
	if s.activeExSet != nil && s.activeExSet.hash == set.hash {
		log.Debug("Exclsion list is already active, skipping", "hash", set.hash.Hex(), "timestamp", set.timestamp)
		return
	}

	// If exclusion set was changed, revalidate blocks up to earliest affected one
	addrToBlock := set.addrToBlockExclusiveDiff(s.activeExSet)
	if len(addrToBlock) > 0 {
		var earliestBlock uint64 = math.MaxUint64
		for _, block := range addrToBlock {
			if block < earliestBlock {
				earliestBlock = block
			}
		}

		// Revalidate in separate goroutine to prevent possible deadlocks
		go s.bc.RevalidateChain(earliestBlock)
	}

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

func (s *RootManager) proposeExclusionSet(set *exclusionSet) (*exclusionSet, error) {
	if !s.isRootNode() {
		return nil, errNotRootNode
	}

	s.exLock.Lock()
	defer s.exLock.Unlock()

	olderThanActive := s.activeExSet != nil && set.timestamp <= s.activeExSet.timestamp
	olderThanDesired := s.desiredExSet != nil && set.timestamp <= s.desiredExSet.timestamp
	if olderThanActive || olderThanDesired {
		return nil, errProposedExclusionListObsolete
	}

	if s.signExclusionSet(set) {
		log.Info("signed desired exclusion list", "hash", set.hash.Hex())
	}

	if s.getActiveRootSet().isEnoughExSetSignatures(set) {
		s.upgradeExclusionSet(set)
	} else {
		s.desiredExSet = set
		s.desiredExFeed.Send(set.copy())

		s.db.saveDesiredExclusionSet(set)
	}

	return set, nil
}

func (s *RootManager) acceptProposedExclusionList() error {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	if s.proposedExSet == nil {
		return errProposedExclusionListEmpty
	}

	if s.desiredExSet != nil && s.proposedExSet.timestamp <= s.desiredExSet.timestamp {
		return errProposedExclusionListObsolete
	}

	if s.signExclusionSet(s.proposedExSet) {
		log.Info("Signed proposed exclusion list", "hash", s.proposedExSet.hash.Hex())
	}

	if s.getActiveRootSet().isEnoughExSetSignatures(s.proposedExSet) {
		s.upgradeExclusionSet(s.proposedExSet)
	} else {
		s.desiredExSet = s.proposedExSet
		s.desiredExFeed.Send(s.desiredExSet.copy())

		s.db.saveDesiredExclusionSet(s.desiredExSet)
	}

	s.proposedExSet = nil
	s.db.deleteProposedExclusionSet()

	return nil
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

func (s *RootManager) proposeRootSet(set *rootSet) (*rootSet, error) {
	if !s.isMember(set.rootAddresses) {
		return nil, errors.New("not a member of new list")
	}

	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	if set.timestamp <= s.active.timestamp || (s.desired != nil && set.timestamp <= s.desired.timestamp) {
		return nil, errProposedRootListObsolete
	}

	if s.signRootSet(set) {
		log.Info("signed desired root list", "hash", set.hash.Hex())
	}

	s.desired = set
	s.desiredRootFeed.Send(set.copy())

	s.db.saveDesiredRootSet(set)

	return set, nil
}

func (s *RootManager) acceptProposedRootList() error {
	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	if s.proposed == nil {
		return errProposedRootListEmpty
	}

	if s.desired != nil && s.proposed.timestamp <= s.desired.timestamp {
		return errProposedRootListObsolete
	}

	if s.signRootSet(s.proposed) {
		log.Info("Signed proposed root list", "hash", s.proposed.hash.Hex())
	}

	if s.active.isAcceptable(s.proposed) {
		s.upgradeRootSet(s.proposed)
	} else {
		s.desired = s.proposed
		s.desiredRootFeed.Send(s.desired.copy())

		s.db.saveDesiredRootSet(s.desired)
	}

	s.proposed = nil
	s.db.deleteProposedRootSet()

	return nil
}

func (s *RootManager) diffRootListByName(nameA, nameB string) ([]DiffEntry, error) {
	setA, err := s.getRootSetByName(nameA)
	if err != nil {
		return nil, err
	}

	setB, err := s.getRootSetByName(nameB)
	if err != nil {
		return nil, err
	}

	return []DiffEntry{
		{
			Name: nameA,
			Diff: s.addressDiff(setA.getAddresses(), setB.getAddresses()),
		},
		{
			Name: nameB,
			Diff: s.addressDiff(setB.getAddresses(), setA.getAddresses()),
		},
	}, nil
}

func (s *RootManager) getRootSetByName(name string) (*rootSet, error) {
	switch strings.ToLower(name) {
	case "active":
		return s.getActiveRootSet(), nil
	case "desired":
		return s.getDesiredRootSet(), nil
	case "proposed":
		return s.getProposedRootSet(), nil
	case "onchain":
		return s.getOnchainRootSet(), nil
	}

	return nil, fmt.Errorf("invalid root set name: %s", name)
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

func (s *RootManager) getOnchainRootSet() *rootSet {
	if s.reg == nil {
		return nil
	}

	roots := s.reg.Roots()
	if roots == nil {
		return nil
	}

	addresses, err := roots.GetMembers(nil)
	if err != nil {
		return nil
	}

	set, err := newRootSet(&common.RootList{
		Timestamp: uint64(time.Now().Unix()),
		Nodes:     addresses,
	})

	if err != nil {
		return nil
	}

	return set
}

func (s *RootManager) diffExclusionListByName(nameA, nameB string) ([]DiffEntry, error) {
	setA, err := s.getExclusionSetByName(nameA)
	if err != nil {
		return nil, err
	}

	setB, err := s.getExclusionSetByName(nameB)
	if err != nil {
		return nil, err
	}

	return []DiffEntry{
		{
			Name: nameA,
			Diff: s.addressDiff(setA.getAddresses(), setB.getAddresses()),
		},
		{
			Name: nameB,
			Diff: s.addressDiff(setB.getAddresses(), setA.getAddresses()),
		},
	}, nil
}

func (s *RootManager) getExclusionSetByName(name string) (*exclusionSet, error) {
	switch strings.ToLower(name) {
	case "active":
		return s.getActiveExclusionSet(), nil
	case "desired":
		return s.getDesiredExclusionSet(), nil
	case "proposed":
		return s.getProposedExclusionSet(), nil
	}

	return nil, fmt.Errorf("invalid exclusion set name: %s", name)
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

// isAcceptableExclusionSet returns true if there us enough signatures and
// exclusion set is not obsolete
func (s *RootManager) isAcceptableExclusionSet(set *exclusionSet) bool {
	if s.activeExSet != nil && set.timestamp <= s.activeExSet.timestamp {
		return false
	}

	return s.getActiveRootSet().isEnoughExSetSignatures(set)
}

// addressDiff returns set of addresses which are only first list but not in second
func (s *RootManager) addressDiff(arrA, arrB []common.Address) []common.Address {
	var diff []common.Address
	for _, addrA := range arrA {
		if !s.addressContains(arrB, addrA) {
			diff = append(diff, addrA)
		}
	}

	return diff
}

// addressContains returns true if address array contains specific address
func (s *RootManager) addressContains(arr []common.Address, addr common.Address) bool {
	for _, item := range arr {
		if item == addr {
			return true
		}
	}

	return false
}
