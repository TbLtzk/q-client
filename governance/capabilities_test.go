package governance

import (
	"reflect"
	"testing"

	"gitlab.com/q-dev/q-client/common"
)

func TestL0GovernanceCapabilities(t *testing.T) {
	t.Run("enabled", func(t *testing.T) {
		rm := newTestRootManager(t, false, true)
		api := newGovernancePublicAPI(t, rm.RootManager)

		cap, err := api.L0GovernanceCapabilities()
		if err != nil {
			t.Fatalf("L0GovernanceCapabilities: %v", err)
		}
		if !cap.ExternalSubmissionEnabled {
			t.Fatal("expected external submission enabled")
		}
		if cap.QgovTypedRelayVersion == nil || *cap.QgovTypedRelayVersion != qgov6 {
			t.Fatalf("qgovTypedRelayVersion %+v want %d", cap.QgovTypedRelayVersion, qgov6)
		}
		assertProposalCaps(t, cap.RootList, true)
	})

	t.Run("disabled external submission", func(t *testing.T) {
		rm := newTestRootManager(t, false, true)
		rm.DisablePublicSubmission = true
		api := newGovernancePublicAPI(t, rm.RootManager)

		cap, err := api.L0GovernanceCapabilities()
		if err != nil {
			t.Fatalf("L0GovernanceCapabilities: %v", err)
		}
		if cap.ExternalSubmissionEnabled {
			t.Fatal("expected external submission disabled")
		}
		if cap.DisableReason != errPublicGovernanceSubmissionDisabled.Error() {
			t.Fatalf("disableReason %q", cap.DisableReason)
		}
		if cap.RootList.RawSubmit || cap.RootList.TypedSubmit {
			t.Fatal("submit flags must be false when external submission is disabled")
		}
	})
}

func assertProposalCaps(t *testing.T, got common.L0GovernanceProposalCapabilities, submitEnabled bool) {
	t.Helper()
	if got.RawSubmit != submitEnabled || got.TypedSubmit != submitEnabled {
		t.Fatalf("submit flags raw=%v typed=%v want %v", got.RawSubmit, got.TypedSubmit, submitEnabled)
	}
	wantSchemes := []string{
		common.GovernanceSigningSchemeRawListHash,
		common.GovernanceSigningSchemeEIP712V1,
	}
	if !reflect.DeepEqual(got.SigningSchemes, wantSchemes) {
		t.Fatalf("signingSchemes %#v", got.SigningSchemes)
	}
}
