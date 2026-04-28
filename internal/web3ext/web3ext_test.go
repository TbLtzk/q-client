package web3ext

import (
	"bytes"
	"testing"
)

func TestGovernanceSubmissionWeb3Extensions(t *testing.T) {
	for _, expected := range []string{
		"name: 'submitSignedRootList'",
		"call: 'govPub_submitSignedRootList'",
		"name: 'submitSignedExclusionList'",
		"call: 'govPub_submitSignedExclusionList'",
	} {
		if !bytes.Contains([]byte(GovPublicJs), []byte(expected)) {
			t.Fatalf("GovPublicJs missing %q", expected)
		}
	}
}
