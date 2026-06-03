package governance

import (
	"fmt"

	"gitlab.com/q-dev/q-client/common"
)

// errTypedMustUseAliasPrefix is the stable substring returned when a typed signature
// was recovered from a root main key but the active set requires the alias address.
const errTypedMustUseAliasPrefix = "signer must use alias"

func (s *rootSet) requiredTypedSigningAddress(queried common.Address) common.Address {
	if s == nil || queried == (common.Address{}) {
		return common.Address{}
	}
	for root, alias := range s.aliases {
		if queried == root || queried == alias {
			if root != alias {
				return alias
			}
			return root
		}
	}
	return common.Address{}
}

func validateTypedActiveSigner(active *rootSet, signer common.Address, proposalLabel string) error {
	if active == nil {
		return fmt.Errorf("active root list is unavailable")
	}
	if len(active.knownSigners(map[common.Address][]byte{signer: nil})) > 0 {
		return nil
	}
	for root, alias := range active.aliases {
		if root != alias && signer == root {
			return fmt.Errorf("typed %s: %s %s, not root %s", proposalLabel, errTypedMustUseAliasPrefix, alias.Hex(), root.Hex())
		}
	}
	return fmt.Errorf("typed %s contains unknown signer", proposalLabel)
}
