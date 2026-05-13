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

	"gitlab.com/q-dev/q-client/accounts"
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
	p1 := newPeer(5, p2p.NewPeer(randomPeerID(), "peer", nil), rw1)
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
	p1 := newPeer(5, p2p.NewPeer(randomPeerID(), "peer", nil), rw1)
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

func TestImportRootList(t *testing.T) {
	tests := []struct {
		name              string
		list              func(t *testing.T, rm *TestRootManager) common.RootList
		beforeImport      func(t *testing.T, rm *TestRootManager, list common.RootList)
		assertAfterImport func(t *testing.T, rm *TestRootManager, list common.RootList)
	}{
		{
			name: "known signer stores proposed list",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
			},
			assertAfterImport: func(t *testing.T, rm *TestRootManager, list common.RootList) {
				if rm.proposed == nil || rm.proposed.hash != list.Hash {
					t.Fatalf("expected proposed root list %s, got %v", list.Hash.Hex(), rm.proposed)
				}
				if rm.active.hash == list.Hash {
					t.Fatal("single-signature root list was upgraded")
				}
			},
		},
		{
			name: "unknown signer is ignored",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, nil, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
			},
			assertAfterImport: func(t *testing.T, rm *TestRootManager, list common.RootList) {
				if rm.proposed != nil {
					t.Fatalf("expected unknown signer root list to be ignored, got %s", rm.proposed.hash.Hex())
				}
			},
		},
		{
			name: "duplicate signature is ignored",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				list := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
				list.Signatures = append(list.Signatures, list.Signatures[0])
				return list
			},
			assertAfterImport: func(t *testing.T, rm *TestRootManager, list common.RootList) {
				if rm.proposed == nil || rm.proposed.hash != list.Hash {
					t.Fatalf("expected proposed root list %s, got %v", list.Hash.Hex(), rm.proposed)
				}
				if len(rm.proposed.signers) != 1 {
					t.Fatalf("expected one unique signature, got %d", len(rm.proposed.signers))
				}
			},
		},
		{
			name: "obsolete timestamp is ignored",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
			},
			beforeImport: func(t *testing.T, rm *TestRootManager, list common.RootList) {
				newer := randomRootList(t, rm, int64(list.Timestamp+1), 10, 1, true)
				set, err := newRootSet(&newer)
				if err != nil {
					t.Fatal(err)
				}
				set.updateAliases(rm.getAliasesOfRoots(set.rootAddresses))
				rm.proposed = set
			},
			assertAfterImport: func(t *testing.T, rm *TestRootManager, list common.RootList) {
				if rm.proposed == nil || rm.proposed.hash == list.Hash {
					t.Fatal("obsolete root list replaced newer proposal")
				}
			},
		},
		{
			name: "threshold signatures upgrade active list",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, defNumAccounts-1, true)
			},
			assertAfterImport: func(t *testing.T, rm *TestRootManager, list common.RootList) {
				if rm.active.hash != list.Hash {
					t.Fatalf("expected active root list %s, got %s", list.Hash.Hex(), rm.active.hash.Hex())
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := newTestRootManager(t, false, true)
			gov, err := New(rm.RootManager, tmpDirName(t))
			if err != nil {
				t.Fatalf("Failed to create Governance: %v", err)
			}
			if err := gov.Start(); err != nil {
				t.Fatalf("Failed to start Governance: %v", err)
			}

			list := tt.list(t, rm)
			if tt.beforeImport != nil {
				tt.beforeImport(t, rm, list)
			}
			if err := gov.handler.importRootList(&list); err != nil {
				t.Fatalf("importRootList returned error: %v", err)
			}
			tt.assertAfterImport(t, rm, list)
		})
	}
}

func TestImportRootListDoesNotRequireLocalSigning(t *testing.T) {
	rm := newTestRootManager(t, true, false)
	gov, err := New(rm.RootManager, tmpDirName(t))
	if err != nil {
		t.Fatalf("Failed to create Governance: %v", err)
	}
	if err := gov.Start(); err != nil {
		t.Fatalf("Failed to start Governance: %v", err)
	}

	list := common.RootList{
		Timestamp: uint64(time.Now().Add(5 * time.Minute).Unix()),
		Nodes: []common.Address{
			common.BytesToAddress(randBytes(common.AddressLength)),
			common.BytesToAddress(randBytes(common.AddressLength)),
			common.BytesToAddress(randBytes(common.AddressLength)),
		},
	}
	set, err := newRootSet(&list)
	if err != nil {
		t.Fatal(err)
	}
	signature, err := rm.SignHash(accounts.Account{Address: rm.active.rootAddresses[0]}, set.hash.Bytes())
	if err != nil {
		t.Fatalf("failed to sign root list: %v", err)
	}
	list.Hash = set.hash
	list.Signatures = [][]byte{signature}

	if err := gov.handler.importRootList(&list); err != nil {
		t.Fatalf("importRootList returned error: %v", err)
	}

	if rm.active.hash == list.Hash {
		t.Fatal("import locally signed the root list and upgraded it")
	}
	if rm.proposed == nil || rm.proposed.hash != list.Hash {
		t.Fatalf("expected imported root list to be proposed, got %v", rm.proposed)
	}
	if len(rm.proposed.signers) != 1 {
		t.Fatalf("expected import to keep only supplied signature, got %d signatures", len(rm.proposed.signers))
	}
}

