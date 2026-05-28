package governance

import (
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
)

// dualVerifyTypedFirst controls list-message signature try order (Phase 2 flips to true).
var dualVerifyTypedFirst = false

func verifyRootListSignature(set *rootSet, sig []byte, networkID uint64, typedFirst bool) (signer common.Address, ok bool) {
	if networkID == 0 {
		typedFirst = false
	}

	tryRaw := func() (common.Address, bool) {
		pubkey, err := crypto.SigToPub(set.hash.Bytes(), sig)
		if err != nil {
			return common.Address{}, false
		}
		return crypto.PubkeyToAddress(*pubkey), true
	}
	tryTyped := func() (common.Address, bool) {
		payloadHash := common.NewRootListSigningPayloadV1(networkID, common.RootList{
			Timestamp: set.timestamp,
			Nodes:     set.rootAddresses,
			Hash:      set.hash,
		}).Hash()
		pubkey, err := crypto.SigToPub(payloadHash.Bytes(), sig)
		if err != nil {
			return common.Address{}, false
		}
		return crypto.PubkeyToAddress(*pubkey), true
	}

	if typedFirst {
		if signer, ok := tryTyped(); ok {
			return signer, true
		}
		return tryRaw()
	}
	if signer, ok := tryRaw(); ok {
		return signer, true
	}
	return tryTyped()
}

func verifyExclusionListSignature(set *exclusionSet, sig []byte, networkID uint64, typedFirst bool) (signer common.Address, ok bool) {
	if networkID == 0 {
		typedFirst = false
	}

	tryRaw := func() (common.Address, bool) {
		pubkey, err := crypto.SigToPub(set.hash.Bytes(), sig)
		if err != nil {
			return common.Address{}, false
		}
		return crypto.PubkeyToAddress(*pubkey), true
	}
	tryTyped := func() (common.Address, bool) {
		unsigned := set.makeList()
		unsigned.Signatures = nil
		payloadHash := common.NewExclusionListSigningPayloadV1(networkID, unsigned).Hash()
		pubkey, err := crypto.SigToPub(payloadHash.Bytes(), sig)
		if err != nil {
			return common.Address{}, false
		}
		return crypto.PubkeyToAddress(*pubkey), true
	}

	if typedFirst {
		if signer, ok := tryTyped(); ok {
			return signer, true
		}
		return tryRaw()
	}
	if signer, ok := tryRaw(); ok {
		return signer, true
	}
	return tryTyped()
}

// hasRawRootListSignature reports whether signer has a signature on set that verifies over list.Hash.
func (s *rootSet) hasRawSignatureFrom(signer common.Address) bool {
	sig, ok := s.signers[signer]
	if !ok {
		return false
	}
	_, err := crypto.SigToPub(s.hash.Bytes(), sig)
	return err == nil
}
