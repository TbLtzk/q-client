package governance

import (
	"math/rand"
	"reflect"
	"testing"

	"gitlab.com/q-dev/q-client/p2p"
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

func TestHandeMessage(t *testing.T) {
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

	//TODO: implement
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

//nolint:unused,deadcode
func runPeer(t *testing.T, protocol p2p.Protocol) *p2p.Peer {
	name := "nodename"
	caps := []p2p.Cap{}
	caps = append(caps, p2p.Cap{Name: protocol.Name, Version: protocol.Version})

	id := randomPeerID()
	p := p2p.NewPeer(id, name, caps)
	if p.ID() != id {
		t.Errorf("ID mismatch: got %v, expected %v", p.ID(), id)
		return nil
	}
	if p.Name() != name {
		t.Errorf("Name mismatch: got %v, expected %v", p.Name(), name)
		return nil
	}
	if !reflect.DeepEqual(p.Caps(), caps) {
		t.Errorf("Caps mismatch: got %v, expected %v", p.Caps(), caps)
		return nil
	}

	return p
}

//nolint:unused
func randomPeerID() (id enode.ID) {
	for i := range id {
		id[i] = byte(rand.Intn(255))
	}
	return id
}
