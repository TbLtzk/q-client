package governance

import (
	"testing"
	"time"

	"gitlab.com/q-dev/q-client/accounts"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/p2p"
)

func TestProtocolVersionsIncludeQgov6(t *testing.T) {
	if len(ProtocolVersions) < 2 || ProtocolVersions[0] != qgov6 {
		t.Fatalf("ProtocolVersions = %v, want primary %d", ProtocolVersions, qgov6)
	}
	if protocolLengths[qgov6] != 9 {
		t.Fatalf("qgov6 length = %d, want 9", protocolLengths[qgov6])
	}
}

func TestTypedRelayDedupRemember(t *testing.T) {
	d := newTypedRelayDedup()
	key := typedRelayKey{proposalHash: common.HexToHash("0x01"), signer: common.HexToAddress("0x02"), kind: 0}
	if !d.remember(key) {
		t.Fatal("first remember failed")
	}
	if d.remember(key) {
		t.Fatal("duplicate remember should return false")
	}
}

func TestDualVerifySkipsInvalidSignature(t *testing.T) {
	rm := newTestRootManager(t, false, true)
	list := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 1, true)
	list.Signatures[0] = randBytes(crypto.SignatureLength)

	set, err := newRootSetForNetwork(&list, rm.networkId)
	if err != nil {
		t.Fatalf("newRootSetForNetwork: %v", err)
	}
	if len(set.signers) != 0 {
		t.Fatalf("invalid signature should be skipped, got signers %v", set.signers)
	}
}

func TestTypedRootRelayIgnoredOnQgov5Peer(t *testing.T) {
	rm := newTestRootManager(t, false, true)
	cm, err := NewConstitutionManager(tmpDirName(t), rm.db, rm.RootManager)
	if err != nil {
		t.Fatalf("NewConstitutionManager: %v", err)
	}
	h := newHandler(rm.RootManager, cm)

	rw1, rw2 := p2p.MsgPipe()
	p1 := newPeer(qgov5, p2p.NewPeer(randomPeerID(), "peer", nil), rw1)
	defer p1.Disconnect(p2p.DiscRequested)

	list := testTypedRootList(t, rm)
	if err := handleMsg(h, TypedRootListRelayMsg, list, p1, rw2); err != nil {
		t.Fatalf("handleMsg: %v", err)
	}
}

func TestSelfBridgeTypedRootListMintsRaw(t *testing.T) {
	rm := newTestRootManager(t, true, false)
	cm, err := NewConstitutionManager(tmpDirName(t), rm.db, rm.RootManager)
	if err != nil {
		t.Fatalf("NewConstitutionManager: %v", err)
	}
	h := newHandler(rm.RootManager, cm)
	h.run()
	defer h.stop()

	list := testTypedRootList(t, rm)
	hash, err := h.ProcessTypedSignedRootList(list)
	if err != nil {
		t.Fatalf("ProcessTypedSignedRootList: %v", err)
	}

	root := rm.active.rootAddresses[0]
	signer := rm.active.aliases[root]
	if !h.rootSetHasRawSignature(hash, signer) {
		t.Fatalf("expected raw signature on proposal %s from signer %s after self-bridge", hash.Hex(), signer.Hex())
	}
}

func testTypedRootList(t *testing.T, rm *TestRootManager) common.RootList {
	t.Helper()

	list := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 0, true)
	unsigned := list
	unsigned.Signatures = nil

	root := rm.active.rootAddresses[0]
	signer := rm.active.aliases[root]

	var signature []byte
	if len(rm.aliasPrivateKeys) > 0 {
		signature = signRootListEIP712(t, rm.networkId, list, rm.aliasPrivateKeys[0])
	} else {
		unsigned := list
		unsigned.Signatures = nil
		set, err := newRootSet(&unsigned)
		if err != nil || set == nil {
			t.Fatalf("newRootSet: %v", err)
		}
		digest, err := rootListEIP712Digest(rm.networkId, set.timestamp, set.hash, set.rootAddresses)
		if err != nil {
			t.Fatalf("rootListEIP712Digest: %v", err)
		}
		signature, err = rm.SignHash(accounts.Account{Address: signer}, digest.Bytes())
		if err != nil {
			t.Fatalf("sign typed payload: %v", err)
		}
	}
	list.Signatures = common.HexSignaturesFromBytes([][]byte{signature})
	return list
}
