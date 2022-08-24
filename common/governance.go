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
	Address  Address `json:"address"`
	Block    uint64  `json:"block"`
	EndBlock uint64  `json:"endBlock,omitempty" rlp:"optional"` //If is set - validator stays banned until reaching this block
}

type BlockRange struct {
	StartAddress uint64
	EndAddress   uint64 //New approach to validator blocking mechanism assumes that there can be end of the ban if the end block is set
}

func (existingBlockRange BlockRange) IsEqualTo(candidate BlockRange) (res bool) {
	return existingBlockRange.StartAddress == candidate.StartAddress && existingBlockRange.EndAddress == candidate.EndAddress
}

func (existingBlockRange BlockRange) IsClosed() bool {
	return existingBlockRange.EndAddress > 0
}

func (existingBlockRange BlockRange) IsValid() bool {
	return existingBlockRange.StartAddress != 0 && (existingBlockRange.EndAddress == 0 || existingBlockRange.EndAddress > existingBlockRange.StartAddress)
}

func (existingBlockRange BlockRange) EndsInFuture(currentBlock uint64) bool {
	return existingBlockRange.EndAddress > currentBlock || existingBlockRange.EndAddress == 0
}

func (existingBlockRange BlockRange) StartsInFuture(currentBlock uint64) bool {
	return existingBlockRange.StartAddress > currentBlock
}

func (existingBlockRange BlockRange) IsAfterBlockRange(candidate BlockRange) bool {
	return candidate.IsClosed() && existingBlockRange.StartAddress > candidate.EndAddress
}

func (existingBlockRange BlockRange) IsBeforeBlockRange(candidate BlockRange) bool {
	return existingBlockRange.IsClosed() && candidate.StartAddress > existingBlockRange.EndAddress
}

func (existingBlockRange BlockRange) StartsWithTheSameBlock(candidate BlockRange) bool {
	return existingBlockRange.StartAddress == candidate.StartAddress
}

func (existingBlockRange BlockRange) IntersectsWithRange(candidate BlockRange) (res bool) {
	return !(existingBlockRange.IsAfterBlockRange(candidate) || candidate.IsAfterBlockRange(existingBlockRange))
}

func (existingBlockRange BlockRange) ContainsAddress(address uint64) (res bool) {
	return (existingBlockRange.IsClosed() && existingBlockRange.StartAddress <= address && existingBlockRange.EndAddress >= address) ||
		(!existingBlockRange.IsClosed() && existingBlockRange.StartAddress <= address)
}
