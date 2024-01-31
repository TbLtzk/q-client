package governance

import (
	"encoding/json"
	"fmt"
	"math/big"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"

	"gitlab.com/q-dev/q-client/accounts"
	"gitlab.com/q-dev/q-client/accounts/keystore"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/contracts"
	"gitlab.com/q-dev/q-client/core"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/eth/downloader"
	"gitlab.com/q-dev/q-client/event"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/params"
	"gitlab.com/q-dev/q-client/sentryMonitor"
)

var (
	errNotRootNode      = errors.New("not a root node")
	errInvalidSignature = errors.New("list contains invalid signature")
	errHashMismatch     = errors.New("hash mismatch")

	errInvalidExclusionList          = errors.New("invalid exclusion list")
	errInvalidExclusionListTimestamp = errors.New("invalid exclusion list timestamp")
	errProposedExclusionListObsolete = errors.New("proposed exclusion list is obsolete")
	errProposedExclusionListEmpty    = errors.New("proposed exclusion list is empty")

	errProposedRootListObsolete = errors.New("proposed root list is obsolete")
	errProposedRootListEmpty    = errors.New("proposed root list is empty")
	errRootManagerCannotSign    = errors.New("RootManager cannot sign hash")

	//errInvalidApprovalList        = errors.New("invalid approval list")
	errProposedApprovalListEmpty  = errors.New("proposed approval list is empty")
	errInvalidApprovalBlockNumber = errors.New("proposed approval list contains wrong block number")
	errExclusionListQuotaExceeded = errors.New("exclusion list quota exceeded")
)

// RootManager stores root and exclusion lists.
// todo: shouldn't be exported
type RootManager struct {
	Keystore
	manager   *accounts.Manager
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

	approvalLock sync.Mutex
	approvalFeed *event.Feed

	// initialized externally
	bc  *core.BlockChain
	reg *contracts.Registry
	dl  *downloader.Downloader

	rewindLimit uint64

	quarantineLock    sync.Mutex
	quarantineEventCh chan *QuarantineEvent

	rootQuotaEntries      map[common.Address][]common.ListQuotaEntry
	exclusionQuotaEntries map[common.Address][]common.ListQuotaEntry

	ProposalQuotaMax        uint
	rootQuotaLock           sync.Mutex
	ProposalQuotaTimeWindow time.Duration
	exclusionQuotaLock      sync.Mutex

	ApprovalMaxFailures    uint64
	transitionBlockChecker *transitionBlockChecker
}

// Config of root manager
type Config struct {
	RootList                      common.RootList `toml:"-"`
	ProposalQuotaMax              uint
	ProposalQuotaTimeWindow       time.Duration
	ApprovalMaxFailures           uint64
	TransitionBlockVerifiedBlocks uint64
}

type DiffEntry struct {
	Name string
	Diff []common.Address
}

type QuarantineEvent struct {
	Set *exclusionSet
}

type Keystore interface {
	IsUnlocked(addr common.Address) bool
	SignHash(a accounts.Account, hash []byte) ([]byte, error)
}

func NewRootManager(am *accounts.Manager, networkId uint64, datadir string, cfg *Config) (*RootManager, error) {
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

	rootQuotaEntries, _ := db.getQuotaEntries(rootQuotaKey)
	exclusionQuotaEntries, _ := db.getQuotaEntries(exclusionQuotaKey)

	manager := &RootManager{
		manager:   am,
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

		approvalFeed: &event.Feed{},

		quarantineEventCh: make(chan *QuarantineEvent),

		rootQuotaEntries:      rootQuotaEntries,
		exclusionQuotaEntries: exclusionQuotaEntries,

		ProposalQuotaMax:        cfg.ProposalQuotaMax,
		ProposalQuotaTimeWindow: cfg.ProposalQuotaTimeWindow,

		ApprovalMaxFailures: cfg.ApprovalMaxFailures,
	}

	manager.transitionBlockChecker = newTransitionBlockChecker(
		int64(cfg.TransitionBlockVerifiedBlocks),
		manager.HandleTransitionBlockSignature,
	)

	manager.setMaxRewindLimit(uint64(params.FullImmutabilityThreshold))

	go manager.startQuarantineRoutine()

	return manager, nil
}

func (s *RootManager) InitDownloader(dl *downloader.Downloader) {
	s.dl = dl
}

func (s *RootManager) InitBlockChain(bc *core.BlockChain) {
	s.bc = bc
	s.transitionBlockChecker.initBlockChain(bc)
}

func (s *RootManager) InitRegistry(reg *contracts.Registry) {
	s.reg = reg
	s.updateAliasesOfRootSets()
}

func (s *RootManager) updateAliasesOfRootSets() {
	s.active.aliases = s.getAliasesOfRoots(s.active.rootAddresses)
	if s.desired != nil {
		s.desired.updateAliases(s.getAliasesOfRoots(s.desired.rootAddresses))
	}
	if s.proposed != nil {
		s.proposed.updateAliases(s.getAliasesOfRoots(s.proposed.rootAddresses))
	}
}

