package governance

import (
	"bytes"
	"compress/gzip"
	"log"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/p2p"
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
	go p.listenForExclusionSets()
	if p.version >= qgov3 {
		go p.listenApprovals()
	}
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

	version int

	rootSetCh           chan *rootSet
	exclusionSetCh      chan *exclusionSet
	approvalCh          chan *common.RootNodeApprovalList
	constitutionFilesCh chan *common.ConstitutionFilesResponse
	done                chan struct{}
}

func newPeer(version int, conn *p2p.Peer, rw p2p.MsgReadWriter) *peer {
	return &peer{
		Peer:                conn,
		id:                  conn.ID().String(),
		rw:                  rw,
		version:             version,
		rootSetCh:           make(chan *rootSet),
		exclusionSetCh:      make(chan *exclusionSet),
		approvalCh:          make(chan *common.RootNodeApprovalList),
		constitutionFilesCh: make(chan *common.ConstitutionFilesResponse),
		done:                make(chan struct{}),
	}
}

func (p *peer) close() {
	close(p.done)
}

type peerStatus struct {
	currentRootSet  *rootSet
	desiredRootSet  *rootSet
	proposedRootSet *rootSet

	currentExSet  *exclusionSet
	desiredExSet  *exclusionSet
	proposedExSet *exclusionSet
}

func (p *peer) handshake(msg statusMsgBody, rm *RootManager) (*peerStatus, error) {
	err := p2p.Send(p.rw, StatusMsg, msg)
	if err != nil {
		return nil, err
	}

	resp, err := p.rw.ReadMsg()
	if err != nil {
		return nil, err
	}

	var status statusMsgBody
	if err := resp.Decode(&status); err != nil {
		return nil, err
	}

	if status.Network != msg.Network {
		return nil, errors.New("invalid network id")
	}

	currentRootSet, err := newRootSet(&status.CurrentRootList)
	if err != nil {
		return nil, errors.Wrap(err, "invalid current root list")
	}

	if currentRootSet == nil {
		return nil, errors.New("empty current root list")
	}
	if currentRootSet != nil {
		currentRootSet.updateAliases(rm.getAliasesOfRoots(currentRootSet.rootAddresses))
	}

	desiredRootSet, err := newRootSet(&status.DesiredRootList)
	if err != nil {
		return nil, errors.Wrap(err, "invalid desired root list")
	}
	if desiredRootSet != nil {
		desiredRootSet.updateAliases(rm.getAliasesOfRoots(desiredRootSet.rootAddresses))
	}

	proposedRootSet, err := newRootSet(&status.ProposedRootList)
	if err != nil {
		return nil, errors.Wrap(err, "invalid proposed root list")
	}
	if proposedRootSet != nil {
		proposedRootSet.updateAliases(rm.getAliasesOfRoots(proposedRootSet.rootAddresses))
	}

	//Validators

	currentExSet, err := newExclusionSet(&status.CurrentExclusionList)
	if err != nil {
		return nil, errors.Wrap(err, "invalid exclusion list")
	}

	desiredExSet, err := newExclusionSet(&status.DesiredExclusionList)
	if err != nil {
		return nil, errors.Wrap(err, "invalid desired exclusion list")
	}

	proposedExSet, err := newExclusionSet(&status.ProposedExclusionList)
	if err != nil {
		return nil, errors.Wrap(err, "invalid proposed exclusion list")
	}

	return &peerStatus{
		currentRootSet:  currentRootSet,
		desiredRootSet:  desiredRootSet,
		proposedRootSet: proposedRootSet,
		currentExSet:    currentExSet,
		desiredExSet:    desiredExSet,
		proposedExSet:   proposedExSet,
	}, nil
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

func (p *peer) listenForExclusionSets() {
	for {
		select {
		case set := <-p.exclusionSetCh:
			if err := p.sendExclusionList(set); err != nil {
				p.Log().Debug("failed to send exclusion set", "err", err)
			}
		case <-p.done:
			return
		}
	}
}

func (p *peer) listenApprovals() {
	for {
		select {
		case approval := <-p.approvalCh:
			if p.version >= qgov3 {
				if err := p.sendApprovalList(approval); err != nil {
					p.Log().Warn("failed to send approval", "err", err)
				}
			}
		case <-p.done:
			return
		}
	}
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

func (p *peer) asyncSendExclusionList(set *exclusionSet) {
	select {
	case p.exclusionSetCh <- set:
	case <-p.done:
		return
	}
}

func (p *peer) asyncSendApprovals(approvalList *common.RootNodeApprovalList) {
	select {
	case p.approvalCh <- approvalList:
	case <-p.done:
		return
	}
}

func (p *peer) asyncSendConstitutionFiles(cm *ConstitutionManager, files []common.ConstitutionFile) {
	var res common.ConstitutionFilesResponse
	for _, file := range files {
		resPath := filepath.Join(cm.baseDir, file.Name)

		if errE := cm.fileExists(resPath); errE != nil {
			if errD := cm.fileExists(filepath.Join(cm.baseDir, draftsDir, file.Name)); errD != nil {
				p.Log().Error("failed to open requested file", "err", errD)
				continue
			}
			resPath = filepath.Join(cm.baseDir, draftsDir, file.Name)
		}

		content, errC := cm.getFileContents(resPath)
		if errC != nil {
			p.Log().Error("failed to open requested file", "err", errC)
			continue
		}

		var b bytes.Buffer
		gz := gzip.NewWriter(&b)
		if _, err := gz.Write(content); err != nil {
			log.Fatal(err)
		}
		if err := gz.Close(); err != nil {
			log.Fatal(err)
		}

		res.Files = append(res.Files, common.ConstitutionFileContent{
			Hash: file.Hash,
			Data: b.Bytes(),
		})
	}
	p.sendConstitutionFiles(&res)
}

func (p *peer) sendExclusionList(set *exclusionSet) error {
	return p2p.Send(p.rw, ExclusionListMsg, set.makeList())
}

func (p *peer) sendApprovalList(approvalList *common.RootNodeApprovalList) error {
	return p2p.Send(p.rw, ApprovalMsg, approvalList)
}

func (p *peer) sendConstitutionFileRequest(request *common.ConstitutionFilesRequest) error {
	if p.version >= qgov4 {
		return p2p.Send(p.rw, ConstitutionFileRequestMsg, request)
	}
	return nil
}

func (p *peer) sendKnownConstitutionFiles(files *common.KnownConstitutionFilesMessage) error {
	if p.version >= qgov4 {
		return p2p.Send(p.rw, KnownConstitutionFilesMsg, files)
	}
	return nil
}

func (p *peer) sendConstitutionFiles(files *common.ConstitutionFilesResponse) error {
	if p.version >= qgov4 {
		return p2p.Send(p.rw, ConstitutionFilesMsg, files)
	}
	return nil
}
