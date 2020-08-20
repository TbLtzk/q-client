package governance

import "gitlab.com/q-dev/go-ethereum/common"

type RootSet struct {
	List RootList
}

type RootList struct {
	Timestamp int64            `json:"timestamp"`
	Nodes     []common.Address `json:"nodes"`

	Signatures []string `json:"signatures"`
}
