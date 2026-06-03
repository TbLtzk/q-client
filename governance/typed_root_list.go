package governance

import (
	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
)

func (s *RootManager) newTypedSignedRootSet(list common.RootList) (*rootSet, error) {
	if len(list.Signatures) == 0 {
		return nil, errors.New("typed root list has no signatures")
	}

	unsigned := list
	unsigned.Signatures = nil

	set, err := newRootSet(&unsigned)
	if err != nil {
		return nil, errors.Wrap(err, "invalid typed root list")
	}
	if set == nil {
		return nil, errors.New("invalid typed root list")
	}

	active := s.getActiveRootSet(true)
	if active == nil {
		return nil, errors.New("active root list is unavailable")
	}

	signers := make(map[common.Address][]byte)
	for _, signature := range list.Signatures {
		signer, ok := verifyRootListEIP712Signature(s.networkId, set, signature)
		if !ok {
			return nil, errInvalidSignature
		}

		if len(active.knownSigners(map[common.Address][]byte{signer: signature})) == 0 {
			return nil, errors.New("typed root list contains unknown signer")
		}
		signers[signer] = signature
	}

	set.signers = signers
	set.updateAliases(s.getAliasesOfRoots(set.rootAddresses))

	return set, nil
}