func TestSubmitSignedRootList(t *testing.T) {
	tests := []struct {
		name              string
		list              func(t *testing.T, rm *TestRootManager) common.RootList
		beforeSubmit      func(t *testing.T, rm *TestRootManager, list common.RootList)
		assertAfterSubmit func(t *testing.T, rm *TestRootManager, list common.RootList, hash common.Hash, err error)
	}{
		{
			name: "disabled submissions are rejected",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
			},
			beforeSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList) {
				rm.DisablePublicSubmission = true
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList, hash common.Hash, err error) {
				if !errors.Is(err, errPublicGovernanceSubmissionDisabled) {
					t.Fatalf("expected disabled submission error, got hash %s err %v", hash.Hex(), err)
				}
				if rm.proposed != nil {
					t.Fatalf("disabled submission changed proposed state: %v", rm.proposed)
				}
			},
		},
		{
			name: "oversized payload is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, nil, time.Now().Add(5*time.Minute).Unix(), 2*maxNRootNodes+1, 2*maxNRootNodes+1, true)
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList, hash common.Hash, err error) {
				if !errors.Is(err, errMsgTooLarge) {
					t.Fatalf("expected oversized root list error, got hash %s err %v", hash.Hex(), err)
				}
				if rm.proposed != nil {
					t.Fatalf("oversized root list changed proposed state: %v", rm.proposed)
				}
			},
		},
		{
			name: "quota exceeded is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
			},
			beforeSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList) {
				rm.ProposalQuotaMax = 1
				rm.ProposalQuotaTimeWindow = time.Hour
				set, err := newRootSet(&list)
				if err != nil {
					t.Fatal(err)
				}
				rm.rootQuotaEntries[rm.active.knownSigners(set.signers)[0]] = []common.ListQuotaEntry{
					{Hash: common.BytesToHash(randBytes(common.HashLength)), Timestamp: uint64(time.Now().Unix())},
				}
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList, hash common.Hash, err error) {
				if err == nil {
					t.Fatalf("expected quota error, got hash %s", hash.Hex())
				}
				if rm.proposed != nil {
					t.Fatalf("quota rejected root list changed proposed state: %v", rm.proposed)
				}
			},
		},
		{
			name: "known signer stores proposed list",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList, hash common.Hash, err error) {
				if err != nil {
					t.Fatalf("SubmitSignedRootList returned error: %v", err)
				}
				if hash != list.Hash {
					t.Fatalf("expected hash %s, got %s", list.Hash.Hex(), hash.Hex())
				}
				if rm.proposed == nil || rm.proposed.hash != list.Hash {
					t.Fatalf("expected proposed root list %s, got %v", list.Hash.Hex(), rm.proposed)
				}
			},
		},
		{
			name: "threshold signatures upgrade active list",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, defNumAccounts-1, true)
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList, hash common.Hash, err error) {
				if err != nil {
					t.Fatalf("SubmitSignedRootList returned error: %v", err)
				}
				if hash != list.Hash {
					t.Fatalf("expected hash %s, got %s", list.Hash.Hex(), hash.Hex())
				}
				if rm.active.hash != list.Hash {
					t.Fatalf("expected active root list %s, got %s", list.Hash.Hex(), rm.active.hash.Hex())
				}
			},
		},
		{
			name: "malformed hash is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				list := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
				list.Hash = common.BytesToHash(randBytes(1))
				return list
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList, hash common.Hash, err error) {
				if !errors.Is(err, errHashMismatch) {
					t.Fatalf("expected hash mismatch, got hash %s err %v", hash.Hex(), err)
				}
				if rm.proposed != nil {
					t.Fatalf("malformed root list changed proposed state: %v", rm.proposed)
				}
			},
		},
		{
			name: "invalid signature is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				list := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
				list.Signatures[0] = randBytes(crypto.SignatureLength)
				return list
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList, hash common.Hash, err error) {
				if !errors.Is(err, errInvalidSignature) {
					t.Fatalf("expected invalid signature, got hash %s err %v", hash.Hex(), err)
				}
				if rm.proposed != nil {
					t.Fatalf("invalid signature changed proposed state: %v", rm.proposed)
				}
			},
		},
		{
			name: "unknown signer is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, nil, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList, hash common.Hash, err error) {
				if err == nil {
					t.Fatalf("expected unknown signer rejection, got hash %s", hash.Hex())
				}
				if rm.proposed != nil {
					t.Fatalf("unknown signer changed proposed state: %v", rm.proposed)
				}
			},
		},
		{
			name: "obsolete timestamp is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
			},
			beforeSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList) {
				newer := randomRootList(t, rm, int64(list.Timestamp+1), 10, 1, true)
				set, err := newRootSet(&newer)
				if err != nil {
					t.Fatal(err)
				}
				set.updateAliases(rm.getAliasesOfRoots(set.rootAddresses))
				rm.proposed = set
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList, hash common.Hash, err error) {
				if err == nil {
					t.Fatalf("expected obsolete root list rejection, got hash %s", hash.Hex())
				}
				if rm.proposed == nil || rm.proposed.hash == list.Hash {
					t.Fatal("obsolete root list replaced newer proposal")
				}
			},
		},
		{
			name: "duplicate is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
			},
			beforeSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList) {
				set, err := newRootSet(&list)
				if err != nil {
					t.Fatal(err)
				}
				set.updateAliases(rm.getAliasesOfRoots(set.rootAddresses))
				rm.proposed = set
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.RootList, hash common.Hash, err error) {
				if err == nil {
					t.Fatalf("expected duplicate root list rejection, got hash %s", hash.Hex())
				}
				if rm.proposed == nil || rm.proposed.hash != list.Hash || len(rm.proposed.signers) != 1 {
					t.Fatalf("duplicate root list changed proposed state: %v", rm.proposed)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := newTestRootManager(t, false, true)
			gov, err := New(rm.RootManager, tmpDirName(t))
			if err != nil {
				t.Fatalf("Failed to create Governance: %v", err)
			}
			if err := gov.Start(); err != nil {
				t.Fatalf("Failed to start Governance: %v", err)
			}
			api := NewGovernancePublicAPI(gov)

			list := tt.list(t, rm)
			if tt.beforeSubmit != nil {
				tt.beforeSubmit(t, rm, list)
			}
			hash, err := api.SubmitSignedRootList(list)
			tt.assertAfterSubmit(t, rm, list, hash, err)
		})
	}
}

