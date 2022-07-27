package governance

import (
	"bytes"
	"gitlab.com/q-dev/q-client/crypto"
	"sync"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/event"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/p2p"
)

// handler of governance protocol messages.
type handler struct {
	rootManager *RootManager

	desiredRootCh  chan *rootSet
	desiredRootSub event.Subscription

	desiredExCh  chan *exclusionSet
	desiredExSub event.Subscription

	rootEventCh     chan *rootSetEvent
	exEventCh       chan *exclusionSetEvent
	approvalEventCh chan *approvalEvent
	done            chan struct{}

	approvalCh  chan *common.RootNodeApprovalList
	approvalSub event.Subscription

	peerWG sync.WaitGroup
	peers  *peerSet
}

func newHandler(rootManager *RootManager) *handler {
	desiredRootCh := make(chan *rootSet)
	desiredRootSub := rootManager.desiredRootFeed.Subscribe(desiredRootCh)

	desiredExCh := make(chan *exclusionSet)
	desiredExSub := rootManager.desiredExFeed.Subscribe(desiredExCh)

	approvalCh := make(chan *common.RootNodeApprovalList)
	approvalSub := rootManager.approvalFeed.Subscribe(approvalCh)

	return &handler{
		rootManager: rootManager,

		desiredRootCh:  desiredRootCh,
		desiredRootSub: desiredRootSub,

		desiredExCh:  desiredExCh,
		desiredExSub: desiredExSub,

		rootEventCh:     make(chan *rootSetEvent),
		exEventCh:       make(chan *exclusionSetEvent),
		approvalEventCh: make(chan *approvalEvent),
		done:            make(chan struct{}),

		approvalCh:  approvalCh,
		approvalSub: approvalSub,

		peerWG: sync.WaitGroup{},
		peers:  newPeerSet(),
	}
}

func (h *handler) run() {
	go h.listenForDesiredRootList()
	go h.broadcastRootSets()

	go h.listenForDesiredExclusionList()
	go h.broadcastExclusionSets()

	go h.listenRNApprovals()
	go h.broadcastApprovals()
}

type rootSetEvent struct {
	fromID string
	set    *rootSet
}

type exclusionSetEvent struct {
	fromID string
	set    *exclusionSet
}

type approvalEvent struct {
	fromID   string
	approval *common.RootNodeApprovalList
}

func (h *handler) listenForDesiredRootList() {
	for {
		select {
		case set := <-h.desiredRootCh:
			for _, p := range h.peers.all() {
				log.Debug("Sending desired root set", "to", p.id)
				p.asyncSendRootList(set)
			}
		case err := <-h.desiredRootSub.Err():
			log.Debug("no new desired roots sets", "err", err)
			return
		case <-h.done:
			return
		}
	}
}

func (h *handler) broadcastRootSets() {
	for {
		select {
		case msg := <-h.rootEventCh:
			for _, p := range h.peers.all() {
				if msg.fromID == p.id {
					continue
				}

				p.asyncSendRootList(msg.set)
			}
		case <-h.done:
			return
		}
	}
}

func (h *handler) listenForDesiredExclusionList() {
	for {
		select {
		case set := <-h.desiredExCh:
			for _, p := range h.peers.all() {
				log.Debug("Sending desired exclusion set", "to", p.id)
				p.asyncSendExclusionList(set)
			}
		case err := <-h.desiredExSub.Err():
			log.Debug("no new desired exclusion sets", "err", err)
			return
		case <-h.done:
			return
		}
	}
}

func (h *handler) broadcastExclusionSets() {
	for {
		select {
		case msg := <-h.exEventCh:
			for _, p := range h.peers.all() {
				if msg.fromID == p.id {
					continue
				}

				p.asyncSendExclusionList(msg.set)
			}
		case <-h.done:
			return
		}
	}
}

func (h *handler) listenRNApprovals() {
	for {
		select {
		case approval := <-h.approvalCh:
			for _, p := range h.peers.all() {
				if p.version < qgov3 {
					continue
				}
				log.Debug("Sending approval list root node approvals", "to", p.id)
				p.asyncSendApprovals(approval)
			}
		case err := <-h.approvalSub.Err():
			log.Debug("no new approvals", "err", err)
			return
		case <-h.done:
			return
		}
	}
}

