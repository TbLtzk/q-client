package governance

import (
	"bytes"
	"fmt"
	"sort"
	"sync"

	mapset "github.com/deckarep/golang-set"
	"gitlab.com/q-dev/go-ethereum/common"
	"gitlab.com/q-dev/go-ethereum/crypto"
)

const validListThresholdPercentage = 75

type rootSet struct {
	timestamp uint64
	hash      common.Hash

	rootAddresses []common.Address
	roots         map[common.Address]struct{}

	lock    sync.Mutex
	signers map[common.Address][]byte
}

func (s *rootSet) addSignature(signer common.Address, signature []byte) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

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
		hash:          list.Hash,
	}

	if (set.hash == common.Hash{}) {
		set.hash = set.calcHash()
	}

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
		if _, ok := roots[addr]; !ok {
			continue
		}

		signers[addr] = sig
	}

	set.signers = signers
	return set, nil
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
		if _, ok := s.roots[signer]; ok {
			sigsCount++
		}
	}

	percentage := (100 * sigsCount) / len(s.rootAddresses)
	return percentage >= validListThresholdPercentage
}

// mergeSignatures saves and returns new signatures found in current
func (s *rootSet) mergeSignatures(hash common.Hash, signatures map[common.Address][]byte) map[common.Address][]byte {
	s.lock.Lock()
	defer s.lock.Unlock()

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

func (s *rootSet) sanitizeSignatures(signatures [][]byte) (map[common.Address][]byte, error) {
	sigs := make(map[common.Address][]byte)
	for _, sig := range signatures {
		pubkey, err := crypto.SigToPub(s.hash.Bytes(), sig)
		if err != nil {
			return nil, err
		}

		addr := crypto.PubkeyToAddress(*pubkey)
		if _, ok := s.roots[addr]; !ok {
			continue
		}

		sigs[addr] = sig
	}

	return sigs, nil
}

func (s *rootSet) copy() *rootSet {
	s.lock.Lock()
	defer s.lock.Unlock()

	signers := make(map[common.Address][]byte, len(s.signers))
	for signer, signature := range s.signers {
		signers[signer] = signature
	}

	return &rootSet{
		hash:          s.hash,
		timestamp:     s.timestamp,
		rootAddresses: s.rootAddresses,
		roots:         s.roots,
		signers:       signers,
	}
}

func (s *rootSet) signatures() [][]byte {
	s.lock.Lock()
	defer s.lock.Unlock()

	sigs := make([][]byte, 0, len(s.signers))
	for _, sig := range s.signers {
		sigs = append(sigs, sig)
	}

	return sigs
}

func (s *rootSet) makeList() common.RootList {
	return common.RootList{
		Timestamp:  s.timestamp,
		Hash:       s.hash,
		Nodes:      s.rootAddresses,
		Signatures: s.signatures(),
	}
}

func (s *rootSet) signersMap() mapset.Set {
	s.lock.Lock()
	defer s.lock.Unlock()

	set := mapset.NewSet()
	for signer := range s.signers {
		set.Add(signer)
	}

	return set
}

func toMapSet(signatures map[common.Address][]byte) mapset.Set {
	set := mapset.NewSet()
	for addr := range signatures {
		set.Add(addr)
	}

	return set
}
