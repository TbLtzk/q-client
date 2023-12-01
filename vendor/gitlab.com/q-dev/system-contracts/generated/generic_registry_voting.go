// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

import (
	"errors"
	"math/big"
	"strings"

	ethereum "gitlab.com/q-dev/q-client"
	"gitlab.com/q-dev/q-client/accounts/abi"
	"gitlab.com/q-dev/q-client/accounts/abi/bind"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ARootNodeApprovalVotingAProposal is an auto generated low-level Go binding around an user-defined struct.
type ARootNodeApprovalVotingAProposal struct {
	Executed          bool
	VotingStartTime   *big.Int
	VotingExpiredTime *big.Int
	RequiredMajority  *big.Int
	CallData          []byte
	Remark            string
}

// ARootNodeApprovalVotingProposalStats is an auto generated low-level Go binding around an user-defined struct.
type ARootNodeApprovalVotingProposalStats struct {
	RequiredMajority *big.Int
	CurrentMajority  *big.Int
}

// GenericContractRegistryVotingMetaData contains all meta data concerning the GenericContractRegistryVoting contract.
var GenericContractRegistryVotingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rootNode\",\"type\":\"address\"}],\"name\":\"RootNodeApproved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"remark_\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"callData_\",\"type\":\"bytes\"}],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingExpiredTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"}],\"internalType\":\"structARootNodeApprovalVoting.AProposal\",\"name\":\"proposal_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"name\":\"getProposalStats\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentMajority\",\"type\":\"uint256\"}],\"internalType\":\"structARootNodeApprovalVoting.ProposalStats\",\"name\":\"stats_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumIVoting.ProposalStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voteCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611726806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a45760003560e01c80630a9f46ad146100a9578063307a064f146100c55780634fc8a20d146100f35780635277b4ae146101135780635c622a0e14610151578063715018a6146101715780638da5cb5b1461017b578063b759f95414610190578063c4d66de8146101a3578063c7f758a8146101b6578063d182fb9e146101d6578063f2fde38b146101e9575b600080fd5b6100b260035481565b6040519081526020015b60405180910390f35b6100d86100d336600461126b565b6101fc565b604080518251815260209283015192810192909252016100bc565b6100b261010136600461126b565b60026020526000908152604090205481565b610141610121366004611299565b600160209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016100bc565b61016461015f36600461126b565b610239565b6040516100bc91906112c9565b6101796103f4565b005b610183610438565b6040516100bc91906112f1565b61017961019e36600461126b565b610447565b6101796101b1366004611305565b61062f565b6101c96101c436600461126b565b6106bf565b6040516100bc919061137e565b6100b26101e4366004611476565b61087c565b6101796101f7366004611305565b6108be565b60408051808201909152600080825260208201526102198261095e565b602080830191909152600092835260049052604090912060030154815290565b6000600354821061024c57506000919050565b6000828152600460208181526040808420815160c081018352815460ff16151581526001820154938101939093526002810154918301919091526003810154606083015291820180549192916080840191906102a790611505565b80601f01602080910402602001604051908101604052809291908181526020018280546102d390611505565b80156103205780601f106102f557610100808354040283529160200191610320565b820191906000526020600020905b81548152906001019060200180831161030357829003601f168201915b5050505050815260200160058201805461033990611505565b80601f016020809104026020016040519081016040528092919081815260200182805461036590611505565b80156103b25780601f10610387576101008083540402835291602001916103b2565b820191906000526020600020905b81548152906001019060200180831161039557829003601f168201915b50505050508152505090508060000151156103d05750600592915050565b80604001514210156103e55750600192915050565b50600792915050565b50919050565b336103fd610438565b6001600160a01b03161461042c5760405162461bcd60e51b81526004016104239061153a565b60405180910390fd5b6104366000610a0c565b565b6038546001600160a01b031690565b8061045181610a5e565b3361045b81610ac2565b6000838152600460205260409020805460ff16156104bb5760405162461bcd60e51b815260206004820152601e60248201527f5b5145432d3033393030305d2d416c72656164792065786563757465642e00006044820152606401610423565b8060020154421061051c5760405162461bcd60e51b815260206004820152602560248201527f5b5145432d3033393030315d2d566f74696e672074696d65206861732065787060448201526434b932b21760d91b6064820152608401610423565b600084815260016020908152604080832033845290915281205460ff166105b6575060008481526001602081815260408084203385528252808420805460ff191684179055878452600290915282208054919261057883611585565b9091555050604080518681523360208201527fdcd66ff3278394a103acd0febedb4f0cfae077df25e5b1e05b6b214f3669dd30910160405180910390a15b60006105c186610bb7565b905081806105cc5750805b6106275760405162461bcd60e51b815260206004820152602660248201527f5b5145432d3033393030325d2d53656e6465722068617320616c7265616479206044820152653b37ba32b21760d11b6064820152608401610423565b505050505050565b600554610100900460ff1680610648575060055460ff16155b6106645760405162461bcd60e51b8152600401610423906115a0565b600554610100900460ff16158015610686576005805461ffff19166101011790555b600080546001600160a01b0319166001600160a01b0384161790556106a9610cb9565b80156106bb576005805461ff00191690555b5050565b6106fa6040518060c0016040528060001515815260200160008152602001600081526020016000815260200160608152602001606081525090565b8161070481610a5e565b600083815260046020818152604092839020835160c081018552815460ff161515815260018201549281019290925260028101549382019390935260038301546060820152908201805491929160808401919061076090611505565b80601f016020809104026020016040519081016040528092919081815260200182805461078c90611505565b80156107d95780601f106107ae576101008083540402835291602001916107d9565b820191906000526020600020905b8154815290600101906020018083116107bc57829003601f168201915b505050505081526020016005820180546107f290611505565b80601f016020809104026020016040519081016040528092919081815260200182805461081e90611505565b801561086b5780601f106108405761010080835404028352916020019161086b565b820191906000526020600020905b81548152906001019060200180831161084e57829003601f168201915b505050505081525050915050919050565b600033610887610438565b6001600160a01b0316146108ad5760405162461bcd60e51b81526004016104239061153a565b6108b78383610d34565b9392505050565b336108c7610438565b6001600160a01b0316146108ed5760405162461bcd60e51b81526004016104239061153a565b6001600160a01b0381166109525760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610423565b61095b81610a0c565b50565b600080610969610f4f565b6001600160a01b031663de8fa4316040518163ffffffff1660e01b815260040160206040518083038186803b1580156109a157600080fd5b505afa1580156109b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d991906115ee565b9050806b033b2e3c9fd0803ce8000000600085815260026020526040902054610a029190611607565b6108b79190611626565b603880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600354811061095b5760405162461bcd60e51b815260206004820152602a60248201527f5b5145432d3033393030345d2d50726f706f73616c20686176656e277420696e60448201526934ba34b0b634bd32b21760b11b6064820152608401610423565b610aca610f4f565b6001600160a01b031663a230c524826040518263ffffffff1660e01b8152600401610af591906112f1565b60206040518083038186803b158015610b0d57600080fd5b505afa158015610b21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b459190611648565b61095b5760405162461bcd60e51b815260206004820152603d60248201527f5b5145432d3033393030335d2d5065726d697373696f6e2064656e696564202d60448201527f206f6e6c7920726f6f74206e6f6465732068617665206163636573732e0000006064820152608401610423565b60008181526004602052604081206003810154610bd38461095e565b1115610caf57805460ff19166001178155600481018054610c7b9190610bf890611505565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2490611505565b8015610c715780601f10610c4657610100808354040283529160200191610c71565b820191906000526020600020905b815481529060010190602001808311610c5457829003601f168201915b5050505050610ffc565b6040518381527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f9060200160405180910390a15b5460ff1692915050565b600554610100900460ff1680610cd2575060055460ff16155b610cee5760405162461bcd60e51b8152600401610423906115a0565b600554610100900460ff16158015610d10576005805461ffff19166101011790555b610d186110bd565b610d20611127565b801561095b576005805461ff001916905550565b6003805460009182919082610d4883611585565b9190505590506000610d58611187565b60405162498bff60e81b815260206004820152601f60248201527f636f6e737469747574696f6e2e70726f706f73616c457865637574696f6e500060448201526001600160a01b03919091169063498bff009060640160206040518083038186803b158015610dc657600080fd5b505afa158015610dda573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dfe91906115ee565b6000838152600460208181526040909220875193945092610e2592918401918801906111d2565b50610e2e611187565b60405162498bff60e81b815260206004820152602260248201527f636f6e737469747574696f6e2e766f74696e672e656d6751557064617465524d60448201526120a560f11b60648201526001600160a01b03919091169063498bff009060840160206040518083038186803b158015610ea757600080fd5b505afa158015610ebb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610edf91906115ee565b60038201554260018201819055610ef790839061166a565b60028201558551610f1190600583019060208901906111d2565b506040518381527fc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe079060200160405180910390a15090949350505050565b600080546040805180820182526014815273676f7665726e616e63652e726f6f744e6f64657360601b60208201529051633fb9027160e01b81526001600160a01b0390921691633fb9027191610fa791600401611682565b60206040518083038186803b158015610fbf57600080fd5b505afa158015610fd3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff79190611695565b905090565b600080546040516001600160a01b039091169061101a9084906116b2565b6000604051808303816000865af19150503d8060008114611057576040519150601f19603f3d011682016040523d82523d6000602084013e61105c565b606091505b50509050806106bb5760405162461bcd60e51b815260206004820152602760248201527f5b5145432d3033393030365d2d50726f706f73616c20657865637574696f6e206044820152663330b4b632b21760c91b6064820152608401610423565b600554610100900460ff16806110d6575060055460ff16155b6110f25760405162461bcd60e51b8152600401610423906115a0565b600554610100900460ff16158015610d20576005805461ffff1916610101179055801561095b576005805461ff001916905550565b600554610100900460ff1680611140575060055460ff16155b61115c5760405162461bcd60e51b8152600401610423906115a0565b600554610100900460ff1615801561117e576005805461ffff19166101011790555b610d2033610a0c565b60008054604080516060810190915260228082526001600160a01b0390921691633fb9027191906116cf60208301396040518263ffffffff1660e01b8152600401610fa79190611682565b8280546111de90611505565b90600052602060002090601f0160209004810192826112005760008555611246565b82601f1061121957805160ff1916838001178555611246565b82800160010185558215611246579182015b8281111561124657825182559160200191906001019061122b565b50611252929150611256565b5090565b5b808211156112525760008155600101611257565b60006020828403121561127d57600080fd5b5035919050565b6001600160a01b038116811461095b57600080fd5b600080604083850312156112ac57600080fd5b8235915060208301356112be81611284565b809150509250929050565b60208101600883106112eb57634e487b7160e01b600052602160045260246000fd5b91905290565b6001600160a01b0391909116815260200190565b60006020828403121561131757600080fd5b81356108b781611284565b60005b8381101561133d578181015183820152602001611325565b8381111561134c576000848401525b50505050565b6000815180845261136a816020860160208601611322565b601f01601f19169290920160200192915050565b602081528151151560208201526020820151604082015260408201516060820152606082015160808201526000608083015160c060a08401526113c460e0840182611352565b905060a0840151601f198483030160c08501526113e18282611352565b95945050505050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff8084111561141b5761141b6113ea565b604051601f8501601f19908116603f01168101908282118183101715611443576114436113ea565b8160405280935085815286868601111561145c57600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561148957600080fd5b823567ffffffffffffffff808211156114a157600080fd5b818501915085601f8301126114b557600080fd5b6114c486833560208501611400565b935060208501359150808211156114da57600080fd5b508301601f810185136114ec57600080fd5b6114fb85823560208401611400565b9150509250929050565b600181811c9082168061151957607f821691505b602082108114156103ee57634e487b7160e01b600052602260045260246000fd5b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052601160045260246000fd5b60006000198214156115995761159961156f565b5060010190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b60006020828403121561160057600080fd5b5051919050565b60008160001904831182151516156116215761162161156f565b500290565b60008261164357634e487b7160e01b600052601260045260246000fd5b500490565b60006020828403121561165a57600080fd5b815180151581146108b757600080fd5b6000821982111561167d5761167d61156f565b500190565b6020815260006108b76020830184611352565b6000602082840312156116a757600080fd5b81516108b781611284565b600082516116c4818460208701611322565b919091019291505056fe676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273a264697066735822122023a5beb2a9bea9af514d076fd56c46c137705debd029c21d242fcd91177fb90f64736f6c63430008090033",
}

// GenericContractRegistryVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use GenericContractRegistryVotingMetaData.ABI instead.
var GenericContractRegistryVotingABI = GenericContractRegistryVotingMetaData.ABI

// GenericContractRegistryVotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GenericContractRegistryVotingMetaData.Bin instead.
var GenericContractRegistryVotingBin = GenericContractRegistryVotingMetaData.Bin

// DeployGenericContractRegistryVoting deploys a new Ethereum contract, binding an instance of GenericContractRegistryVoting to it.
func DeployGenericContractRegistryVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GenericContractRegistryVoting, error) {
	parsed, err := GenericContractRegistryVotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GenericContractRegistryVotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenericContractRegistryVoting{GenericContractRegistryVotingCaller: GenericContractRegistryVotingCaller{contract: contract}, GenericContractRegistryVotingTransactor: GenericContractRegistryVotingTransactor{contract: contract}, GenericContractRegistryVotingFilterer: GenericContractRegistryVotingFilterer{contract: contract}}, nil
}

// GenericContractRegistryVoting is an auto generated Go binding around an Ethereum contract.
type GenericContractRegistryVoting struct {
	GenericContractRegistryVotingCaller     // Read-only binding to the contract
	GenericContractRegistryVotingTransactor // Write-only binding to the contract
	GenericContractRegistryVotingFilterer   // Log filterer for contract events
}

// GenericContractRegistryVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenericContractRegistryVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericContractRegistryVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericContractRegistryVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericContractRegistryVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericContractRegistryVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericContractRegistryVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericContractRegistryVotingSession struct {
	Contract     *GenericContractRegistryVoting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// GenericContractRegistryVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericContractRegistryVotingCallerSession struct {
	Contract *GenericContractRegistryVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// GenericContractRegistryVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericContractRegistryVotingTransactorSession struct {
	Contract     *GenericContractRegistryVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// GenericContractRegistryVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenericContractRegistryVotingRaw struct {
	Contract *GenericContractRegistryVoting // Generic contract binding to access the raw methods on
}

// GenericContractRegistryVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericContractRegistryVotingCallerRaw struct {
	Contract *GenericContractRegistryVotingCaller // Generic read-only contract binding to access the raw methods on
}

// GenericContractRegistryVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericContractRegistryVotingTransactorRaw struct {
	Contract *GenericContractRegistryVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericContractRegistryVoting creates a new instance of GenericContractRegistryVoting, bound to a specific deployed contract.
func NewGenericContractRegistryVoting(address common.Address, backend bind.ContractBackend) (*GenericContractRegistryVoting, error) {
	contract, err := bindGenericContractRegistryVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericContractRegistryVoting{GenericContractRegistryVotingCaller: GenericContractRegistryVotingCaller{contract: contract}, GenericContractRegistryVotingTransactor: GenericContractRegistryVotingTransactor{contract: contract}, GenericContractRegistryVotingFilterer: GenericContractRegistryVotingFilterer{contract: contract}}, nil
}

