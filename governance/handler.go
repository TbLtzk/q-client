package governance

import (
	"bytes"
	"compress/gzip"
	"io"
	"sync"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/event"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/p2p"
)

var (
	errUnknownMsgCode = errors.New("unknown msg code")
	errMsgTooLarge    = errors.New("message too large")
)

// handler of governance protocol messages.
type handler struct {
	rootManager         *RootManager
	constitutionManager *ConstitutionManager

	desiredRootCh  chan *rootSet
	desiredRootSub event.Subscription

	desiredExCh  chan *exclusionSet
	desiredExSub event.Subscription

	rootEventCh     chan *rootSetEvent
	exEventCh       chan *exclusionSetEvent
	approvalEventCh chan *approvalEvent
	done            chan struct{}

	approvalCh  chan *common.RootNodeApprovalList
	approvalSub event.Subscription

	peerWG sync.WaitGroup
	peers  *peerSet

	failureCounts map[string]map[uint64]uint64 // peer id => msg code => failures count
	maxFailures   uint64

	typedRelayDedup       *typedRelayDedup
	typedRootRelayCh      chan *typedRootRelayEvent
	typedExclusionRelayCh chan *typedExclusionRelayEvent
}

func newHandler(rootManager *RootManager, cm *ConstitutionManager) *handler {
	desiredRootCh := make(chan *rootSet)
	desiredRootSub := rootManager.desiredRootFeed.Subscribe(desiredRootCh)

	desiredExCh := make(chan *exclusionSet)
	desiredExSub := rootManager.desiredExFeed.Subscribe(desiredExCh)

	approvalCh := make(chan *common.RootNodeApprovalList)
	approvalSub := rootManager.approvalFeed.Subscribe(approvalCh)

	return &handler{
		constitutionManager: cm,

		rootManager: rootManager,

		desiredRootCh:  desiredRootCh,
		desiredRootSub: desiredRootSub,

		desiredExCh:  desiredExCh,
		desiredExSub: desiredExSub,

		rootEventCh:     make(chan *rootSetEvent),
		exEventCh:       make(chan *exclusionSetEvent),
		approvalEventCh: make(chan *approvalEvent),
		done:            make(chan struct{}),

		approvalCh:  approvalCh,
		approvalSub: approvalSub,

		peerWG: sync.WaitGroup{},
		peers:  newPeerSet(),

		failureCounts: make(map[string]map[uint64]uint64),
		maxFailures:   rootManager.ApprovalMaxFailures,

		typedRelayDedup:       newTypedRelayDedup(),
		typedRootRelayCh:      make(chan *typedRootRelayEvent),
		typedExclusionRelayCh: make(chan *typedExclusionRelayEvent),
	}
}

func (h *handler) run() {
	go h.listenForDesiredRootList()
	go h.broadcastRootSets()

	go h.listenForDesiredExclusionList()
	go h.broadcastExclusionSets()

	go h.listenForRNApprovals()
	go h.broadcastApprovals()

	go h.broadcastTypedRootRelays()
	go h.broadcastTypedExclusionRelays()
}

type rootSetEvent struct {
	fromID string
	set    *rootSet
}

type exclusionSetEvent struct {
	fromID string
	set    *exclusionSet
}

type approvalEvent struct {
	fromID   string
	approval *common.RootNodeApprovalList
}

func (h *handler) listenForDesiredRootList() {
	for {
		select {
		case set := <-h.desiredRootCh:
			for _, p := range h.peers.all() {
				log.Debug("Sending desired root set", "to", p.id)
				p.asyncSendRootList(set)
			}
		case err := <-h.desiredRootSub.Err():
			log.Debug("no new desired roots sets", "err", err)
			return
		case <-h.done:
			return
		}
	}
}

