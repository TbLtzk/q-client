package governance

import (
	"testing"
	"time"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
)

func TestGetGovernanceProposalStatus(t *testing.T) {
	t.Run("active root list", func(t *testing.T) {
		rm := newTestRootManager(t, false, true)
		api := newGovernancePublicAPI(t, rm.RootManager)

		hash := rm.active.hash
		status, err := api.GetGovernanceProposalStatus(common.GovernanceProposalTypeRootList, hash, common.Address{})
		if err != nil {
			t.Fatalf("GetGovernanceProposalStatus: %v", err)
		}
		if status.Phase != governanceProposalPhaseActive {
			t.Fatalf("phase %q want active", status.Phase)
		}
		if status.Hash != hash {
			t.Fatalf("hash mismatch")
		}
		if status.Threshold.ThresholdPercent != rootListThresholdPercentage {
			t.Fatalf("threshold percent %d", status.Threshold.ThresholdPercent)
		}
	})

	t.Run("proposed root list", func(t *testing.T) {
		rm := newTestRootManager(t, false, true)
		api := newGovernancePublicAPI(t, rm.RootManager)

		rootList := randomRootList(t, rm, time.Now().Add(10*time.Minute).Unix(), 5, 1, true)
		set, err := newRootSet(&rootList)
		if err != nil {
			t.Fatal(err)
		}
		rm.proposed = set

		status, err := api.GetGovernanceProposalStatus(common.GovernanceProposalTypeRootList, set.hash, common.Address{})
		if err != nil {
			t.Fatalf("GetGovernanceProposalStatus: %v", err)
		}
		if status.Phase != governanceProposalPhaseProposed {
			t.Fatalf("phase %q want proposed", status.Phase)
		}
	})

	t.Run("unknown hash", func(t *testing.T) {
		rm := newTestRootManager(t, false, true)
		api := newGovernancePublicAPI(t, rm.RootManager)

		status, err := api.GetGovernanceProposalStatus(
			common.GovernanceProposalTypeRootList,
			common.BytesToHash([]byte{9, 9, 9}),
			common.Address{},
		)
		if err != nil {
			t.Fatalf("GetGovernanceProposalStatus: %v", err)
		}
		if status.Phase != governanceProposalPhaseUnknown {
			t.Fatalf("phase %q want unknown", status.Phase)
		}
	})

	t.Run("threshold matches isAcceptable at 75 percent", func(t *testing.T) {
		rm := newTestRootManager(t, false, true)
		active := rm.active.copy()
		active.aliases = rm.getAliasesOfRoots(active.rootAddresses)

		proposed := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 5, defNumAccounts-1, true)
		set, err := newRootSet(&proposed)
		if err != nil {
			t.Fatal(err)
		}

		status := buildRootListProposalStatus(active, set, governanceProposalPhaseProposed, common.Address{})
		if status.Threshold.MeetsThreshold != active.isAcceptable(set) {
			t.Fatalf("MeetsThreshold=%v isAcceptable=%v", status.Threshold.MeetsThreshold, active.isAcceptable(set))
		}
		if status.Threshold.ThresholdPercent != rootListThresholdPercentage {
			t.Fatalf("threshold percent %d", status.Threshold.ThresholdPercent)
		}
		if status.Threshold.SignedPercent < rootListThresholdPercentage {
			t.Fatalf("expected at least %d%% signed, got %d%%", rootListThresholdPercentage, status.Threshold.SignedPercent)
		}
	})

	t.Run("alias still needs signature", func(t *testing.T) {
		rm := newTestRootManager(t, false, true)
		active := rm.active.copy()
		active.aliases = rm.getAliasesOfRoots(active.rootAddresses)

		proposed := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 5, 0, true)
		set, err := newRootSet(&proposed)
		if err != nil {
			t.Fatal(err)
		}

		alias := rm.aliasPrivateKeys[0]
		queried := crypto.PubkeyToAddress(alias.PublicKey)
		status := buildRootListProposalStatus(active, set, governanceProposalPhaseProposed, queried)
		if !status.NeedsSignature {
			t.Fatal("expected alias to still need signature")
		}
	})

	t.Run("root does not need signature when alias already signed", func(t *testing.T) {
		rm := newTestRootManager(t, false, true)
		active := rm.active.copy()
		active.aliases = rm.getAliasesOfRoots(active.rootAddresses)

		proposed := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 5, 1, true)
		set, err := newRootSet(&proposed)
		if err != nil {
			t.Fatal(err)
		}

		rootAddr := rm.rootPrivateKeys[0]
		queried := crypto.PubkeyToAddress(rootAddr.PublicKey)
		status := buildRootListProposalStatus(active, set, governanceProposalPhaseProposed, queried)
		if status.NeedsSignature {
			t.Fatal("alias signature should satisfy the root signer slot")
		}
	})

	t.Run("alias still needs signature when only root key signed", func(t *testing.T) {
		rm := newTestRootManager(t, false, true)
		active := rm.active.copy()
		active.aliases = rm.getAliasesOfRoots(active.rootAddresses)

		proposed := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 5, 1, false)
		set, err := newRootSet(&proposed)
		if err != nil {
			t.Fatal(err)
		}

		alias := rm.aliasPrivateKeys[0]
		queried := crypto.PubkeyToAddress(alias.PublicKey)
		status := buildRootListProposalStatus(active, set, governanceProposalPhaseProposed, queried)
		if !status.NeedsSignature {
			t.Fatal("expected alias to still need signature when proposal only has root-key signature")
		}
	})

	t.Run("non-root address does not need signature", func(t *testing.T) {
		rm := newTestRootManager(t, false, true)
		active := rm.active.copy()
		active.aliases = rm.getAliasesOfRoots(active.rootAddresses)

		proposed := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 5, 0, true)
		set, err := newRootSet(&proposed)
		if err != nil {
			t.Fatal(err)
		}

		unknown, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		queried := crypto.PubkeyToAddress(unknown.PublicKey)
		status := buildRootListProposalStatus(active, set, governanceProposalPhaseProposed, queried)
		if status.NeedsSignature {
			t.Fatal("non-root address should not need signature")
		}
	})

	t.Run("proposed exclusion list", func(t *testing.T) {
		rm := newTestRootManager(t, false, true)
		api := newGovernancePublicAPI(t, rm.RootManager)

		list := signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, 1, true, 6000)
		set, err := newExclusionSet(&list)
		if err != nil {
			t.Fatal(err)
		}
		rm.proposedExSet = set

		status, err := api.GetGovernanceProposalStatus(common.GovernanceProposalTypeExclusionList, set.hash, common.Address{})
		if err != nil {
			t.Fatalf("GetGovernanceProposalStatus: %v", err)
		}
		if status.Phase != governanceProposalPhaseProposed {
			t.Fatalf("phase %q want proposed", status.Phase)
		}
		if status.ProposalType != common.GovernanceProposalTypeExclusionList {
			t.Fatalf("proposal type %q", status.ProposalType)
		}
	})

	t.Run("exclusion threshold matches isEnoughExSetSignatures", func(t *testing.T) {
		rm := newTestRootManager(t, false, true)
		active := rm.active.copy()
		active.aliases = rm.getAliasesOfRoots(active.rootAddresses)

		list := signedExclusionList(t, rm, uint64(time.Now().Add(5*time.Minute).Unix()), 2, defNumAccounts-1, true, 7000)
		set, err := newExclusionSet(&list)
		if err != nil {
			t.Fatal(err)
		}

		status := buildExclusionListProposalStatus(active, set, governanceProposalPhaseProposed, common.Address{})
		if status.Threshold.MeetsThreshold != active.isEnoughExSetSignatures(set) {
			t.Fatalf("MeetsThreshold=%v isEnoughExSetSignatures=%v", status.Threshold.MeetsThreshold, active.isEnoughExSetSignatures(set))
		}
	})

	t.Run("desired root list", func(t *testing.T) {
		rm := newTestRootManager(t, false, true)
		api := newGovernancePublicAPI(t, rm.RootManager)

		rootList := randomRootList(t, rm, time.Now().Add(10*time.Minute).Unix(), 5, defNumAccounts-1, true)
		set, err := newRootSet(&rootList)
		if err != nil {
			t.Fatal(err)
		}
		rm.desired = set

		status, err := api.GetGovernanceProposalStatus(common.GovernanceProposalTypeRootList, set.hash, common.Address{})
		if err != nil {
			t.Fatalf("GetGovernanceProposalStatus: %v", err)
		}
		if status.Phase != governanceProposalPhaseDesired {
			t.Fatalf("phase %q want desired", status.Phase)
		}
	})
}

func newGovernancePublicAPI(t *testing.T, rm *RootManager) *GovernancePublicAPI {
	t.Helper()
	gov, err := New(rm, tmpDirName(t))
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return NewGovernancePublicAPI(gov)
}
