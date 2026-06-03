package governance

import (
	"testing"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
)

func TestNormalizeECDSASignatureV_walletV(t *testing.T) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}

	list := common.RootList{
		Timestamp: 1715072400,
		Nodes:     []common.Address{common.HexToAddress("0xA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B")},
	}
	unsigned := list
	set, err := newRootSet(&unsigned)
	if err != nil || set == nil {
		t.Fatalf("newRootSet: %v", err)
	}
	list.Hash = set.hash

	digest, err := rootListEIP712Digest(1337, set.timestamp, set.hash, set.rootAddresses)
	if err != nil {
		t.Fatalf("digest: %v", err)
	}
	sig, err := crypto.Sign(digest.Bytes(), key)
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}

	addr0, ok0 := verifyRootListEIP712Signature(1337, set, sig)
	if !ok0 {
		t.Fatal("expected v=0/1 signature to verify")
	}

	walletSig := append([]byte(nil), sig...)
	walletSig[crypto.RecoveryIDOffset] += 27

	addr27, ok27 := verifyRootListEIP712Signature(1337, set, walletSig)
	if !ok27 {
		t.Fatal("expected v=27/28 wallet signature to verify after normalization")
	}
	if addr27 != addr0 {
		t.Fatalf("signer mismatch: %s vs %s", addr27.Hex(), addr0.Hex())
	}

	norm := normalizeECDSASignatureV(walletSig)
	if norm[crypto.RecoveryIDOffset] != sig[crypto.RecoveryIDOffset] {
		t.Fatalf("normalized v=%d want %d", norm[crypto.RecoveryIDOffset], sig[crypto.RecoveryIDOffset])
	}
}

func TestVerifyRootListSignature_walletV_rawPath(t *testing.T) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}

	list := common.RootList{Timestamp: 1715072401, Nodes: []common.Address{common.HexToAddress("0x01")}}
	set, err := newRootSet(&list)
	if err != nil || set == nil {
		t.Fatalf("newRootSet: %v", err)
	}

	sig, err := crypto.Sign(set.hash.Bytes(), key)
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}
	walletSig := append([]byte(nil), sig...)
	walletSig[crypto.RecoveryIDOffset] += 27

	_, ok := verifyRootListSignature(set, walletSig, 0, false)
	if !ok {
		t.Fatal("raw dual-verify path should accept wallet v on list hash")
	}
}