func (h *handler) broadcastRootSets() {
	for {
		select {
		case msg := <-h.rootEventCh:
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

func (h *handler) listenForDesiredExclusionList() {
	for {
		select {
		case set := <-h.desiredExCh:
			for _, p := range h.peers.all() {
				log.Debug("Sending desired exclusion set", "to", p.id)
				p.asyncSendExclusionList(set)
			}
		case err := <-h.desiredExSub.Err():
			log.Debug("no new desired exclusion sets", "err", err)
			return
		case <-h.done:
			return
		}
	}
}

func (h *handler) broadcastExclusionSets() {
	for {
		select {
		case msg := <-h.exEventCh:
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

func (h *handler) listenForRNApprovals() {
	for {
		select {
		case approval := <-h.approvalCh:
			for _, p := range h.peers.all() {
				log.Debug("Sending approval list root node approvals", "to", p.id)
				p.asyncSendApprovals(approval)
			}
		case err := <-h.approvalSub.Err():
			log.Debug("no new approvals", "err", err)
			return
		case <-h.done:
			return
		}
	}
}

func (h *handler) broadcastApprovals() {
	for {
		select {
		case msg := <-h.approvalEventCh:
			for _, p := range h.peers.all() {
				if msg.fromID == p.id {
					continue
				}

				p.asyncSendApprovals(msg.approval)
			}
		case <-h.done:
			return
		}
	}
}

func (h *handler) broadcastConstitutionRequest(request *common.ConstitutionFilesRequest) {
	for _, p := range h.peers.all() {
		p.sendConstitutionFileRequest(request)
	}
}

//nolint:unused
func (h *handler) broadcastKnownConstitutionFiles(files *common.KnownConstitutionFilesMessage) {
	for _, p := range h.peers.all() {
		p.sendKnownConstitutionFiles(files)
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
			return h.runPeer(newPeer(int(version), p, rw))
		},
		NodeInfo: func() interface{} {
			active := h.rootManager.getActiveRootSet(true)
			return struct {
				Timestamp uint64
			}{
				Timestamp: active.timestamp,
			}
		},
	}
}

func (h *handler) makeStatusBody(rm *RootManager, cm *ConstitutionManager) statusMsgBody {
	rm.exLock.Lock()
	defer rm.exLock.Unlock()

	rm.rootLock.Lock()
	defer rm.rootLock.Unlock()

	rm.active.aliases = rm.getAliasesOfRoots(rm.active.rootAddresses)
	if rm.desired != nil {
		rm.desired.aliases = rm.getAliasesOfRoots(rm.desired.rootAddresses)
	}
	if rm.proposed != nil {
		rm.proposed.aliases = rm.getAliasesOfRoots(rm.proposed.rootAddresses)
	}

	return statusMsgBody{
		CurrentRootList:  rm.active.copy().makeList(),
		DesiredRootList:  rm.desired.copy().makeList(),
		ProposedRootList: rm.proposed.copy().makeList(),

		CurrentExclusionList:  rm.activeExSet.copy().makeList(),
		DesiredExclusionList:  rm.desiredExSet.copy().makeList(),
		ProposedExclusionList: rm.proposedExSet.copy().makeList(),

		Network: rm.networkId,
	}
}

func (h *handler) runPeer(p *peer) error {
	rm := h.rootManager
	cm := h.constitutionManager

	statusBody := h.makeStatusBody(rm, cm)
	status, err := p.handshake(statusBody, h.rootManager)
	if err != nil {
		return err
	}
	if status == nil {
		return errors.New("status is nil")
	}

	// Propagate current root set to newly connected peers
	if shouldPropagateRoots(rm.desired, status.desiredRootSet, rm.active) {
		h.propagateRootSet(status.desiredRootSet)
	}
	if shouldPropagateRoots(rm.proposed, status.proposedRootSet, rm.active) {
		h.propagateRootSet(status.proposedRootSet)
	}
	if shouldPropagateRoots(rm.active, status.currentRootSet, rm.active) {
		h.propagateRootSet(status.currentRootSet)
	}

	// Propagate current exclusion set to newly connected peers
	if shouldPropagateExcl(rm.desiredExSet, status.desiredExSet, rm.active) {
		h.propagateExclusionSet(status.desiredExSet)
	}
	if shouldPropagateExcl(rm.proposedExSet, status.proposedExSet, rm.active) {
		h.propagateExclusionSet(status.proposedExSet)
	}
	if shouldPropagateExcl(rm.activeExSet, status.currentExSet, rm.active) {
		h.propagateExclusionSet(status.currentExSet)
	}

	approvals := rm.db.getLastApprovals()
	if err = p.sendApprovalList(approvals); err != nil {
		p.Log().Warn("failed to send approval", "err", err, "root-set", approvals)
	}

	h.peers.register(p)
	defer h.peers.unregister(p)

	// Status nil check goes higher
	for _, set := range []struct {
		our   *rootSet
		their *rootSet
	}{
		{rm.active, status.currentRootSet},
		{rm.desired, status.desiredRootSet},
		{rm.proposed, status.proposedRootSet},
	} {
		our, their := set.our, set.their
		if shouldPropagateRoots(our, their, rm.active) {
			if err = h.importRootSet(p.id, their, true); err != nil {
				return err
			}
		}
	}

	for _, set := range []struct {
		our   *exclusionSet
		their *exclusionSet
	}{
		{rm.activeExSet, status.currentExSet},
		{rm.desiredExSet, status.desiredExSet},
		{rm.proposedExSet, status.proposedExSet},
	} {
		our, their := set.our, set.their
		if shouldPropagateExcl(our, their, rm.active) {
			if err = h.importExclusionSet(p.id, their, true); err != nil {
				return err
			}
		}
	}

	h.constitutionManager.insureLatestHashIsHandled()

	newReq := common.ConstitutionFilesRequest{Hashes: h.constitutionManager.requiredHashes}
	p.sendConstitutionFileRequest(&newReq)

	filesReq := common.KnownConstitutionFilesMessage{Hashes: h.constitutionManager.knownHashes}
	p.sendKnownConstitutionFiles(&filesReq)

	h.peerWG.Add(1)
	defer h.peerWG.Done()

	for {
		if err = h.handleMsg(p); err != nil {
			p.Log().Debug("Governance message handling failed", "err", err)
			return err
		}
	}
}

func shouldPropagateExcl(our, their *exclusionSet, active *rootSet) bool {
	if their == nil {
		return false
	}
	if our == nil {
		return len(active.knownSigners(their.signers)) != 0 // signatures were checked in handshake func
	}

	if our.hash == their.hash {
		return !areSignersEqual(our.signers, their.signers)
	}

	if their.timestamp > our.timestamp {
		return len(active.knownSigners(their.signers)) != 0 // signatures were checked in handshake func
	}

	return false
}

func shouldPropagateRoots(our, their, active *rootSet) bool {
	if their == nil {
		return false
	}

	if our == nil {
		return len(active.knownSigners(their.signers)) != 0 // signatures were checked in handshake func
	}

	if our.hash == their.hash {
		return !areSignersEqual(our.signers, their.signers)
	}

	if their.timestamp > our.timestamp {
		return len(active.knownSigners(their.signers)) != 0 // signatures were checked in handshake func
	}

	return false
}

func areSignersEqual(ourSigners, theirSigners map[common.Address][]byte) bool {
	if len(ourSigners) != len(theirSigners) {
		return false
	}

	for addr := range theirSigners {
		if _, ok := ourSigners[addr]; ok {
			continue
		}

		return false
	}

	return true
}

type rootListImportSnapshot struct {
	activeSigners   int
	desiredSigners  int
	proposedSigners int

	activeHash   common.Hash
	desiredHash  common.Hash
	proposedHash common.Hash
}

func (s rootListImportSnapshot) changedFrom(before rootListImportSnapshot) bool {
	return s.activeHash != before.activeHash ||
		s.desiredHash != before.desiredHash ||
		s.proposedHash != before.proposedHash ||
		s.activeSigners != before.activeSigners ||
		s.desiredSigners != before.desiredSigners ||
		s.proposedSigners != before.proposedSigners
}

func (s *RootManager) snapshotRootListImport(set *rootSet) (rootListImportSnapshot, error) {
	if set == nil {
		return rootListImportSnapshot{}, errProposedRootListEmpty
	}

	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	set.updateAliases(s.getAliasesOfRoots(set.rootAddresses))
	if len(s.active.knownSigners(set.signers)) == 0 {
		return rootListImportSnapshot{}, errors.New("signed root list has no active root signatures")
	}
	if set.hash != s.active.hash && set.timestamp <= s.active.timestamp {
		return rootListImportSnapshot{}, errProposedRootListObsolete
	}
	if s.desired != nil && set.hash != s.desired.hash && set.timestamp <= s.desired.timestamp {
		return rootListImportSnapshot{}, errProposedRootListObsolete
	}
	if s.proposed != nil && set.hash != s.proposed.hash && set.timestamp <= s.proposed.timestamp {
		return rootListImportSnapshot{}, errProposedRootListObsolete
	}
	if s.active == nil || s.active.hash != set.hash {
		exceeded, err := s.isSetQuotaExceededDryRun(set)
		if err != nil {
			return rootListImportSnapshot{}, err
		}
		if exceeded {
			return rootListImportSnapshot{}, errors.New("root list quota exceeded")
		}
	}

	return s.rootListImportSnapshotLocked(set.hash), nil
}

func (s *RootManager) rootListImportSnapshot(hash common.Hash) rootListImportSnapshot {
	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	return s.rootListImportSnapshotLocked(hash)
}

func (s *RootManager) rootListImportSnapshotLocked(hash common.Hash) rootListImportSnapshot {
	snapshot := rootListImportSnapshot{}
	if s.active != nil {
		snapshot.activeHash = s.active.hash
		if s.active.hash == hash {
			snapshot.activeSigners = len(s.active.signers)
		}
	}
	if s.desired != nil {
		snapshot.desiredHash = s.desired.hash
		if s.desired.hash == hash {
			snapshot.desiredSigners = len(s.desired.signers)
		}
	}
	if s.proposed != nil {
		snapshot.proposedHash = s.proposed.hash
		if s.proposed.hash == hash {
			snapshot.proposedSigners = len(s.proposed.signers)
		}
	}
	return snapshot
}

type exclusionListImportSnapshot struct {
	activeSigners   int
	desiredSigners  int
	proposedSigners int
	quarantineSize  int

	activeHash   common.Hash
	desiredHash  common.Hash
	proposedHash common.Hash
}

func (s exclusionListImportSnapshot) changedFrom(before exclusionListImportSnapshot) bool {
	return s.activeHash != before.activeHash ||
		s.desiredHash != before.desiredHash ||
		s.proposedHash != before.proposedHash ||
		s.activeSigners != before.activeSigners ||
		s.desiredSigners != before.desiredSigners ||
		s.proposedSigners != before.proposedSigners ||
		s.quarantineSize != before.quarantineSize
}

func (s *RootManager) snapshotExclusionListImport(set *exclusionSet) (exclusionListImportSnapshot, error) {
	if set == nil {
		return exclusionListImportSnapshot{}, errProposedExclusionListEmpty
	}

	if err := validateExclusionSetRanges(set); err != nil {
		return exclusionListImportSnapshot{}, err
	}

	s.exLock.Lock()
	defer s.exLock.Unlock()

	s.active.aliases = s.getAliasesOfRoots(s.active.rootAddresses)
	if len(s.getActiveRootSet(true).knownSigners(set.signers)) == 0 {
		return exclusionListImportSnapshot{}, errors.New("signed exclusion list has no active root signatures")
	}

	obsoleteByActive := s.activeExSet != nil && set.hash != s.activeExSet.hash && set.timestamp <= s.activeExSet.timestamp
	obsoleteByDesired := s.desiredExSet != nil && set.hash != s.desiredExSet.hash && set.timestamp <= s.desiredExSet.timestamp
	obsoleteByProposed := s.proposedExSet != nil && set.hash != s.proposedExSet.hash && set.timestamp <= s.proposedExSet.timestamp
	if obsoleteByActive || obsoleteByDesired || obsoleteByProposed {
		return exclusionListImportSnapshot{}, errProposedExclusionListObsolete
	}

	if s.isExclusionSetInQuarantine(set) {
		return exclusionListImportSnapshot{}, errors.New("signed exclusion list is quarantined")
	}

	if s.getActiveRootSet(true).isEnoughExSetSignatures(set) && s.exclusionSetWouldQuarantine(set) {
		return exclusionListImportSnapshot{}, errors.New("signed exclusion list would be quarantined")
	}

	if s.activeExSet == nil || s.activeExSet.hash != set.hash {
		exceeded, err := s.isSetQuotaExceededDryRun(set)
		if err != nil {
			return exclusionListImportSnapshot{}, err
		}
		if exceeded {
			return exclusionListImportSnapshot{}, errors.New("exclusion list quota exceeded")
		}
	}

	return s.exclusionListImportSnapshotLocked(set.hash), nil
}

func validateExclusionSetRanges(set *exclusionSet) error {
	for address, ranges := range set.blockRanges {
		for _, blockRange := range ranges {
			if !blockRange.IsValid() {
				return errors.Errorf("invalid exclusion list range for %s", address.Hex())
			}
		}
	}
	return nil
}

func (s *RootManager) exclusionSetWouldQuarantine(set *exclusionSet) bool {
	if s.bc == nil {
		return false
	}
	if len(set.addrToBlockRangeExclusiveDiff(s.activeExSet)) == 0 {
		return false
	}
	return s.isExclusionSetMeetsQuarantineCriteria(set.earliestBlockFromDiff(s.activeExSet))
}

func (s *RootManager) exclusionListImportSnapshot(hash common.Hash) exclusionListImportSnapshot {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	return s.exclusionListImportSnapshotLocked(hash)
}

func (s *RootManager) exclusionListImportSnapshotLocked(hash common.Hash) exclusionListImportSnapshot {
	snapshot := exclusionListImportSnapshot{}
	if s.activeExSet != nil {
		snapshot.activeHash = s.activeExSet.hash
		if s.activeExSet.hash == hash {
			snapshot.activeSigners = len(s.activeExSet.signers)
		}
	}
	if s.desiredExSet != nil {
		snapshot.desiredHash = s.desiredExSet.hash
		if s.desiredExSet.hash == hash {
			snapshot.desiredSigners = len(s.desiredExSet.signers)
		}
	}
	if s.proposedExSet != nil {
		snapshot.proposedHash = s.proposedExSet.hash
		if s.proposedExSet.hash == hash {
			snapshot.proposedSigners = len(s.proposedExSet.signers)
		}
	}
	if qSets, err := s.db.getExclusionSetsFromQuarantine(); err == nil {
		snapshot.quarantineSize = len(qSets)
	}
	return snapshot
}

func (h *handler) propagateRootSet(set *rootSet) {
	if set != nil {
		h.rootEventCh <- &rootSetEvent{set: set}
	}
}

func (h *handler) propagateExclusionSet(set *exclusionSet) {
	if set != nil {
		h.exEventCh <- &exclusionSetEvent{set: set}
	}
}

func (h *handler) handleMsg(p *peer) error {
	msg, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}

	if msg.Code != ConstitutionFilesMsg {
		if msg.Size > protocolMaxMsgSize {
			return errMsgTooLarge
		}
	} else {
		if msg.Size > maxConstitutionFileSize {
			return errMsgTooLarge
		}
	}

	defer func() { _ = msg.Discard() }()

	switch msg.Code {
	case StatusMsg:
		return errors.New("unexpected msg code")
	case RootListMsg:
		return h.handleRootListMsg(p, msg)
	case ExclusionListMsg:
		return h.handleExclusionList(p, msg)
	case ApprovalMsg:
		return h.handleWithRetry(p, msg, h.handleApprovalMsg)
	case ConstitutionFileRequestMsg:
		return h.handleConstitutionRequestMsg(p, msg)
	case ConstitutionFilesMsg:
		return h.handleConstitutionFilesMsg(p, msg)
	case KnownConstitutionFilesMsg:
		return h.handleKnownFilesMsg(p, msg)
	case TypedRootListRelayMsg:
		return h.handleTypedRootListRelay(p, msg)
	case TypedExclusionListRelayMsg:
		return h.handleTypedExclusionListRelay(p, msg)
	default:
		return errUnknownMsgCode
	}
}

func (h *handler) handleWithRetry(p *peer, msg p2p.Msg, handlerFunc func(p *peer, msg p2p.Msg) error) error {
	if _, ok := h.failureCounts[p.id]; !ok {
		h.failureCounts[p.id] = make(map[uint64]uint64)
	}

	err := handlerFunc(p, msg)
	if err != nil {
		h.failureCounts[p.id][msg.Code]++
		p.Log().Error("Failed to handle governance message", "err", err)
	} else {
		h.failureCounts[p.id][msg.Code] = 0
	}

	if h.failureCounts[p.id][msg.Code] >= h.maxFailures {
		h.failureCounts[p.id][msg.Code] = 0

		// The error will reset a peer, so release the memory
		delete(h.failureCounts, p.id)

		return err
	}

	return nil
}

func (h *handler) handleRootListMsg(p *peer, msg p2p.Msg) error {
	var list common.RootList
	if err := msg.Decode(&list); err != nil {
		return err
	}

	return h.importRootListFrom(p.id, &list, true)
}

func (h *handler) handleApprovalMsg(p *peer, msg p2p.Msg) error {
	var approval common.RootNodeApprovalList
	if err := msg.Decode(&approval); err != nil {
		return err
	}

	received := approval
	return h.handleIncomingApproval(p, &received)
}

func (h *handler) handleConstitutionRequestMsg(p *peer, msg p2p.Msg) error {
	var request common.ConstitutionFilesRequest
	if err := msg.Decode(&request); err != nil {
		return err
	}
	return h.handleConstitutionFileRequest(p, &request)
}

func (h *handler) handleKnownFilesMsg(p *peer, msg p2p.Msg) error {
	var request common.KnownConstitutionFilesMessage
	if err := msg.Decode(&request); err != nil {
		return err
	}
	return h.handleKnownConstitutionFiles(p, &request)
}

func (h *handler) handleConstitutionFilesMsg(p *peer, msg p2p.Msg) error {
	var response common.ConstitutionFilesResponse
	if err := msg.Decode(&response); err != nil {
		return err
	}

	for _, file := range response.Files {
		if len(file.Data) == 0 {
			log.Debug("Received file with zero length", "hash", file.Hash)
			return errors.New("Received file with zero length")
		}
	}

	h.constitutionManager.storageLock.Lock()
	defer h.constitutionManager.storageLock.Unlock()

	var fulFilledRequests = make(map[common.Hash]struct{})
	var requiredHashes = make(map[common.Hash]struct{})
	for _, hash := range h.constitutionManager.requiredHashes {
		requiredHashes[hash] = struct{}{}
	}

	var totalDecompressedSize int64

	for _, file := range response.Files {
		if _, ok := requiredHashes[file.Hash]; !ok {
			log.Error("Received constitution file with non-requested hash", "hash", file.Hash)
			continue
		}

		reader := bytes.NewReader(file.Data)
		gzreader, e1 := gzip.NewReader(reader)
		if e1 != nil {
			return e1
		}

		// Protect against gzip bombs by limiting decompressed size.
		limitedReader := &io.LimitedReader{
			R: gzreader,
			// Allow one extra byte to reliably detect overflow beyond the limit.
			N: maxConstitutionDecompressedSize + 1,
		}

		output, e2 := io.ReadAll(limitedReader)
		_ = gzreader.Close()
		if e2 != nil {
			return e2
		}

		if int64(len(output)) > maxConstitutionDecompressedSize {
			log.Warn("Received constitution file exceeds decompressed size limit",
				"hash", file.Hash,
				"limit", maxConstitutionDecompressedSize,
				"size", len(output))
			return errors.New("received constitution file exceeds decompressed size limit")
		}

		totalDecompressedSize += int64(len(output))
		if totalDecompressedSize > maxConstitutionTotalDecompressedSize {
			log.Warn("Total decompressed size of constitution files in message exceeds limit",
				"limit", maxConstitutionTotalDecompressedSize,
				"size", totalDecompressedSize)
			return errors.New("total decompressed size of constitution files in message exceeds limit")
		}

		receivedHash := h.constitutionManager.getHashByFileContent(output)

		if receivedHash != file.Hash {
			log.Error("Received file hash doesn't match the requested one", "requested", file.Hash, "got", receivedHash)
			return errors.New("Received file hash doesn't match the requested one")
		}

		cFile := common.ConstitutionFile{
			Name:      h.constitutionManager.filenameFromHash(file.Hash),
			Hash:      file.Hash,
			CreatedAt: time.Now().Unix(),
		}

		// Received file can be the draft (if was requested previously)
		legit, errV := h.constitutionManager.isHashValid(file.Hash)
		if errV != nil {
			// If we receive a constitution, and our registry is not initialized, most likely
			// we haven't synced the state yet
			if errors.Is(errV, vcHasntDeployedErr) {
				break
			}

			return errV
		}

		errStore := h.constitutionManager.storeConstitutionFile(output, cFile, legit)
		if errStore != nil {
			return errStore
		}
		fulFilledRequests[file.Hash] = struct{}{}
	}

	// Store missing hashes to ask new peers about them
	var missingHashes []common.Hash
	for _, hash := range h.constitutionManager.requiredHashes {
		if _, ok := fulFilledRequests[hash]; !ok {
			missingHashes = append(missingHashes, hash)
		}
	}
	if len(missingHashes) != len(h.constitutionManager.requiredHashes) {
		if errSave := h.constitutionManager.db.saveConstitutionFileRequests(missingHashes); errSave != nil {
			return errors.Wrap(errSave, "Failed to save constitution file requests to the database")
		}
	}

	return nil
}

func (h *handler) importRootList(list *common.RootList) error {
	return h.importRootListFrom("", list, false)
}

func (h *handler) importRootListFrom(fromID string, list *common.RootList, signLocal bool) error {
	received, err := newRootSetForNetwork(list, h.rootManager.networkId)
	if err != nil {
		return err
	}
	return h.importRootSet(fromID, received, signLocal)
}

func (h *handler) importRootSet(fromID string, received *rootSet, signLocal bool) error {
	rm := h.rootManager
	rm.rootLock.Lock()
	defer rm.rootLock.Unlock()

	received.updateAliases(rm.getAliasesOfRoots(received.rootAddresses))

	// Check if the received root set is acceptable (spam protection)
	// Skip this check if the received root set is the same as the active one
	if rm.active == nil || rm.active.hash != received.hash {
		exceeded, err := rm.isSetQuotaExceeded(received)
		if err != nil {
			log.Error("Failed to check root set quota", "error", err)
			return nil
		}
		if exceeded {
			log.Warn("Received root set's author has exceeded quota", "hash", received.hash)
			return nil
		}
	}

	switch {
	case rm.active.isAcceptable(received) && (rm.desired == nil || rm.desired.hash != received.hash):
		if signLocal && rm.isRootNode(false) {
			rm.signRootSet(received)
		}

		rm.upgradeRootSet(received)

		h.rootEventCh <- &rootSetEvent{set: received}
	case rm.active.hash == received.hash:
		if signLocal && rm.isRootNode(false) {
			rm.signRootSet(rm.active)
		}
		newSignatures := rm.active.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		log.Debug("Received active root list signatures", "from", fromID, "signers", toSigners(newSignatures))
		h.rootEventCh <- &rootSetEvent{fromID: fromID, set: rm.active.copy()}

		rm.db.saveActiveRootSet(rm.active)
	case rm.desired != nil && rm.desired.hash == received.hash:
		newSignatures := rm.desired.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		log.Debug("Received desired root list signatures", "from", fromID, "signers", toSigners(newSignatures))
		if !rm.active.isAcceptable(rm.desired) {
			h.rootEventCh <- &rootSetEvent{fromID: fromID, set: rm.desired.copy()}
			rm.db.saveDesiredRootSet(rm.desired)
			return nil
		}

		rm.upgradeRootSet(rm.desired)
		h.rootEventCh <- &rootSetEvent{set: rm.active.copy()}
	case rm.proposed != nil && rm.proposed.hash == received.hash:
		newSignatures := rm.proposed.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		log.Debug("Received proposed root list signatures", "from", fromID, "signers", toSigners(newSignatures))
		if !rm.active.isAcceptable(rm.proposed) {
			h.rootEventCh <- &rootSetEvent{fromID: fromID, set: rm.proposed.copy()}
			rm.db.saveProposedRootSet(rm.proposed)
			return nil
		}

		rm.upgradeRootSet(rm.proposed)
		h.rootEventCh <- &rootSetEvent{set: rm.active.copy()}
	default:
		signers := rm.active.knownSigners(received.signers)
		if len(signers) == 0 {
			log.Debug("Ignoring proposed root list: not signed by any known root node",
				"hash", received.hash.Hex(),
				"from", fromID)
			return nil
		}

		if rm.proposed != nil && rm.proposed.timestamp >= received.timestamp {
			log.Debug("Ignoring proposed root list: obsolete", "timestamp", received.timestamp, "hash", received.hash.Hex())
			return nil
		}

		rm.proposed = received
		rm.db.saveProposedRootSet(rm.proposed)
		h.rootEventCh <- &rootSetEvent{fromID: fromID, set: received}
		log.Warn("Received a new proposed root list", "hash", received.hash.Hex(), "timestamp", received.timestamp)
	}

	return nil
}

func (h *handler) handleExclusionList(p *peer, msg p2p.Msg) error {
	var list common.ValidatorExclusionList
	if err := msg.Decode(&list); err != nil {
		return err
	}

	return h.importExclusionListFrom(p.id, &list, true)
}

func (h *handler) importExclusionList(list *common.ValidatorExclusionList) error {
	return h.importExclusionListFrom("", list, false)
}

func (h *handler) importExclusionListFrom(fromID string, list *common.ValidatorExclusionList, signLocal bool) error {
	received, err := newExclusionSetForNetwork(list, h.rootManager.networkId)
	if err != nil {
		return err
	}
	return h.importExclusionSet(fromID, received, signLocal)
}

func (h *handler) importExclusionSet(fromID string, received *exclusionSet, signLocal bool) error {
	if received.blockRanges == nil && received.addrToBlock != nil {
		for address, blockAddress := range received.addrToBlock {
			received.blockRanges[address] = []common.BlockRange{
				{
					StartAddress: blockAddress,
				},
			}
		}
	}
	rm := h.rootManager
	rm.exLock.Lock()
	defer rm.exLock.Unlock()

	rm.active.aliases = rm.getAliasesOfRoots(rm.active.rootAddresses)

	// Check if the received exclusion set is acceptable (spam protection)
	// Skip this check if the received exclusion set is the same as the active one
	if rm.activeExSet == nil || rm.activeExSet.hash != received.hash {
		exceeded, err := rm.isSetQuotaExceeded(received)
		if err != nil {
			log.Error("Failed to check exclusion set quota", "error", err)
			return nil
		}
		if exceeded {
			log.Warn("Received exclusion set's author has exceeded quota", "hash", received.hash)
			return nil
		}
	}

	switch true {
	case rm.isAcceptableExclusionSet(received) && !rm.isExclusionSetInQuarantine(received):
		// ensure locally stored signatures are not lost
		if rm.desiredExSet != nil && rm.desiredExSet.hash == received.hash {
			received.mergeSignatures(rm.desiredExSet.hash, rm.desiredExSet.signers)
		}

		if signLocal && rm.isRootNode(false) {
			rm.signExclusionSet(received)
		}

		rm.upgradeExclusionSet(received, rm.activeExSet == nil)

		h.exEventCh <- &exclusionSetEvent{set: received}
	case rm.activeExSet != nil && rm.activeExSet.hash == received.hash:
		// On very start of the node account can be not unlocked, so isRootNode can return false
		signatureAdded := false // In case when alias changed
		if signLocal && rm.isRootNode(false) {
			signatureAdded = rm.signExclusionSet(rm.activeExSet)
		}
		newSignatures := rm.activeExSet.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 && !signatureAdded {
			return nil
		}

		log.Debug("Received new exclusion list signatures", "from", fromID, "singers", toSigners(newSignatures))
		h.exEventCh <- &exclusionSetEvent{fromID: fromID, set: rm.activeExSet.copy()}
		rm.db.saveActiveExclusionSet(rm.activeExSet)
	case rm.desiredExSet != nil && rm.desiredExSet.hash == received.hash:
		newSignatures := rm.desiredExSet.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		log.Debug("Received new desired exclusion list signatures", "from", fromID, "singers", toSigners(newSignatures))
		if !rm.getActiveRootSet(true).isEnoughExSetSignatures(rm.desiredExSet) {
			h.exEventCh <- &exclusionSetEvent{fromID: fromID, set: rm.desiredExSet}
			rm.db.saveDesiredExclusionSet(rm.desiredExSet)
			return nil
		}

		rm.upgradeExclusionSet(rm.desiredExSet, false)
		h.exEventCh <- &exclusionSetEvent{set: rm.activeExSet}

	case rm.proposedExSet != nil && rm.proposedExSet.hash == received.hash:
		newSignatures := rm.proposedExSet.mergeSignatures(received.hash, received.signers)
		if len(newSignatures) == 0 {
			return nil
		}

		if !rm.isAcceptableExclusionSet(rm.proposedExSet) {
			h.exEventCh <- &exclusionSetEvent{fromID: fromID, set: rm.proposedExSet.copy()}
			rm.db.saveProposedExclusionSet(rm.proposedExSet)
			return nil
		}

		rm.upgradeExclusionSet(rm.proposedExSet, false)
		h.exEventCh <- &exclusionSetEvent{set: rm.activeExSet.copy()}
	default:
		if len(rm.getActiveRootSet(true).knownSigners(received.signers)) == 0 {
			log.Debug("Ignoring proposed exclusion list: no active root node signatures")
			return nil
		}

		obsoleteByActive := rm.activeExSet != nil && received.timestamp <= rm.activeExSet.timestamp
		obsoleteByProposed := rm.proposedExSet != nil && received.timestamp <= rm.proposedExSet.timestamp
		if obsoleteByActive || obsoleteByProposed {
			return nil
		}

		if h.rootManager.isExclusionSetInQuarantine(received) {
			log.Debug("Received proposed exclusion list already quarantined", "hash", received.hash.String())
			return nil
		}

		rm.proposedExSet = received
		rm.db.saveProposedExclusionSet(rm.proposedExSet)
		h.exEventCh <- &exclusionSetEvent{fromID: fromID, set: received}

		log.Warn("Received a new proposed exclusion list", "hash", received.hash.Hex(), "timestamp", received.timestamp)
	}

	return nil
}

func (h *handler) handleIncomingApproval(p *peer, received *common.RootNodeApprovalList) error {
	rm := h.rootManager
	rm.approvalLock.Lock()
	defer rm.approvalLock.Unlock()

	if received == nil {
		return errProposedApprovalListEmpty
	}
	if rm == nil || rm.bc == nil {
		return nil
	}

	if received.BlockNumber.Uint64()%rm.bc.Config().Clique.Epoch != 0 {
		log.Error("Received root node approval list contains invalid block number", "blockNumber", received.BlockNumber)
		return errInvalidApprovalBlockNumber
	}
	if received.BlockNumber.Uint64() == 0 {
		return nil
	}
	if received.Approvals == nil {
		return nil
	}

	exApprovals, errEx := rm.db.getApprovalRecordsByBlockNumber(received.BlockNumber)
	if errEx != nil {
		return errEx
	}

	toPropagate := false
	for _, approval := range received.Approvals {
		pubkey, err := crypto.SigToPub(approval.Hash.Bytes(), approval.Signature)
		if err != nil {
			return errInvalidSignature
		}

		addr := crypto.PubkeyToAddress(*pubkey)

		// TODO Refactor
		ok := false
		for address, alias := range rm.active.aliases {
			if address == addr || alias == addr {
				ok = true
				break
			}
		}

		if !ok {
			continue
		}

		approvalExists := false
		for _, exApproval := range exApprovals {
			if bytes.Equal(approval.Signature, exApproval.Signature) {
				log.Debug("Received root node approval already exists in the DB, skipping", "addr", approval.Signer.String())
				approvalExists = true
				break
			}
		}
		if approvalExists {
			continue
		}

		if errSave := rm.db.saveApprovalRecord(approval); errSave != nil {
			return errSave
		}

		toPropagate = true
	}

	if toPropagate {
		h.approvalEventCh <- &approvalEvent{fromID: p.id, approval: received}
	}

	return nil
}

func (h *handler) handleConstitutionFileRequest(p *peer, received *common.ConstitutionFilesRequest) error {
	cm := h.constitutionManager

	if cm == nil {
		log.Error("Constitution manager is not initialized")
		return nil
	}
	cm.storageLock.Lock()
	defer cm.storageLock.Unlock()

	if received == nil || len(received.Hashes) == 0 {
		return nil
	}

	exFiles, err := cm.db.getConstitutionFiles()
	if err != nil {
		return err
	}

	drafts, err := cm.getDraftFiles()
	if err != nil {
		log.Error("Error getting draft files", "error", err)
	}

	var presentFiles []common.ConstitutionFile
	for _, hash := range received.Hashes {
		ok, errV := h.constitutionManager.isHashValid(hash)
		if errV != nil {
			// If we receive a constitution, and our registry is not initialized, most likely
			// we haven't synced the state yet
			if errors.Is(errV, vcHasntDeployedErr) {
				break
			}

			return errV
		}

		// If the requested file is not on the list but node has it, it answers anyway (this will allow for draft constitution to be stored in the file system)
		if !ok {
			log.Debug("Requested file hash doesn't belong to history")
		}

		foundInFiles := false
		for _, exFile := range exFiles {
			if exFile.Hash == hash {
				presentFiles = append(presentFiles, exFile)
				foundInFiles = true
			}
		}

		if !foundInFiles {
			for _, dFile := range drafts {
				if dFile.Hash == hash {
					presentFiles = append(presentFiles, dFile)
				}
			}
		}
	}

	if len(presentFiles) > 0 {
		// Actually run asynchronously to avoid blocking the handler
		go p.asyncSendConstitutionFiles(cm, presentFiles) // TODO check constitution drafts
	}

	return nil
}

func (h *handler) handleKnownConstitutionFiles(p *peer, received *common.KnownConstitutionFilesMessage) error {
	cm := h.constitutionManager

	cm.storageLock.Lock()
	defer cm.storageLock.Unlock()

	if received == nil || len(received.Hashes) == 0 {
		return nil
	}

	exFiles, err := cm.db.getKnownConstitutionFiles()
	if err != nil {
		return err
	}
	if len(exFiles) == 0 {
		return nil
	}

	var exFilesSet = make(map[common.Hash]struct{})
	for _, file := range exFiles {
		exFilesSet[file] = struct{}{}
	}

	resNewFiles := make([]common.Hash, 0)
	for _, hash := range received.Hashes {
		if _, ok := exFilesSet[hash]; !ok {
			resNewFiles = append(resNewFiles, hash)
			log.Info("New constitution file found. You can request it by using command: gov.requestForConstitutionFile(\"" + hash.String() + "\")")
		}
	}
	if len(resNewFiles) > 0 {
		err = cm.updateKnownConstitutionFiles(resNewFiles)
		if err != nil {
			if errors.Is(err, vcHasntDeployedErr) {
				return nil
			}

			return err
		}
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