func TestSubmitSignedRootListEquivalentToPeerImport(t *testing.T) {
	tests := []struct {
		name           string
		list           func(t *testing.T, rm *TestRootManager) common.RootList
		beforeImport   func(t *testing.T, rm *TestRootManager, list common.RootList)
		wantRPCError   bool
		compareStates  bool
		compareNoState bool
	}{
		{
			name: "valid proposal",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
			},
			compareStates: true,
		},
		{
			name: "threshold upgrade",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, defNumAccounts-1, true)
			},
			compareStates: true,
		},
		{
			name: "duplicate proposal",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
			},
			beforeImport: func(t *testing.T, rm *TestRootManager, list common.RootList) {
				set, err := newRootSet(&list)
				if err != nil {
					t.Fatal(err)
				}
				set.updateAliases(rm.getAliasesOfRoots(set.rootAddresses))
				rm.proposed = set
			},
			wantRPCError:   true,
			compareNoState: true,
		},
		{
			name: "stale proposal",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
			},
			beforeImport: func(t *testing.T, rm *TestRootManager, list common.RootList) {
				newer := randomRootList(t, rm, int64(list.Timestamp+1), 10, 1, true)
				set, err := newRootSet(&newer)
				if err != nil {
					t.Fatal(err)
				}
				set.updateAliases(rm.getAliasesOfRoots(set.rootAddresses))
				rm.proposed = set
			},
			wantRPCError:   true,
			compareNoState: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peerRM := newTestRootManager(t, false, true)
			rpcRM := newTestRootManager(t, false, true)

			list := tt.list(t, peerRM)
			cloneRootManagerRootState(rpcRM.RootManager, peerRM.RootManager)
			if tt.beforeImport != nil {
				tt.beforeImport(t, peerRM, list)
				tt.beforeImport(t, rpcRM, list)
			}

			peerGov, err := New(peerRM.RootManager, tmpDirName(t))
			if err != nil {
				t.Fatalf("Failed to create peer Governance: %v", err)
			}
			rpcGov, err := New(rpcRM.RootManager, tmpDirName(t))
			if err != nil {
				t.Fatalf("Failed to create rpc Governance: %v", err)
			}
			if err := peerGov.Start(); err != nil {
				t.Fatalf("Failed to start peer Governance: %v", err)
			}
			if err := rpcGov.Start(); err != nil {
				t.Fatalf("Failed to start rpc Governance: %v", err)
			}

			peerBefore := peerRM.rootListImportSnapshot(list.Hash)
			rpcBefore := rpcRM.rootListImportSnapshot(list.Hash)
			if err := peerGov.handler.importRootListFrom("peer", &list, false); err != nil {
				t.Fatalf("peer import failed: %v", err)
			}
			_, rpcErr := NewGovernancePublicAPI(rpcGov).SubmitSignedRootList(list)
			if tt.wantRPCError && rpcErr == nil {
				t.Fatal("expected RPC import error")
			}
			if !tt.wantRPCError && rpcErr != nil {
				t.Fatalf("RPC import failed: %v", rpcErr)
			}

			if tt.compareStates {
				assertRootImportStateEqual(t, peerRM.RootManager, rpcRM.RootManager, list.Hash)
			}
			if tt.compareNoState {
				if peerRM.rootListImportSnapshot(list.Hash).changedFrom(peerBefore) {
					t.Fatal("peer path mutated state for ignored root list")
				}
				if rpcRM.rootListImportSnapshot(list.Hash).changedFrom(rpcBefore) {
					t.Fatal("RPC path mutated state for ignored root list")
				}
			}
		})
	}
}

func TestSubmitTypedSignedRootList(t *testing.T) {
	makeTypedRootList := func(t *testing.T, rm *TestRootManager, signByAlias bool) common.RootList {
		t.Helper()

		list := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 0, true)
		unsigned := list
		unsigned.Signatures = nil

		payloadHash := common.NewRootListSigningPayloadV1(rm.networkId, unsigned).Hash()

		var key *ecdsa.PrivateKey
		if signByAlias {
			key = rm.aliasPrivateKeys[0]
		} else {
			key = rm.rootPrivateKeys[0]
		}
		signature, err := crypto.Sign(payloadHash.Bytes(), key)
		if err != nil {
			t.Fatalf("failed to sign typed root payload: %v", err)
		}
		list.Signatures = [][]byte{signature}
		return list
	}

	makeTypedRootListWithChainID := func(t *testing.T, rm *TestRootManager, chainID uint64) common.RootList {
		t.Helper()

		list := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 0, true)
		unsigned := list
		unsigned.Signatures = nil

		payloadHash := common.NewRootListSigningPayloadV1(chainID, unsigned).Hash()
		signature, err := crypto.Sign(payloadHash.Bytes(), rm.aliasPrivateKeys[0])
		if err != nil {
			t.Fatalf("failed to sign typed root payload: %v", err)
		}
		list.Signatures = [][]byte{signature}
		return list
	}

	makeTypedRootListWithAlteredPayload := func(t *testing.T, rm *TestRootManager) common.RootList {
		t.Helper()

		list := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 0, true)
		unsigned := list
		unsigned.Signatures = nil
		unsigned.Timestamp++

		payloadHash := common.NewRootListSigningPayloadV1(rm.networkId, unsigned).Hash()
		signature, err := crypto.Sign(payloadHash.Bytes(), rm.aliasPrivateKeys[0])
		if err != nil {
			t.Fatalf("failed to sign typed root payload: %v", err)
		}
		list.Signatures = [][]byte{signature}
		return list
	}

	makeTypedRootListWithWrongProposalType := func(t *testing.T, rm *TestRootManager) common.RootList {
		t.Helper()

		list := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 0, true)
		unsigned := list
		unsigned.Signatures = nil

		payload := common.NewRootListSigningPayloadV1(rm.networkId, unsigned)
		payload.Metadata.ProposalType = common.GovernanceProposalTypeExclusionList
		signature, err := crypto.Sign(payload.Hash().Bytes(), rm.aliasPrivateKeys[0])
		if err != nil {
			t.Fatalf("failed to sign typed root payload: %v", err)
		}
		list.Signatures = [][]byte{signature}
		return list
	}

	tests := []struct {
		name    string
		list    func(t *testing.T, rm *TestRootManager) common.RootList
		wantErr bool
	}{
		{
			name: "valid typed signature from active alias is accepted",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return makeTypedRootList(t, rm, true)
			},
		},
		{
			name: "typed signature from root key instead of alias is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return makeTypedRootList(t, rm, false)
			},
			wantErr: true,
		},
		{
			name: "typed signature with wrong chain id is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return makeTypedRootListWithChainID(t, rm, rm.networkId+1)
			},
			wantErr: true,
		},
		{
			name: "typed signature over altered payload fields is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return makeTypedRootListWithAlteredPayload(t, rm)
			},
			wantErr: true,
		},
		{
			name: "typed signature with wrong proposal type is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.RootList {
				return makeTypedRootListWithWrongProposalType(t, rm)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := newTestRootManager(t, false, true)
			gov, err := New(rm.RootManager, tmpDirName(t))
			if err != nil {
				t.Fatalf("Failed to create Governance: %v", err)
			}
			if err := gov.Start(); err != nil {
				t.Fatalf("Failed to start Governance: %v", err)
			}

			list := tt.list(t, rm)
			hash, err := NewGovernancePublicAPI(gov).SubmitTypedSignedRootList(list)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got hash %s", hash.Hex())
				}
				if rm.proposed != nil && rm.proposed.hash == list.Hash {
					t.Fatalf("rejected typed root list changed proposed state: %v", rm.proposed)
				}
				return
			}

			if err != nil {
				t.Fatalf("SubmitTypedSignedRootList returned error: %v", err)
			}
			if hash != list.Hash {
				t.Fatalf("expected hash %s, got %s", list.Hash.Hex(), hash.Hex())
			}
			if rm.proposed == nil || rm.proposed.hash != list.Hash {
				t.Fatalf("expected proposed root list %s, got %v", list.Hash.Hex(), rm.proposed)
			}
		})
	}
}

