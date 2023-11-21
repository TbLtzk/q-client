package governance

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	rand2 "math/rand"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/pkg/errors"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/p2p"
)

type ByteSize uint64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
)

func TestHandshake(t *testing.T) {
	rm := newTestRootManager(t, false, true)
	cm, errCm := NewConstitutionManager(os.TempDir(), rm.db, rm.RootManager)
	if errCm != nil {
		log.Error("Can't create ConstitutionManager: %v", errCm)
	}

	tests := []struct {
		name string
		code uint64
		msg  interface{}
		err  error
	}{
		{
			name: "Wrong network id",
			code: StatusMsg,
			msg: statusMsgBody{
				CurrentRootList:       rm.active.copy().makeList(),
				DesiredRootList:       rm.desired.copy().makeList(),
				ProposedRootList:      rm.proposed.copy().makeList(),
				CurrentExclusionList:  rm.activeExSet.copy().makeList(),
				DesiredExclusionList:  rm.desiredExSet.copy().makeList(),
				ProposedExclusionList: rm.proposedExSet.copy().makeList(),
				Network:               1,
			},
			err: errors.New("invalid network id"),
		}, {
			name: "Empty status body",
			code: StatusMsg,
			msg:  statusMsgBody{},
			err:  errInvalidNetworkId,
		}, {
			name: "invalid message",
			code: RootListMsg,
			msg:  "wrong message",
			err:  errors.New("invalid message: (code 1) (size 14) rlp: expected input list for governance.statusMsgBody"),
		}, {
			name: "Success hadshake",
			code: StatusMsg,
			msg: statusMsgBody{
				CurrentRootList:       rm.active.copy().makeList(),
				DesiredRootList:       rm.desired.copy().makeList(),
				ProposedRootList:      rm.proposed.copy().makeList(),
				CurrentExclusionList:  rm.activeExSet.copy().makeList(),
				DesiredExclusionList:  rm.desiredExSet.copy().makeList(),
				ProposedExclusionList: rm.proposedExSet.copy().makeList(),
				Network:               123456,
			},
			err: nil, //Timeout
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := handshake(rm.RootManager, cm, test.code, test.msg)
			if (test.err != nil && err == nil) || (test.err == nil && err != nil) || (test.err != nil && err != nil && err.Error() != test.err.Error()) {
				t.Errorf("Expected error '%v', got '%v'", test.err, err)
			}
		})
	}
}

func TestHandleMessage(t *testing.T) {
	rm := newTestRootManager(t, false, true)
	cm, errCm := NewConstitutionManager(os.TempDir(), rm.db, rm.RootManager)
	if errCm != nil {
		log.Error("Can't create ConstitutionManager: %v", errCm)
	}
	h := newHandler(rm.RootManager, cm)
	rw1, rw2 := p2p.MsgPipe()
	p1 := newPeer(2, p2p.NewPeer(randomPeerID(), "peer", nil), rw1)
	defer p1.Disconnect(p2p.DiscRequested)

	tests := []struct {
		name string
		code uint64
		msg  interface{}
		err  error
	}{
		{
			name: "Unknown message",
			code: 0x123,
			msg:  "any",
			err:  errUnknownMsgCode,
		}, {
			name: "Too long message",
			code: StatusMsg,
			msg:  randString(10000000),
			err:  errMsgTooLarge,
		}, {
			name: "Large root list",
			code: RootListMsg,
			msg:  randomRootList(t, nil, time.Now().Unix(), 2*maxNRootNodes+1, 2*maxNRootNodes+1, false),
			err:  errMsgTooLarge,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := handleMsg(h, test.code, test.msg, p1, rw2)
			if !errors.Is(err, test.err) {
				t.Errorf("Expected error '%v', got '%v'", test.err, err)
			}
		})
	}
}