// ExclusionSetValidators returns set of excluded validators addresses.
func (s *RootManager) ExclusionSetValidators() map[common.Address][]common.BlockRange {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	set := make(map[common.Address][]common.BlockRange)
	if s.activeExSet == nil {
		return set
	}

	for addr, br := range s.activeExSet.blockRanges {
		set[addr] = append(set[addr], br...)
	}

	return set
}

func (s *RootManager) ExclusionSetTimestamp() uint64 {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	if s.activeExSet == nil {
		return 0
	}

	return s.activeExSet.timestamp
}

func (s *RootManager) isRootNode(lock bool) bool {
	if lock {
		s.rootLock.Lock()
		defer s.rootLock.Unlock()
	}

	log.Debug("isRootNode", "active root node aliases", s.active.aliases)

	for _, alias := range s.active.aliases {
		if s.IsUnlocked(alias) {
			return true
		}
	}

	log.Debug("no unlocked root nodes found")

	return false
}

func (s *RootManager) signRootSet(set *rootSet) bool {
	var isMember bool
	aliases := s.getAliasesOfRoots(s.active.rootAddresses)

	for _, addr := range s.active.rootAddresses {
		aliasedAddr := aliases[addr]

		if !s.IsUnlocked(aliasedAddr) {
			continue
		}

		//log.Info("Attempting to sign root set")

		isMember = true
		signature, err := s.SignHash(accounts.Account{Address: aliasedAddr}, set.hash.Bytes())
		if err != nil {
			log.Error("Failed to sign root set", "err", err)
			continue
		}

		if set.addSignature(aliasedAddr, signature) {
			log.Info("Signed root list", "hash", set.hash.Hex(), "signer", aliasedAddr.Hex(), "root", addr.Hex())
		}
	}

	return isMember
}

func (s *RootManager) signExclusionSet(set *exclusionSet) bool {
	var isSigned bool
	rootsWithAliases := s.getActiveRootSet(true)

	for _, addr := range rootsWithAliases.rootAddresses {
		aliasedAddr := rootsWithAliases.aliases[addr]

		if !s.IsUnlocked(aliasedAddr) {
			continue
		}

		signature, err := s.SignHash(accounts.Account{Address: aliasedAddr}, set.hash.Bytes())
		if err != nil {
			log.Error("Failed to sign exclusion list", "err", err)
			continue
		}

		isSigned = set.addSignature(aliasedAddr, signature)
		if isSigned {
			log.Info("Signed exclusion list", "hash", set.hash.Hex(), "signer", aliasedAddr.Hex(), "root", addr.Hex())
		}
	}

	return isSigned
}

func (s *RootManager) isExclusionSetMeetsQuarantineCriteria(set *exclusionSet, blockNumber uint64) bool {
	rewindLimit := s.maxRewindLimit() //TODO reconsider this logic

	//Check if set is already in the quarantine or can cause a rewind
	return s.isExclusionSetInQuarantine(set) || s.bc.CurrentBlock().Number().Uint64()-blockNumber > rewindLimit
}

