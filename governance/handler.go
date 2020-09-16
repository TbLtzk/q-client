package governance

import (
	"sync"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/go-ethereum/common"
	"gitlab.com/q-dev/go-ethereum/event"
	"gitlab.com/q-dev/go-ethereum/log"
	"gitlab.com/q-dev/go-ethereum/p2p"
)

// handler of protocol messages.
type handler struct {
	rootsMgr *RootManager

	targetRootListCh     chan *rootSet
	newTargetRootListSub event.Subscription

	rootListCh chan *rootSetEvent
	done       chan struct{}

	peerWG sync.WaitGroup
	peers  *peerSet
}

func newHandler(roots *RootManager) *handler {
	newRoots := make(chan *rootSet)
	newRootsSub := roots.targetRootFeed.Subscribe(newRoots)

	return &handler{
		rootsMgr:             roots,
		targetRootListCh:     newRoots,
		newTargetRootListSub: newRootsSub,
		rootListCh:           make(chan *rootSetEvent),
		done:                 make(chan struct{}),
		peers:                newPeerSet(),
	}
}

func (h *handler) run() {
	go h.listenForTargetRootList()
	go h.broadcastRootSets()
}

type rootSetEvent struct {
	fromID string
	set    *rootSet
}

func (h *handler) listenForTargetRootList() {
	for {
		select {
		case set := <-h.targetRootListCh:
			for _, p := range h.peers.all() {
				log.Debug("SENDING target set", "to", p.id)
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

func (h *handler) broadcastTargetRootSet(set *rootSet) {
	select {
	case h.targetRootListCh <- set:
	case <-h.done:
		return
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
	currentRootSet := h.rootsMgr.currentSet()
	status, err := p.handshake(currentRootSet.makeList())
	if err != nil {
		return err
	}

	h.peers.register(p)
	defer h.peers.unregister(p)

	h.syncRootLists(p, currentRootSet, status.rootSet)

	h.peerWG.Add(1)
	defer h.peerWG.Done()

	for {
		if err := h.handleMsg(p); err != nil {
			p.Log().Debug("Governance message handling failed", "err", err)
			return err
		}
	}
}

func (h *handler) syncRootLists(p *peer, current, incoming *rootSet) {
	if current.isAcceptable(incoming) {
		h.rootsMgr.upgrade(incoming)
		return
	}

	if current.hash != incoming.hash || current.timestamp != incoming.timestamp {
		return
	}

	newSignatures := current.mergeSignatures(incoming.hash, incoming.signers)
	if len(newSignatures) != 0 {
		logReceivedSigs(p.id, newSignatures)

		// todo: this may be unnecessary with current approach
		h.rootListCh <- &rootSetEvent{fromID: p.id, set: current.copy()}
	}

	// push target root list if have any locally
	target := h.rootsMgr.targetSet()
	if target != nil {
		p.asyncSendRootList(target)
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

	current := h.rootsMgr.currentSet()
	// check if it's 'push'
	if current.isAcceptable(received) && h.rootsMgr.upgrade(received) {
		log.Info("upgraded root list", "hash", received.hash, "timestamp", received.timestamp)
		h.rootListCh <- &rootSetEvent{fromID: p.id, set: received}
		return nil
	}

	if current.hash == received.hash {
		newSignatures := current.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) != 0 {
			logReceivedSigs(p.id, newSignatures)
			h.rootListCh <- &rootSetEvent{fromID: p.id, set: current.copy()}
		}

		return nil
	}

	target := h.rootsMgr.targetSet()
	// todo: save and wait action from the operator
	if target == nil {
		log.Debug("received target root list but has none locally", "from", p.id)
		return nil
	}

	// todo: also check set itself, hash can be different because of timestamp
	if target.hash != received.hash {
		log.Warn("received target root list does not match local", "hash", received.hash, "from", p.id)
		return nil
	}

	log.Debug("received target root list", "hash", received.hash, "from", p.id)
	newSignatures := target.mergeSignatures(received.hash, received.signers)
	if len(newSignatures) == 0 {
		return nil
	}

	logReceivedSigs(p.id, newSignatures)

	// trigger 'push'
	if h.rootsMgr.currentSet().isAcceptable(target) && h.rootsMgr.upgrade(target) {
		log.Info("upgraded root list", "hash", target.hash, "timestamp", target.timestamp)

		// todo: probably it's also make sense to send it back
		h.rootListCh <- &rootSetEvent{fromID: p.id, set: target.copy()}
		return nil
	}

	h.rootListCh <- &rootSetEvent{fromID: p.id, set: target.copy()}
	return nil
}

func logReceivedSigs(from string, sigs map[common.Address][]byte) {
	var strs []string
	for addr := range sigs {
		strs = append(strs, addr.Hex())
	}

	log.Debug("received new root list signatures", "signers", strs, "from", from)
}
