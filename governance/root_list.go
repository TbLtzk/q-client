package governance

import (
	"gitlab.com/q-dev/go-ethereum/common"
	"sync"
)

type RootSet struct {
	lock sync.Mutex
	List RootList
}

func (s *RootSet) CurrentList() RootList {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.List
}

type RootList struct {
	Timestamp int64            `json:"timestamp"`
	Nodes     []common.Address `json:"nodes"`
	Digest    common.Hash      `json:"digest"`

	Signatures []string `json:"signatures"`
}

func (l *RootList) SignData() []byte {
	// TODO: implement me
	// build {timestamp + ordered list of addresses}

	return nil
}