// unsafe for concurrent usage
// lock exLock externally first
func (s *RootManager) upgradeExclusionSet(set *exclusionSet, forceUpgrade bool) {
	if s.activeExSet != nil && s.activeExSet.hash == set.hash {
		log.Debug("Exclusion list is already active, skipping", "hash", set.hash.Hex(), "timestamp", set.timestamp)
		log.Info("Exclusion list is already active, skipping", "hash", set.hash.Hex(), "timestamp", set.timestamp)
		return
	}

	// If exclusion set was changed, revalidate blocks up to the earliest affected one
	addrToBlockRange := set.addrToBlockRangeExclusiveDiff(s.activeExSet)

	if len(addrToBlockRange) > 0 {
		earliestBlock := set.earliestBlockFromDiff(s.activeExSet)

		//If the exclusion set contains blocks in the past, and applying it can cause the rewind, quarantine it.
		//Otherwise, revalidate the chain.
		//If upgrade is forced, we don't care about the age of the exclusion set
		//Verification goes only here, because we don't want to quarantine exclusion set if there's no need to
		if !forceUpgrade && s.isExclusionSetMeetsQuarantineCriteria(set, earliestBlock) {
			//Notification about quarantine is sent in startQuarantineRoutine
			if err := s.initiateExclusionSetQuarantine(set); err != nil {
				if err != nil {
					//Quarantine failed, but we still need to return from here to avoid updating exclusion set
					log.Error("Failed to quarantine exclusion set", "err", err)
					return
				}
			}
			log.Info("Exclusion list upgrade failed (!forceUpgrade)", "set", set)
			return
		}

		// Revalidate in separate goroutine to prevent possible deadlocks
		go func() {
			err := s.bc.RevalidateChain(earliestBlock)
			if err != nil {
				log.Error("Failed to revalidate chain after updating the exclusion list", "err", err)
			}
		}()
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

func (s *RootManager) validateOldExclusionSet(set *exclusionSet) error {
	if set == nil {
		return nil
	}

	currentBlock := s.bc.CurrentBlock().Number().Uint64()

	if s.activeExSet != nil {
		// current members of exclusion list should be left unchanged
		for addr, activeBanBlock := range s.activeExSet.addrToBlock {
			newBanBlock, ok := set.addrToBlock[addr]
			if !ok && activeBanBlock <= currentBlock {
				return fmt.Errorf("cannot remove banned validator: %s", addr.String())
			}

			if ok && activeBanBlock > currentBlock && newBanBlock > currentBlock {
				continue
			}

			if ok && newBanBlock != activeBanBlock {
				return fmt.Errorf("cannot change banned validator block: %s", addr.String())
			}
		}
	}

	// new members of exclusion list should be in future
	for addr, newBanBlock := range set.addrToBlock {
		if s.activeExSet != nil {
			_, ok := s.activeExSet.addrToBlock[addr]
			if ok {
				continue
			}
		}

		if newBanBlock <= currentBlock {
			return fmt.Errorf("cannot ban validator in past: %s", addr.String())
		}
	}

	return nil
}

func (s *RootManager) validateExclusionSet(proposedSet *exclusionSet) error {
	if proposedSet == nil {
		return nil
	}

	if !s.isAthosReached() {
		return s.validateOldExclusionSet(proposedSet)
	} else {
		return s.validateNewExclusionSet(proposedSet)
	}
}

func (s *RootManager) validateNewExclusionSet(proposedSet *exclusionSet) error {
	if proposedSet == nil {
		return nil
	}

	currentBlock := s.bc.CurrentBlock().Number()

	if s.activeExSet != nil {
		//Upgraded L0 governance
		for addr, currentBanBlockRanges := range s.activeExSet.blockRanges {
			newBanBlockRanges, newSetContainsRanges := proposedSet.blockRanges[addr]

			if !newSetContainsRanges {
				return fmt.Errorf("proposed exclusion set doesn't contain blocks for address %s", addr.String())
			}

			if len(newBanBlockRanges) == 0 {
				return fmt.Errorf("couldn't find any blocks in proposal for address %s", addr.String())
			}

			//TODO Passed range can be increased& its bad

			for _, exBlockRange := range currentBanBlockRanges {
				inNewSet := false

				for _, newBlockRange := range newBanBlockRanges {
					if exBlockRange.StartsWithTheSameBlock(newBlockRange) {
						if !newBlockRange.IsValid() {
							return fmt.Errorf("specified block range in proposal is invalid: %d - %d for address %s", newBlockRange.StartAddress,
								newBlockRange.EndAddress,
								addr.String())
						}

						if newBlockRange.IsEqualTo(exBlockRange) {
							inNewSet = true
							continue
						}

						//Attempt to reopen closed ban
						if !exBlockRange.EndsInFuture(currentBlock.Uint64()) && newBlockRange.EndsInFuture(currentBlock.Uint64()) {
							return fmt.Errorf("cannot reopen closed ban: %d - %d for address %s",
								exBlockRange.StartAddress,
								exBlockRange.EndAddress,
								addr.String())
						}

						//Attempt to close existing range in past
						if !newBlockRange.EndsInFuture(currentBlock.Uint64()) {
							return fmt.Errorf("cannot close ban in past: %d - %d for address %s",
								newBlockRange.StartAddress,
								newBlockRange.EndAddress,
								addr.String())
						}

						inNewSet = true
					}
				}
				if !inNewSet {
					return fmt.Errorf("active exclusion record for range %d - %d for address  %s doesn't exist in proposed set",
						exBlockRange.StartAddress,
						exBlockRange.EndAddress,
						addr.String())
				}
			}
		}
	}

	for addr, newBanBlockRanges := range proposedSet.blockRanges {
		if len(newBanBlockRanges) == 0 {
			return fmt.Errorf("set should contain at least 1 record")
		}

		for i, newBlockRange := range newBanBlockRanges {
			if !newBlockRange.IsValid() {
				return fmt.Errorf("invalid block addresses in proposal: %d - %d for address %s", newBlockRange.StartAddress,
					newBlockRange.EndAddress,
					addr.String())
			}

			for j, newBlockRangeAlt := range newBanBlockRanges {
				if i != j && newBlockRangeAlt.IntersectsWithRange(newBlockRange) {
					return fmt.Errorf("proposed block range %d - %d intersects with proposed block range  %d - %d for address %s",
						newBlockRange.StartAddress,
						newBlockRange.EndAddress,
						newBlockRangeAlt.StartAddress,
						newBlockRangeAlt.EndAddress,
						addr.String())
				}
			}

			if !newBlockRange.StartsInFuture(currentBlock.Uint64()) && s.activeExSet.blockRanges != nil {
				inValidRanges := false

				for _, exBlockRange := range s.activeExSet.blockRanges[addr] {
					if newBlockRange.StartsWithTheSameBlock(exBlockRange) {
						inValidRanges = true
						break
					}
				}
				if !inValidRanges {
					return fmt.Errorf("cannot add bans in past: %d - %d for address %s",
						newBlockRange.StartAddress,
						newBlockRange.EndAddress,
						addr.String())
				}
			}
		}
	}
	return nil
}

func (s *RootManager) proposeExclusionSet(set *exclusionSet, force bool) (*exclusionSet, error) {
	if !s.isRootNode(true) {
		return nil, errNotRootNode
	}

	exceeded, err := s.isSetQuotaExceeded(set)
	if err != nil {
		log.Warn("Failed to check quota for the proposed list", "err", err)
	}
	if exceeded {
		return nil, errExclusionListQuotaExceeded
	}

	s.exLock.Lock()
	defer s.exLock.Unlock()

	olderThanActive := s.activeExSet != nil && set.timestamp <= s.activeExSet.timestamp
	olderThanDesired := s.desiredExSet != nil && set.timestamp <= s.desiredExSet.timestamp
	if olderThanActive || olderThanDesired {
		return nil, errProposedExclusionListObsolete
	}

	ts := []uint64{uint64(time.Now().Unix())}

	if s.activeExSet != nil {
		ts = append(ts, s.activeExSet.timestamp)
	}
	if s.desiredExSet != nil {
		ts = append(ts, s.desiredExSet.timestamp)
	}
	if s.proposedExSet != nil {
		ts = append(ts, s.proposedExSet.timestamp)
	}
	sort.Slice(ts, func(i int, j int) bool {
		return ts[i] > ts[j]
	})

	if len(strconv.Itoa(int(set.timestamp))) != len(strconv.Itoa(int(ts[0]))) {
		return nil, errInvalidExclusionListTimestamp
	}

	err = s.validateExclusionSet(set)
	if err != nil {
		return nil, err
	}

	if !force {
		if s.proposedExSet != nil && areExclusionListsEqual(set, s.proposedExSet) {
			log.Warn("The proposed list is a duplicate of the current proposed list")
			return set, nil
		}
		if s.activeExSet != nil && areExclusionListsEqual(set, s.activeExSet) {
			log.Warn("The proposed list is a duplicate of the active list")
			return set, nil
		}
	}

	if s.signExclusionSet(set) {
		log.Info("Signed desired exclusion list", "hash", set.hash.Hex())
	}
	s.active.aliases = s.getAliasesOfRoots(s.active.rootAddresses)

	if s.getActiveRootSet(true).isEnoughExSetSignatures(set) {
		s.upgradeExclusionSet(set, false)
	}

	s.proposedExSet = set
	s.db.saveProposedExclusionSet(set)
	err = s.acceptProposedExclusionList(false)
	if err != nil {
		return nil, err
	}

	return set, nil
}

func areExclusionListsEqual(a, b *exclusionSet) bool {
	if a == nil || b == nil {
		return false
	}

	if !reflect.DeepEqual(a.addresses, b.addresses) || !reflect.DeepEqual(a.blockRanges, b.blockRanges) {
		return false
	}

	return true
}

func (s *RootManager) acceptProposedExclusionList(lock bool) error {
	if lock {
		s.exLock.Lock()
		defer s.exLock.Unlock()
	}

	if s.proposedExSet == nil {
		return errProposedExclusionListEmpty
	}

	if s.desiredExSet != nil && s.proposedExSet.timestamp <= s.desiredExSet.timestamp {
		return errProposedExclusionListObsolete
	}

	if s.activeExSet != nil && s.proposedExSet.timestamp <= s.activeExSet.timestamp {
		return errProposedExclusionListObsolete
	}

	err := s.validateExclusionSet(s.proposedExSet)
	if err != nil {
		return err
	}

	if s.signExclusionSet(s.proposedExSet) {
		log.Info("Signed proposed exclusion list", "hash", s.proposedExSet.hash.Hex())
	}

	exSetToSend := s.proposedExSet

	s.active.aliases = s.getAliasesOfRoots(s.active.rootAddresses)

	if s.getActiveRootSet(true).isEnoughExSetSignatures(s.proposedExSet) {
		s.upgradeExclusionSet(s.proposedExSet, false)
	} else {
		s.desiredExSet = s.proposedExSet
		s.db.saveDesiredExclusionSet(s.desiredExSet)
	}

	s.proposedExSet = nil
	s.db.deleteProposedExclusionSet()
	s.desiredExFeed.Send(exSetToSend.copy())

	return nil
}

// unsafe for concurrent calls.
// must be locked before calling
func (s *RootManager) upgradeRootSet(set *rootSet) {
	s.active = set
	s.db.saveActiveRootSet(s.active)

	log.Info("Upgraded root list", "hash", set.hash.Hex(), "timestamp", set.timestamp)

	if s.proposed != nil && s.proposed.timestamp <= set.timestamp {
		log.Info("Dropping obsolete proposed root set", "timestamp", set.timestamp)

		s.proposed = nil
		s.db.deleteProposedRootSet()
	}

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

//nolint:unused
func (s *RootManager) validateRootSet(addresses []common.Address, lock bool) error {
	onChainRootSet := s.getOnchainRootSet(lock)
	diff := s.addressDiff(onChainRootSet.getAddresses(), addresses)

	if len(diff) > 0 {
		return errors.New("Dropping root list that removes on-chain nodes")
	}

	return nil
}

func (s *RootManager) proposeRootSet(set *rootSet, force bool) (*rootSet, error) {
	if !s.isRootNode(false) {
		return nil, errors.New("not a root node")
	}

	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	if set.timestamp <= s.active.timestamp || (s.desired != nil && set.timestamp <= s.desired.timestamp) {
		return nil, errProposedRootListObsolete
	}

	ts := []uint64{uint64(time.Now().Unix())}

	if s.active != nil {
		ts = append(ts, s.active.timestamp)
	}
	if s.desired != nil {
		ts = append(ts, s.desired.timestamp)
	}
	if s.proposed != nil {
		ts = append(ts, s.proposed.timestamp)
	}
	sort.Slice(ts, func(i int, j int) bool {
		return ts[i] > ts[j]
	})

	if len(strconv.Itoa(int(set.timestamp))) != len(strconv.Itoa(int(ts[0]))) {
		return nil, errInvalidExclusionListTimestamp
	}

	if !force {
		if s.proposed != nil && reflect.DeepEqual(set.rootAddresses, s.proposed.rootAddresses) {
			log.Warn("The proposed list is a duplicate of the current proposed list")
			return set, nil
		}
		if s.active != nil && reflect.DeepEqual(set.rootAddresses, s.active.rootAddresses) {
			log.Warn("The proposed list is a duplicate of the active list")
			return set, nil
		}
	}

	if s.signRootSet(set) {
		log.Info("Signed desired root list", "hash", set.hash.Hex())
	}

	s.proposed = set
	s.db.saveProposedRootSet(set)
	err := s.acceptProposedRootList(false)
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (s *RootManager) acceptProposedRootList(lock bool) error {
	if !s.isRootNode(false) {
		return errors.New("not a root node")
	}

	if lock {
		s.rootLock.Lock()
		defer s.rootLock.Unlock()
	}

	if s.proposed == nil {
		return errProposedRootListEmpty
	}

	if s.desired != nil && s.proposed.timestamp <= s.desired.timestamp {
		return errProposedRootListObsolete
	}

	if s.signRootSet(s.proposed) {
		log.Info("Signed proposed root list", "hash", s.proposed.hash.Hex())
	} else {
		log.Info("Failed to sign proposed root list", "list", s.proposed)
	}

	rootSetToSend := s.proposed

	s.proposed.aliases = s.getAliasesOfRoots(s.proposed.rootAddresses)

	if s.active.isAcceptable(s.proposed) {
		s.upgradeRootSet(s.proposed)
	} else {
		s.desired = s.proposed
		s.db.saveDesiredRootSet(s.desired)
	}

	s.proposed = nil
	s.db.deleteProposedRootSet()

	s.desiredRootFeed.Send(rootSetToSend.copy())

	return nil
}

func (s *RootManager) diffRootListByName(nameA, nameB string, lock bool) ([]DiffEntry, error) {
	if lock {
		s.rootLock.Lock()
		defer s.rootLock.Unlock()
	}

	setA, err := s.getRootSetByName(nameA, false)
	if err != nil {
		return nil, err
	}

	setB, err := s.getRootSetByName(nameB, false)
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

func (s *RootManager) getRootSetByName(name string, lock bool) (*rootSet, error) {
	switch strings.ToLower(name) {
	case "active":
		return s.getActiveRootSet(lock), nil
	case "desired":
		return s.getDesiredRootSet(lock), nil
	case "proposed":
		return s.getProposedRootSet(lock), nil
	case "onchain":
		return s.getOnchainRootSet(lock), nil
	}

	return nil, fmt.Errorf("invalid root set name: %s", name)
}

func (s *RootManager) getActiveRootSet(lock bool) *rootSet {
	if lock {
		s.rootLock.Lock()
		defer s.rootLock.Unlock()
	}

	if s.active != nil {
		s.active.aliases = s.getAliasesOfRoots(s.active.rootAddresses)
	}

	return s.active.copy()
}

func (s *RootManager) getDesiredRootSet(lock bool) *rootSet {
	if lock {
		s.rootLock.Lock()
		defer s.rootLock.Unlock()
	}

	if s.desired != nil {
		s.desired.aliases = s.getAliasesOfRoots(s.desired.rootAddresses)
	}

	return s.desired.copy()
}

func (s *RootManager) getProposedRootSet(lock bool) *rootSet {
	if lock {
		s.rootLock.Lock()
		defer s.rootLock.Unlock()
	}

	if s.proposed != nil {
		s.proposed.aliases = s.getAliasesOfRoots(s.proposed.rootAddresses)
	}

	return s.proposed.copy()
}

func (s *RootManager) getOnchainRootSet(lock bool) *rootSet {
	if lock {
		s.rootLock.Lock()
		defer s.rootLock.Unlock()
	}

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
	set.updateAliases(s.getAliasesOfRoots(set.rootAddresses))

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

func (s *RootManager) getActiveApprovalList(blockNumber *big.Int, hash *common.Hash) (*common.RootNodeApprovalList, error) {
	s.approvalLock.Lock()
	defer s.approvalLock.Unlock()

	switch {
	case blockNumber != nil && hash != nil:
		return &common.RootNodeApprovalList{}, errors.New("Block number and hash cannot be specified at the same time")
	case blockNumber != nil:
		if s.bc.GetBlockByNumber(blockNumber.Uint64()) == nil {
			return nil, errors.New("Specified block number doesn't exit")
		}
		return s.getActiveApprovalListByBlockNumber(blockNumber)
	case hash != nil:
		block := s.bc.GetBlockByHash(*hash)
		if block == nil {
			return nil, errors.New("Can't find block by specified hash")
		}
		return s.getActiveApprovalListByBlockNumberAndHash(block.Number(), *hash)
	default:
		return s.db.getLastApprovals().Copy(), nil
	}
}

func (s *RootManager) getActiveApprovalListByBlockNumber(blockNumber *big.Int) (*common.RootNodeApprovalList, error) {
	approvals, err := s.db.getApprovalRecordsByBlockNumber(blockNumber)
	if err != nil {
		return nil, err
	}
	res := &common.RootNodeApprovalList{}
	return res.FillFromArray(approvals), nil
}

func (s *RootManager) getActiveApprovalListByBlockNumberAndHash(blockNumber *big.Int, hash common.Hash) (*common.RootNodeApprovalList, error) {
	approvals, err := s.db.getApprovalRecordsByBlockNumber(blockNumber)
	if err != nil {
		return nil, err
	}
	res := &common.RootNodeApprovalList{}
	var tApprovals []common.RootNodeApproval
	for _, approval := range approvals {
		if approval.Hash == hash {
			tApprovals = append(tApprovals, approval)
		}
	}
	return res.FillFromArray(tApprovals), nil
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

// isAcceptableExclusionSet returns true if there are enough signatures and
// exclusion set is not obsolete
func (s *RootManager) isAcceptableExclusionSet(set *exclusionSet) bool {
	if s.activeExSet != nil && set.timestamp <= s.activeExSet.timestamp {
		return false
	}

	return s.getActiveRootSet(true).isEnoughExSetSignatures(set)
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

func (s *RootManager) IsUnlocked(addr common.Address) bool {
	if _ks := s.manager.Backends(keystore.KeyStoreType); len(_ks) > 0 {
		ks := _ks[0].(*keystore.KeyStore)
		return ks.IsUnlocked(addr)
	}
	return false
}

func (s *RootManager) getAliasesOfRoots(addresses []common.Address) map[common.Address]common.Address {
	res := make(map[common.Address]common.Address)

	if !s.isAthosReached() {
		for _, address := range addresses {
			res[address] = address
		}
		return res
	}

	providerAliases := s.reg.AccountAliases(nil)
	if providerAliases == nil { //signers are set already
		log.Warn("failed to get aliases list from smart contract or smart contract not deployed")
		for _, address := range addresses {
			res[address] = address
		}
		return res
	}
	rnOperationPurpose, _ := new(big.Int).SetString("33a9d3006f267399569cda2996bb19776f92c98b990053176d19c710ed251a5d", 16) //crypto.Keccak256([]byte("ROOT_NODE_OPERATION")

	var purposes []*big.Int
	for range addresses {
		purposes = append(purposes, rnOperationPurpose)
	}

	aliases, errAlias := providerAliases.ResolveBatch(nil, addresses, purposes)
	if errAlias != nil {
		log.Error("failed to get root list account aliases from smart contract", "error", errAlias)
		for _, address := range addresses {
			res[address] = address
		}
		return res
	}

	for i, address := range addresses {
		res[address] = aliases[i]
	}

	return res
}

func (s *RootManager) SignHash(a accounts.Account, hash []byte) ([]byte, error) {
	if _ks := s.manager.Backends(keystore.KeyStoreType); len(_ks) > 0 {
		ks := _ks[0].(*keystore.KeyStore)
		return ks.SignHash(a, hash)
	}
	return nil, errRootManagerCannotSign
}

func (s *RootManager) ValidatePreviousTransitionBlockSignature() {
	if s.bc == nil {
		return
	}
	currentBlock := s.bc.CurrentBlock().Number().Uint64()
	previousTransitionBlock := currentBlock - currentBlock%s.bc.Config().Clique.Epoch
	s.HandleTransitionBlockSignature(s.bc.GetBlockByNumber(previousTransitionBlock).Header())
}

func (s *RootManager) HandleTransitionBlockSignature(header *types.Header) {
	checkLater, err := s.transitionBlockChecker.checkTransitionBlockLater(header)
	if err != nil {
		log.Error("Failed to check transition block", "err", err)
		return
	}
	if checkLater {
		return
	}

	if header == nil {
		log.Debug("Nil header is passed to HandleTransitionBlockSignature")
		return
	}

	s.approvalLock.Lock()
	defer s.approvalLock.Unlock()

	roots := s.active.aliases
	var unlockedRoots []common.Address
	for _, addr := range roots {
		if s.IsUnlocked(addr) {
			unlockedRoots = append(unlockedRoots, addr)
		}
	}
	if len(unlockedRoots) == 0 {
		log.Debug("There is no any unlocked root")
		return
	}
	if s.bc == nil && s.dl == nil {
		log.Debug("Blockchain and downloader are nil in HandleTransitionBlockSignature")
		return
	}

	currentBlock := s.bc.CurrentBlock().Number().Uint64()
	if s.dl != nil && s.dl.Progress().HighestBlock > currentBlock {
		currentBlock = s.dl.Progress().HighestBlock
		log.Debug("Co-sign during sync", "block_to_co-sign", header.Number.Uint64(), "latest_block", currentBlock)
	}

	//No need to sign blocks that are not fresh enough
	if (currentBlock-s.bc.Config().Clique.Epoch) < header.Number.Uint64() && (currentBlock+s.bc.Config().Clique.Epoch) > header.Number.Uint64() {
		log.Info("Handling new transition block", "block number", header.Number.Uint64())

		prevBlockAddress := new(big.Int).SetUint64(header.Number.Uint64() - s.bc.Config().Clique.Epoch)
		if recs, errRecs := s.db.getApprovalRecordsByBlockNumber(prevBlockAddress); errRecs == nil {
			percentage := 100 * len(recs) / len(roots)
			if percentage < approvalsThresholdPercentage {
				log.Warn("Root node approval list contains less than 75% records!")
			}
		}

		for _, addr := range unlockedRoots {
			signature, err := s.SignHash(accounts.Account{Address: addr}, header.Hash().Bytes())
			if err != nil {
				log.Error("Failed to co-sign transition block by root node", "err", err)
				continue
			}

			approval := common.RootNodeApproval{
				BlockNumber: header.Number,
				Hash:        header.Hash(),
				Signature:   signature,
				Signer:      addr,
			}

			if err := s.db.saveApprovalRecord(approval); err != nil {
				log.Error("Failed to save approval of the transition block", "err", err)
			} else {
				resList := common.RootNodeApprovalList{
					BlockNumber: approval.BlockNumber,
					Approvals:   []common.RootNodeApproval{approval},
				}
				s.approvalFeed.Send(&resList)
			}
		}
	}
}

func (s *RootManager) isAthosReached() bool {
	if s.bc == nil {
		return false
	}
	currentBlock := s.bc.CurrentBlock().Number()
	return s.bc.Config().IsAthos(currentBlock)
}

// received exclusion set cam cause huge rewind of the blockchain. It is very undesirable
// instead of updating active exclusion set we will create new one and start quarantining it
func (s *RootManager) initiateExclusionSetQuarantine(set *exclusionSet) error {
	if s.isExclusionSetInQuarantine(set) {
		return nil
	}

	s.quarantineLock.Lock()
	defer s.quarantineLock.Unlock()

	s.quarantineEventCh <- &QuarantineEvent{
		set,
	}

	err := s.db.addExclusionSetToQuarantine(set)
	if err != nil {
		return errors.Wrap(err, "failed to add exclusion set to quarantine")
	}

	return nil
}

func (s *RootManager) notifyExclusionSetIsQuarantined(set *exclusionSet, currentBlock uint64, rewindDepth uint64, rewindLimit uint64) {
	b, err := json.MarshalIndent(set.makeList(), "", "\t")
	if err != nil {
		log.Error("Failed to marshal exclusion set", "err", err)
	}
	msg := "" +
		"There is an exclusion set in the quarantine. Accepting this set will cause huge rewind of the blockchain\n" +
		"		Blockchain will be rewound approx to block #" + fmt.Sprint(rewindDepth) + "). \n" +
		"		Current block is " + fmt.Sprint(currentBlock) + ".\n" +
		"		The Current allowed rewind limit:" + fmt.Sprint(rewindLimit) + ".\n" +
		"		Quarantined exclusion set:\n\n" +
		"" + string(b) + "\n\n" +
		"		If you still want to accept this list and do realize potential harm, execute the following command:\n\n" +
		"\tgov.acceptQuarantinedExclusionList(\"" + set.hash.String() + "\")\n\n"
	sentryMonitor.HandleMessage(msg)
}

func (s *RootManager) startQuarantineRoutine() {
	for {
		if s.bc == nil {
			time.Sleep(5 * time.Second)
			continue
		}
		select {
		case setEvent := <-s.quarantineEventCh:
			if setEvent.Set != nil {
				earliestBlock := setEvent.Set.earliestBlockFromDiff(s.activeExSet)
				s.notifyExclusionSetIsQuarantined(setEvent.Set, s.bc.CurrentBlock().Number().Uint64(), earliestBlock, s.maxRewindLimit())
			}
		case <-time.After(time.Minute):
			s.quarantineLock.Lock()
			sets, err := s.db.getExclusionSetsFromQuarantine()
			if err != nil {
				log.Error("Failed to get exclusion sets from quarantine", "err", err)
			}
			if sets != nil {
				log.Warn("You have exclusion lists in the quarantine. You can see them with command: gov.quarantinedExclusionLists()")
				for i := range sets {
					earliestBlock := sets[i].earliestBlockFromDiff(s.activeExSet)
					s.notifyExclusionSetIsQuarantined(&sets[i], s.bc.CurrentBlock().Number().Uint64(), earliestBlock, s.maxRewindLimit())
				}
			}
			s.quarantineLock.Unlock()
		}
	}
}

func (s *RootManager) isExclusionSetInQuarantine(set *exclusionSet) bool {
	s.quarantineLock.Lock()
	defer s.quarantineLock.Unlock()

	if set == nil {
		return false
	}
	sets, err := s.db.getExclusionSetsFromQuarantine()
	if err != nil {
		log.Error("Failed to get exclusion set from quarantine", "err", err)
		return false
	}
	for _, s := range sets {
		if s.hash == set.hash {
			return true
		}
	}
	return false
}

// TODO rethink this
func (s *RootManager) maxRewindLimit() uint64 {
	return s.rewindLimit
}

func (s *RootManager) setMaxRewindLimit(limit uint64) {
	s.rewindLimit = limit
}

func (s *RootManager) acceptQuarantinedExclusionSet(hash *common.Hash) error {
	s.quarantineLock.Lock()
	defer s.quarantineLock.Unlock()

	set, err := s.db.getQuarantinedExclusionSetByHash(hash)
	if err != nil {
		return errors.Wrap(err, "failed to get exclusion set from quarantine")
	}
	if set == nil {
		return errors.New("exclusion set not found in quarantine")
	}
	log.Warn("Accepting exclusion set from the quarantine", "hash", hash)
	if _, err := s.db.removeExclusionSetFromQuarantine(set); err != nil {
		return errors.Wrap(err, "failed to remove exclusion set from quarantine")
	}
	log.Warn("Exclusion set has been removed from the quarantine", "hash", hash)

	s.upgradeExclusionSet(set, true)
	return nil
}

func (s *RootManager) isSetQuotaExceeded(received interface{}) (bool, error) {
	var (
		prefix         []byte
		hash           common.Hash
		signers        map[common.Address][]byte
		currentEntries map[common.Address][]common.ListQuotaEntry
		resEntries     = make(map[common.Address][]common.ListQuotaEntry)
	)
	switch set := received.(type) {
	case *rootSet:
		hash = set.hash
		signers = set.signers
		prefix = rootQuotaKey
		currentEntries = s.rootQuotaEntries
		s.rootQuotaLock.Lock()
		defer s.rootQuotaLock.Unlock()
	case *exclusionSet:
		hash = set.hash
		signers = set.signers
		prefix = exclusionQuotaKey
		currentEntries = s.exclusionQuotaEntries
		s.exclusionQuotaLock.Lock()
		defer s.exclusionQuotaLock.Unlock()
	default:
		return false, errors.New("unknown set type")
	}

	//First, remove expired entries. No matter who sent them
	//Also, remove entries if they are already in the list and the list was signed by more than one root node
	for address, entries := range currentEntries {
		for _, entry := range entries {
			entryTime := time.Unix(int64(entry.Timestamp), 0)
			if entryTime.Add(s.ProposalQuotaTimeWindow).Before(time.Now()) {
				continue //Record expired, ignore it
			}
			//List can be signed by more than one root node, so if it is already in the DB, we need to remove it
			if len(signers) > 1 && entry.Hash == hash {
				continue //Record is already in the list, ignore it
			}
			//Record is still valid, keep it
			resEntries[address] = append(resEntries[address], entry)
		}
	}

	//Need to save cleaned quotas anyway.
	if !reflect.DeepEqual(currentEntries, resEntries) {
		if err := s.saveQuotas(received, resEntries, prefix); err != nil {
			return false, errors.Wrap(err, "failed to save quota entries")
		}
	}

	//No signatures or more than one signature.
	if len(signers) == 0 || len(signers) > 1 {
		return false, nil
	}

	//We have one signature, so we need to check if it is already in the DB
	//Also, resulting list is cleaned up from expired entries and entries signed by more than one root node
	keys := reflect.ValueOf(signers).MapKeys()
	signer := keys[0].Interface().(common.Address)

	//Check if there's already an entry for this signer or if the quota is exceeded
	if entries, ok := resEntries[signer]; ok {
		if len(entries) >= int(s.ProposalQuotaMax) {
			//Quota exceeded
			return true, nil
		}
		for _, entry := range entries {
			if entry.Hash == hash {
				//Entry already exists
				return false, nil
			}
		}
	}

	//We are here, this means that no entry found and quota is not exceeded. Add new entry
	entry := common.ListQuotaEntry{
		Hash:      hash,
		Timestamp: uint64(time.Now().Unix()),
		Author:    signer,
	}

	resEntries[signer] = append(resEntries[signer], entry)
	if err := s.saveQuotas(received, resEntries, prefix); err != nil {
		return false, errors.Wrap(err, "failed to save quota entries")
	}

	return len(resEntries) > 3, nil
}

// helper function just to avoid code duplication
func (s *RootManager) saveQuotas(received interface{}, currentEntries map[common.Address][]common.ListQuotaEntry, prefix []byte) error {
	if err := s.db.saveQuotaEntries(currentEntries, prefix); err != nil {
		log.Error("Failed to save quota entries", "err", err)
		return err
	}
	switch received.(type) {
	case rootSet, *rootSet:
		s.rootQuotaEntries = currentEntries
	case exclusionSet, *exclusionSet:
		s.exclusionQuotaEntries = currentEntries
	}
	return nil
}