func TestHandleRootSet(t *testing.T) {
	rm := newTestRootManager(t, false, true)
	bc := newTestChain(t, rm.RootManager)
	defer bc.Stop()
	rm.InitBlockChain(bc)

	gov, err := New(rm.RootManager, tmpDirName(t))
	if err != nil {
		t.Fatalf("Failed to create Governance: %v", err)
	}
	err = gov.Start()
	if err != nil {
		t.Fatalf("Failed to start Governance: %v", err)
	}

	rw1, rw2 := p2p.MsgPipe()
	p1 := newPeer(2, p2p.NewPeer(randomPeerID(), "peer", nil), rw1)
	defer p1.Disconnect(p2p.DiscRequested)

	startRootList := rm.active.copy()

	listWithWrongHash := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 3, defNumAccounts, true)
	listWithWrongHash.Hash = common.BytesToHash(randBytes(1))

	tests := []struct {
		name         string
		code         uint64
		msg          common.RootList
		err          error
		shouldAccept bool
	}{
		{
			name:         "Root list with 100% signatures",
			code:         RootListMsg,
			msg:          randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, defNumAccounts, true),
			err:          nil,
			shouldAccept: true,
		},
		{
			name:         "Root list with 100% signatures, signed by main account",
			code:         RootListMsg,
			msg:          randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, defNumAccounts, false),
			err:          nil,
			shouldAccept: false,
		},
		{
			name:         "Root list with 75% signatures",
			code:         RootListMsg,
			msg:          randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, defNumAccounts-1, true),
			err:          nil,
			shouldAccept: true,
		},
		{
			name:         "Root list with not enough signatures",
			code:         RootListMsg,
			msg:          randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true),
			err:          nil,
			shouldAccept: false,
		},
		{
			name:         "Root list with wrong signatures",
			code:         RootListMsg,
			msg:          randomRootList(t, nil, time.Now().Add(5*time.Minute).Unix(), 10, 1, false),
			err:          nil,
			shouldAccept: false,
		},
		{
			name:         "Large root list",
			code:         RootListMsg,
			msg:          randomRootList(t, nil, time.Now().Unix(), 2*maxNRootNodes+1, 2*maxNRootNodes+1, false),
			err:          errMsgTooLarge,
			shouldAccept: false,
		},
		{
			name:         "Root list with wrong hash",
			code:         RootListMsg,
			msg:          listWithWrongHash,
			err:          errHashMismatch,
			shouldAccept: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := handleMsg(gov.handler, test.code, test.msg, p1, rw2)
			if !errors.Is(err, test.err) {
				t.Errorf("Expected error '%v', got '%v'", test.err, err)
			}
			if test.shouldAccept && test.msg.Hash != rm.active.hash {
				t.Error("List wasn't accepted")
			}
			if !test.shouldAccept && rm.active.hash != startRootList.hash {
				t.Error("List was accepted")
			}
		})
		rm.active = startRootList
	}
}

