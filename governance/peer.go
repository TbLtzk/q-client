package governance

import (
	"sync"

	"gitlab.com/q-dev/go-ethereum/p2p"
)

type peerSet struct {
	peers map[string]*peer
	lock  sync.Mutex
}

//func newPeerSet() *peerSet {
//	return &peerSet{
//		peers: make(map[string]*peer),
//	}
//}

func (s *peerSet) register(p *peer) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.peers[p.id] = p
}

func (s *peerSet) unregister(p *peer) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.peers, p.id)
}

// wrapper around the p2p peer.
type peer struct {
	*p2p.Peer

	id string
	rw p2p.MsgReadWriter

	// filled after handshake
	rootListTimestamp int64
}

func newPeer(conn *p2p.Peer, rw p2p.MsgReadWriter) *peer {
	return &peer{
		Peer: conn,
		id:   conn.ID().ShortString(),
		rw:   rw,
	}
}

func (p *peer) getRootList() error {
	return p2p.Send(p.rw, GetRootListMsg, struct{}{})
}