// NewGenericContractRegistryVotingCaller creates a new read-only instance of GenericContractRegistryVoting, bound to a specific deployed contract.
func NewGenericContractRegistryVotingCaller(address common.Address, caller bind.ContractCaller) (*GenericContractRegistryVotingCaller, error) {
	contract, err := bindGenericContractRegistryVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericContractRegistryVotingCaller{contract: contract}, nil
}

// NewGenericContractRegistryVotingTransactor creates a new write-only instance of GenericContractRegistryVoting, bound to a specific deployed contract.
func NewGenericContractRegistryVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*GenericContractRegistryVotingTransactor, error) {
	contract, err := bindGenericContractRegistryVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericContractRegistryVotingTransactor{contract: contract}, nil
}

// NewGenericContractRegistryVotingFilterer creates a new log filterer instance of GenericContractRegistryVoting, bound to a specific deployed contract.
func NewGenericContractRegistryVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*GenericContractRegistryVotingFilterer, error) {
	contract, err := bindGenericContractRegistryVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericContractRegistryVotingFilterer{contract: contract}, nil
}

// bindGenericContractRegistryVoting binds a generic wrapper to an already deployed contract.
func bindGenericContractRegistryVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericContractRegistryVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericContractRegistryVoting *GenericContractRegistryVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericContractRegistryVoting.Contract.GenericContractRegistryVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericContractRegistryVoting *GenericContractRegistryVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.GenericContractRegistryVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericContractRegistryVoting *GenericContractRegistryVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.GenericContractRegistryVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericContractRegistryVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericContractRegistryVoting *GenericContractRegistryVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericContractRegistryVoting *GenericContractRegistryVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.contract.Transact(opts, method, params...)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 id_) view returns((bool,uint256,uint256,uint256,bytes,string) proposal_)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCaller) GetProposal(opts *bind.CallOpts, id_ *big.Int) (ARootNodeApprovalVotingAProposal, error) {
	var out []interface{}
	err := _GenericContractRegistryVoting.contract.Call(opts, &out, "getProposal", id_)

	if err != nil {
		return *new(ARootNodeApprovalVotingAProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(ARootNodeApprovalVotingAProposal)).(*ARootNodeApprovalVotingAProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 id_) view returns((bool,uint256,uint256,uint256,bytes,string) proposal_)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingSession) GetProposal(id_ *big.Int) (ARootNodeApprovalVotingAProposal, error) {
	return _GenericContractRegistryVoting.Contract.GetProposal(&_GenericContractRegistryVoting.CallOpts, id_)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 id_) view returns((bool,uint256,uint256,uint256,bytes,string) proposal_)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCallerSession) GetProposal(id_ *big.Int) (ARootNodeApprovalVotingAProposal, error) {
	return _GenericContractRegistryVoting.Contract.GetProposal(&_GenericContractRegistryVoting.CallOpts, id_)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 id_) view returns((uint256,uint256) stats_)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCaller) GetProposalStats(opts *bind.CallOpts, id_ *big.Int) (ARootNodeApprovalVotingProposalStats, error) {
	var out []interface{}
	err := _GenericContractRegistryVoting.contract.Call(opts, &out, "getProposalStats", id_)

	if err != nil {
		return *new(ARootNodeApprovalVotingProposalStats), err
	}

	out0 := *abi.ConvertType(out[0], new(ARootNodeApprovalVotingProposalStats)).(*ARootNodeApprovalVotingProposalStats)

	return out0, err

}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 id_) view returns((uint256,uint256) stats_)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingSession) GetProposalStats(id_ *big.Int) (ARootNodeApprovalVotingProposalStats, error) {
	return _GenericContractRegistryVoting.Contract.GetProposalStats(&_GenericContractRegistryVoting.CallOpts, id_)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 id_) view returns((uint256,uint256) stats_)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCallerSession) GetProposalStats(id_ *big.Int) (ARootNodeApprovalVotingProposalStats, error) {
	return _GenericContractRegistryVoting.Contract.GetProposalStats(&_GenericContractRegistryVoting.CallOpts, id_)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 id_) view returns(uint8)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCaller) GetStatus(opts *bind.CallOpts, id_ *big.Int) (uint8, error) {
	var out []interface{}
	err := _GenericContractRegistryVoting.contract.Call(opts, &out, "getStatus", id_)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 id_) view returns(uint8)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingSession) GetStatus(id_ *big.Int) (uint8, error) {
	return _GenericContractRegistryVoting.Contract.GetStatus(&_GenericContractRegistryVoting.CallOpts, id_)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 id_) view returns(uint8)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCallerSession) GetStatus(id_ *big.Int) (uint8, error) {
	return _GenericContractRegistryVoting.Contract.GetStatus(&_GenericContractRegistryVoting.CallOpts, id_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GenericContractRegistryVoting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingSession) Owner() (common.Address, error) {
	return _GenericContractRegistryVoting.Contract.Owner(&_GenericContractRegistryVoting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCallerSession) Owner() (common.Address, error) {
	return _GenericContractRegistryVoting.Contract.Owner(&_GenericContractRegistryVoting.CallOpts)
}

// ProposalsCount is a free data retrieval call binding the contract method 0x0a9f46ad.
//
// Solidity: function proposalsCount() view returns(uint256)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCaller) ProposalsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GenericContractRegistryVoting.contract.Call(opts, &out, "proposalsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalsCount is a free data retrieval call binding the contract method 0x0a9f46ad.
//
// Solidity: function proposalsCount() view returns(uint256)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingSession) ProposalsCount() (*big.Int, error) {
	return _GenericContractRegistryVoting.Contract.ProposalsCount(&_GenericContractRegistryVoting.CallOpts)
}

// ProposalsCount is a free data retrieval call binding the contract method 0x0a9f46ad.
//
// Solidity: function proposalsCount() view returns(uint256)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCallerSession) ProposalsCount() (*big.Int, error) {
	return _GenericContractRegistryVoting.Contract.ProposalsCount(&_GenericContractRegistryVoting.CallOpts)
}

