package governance

import (
	"math/rand"
	"testing"

	"gitlab.com/q-dev/q-client/p2p/enode"
)

func TestNewGovernanceLifecycle(t *testing.T) {
	var gov = newGovernance(t)

	defer func(gov *Governance) {
		err := gov.Stop()
		if err != nil {
			t.Fatalf("Failed to stop Governance: %v", err)
		}
	}(gov)
	if gov == nil {
		t.Fatalf("Failed to create Governance instance")
	}
}

func TestRunPeer(t *testing.T) {
	var gov = newGovernance(t)

	defer func(gov *Governance) {
		err := gov.Stop()
		if err != nil {
			t.Fatalf("Failed to stop Governance: %v", err)
		}
	}(gov)
	if gov == nil {
		t.Fatalf("Failed to create Governance instance")
	}
}

func newGovernance(t *testing.T) *Governance {
	rm := newTestRootManager(t, true)

	bc := newTestChain(t, rm)
	defer bc.Stop()
	rm.InitBlockChain(bc)

	gov, err := New(rm, tmpDirName(t))
	if err != nil {
		t.Fatalf("Failed to create Governance: %v", err)
	}

	err = gov.Start()
	if err != nil {
		t.Fatalf("Failed to start Governance: %v", err)
		return nil
	}

	return gov
}

func randomPeerID() (id enode.ID) {
	for i := range id {
		id[i] = byte(rand.Intn(255))
	}
	return id
}
