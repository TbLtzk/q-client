package governance

import (
	"sync"

	"gitlab.com/q-dev/q-client/accounts"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/p2p"
)

const maxTypedRelayDedupEntries = 10000

type typedRelayDedup struct {
	mu    sync.Mutex
	seen  map[typedRelayKey]struct{}
	order []typedRelayKey
}

type typedRelayKey struct {
	proposalHash common.Hash
	signer       common.Address
	kind         uint8 // 0 root, 1 exclusion
}

func newTypedRelayDedup() *typedRelayDedup {
	return &typedRelayDedup{seen: make(map[typedRelayKey]struct{})}
}

func (d *typedRelayDedup) remember(key typedRelayKey) bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.seen[key]; ok {
		return false
	}
	d.seen[key] = struct{}{}
	d.order = append(d.order, key)
	if len(d.order) > maxTypedRelayDedupEntries {
		evict := d.order[0]
		d.order = d.order[1:]
		delete(d.seen, evict)
	}
	return true
}

type typedRootRelayEvent struct {
	fromID string
	list   common.RootList
}

type typedExclusionRelayEvent struct {
	fromID string
	list   common.ValidatorExclusionList
}

func (h *handler) broadcastTypedRootRelays() {
	for {
		select {
		case msg := <-h.typedRootRelayCh:
			for _, p := range h.peers.all() {
				if msg.fromID == p.id || p.version < qgov6 {
					continue
				}
				if err := p.sendTypedRootListRelay(&msg.list); err != nil {
					p.Log().Warn("failed to send typed root list relay", "err", err)
				}
			}
		case <-h.done:
			return
		}
	}
}

func (h *handler) broadcastTypedExclusionRelays() {
	for {
		select {
		case msg := <-h.typedExclusionRelayCh:
			for _, p := range h.peers.all() {
				if msg.fromID == p.id || p.version < qgov6 {
					continue
				}
				if err := p.sendTypedExclusionListRelay(&msg.list); err != nil {
					p.Log().Warn("failed to send typed exclusion list relay", "err", err)
				}
			}
		case <-h.done:
			return
		}
	}
}

func (h *handler) ProcessTypedSignedRootList(list common.RootList) (common.Hash, error) {
	return h.processTypedRootListAttestation(list, "")
}

func (h *handler) ProcessTypedSignedExclusionList(list common.ValidatorExclusionList) (common.Hash, error) {
	return h.processTypedExclusionListAttestation(list, "")
}

func (h *handler) processTypedRootListAttestation(list common.RootList, fromID string) (common.Hash, error) {
	if err := h.rootManager.validatePublicSubmissionAllowed(list); err != nil {
		return common.Hash{}, err
	}

	set, err := h.rootManager.newTypedSignedRootSet(list)
	if err != nil {
		return common.Hash{}, err
	}

	h.selfBridgeTypedRootList(set)
	h.relayTypedRootList(list, fromID, set.signers)

	return set.hash, nil
}

func (h *handler) processTypedExclusionListAttestation(list common.ValidatorExclusionList, fromID string) (common.Hash, error) {
	if err := h.rootManager.validatePublicSubmissionAllowed(list); err != nil {
		return common.Hash{}, err
	}

	set, err := h.rootManager.newTypedSignedExclusionSet(list)
	if err != nil {
		return common.Hash{}, err
	}

	h.selfBridgeTypedExclusionList(set)
	h.relayTypedExclusionList(list, fromID, set.signers)

	return set.hash, nil
}

func (h *handler) relayTypedRootList(list common.RootList, fromID string, signers map[common.Address][]byte) {
	var relay bool
	for signer := range signers {
		key := typedRelayKey{proposalHash: list.Hash, signer: signer, kind: 0}
		if h.typedRelayDedup.remember(key) {
			relay = true
		}
	}
	if !relay {
		return
	}
	h.typedRootRelayCh <- &typedRootRelayEvent{fromID: fromID, list: list}
}

func (h *handler) relayTypedExclusionList(list common.ValidatorExclusionList, fromID string, signers map[common.Address][]byte) {
	var relay bool
	for signer := range signers {
		key := typedRelayKey{proposalHash: list.Hash, signer: signer, kind: 1}
		if h.typedRelayDedup.remember(key) {
			relay = true
		}
	}
	if !relay {
		return
	}
	h.typedExclusionRelayCh <- &typedExclusionRelayEvent{fromID: fromID, list: list}
}

