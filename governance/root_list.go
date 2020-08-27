package governance

import (
	"bytes"
	"fmt"
	"gitlab.com/q-dev/go-ethereum/crypto"
	"sort"
	"sync"

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
	listBytesOfHash := list.BuildSignData()

	if bytes.Compare(listBytesOfHash, list.Hash.Bytes()) != 0 {
		return errors.New("Hash doesn't match")
	}

	percentOfSignatures := (100 * len(list.Signatures)) / len(s.List.Nodes)
	// rounding to the greater
	if float64((100*len(list.Signatures))/len(s.List.Nodes))+float64(0.5) < (float64(100*len(list.Signatures)) / float64(len(s.List.Nodes))) {
		percentOfSignatures++
	}

	if percentOfSignatures < validThreshold {
		return ErrIncomplete
	}

	for _, signature := range list.Signatures {
		pubKey, err := crypto.Ecrecover(listBytesOfHash, signature)
		if err != nil {
			return err
		}

		if !crypto.VerifySignature(pubKey, listBytesOfHash, signature) {
			return ErrInvalidSignature
		}
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

type RootAddresses []common.Address

func (a RootAddresses) Len() int {
	return a.Len()
}

func (a RootAddresses) Less(i, j int) bool {
	return bytes.Compare(a[i].Bytes(), a[j].Bytes()) >= 0
}

func (a RootAddresses) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
