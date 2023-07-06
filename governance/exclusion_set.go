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
	blockRanges map[common.Address][]common.BlockRange //Several address ranges can be assigned to one validator

	signers map[common.Address][]byte
}

func newExclusionSet(list *common.ValidatorExclusionList) (*exclusionSet, error) {
	if list == nil {
		return nil, errInvalidExclusionList
	}

	var validators []common.Address
	validatorsSet := make(map[common.Address][]common.BlockRange)

	for _, val := range list.Validators {
		// consider duplicates as error since
		// they can lead to ambiguity in block number
		if banBlocks, ok := validatorsSet[val.Address]; ok {
			contains := false
			for _, block := range banBlocks {
				if block.StartAddress == val.Block && block.EndAddress == val.EndBlock {
					contains = true
					break
				}
			}
			if contains {
				return nil, errors.Errorf("duplicated address range %s", val.Address.Hex())
			}
		}

		validators = append(validators, val.Address)
		validatorsSet[val.Address] = append(validatorsSet[val.Address], common.BlockRange{StartAddress: val.Block,
			EndAddress: val.EndBlock})
	}

	sort.SliceStable(validators, func(i, j int) bool {
		return bytes.Compare(validators[i].Bytes(), validators[j].Bytes()) > 0
	})

	addrs := make(map[common.Address]uint64)
	for address, br := range validatorsSet {
		var ba uint64
		if len(br) > 0 {
			ba = br[0].StartAddress
		}
		addrs[address] = ba
	}

	set := &exclusionSet{
		timestamp:   list.Timestamp,
		addresses:   validators,
		addrToBlock: addrs,
		blockRanges: validatorsSet,
	}

	if (list.Hash != common.Hash{}) {
		set.hash = set.calcFullHash() //new hash rule only for validation (for now). Try it first
	} else {
		set.hash = set.calcHash() //old hash rule is applied if list is new and hash is empty
	}

	set.fixTimestamp()

	if set.hash != list.Hash && (list.Hash != common.Hash{}) {
		set.hash = set.calcHash()
		set.fixTimestamp()
		if set.hash != list.Hash && (list.Hash != common.Hash{}) {
			return nil, errHashMismatch
		}
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

func (s *exclusionSet) calcFullHash() common.Hash {
	msg := append([]byte{}, fmt.Sprint(s.timestamp)...)

	keys := make([]common.Address, 0, len(s.blockRanges))

	for k := range s.blockRanges {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].String() < keys[j].String()
	})

	for _, address := range keys {
		msg = append(msg, address.Bytes()...)
		ranges := s.blockRanges[address]
		for _, r := range ranges {
			msg = append(msg, fmt.Sprint(r.StartAddress)...)
			msg = append(msg, fmt.Sprint(r.EndAddress)...)
		}
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
	for address, blocks := range s.blockRanges {
		for _, blockAddr := range blocks {
			validators = append(validators, common.ExcludedValidator{
				Address:  address,
				Block:    blockAddr.StartAddress,
				EndBlock: blockAddr.EndAddress,
			})
		}
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
	addrs := make(map[common.Address]uint64)
	for address, br := range s.blockRanges {
		var ba uint64
		if len(br) > 0 {
			ba = br[0].StartAddress
		}
		addrs[address] = ba
	}

	return &exclusionSet{
		timestamp:   s.timestamp,
		hash:        s.hash,
		addresses:   s.addresses,
		addrToBlock: addrs,
		blockRanges: s.blockRanges,
		signers:     signers,
	}
}

// addrToBlockRangeExclusiveDiff returns map with exclusive set of address-to-block pairs from two exclusion sets
func (s1 *exclusionSet) addrToBlockRangeExclusiveDiff(s2 *exclusionSet) map[common.Address][]common.BlockRange {
	if s2 == nil {
		return s1.blockRanges
	}
	if s1 == nil {
		return nil
	}

	res := make(map[common.Address][]common.BlockRange)

	// add addess-to-block that is only in s1, but not in s2
	for addr, blockRangesFromS1 := range s1.blockRanges {
		for _, blockRangeFromS1 := range blockRangesFromS1 {
			exist := false
			for _, blockRangeFromS2 := range s2.blockRanges[addr] {
				if blockRangeFromS1.IsEqualTo(blockRangeFromS2) {
					{
						exist = true
						break
					}
				}
			}

			if !exist {
				res[addr] = append(res[addr], blockRangeFromS1)
			}
		}
	}

	// add addess-to-block that is only in s2, but not in s1
	for addr, blockRangesFromS2 := range s2.blockRanges {
		for _, blockRangeFromS2 := range blockRangesFromS2 {
			exist := false
			for _, blockRangeFromS1 := range s1.blockRanges[addr] {
				if blockRangeFromS2.IsEqualTo(blockRangeFromS1) {
					{
						exist = true
						break
					}
				}
			}

			if !exist {
				res[addr] = append(res[addr], blockRangeFromS2)
			}
		}
	}
	return res
}

func (s *exclusionSet) getAddresses() []common.Address {
	if s == nil {
		return nil
	}

	return s.addresses
}

// Fix after unhappy testing
func (s *exclusionSet) fixTimestamp() {
	if s.hash.String() == "0x0e93f67e14240ec1e3aa6f0e316fc331cd3197816e88a7dc21b8749d092e1762" {
		s.timestamp = 1669136688
	}
	if s.hash.String() == "0x1241ef2c59655851eab2d85deedf22812a759aefdf63f22658ac878c5db6819c" {
		s.hash = common.HexToHash("0x0e93f67e14240ec1e3aa6f0e316fc331cd3197816e88a7dc21b8749d092e1762")
	}
}