func TestSubmitTypedSignedExclusionList(t *testing.T) {
	makeTypedExclusionList := func(t *testing.T, rm *TestRootManager, signByAlias bool) common.ValidatorExclusionList {
		t.Helper()

		list := signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
		unsigned := list
		unsigned.Signatures = nil

		payloadHash := common.NewExclusionListSigningPayloadV1(rm.networkId, unsigned).Hash()

		var key *ecdsa.PrivateKey
		if signByAlias {
			key = rm.aliasPrivateKeys[0]
		} else {
			key = rm.rootPrivateKeys[0]
		}
		signature, err := crypto.Sign(payloadHash.Bytes(), key)
		if err != nil {
			t.Fatalf("failed to sign typed exclusion payload: %v", err)
		}
		list.Signatures = [][]byte{signature}
		return list
	}

	makeTypedExclusionListWithChainID := func(t *testing.T, rm *TestRootManager, chainID uint64) common.ValidatorExclusionList {
		t.Helper()

		list := signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
		unsigned := list
		unsigned.Signatures = nil

		payloadHash := common.NewExclusionListSigningPayloadV1(chainID, unsigned).Hash()
		signature, err := crypto.Sign(payloadHash.Bytes(), rm.aliasPrivateKeys[0])
		if err != nil {
			t.Fatalf("failed to sign typed exclusion payload: %v", err)
		}
		list.Signatures = [][]byte{signature}
		return list
	}

	makeTypedExclusionListWithAlteredTimestamp := func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
		t.Helper()

		list := signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
		unsigned := list
		unsigned.Signatures = nil
		unsigned.Timestamp++

		payloadHash := common.NewExclusionListSigningPayloadV1(rm.networkId, unsigned).Hash()
		signature, err := crypto.Sign(payloadHash.Bytes(), rm.aliasPrivateKeys[0])
		if err != nil {
			t.Fatalf("failed to sign typed exclusion payload: %v", err)
		}
		list.Signatures = [][]byte{signature}
		return list
	}

	makeTypedExclusionListWithWrongProposalType := func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
		t.Helper()

		list := signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
		unsigned := list
		unsigned.Signatures = nil

		payload := common.NewExclusionListSigningPayloadV1(rm.networkId, unsigned)
		payload.Metadata.ProposalType = common.GovernanceProposalTypeRootList
		signature, err := crypto.Sign(payload.Hash().Bytes(), rm.aliasPrivateKeys[0])
		if err != nil {
			t.Fatalf("failed to sign typed exclusion payload: %v", err)
		}
		list.Signatures = [][]byte{signature}
		return list
	}

	makeTypedExclusionListWithAlteredValidators := func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
		t.Helper()

		list := signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
		unsigned := list
		unsigned.Signatures = nil
		unsigned.Validators = append([]common.ExcludedValidator(nil), list.Validators...)
		unsigned.Validators[0].EndBlock++

		payloadHash := common.NewExclusionListSigningPayloadV1(rm.networkId, unsigned).Hash()
		signature, err := crypto.Sign(payloadHash.Bytes(), rm.aliasPrivateKeys[0])
		if err != nil {
			t.Fatalf("failed to sign typed exclusion payload: %v", err)
		}
		list.Signatures = [][]byte{signature}
		return list
	}

	tests := []struct {
		name    string
		list    func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList
		wantErr bool
	}{
		{
			name: "valid typed signature from active alias is accepted",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return makeTypedExclusionList(t, rm, true)
			},
		},
		{
			name: "typed signature from root key instead of alias is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return makeTypedExclusionList(t, rm, false)
			},
			wantErr: true,
		},
		{
			name: "typed signature with wrong chain id is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return makeTypedExclusionListWithChainID(t, rm, rm.networkId+1)
			},
			wantErr: true,
		},
		{
			name: "typed signature over altered timestamp is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return makeTypedExclusionListWithAlteredTimestamp(t, rm)
			},
			wantErr: true,
		},
		{
			name: "typed signature with wrong proposal type is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return makeTypedExclusionListWithWrongProposalType(t, rm)
			},
			wantErr: true,
		},
		{
			name: "typed signature over altered validator ranges is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return makeTypedExclusionListWithAlteredValidators(t, rm)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := newTestRootManager(t, false, true)
			gov, err := New(rm.RootManager, tmpDirName(t))
			if err != nil {
				t.Fatalf("Failed to create Governance: %v", err)
			}
			if err := gov.Start(); err != nil {
				t.Fatalf("Failed to start Governance: %v", err)
			}

			list := tt.list(t, rm)
			hash, err := NewGovernancePublicAPI(gov).SubmitTypedSignedExclusionList(list)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got hash %s", hash.Hex())
				}
				if rm.proposedExSet != nil && rm.proposedExSet.hash == list.Hash {
					t.Fatalf("rejected typed exclusion list changed proposed state: %v", rm.proposedExSet)
				}
				return
			}

			if err != nil {
				t.Fatalf("SubmitTypedSignedExclusionList returned error: %v", err)
			}
			if hash != list.Hash {
				t.Fatalf("expected hash %s, got %s", list.Hash.Hex(), hash.Hex())
			}
			if rm.proposedExSet == nil || rm.proposedExSet.hash != list.Hash {
				t.Fatalf("expected proposed exclusion list %s, got %v", list.Hash.Hex(), rm.proposedExSet)
			}
		})
	}
}