func TestHandleExclusionSet(t *testing.T) {
	rm := newTestRootManager(t, false, true)
	bc := newTestChain(t, rm.RootManager)
	defer bc.Stop()
	rm.InitBlockChain(bc)

	gov, err := New(rm.RootManager, tmpDirName(t))
	if err != nil {
		t.Fatalf("Failed to create Governance: %v", err)
	}
	err = gov.Start()
	if err != nil {
		t.Fatalf("Failed to start Governance: %v", err)
	}

	rw1, rw2 := p2p.MsgPipe()
	p1 := newPeer(2, p2p.NewPeer(randomPeerID(), "peer", nil), rw1)
	defer p1.Disconnect(p2p.DiscRequested)

	listWithWrongHash := randomExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 3, defNumAccounts, true)
	listWithWrongHash.Hash = common.BytesToHash(randBytes(1))

	tests := []struct {
		name         string
		code         uint64
		msg          common.ValidatorExclusionList
		err          error
		shouldAccept bool
	}{
		{
			name:         "Exclusion list with 100% signatures",
			code:         ExclusionListMsg,
			msg:          randomExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 3, defNumAccounts, true),
			err:          nil,
			shouldAccept: true,
		},
		{
			name:         "Exclusion list with 100% signatures, signed by main accounts",
			code:         ExclusionListMsg,
			msg:          randomExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 3, defNumAccounts, false),
			err:          nil,
			shouldAccept: false,
		},
		{
			name:         "Exclusion list with 75% signatures",
			code:         ExclusionListMsg,
			msg:          randomExclusionList(t, rm, uint64(time.Now().Add(5*time.Hour).Unix()), 3, defNumAccounts, true),
			err:          nil,
			shouldAccept: true,
		},
		{
			name:         "Exclusion list with not enough signatures",
			code:         ExclusionListMsg,
			msg:          randomExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 3, 1, true),
			err:          nil,
			shouldAccept: false,
		},
		{
			name:         "Large exclusion list",
			code:         ExclusionListMsg,
			msg:          randomExclusionList(t, nil, uint64(time.Now().Add(5*time.Minute).Unix()), 2*maxNRootNodes+1, 2*maxNRootNodes+1, true),
			err:          errMsgTooLarge,
			shouldAccept: false,
		},
		{
			name:         "Exclusion list with wrong hash",
			code:         ExclusionListMsg,
			msg:          listWithWrongHash,
			err:          errHashMismatch,
			shouldAccept: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := handleMsg(gov.handler, test.code, test.msg, p1, rw2)
			if !errors.Is(err, test.err) {
				t.Errorf("Expected error '%v', got '%v'", test.err, err)
			}
			if test.shouldAccept && test.msg.Hash != rm.activeExSet.hash {
				t.Error("List wasn't accepted")
			}
			if (!test.shouldAccept && rm.activeExSet != nil && rm.activeExSet.hash != common.Hash{}) {
				t.Error("List was accepted")
			}
		})
		rm.activeExSet = nil
	}
}

func TestHandleConstitution(t *testing.T) {
	rm := newTestRootManager(t, false, true)
	bc := newTestChain(t, rm.RootManager)
	defer bc.Stop()
	rm.InitBlockChain(bc)

	gov, err := New(rm.RootManager, tmpDirName(t))
	if err != nil {
		t.Fatalf("Failed to create Governance: %v", err)
	}
	err = gov.Start()
	if err != nil {
		t.Fatalf("Failed to start Governance: %v", err)
	}

	rw1, rw2 := p2p.MsgPipe()
	p1 := newPeer(2, p2p.NewPeer(randomPeerID(), "peer", nil), rw1)
	defer p1.Disconnect(p2p.DiscRequested)

	tests := []struct {
		name               string
		code               uint64
		filesNum           int
		minFileSize        int
		maxFileSize        int
		msg                interface{}
		requestHashesCount int
		err                error
	}{
		{
			name:        "Receive constitution files",
			code:        ConstitutionFilesMsg,
			filesNum:    3,
			minFileSize: int(2 * KB),
			maxFileSize: int(3 * KB),
			msg:         common.ConstitutionFilesResponse{},
			err:         nil,
		},
		{
			name:        "Receive big file",
			code:        ConstitutionFilesMsg,
			filesNum:    1,
			minFileSize: maxConstitutionFileSize,
			maxFileSize: maxConstitutionFileSize + 1,
			msg:         common.ConstitutionFilesResponse{},
			err:         errMsgTooLarge,
		},
		{
			name:               "Request constitution files",
			code:               ConstitutionFileRequestMsg,
			filesNum:           10,
			minFileSize:        int(2 * MB),
			maxFileSize:        int(3 * MB),
			msg:                common.ConstitutionFilesRequest{},
			requestHashesCount: 5,
			err:                nil,
		},
		{
			name:               "Request a lot of files",
			code:               ConstitutionFileRequestMsg,
			filesNum:           1,
			minFileSize:        int(1 * KB),
			maxFileSize:        int(2 * KB),
			msg:                common.ConstitutionFilesRequest{},
			requestHashesCount: 100000,
			err:                errMsgTooLarge,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gov.ConstitutionManager.requiredHashes = gov.ConstitutionManager.requiredHashes[:0]
			gov.ConstitutionManager.knownHashes = gov.ConstitutionManager.knownHashes[:0]

			files, contents := randomConstitutionFiles(test.filesNum, test.minFileSize, test.maxFileSize)

			switch test.msg.(type) {
			case common.ConstitutionFilesRequest:
				for i := 0; i < test.filesNum; i++ {
					err := gov.ConstitutionManager.storeConstitutionFile(contents[i].Data, files[i], true)
					if err != nil {
						t.Errorf("Failed to save constitution file: %v", err)
					}
				}

				hashesToRequest := make([]common.Hash, 0, test.requestHashesCount)

				if test.requestHashesCount <= len(files) {
					for i := 0; i < test.requestHashesCount; i++ {
						hashesToRequest = append(hashesToRequest, files[i].Hash)
					}
				}
				diff := test.requestHashesCount - len(files)
				if diff > 0 {
					for i := 0; i < diff; i++ {
						hashesToRequest = append(hashesToRequest, common.BytesToHash(randBytes(32)))
					}
				}

				test.msg = common.ConstitutionFilesRequest{
					Hashes: hashesToRequest,
				}

			case common.ConstitutionFilesResponse:
				test.msg = common.ConstitutionFilesResponse{Files: contents}

			default:
				t.Error("Wrong msg type")
			}

			err := handleMsg(gov.handler, test.code, test.msg, p1, rw2)
			if !errors.Is(err, test.err) {
				t.Errorf("Expected error '%v', got '%v'", test.err, err)
			}
		})
		rm.activeExSet = nil
	}
}

