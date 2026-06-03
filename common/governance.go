package common

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"gitlab.com/q-dev/q-client/common/hexutil"
)

// RootList.
type RootList struct {
	Timestamp uint64    `json:"timestamp"`
	Nodes     []Address `json:"nodes"`
	Hash      Hash      `json:"hash"`

	// Signatures use 0x hex in JSON-RPC (hexutil.Bytes); not base64.
	Signatures []hexutil.Bytes `json:"signatures"`
}

// HexSignaturesFromBytes converts raw signature slices for JSON-RPC list payloads.
func HexSignaturesFromBytes(sigs [][]byte) []hexutil.Bytes {
	if len(sigs) == 0 {
		return nil
	}
	out := make([]hexutil.Bytes, len(sigs))
	for i, b := range sigs {
		out[i] = hexutil.Bytes(b)
	}
	return out
}

// ValidatorExclusionList.
type ValidatorExclusionList struct {
	Timestamp  uint64              `json:"timestamp"`
	Validators []ExcludedValidator `json:"validators"`
	Hash       Hash                `json:"hash"`

	// Signatures use 0x hex in JSON-RPC (hexutil.Bytes); not base64.
	Signatures []hexutil.Bytes `json:"signatures"`
}

func (el ValidatorExclusionList) IsEmpty() bool {
	return el.Timestamp == 0 || len(el.Validators) == 0 || el.Hash == Hash{}
}

// ExcludedValidator.
type ExcludedValidator struct {
	Address  Address `json:"address"`
	Block    uint64  `json:"block"`
	EndBlock uint64  `json:"endBlock,omitempty" rlp:"optional"` // If is set - validator stays banned until reaching this block
}

type BlockRange struct {
	StartAddress uint64
	EndAddress   uint64 // New approach to validator blocking mechanism assumes that there can be end of the ban if the end block is set
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

type RootNodeApprovalList struct {
	BlockNumber *big.Int `json:"blockNumber"`
	Approvals   []RootNodeApproval
}

func (l *RootNodeApprovalList) String() string {
	var approvalsString []string
	for _, approval := range l.Approvals {
		approvalsString = append(approvalsString, approval.String())
	}

	return fmt.Sprintf("block number: %s, approvals: [%s]", l.BlockNumber.String(), strings.Join(approvalsString, ", "))
}

type RootNodeApproval struct {
	BlockNumber *big.Int `json:"blockNumber"`
	Hash        Hash     `json:"hash"`
	Signer      Address  `json:"signer"`
	Signature   []byte   `json:"signature"`
}

func (a *RootNodeApproval) String() string {
	return fmt.Sprintf("block number: %s, hash: %s, signer: %s", a.BlockNumber.String(), a.Hash.String(), a.Signer)
}

func (approval RootNodeApproval) GetApprovalDbKey(prefix []byte) (key []byte) {
	key = append(prefix, approval.BlockNumber.Bytes()...)
	return key
}

func (signature RootNodeApproval) Equals(in RootNodeApproval) bool {
	res := true

	res = res && (signature.Hash == in.Hash)
	res = res && signature.BlockNumber == in.BlockNumber
	res = res && signature.Signer == in.Signer
	res = res && bytes.Equal(signature.Signature, in.Signature)

	return res
}

func (list *RootNodeApprovalList) Copy() *RootNodeApprovalList {
	if list == nil {
		return nil
	}

	// copy concurrently mutable list

	return &RootNodeApprovalList{
		BlockNumber: list.BlockNumber,
		Approvals:   list.Approvals,
	}
}

func (RootNodeApprovalList) FillFromArray(arr []RootNodeApproval) *RootNodeApprovalList {
	var res RootNodeApprovalList
	if len(arr) == 0 {
		return nil
	}
	res.BlockNumber = arr[0].BlockNumber
	res.Approvals = append(res.Approvals, arr...)
	return &res
}

type ConstitutionFilesRequest struct {
	Hashes []Hash `json:"hashes"`
}

type KnownConstitutionFilesMessage struct {
	Hashes []Hash `json:"hashes"`
}

type ConstitutionFilesResponse struct {
	Files []ConstitutionFileContent `json:"files"`
}

type ConstitutionFile struct {
	Name      string `json:"name"`
	Hash      Hash   `json:"hash"`
	CreatedAt int64
}

type ConstitutionFileContent struct {
	Hash Hash   `json:"hash"`
	Data []byte `json:"data"`
}

type ListQuotaEntry struct {
	Hash      Hash    `json:"hash"`
	Timestamp uint64  `json:"timestamp"` // determines when the entry was added, not the timestamp of the list as it can be not correct
	Author    Address `json:"author"`
}
