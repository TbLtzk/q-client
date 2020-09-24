package common

// RootList.
type RootList struct {
	Timestamp uint64    `json:"timestamp"`
	Nodes     []Address `json:"nodes"`
	Hash      Hash      `json:"hash"`

	Signatures [][]byte `json:"signatures"`
}

// ValidatorExclusionList.
type ValidatorExclusionList struct {
	Timestamp  uint64              `json:"timestamp"`
	Validators []ExcludedValidator `json:"validators"`
	Hash       Hash                `json:"hash"`

	Signatures [][]byte `json:"signatures"`
}

// ExcludedValidator.
type ExcludedValidator struct {
	Address Address `json:"address"`
	Block   uint64  `json:"block"`
}
