package governance

import (
	"sync"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/go-ethereum/common"
)

var (
	ErrInvalidSignature = errors.New("list contains invalid signature")
	ErrIncomplete       = errors.New("not enough signatures")
)

//const validThreshold = 75

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

func (s *RootList) Validate(list RootList) error {
	// todo: implement me
	return nil
}

type RootList struct {
	Timestamp uint64           `json:"timestamp"`
	Nodes     []common.Address `json:"nodes"`
	Hash      common.Hash      `json:"hash"`

	Signatures [][]byte `json:"signatures"`
}

func (l *RootList) SignData() []byte {
	// TODO: implement me
	// build {timestamp + ordered list of addresses}

	return nil
}

type RootAddresses []common.Address

// TODO: implement interface
func (a RootAddresses) Len() int {
	panic("implement me")
}

func (a RootAddresses) Less(i, j int) bool {
	panic("implement me")
}

func (a RootAddresses) Swap(i, j int) {
	panic("implement me")
}