func cloneRootManagerRootState(dst, src *RootManager) {
	dst.active = src.active.copy()
	dst.desired = src.desired.copy()
	dst.proposed = src.proposed.copy()
	dst.reg = newTestRegistry()
	dst.rootQuotaEntries = map[common.Address][]common.ListQuotaEntry{}
	for signer, entries := range src.rootQuotaEntries {
		dst.rootQuotaEntries[signer] = append([]common.ListQuotaEntry(nil), entries...)
	}
}

func assertRootImportStateEqual(t *testing.T, peerRM, rpcRM *RootManager, hash common.Hash) {
	t.Helper()
	if peerRM.rootListImportSnapshot(hash) != rpcRM.rootListImportSnapshot(hash) {
		t.Fatalf("root import state mismatch: peer=%+v rpc=%+v", peerRM.rootListImportSnapshot(hash), rpcRM.rootListImportSnapshot(hash))
	}
}

func assertExclusionImportStateEqual(t *testing.T, peerRM, rpcRM *RootManager, hash common.Hash) {
	t.Helper()
	if peerRM.exclusionListImportSnapshot(hash) != rpcRM.exclusionListImportSnapshot(hash) {
		t.Fatalf("exclusion import state mismatch: peer=%+v rpc=%+v", peerRM.exclusionListImportSnapshot(hash), rpcRM.exclusionListImportSnapshot(hash))
	}
}

func TestSubmitSignedExclusionListEquivalentToPeerImport(t *testing.T) {
	tests := []struct {
		name           string
		initChain      bool
		list           func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList
		beforeImport   func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList)
		wantRPCError   bool
		compareStates  bool
		compareNoState bool
		compareQuarant bool
	}{
		{
			name: "valid proposal",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
			},
			compareStates: true,
		},
		{
			name: "threshold upgrade",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedActiveExclusionList(t, rm, defNumAccounts-1, true)
			},
			compareStates: true,
		},
		{
			name: "duplicate proposal",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
			},
			beforeImport: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList) {
				set, err := newExclusionSet(&list)
				if err != nil {
					t.Fatal(err)
				}
				rm.proposedExSet = set
			},
			wantRPCError:   true,
			compareNoState: true,
		},
		{
			name: "stale proposal",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
			},
			beforeImport: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList) {
				newer := signedExclusionList(t, rm, list.Timestamp+1, 2, 1, true, 7000)
				set, err := newExclusionSet(&newer)
				if err != nil {
					t.Fatal(err)
				}
				rm.proposedExSet = set
			},
			wantRPCError:   true,
			compareNoState: true,
		},
		{
			name:      "quarantine",
			initChain: true,
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				rm.setMaxRewindLimit(100)
				list := rm.activeExSet.copy().makeList()
				list.Timestamp++
				list.Signatures = nil
				list.Hash = common.Hash{}
				list.Validators = append(list.Validators, common.ExcludedValidator{
					Address: common.HexToAddress("0x2B01035cDa82a02fb135EBd68676Fa17FdcAD365"),
					Block:   2,
				})
				set, err := newExclusionSet(&list)
				if err != nil {
					t.Fatal(err)
				}
				rm.signExclusionSet(set)
				return set.makeList()
			},
			wantRPCError:   true,
			compareQuarant: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peerRM := newTestRootManager(t, false, true)
			rpcRM := newTestRootManager(t, false, true)
			if tt.initChain {
				peerRM = newTestRootManager(t, true, false)
				rpcRM = newTestRootManager(t, true, false)
				peerBC := newTestChain(t, peerRM.RootManager)
				defer peerBC.Stop()
				rpcBC := newTestChain(t, rpcRM.RootManager)
				defer rpcBC.Stop()
				peerRM.InitBlockChain(peerBC)
				rpcRM.InitBlockChain(rpcBC)
			}

			list := tt.list(t, peerRM)
			cloneRootManagerRootState(rpcRM.RootManager, peerRM.RootManager)
			rpcRM.activeExSet = peerRM.activeExSet.copy()
			rpcRM.desiredExSet = peerRM.desiredExSet.copy()
			rpcRM.proposedExSet = peerRM.proposedExSet.copy()
			rpcRM.rewindLimit = peerRM.rewindLimit
			if tt.beforeImport != nil {
				tt.beforeImport(t, peerRM, list)
				tt.beforeImport(t, rpcRM, list)
			}

			peerGov, err := New(peerRM.RootManager, tmpDirName(t))
			if err != nil {
				t.Fatalf("Failed to create peer Governance: %v", err)
			}
			rpcGov, err := New(rpcRM.RootManager, tmpDirName(t))
			if err != nil {
				t.Fatalf("Failed to create rpc Governance: %v", err)
			}
			if err := peerGov.Start(); err != nil {
				t.Fatalf("Failed to start peer Governance: %v", err)
			}
			if err := rpcGov.Start(); err != nil {
				t.Fatalf("Failed to start rpc Governance: %v", err)
			}

			peerBefore := peerRM.exclusionListImportSnapshot(list.Hash)
			rpcBefore := rpcRM.exclusionListImportSnapshot(list.Hash)
			if err := peerGov.handler.importExclusionListFrom("peer", &list, false); err != nil {
				t.Fatalf("peer import failed: %v", err)
			}
			_, rpcErr := NewGovernancePublicAPI(rpcGov).SubmitSignedExclusionList(list)
			if tt.wantRPCError && rpcErr == nil {
				t.Fatal("expected RPC import error")
			}
			if !tt.wantRPCError && rpcErr != nil {
				t.Fatalf("RPC import failed: %v", rpcErr)
			}

			if tt.compareStates {
				assertExclusionImportStateEqual(t, peerRM.RootManager, rpcRM.RootManager, list.Hash)
			}
			if tt.compareNoState {
				if peerRM.exclusionListImportSnapshot(list.Hash).changedFrom(peerBefore) {
					t.Fatal("peer path mutated state for ignored exclusion list")
				}
				if rpcRM.exclusionListImportSnapshot(list.Hash).changedFrom(rpcBefore) {
					t.Fatal("RPC path mutated state for ignored exclusion list")
				}
			}
			if tt.compareQuarant {
				peerChanged := peerRM.exclusionListImportSnapshot(list.Hash).changedFrom(peerBefore)
				rpcChanged := rpcRM.exclusionListImportSnapshot(list.Hash).changedFrom(rpcBefore)
				if !peerChanged || rpcChanged {
					t.Fatalf("expected peer quarantine mutation and RPC rejection, peerChanged=%v rpcChanged=%v", peerChanged, rpcChanged)
				}
			}
		})
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
	p1 := newPeer(5, p2p.NewPeer(randomPeerID(), "peer", nil), rw1)
	defer p1.Disconnect(p2p.DiscRequested)

	listWithWrongHash := randomExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 3, defNumAccounts, true)
	// Calculate the correct hash before changing it
	correctHash := listWithWrongHash.Hash
	listWithWrongHash.Hash = common.BytesToHash(randBytes(1))
	correctHashPtr := &correctHash

	tests := []struct {
		name         string
		code         uint64
		msg          common.ValidatorExclusionList
		err          error
		shouldAccept bool
		expectedHash *common.Hash // Optional: expected hash if different from msg.Hash (e.g., for wrong hash test)
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
			err:          nil,            // Hash mismatch now logs a warning instead of returning an error
			shouldAccept: true,           // List is accepted if signatures are valid (signatures are validated against calculated hash)
			expectedHash: correctHashPtr, // The accepted hash should be the calculated hash, not the wrong provided hash
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := handleMsg(gov.handler, test.code, test.msg, p1, rw2)
			if !errors.Is(err, test.err) {
				t.Errorf("Expected error '%v', got '%v'", test.err, err)
			}
			if test.shouldAccept {
				expectedHash := test.msg.Hash
				if test.expectedHash != nil {
					expectedHash = *test.expectedHash
				}
				if rm.activeExSet == nil || rm.activeExSet.hash != expectedHash {
					t.Errorf("List wasn't accepted or has wrong hash. Expected: %s, Got: %v", expectedHash.Hex(), func() string {
						if rm.activeExSet == nil {
							return "nil"
						}
						return rm.activeExSet.hash.Hex()
					}())
				}
			}
			if (!test.shouldAccept && rm.activeExSet != nil && rm.activeExSet.hash != common.Hash{}) {
				t.Error("List was accepted")
			}
		})
		rm.activeExSet = nil
	}
}