func (h *handler) broadcastApprovals() {
	for {
		select {
		case msg := <-h.approvalEventCh:
			for _, p := range h.peers.all() {
				if p.version < qgov3 {
					continue
				}
				if msg.fromID == p.id {
					continue
				}

				p.asyncSendApprovals(msg.approval)
			}
		case <-h.done:
			return
		}
	}
}

func (h *handler) stop() {
	h.peers.close()
	h.peerWG.Wait()

	close(h.done)
}

func (h *handler) makeProtocol(version uint) p2p.Protocol {
	length, ok := protocolLengths[version]
	if !ok {
		panic("makeProtocol for unknown version")
	}

	return p2p.Protocol{
		Name:    protocolName,
		Version: version,
		Length:  length,
		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			return h.runPeer(newPeer(int(version), p, rw))
		},
		NodeInfo: func() interface{} {
			active := h.rootManager.getActiveRootSet(true)
			return struct {
				Timestamp uint64
			}{
				Timestamp: active.timestamp,
			}
		},
	}
}

func (h *handler) makeStatusBody(rm *RootManager) statusMsgBody {
	rm.rootLock.Lock()
	defer rm.rootLock.Unlock()

	rm.exLock.Lock()
	defer rm.exLock.Unlock()

	return statusMsgBody{
		CurrentRootList:  rm.active.copy().makeList(),
		DesiredRootList:  rm.desired.copy().makeList(),
		ProposedRootList: rm.proposed.copy().makeList(),

		CurrentExclusionList:  rm.activeExSet.copy().makeList(),
		DesiredExclusionList:  rm.desiredExSet.copy().makeList(),
		ProposedExclusionList: rm.proposedExSet.copy().makeList(),

		Network: rm.networkId,
	}
}

func (h *handler) runPeer(p *peer) error {
	rm := h.rootManager

	statusBody := h.makeStatusBody(rm)
	status, err := p.handshake(statusBody)
	if err != nil {
		return err
	}
	if status == nil {
		return errors.New("status is nil")
	}

	// Propagate current root set to newly connected peers
	h.propagateRootSet(status.desiredRootSet)
	h.propagateRootSet(status.proposedRootSet)
	h.propagateRootSet(status.currentRootSet)

	// Propagate current exclusion set to newly connected peers
	h.propagateExclusionSet(status.desiredExSet)
	h.propagateExclusionSet(status.proposedExSet)
	h.propagateExclusionSet(status.currentExSet)

	if p.version >= qgov3 {
		h.propagateApprovals(rm.db.getLastApprovals())
	}

	h.peers.register(p)
	defer h.peers.unregister(p)

	if status != nil {
		for _, set := range []*rootSet{status.currentRootSet, status.desiredRootSet, status.proposedRootSet} {
			if set == nil {
				continue
			}

			if err := h.handleRootSet(p, set); err != nil {
				return err
			}
		}

		for _, set := range []*exclusionSet{status.currentExSet, status.desiredExSet, status.proposedExSet} {
			if set == nil {
				continue
			}

			if err := h.handleExclusionSet(p, set); err != nil {
				return err
			}
		}

	}

	h.peerWG.Add(1)
	defer h.peerWG.Done()

	for {
		if err := h.handleMsg(p); err != nil {
			p.Log().Error("Governance message handling failed", "err", err)
			return err
		}
	}
}

func (h *handler) propagateRootSet(set *rootSet) {
	if set != nil {
		h.rootEventCh <- &rootSetEvent{set: set}
	}
}

func (h *handler) propagateExclusionSet(set *exclusionSet) {
	if set != nil {
		h.exEventCh <- &exclusionSetEvent{set: set}
	}
}

func (h *handler) propagateApprovals(approvals *common.RootNodeApprovalList) {
	if approvals != nil {
		h.approvalEventCh <- &approvalEvent{approval: approvals}
	}
}

