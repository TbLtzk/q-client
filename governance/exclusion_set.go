package governance

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
)

type exclusionSet struct {
	timestamp uint64
	hash      common.Hash

	addresses   []common.Address
	addrToBlock map[common.Address]uint64

	signers map[common.Address][]byte
}

func newExclusionSet(list *common.ValidatorExclusionList) (*exclusionSet, error) {
	if list == nil {
		return nil, errInvalidExclusionList
	}

	var validators []common.Address
	validatorsSet := make(map[common.Address]uint64)
	for _, val := range list.Validators {
		// consider duplicates as error since
		// they can lead to ambiguity in block number
		if _, ok := validatorsSet[val.Address]; ok {
			return nil, errors.Errorf("duplicated address %s", val.Address.Hex())
		}

		validators = append(validators, val.Address)
		validatorsSet[val.Address] = val.Block
	}

	sort.SliceStable(validators, func(i, j int) bool {
		return bytes.Compare(validators[i].Bytes(), validators[j].Bytes()) > 0
	})

	set := &exclusionSet{
		timestamp:   list.Timestamp,
		addresses:   validators,
		addrToBlock: validatorsSet,
	}
	set.hash = set.calcHash()

	if set.hash != list.Hash && (list.Hash != common.Hash{}) {
		return nil, errHashMismatch
	}

	signers := make(map[common.Address][]byte)
	for _, sig := range list.Signatures {
		signer, err := crypto.SigToPub(set.hash.Bytes(), sig)
		if err != nil {
			return nil, errInvalidSignature
		}

		signers[crypto.PubkeyToAddress(*signer)] = sig
	}

	set.signers = signers
	return set, nil
}

func (s *exclusionSet) calcHash() common.Hash {
	msg := append([]byte{}, fmt.Sprint(s.timestamp)...)
	for _, addr := range s.addresses {
		msg = append(msg, addr.Bytes()...)
	}

	return crypto.Keccak256Hash(msg)
}

func (s *exclusionSet) mergeSignatures(hash common.Hash, signers map[common.Address][]byte) map[common.Address][]byte {
	if s.hash != hash {
		return nil
	}

	newSignatures := make(map[common.Address][]byte)
	for signer, signature := range signers {
		if _, ok := s.signers[signer]; ok {
			continue
		}

		s.signers[signer] = signature
		newSignatures[signer] = signature
	}

	return newSignatures
}

func (s *exclusionSet) addSignature(signer common.Address, signature []byte) bool {
	if _, ok := s.signers[signer]; ok {
		return false
	}

	s.signers[signer] = signature
	return true
}

func (s *exclusionSet) makeList() common.ValidatorExclusionList {
	if s == nil {
		return common.ValidatorExclusionList{}
	}

	var signatures [][]byte
	for _, sig := range s.signers {
		signatures = append(signatures, sig)
	}

	var validators []common.ExcludedValidator
	for addr, block := range s.addrToBlock {
		validators = append(validators, common.ExcludedValidator{
			Address: addr,
			Block:   block,
		})
	}

	return common.ValidatorExclusionList{
		Timestamp:  s.timestamp,
		Hash:       s.hash,
		Signatures: signatures,
		Validators: validators,
	}
}

func (s *exclusionSet) copy() *exclusionSet {
	if s == nil {
		return nil
	}

	// copy concurrently mutable signers map
	signers := make(map[common.Address][]byte)
	for signer, sig := range s.signers {
		signers[signer] = sig
	}

	return &exclusionSet{
		timestamp:   s.timestamp,
		hash:        s.hash,
		addresses:   s.addresses,
		addrToBlock: s.addrToBlock,
		signers:     signers,
	}
}

// addrToBlockExclusiveDiff returns map with exclusive set of address-to-block pairs from two exclusion sets
func (s1 *exclusionSet) addrToBlockExclusiveDiff(s2 *exclusionSet) map[common.Address]uint64 {
	if s2 == nil {
		return s1.addrToBlock
	}

	res := make(map[common.Address]uint64)

	// add addess-to-block that is only in s1, but not in s2
	for _, addr := range s1.addresses {
		if block, ok := s2.addrToBlock[addr]; ok && block == s1.addrToBlock[addr] {
			continue
		}

		res[addr] = s1.addrToBlock[addr]
	}

	// add addess-to-block that is only in s2, but not in s1
	for _, addr := range s2.addresses {
		if block, ok := s1.addrToBlock[addr]; ok && block == s2.addrToBlock[addr] {
			continue
		}

		res[addr] = s2.addrToBlock[addr]
	}

	return res
}

func (s *exclusionSet) getAddresses() []common.Address {
	if s == nil {
		return nil
	}

	return s.addresses
}