func TestApproval(t *testing.T) {
	rm := newTestRootManager(t, false, true)
	bc := newTestChain(t, rm.RootManager)
	defer bc.Stop()
	rm.InitBlockChain(bc)

	gov, err := New(rm.RootManager, tmpDirName(t))
	if err != nil {
		t.Fatalf("Failed to create Governance: %v", err)
	}
	err = gov.Start()
	if err != nil {
		t.Fatalf("Failed to start Governance: %v", err)
	}

	rw1, rw2 := p2p.MsgPipe()
	p1 := newPeer(2, p2p.NewPeer(randomPeerID(), "peer", nil), rw1)
	defer p1.Disconnect(p2p.DiscRequested)

	tests := []struct {
		name string
		code uint64
		msg  common.RootNodeApprovalList
		err  error
	}{
		{
			name: "Approval list with empty hash",
			code: ApprovalMsg,
			msg:  randomApprovalList(t, rm, common.Hash{}, int64(rm.bc.Config().Clique.Epoch), defNumAccounts),
			err:  nil,
		},
		{
			name: "Approval on non-epoch transition block",
			code: ApprovalMsg,
			msg:  randomApprovalList(t, rm, common.Hash{}, int64(rm.bc.Config().Clique.Epoch)+1, defNumAccounts),
			err:  errInvalidApprovalBlockNumber,
		},
		{
			name: "Approval msg is too large",
			code: ApprovalMsg,
			msg:  randomApprovalList(t, rm, common.Hash{}, int64(rm.bc.Config().Clique.Epoch), maxNRootNodes*2+1),
			err:  errMsgTooLarge,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := handleMsg(gov.handler, test.code, test.msg, p1, rw2)
			if !errors.Is(err, test.err) {
				t.Errorf("Expected error '%v', got '%v'", test.err, err)
			}
		})
	}
}

