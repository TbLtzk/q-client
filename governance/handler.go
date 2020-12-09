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
			current := h.rootsMgr.CurrentList()
			return struct {
				Timestamp uint64
			}{
				Timestamp: current.Timestamp,
			}
		},
	}
}

func (h *handler) runPeer(p *peer) error {
	currentRootSet := h.rootsMgr.currentRootSet()
	currentExSet := h.rootsMgr.currentExclusionSet()
	status, err := p.handshake(currentRootSet.makeList(), currentExSet.makeList())
	if err != nil {
		return err
	}

	h.peers.register(p)
	defer h.peers.unregister(p)

	h.syncRootLists(p, status.rootSet)
	h.syncExclusionLists(p, status.exclusionSet)

	h.peerWG.Add(1)
	defer h.peerWG.Done()

	for {
		if err := h.handleMsg(p); err != nil {
			p.Log().Debug("Governance message handling failed", "err", err)
			return err
		}
	}
}

func (h *handler) syncRootLists(p *peer, incoming *rootSet) {
	defer func() {
		// push target root list if have any locally
		target := h.rootsMgr.targetRootSet()
		if target != nil {
			p.asyncSendRootList(target)
		}
	}()

	rm := h.rootsMgr
	rm.lock.Lock()
	defer rm.lock.Unlock()

	if rm.current.hash == incoming.hash {
		newSignatures := rm.current.mergeSignatures(incoming.hash, incoming.signers)
		if len(newSignatures) != 0 {
			log.Debug("received root list signatures", "from", p.id, "signers", toSigners(incoming.signers))
			logReceivedSigs(p.id, newSignatures)

			h.rootListCh <- &rootSetEvent{fromID: p.id, set: rm.current.copy()}
		}

		return
	}

	if !rm.current.isAcceptable(incoming) {
		return
	}

	rm.current = incoming
	rm.db.saveCurrentRootSet(rm.current)

	rm.isRoot = rm.isMember(rm.current.rootAddresses)

	log.Info("upgraded root list", "hash", incoming.hash.Hex(), "timestamp", incoming.timestamp)
	h.rootListCh <- &rootSetEvent{set: incoming}

	if rm.target != nil && rm.target.timestamp <= incoming.timestamp {
		log.Info("dropping outdated desired root list",
			"hash", rm.target.hash.Hex(),
			"timestamp", rm.target.timestamp)

		rm.target = nil
		rm.db.deleteDesiredRootSet()
	}

	if rm.proposed != nil && rm.proposed.timestamp <= incoming.timestamp {
		rm.proposed = nil
		rm.db.deleteProposedRootSet()
	}
}

func (h *handler) syncExclusionLists(p *peer, incoming *exclusionSet) {
	defer func() {
		target := h.rootsMgr.currentExclusionSet()
		if target != nil {
			p.asyncSendExclusionList(target)
		}
	}()

	if incoming == nil {
		return
	}

	if h.rootsMgr.tryUpgradeExclusionList(incoming) {
		log.Info("upgraded exclusion list", "hash", incoming.hash, "timestamp", incoming.timestamp)
		return
	}

	current := h.rootsMgr.currentExclusionSet()
	if current == nil || current.hash != incoming.hash {
		return
	}

	newSignatures := current.mergeSignatures(incoming.hash, incoming.signers)
	if len(newSignatures) != 0 {
		logReceivedSigs(p.id, newSignatures)

		h.exListCh <- &exclusionSetEvent{fromID: p.id, set: current}
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
		return h.handleRootList(p, msg)
	case ExclusionListMsg:
		return h.handleExclusionList(p, msg)
	default:
		return errors.New("unknown msg code")
	}
}

func (h *handler) handleRootList(p *peer, msg p2p.Msg) error {
	var list common.RootList
	if err := msg.Decode(&list); err != nil {
		return err
	}

	received, err := newRootSet(&list)
	if err != nil {
		return err
	}

	rm := h.rootsMgr
	rm.lock.Lock()
	defer rm.lock.Unlock()

	switch {
	case rm.current.isAcceptable(received) && (rm.target == nil || rm.target.hash != received.hash):
		if rm.isMember(received.rootAddresses) {
			rm.signRootSet(received)
		}

		rm.upgrade(received)
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

	if h.rootsMgr.tryUpgradeExclusionList(received) {
		log.Info("upgraded exclusion list", "hash", received.hash.Hex(), "timestamp", received.timestamp)
		h.exListCh <- &exclusionSetEvent{fromID: p.id, set: received}
		return nil
	}

	current := h.rootsMgr.currentExclusionSet()
	if current != nil && current.hash == received.hash {
		newSignatures := current.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) != 0 {
			logReceivedSigs(p.id, newSignatures)
			h.exListCh <- &exclusionSetEvent{fromID: p.id, set: current}
		}

		return nil
	}

	target := h.rootsMgr.targetExclusionSet()
	if target == nil {
		log.Debug("received target exclusion list but has none locally", "from", p.id)
		return nil
	}

	if target.hash != received.hash {
		log.Warn("received target exclusion set does not match local", "hash", received.hash.Hex(), "from", p.id)
		return nil
	}

	newSignatures := target.mergeSignatures(received.hash, received.signers)
	if len(newSignatures) == 0 {
		return nil
	}

	logReceivedSigs(p.id, newSignatures)
	if h.rootsMgr.tryUpgradeExclusionList(target) {
		log.Info("upgraded exclusion list", "hash", target.hash.Hex(), "timestamp", target.timestamp)
		h.exListCh <- &exclusionSetEvent{fromID: p.id, set: target}

		return nil
	}

	h.exListCh <- &exclusionSetEvent{fromID: p.id, set: target}
	return nil
}

func logReceivedSigs(from string, sigs map[common.Address][]byte) {
	var strs []string
	for addr := range sigs {
		strs = append(strs, addr.Hex())
	}

	log.Debug("received new signatures", "signers", strs, "from", from)
}

func toSigners(sigs map[common.Address][]byte) []string {
	var strs []string
	for addr := range sigs {
		strs = append(strs, addr.Hex())
	}

	return strs
}