func (h *handler) handleMsg(p *peer) error {
	msg, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}

	if msg.Size > protocolMaxMsgSize {
		return errors.Wrap(err, "message too large")
	}

	defer func() { _ = msg.Discard() }()

	switch msg.Code {
	case StatusMsg:
		return errors.New("unexpected msg code")
	case RootListMsg:
		return h.handleRootListMsg(p, msg)
	case ExclusionListMsg:
		return h.handleExclusionList(p, msg)
	case ApprovalMsg:
		return h.handleApprovalMsg(p, msg)
	default:
		return errors.New("unknown msg code")
	}
}

func (h *handler) handleRootListMsg(p *peer, msg p2p.Msg) error {
	var list common.RootList
	if err := msg.Decode(&list); err != nil {
		return err
	}

	received, err := newRootSet(&list)
	if err != nil {
		return err
	}

	return h.handleRootSet(p, received)
}

func (h *handler) handleApprovalMsg(p *peer, msg p2p.Msg) error {
	var approval common.RootNodeApprovalList
	if err := msg.Decode(&approval); err != nil {
		return err
	}

	received := approval
	return h.handleIncomingApproval(p, &received)
}

func (h *handler) handleRootSet(p *peer, received *rootSet) error {
	rm := h.rootManager
	rm.rootLock.Lock()
	defer rm.rootLock.Unlock()

	switch {
	case rm.active.isAcceptable(received) && (rm.desired == nil || rm.desired.hash != received.hash):
		if rm.isMember(received.rootAddresses) {
			rm.signRootSet(received)
		}

		rm.upgradeRootSet(received)
		h.rootEventCh <- &rootSetEvent{set: received}
	case rm.active.hash == received.hash:
		newSignatures := rm.active.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		log.Debug("Received active root list signatures", "from", p.id, "signers", toSigners(newSignatures))
		h.rootEventCh <- &rootSetEvent{fromID: p.id, set: rm.active.copy()}

		rm.db.saveActiveRootSet(rm.active)
	case rm.desired != nil && rm.desired.hash == received.hash:
		newSignatures := rm.desired.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		log.Debug("Received desired root list signatures", "from", p.id, "signers", toSigners(newSignatures))
		if !rm.active.isAcceptable(rm.desired) {
			h.rootEventCh <- &rootSetEvent{fromID: p.id, set: rm.desired.copy()}
			rm.db.saveDesiredRootSet(rm.desired)
			return nil
		}

		rm.upgradeRootSet(rm.desired)
		h.rootEventCh <- &rootSetEvent{set: rm.active.copy()}
	case rm.proposed != nil && rm.proposed.hash == received.hash:
		newSignatures := rm.proposed.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		log.Debug("Received proposed root list signatures", "from", p.id, "signers", toSigners(newSignatures))
		if !rm.active.isAcceptable(rm.proposed) {
			h.rootEventCh <- &rootSetEvent{fromID: p.id, set: rm.proposed.copy()}
			rm.db.saveProposedRootSet(rm.proposed)
			return nil
		}

		rm.upgradeRootSet(rm.proposed)
		h.rootEventCh <- &rootSetEvent{set: rm.active.copy()}
	default:
		signers := rm.active.knownSigners(received.signers)
		if len(signers) == 0 {
			log.Debug("Ignoring proposed root list: not signed by any known root node",
				"hash", received.hash.Hex(),
				"from", p.id)
			return nil
		}

		if rm.proposed != nil && rm.proposed.timestamp >= received.timestamp {
			log.Debug("Ignoring proposed root list: obsolete", "timestamp", received.timestamp, "hash", received.hash.Hex())
			return nil
		}

		rm.proposed = received
		rm.db.saveProposedRootSet(rm.proposed)
		h.rootEventCh <- &rootSetEvent{fromID: p.id, set: received}
		log.Warn("Received a new proposed root list", "hash", received.hash.Hex(), "timestamp", received.timestamp)
	}

	return nil
}

func (h *handler) handleExclusionList(p *peer, msg p2p.Msg) error {
	var list common.ValidatorExclusionList
	if err := msg.Decode(&list); err != nil {
		return err
	}

	received, err := newExclusionSet(&list)
	if err != nil {
		return err
	}

	return h.handleExclusionSet(p, received)
}

