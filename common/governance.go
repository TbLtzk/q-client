package common

// RootList.
type RootList struct {
	Timestamp uint64    `json:"timestamp"`
	Nodes     []Address `json:"nodes"`
	Hash      Hash      `json:"hash"`

	Signatures [][]byte `json:"signatures"`
}