// VoteCount is a free data retrieval call binding the contract method 0x4fc8a20d.
//
// Solidity: function voteCount(uint256 ) view returns(uint256)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCaller) VoteCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GenericContractRegistryVoting.contract.Call(opts, &out, "voteCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteCount is a free data retrieval call binding the contract method 0x4fc8a20d.
//
// Solidity: function voteCount(uint256 ) view returns(uint256)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingSession) VoteCount(arg0 *big.Int) (*big.Int, error) {
	return _GenericContractRegistryVoting.Contract.VoteCount(&_GenericContractRegistryVoting.CallOpts, arg0)
}

// VoteCount is a free data retrieval call binding the contract method 0x4fc8a20d.
//
// Solidity: function voteCount(uint256 ) view returns(uint256)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCallerSession) VoteCount(arg0 *big.Int) (*big.Int, error) {
	return _GenericContractRegistryVoting.Contract.VoteCount(&_GenericContractRegistryVoting.CallOpts, arg0)
}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCaller) Voted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _GenericContractRegistryVoting.contract.Call(opts, &out, "voted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingSession) Voted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _GenericContractRegistryVoting.Contract.Voted(&_GenericContractRegistryVoting.CallOpts, arg0, arg1)
}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingCallerSession) Voted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _GenericContractRegistryVoting.Contract.Voted(&_GenericContractRegistryVoting.CallOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(uint256 id_) returns()
func (_GenericContractRegistryVoting *GenericContractRegistryVotingTransactor) Approve(opts *bind.TransactOpts, id_ *big.Int) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.contract.Transact(opts, "approve", id_)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(uint256 id_) returns()
func (_GenericContractRegistryVoting *GenericContractRegistryVotingSession) Approve(id_ *big.Int) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.Approve(&_GenericContractRegistryVoting.TransactOpts, id_)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(uint256 id_) returns()
func (_GenericContractRegistryVoting *GenericContractRegistryVotingTransactorSession) Approve(id_ *big.Int) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.Approve(&_GenericContractRegistryVoting.TransactOpts, id_)
}

