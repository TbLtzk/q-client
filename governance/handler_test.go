package governance

import (
	"crypto/rand"
	"errors"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/p2p"
	"os"
	"testing"
	"time"
)

func TestHandshake(t *testing.T) {
	rm := newTestRootManager(t, false)
	cm, errCm := NewConstitutionManager(os.TempDir(), rm.db, rm)
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
			err:  errors.New("invalid network id"),
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
			err := handshake(rm, cm, test.code, test.msg)
			if (test.err != nil && err == nil) || (test.err == nil && err != nil) || (test.err != nil && err != nil && err.Error() != test.err.Error()) {
				t.Errorf("Expected error '%v', got '%v'", test.err, err)
			}

		})
	}

}

func TestHandleMessage(t *testing.T) {
	rm := newTestRootManager(t, false)
	cm, errCm := NewConstitutionManager(os.TempDir(), rm.db, rm)
	if errCm != nil {
		log.Error("Can't create ConstitutionManager: %v", errCm)
	}
	h := newHandler(rm, cm)
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
			err:  errors.New("unknown msg code"),
		}, {
			name: "Too long message",
			code: StatusMsg,
			msg:  randString(10000000),
			err:  errors.New("message too large"),
		}, {
			name: "Large root list",
			code: RootListMsg,
			msg:  randomRootList(2*maxNRootNodes + 1),
			err:  errors.New("message too large"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := handleMsg(h, test.code, test.msg, p1, rw2)
			if (test.err != nil && err == nil) || (test.err == nil && err != nil) || (test.err != nil && err != nil && err.Error() != test.err.Error()) {
				t.Errorf("Expected error '%v', got '%v'", test.err, err)
			}
		})
	}

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

func randomRootList(size int) common.RootList {
	rootList := common.RootList{
		Timestamp:  1680255000,
		Nodes:      make([]common.Address, size),
		Hash:       common.BytesToHash(randBytes(common.HashLength)),
		Signatures: make([][]byte, size),
	}
	for i := 0; i < size; i++ {
		rootList.Nodes[i] = common.BytesToAddress(randBytes(common.AddressLength))
		rootList.Signatures = append(rootList.Signatures, randBytes(crypto.SignatureLength))
	}

	return rootList
}

//TODO tests for handleRootSet,handleExclusionSet, constitutionManager and approvalMessages