func (h *handler) handleTypedRootListRelay(p *peer, msg p2p.Msg) error {
	if p.version < qgov6 {
		return nil
	}
	var list common.RootList
	if err := msg.Decode(&list); err != nil {
		return err
	}
	if _, err := h.processTypedRootListAttestation(list, p.id); err != nil {
		p.Log().Debug("Ignoring invalid typed root list relay", "err", err)
	}
	return nil
}

func (h *handler) handleTypedExclusionListRelay(p *peer, msg p2p.Msg) error {
	if p.version < qgov6 {
		return nil
	}
	var list common.ValidatorExclusionList
	if err := msg.Decode(&list); err != nil {
		return err
	}
	if _, err := h.processTypedExclusionListAttestation(list, p.id); err != nil {
		p.Log().Debug("Ignoring invalid typed exclusion list relay", "err", err)
	}
	return nil
}

func (rm *RootManager) localGovernanceSigningAddresses() []common.Address {
	active := rm.getActiveRootSet(true)
	if active == nil {
		return nil
	}

	var addrs []common.Address
	for _, root := range active.rootAddresses {
		alias := active.aliases[root]
		if rm.IsUnlocked(alias) {
			addrs = append(addrs, alias)
		}
	}
	return addrs
}

func (h *handler) selfBridgeTypedRootList(typed *rootSet) {
	localAddrs := h.rootManager.localGovernanceSigningAddresses()
	if len(localAddrs) == 0 {
		return
	}

	for _, signer := range localAddrs {
		if _, ok := typed.signers[signer]; !ok {
			continue
		}
		if h.rootSetHasRawSignature(typed.hash, signer) {
			continue
		}
		if err := h.mintAndImportRawRootSignature(typed, signer); err != nil {
			log.Error("Failed to self-bridge typed root list attestation", "err", err, "hash", typed.hash.Hex(), "signer", signer.Hex())
		}
	}
}

func (h *handler) selfBridgeTypedExclusionList(typed *exclusionSet) {
	localAddrs := h.rootManager.localGovernanceSigningAddresses()
	if len(localAddrs) == 0 {
		return
	}

	for _, signer := range localAddrs {
		if _, ok := typed.signers[signer]; !ok {
			continue
		}
		if h.exclusionSetHasRawSignature(typed.hash, signer) {
			continue
		}
		if err := h.mintAndImportRawExclusionSignature(typed, signer); err != nil {
			log.Error("Failed to self-bridge typed exclusion list attestation", "err", err, "hash", typed.hash.Hex(), "signer", signer.Hex())
		}
	}
}

func (h *handler) rootSetHasRawSignature(hash common.Hash, signer common.Address) bool {
	rm := h.rootManager
	rm.rootLock.Lock()
	defer rm.rootLock.Unlock()

	for _, set := range []*rootSet{rm.active, rm.desired, rm.proposed} {
		if set == nil || set.hash != hash {
			continue
		}
		if set.hasRawSignatureFrom(signer) {
			return true
		}
	}
	return false
}

func (h *handler) exclusionSetHasRawSignature(hash common.Hash, signer common.Address) bool {
	rm := h.rootManager
	rm.exLock.Lock()
	defer rm.exLock.Unlock()

	for _, set := range []*exclusionSet{rm.activeExSet, rm.desiredExSet, rm.proposedExSet} {
		if set == nil || set.hash != hash {
			continue
		}
		if set.hasRawSignatureFrom(signer) {
			return true
		}
	}
	return false
}

func (s *exclusionSet) hasRawSignatureFrom(signer common.Address) bool {
	sig, ok := s.signers[signer]
	if !ok {
		return false
	}
	_, err := crypto.SigToPub(s.hash.Bytes(), normalizeECDSASignatureV(sig))
	return err == nil
}

func (h *handler) mintAndImportRawRootSignature(proposal *rootSet, signer common.Address) error {
	sig, err := h.rootManager.SignHash(accounts.Account{Address: signer}, proposal.hash.Bytes())
	if err != nil {
		return err
	}

	list := proposal.makeList()
	list.Signatures = common.HexSignaturesFromBytes([][]byte{sig})
	return h.importRootListFrom("", &list, false)
}

func (h *handler) mintAndImportRawExclusionSignature(proposal *exclusionSet, signer common.Address) error {
	sig, err := h.rootManager.SignHash(accounts.Account{Address: signer}, proposal.hash.Bytes())
	if err != nil {
		return err
	}

	list := proposal.makeList()
	list.Signatures = common.HexSignaturesFromBytes([][]byte{sig})
	return h.importExclusionListFrom("", &list, false)
}