func TestImportExclusionList(t *testing.T) {
	tests := []struct {
		name              string
		initChain         bool
		list              func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList
		beforeImport      func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList)
		assertAfterImport func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList)
	}{
		{
			name: "known signer stores proposed list",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
			},
			assertAfterImport: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList) {
				if rm.proposedExSet == nil || rm.proposedExSet.hash != list.Hash {
					t.Fatalf("expected proposed exclusion list %s, got %v", list.Hash.Hex(), rm.proposedExSet)
				}
				if rm.activeExSet.hash == list.Hash {
					t.Fatal("single-signature exclusion list was upgraded")
				}
			},
		},
		{
			name: "unknown signer is ignored",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedExclusionList(t, nil, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
			},
			assertAfterImport: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList) {
				if rm.proposedExSet != nil {
					t.Fatalf("expected unknown signer exclusion list to be ignored, got %s", rm.proposedExSet.hash.Hex())
				}
			},
		},
		{
			name: "duplicate signature is ignored",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				list := signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
				list.Signatures = append(list.Signatures, list.Signatures[0])
				return list
			},
			assertAfterImport: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList) {
				if rm.proposedExSet == nil || rm.proposedExSet.hash != list.Hash {
					t.Fatalf("expected proposed exclusion list %s, got %v", list.Hash.Hex(), rm.proposedExSet)
				}
				if len(rm.proposedExSet.signers) != 1 {
					t.Fatalf("expected one unique signature, got %d", len(rm.proposedExSet.signers))
				}
			},
		},
		{
			name: "obsolete timestamp is ignored",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
			},
			beforeImport: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList) {
				newer := signedExclusionList(t, rm, list.Timestamp+1, 2, 1, true, 7000)
				set, err := newExclusionSet(&newer)
				if err != nil {
					t.Fatal(err)
				}
				rm.proposedExSet = set
			},
			assertAfterImport: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList) {
				if rm.proposedExSet == nil || rm.proposedExSet.hash == list.Hash {
					t.Fatal("obsolete exclusion list replaced newer proposal")
				}
			},
		},
		{
			name: "threshold signatures upgrade active list",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedActiveExclusionList(t, rm, defNumAccounts-1, true)
			},
			assertAfterImport: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList) {
				if rm.activeExSet == nil || rm.activeExSet.hash != list.Hash {
					t.Fatalf("expected active exclusion list %s, got %v", list.Hash.Hex(), rm.activeExSet)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := newTestRootManager(t, false, true)
			if tt.initChain {
				bc := newTestChain(t, rm.RootManager)
				defer bc.Stop()
				rm.InitBlockChain(bc)
			}
			gov, err := New(rm.RootManager, tmpDirName(t))
			if err != nil {
				t.Fatalf("Failed to create Governance: %v", err)
			}
			if err := gov.Start(); err != nil {
				t.Fatalf("Failed to start Governance: %v", err)
			}

			list := tt.list(t, rm)
			if tt.beforeImport != nil {
				tt.beforeImport(t, rm, list)
			}
			if err := gov.handler.importExclusionList(&list); err != nil {
				t.Fatalf("importExclusionList returned error: %v", err)
			}
			tt.assertAfterImport(t, rm, list)
		})
	}
}

