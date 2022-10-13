package governance

import (
	"bytes"
	"fmt"
	"sort"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/log"
)

const (
	rootListThresholdPercentage      = 75
	exclusionListThresholdPercentage = 75
	approvalsThresholdPercentage     = 75
)

type rootSet struct {
	timestamp uint64
	hash      common.Hash

	rootAddresses []common.Address
	aliases       map[common.Address]common.Address
	roots         map[common.Address]struct{}

	signers map[common.Address][]byte //who actually signed list
}

func (s *rootSet) addSignature(signer common.Address, signature []byte) bool {
	if _, ok := s.signers[signer]; ok {
		return false
	}

	s.signers[signer] = signature
	return true
}

func (s *rootSet) calcHash() common.Hash {
	msg := append([]byte{}, fmt.Sprint(s.timestamp)...)
	for _, addr := range s.rootAddresses {
		msg = append(msg, addr.Bytes()...)
	}

	return crypto.Keccak256Hash(msg)
}

func newRootSet(list *common.RootList) (*rootSet, error) {
	if list == nil || len(list.Nodes) == 0 {
		return nil, nil
	}

	roots := make(map[common.Address]struct{})
	var rootAddrs []common.Address
	for _, addr := range list.Nodes {
		if _, ok := roots[addr]; ok {
			continue
		}

		roots[addr] = struct{}{}
		rootAddrs = append(rootAddrs, addr)
	}

	// normalize root list
	sort.SliceStable(rootAddrs, func(i, j int) bool {
		return bytes.Compare(rootAddrs[i].Bytes(), rootAddrs[j].Bytes()) > 0
	})

	set := &rootSet{
		timestamp:     list.Timestamp,
		roots:         roots,
		rootAddresses: rootAddrs,
	}
	set.hash = set.calcHash()

	if set.hash != list.Hash && (list.Hash != common.Hash{}) {
		return nil, errHashMismatch
	}

	signers := make(map[common.Address][]byte)
	hash := set.hash.Bytes()
	for _, sig := range list.Signatures {
		pubkey, err := crypto.SigToPub(hash, sig)
		if err != nil {
			return nil, errInvalidSignature
		}

		addr := crypto.PubkeyToAddress(*pubkey)
		//if _, ok := roots[addr]; !ok {
		//	continue
		//}

		signers[addr] = sig
	}

	set.signers = signers
	return set, nil
}

func (s *rootSet) validateSignatures() error {
	if s.aliases == nil {
		log.Error("Empty account aliases")
		return nil
	}

	signers := make(map[common.Address][]byte)
	hash := s.hash.Bytes()

	for _, sig := range s.signers {
		pubkey, err := crypto.SigToPub(hash, sig)
		if err != nil {
			return err
		}

		addr := crypto.PubkeyToAddress(*pubkey)

		exists := false
		for _, alias := range s.aliases {
			if alias == addr {
				exists = true
				break
			}
		}

		if !exists {
			continue
		}

		signers[addr] = sig
	}

	s.signers = signers
	return nil
}

func (s *rootSet) updateAliases(aliases map[common.Address]common.Address) {
	s.aliases = aliases
}

func (s *rootSet) isAcceptable(set *rootSet) bool {
	if set.timestamp <= s.timestamp {
		return false
	}

	if len(s.rootAddresses) == 0 {
		return false
	}

	var sigsCount int
	for signer := range set.signers {
		for root, alias := range s.aliases {
			if (signer == root && signer == alias) || (root != alias && signer == alias) {
				sigsCount++
			}
		}
	}

	percentage := (100 * sigsCount) / len(s.rootAddresses)
	return percentage >= rootListThresholdPercentage
}

func (s *rootSet) isEnoughExSetSignatures(set *exclusionSet) bool {
	var signaturesCount int
	for signer := range set.signers {
		for root, alias := range s.aliases {
			if (signer == root && signer == alias) || (root != alias && signer == alias) {
				signaturesCount++
			}
		}
	}

	percentage := (100 * signaturesCount) / len(s.rootAddresses)
	return percentage >= exclusionListThresholdPercentage
}

// mergeSignatures saves and returns new signatures found in current
func (s *rootSet) mergeSignatures(hash common.Hash, signatures map[common.Address][]byte) map[common.Address][]byte {
	if s.hash != hash {
		return nil
	}

	newSigs := make(map[common.Address][]byte)
	for addr, sig := range signatures {
		if _, ok := s.signers[addr]; ok {
			continue
		}

		s.signers[addr] = sig
		newSigs[addr] = sig
	}

	return newSigs
}

func (s *rootSet) knownSigners(signers map[common.Address][]byte) []common.Address {
	var intersection []common.Address
	for addr := range signers {
		for root, alias := range s.aliases {
			if (addr == root && addr == alias) || (root != alias && addr == alias) {
				intersection = append(intersection, addr)
			}
		}
	}

	return intersection
}

func (s *rootSet) copy() *rootSet {
	if s == nil {
		return nil
	}

	signers := make(map[common.Address][]byte, len(s.signers))
	for signer, signature := range s.signers {
		signers[signer] = signature
	}

	return &rootSet{
		hash:          s.hash,
		timestamp:     s.timestamp,
		rootAddresses: s.rootAddresses,
		aliases:       s.aliases,
		roots:         s.roots,
		signers:       signers,
	}
}

func (s *rootSet) getAddresses() []common.Address {
	if s == nil {
		return nil
	}

	return s.rootAddresses
}

func (s *rootSet) signatures() [][]byte {
	sigs := make([][]byte, 0, len(s.signers))
	for _, sig := range s.signers {
		sigs = append(sigs, sig)
	}

	return sigs
}

func (s *rootSet) makeList() common.RootList {
	if s == nil {
		return common.RootList{}
	}

	return common.RootList{
		Timestamp:  s.timestamp,
		Hash:       s.hash,
		Nodes:      s.rootAddresses,
		Signatures: s.signatures(),
	}
}