// CreateProposal is a paid mutator transaction binding the contract method 0xd182fb9e.
//
// Solidity: function createProposal(string remark_, bytes callData_) returns(uint256)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingTransactor) CreateProposal(opts *bind.TransactOpts, remark_ string, callData_ []byte) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.contract.Transact(opts, "createProposal", remark_, callData_)
}

// CreateProposal is a paid mutator transaction binding the contract method 0xd182fb9e.
//
// Solidity: function createProposal(string remark_, bytes callData_) returns(uint256)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingSession) CreateProposal(remark_ string, callData_ []byte) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.CreateProposal(&_GenericContractRegistryVoting.TransactOpts, remark_, callData_)
}

// CreateProposal is a paid mutator transaction binding the contract method 0xd182fb9e.
//
// Solidity: function createProposal(string remark_, bytes callData_) returns(uint256)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingTransactorSession) CreateProposal(remark_ string, callData_ []byte) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.CreateProposal(&_GenericContractRegistryVoting.TransactOpts, remark_, callData_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registry_) returns()
func (_GenericContractRegistryVoting *GenericContractRegistryVotingTransactor) Initialize(opts *bind.TransactOpts, registry_ common.Address) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.contract.Transact(opts, "initialize", registry_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registry_) returns()
func (_GenericContractRegistryVoting *GenericContractRegistryVotingSession) Initialize(registry_ common.Address) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.Initialize(&_GenericContractRegistryVoting.TransactOpts, registry_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registry_) returns()
func (_GenericContractRegistryVoting *GenericContractRegistryVotingTransactorSession) Initialize(registry_ common.Address) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.Initialize(&_GenericContractRegistryVoting.TransactOpts, registry_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GenericContractRegistryVoting *GenericContractRegistryVotingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GenericContractRegistryVoting *GenericContractRegistryVotingSession) RenounceOwnership() (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.RenounceOwnership(&_GenericContractRegistryVoting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GenericContractRegistryVoting *GenericContractRegistryVotingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.RenounceOwnership(&_GenericContractRegistryVoting.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GenericContractRegistryVoting *GenericContractRegistryVotingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GenericContractRegistryVoting *GenericContractRegistryVotingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.TransferOwnership(&_GenericContractRegistryVoting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GenericContractRegistryVoting *GenericContractRegistryVotingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GenericContractRegistryVoting.Contract.TransferOwnership(&_GenericContractRegistryVoting.TransactOpts, newOwner)
}

// GenericContractRegistryVotingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GenericContractRegistryVoting contract.
type GenericContractRegistryVotingOwnershipTransferredIterator struct {
	Event *GenericContractRegistryVotingOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GenericContractRegistryVotingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericContractRegistryVotingOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GenericContractRegistryVotingOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GenericContractRegistryVotingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericContractRegistryVotingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericContractRegistryVotingOwnershipTransferred represents a OwnershipTransferred event raised by the GenericContractRegistryVoting contract.
type GenericContractRegistryVotingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GenericContractRegistryVotingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GenericContractRegistryVoting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GenericContractRegistryVotingOwnershipTransferredIterator{contract: _GenericContractRegistryVoting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GenericContractRegistryVotingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GenericContractRegistryVoting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericContractRegistryVotingOwnershipTransferred)
				if err := _GenericContractRegistryVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingFilterer) ParseOwnershipTransferred(log types.Log) (*GenericContractRegistryVotingOwnershipTransferred, error) {
	event := new(GenericContractRegistryVotingOwnershipTransferred)
	if err := _GenericContractRegistryVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenericContractRegistryVotingProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the GenericContractRegistryVoting contract.
type GenericContractRegistryVotingProposalCreatedIterator struct {
	Event *GenericContractRegistryVotingProposalCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GenericContractRegistryVotingProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericContractRegistryVotingProposalCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GenericContractRegistryVotingProposalCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GenericContractRegistryVotingProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericContractRegistryVotingProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericContractRegistryVotingProposalCreated represents a ProposalCreated event raised by the GenericContractRegistryVoting contract.
type GenericContractRegistryVotingProposalCreated struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 proposalId)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*GenericContractRegistryVotingProposalCreatedIterator, error) {

	logs, sub, err := _GenericContractRegistryVoting.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &GenericContractRegistryVotingProposalCreatedIterator{contract: _GenericContractRegistryVoting.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 proposalId)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *GenericContractRegistryVotingProposalCreated) (event.Subscription, error) {

	logs, sub, err := _GenericContractRegistryVoting.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericContractRegistryVotingProposalCreated)
				if err := _GenericContractRegistryVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCreated is a log parse operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 proposalId)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingFilterer) ParseProposalCreated(log types.Log) (*GenericContractRegistryVotingProposalCreated, error) {
	event := new(GenericContractRegistryVotingProposalCreated)
	if err := _GenericContractRegistryVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenericContractRegistryVotingProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the GenericContractRegistryVoting contract.
type GenericContractRegistryVotingProposalExecutedIterator struct {
	Event *GenericContractRegistryVotingProposalExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GenericContractRegistryVotingProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericContractRegistryVotingProposalExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GenericContractRegistryVotingProposalExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GenericContractRegistryVotingProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericContractRegistryVotingProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericContractRegistryVotingProposalExecuted represents a ProposalExecuted event raised by the GenericContractRegistryVoting contract.
type GenericContractRegistryVotingProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*GenericContractRegistryVotingProposalExecutedIterator, error) {

	logs, sub, err := _GenericContractRegistryVoting.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &GenericContractRegistryVotingProposalExecutedIterator{contract: _GenericContractRegistryVoting.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GenericContractRegistryVotingProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _GenericContractRegistryVoting.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericContractRegistryVotingProposalExecuted)
				if err := _GenericContractRegistryVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingFilterer) ParseProposalExecuted(log types.Log) (*GenericContractRegistryVotingProposalExecuted, error) {
	event := new(GenericContractRegistryVotingProposalExecuted)
	if err := _GenericContractRegistryVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenericContractRegistryVotingRootNodeApprovedIterator is returned from FilterRootNodeApproved and is used to iterate over the raw logs and unpacked data for RootNodeApproved events raised by the GenericContractRegistryVoting contract.
type GenericContractRegistryVotingRootNodeApprovedIterator struct {
	Event *GenericContractRegistryVotingRootNodeApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GenericContractRegistryVotingRootNodeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericContractRegistryVotingRootNodeApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GenericContractRegistryVotingRootNodeApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GenericContractRegistryVotingRootNodeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericContractRegistryVotingRootNodeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericContractRegistryVotingRootNodeApproved represents a RootNodeApproved event raised by the GenericContractRegistryVoting contract.
type GenericContractRegistryVotingRootNodeApproved struct {
	ProposalId *big.Int
	RootNode   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRootNodeApproved is a free log retrieval operation binding the contract event 0xdcd66ff3278394a103acd0febedb4f0cfae077df25e5b1e05b6b214f3669dd30.
//
// Solidity: event RootNodeApproved(uint256 proposalId, address rootNode)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingFilterer) FilterRootNodeApproved(opts *bind.FilterOpts) (*GenericContractRegistryVotingRootNodeApprovedIterator, error) {

	logs, sub, err := _GenericContractRegistryVoting.contract.FilterLogs(opts, "RootNodeApproved")
	if err != nil {
		return nil, err
	}
	return &GenericContractRegistryVotingRootNodeApprovedIterator{contract: _GenericContractRegistryVoting.contract, event: "RootNodeApproved", logs: logs, sub: sub}, nil
}

// WatchRootNodeApproved is a free log subscription operation binding the contract event 0xdcd66ff3278394a103acd0febedb4f0cfae077df25e5b1e05b6b214f3669dd30.
//
// Solidity: event RootNodeApproved(uint256 proposalId, address rootNode)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingFilterer) WatchRootNodeApproved(opts *bind.WatchOpts, sink chan<- *GenericContractRegistryVotingRootNodeApproved) (event.Subscription, error) {

	logs, sub, err := _GenericContractRegistryVoting.contract.WatchLogs(opts, "RootNodeApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericContractRegistryVotingRootNodeApproved)
				if err := _GenericContractRegistryVoting.contract.UnpackLog(event, "RootNodeApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRootNodeApproved is a log parse operation binding the contract event 0xdcd66ff3278394a103acd0febedb4f0cfae077df25e5b1e05b6b214f3669dd30.
//
// Solidity: event RootNodeApproved(uint256 proposalId, address rootNode)
func (_GenericContractRegistryVoting *GenericContractRegistryVotingFilterer) ParseRootNodeApproved(log types.Log) (*GenericContractRegistryVotingRootNodeApproved, error) {
	event := new(GenericContractRegistryVotingRootNodeApproved)
	if err := _GenericContractRegistryVoting.contract.UnpackLog(event, "RootNodeApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