func TestImportExclusionListDoesNotRequireLocalSigning(t *testing.T) {
	rm := newTestRootManager(t, true, false)
	gov, err := New(rm.RootManager, tmpDirName(t))
	if err != nil {
		t.Fatalf("Failed to create Governance: %v", err)
	}
	if err := gov.Start(); err != nil {
		t.Fatalf("Failed to start Governance: %v", err)
	}

	list := signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, false, 6000)
	if err := gov.handler.importExclusionList(&list); err != nil {
		t.Fatalf("importExclusionList returned error: %v", err)
	}

	if rm.activeExSet != nil && rm.activeExSet.hash == list.Hash {
		t.Fatal("import locally signed the exclusion list and upgraded it")
	}
	if rm.proposedExSet == nil || rm.proposedExSet.hash != list.Hash {
		t.Fatalf("expected imported exclusion list to be proposed, got %v", rm.proposedExSet)
	}
	if len(rm.proposedExSet.signers) != 1 {
		t.Fatalf("expected import to keep only supplied signature, got %d signatures", len(rm.proposedExSet.signers))
	}
}

func TestImportExclusionListPreservesQuarantine(t *testing.T) {
	earliestBlock := uint64(2)
	rm := newTestRootManager(t, true, false)
	bc := newTestChain(t, rm.RootManager)
	defer bc.Stop()
	rm.InitBlockChain(bc)
	rm.setMaxRewindLimit(100)

	list := rm.activeExSet.copy().makeList()
	list.Timestamp++
	list.Signatures = nil
	list.Hash = common.Hash{}
	list.Validators = append(list.Validators, common.ExcludedValidator{
		Address: common.HexToAddress("0x2B01035cDa82a02fb135EBd68676Fa17FdcAD365"),
		Block:   earliestBlock,
	})
	set, err := newExclusionSet(&list)
	if err != nil {
		t.Fatalf("Can't create new test exclusion set: %v", err)
	}
	rm.signExclusionSet(set)
	list = set.makeList()

	gov, err := New(rm.RootManager, tmpDirName(t))
	if err != nil {
		t.Fatalf("Failed to create Governance: %v", err)
	}
	if err := gov.Start(); err != nil {
		t.Fatalf("Failed to start Governance: %v", err)
	}

	currentExclusionSet := rm.activeExSet.copy()
	if err := gov.handler.importExclusionList(&list); err != nil {
		t.Fatalf("importExclusionList returned error: %v", err)
	}
	if currentExclusionSet.hash != rm.activeExSet.hash {
		t.Fatalf("active exclusion set changed for quarantined import")
	}
	qSets, err := rm.db.getExclusionSetsFromQuarantine()
	if err != nil {
		t.Fatalf("Failed to get exclusion sets from quarantine: %v", err)
	}
	if len(qSets) != 1 || qSets[0].hash != set.hash {
		t.Fatalf("expected quarantined exclusion set %s, got %v", set.hash.Hex(), qSets)
	}
}

