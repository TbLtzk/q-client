package governance

import (
	"testing"

	"gitlab.com/q-dev/q-client/accounts/keystore"
)

func catchUpTestHandler(t *testing.T, rm *TestRootManager) *handler {
	t.Helper()

	gov, err := New(rm.RootManager, tmpDirName(t))
	if err != nil {
		t.Fatalf("failed to create governance: %v", err)
	}
	// catchUpSignActiveLists propagates via rootEventCh/exEventCh; consumers must be running.
	gov.handler.run()
	t.Cleanup(func() { gov.handler.stop() })
	return gov.handler
}

func lockAllKeystoreAccounts(t *testing.T, rm *RootManager) *keystore.KeyStore {
	t.Helper()

	backends := rm.manager.Backends(keystore.KeyStoreType)
	if len(backends) == 0 {
		t.Fatal("keystore backend missing")
	}
	ks := backends[0].(*keystore.KeyStore)
	for _, acc := range ks.Accounts() {
		if err := ks.Lock(acc.Address); err != nil {
			t.Fatalf("failed to lock account %s: %v", acc.Address.Hex(), err)
		}
	}
	return ks
}

func TestCatchUpSignActiveLists(t *testing.T) {
	t.Run("unlocked root alias", func(t *testing.T) {
		rm := newTestRootManager(t, true, false)
		h := catchUpTestHandler(t, rm)
		ks := lockAllKeystoreAccounts(t, rm.RootManager)

		acc := ks.Accounts()[0]
		if _, ok := rm.active.signers[acc.Address]; ok {
			t.Fatal("precondition: active root list should not contain local signature")
		}
		if _, ok := rm.activeExSet.signers[acc.Address]; ok {
			t.Fatal("precondition: active exclusion list should not contain local signature")
		}

		if err := ks.Unlock(acc, ""); err != nil {
			t.Fatalf("failed to unlock account: %v", err)
		}
		h.catchUpSignActiveLists()

		rm.rootLock.Lock()
		if _, ok := rm.active.signers[acc.Address]; !ok {
			t.Fatal("expected catch-up signature on active root list")
		}
		rm.rootLock.Unlock()

		rm.exLock.Lock()
		if _, ok := rm.activeExSet.signers[acc.Address]; !ok {
			t.Fatal("expected catch-up signature on active exclusion list")
		}
		rm.exLock.Unlock()

		savedRoot, err := rm.db.getActiveRootSet()
		if err != nil {
			t.Fatalf("failed to load active root set from db: %v", err)
		}
		if _, ok := savedRoot.signers[acc.Address]; !ok {
			t.Fatal("expected catch-up root signature persisted to db")
		}

		savedEx := rm.db.getActiveExclusionSet()
		if savedEx == nil {
			t.Fatal("expected active exclusion set in db")
		}
		if _, ok := savedEx.signers[acc.Address]; !ok {
			t.Fatal("expected catch-up exclusion signature persisted to db")
		}

		signerCount := len(rm.active.signers)
		h.catchUpSignActiveLists()
		rm.rootLock.Lock()
		defer rm.rootLock.Unlock()
		if len(rm.active.signers) != signerCount {
			t.Fatalf("duplicate catch-up added signatures: got %d, want %d", len(rm.active.signers), signerCount)
		}
	})

	t.Run("non-root unchanged", func(t *testing.T) {
		rm := newTestRootManager(t, false, false)
		h := catchUpTestHandler(t, rm)

		h.catchUpSignActiveLists()

		rm.rootLock.Lock()
		defer rm.rootLock.Unlock()
		if len(rm.active.signers) != 0 {
			t.Fatalf("non-root node should not add catch-up signatures, got %d", len(rm.active.signers))
		}
	})

	t.Run("locked alias skipped", func(t *testing.T) {
		rm := newTestRootManager(t, true, false)
		h := catchUpTestHandler(t, rm)
		lockAllKeystoreAccounts(t, rm.RootManager)

		h.catchUpSignActiveLists()

		rm.rootLock.Lock()
		defer rm.rootLock.Unlock()
		if len(rm.active.signers) != 0 {
			t.Fatalf("locked alias should not be signed by catch-up, got %d signatures", len(rm.active.signers))
		}
	})
}
