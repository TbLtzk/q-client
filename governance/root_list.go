package governance

import (
	"bytes"
	"fmt"
	"sort"
	"sync"

	"gitlab.com/q-dev/go-ethereum/crypto"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/go-ethereum/common"
)

var (
	ErrInvalidSignature = errors.New("list contains invalid signature")
	ErrIncomplete       = errors.New("not enough signatures")
)

const validThreshold = 75

type RootSet struct {
	lock sync.Mutex
	List RootList
}

func (s *RootSet) CurrentList() RootList {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.List
}

func (s *RootSet) UpdateList(newList RootList) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.List = newList
}

func (s *RootSet) Validate(list RootList) error {
	listBytesOfHash := crypto.Keccak256(list.BuildSignData())

	if !bytes.Equal(listBytesOfHash, list.Hash.Bytes()) {
		return errors.New("Hash doesn't match")
	}

	rootKeys := make(map[common.Address]bool)
	for _, root := range s.CurrentList().Nodes {
		rootKeys[root] = true
	}
	var successfullSigantures int
	for _, signature := range list.Signatures {
		pubkey, err := crypto.Ecrecover(listBytesOfHash, signature)
		if err != nil {
			return err
		}

		if rootKeys[common.BytesToAddress(pubkey)] {
			successfullSigantures++
		} else {
			return ErrInvalidSignature
		}
	}

	percentOfSignatures := (100 * successfullSigantures) / len(s.CurrentList().Nodes)

	if percentOfSignatures < validThreshold {
		return ErrIncomplete
	}

	return nil
}

type RootList struct {
	Timestamp uint64           `json:"timestamp"`
	Nodes     []common.Address `json:"nodes"`
	Hash      common.Hash      `json:"hash"`

	Signatures [][]byte `json:"signatures"`
}

func (l *RootList) BuildSignData() []byte {
	// build {timestamp + ordered list of addresses}
	var toSign []byte

	toSign = append(toSign, fmt.Sprint(l.Timestamp)...)

	sort.SliceStable(l.Nodes, func(i, j int) bool { return bytes.Compare(l.Nodes[i].Bytes(), l.Nodes[j].Bytes()) >= 0 })
	for _, node := range l.Nodes {
		toSign = append(toSign, node.Bytes()...)
	}

	return toSign
}
