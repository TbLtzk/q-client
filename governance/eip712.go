package governance

import (
	"fmt"
	"math/big"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/common/math"
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/signer/core/apitypes"
)

const governanceEIP712ZeroContract = "0x0000000000000000000000000000000000000000"

var (
	eip712DomainType = []apitypes.Type{
		{Name: "name", Type: "string"},
		{Name: "version", Type: "string"},
		{Name: "chainId", Type: "uint256"},
		{Name: "verifyingContract", Type: "address"},
	}
	rootListProposalType = []apitypes.Type{
		{Name: "proposalType", Type: "string"},
		{Name: "timestamp", Type: "uint256"},
		{Name: "payloadHash", Type: "bytes32"},
		{Name: "nodes", Type: "address[]"},
	}
	excludedValidatorType = []apitypes.Type{
		{Name: "address", Type: "address"},
		{Name: "block", Type: "uint256"},
		{Name: "endBlock", Type: "uint256"},
	}
	exclusionListProposalType = []apitypes.Type{
		{Name: "proposalType", Type: "string"},
		{Name: "timestamp", Type: "uint256"},
		{Name: "payloadHash", Type: "bytes32"},
		{Name: "validators", Type: "QGOVL0ExcludedValidator[]"},
	}
)

func governanceEIP712Domain(chainID uint64) apitypes.TypedDataDomain {
	return apitypes.TypedDataDomain{
		Name:              common.GovernanceSigningDomain,
		Version:           common.GovernanceSigningVersionV1,
		ChainId:           math.NewHexOrDecimal256(int64(chainID)),
		VerifyingContract: governanceEIP712ZeroContract,
	}
}

func rootListEIP712TypedData(chainID uint64, timestamp uint64, payloadHash common.Hash, nodes []common.Address) apitypes.TypedData {
	nodeAddrs := make([]interface{}, len(nodes))
	for i, node := range nodes {
		nodeAddrs[i] = node.Hex()
	}

	return apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain":                          eip712DomainType,
			common.GovernanceEIP712RootListTypeName: rootListProposalType,
		},
		PrimaryType: common.GovernanceEIP712RootListTypeName,
		Domain:      governanceEIP712Domain(chainID),
		Message: map[string]interface{}{
			"proposalType": common.GovernanceProposalTypeRootList,
			"timestamp":    uint256Value(timestamp),
			"payloadHash":  payloadHash.Hex(),
			"nodes":        nodeAddrs,
		},
	}
}

func exclusionListEIP712TypedData(chainID uint64, timestamp uint64, payloadHash common.Hash, validators []common.ExcludedValidator) apitypes.TypedData {
	entries := make([]interface{}, len(validators))
	for i, val := range validators {
		entries[i] = map[string]interface{}{
			"address":  val.Address.Hex(),
			"block":    uint256Value(val.Block),
			"endBlock": uint256Value(val.EndBlock),
		}
	}

	return apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain":                               eip712DomainType,
			"QGOVL0ExcludedValidator":                    excludedValidatorType,
			common.GovernanceEIP712ExclusionListTypeName: exclusionListProposalType,
		},
		PrimaryType: common.GovernanceEIP712ExclusionListTypeName,
		Domain:      governanceEIP712Domain(chainID),
		Message: map[string]interface{}{
			"proposalType": common.GovernanceProposalTypeExclusionList,
			"timestamp":    uint256Value(timestamp),
			"payloadHash":  payloadHash.Hex(),
			"validators":   entries,
		},
	}
}

func uint256Value(v uint64) *math.HexOrDecimal256 {
	return math.NewHexOrDecimal256(int64(v))
}

func eip712Digest(td apitypes.TypedData) (common.Hash, error) {
	hash, _, err := apitypes.TypedDataAndHash(td)
	if err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(hash), nil
}

func buildRootListEIP712(chainID uint64, list common.RootList) (common.GovernanceEIP712TypedData, error) {
	unsigned := list
	unsigned.Signatures = nil

	set, err := newRootSet(&unsigned)
	if err != nil {
		return common.GovernanceEIP712TypedData{}, err
	}
	if set == nil {
		return common.GovernanceEIP712TypedData{}, fmt.Errorf("root list has no nodes")
	}
	if list.Hash != (common.Hash{}) && set.hash != list.Hash {
		return common.GovernanceEIP712TypedData{}, errHashMismatch
	}

	td := rootListEIP712TypedData(chainID, set.timestamp, set.hash, set.rootAddresses)
	return toPublicEIP712TypedData(td), nil
}