func TestSubmitSignedExclusionList(t *testing.T) {
	tests := []struct {
		name              string
		initChain         bool
		list              func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList
		beforeSubmit      func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList)
		assertAfterSubmit func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList, hash common.Hash, err error)
		wantErr           bool
	}{
		{
			name: "disabled submissions are rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
			},
			beforeSubmit: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList) {
				rm.DisablePublicSubmission = true
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList, hash common.Hash, err error) {
				if !errors.Is(err, errPublicGovernanceSubmissionDisabled) {
					t.Fatalf("expected disabled submission error, got hash %s err %v", hash.Hex(), err)
				}
				if rm.proposedExSet != nil {
					t.Fatalf("disabled submission changed proposed state: %v", rm.proposedExSet)
				}
			},
		},
		{
			name: "oversized payload is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return randomExclusionList(t, nil, uint64(time.Now().Add(5*time.Minute).Unix()), 2*maxNRootNodes+1, 2*maxNRootNodes+1, true)
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList, hash common.Hash, err error) {
				if !errors.Is(err, errMsgTooLarge) {
					t.Fatalf("expected oversized exclusion list error, got hash %s err %v", hash.Hex(), err)
				}
				if rm.proposedExSet != nil {
					t.Fatalf("oversized exclusion list changed proposed state: %v", rm.proposedExSet)
				}
			},
		},
		{
			name: "quota exceeded is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
			},
			beforeSubmit: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList) {
				rm.ProposalQuotaMax = 1
				rm.ProposalQuotaTimeWindow = time.Hour
				set, err := newExclusionSet(&list)
				if err != nil {
					t.Fatal(err)
				}
				rm.exclusionQuotaEntries[rm.getActiveRootSet(true).knownSigners(set.signers)[0]] = []common.ListQuotaEntry{
					{Hash: common.BytesToHash(randBytes(common.HashLength)), Timestamp: uint64(time.Now().Unix())},
				}
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList, hash common.Hash, err error) {
				if err == nil {
					t.Fatalf("expected quota error, got hash %s", hash.Hex())
				}
				if rm.proposedExSet != nil {
					t.Fatalf("quota rejected exclusion list changed proposed state: %v", rm.proposedExSet)
				}
			},
		},
		{
			name: "known signer stores proposed list",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList, hash common.Hash, err error) {
				if err != nil {
					t.Fatalf("SubmitSignedExclusionList returned error: %v", err)
				}
				if hash != list.Hash {
					t.Fatalf("expected submitted hash %s, got %s", list.Hash.Hex(), hash.Hex())
				}
				if rm.proposedExSet == nil || rm.proposedExSet.hash != list.Hash {
					t.Fatalf("expected proposed exclusion list %s, got %v", list.Hash.Hex(), rm.proposedExSet)
				}
			},
		},
		{
			name: "threshold signatures upgrade active list",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedActiveExclusionList(t, rm, defNumAccounts-1, true)
			},
			assertAfterSubmit: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList, hash common.Hash, err error) {
				if err != nil {
					t.Fatalf("SubmitSignedExclusionList returned error: %v", err)
				}
				if rm.activeExSet == nil || rm.activeExSet.hash != list.Hash {
					t.Fatalf("expected active exclusion list %s, got %v", list.Hash.Hex(), rm.activeExSet)
				}
			},
		},
		{
			name: "unknown signer is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedExclusionList(t, nil, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
			},
			wantErr: true,
		},
		{
			name: "unsigned list is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				list := signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
				list.Signatures = nil
				return list
			},
			wantErr: true,
		},
		{
			name: "malformed signature is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				list := signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
				list.Signatures[0] = randBytes(crypto.SignatureLength)
				return list
			},
			wantErr: true,
		},
		{
			name: "invalid range is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				list := signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
				list.Validators[0].EndBlock = list.Validators[0].Block
				list.Hash = common.Hash{}
				list.Signatures = nil
				set, err := newExclusionSet(&list)
				if err != nil {
					t.Fatal(err)
				}
				rm.signExclusionSet(set)
				return set.makeList()
			},
			wantErr: true,
		},
		{
			name: "obsolete timestamp is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
			},
			beforeSubmit: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList) {
				newer := signedExclusionList(t, rm, list.Timestamp+1, 2, 1, true, 7000)
				set, err := newExclusionSet(&newer)
				if err != nil {
					t.Fatal(err)
				}
				rm.proposedExSet = set
			},
			wantErr: true,
		},
		{
			name: "duplicate is rejected",
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				return signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
			},
			beforeSubmit: func(t *testing.T, rm *TestRootManager, list common.ValidatorExclusionList) {
				set, err := newExclusionSet(&list)
				if err != nil {
					t.Fatal(err)
				}
				rm.proposedExSet = set
			},
			wantErr: true,
		},
		{
			name:      "quarantined list is rejected",
			initChain: true,
			list: func(t *testing.T, rm *TestRootManager) common.ValidatorExclusionList {
				rm.setMaxRewindLimit(100)
				list := rm.activeExSet.copy().makeList()
				list.Timestamp++
				list.Signatures = nil
				list.Hash = common.Hash{}
				list.Validators = append(list.Validators, common.ExcludedValidator{
					Address: common.HexToAddress("0x2B01035cDa82a02fb135EBd68676Fa17FdcAD365"),
					Block:   2,
				})
				set, err := newExclusionSet(&list)
				if err != nil {
					t.Fatal(err)
				}
				rm.signExclusionSet(set)
				return set.makeList()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := newTestRootManager(t, false, true)
			if tt.initChain {
				rm = newTestRootManager(t, true, false)
				bc := newTestChain(t, rm.RootManager)
				defer bc.Stop()
				rm.InitBlockChain(bc)
			}
			gov, err := New(rm.RootManager, tmpDirName(t))
			if err != nil {
				t.Fatalf("Failed to create Governance: %v", err)
			}
			if err := gov.Start(); err != nil {
				t.Fatalf("Failed to start Governance: %v", err)
			}

			api := NewGovernancePublicAPI(gov)
			list := tt.list(t, rm)
			if tt.beforeSubmit != nil {
				tt.beforeSubmit(t, rm, list)
			}
			hash, err := api.SubmitSignedExclusionList(list)
			if tt.assertAfterSubmit != nil {
				tt.assertAfterSubmit(t, rm, list, hash, err)
				return
			}
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected SubmitSignedExclusionList error")
				}
				return
			}
			if err != nil {
				t.Fatalf("SubmitSignedExclusionList returned error: %v", err)
			}
			tt.assertAfterSubmit(t, rm, list, hash, err)
		})
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
	p1 := newPeer(5, p2p.NewPeer(randomPeerID(), "peer", nil), rw1)
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
	p1 := newPeer(5, p2p.NewPeer(randomPeerID(), "peer", nil), rw1)
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

	p1 := newPeer(5, p2p.NewPeer(randomPeerID(), "peer", nil), rw1)

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

func signedExclusionList(t *testing.T, rm *TestRootManager, timestamp uint64, validatorsCount, signersCount int, signByAlias bool, blockStart uint64) common.ValidatorExclusionList {
	exclusionList := common.ValidatorExclusionList{
		Timestamp:  timestamp,
		Validators: make([]common.ExcludedValidator, 0, validatorsCount),
		Signatures: make([][]byte, 0, signersCount),
	}
	for i := 0; i < validatorsCount; i++ {
		block := blockStart + uint64(i)
		exclusionList.Validators = append(exclusionList.Validators, common.ExcludedValidator{
			Address:  common.BytesToAddress(randBytes(common.AddressLength)),
			Block:    block,
			EndBlock: block + 100,
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
	} else if rm != nil && signersCount <= len(rm.RootManager.manager.Accounts()) {
		for _, account := range rm.RootManager.manager.Accounts()[:signersCount] {
			signature, err := rm.SignHash(accounts.Account{Address: account}, exclusionList.Hash.Bytes())
			if err != nil {
				t.Error(errors.Wrap(err, "failed to sign hash"))
			}
			exclusionList.Signatures = append(exclusionList.Signatures, signature)
		}
		return exclusionList
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

func signedActiveExclusionListWithTimestamp(t *testing.T, rm *TestRootManager, timestamp uint64, signersCount int, signByAlias bool) common.ValidatorExclusionList {
	list := rm.activeExSet.copy().makeList()
	list.Timestamp = timestamp
	list.Signatures = nil
	list.Hash = common.Hash{}
	set, err := newExclusionSet(&list)
	if err != nil {
		t.Fatal(err)
	}

	privateKeys := make([]*ecdsa.PrivateKey, 0, signersCount)
	if signByAlias {
		privateKeys = append(privateKeys, rm.aliasPrivateKeys[:signersCount]...)
	} else {
		privateKeys = append(privateKeys, rm.rootPrivateKeys[:signersCount]...)
	}
	for _, key := range privateKeys {
		sig, err := crypto.Sign(set.hash.Bytes(), key)
		if err != nil {
			t.Error(errors.Wrap(err, "failed to sign hash"))
		}
		list.Signatures = append(list.Signatures, sig)
	}
	list.Hash = set.hash
	return list
}

func signedActiveExclusionList(t *testing.T, rm *TestRootManager, signersCount int, signByAlias bool) common.ValidatorExclusionList {
	return signedActiveExclusionListWithTimestamp(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), signersCount, signByAlias)
}
