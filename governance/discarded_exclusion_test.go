package governance

import (
	"testing"
	"time"
)

func TestDiscardedExclusionListImportNoOp(t *testing.T) {
	rm := newTestRootManager(t, false, true)
	bc := newTestChain(t, rm.RootManager)
	defer bc.Stop()
	rm.InitBlockChain(bc)

	list := signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)

	gov, err := New(rm.RootManager, tmpDirName(t))
	if err != nil {
		t.Fatalf("Failed to create Governance: %v", err)
	}
	startGovernance(t, gov)

	if err := gov.handler.importExclusionList(&list); err != nil {
		t.Fatalf("first import returned error: %v", err)
	}
	if rm.proposedExSet == nil {
		t.Fatalf("expected proposed exclusion list after first import")
	}
	proposedHash := rm.proposedExSet.hash

	if err := rm.discardExclusionList(&proposedHash); err != nil {
		t.Fatalf("discardExclusionList returned error: %v", err)
	}
	if rm.proposedExSet != nil {
		t.Fatalf("expected proposed exclusion list cleared after discard")
	}

	if err := gov.handler.importExclusionList(&list); err != nil {
		t.Fatalf("second import returned error: %v", err)
	}
	if rm.proposedExSet != nil {
		t.Fatalf("discarded exclusion list import changed proposed state: %v", rm.proposedExSet)
	}
}
