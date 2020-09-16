package governance

import (
	"sync"

	"gitlab.com/q-dev/go-ethereum/common"
	"gitlab.com/q-dev/go-ethereum/p2p"
)

type peerSet struct {
	closed bool
	peers  map[string]*peer
	lock   sync.Mutex
}

func newPeerSet() *peerSet {
	return &peerSet{
		peers: make(map[string]*peer),
	}
}

func (s *peerSet) register(p *peer) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.closed {
		p.Disconnect(p2p.DiscQuitting)
		return
	}

	s.peers[p.id] = p
	go p.listenForRootSets()
}

func (s *peerSet) unregister(p *peer) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.peers, p.id)
	p.close()
}

func (s *peerSet) all() []*peer {
	s.lock.Lock()
	defer s.lock.Unlock()

	peers := make([]*peer, 0, len(s.peers))
	for _, p := range s.peers {
		peers = append(peers, p)
	}

	return peers
}

func (s *peerSet) close() {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, p := range s.peers {
		p.Disconnect(p2p.DiscQuitting)
	}
	s.closed = true
}

// wrapper around the p2p peer.
type peer struct {
	*p2p.Peer

	id string
	rw p2p.MsgReadWriter

	rootSetCh chan *rootSet
	done      chan struct{}
}

func newPeer(conn *p2p.Peer, rw p2p.MsgReadWriter) *peer {
	return &peer{
		Peer:      conn,
		id:        conn.ID().ShortString(),
		rw:        rw,
		rootSetCh: make(chan *rootSet),
		done:      make(chan struct{}),
	}
}

func (p *peer) close() {
	close(p.done)
}

type peerStatus struct {
	rootSet *rootSet
}

func (p *peer) handshake(rootList common.RootList) (*peerStatus, error) {
	err := p2p.Send(p.rw, StatusMsg, statusMsgBody{RootList: rootList})
	if err != nil {
		return nil, err
	}

	msg, err := p.rw.ReadMsg()
	if err != nil {
		return nil, err
	}

	var status statusMsgBody
	if err := msg.Decode(&status); err != nil {
		return nil, err
	}

	incoming, err := newRootSet(&status.RootList)
	if err != nil {
		return nil, err
	}

	return &peerStatus{rootSet: incoming}, nil
}

func (p *peer) listenForRootSets() {
	for {
		select {
		case set := <-p.rootSetCh:
			if err := p.sendRootList(set); err != nil {
				p.Log().Debug("failed to send target root set", "err", err)
			}
		case <-p.done:
			return
		}
	}
}

func (p *peer) sendStatus(rootList common.RootList) error {
	return p2p.Send(p.rw, StatusMsg, statusMsgBody{RootList: rootList})
}

func (p *peer) asyncSendRootList(set *rootSet) {
	select {
	case p.rootSetCh <- set:
	case <-p.done:
		return
	}
}

func (p *peer) sendRootList(set *rootSet) error {
	return p2p.Send(p.rw, RootListMsg, set.makeList())
}