func (h *handler) handleExclusionSet(p *peer, received *exclusionSet) error {
	rm := h.rootManager
	rm.exLock.Lock()
	defer rm.exLock.Unlock()

	switch true {
	case rm.isAcceptableExclusionSet(received):
		// ensure locally stored signatures are not lost
		if rm.desiredExSet != nil && rm.desiredExSet.hash == received.hash {
			received.mergeSignatures(rm.desiredExSet.hash, rm.desiredExSet.signers)
		}

		if rm.isRootNode() {
			rm.signExclusionSet(received)
		}

		rm.upgradeExclusionSet(received)
		h.exEventCh <- &exclusionSetEvent{set: received}
	case rm.activeExSet != nil && rm.activeExSet.hash == received.hash:
		newSignatures := rm.activeExSet.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		log.Debug("Received new exclusion list signatures", "from", p.id, "singers", toSigners(newSignatures))
		h.exEventCh <- &exclusionSetEvent{fromID: p.id, set: rm.activeExSet.copy()}
		rm.db.saveActiveExclusionSet(rm.activeExSet)
	case rm.desiredExSet != nil && rm.desiredExSet.hash == received.hash:
		newSignatures := rm.desiredExSet.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		log.Debug("Received new desired exclusion list signatures", "from", p.id, "singers", toSigners(newSignatures))
		if !rm.getActiveRootSet(true).isEnoughExSetSignatures(rm.desiredExSet) {
			h.exEventCh <- &exclusionSetEvent{fromID: p.id, set: rm.desiredExSet}
			rm.db.saveDesiredExclusionSet(rm.desiredExSet)
			return nil
		}

		rm.upgradeExclusionSet(rm.desiredExSet)
		h.exEventCh <- &exclusionSetEvent{set: rm.activeExSet}
	default:
		if len(rm.getActiveRootSet(true).knownSigners(received.signers)) == 0 {
			log.Debug("Ignoring proposed exclusion list: no active root node signatures")
			return nil
		}

		obsoleteByActive := rm.activeExSet != nil && received.timestamp <= rm.activeExSet.timestamp
		obsoleteByProposed := rm.proposedExSet != nil && received.timestamp <= rm.proposedExSet.timestamp
		if obsoleteByActive || obsoleteByProposed {
			log.Debug("Ignoring proposed exclusion list: obsolete")
			return nil
		}

		rm.proposedExSet = received
		rm.db.saveProposedExclusionSet(rm.proposedExSet)

		log.Warn("Received a new proposed exclusion list", "hash", received.hash.Hex(), "timestamp", received.timestamp)
	}

	return nil
}

func (h *handler) handleIncomingApproval(p *peer, received *common.RootNodeApprovalList) error {
	rm := h.rootManager
	rm.approvalLock.Lock()
	defer rm.approvalLock.Unlock()

	if received == nil {
		return errProposedApprovalListEmpty
	}

	if received.BlockNumber.Uint64()%rm.bc.Config().Clique.Epoch != 0 {
		log.Error("Received root node approval list contains invalid block number", "blockNumber", received.BlockNumber)
		return errInvalidApprovalBlockNumber
	}

	exApprovals, errEx := rm.db.getApprovalRecordsByBlockNumber(received.BlockNumber)
	if errEx != nil {
		log.Warn(errEx.Error()) //todo
	}

	for _, approval := range received.Approvals {

		pubkey, err := crypto.SigToPub(approval.Hash.Bytes(), approval.Signature)
		if err != nil {
			return errInvalidSignature
		}

		addr := crypto.PubkeyToAddress(*pubkey)
		if _, ok := rm.active.roots[addr]; !ok {
			log.Warn("Received root node approval contains non-root signature", "addr", addr, "blockNumber", received.BlockNumber)
			continue
		}

		for _, signature := range exApprovals {
			if bytes.Equal(approval.Signature, signature.Signature) {
				continue
			}
		}

		if errSave := rm.db.saveApprovalRecord(approval, false); errSave != nil {
			//TODO
		}

	}

	h.approvalEventCh <- &approvalEvent{fromID: p.id, approval: received}

	return nil
}

func toSigners(sigs map[common.Address][]byte) []string {
	var strs []string
	for addr := range sigs {
		strs = append(strs, addr.Hex())
	}

	return strs
}