func buildExclusionListEIP712(chainID uint64, list common.ValidatorExclusionList) (common.GovernanceEIP712TypedData, error) {
	unsigned := list
	unsigned.Signatures = nil

	set, err := newExclusionSet(&unsigned)
	if err != nil {
		return common.GovernanceEIP712TypedData{}, err
	}
	if set == nil {
		return common.GovernanceEIP712TypedData{}, fmt.Errorf("invalid exclusion list proposal")
	}
	if list.Hash != (common.Hash{}) && set.hash != list.Hash {
		return common.GovernanceEIP712TypedData{}, errHashMismatch
	}

	source := set.makeList()
	source.Signatures = nil
	validators := canonicalExcludedValidatorsForEIP712(source.Validators)

	td := exclusionListEIP712TypedData(chainID, set.timestamp, set.hash, validators)
	return toPublicEIP712TypedData(td), nil
}

func rootListEIP712Digest(chainID uint64, timestamp uint64, payloadHash common.Hash, nodes []common.Address) (common.Hash, error) {
	return eip712Digest(rootListEIP712TypedData(chainID, timestamp, payloadHash, nodes))
}

func exclusionListEIP712Digest(chainID uint64, timestamp uint64, payloadHash common.Hash, validators []common.ExcludedValidator) (common.Hash, error) {
	return eip712Digest(exclusionListEIP712TypedData(chainID, timestamp, payloadHash, validators))
}

func verifyRootListEIP712Signature(chainID uint64, set *rootSet, sig []byte) (common.Address, bool) {
	if chainID == 0 || set == nil {
		return common.Address{}, false
	}
	digest, err := rootListEIP712Digest(chainID, set.timestamp, set.hash, set.rootAddresses)
	if err != nil {
		return common.Address{}, false
	}
	pubkey, err := crypto.SigToPub(digest.Bytes(), normalizeECDSASignatureV(sig))
	if err != nil {
		return common.Address{}, false
	}
	return crypto.PubkeyToAddress(*pubkey), true
}

func verifyExclusionListEIP712Signature(chainID uint64, set *exclusionSet, sig []byte) (common.Address, bool) {
	if chainID == 0 || set == nil {
		return common.Address{}, false
	}
	unsigned := set.makeList()
	unsigned.Signatures = nil
	validators := canonicalExcludedValidatorsForEIP712(unsigned.Validators)

	digest, err := exclusionListEIP712Digest(chainID, set.timestamp, set.hash, validators)
	if err != nil {
		return common.Address{}, false
	}
	pubkey, err := crypto.SigToPub(digest.Bytes(), normalizeECDSASignatureV(sig))
	if err != nil {
		return common.Address{}, false
	}
	return crypto.PubkeyToAddress(*pubkey), true
}

// canonicalExcludedValidatorsForEIP712 uses the same validator ordering as list parsing.
func canonicalExcludedValidatorsForEIP712(validators []common.ExcludedValidator) []common.ExcludedValidator {
	return common.NewExclusionListSigningPayloadV1(0, common.ValidatorExclusionList{Validators: validators}).Validators
}

func toPublicEIP712TypedData(td apitypes.TypedData) common.GovernanceEIP712TypedData {
	out := common.GovernanceEIP712TypedData{
		PrimaryType: td.PrimaryType,
		Message:     td.Message,
		Types:       make(map[string][]common.GovernanceEIP712TypeField, len(td.Types)),
	}
	for typeName, fields := range td.Types {
		typedFields := make([]common.GovernanceEIP712TypeField, len(fields))
		for i, field := range fields {
			typedFields[i] = common.GovernanceEIP712TypeField{Name: field.Name, Type: field.Type}
		}
		out.Types[typeName] = typedFields
	}
	chainID := ""
	if td.Domain.ChainId != nil {
		chainID = (*big.Int)(td.Domain.ChainId).String()
	}
	out.Domain = common.GovernanceEIP712Domain{
		Name:              td.Domain.Name,
		Version:           td.Domain.Version,
		ChainId:           chainID,
		VerifyingContract: td.Domain.VerifyingContract,
	}
	return out
}

func fromPublicEIP712TypedData(data common.GovernanceEIP712TypedData) (apitypes.TypedData, error) {
	types := make(apitypes.Types, len(data.Types))
	for typeName, fields := range data.Types {
		typedFields := make([]apitypes.Type, len(fields))
		for i, field := range fields {
			typedFields[i] = apitypes.Type{Name: field.Name, Type: field.Type}
		}
		types[typeName] = typedFields
	}

	chainID, ok := math.ParseBig256(data.Domain.ChainId)
	if !ok {
		return apitypes.TypedData{}, fmt.Errorf("invalid chainId: %s", data.Domain.ChainId)
	}

	return apitypes.TypedData{
		Types:       types,
		PrimaryType: data.PrimaryType,
		Domain: apitypes.TypedDataDomain{
			Name:              data.Domain.Name,
			Version:           data.Domain.Version,
			ChainId:           (*math.HexOrDecimal256)(chainID),
			VerifyingContract: data.Domain.VerifyingContract,
		},
		Message: data.Message,
	}, nil
}
