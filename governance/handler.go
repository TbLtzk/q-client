package governance

import (
	"sync"

	"gitlab.com/q-dev/go-ethereum/log"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/go-ethereum/p2p"
)

// handler of protocol messages.
type handler struct {
	roots *RootSet

	peerWG sync.WaitGroup
	peers  *peerSet
}

//func newHandler(roots *RootSet) *handler {
//	return &handler{
//		roots: roots,
//		peers: newPeerSet(),
//	}
//}

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
				Timestamp uint64
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
		return h.handleGetRootList(p, msg)
	case RootListMsg:
		return h.handleRootList(p, msg)
	case NewRootListMsg:
		return h.handleNewRootList(p, msg)
	default:
		return errors.New("unknown msg code")
	}
}

func (h *handler) handleGetRootList(p *peer, msg p2p.Msg) error {
	err := p.sendRootList(h.roots.CurrentList())
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) handleRootList(p *peer, msg p2p.Msg) error {
	var (
		list RootList
		err  error
	)
	err = msg.Decode(&list)
	if err != nil {
		return err
	}

	if list.Timestamp < h.roots.CurrentList().Timestamp {
		return nil
	}

	err = h.roots.Validate(list)
	if err != nil {
		return err
	}

	h.roots.UpdateList(list)

	return nil
}

func (h *handler) handleNewRootList(p *peer, msg p2p.Msg) error {
	var (
		list RootList
		err  error
	)
	err = msg.Decode(&list)
	if err != nil {
		return err
	}

	if list.Timestamp < h.roots.CurrentList().Timestamp {
		return nil
	}

	err = h.roots.Validate(list)
	if err != nil {
		return err
	}

	h.roots.UpdateList(list)

	err = h.sendNewRootList(p)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) sendNewRootList(receivedFrom *peer) error {
	for _, peer := range h.peers.peers {
		if receivedFrom.id == peer.id {
			continue
		}
		err := peer.sendRootList(h.roots.CurrentList())
		if err != nil {
			log.Warn("faulty peer", err)
		}
	}
	return nil
}