func randomApprovalList(t *testing.T, rm *TestRootManager, hash common.Hash, blockNum int64, approvalsCount int) common.RootNodeApprovalList {
	privateKeys := make([]*ecdsa.PrivateKey, 0)
	if rm != nil && approvalsCount <= len(rm.rootPrivateKeys) {
		privateKeys = append(privateKeys, rm.rootPrivateKeys[:approvalsCount]...)
	} else {
		for i := 0; i < approvalsCount; i++ {
			privateKey, err := crypto.GenerateKey()
			if err != nil {
				t.Error(errors.Wrap(err, "failed to generate random private key"))
			}
			privateKeys = append(privateKeys, privateKey)
		}
	}

	approvals := make([]common.RootNodeApproval, 0, approvalsCount)
	for i := 0; i < approvalsCount; i++ {
		apprBlockNum, err := rand.Int(rand.Reader, big.NewInt(blockNum))
		if err != nil {
			t.Errorf("Failed to generate random block number: %v", err)
		}

		sig, err := crypto.Sign(hash.Bytes(), privateKeys[i])
		if err != nil {
			t.Error(errors.Wrap(err, "failed to sign hash"))
		}

		approvals = append(approvals, common.RootNodeApproval{
			BlockNumber: apprBlockNum,
			Hash:        hash,
			Signer:      crypto.PubkeyToAddress(privateKeys[i].PublicKey),
			Signature:   sig,
		})
	}

	return common.RootNodeApprovalList{
		BlockNumber: big.NewInt(blockNum),
		Approvals:   approvals,
	}
}

func randomConstitutionFiles(count int, minSize, maxSize int) ([]common.ConstitutionFile, []common.ConstitutionFileContent) {
	files := make([]common.ConstitutionFile, 0, count)
	contents := make([]common.ConstitutionFileContent, 0, count)

	for i := 0; i < count; i++ {
		hash := common.BytesToHash(randBytes(32))

		files = append(files, common.ConstitutionFile{
			Name:      randString(10),
			Hash:      hash,
			CreatedAt: int64(rand2.Int()),
		})

		contents = append(contents, common.ConstitutionFileContent{
			Hash: hash,
			Data: randBytes(rand2.Intn(maxSize-minSize) + minSize),
		})
	}

	return files, contents
}

func handleMsg(h *handler, code uint64, msg interface{}, p1 *peer, rw2 *p2p.MsgPipeRW) error {
	errCh := make(chan error)

	go func() {
		err := p2p.Send(rw2, code, msg)
		if err != nil {
			errCh <- err
		}
	}()

	go func() {
		err := h.handleMsg(p1)
		errCh <- err
	}()

	for {
		select {
		case err := <-errCh:
			return err
		case <-time.After(5 * time.Second):
			return errors.New("timeout")
		}
	}
}

func handshake(rm *RootManager, cm *ConstitutionManager, code uint64, msg interface{}) error {
	h := newHandler(rm, cm)

	p1, rw1, rw2, errCh := makeHandshake(h, code, msg)
	defer rw1.Close()
	defer rw2.Close()

	for {
		select {
		case errIn := <-errCh:
			return errIn
		case <-time.After(1 * time.Second):
			p1.Disconnect(p2p.DiscRequested)
			return nil
		}
	}
}

func makeHandshake(h *handler, code uint64, msg interface{}) (*peer, *p2p.MsgPipeRW, *p2p.MsgPipeRW, chan error) {
	rw1, rw2 := p2p.MsgPipe()

	p1 := newPeer(2, p2p.NewPeer(randomPeerID(), "peer", nil), rw1)

	errCh := make(chan error)

	go func() {
		err := h.runPeer(p1)
		if err != nil {
			errCh <- err
		}
	}()

	go func() {
		err := p2p.Send(rw2, code, msg)
		if err != nil {
			errCh <- err
		}
	}()

	go func() {
		msgR, err := rw2.ReadMsg()
		if err != nil {
			errCh <- err
		}
		var data statusMsgBody
		err = msgR.Decode(&data)
		if err != nil {
			errCh <- err
		}
	}()

	return p1, rw1, rw2, errCh
}

func randString(n int) string {
	return string(randBytes(n))
}

