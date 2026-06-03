package governance

import "gitlab.com/q-dev/q-client/crypto"

// normalizeECDSASignatureV returns a copy of sig with Ethereum Yellow Paper v (27/28)
// converted to the recovery id form (0/1) expected by crypto.SigToPub. Values already
// in 0/1 are unchanged. Caller-owned slices are not mutated.
func normalizeECDSASignatureV(sig []byte) []byte {
	if len(sig) != crypto.SignatureLength {
		return sig
	}
	out := make([]byte, len(sig))
	copy(out, sig)
	if out[crypto.RecoveryIDOffset] >= 27 {
		out[crypto.RecoveryIDOffset] -= 27
	}
	return out
}
