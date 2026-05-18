package web3ext

import (
	"bytes"
	"testing"
)

func TestGovernanceSubmissionWeb3Extensions(t *testing.T) {
	for _, expected := range []string{
		"name: 'submitSignedRootList'",
		"call: 'govPub_submitSignedRootList'",
		"name: 'signingPayloadRootListV1'",
		"call: 'govPub_signingPayloadRootListV1'",
		"name: 'signingPayloadRootListV1WithDigest'",
		"call: 'govPub_signingPayloadRootListV1WithDigest'",
		"name: 'submitSignedExclusionList'",
		"call: 'govPub_submitSignedExclusionList'",
		"name: 'signingPayloadExclusionListV1'",
		"call: 'govPub_signingPayloadExclusionListV1'",
		"name: 'signingPayloadExclusionListV1WithDigest'",
		"call: 'govPub_signingPayloadExclusionListV1WithDigest'",
	} {
		if !bytes.Contains([]byte(GovPublicJs), []byte(expected)) {
			t.Fatalf("GovPublicJs missing %q", expected)
		}
	}
}