func randBytes(n int) []byte {
	const alphanum = "0123456789ABCDEF"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return bytes
}

func randomRootList(t *testing.T, rm *TestRootManager, timestamp int64, nodesCount, signersCount int, signByAlias bool) common.RootList {
	rootList := common.RootList{
		Timestamp:  uint64(timestamp),
		Nodes:      make([]common.Address, 0, nodesCount),
		Signatures: make([][]byte, 0, signersCount),
	}
	for i := 0; i < nodesCount; i++ {
		rootList.Nodes = append(rootList.Nodes, common.BytesToAddress(randBytes(common.AddressLength)))
	}

	sort.SliceStable(rootList.Nodes, func(i, j int) bool {
		return bytes.Compare(rootList.Nodes[i].Bytes(), rootList.Nodes[j].Bytes()) > 0
	})
	msg := append([]byte{}, fmt.Sprint(rootList.Timestamp)...)
	for _, addr := range rootList.Nodes {
		msg = append(msg, addr.Bytes()...)
	}
	rootList.Hash = crypto.Keccak256Hash(msg)

	privateKeys := make([]*ecdsa.PrivateKey, 0)
	if rm != nil && signersCount <= len(rm.rootPrivateKeys) {
		if signByAlias {
			privateKeys = append(privateKeys, rm.aliasPrivateKeys[:signersCount]...)
		} else {
			privateKeys = append(privateKeys, rm.rootPrivateKeys[:signersCount]...)
		}
	} else {
		for i := 0; i < signersCount; i++ {
			privateKey, err := crypto.GenerateKey()
			if err != nil {
				t.Error(errors.Wrap(err, "failed to generate random private key"))
			}
			privateKeys = append(privateKeys, privateKey)
		}
	}

	for _, key := range privateKeys {
		sig, err := crypto.Sign(rootList.Hash.Bytes(), key)
		if err != nil {
			t.Error(errors.Wrap(err, "failed to sign hash"))
		}
		rootList.Signatures = append(rootList.Signatures, sig)
	}

	return rootList
}

func randomExclusionList(t *testing.T, rm *TestRootManager, timestamp uint64, validatorsCount, signersCount int, signByAlias bool) common.ValidatorExclusionList {
	exclusionList := common.ValidatorExclusionList{
		Timestamp:  timestamp,
		Validators: make([]common.ExcludedValidator, 0, validatorsCount),
		Signatures: make([][]byte, 0, signersCount),
	}
	for i := 0; i < validatorsCount; i++ {
		exclusionList.Validators = append(exclusionList.Validators, common.ExcludedValidator{
			Address:  common.BytesToAddress(randBytes(common.AddressLength)),
			Block:    rand2.Uint64(),
			EndBlock: rand2.Uint64(),
		})
	}

	exclusionSet, err := newExclusionSet(&exclusionList)
	if err != nil {
		t.Error(err)
	}
	exclusionList.Hash = exclusionSet.hash

	privateKeys := make([]*ecdsa.PrivateKey, 0)
	if rm != nil && signersCount <= len(rm.rootPrivateKeys) {
		if signByAlias {
			privateKeys = append(privateKeys, rm.aliasPrivateKeys[:signersCount]...)
		} else {
			privateKeys = append(privateKeys, rm.rootPrivateKeys[:signersCount]...)
		}
	} else {
		for i := 0; i < signersCount; i++ {
			privateKey, err := crypto.GenerateKey()
			if err != nil {
				t.Error(errors.Wrap(err, "failed to generate random private key"))
			}
			privateKeys = append(privateKeys, privateKey)
		}
	}

	for _, key := range privateKeys {
		sig, err := crypto.Sign(exclusionList.Hash.Bytes(), key)
		if err != nil {
			t.Error(errors.Wrap(err, "failed to sign hash"))
		}
		exclusionList.Signatures = append(exclusionList.Signatures, sig)
	}

	return exclusionList
}
