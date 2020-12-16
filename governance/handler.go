package governance

import (
	"sync"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/event"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/p2p"
)

// handler of protocol messages.
type handler struct {
	rootsMgr *RootManager

	targetRootListCh     chan *rootSet
	newTargetRootListSub event.Subscription

	targetExListCh  chan *exclusionSet
	targetExListSub event.Subscription

	rootListCh chan *rootSetEvent
	exListCh   chan *exclusionSetEvent
	done       chan struct{}

	peerWG sync.WaitGroup
	peers  *peerSet
}

func newHandler(roots *RootManager) *handler {
	newRoots := make(chan *rootSet)
	newRootsSub := roots.targetRootFeed.Subscribe(newRoots)

	targetEx := make(chan *exclusionSet)
	targetExSub := roots.targetExListFeed.Subscribe(targetEx)

	return &handler{
		rootsMgr:             roots,
		targetRootListCh:     newRoots,
		newTargetRootListSub: newRootsSub,
		targetExListCh:       targetEx,
		targetExListSub:      targetExSub,
		rootListCh:           make(chan *rootSetEvent),
		exListCh:             make(chan *exclusionSetEvent),
		done:                 make(chan struct{}),
		peers:                newPeerSet(),
	}
}

func (h *handler) run() {
	go h.listenForTargetRootList()
	go h.broadcastRootSets()

	go h.listenForTargetExclusionList()
	go h.broadcastExclusionSets()
}

type rootSetEvent struct {
	fromID string
	set    *rootSet
}

type exclusionSetEvent struct {
	fromID string
	set    *exclusionSet
}

func (h *handler) listenForTargetRootList() {
	for {
		select {
		case set := <-h.targetRootListCh:
			for _, p := range h.peers.all() {
				log.Debug("Sending target set", "to", p.id)
				p.asyncSendRootList(set)
			}
		case err := <-h.newTargetRootListSub.Err():
			log.Debug("no new target roots sets", "err", err)
			return
		case <-h.done:
			return
		}
	}
}

