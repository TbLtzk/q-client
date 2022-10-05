package common

import (
	"bytes"
	"math/big"
)

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

type RootNodeApprovalList struct {
	BlockNumber *big.Int `json:"blockNumber"`
	Approvals   []RootNodeApproval
}

type ConstitutionFilesRequest struct {
	//TODO do we need to sign the request? Regular node can start client without the unlocking any account
	Hashes []Hash `json:"hashes"`
}

type ConstitutionFilesResponse struct {
	Files []ConstitutionFileContent `json:"files"`
}

type RootNodeApproval struct {
	BlockNumber *big.Int `json:"blockNumber"`
	Hash        Hash     `json:"hash"`
	Signer      Address  `json:"signer"`
	Signature   []byte   `json:"signature"`
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
	for _, approval := range arr {
		res.Approvals = append(res.Approvals, approval)
	}
	return &res
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
