package governance

import (
	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/params"
)

// defaultDiscardedExclusionHashes lists built-in exclusion-list hashes to ignore per network.
// Keys are wire or calculated hashes; both are stored so gossip can be dropped before parsing.
var defaultDiscardedExclusionHashes = map[uint64][]common.Hash{
	params.MainnetChainConfig.ChainID.Uint64(): {
		// 2022-era list still gossiped on mainnet (calculated hash).
		common.HexToHash("0x61b7caed3930d652bf98942ab4dfaaa5a948736e87de11a013556e5895a09a06"),
		// Stale provided hash from pre-block-range hash rule.
		common.HexToHash("0x37809e789038d75ab188f9e2112495b088e732282c8343741c5902540dd01507"),
	},
}

func (s *RootManager) loadDiscardedExclusionHashes() error {
	hashes, err := s.db.getDiscardedExclusionHashes()
	if err != nil {
		return err
	}
	cache := make(map[common.Hash]struct{}, len(hashes))
	for _, hash := range hashes {
		cache[hash] = struct{}{}
	}
	s.discardedExclusionLock.Lock()
	s.discardedExclusionHashes = cache
	s.discardedExclusionLock.Unlock()
	return nil
}

func (s *RootManager) seedDiscardedExclusionHashes() {
	for _, hash := range defaultDiscardedExclusionHashes[s.networkId] {
		if err := s.addDiscardedExclusionHash(hash); err != nil {
			log.Error("Failed to seed discarded exclusion list hash", "hash", hash.Hex(), "err", err)
		}
	}
	s.clearDiscardedExclusionState()
}

func (s *RootManager) clearDiscardedExclusionState() {
	s.discardedExclusionLock.RLock()
	hashes := make([]common.Hash, 0, len(s.discardedExclusionHashes))
	for hash := range s.discardedExclusionHashes {
		hashes = append(hashes, hash)
	}
	s.discardedExclusionLock.RUnlock()

	for _, hash := range hashes {
		s.quarantineLock.Lock()
		_, err := s.db.removeExclusionSetFromQuarantine(&exclusionSet{hash: hash})
		s.quarantineLock.Unlock()
		if err != nil {
			log.Error("Failed to remove discarded exclusion list from quarantine", "hash", hash.Hex(), "err", err)
		}
	}

	s.exLock.Lock()
	defer s.exLock.Unlock()
	if s.proposedExSet != nil && s.isDiscardedExclusionHash(s.proposedExSet.hash) {
		s.proposedExSet = nil
		s.db.deleteProposedExclusionSet()
	}
}

func (s *RootManager) isDiscardedExclusionHash(hash common.Hash) bool {
	if hash == (common.Hash{}) {
		return false
	}
	s.discardedExclusionLock.RLock()
	defer s.discardedExclusionLock.RUnlock()
	_, ok := s.discardedExclusionHashes[hash]
	return ok
}

func (s *RootManager) addDiscardedExclusionHash(hash common.Hash) error {
	if hash == (common.Hash{}) {
		return errors.New("invalid hash")
	}
	if err := s.db.addDiscardedExclusionHash(hash); err != nil {
		return err
	}
	s.discardedExclusionLock.Lock()
	if s.discardedExclusionHashes == nil {
		s.discardedExclusionHashes = make(map[common.Hash]struct{})
	}
	s.discardedExclusionHashes[hash] = struct{}{}
	s.discardedExclusionLock.Unlock()
	return nil
}

func (s *RootManager) isIgnoredExclusionHash(hash common.Hash) bool {
	if s.isDiscardedExclusionHash(hash) {
		return true
	}
	return s.isExclusionSetInQuarantine(&exclusionSet{hash: hash})
}

// parseExclusionListFromWire parses a gossiped exclusion list unless it is discarded or quarantined.
func (s *RootManager) parseExclusionListFromWire(list *common.ValidatorExclusionList) (*exclusionSet, error) {
	if list == nil {
		return nil, errInvalidExclusionList
	}
	if list.IsEmpty() {
		return nil, nil
	}
	if s.isDiscardedExclusionHash(list.Hash) {
		return nil, nil
	}

	set, err := newExclusionSetForNetwork(list, s.networkId)
	if err != nil {
		return nil, err
	}
	s.logExclusionListHashMismatchIfNeeded(list, set)
	if s.isIgnoredExclusionHash(set.hash) {
		return nil, nil
	}
	return set, nil
}

func (s *RootManager) logExclusionListHashMismatchIfNeeded(list *common.ValidatorExclusionList, set *exclusionSet) {
	if list == nil || set == nil || list.Hash == (common.Hash{}) {
		return
	}
	if list.Hash == set.hash {
		return
	}
	if s.isIgnoredExclusionHash(set.hash) || s.isDiscardedExclusionHash(list.Hash) {
		return
	}
	log.Warn("Exclusion list hash mismatch", "provided", list.Hash.Hex(), "calculated", set.hash.Hex(), "timestamp", set.timestamp)
}

func (s *RootManager) discardExclusionList(hash *common.Hash) error {
	if hash == nil || *hash == (common.Hash{}) {
		return errors.New("invalid hash")
	}

	s.exLock.Lock()
	defer s.exLock.Unlock()

	if s.activeExSet != nil && s.activeExSet.hash == *hash {
		return errors.New("cannot discard active exclusion list")
	}

	if err := s.addDiscardedExclusionHash(*hash); err != nil {
		return errors.Wrap(err, "failed to discard exclusion list")
	}

	s.quarantineLock.Lock()
	if _, err := s.db.removeExclusionSetFromQuarantine(&exclusionSet{hash: *hash}); err != nil {
		s.quarantineLock.Unlock()
		return errors.Wrap(err, "failed to remove exclusion list from quarantine")
	}
	s.quarantineLock.Unlock()

	if s.proposedExSet != nil && s.proposedExSet.hash == *hash {
		s.proposedExSet = nil
		s.db.deleteProposedExclusionSet()
	}

	log.Info("Discarded exclusion list", "hash", hash.Hex())
	return nil
}
