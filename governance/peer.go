package governance

import (
	"sync"

	mapset "github.com/deckarep/golang-set"

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
	go p.broadcastSignatures()
}

func (s *peerSet) unregister(p *peer) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.peers, p.id)
	p.close()
}

func (s *peerSet) peersWithoutSignature(event newSignaturesEvent) []*peer {
	s.lock.Lock()
	defer s.lock.Unlock()

	var peers []*peer
	signers := toMapSet(event.signatures)
	for _, p := range s.peers {
		if !p.isKnownSigners(event.hash, signers) {
			peers = append(peers, p)
		}
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

	lock         sync.Mutex
	hash         common.Hash
	knownSigners mapset.Set

	signaturesCh chan newSignaturesEvent
	done         chan struct{}
}

func newPeer(conn *p2p.Peer, rw p2p.MsgReadWriter) *peer {
	return &peer{
		Peer:         conn,
		id:           conn.ID().ShortString(),
		rw:           rw,
		knownSigners: mapset.NewSet(),
		signaturesCh: make(chan newSignaturesEvent),
		done:         make(chan struct{}),
	}
}

func (p *peer) close() {
	close(p.done)
}

type peerStatus struct {
	rootSet *rootSet
}

func (p *peer) handshake(rootList common.RootList) (*peerStatus, error) {
	err := p2p.Send(p.rw, StatusMsg, statusMsgData{RootList: rootList})
	if err != nil {
		return nil, err
	}

	msg, err := p.rw.ReadMsg()
	if err != nil {
		return nil, err
	}

	var status statusMsgData
	if err := msg.Decode(&status); err != nil {
		return nil, err
	}

	incoming, err := newRootSet(&status.RootList)
	if err != nil {
		return nil, err
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	p.hash = incoming.hash
	p.knownSigners = incoming.signersMap()

	sent, _ := newRootSet(&rootList)
	if incoming.isAcceptable(sent) {
		p.hash = sent.hash
		p.knownSigners = sent.signersMap()
	} else if incoming.hash == sent.hash {
		p.knownSigners = p.knownSigners.Union(sent.signersMap())
	}

	return &peerStatus{rootSet: incoming}, nil
}

func (p *peer) broadcastSignatures() {
	for {
		select {
		case event := <-p.signaturesCh:
			if err := p.sendNewSignatures(event); err != nil {
				p.Log().Debug("failed to send new signatures", "err", err)
			}
		case <-p.done:
			return
		}
	}
}

func (p *peer) isKnownSigners(hash common.Hash, signers mapset.Set) bool {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.hash != hash {
		return false
	}

	return signers.IsSubset(p.knownSigners)
}

func (p *peer) currentHash() common.Hash {
	p.lock.Lock()
	defer p.lock.Unlock()

	return p.hash
}

func (p *peer) setCurrentRootList(hash common.Hash, signers mapset.Set) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.hash = hash
	p.knownSigners = signers
}

func (p *peer) sendStatus(rootList common.RootList) error {
	return p2p.Send(p.rw, StatusMsg, statusMsgData{RootList: rootList})
}

func (p *peer) sendRootList(list common.RootList) error {
	return p2p.Send(p.rw, NewRootListMsg, list)
}

func (p *peer) sendNewRootList(list common.RootList) error {
	return p2p.Send(p.rw, NewRootListMsg, list)
}

func (p *peer) asyncSendNewSignatures(event newSignaturesEvent) {
	select {
	case p.signaturesCh <- event:
	case <-p.done:
		return
	}
}

func (p *peer) sendNewSignatures(msg newSignaturesEvent) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if msg.hash != p.hash {
		return nil
	}

	set := toMapSet(msg.signatures)
	if set.IsSubset(p.knownSigners) {
		return nil
	}

	err := p2p.Send(p.rw, NewSignaturesMsg, initNewSignaturesMsg(msg.hash, msg.signatures))
	if err != nil {
		return err
	}

	p.knownSigners = p.knownSigners.Union(set)
	return nil
}
