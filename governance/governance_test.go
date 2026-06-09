package governance

import (
	"math/rand"
	"testing"

	"gitlab.com/q-dev/q-client/p2p/enode"
)

func TestNewGovernanceLifecycle(t *testing.T) {
	if newGovernance(t) == nil {
		t.Fatalf("Failed to create Governance instance")
	}
}

func TestRunPeer(t *testing.T) {
	if newGovernance(t) == nil {
		t.Fatalf("Failed to create Governance instance")
	}
}

// startGovernance starts the governance service and registers t.Cleanup to stop it.
func startGovernance(t *testing.T, gov *Governance) {
	t.Helper()
	if err := gov.Start(); err != nil {
		t.Fatalf("Failed to start Governance: %v", err)
	}
	t.Cleanup(func() {
		if err := gov.Stop(); err != nil {
			t.Fatalf("Failed to stop Governance: %v", err)
		}
	})
}

func newGovernance(t *testing.T) *Governance {
	rm := newTestRootManager(t, true, false)

	bc := newTestChain(t, rm.RootManager)
	t.Cleanup(func() { bc.Stop() })
	rm.InitBlockChain(bc)

	gov, err := New(rm.RootManager, tmpDirName(t))
	if err != nil {
		t.Fatalf("Failed to create Governance: %v", err)
	}

	startGovernance(t, gov)
	return gov
}

func randomPeerID() (id enode.ID) {
	for i := range id {
		id[i] = byte(rand.Intn(255))
	}
	return id
}
