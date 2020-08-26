package governance

import (
	"github.com/pkg/errors"
	"gitlab.com/q-dev/go-ethereum/p2p"
	"sync"
)

// handler of protocol messages.
type handler struct {
	roots *RootSet

	peerWG sync.WaitGroup
	peers  *peerSet
}

func newHandler(roots *RootSet) *handler {
	return &handler{
		roots: roots,
		peers: newPeerSet(),
	}
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
			current := h.roots.CurrentList()
			return struct {
				Timestamp int64
			}{
				Timestamp: current.Timestamp,
			}
		},
	}
}

func (h *handler) runPeer(pr *peer) error {
	// just request current list as handshake
	// todo: consider exchanging statuses and requesting lists as needed instead
	if err := pr.getRootList(); err != nil {
		return err
	}

	h.peers.register(pr)
	defer h.peers.unregister(pr)

	h.peerWG.Add(1)
	defer h.peerWG.Done()

	for {
		if err := h.handleMsg(pr); err != nil {
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
	case GetRootListMsg:
		return h.handleGetRootList(msg)
	case RootListMsg:
		return h.handleRootList(msg)
	default:
		return errors.New("unknown msg code")
	}
}

func (h *handler) handleGetRootList(msg p2p.Msg) error {
	// TODO: implement me
	return nil
}

func (h *handler) handleRootList(msg p2p.Msg) error {
	// TODO: implement me
	return nil
}

func (h *handler) handleNewRootList(msg p2p.Msg) error {
	// TODO: implement me
	return nil
}
