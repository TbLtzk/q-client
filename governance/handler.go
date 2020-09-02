package governance

import (
	"sync"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/go-ethereum/common"
	"gitlab.com/q-dev/go-ethereum/log"
	"gitlab.com/q-dev/go-ethereum/p2p"
)

const (
	// not really important, just need some buffer for robustness
	signaturesBuff = 4
)

// handler of protocol messages.
type handler struct {
	rootsMgr *RootManager

	newSignaturesCh chan newSignaturesEvent
	done            chan struct{}

	peerWG sync.WaitGroup
	peers  *peerSet
}

type newSignaturesEvent struct {
	hash       common.Hash
	signatures map[common.Address][]byte
}

func newHandler(roots *RootManager) *handler {
	return &handler{
		rootsMgr:        roots,
		newSignaturesCh: make(chan newSignaturesEvent, signaturesBuff),
		done:            make(chan struct{}),
		peers:           newPeerSet(),
	}
}

func (h *handler) run() {
	go h.broadcastSignatures()
}

func (h *handler) broadcastSignatures() {
	for {
		select {
		case event := <-h.newSignaturesCh:
			peers := h.peers.peersWithoutSignature(event)
			for _, p := range peers {
				p.asyncSendNewSignatures(event)
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

func (h *handler) sendNewSignatures(event newSignaturesEvent) {
	select {
	case h.newSignaturesCh <- event:
	case <-h.done:
		return
	}
}

func (h *handler) runPeer(p *peer) error {
	rootList := h.rootsMgr.CurrentList()
	status, err := p.handshake(rootList)
	if err != nil {
		return err
	}

	h.peers.register(p)
	defer h.peers.unregister(p)

	currentSet, _ := newRootSet(&rootList)
	h.syncRootLists(p, currentSet, status.rootSet)

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

	newSignatures := h.rootsMgr.currentSet().mergeSignatures(incoming.hash, incoming.signers)
	if len(newSignatures) != 0 {
		logReceivedSigs(p.id, newSignatures)
		h.sendNewSignatures(newSignaturesEvent{hash: incoming.hash, signatures: newSignatures})
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
	case NewSignaturesMsg:
		return h.handleNewSignatures(p, msg)
	//case GetRootListMsg:
	//	return h.handleGetRootList(p, msg)
	//case RootListMsg:
	//	return h.handleRootList(p, msg)
	//case NewRootListMsg:
	//	return h.handleNewRootList(p, msg)
	default:
		return errors.New("unknown msg code")
	}
}

func (h *handler) handleNewSignatures(p *peer, msg p2p.Msg) error {
	var data newSignaturesMsg
	if err := msg.Decode(&data); err != nil {
		return err
	}

	current := h.rootsMgr.currentSet()
	if data.Hash != current.hash {
		return nil
	}

	// todo: send info about what is signed for validation
	sigs, err := current.sanitizeSignatures(data.Signatures)
	if err != nil {
		return err
	}

	newSigs := current.mergeSignatures(data.Hash, sigs)
	if len(newSigs) == 0 {
		return nil
	}

	logReceivedSigs(p.id, newSigs)
	h.sendNewSignatures(newSignaturesEvent{hash: data.Hash, signatures: newSigs})

	return nil
}

func (h *handler) handleNewRootList(p *peer, msg p2p.Msg) error {
	var list common.RootList
	if err := msg.Decode(&list); err != nil {
		return err
	}

	set, err := newRootSet(&list)
	if err != nil {
		return err
	}

	current := h.rootsMgr.currentSet()
	if set.timestamp <= current.timestamp {
		return nil
	}

	if !current.isAcceptable(set) {
		return errIncomplete
	}

	h.rootsMgr.upgrade(set)
	_ = h.sendNewRootList(p)

	return nil
}

func (h *handler) sendNewRootList(receivedFrom *peer) error {
	for _, peer := range h.peers.peers {
		if receivedFrom.id == peer.id {
			continue
		}
		err := peer.sendRootList(h.rootsMgr.CurrentList())
		if err != nil {
			log.Warn("faulty peer", err)
		}
	}
	return nil
}

func logReceivedSigs(from string, sigs map[common.Address][]byte) {
	var strs []string
	for addr := range sigs {
		strs = append(strs, addr.Hex())
	}

	log.Debug("received new root list signatures", "signers", strs, "from", from)
}