func (h *handler) broadcastRootSets() {
	for {
		select {
		case msg := <-h.rootListCh:
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

func (h *handler) listenForTargetExclusionList() {
	for {
		select {
		case set := <-h.targetExListCh:
			for _, p := range h.peers.all() {
				log.Debug("SENDING exclusion set", "to", p.id)
				p.asyncSendExclusionList(set)
			}
		case err := <-h.targetExListSub.Err():
			log.Debug("no new target exclusion sets", "err", err)
			return
		case <-h.done:
			return
		}
	}
}

func (h *handler) broadcastExclusionSets() {
	for {
		select {
		case msg := <-h.exListCh:
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
			return h.runPeer(newPeer(p, rw))
		},
		NodeInfo: func() interface{} {
			current := h.rootsMgr.currentRootSet()
			return struct {
				Timestamp uint64
			}{
				Timestamp: current.timestamp,
			}
		},
	}
}

func (h *handler) runPeer(p *peer) error {
	rm := h.rootsMgr
	// lock in order to avoid sending
	// partially upgraded state
	rm.lock.Lock()
	rm.exLock.Lock()
	status, err := p.handshake(statusMsgBody{
		CurrentRootList:       rm.current.copy().makeList(),
		DesiredRootList:       rm.target.copy().makeList(),
		ProposedRootList:      rm.proposed.copy().makeList(),
		CurrentExclusionList:  rm.exclusionSet.copy().makeList(),
		DesiredExclusionList:  rm.targetExSet.copy().makeList(),
		ProposedExclusionList: rm.proposedExSet.copy().makeList(),
	})
	rm.lock.Unlock()
	rm.exLock.Unlock()

	if err != nil {
		return err
	}

	h.peers.register(p)
	defer h.peers.unregister(p)

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

	h.peerWG.Add(1)
	defer h.peerWG.Done()

	for {
		if err := h.handleMsg(p); err != nil {
			p.Log().Debug("Governance message handling failed", "err", err)
			return err
		}
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

func (h *handler) handleRootSet(p *peer, received *rootSet) error {
	rm := h.rootsMgr
	rm.lock.Lock()
	defer rm.lock.Unlock()

	switch {
	case rm.current.isAcceptable(received) && (rm.target == nil || rm.target.hash != received.hash):
		if rm.isMember(received.rootAddresses) {
			rm.signRootSet(received)
		}

		rm.upgradeRootSet(received)
		h.rootListCh <- &rootSetEvent{set: received}
	case rm.current.hash == received.hash:
		newSignatures := rm.current.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		log.Debug("Received current root list signatures", "from", p.id, "signers", toSigners(newSignatures))
		h.rootListCh <- &rootSetEvent{fromID: p.id, set: rm.current.copy()}

		rm.db.saveCurrentRootSet(rm.current)
	case rm.target != nil && rm.target.hash == received.hash:
		newSignatures := rm.target.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		log.Debug("Received desired root list signatures", "from", p.id, "signers", toSigners(newSignatures))
		if !rm.current.isAcceptable(rm.target) {
			h.rootListCh <- &rootSetEvent{fromID: p.id, set: rm.target.copy()}
			rm.db.saveDesiredRootSet(rm.target)
			return nil
		}

		rm.upgradeRootSet(rm.target)
		h.rootListCh <- &rootSetEvent{set: rm.current.copy()}
	default:
		if !rm.isMember(received.rootAddresses) {
			log.Debug("Ignoring proposed root list: not a member of the new list", "hash", received.hash.Hex())
			return nil
		}

		signers := rm.current.knownSigners(received.signers)
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
		h.rootListCh <- &rootSetEvent{fromID: p.id, set: received}
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
	rm := h.rootsMgr
	rm.exLock.Lock()
	defer rm.exLock.Unlock()

	switch true {
	case rm.isAcceptableExclusionSet(received):
		// ensure locally stored signatures are not lost
		if rm.targetExSet != nil && rm.targetExSet.hash == received.hash {
			received.mergeSignatures(rm.targetExSet.hash, rm.targetExSet.signers)
		}

		if rm.isRootNode() {
			rm.signExclusionSet(received)
		}

		rm.upgradeExclusionSet(received)
		h.exListCh <- &exclusionSetEvent{set: received}
	case rm.exclusionSet != nil && rm.exclusionSet.hash == received.hash:
		newSignatures := rm.exclusionSet.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		log.Debug("Received new exclusion list signatures", "from", p.id, "singers", toSigners(newSignatures))
		h.exListCh <- &exclusionSetEvent{fromID: p.id, set: rm.exclusionSet.copy()}
		rm.db.saveCurrentExclusionSet(rm.exclusionSet)
	case rm.targetExSet != nil && rm.targetExSet.hash == received.hash:
		newSignatures := rm.targetExSet.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		log.Debug("Received new desired exclusion list signatures", "from", p.id, "singers", toSigners(newSignatures))
		if !rm.currentRootSet().isEnoughExSetSignatures(rm.targetExSet) {
			h.exListCh <- &exclusionSetEvent{fromID: p.id, set: rm.targetExSet}
			rm.db.saveDesiredExclusionSet(rm.targetExSet)
			return nil
		}

		rm.upgradeExclusionSet(rm.targetExSet)
		h.exListCh <- &exclusionSetEvent{set: rm.exclusionSet}
	default:
		if !rm.isRootNode() {
			log.Debug("Ignoring proposed exclusion list: not a root node")
			return nil
		}

		if len(rm.currentRootSet().knownSigners(received.signers)) == 0 {
			log.Debug("Ignoring proposed exclusion list: no current root node signatures")
			return nil
		}

		if rm.proposedExSet != nil && received.timestamp <= rm.proposedExSet.timestamp {
			log.Debug("Ignoring proposed exclusion list: obsolete")
			return nil
		}

		rm.proposedExSet = received
		rm.db.saveProposedExclusionSet(rm.proposedExSet)

		log.Warn("Received a new proposed exclusion list", "hash", received.hash.Hex(), "timestamp", received.timestamp)
	}

	return nil
}

func toSigners(sigs map[common.Address][]byte) []string {
	var strs []string
	for addr := range sigs {
		strs = append(strs, addr.Hex())
	}

	return strs
}
