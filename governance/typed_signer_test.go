package governance

import (
	"strings"
	"testing"

	"gitlab.com/q-dev/q-client/crypto"
)

func TestValidateTypedActiveSignerAliasRequired(t *testing.T) {
	rm := newTestRootManager(t, false, true)
	active := rm.active.copy()

	rootAddr := crypto.PubkeyToAddress(rm.rootPrivateKeys[0].PublicKey)
	aliasAddr := crypto.PubkeyToAddress(rm.aliasPrivateKeys[0].PublicKey)

	err := validateTypedActiveSigner(active, rootAddr, "root list")
	if err == nil {
		t.Fatal("expected error for root key when alias is required")
	}
	if !strings.Contains(err.Error(), errTypedMustUseAliasPrefix) {
		t.Fatalf("error %q missing %q", err, errTypedMustUseAliasPrefix)
	}
	if !strings.Contains(err.Error(), aliasAddr.Hex()) {
		t.Fatalf("error %q missing alias %s", err, aliasAddr.Hex())
	}
	if !strings.Contains(err.Error(), rootAddr.Hex()) {
		t.Fatalf("error %q missing root %s", err, rootAddr.Hex())
	}

	if err := validateTypedActiveSigner(active, aliasAddr, "exclusion list"); err != nil {
		t.Fatalf("alias signer should be accepted: %v", err)
	}
}

func TestRequiredTypedSigningAddress(t *testing.T) {
	rm := newTestRootManager(t, false, true)
	active := rm.active.copy()

	rootAddr := crypto.PubkeyToAddress(rm.rootPrivateKeys[0].PublicKey)
	aliasAddr := crypto.PubkeyToAddress(rm.aliasPrivateKeys[0].PublicKey)

	if got := active.requiredTypedSigningAddress(rootAddr); got != aliasAddr {
		t.Fatalf("root queried: got %s want alias %s", got.Hex(), aliasAddr.Hex())
	}
	if got := active.requiredTypedSigningAddress(aliasAddr); got != aliasAddr {
		t.Fatalf("alias queried: got %s want %s", got.Hex(), aliasAddr.Hex())
	}
}
