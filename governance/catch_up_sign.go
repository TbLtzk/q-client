package governance

import (
	"time"

	"gitlab.com/q-dev/q-client/accounts"
	"gitlab.com/q-dev/q-client/log"
)

const catchUpSignTickerInterval = 60 * time.Second

func (s *RootManager) catchUpSignRootSet(set *rootSet) bool {
	if s == nil || set == nil || s.active == nil {
		return false
	}

	var added bool
	aliases := s.getAliasesOfRoots(s.active.rootAddresses)

	for _, addr := range s.active.rootAddresses {
		aliasedAddr := aliases[addr]

		if !s.IsUnlocked(aliasedAddr) {
			continue
		}
		if _, ok := set.signers[aliasedAddr]; ok {
			continue
		}

		signature, err := s.SignHash(accounts.Account{Address: aliasedAddr}, set.hash.Bytes())
		if err != nil {
			log.Error("Failed to catch-up sign active root list", "err", err)
			continue
		}

		if set.addSignature(aliasedAddr, signature) {
			log.Info("Signed active root list (catch-up)", "hash", set.hash.Hex(), "signer", aliasedAddr.Hex(), "root", addr.Hex())
			added = true
		}
	}

	return added
}

func (s *RootManager) catchUpSignActiveExclusionSet(set *exclusionSet) bool {
	if s == nil || set == nil {
		return false
	}

	rootsWithAliases := s.getActiveRootSet(true)
	var added bool

	for _, addr := range rootsWithAliases.rootAddresses {
		aliasedAddr := rootsWithAliases.aliases[addr]

		if !s.IsUnlocked(aliasedAddr) {
			continue
		}
		if _, ok := set.signers[aliasedAddr]; ok {
			continue
		}

		signature, err := s.SignHash(accounts.Account{Address: aliasedAddr}, set.hash.Bytes())
		if err != nil {
			log.Error("Failed to catch-up sign active exclusion list", "err", err)
			continue
		}

		if set.addSignature(aliasedAddr, signature) {
			log.Info("Signed active exclusion list (catch-up)", "hash", set.hash.Hex(), "signer", aliasedAddr.Hex(), "root", addr.Hex())
			added = true
		}
	}

	return added
}

func (h *handler) catchUpSignActiveLists() {
	rm := h.rootManager

	var rootSet *rootSet
	rm.rootLock.Lock()
	rm.refreshActiveAliases()
	if rm.active != nil && rm.isRootNode(false) {
		if rm.catchUpSignRootSet(rm.active) {
			rm.db.saveActiveRootSet(rm.active)
			rootSet = rm.active.copy()
		}
	}
	rm.rootLock.Unlock()
	if rootSet != nil {
		h.propagateRootSet(rootSet)
	}

	var exSet *exclusionSet
	rm.exLock.Lock()
	if rm.activeExSet != nil && rm.isRootNode(false) {
		if rm.catchUpSignActiveExclusionSet(rm.activeExSet) {
			rm.db.saveActiveExclusionSet(rm.activeExSet)
			exSet = rm.activeExSet.copy()
		}
	}
	rm.exLock.Unlock()
	if exSet != nil {
		h.propagateExclusionSet(exSet)
	}
}

func (h *handler) catchUpSignActiveListsTicker() {
	ticker := time.NewTicker(catchUpSignTickerInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			h.catchUpSignActiveLists()
		case <-h.done:
			return
		}
	}
}
