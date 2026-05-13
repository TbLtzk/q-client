package governance

import (
	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
)

func (s *RootManager) newTypedSignedExclusionSet(list common.ValidatorExclusionList) (*exclusionSet, error) {
	if len(list.Signatures) == 0 {
		return nil, errors.New("typed exclusion list has no signatures")
	}

	unsigned := list
	unsigned.Signatures = nil

	set, err := newExclusionSet(&unsigned)
	if err != nil {
		return nil, errors.Wrap(err, "invalid typed exclusion list")
	}
	if set == nil {
		return nil, errors.New("invalid typed exclusion list")
	}

	unsignedList := set.makeList()
	unsignedList.Signatures = nil

	payloadHash := common.NewExclusionListSigningPayloadV1(s.networkId, unsignedList).Hash()

	active := s.getActiveRootSet(true)
	if active == nil {
		return nil, errors.New("active root list is unavailable")
	}

	signers := make(map[common.Address][]byte)
	for _, signature := range list.Signatures {
		pubkey, err := crypto.SigToPub(payloadHash.Bytes(), signature)
		if err != nil {
			return nil, errInvalidSignature
		}
		signer := crypto.PubkeyToAddress(*pubkey)

		if len(active.knownSigners(map[common.Address][]byte{signer: signature})) == 0 {
			return nil, errors.New("typed exclusion list contains unknown signer")
		}
		signers[signer] = signature
	}

	set.signers = signers
	return set, nil
}
