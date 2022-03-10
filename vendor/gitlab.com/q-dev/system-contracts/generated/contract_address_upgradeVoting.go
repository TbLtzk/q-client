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

// ContractRegistryAddressVotingSetKeyProposal is an auto generated low-level Go binding around an user-defined struct.
type ContractRegistryAddressVotingSetKeyProposal struct {
	Executed          bool
	VotingStartTime   *big.Int
	VotingExpiredTime *big.Int
	Key               string
	Proxy             common.Address
}

// ContractRegistryAddressVotingMetaData contains all meta data concerning the ContractRegistryAddressVoting contract.
var ContractRegistryAddressVotingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_rootNode\",\"type\":\"address\"}],\"name\":\"RootNodeApproved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getProposalStats\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentMajority\",\"type\":\"uint256\"}],\"internalType\":\"structARootNodeApprovalVoting.ProposalStats\",\"name\":\"_stats\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumARootNodeApprovalVoting.ProposalStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voteCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingExpiredTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"internalType\":\"structContractRegistryAddressVoting.SetKeyProposal\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611733806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a45760003560e01c80630a9f46ad146100a9578063307a064f146100c55780633ed8e9b4146100f35780634fc8a20d146101065780635277b4ae146101265780635c622a0e14610164578063715018a6146101845780638da5cb5b1461018e578063b759f954146101a3578063c4d66de8146101b6578063c7f758a8146101c9578063f2fde38b146101e9575b600080fd5b6100b260035481565b6040519081526020015b60405180910390f35b6100d86100d33660046111a6565b6101fc565b604080518251815260209283015192810192909252016100bc565b6100b26101013660046111d4565b610239565b6100b26101143660046111a6565b60026020526000908152604090205481565b610154610134366004611257565b600160209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016100bc565b6101776101723660046111a6565b6102a7565b6040516100bc9190611287565b61018c6103d0565b005b61019661040b565b6040516100bc91906112af565b61018c6101b13660046111a6565b61041a565b61018c6101c43660046112c3565b6107d9565b6101dc6101d73660046111a6565b610869565b6040516100bc919061133c565b61018c6101f73660046112c3565b6109a1565b604080518082019091526000808252602082015261021982610a41565b602080830191909152600092835260049052604090912060030154815290565b60003361024461040b565b6001600160a01b0316146102735760405162461bcd60e51b815260040161026a90611399565b60405180910390fd5b61029f84848460405160200161028b939291906113ce565b604051602081830303815290604052610af6565b949350505050565b600060035482106102ba57506000919050565b6000828152600460208181526040808420815160a081018352815460ff16151581526001820154938101939093526002810154918301919091526003810154606083015291820180549192916080840191906103159061140f565b80601f01602080910402602001604051908101604052809291908181526020018280546103419061140f565b801561038e5780601f106103635761010080835404028352916020019161038e565b820191906000526020600020905b81548152906001019060200180831161037157829003601f168201915b50505050508152505090508060000151156103ac5750600392915050565b80604001514210156103c15750600192915050565b50600292915050565b50919050565b336103d961040b565b6001600160a01b0316146103ff5760405162461bcd60e51b815260040161026a90611399565b6104096000610cfa565b565b6038546001600160a01b031690565b80600354811061043c5760405162461bcd60e51b815260040161026a90611444565b33610445610d4c565b6001600160a01b031663a230c524826040518263ffffffff1660e01b815260040161047091906112af565b60206040518083038186803b15801561048857600080fd5b505afa15801561049c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104c09190611488565b6105325760405162461bcd60e51b815260206004820152603d60248201527f5b5145432d3033393030305d2d5065726d697373696f6e2064656e696564202d60448201527f206f6e6c7920726f6f74206e6f6465732068617665206163636573732e000000606482015260840161026a565b6000838152600460208181526040808420815160a081018352815460ff161515815260018201549381019390935260028101549183019190915260038101546060830152918201805491929160808401919061058d9061140f565b80601f01602080910402602001604051908101604052809291908181526020018280546105b99061140f565b80156106065780601f106105db57610100808354040283529160200191610606565b820191906000526020600020905b8154815290600101906020018083116105e957829003601f168201915b50505050508152505090508060000151156106635760405162461bcd60e51b815260206004820152601e60248201527f5b5145432d3033393030315d2d416c72656164792065786563757465642e0000604482015260640161026a565b806040015142106106c45760405162461bcd60e51b815260206004820152602560248201527f5b5145432d3033393030325d2d566f74696e672074696d65206861732065787060448201526434b932b21760d91b606482015260840161026a565b6000848152600160209081526040808320338452909152812054819060ff166107635760008681526001602081815260408084203385528252808420805460ff1916909317909255888352600290528120805491610721836114c0565b9091555050604080518781523360208201527fdcd66ff3278394a103acd0febedb4f0cfae077df25e5b1e05b6b214f3669dd30910160405180910390a1600191505b61076c86610df9565b905081806107775750805b6107d15760405162461bcd60e51b815260206004820152602560248201527f5b5145432d3033393030345d2d53656e6465722068617320616c7265616479206044820152641d9bdd195960da1b606482015260840161026a565b505050505050565b600554610100900460ff16806107f2575060055460ff16155b61080e5760405162461bcd60e51b815260040161026a906114db565b600554610100900460ff16158015610830576005805461ffff19166101011790555b600080546001600160a01b0319166001600160a01b038416179055610853610ef9565b8015610865576005805461ff00191690555b5050565b6040805160a08101825260008082526020820181905291810182905260608082015260808101919091528160035481106108b55760405162461bcd60e51b815260040161026a90611444565b600083815260046020818152604092839020805460ff1615158652600181015491860191909152600281015492850192909252810180546108f59061140f565b80601f01602080910402602001604051908101604052809291908181526020018280546109219061140f565b801561096e5780601f106109435761010080835404028352916020019161096e565b820191906000526020600020905b81548152906001019060200180831161095157829003601f168201915b5050505050806020019051810190610986919061154f565b6001600160a01b031660808501526060840152509092915050565b336109aa61040b565b6001600160a01b0316146109d05760405162461bcd60e51b815260040161026a90611399565b6001600160a01b038116610a355760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161026a565b610a3e81610cfa565b50565b600080610a4c610d4c565b6001600160a01b031663de8fa4316040518163ffffffff1660e01b815260040160206040518083038186803b158015610a8457600080fd5b505afa158015610a98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610abc919061160f565b9050806b033b2e3c9fd0803ce8000000600085815260026020526040902054610ae59190611628565b610aef9190611647565b9392505050565b6003805460009182919082610b0a836114c0565b9190505590506000610b1a610f74565b60405162498bff60e81b815260206004820152601f60248201527f636f6e737469747574696f6e2e70726f706f73616c457865637574696f6e500060448201526001600160a01b03919091169063498bff009060640160206040518083038186803b158015610b8857600080fd5b505afa158015610b9c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc0919061160f565b6000838152600460208181526040909220875193945092610be7929184019188019061110d565b50610bf0610f74565b60405162498bff60e81b815260206004820152602260248201527f636f6e737469747574696f6e2e766f74696e672e656d6751557064617465524d60448201526120a560f11b60648201526001600160a01b03919091169063498bff009060840160206040518083038186803b158015610c6957600080fd5b505afa158015610c7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca1919061160f565b60038201554260018201819055610cb9908390611669565b60028201556040518381527fc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe079060200160405180910390a150909392505050565b603880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600080546040805180820182526014815273676f7665726e616e63652e726f6f744e6f64657360601b60208201529051633fb9027160e01b81526001600160a01b0390921691633fb9027191610da491600401611681565b60206040518083038186803b158015610dbc57600080fd5b505afa158015610dd0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df49190611694565b905090565b60008181526004602052604081206003810154610e1584610a41565b1115610eef57610eb0816004018054610e2d9061140f565b80601f0160208091040260200160405190810160405280929190818152602001828054610e599061140f565b8015610ea65780601f10610e7b57610100808354040283529160200191610ea6565b820191906000526020600020905b815481529060010190602001808311610e8957829003601f168201915b5050505050610fbf565b6040518381527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f9060200160405180910390a1805460ff191660011781555b5460ff1692915050565b600554610100900460ff1680610f12575060055460ff16155b610f2e5760405162461bcd60e51b815260040161026a906114db565b600554610100900460ff16158015610f50576005805461ffff19166101011790555b610f58611043565b610f606110ad565b8015610a3e576005805461ff001916905550565b60008054604080516060810190915260228082526001600160a01b0390921691633fb9027191906116dc60208301396040518263ffffffff1660e01b8152600401610da49190611681565b60008082806020019051810190610fd6919061154f565b600054604051639b2ea4bd60e01b81529294509092506001600160a01b031690639b2ea4bd9061100c90859085906004016116b1565b600060405180830381600087803b15801561102657600080fd5b505af115801561103a573d6000803e3d6000fd5b50505050505050565b600554610100900460ff168061105c575060055460ff16155b6110785760405162461bcd60e51b815260040161026a906114db565b600554610100900460ff16158015610f60576005805461ffff19166101011790558015610a3e576005805461ff001916905550565b600554610100900460ff16806110c6575060055460ff16155b6110e25760405162461bcd60e51b815260040161026a906114db565b600554610100900460ff16158015611104576005805461ffff19166101011790555b610f6033610cfa565b8280546111199061140f565b90600052602060002090601f01602090048101928261113b5760008555611181565b82601f1061115457805160ff1916838001178555611181565b82800160010185558215611181579182015b82811115611181578251825591602001919060010190611166565b5061118d929150611191565b5090565b5b8082111561118d5760008155600101611192565b6000602082840312156111b857600080fd5b5035919050565b6001600160a01b0381168114610a3e57600080fd5b6000806000604084860312156111e957600080fd5b833567ffffffffffffffff8082111561120157600080fd5b818601915086601f83011261121557600080fd5b81358181111561122457600080fd5b87602082850101111561123657600080fd5b6020928301955093505084013561124c816111bf565b809150509250925092565b6000806040838503121561126a57600080fd5b82359150602083013561127c816111bf565b809150509250929050565b60208101600483106112a957634e487b7160e01b600052602160045260246000fd5b91905290565b6001600160a01b0391909116815260200190565b6000602082840312156112d557600080fd5b8135610aef816111bf565b60005b838110156112fb5781810151838201526020016112e3565b8381111561130a576000848401525b50505050565b600081518084526113288160208601602086016112e0565b601f01601f19169290920160200192915050565b6020815281511515602082015260208201516040820152604082015160608201526000606083015160a0608084015261137860c0840182611310565b608094909401516001600160a01b031660a093909301929092525090919050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6040815282604082015282846060830137600060608483018101919091526001600160a01b03929092166020820152601f909201601f191690910101919050565b600181811c9082168061142357607f821691505b602082108114156103ca57634e487b7160e01b600052602260045260246000fd5b60208082526024908201527f5b5145432d3033393030355d2d50726f706f73616c206861766e277420696e696040820152633a32b21760e11b606082015260800190565b60006020828403121561149a57600080fd5b81518015158114610aef57600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156114d4576114d46114aa565b5060010190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b805161154a816111bf565b919050565b6000806040838503121561156257600080fd5b825167ffffffffffffffff8082111561157a57600080fd5b818501915085601f83011261158e57600080fd5b8151818111156115a0576115a0611529565b604051601f8201601f19908116603f011681019083821181831017156115c8576115c8611529565b816040528281528860208487010111156115e157600080fd5b6115f28360208301602088016112e0565b80965050505050506116066020840161153f565b90509250929050565b60006020828403121561162157600080fd5b5051919050565b6000816000190483118215151615611642576116426114aa565b500290565b60008261166457634e487b7160e01b600052601260045260246000fd5b500490565b6000821982111561167c5761167c6114aa565b500190565b602081526000610aef6020830184611310565b6000602082840312156116a657600080fd5b8151610aef816111bf565b6040815260006116c46040830185611310565b905060018060a01b0383166020830152939250505056fe676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273a264697066735822122004a34bd56856067b8d194dd81d32c84a925a43133194b002c7b7d8602b7d509164736f6c63430008090033",
}

// ContractRegistryAddressVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractRegistryAddressVotingMetaData.ABI instead.
var ContractRegistryAddressVotingABI = ContractRegistryAddressVotingMetaData.ABI

// ContractRegistryAddressVotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractRegistryAddressVotingMetaData.Bin instead.
var ContractRegistryAddressVotingBin = ContractRegistryAddressVotingMetaData.Bin

// DeployContractRegistryAddressVoting deploys a new Ethereum contract, binding an instance of ContractRegistryAddressVoting to it.
func DeployContractRegistryAddressVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractRegistryAddressVoting, error) {
	parsed, err := ContractRegistryAddressVotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractRegistryAddressVotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractRegistryAddressVoting{ContractRegistryAddressVotingCaller: ContractRegistryAddressVotingCaller{contract: contract}, ContractRegistryAddressVotingTransactor: ContractRegistryAddressVotingTransactor{contract: contract}, ContractRegistryAddressVotingFilterer: ContractRegistryAddressVotingFilterer{contract: contract}}, nil
}

// ContractRegistryAddressVoting is an auto generated Go binding around an Ethereum contract.
type ContractRegistryAddressVoting struct {
	ContractRegistryAddressVotingCaller     // Read-only binding to the contract
	ContractRegistryAddressVotingTransactor // Write-only binding to the contract
	ContractRegistryAddressVotingFilterer   // Log filterer for contract events
}

// ContractRegistryAddressVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractRegistryAddressVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryAddressVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractRegistryAddressVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryAddressVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractRegistryAddressVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryAddressVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractRegistryAddressVotingSession struct {
	Contract     *ContractRegistryAddressVoting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ContractRegistryAddressVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractRegistryAddressVotingCallerSession struct {
	Contract *ContractRegistryAddressVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// ContractRegistryAddressVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractRegistryAddressVotingTransactorSession struct {
	Contract     *ContractRegistryAddressVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// ContractRegistryAddressVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRegistryAddressVotingRaw struct {
	Contract *ContractRegistryAddressVoting // Generic contract binding to access the raw methods on
}

// ContractRegistryAddressVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractRegistryAddressVotingCallerRaw struct {
	Contract *ContractRegistryAddressVotingCaller // Generic read-only contract binding to access the raw methods on
}

// ContractRegistryAddressVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractRegistryAddressVotingTransactorRaw struct {
	Contract *ContractRegistryAddressVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractRegistryAddressVoting creates a new instance of ContractRegistryAddressVoting, bound to a specific deployed contract.
func NewContractRegistryAddressVoting(address common.Address, backend bind.ContractBackend) (*ContractRegistryAddressVoting, error) {
	contract, err := bindContractRegistryAddressVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryAddressVoting{ContractRegistryAddressVotingCaller: ContractRegistryAddressVotingCaller{contract: contract}, ContractRegistryAddressVotingTransactor: ContractRegistryAddressVotingTransactor{contract: contract}, ContractRegistryAddressVotingFilterer: ContractRegistryAddressVotingFilterer{contract: contract}}, nil
}

// NewContractRegistryAddressVotingCaller creates a new read-only instance of ContractRegistryAddressVoting, bound to a specific deployed contract.
func NewContractRegistryAddressVotingCaller(address common.Address, caller bind.ContractCaller) (*ContractRegistryAddressVotingCaller, error) {
	contract, err := bindContractRegistryAddressVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryAddressVotingCaller{contract: contract}, nil
}

// NewContractRegistryAddressVotingTransactor creates a new write-only instance of ContractRegistryAddressVoting, bound to a specific deployed contract.
func NewContractRegistryAddressVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractRegistryAddressVotingTransactor, error) {
	contract, err := bindContractRegistryAddressVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryAddressVotingTransactor{contract: contract}, nil
}

// NewContractRegistryAddressVotingFilterer creates a new log filterer instance of ContractRegistryAddressVoting, bound to a specific deployed contract.
func NewContractRegistryAddressVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractRegistryAddressVotingFilterer, error) {
	contract, err := bindContractRegistryAddressVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryAddressVotingFilterer{contract: contract}, nil
}

// bindContractRegistryAddressVoting binds a generic wrapper to an already deployed contract.
func bindContractRegistryAddressVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractRegistryAddressVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractRegistryAddressVoting.Contract.ContractRegistryAddressVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.ContractRegistryAddressVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.ContractRegistryAddressVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractRegistryAddressVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.contract.Transact(opts, method, params...)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _id) view returns((bool,uint256,uint256,string,address) proposal)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCaller) GetProposal(opts *bind.CallOpts, _id *big.Int) (ContractRegistryAddressVotingSetKeyProposal, error) {
	var out []interface{}
	err := _ContractRegistryAddressVoting.contract.Call(opts, &out, "getProposal", _id)

	if err != nil {
		return *new(ContractRegistryAddressVotingSetKeyProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(ContractRegistryAddressVotingSetKeyProposal)).(*ContractRegistryAddressVotingSetKeyProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _id) view returns((bool,uint256,uint256,string,address) proposal)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingSession) GetProposal(_id *big.Int) (ContractRegistryAddressVotingSetKeyProposal, error) {
	return _ContractRegistryAddressVoting.Contract.GetProposal(&_ContractRegistryAddressVoting.CallOpts, _id)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _id) view returns((bool,uint256,uint256,string,address) proposal)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCallerSession) GetProposal(_id *big.Int) (ContractRegistryAddressVotingSetKeyProposal, error) {
	return _ContractRegistryAddressVoting.Contract.GetProposal(&_ContractRegistryAddressVoting.CallOpts, _id)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _id) view returns((uint256,uint256) _stats)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCaller) GetProposalStats(opts *bind.CallOpts, _id *big.Int) (ARootNodeApprovalVotingProposalStats, error) {
	var out []interface{}
	err := _ContractRegistryAddressVoting.contract.Call(opts, &out, "getProposalStats", _id)

	if err != nil {
		return *new(ARootNodeApprovalVotingProposalStats), err
	}

	out0 := *abi.ConvertType(out[0], new(ARootNodeApprovalVotingProposalStats)).(*ARootNodeApprovalVotingProposalStats)

	return out0, err

}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _id) view returns((uint256,uint256) _stats)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingSession) GetProposalStats(_id *big.Int) (ARootNodeApprovalVotingProposalStats, error) {
	return _ContractRegistryAddressVoting.Contract.GetProposalStats(&_ContractRegistryAddressVoting.CallOpts, _id)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _id) view returns((uint256,uint256) _stats)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCallerSession) GetProposalStats(_id *big.Int) (ARootNodeApprovalVotingProposalStats, error) {
	return _ContractRegistryAddressVoting.Contract.GetProposalStats(&_ContractRegistryAddressVoting.CallOpts, _id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _id) view returns(uint8)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCaller) GetStatus(opts *bind.CallOpts, _id *big.Int) (uint8, error) {
	var out []interface{}
	err := _ContractRegistryAddressVoting.contract.Call(opts, &out, "getStatus", _id)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _id) view returns(uint8)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingSession) GetStatus(_id *big.Int) (uint8, error) {
	return _ContractRegistryAddressVoting.Contract.GetStatus(&_ContractRegistryAddressVoting.CallOpts, _id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _id) view returns(uint8)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCallerSession) GetStatus(_id *big.Int) (uint8, error) {
	return _ContractRegistryAddressVoting.Contract.GetStatus(&_ContractRegistryAddressVoting.CallOpts, _id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryAddressVoting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingSession) Owner() (common.Address, error) {
	return _ContractRegistryAddressVoting.Contract.Owner(&_ContractRegistryAddressVoting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCallerSession) Owner() (common.Address, error) {
	return _ContractRegistryAddressVoting.Contract.Owner(&_ContractRegistryAddressVoting.CallOpts)
}

// ProposalsCount is a free data retrieval call binding the contract method 0x0a9f46ad.
//
// Solidity: function proposalsCount() view returns(uint256)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCaller) ProposalsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryAddressVoting.contract.Call(opts, &out, "proposalsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalsCount is a free data retrieval call binding the contract method 0x0a9f46ad.
//
// Solidity: function proposalsCount() view returns(uint256)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingSession) ProposalsCount() (*big.Int, error) {
	return _ContractRegistryAddressVoting.Contract.ProposalsCount(&_ContractRegistryAddressVoting.CallOpts)
}

// ProposalsCount is a free data retrieval call binding the contract method 0x0a9f46ad.
//
// Solidity: function proposalsCount() view returns(uint256)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCallerSession) ProposalsCount() (*big.Int, error) {
	return _ContractRegistryAddressVoting.Contract.ProposalsCount(&_ContractRegistryAddressVoting.CallOpts)
}

// VoteCount is a free data retrieval call binding the contract method 0x4fc8a20d.
//
// Solidity: function voteCount(uint256 ) view returns(uint256)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCaller) VoteCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryAddressVoting.contract.Call(opts, &out, "voteCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteCount is a free data retrieval call binding the contract method 0x4fc8a20d.
//
// Solidity: function voteCount(uint256 ) view returns(uint256)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingSession) VoteCount(arg0 *big.Int) (*big.Int, error) {
	return _ContractRegistryAddressVoting.Contract.VoteCount(&_ContractRegistryAddressVoting.CallOpts, arg0)
}

// VoteCount is a free data retrieval call binding the contract method 0x4fc8a20d.
//
// Solidity: function voteCount(uint256 ) view returns(uint256)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCallerSession) VoteCount(arg0 *big.Int) (*big.Int, error) {
	return _ContractRegistryAddressVoting.Contract.VoteCount(&_ContractRegistryAddressVoting.CallOpts, arg0)
}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCaller) Voted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractRegistryAddressVoting.contract.Call(opts, &out, "voted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingSession) Voted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ContractRegistryAddressVoting.Contract.Voted(&_ContractRegistryAddressVoting.CallOpts, arg0, arg1)
}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingCallerSession) Voted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ContractRegistryAddressVoting.Contract.Voted(&_ContractRegistryAddressVoting.CallOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(uint256 _id) returns()
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingTransactor) Approve(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.contract.Transact(opts, "approve", _id)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(uint256 _id) returns()
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingSession) Approve(_id *big.Int) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.Approve(&_ContractRegistryAddressVoting.TransactOpts, _id)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(uint256 _id) returns()
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingTransactorSession) Approve(_id *big.Int) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.Approve(&_ContractRegistryAddressVoting.TransactOpts, _id)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x3ed8e9b4.
//
// Solidity: function createProposal(string _key, address _proxy) returns(uint256)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingTransactor) CreateProposal(opts *bind.TransactOpts, _key string, _proxy common.Address) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.contract.Transact(opts, "createProposal", _key, _proxy)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x3ed8e9b4.
//
// Solidity: function createProposal(string _key, address _proxy) returns(uint256)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingSession) CreateProposal(_key string, _proxy common.Address) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.CreateProposal(&_ContractRegistryAddressVoting.TransactOpts, _key, _proxy)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x3ed8e9b4.
//
// Solidity: function createProposal(string _key, address _proxy) returns(uint256)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingTransactorSession) CreateProposal(_key string, _proxy common.Address) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.CreateProposal(&_ContractRegistryAddressVoting.TransactOpts, _key, _proxy)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.Initialize(&_ContractRegistryAddressVoting.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.Initialize(&_ContractRegistryAddressVoting.TransactOpts, _registry)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.RenounceOwnership(&_ContractRegistryAddressVoting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.RenounceOwnership(&_ContractRegistryAddressVoting.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.TransferOwnership(&_ContractRegistryAddressVoting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistryAddressVoting.Contract.TransferOwnership(&_ContractRegistryAddressVoting.TransactOpts, newOwner)
}

// ContractRegistryAddressVotingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractRegistryAddressVoting contract.
type ContractRegistryAddressVotingOwnershipTransferredIterator struct {
	Event *ContractRegistryAddressVotingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractRegistryAddressVotingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryAddressVotingOwnershipTransferred)
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
		it.Event = new(ContractRegistryAddressVotingOwnershipTransferred)
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
func (it *ContractRegistryAddressVotingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryAddressVotingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryAddressVotingOwnershipTransferred represents a OwnershipTransferred event raised by the ContractRegistryAddressVoting contract.
type ContractRegistryAddressVotingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractRegistryAddressVotingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractRegistryAddressVoting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryAddressVotingOwnershipTransferredIterator{contract: _ContractRegistryAddressVoting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractRegistryAddressVotingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractRegistryAddressVoting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryAddressVotingOwnershipTransferred)
				if err := _ContractRegistryAddressVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingFilterer) ParseOwnershipTransferred(log types.Log) (*ContractRegistryAddressVotingOwnershipTransferred, error) {
	event := new(ContractRegistryAddressVotingOwnershipTransferred)
	if err := _ContractRegistryAddressVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryAddressVotingProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the ContractRegistryAddressVoting contract.
type ContractRegistryAddressVotingProposalCreatedIterator struct {
	Event *ContractRegistryAddressVotingProposalCreated // Event containing the contract specifics and raw log

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
func (it *ContractRegistryAddressVotingProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryAddressVotingProposalCreated)
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
		it.Event = new(ContractRegistryAddressVotingProposalCreated)
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
func (it *ContractRegistryAddressVotingProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryAddressVotingProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryAddressVotingProposalCreated represents a ProposalCreated event raised by the ContractRegistryAddressVoting contract.
type ContractRegistryAddressVotingProposalCreated struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 _proposalId)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*ContractRegistryAddressVotingProposalCreatedIterator, error) {

	logs, sub, err := _ContractRegistryAddressVoting.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryAddressVotingProposalCreatedIterator{contract: _ContractRegistryAddressVoting.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 _proposalId)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *ContractRegistryAddressVotingProposalCreated) (event.Subscription, error) {

	logs, sub, err := _ContractRegistryAddressVoting.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryAddressVotingProposalCreated)
				if err := _ContractRegistryAddressVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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
// Solidity: event ProposalCreated(uint256 _proposalId)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingFilterer) ParseProposalCreated(log types.Log) (*ContractRegistryAddressVotingProposalCreated, error) {
	event := new(ContractRegistryAddressVotingProposalCreated)
	if err := _ContractRegistryAddressVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryAddressVotingProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the ContractRegistryAddressVoting contract.
type ContractRegistryAddressVotingProposalExecutedIterator struct {
	Event *ContractRegistryAddressVotingProposalExecuted // Event containing the contract specifics and raw log

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
func (it *ContractRegistryAddressVotingProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryAddressVotingProposalExecuted)
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
		it.Event = new(ContractRegistryAddressVotingProposalExecuted)
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
func (it *ContractRegistryAddressVotingProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryAddressVotingProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryAddressVotingProposalExecuted represents a ProposalExecuted event raised by the ContractRegistryAddressVoting contract.
type ContractRegistryAddressVotingProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 _proposalId)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*ContractRegistryAddressVotingProposalExecutedIterator, error) {

	logs, sub, err := _ContractRegistryAddressVoting.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryAddressVotingProposalExecutedIterator{contract: _ContractRegistryAddressVoting.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 _proposalId)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *ContractRegistryAddressVotingProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _ContractRegistryAddressVoting.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryAddressVotingProposalExecuted)
				if err := _ContractRegistryAddressVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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
// Solidity: event ProposalExecuted(uint256 _proposalId)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingFilterer) ParseProposalExecuted(log types.Log) (*ContractRegistryAddressVotingProposalExecuted, error) {
	event := new(ContractRegistryAddressVotingProposalExecuted)
	if err := _ContractRegistryAddressVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryAddressVotingRootNodeApprovedIterator is returned from FilterRootNodeApproved and is used to iterate over the raw logs and unpacked data for RootNodeApproved events raised by the ContractRegistryAddressVoting contract.
type ContractRegistryAddressVotingRootNodeApprovedIterator struct {
	Event *ContractRegistryAddressVotingRootNodeApproved // Event containing the contract specifics and raw log

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
func (it *ContractRegistryAddressVotingRootNodeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryAddressVotingRootNodeApproved)
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
		it.Event = new(ContractRegistryAddressVotingRootNodeApproved)
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
func (it *ContractRegistryAddressVotingRootNodeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryAddressVotingRootNodeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryAddressVotingRootNodeApproved represents a RootNodeApproved event raised by the ContractRegistryAddressVoting contract.
type ContractRegistryAddressVotingRootNodeApproved struct {
	ProposalId *big.Int
	RootNode   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRootNodeApproved is a free log retrieval operation binding the contract event 0xdcd66ff3278394a103acd0febedb4f0cfae077df25e5b1e05b6b214f3669dd30.
//
// Solidity: event RootNodeApproved(uint256 _proposalId, address _rootNode)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingFilterer) FilterRootNodeApproved(opts *bind.FilterOpts) (*ContractRegistryAddressVotingRootNodeApprovedIterator, error) {

	logs, sub, err := _ContractRegistryAddressVoting.contract.FilterLogs(opts, "RootNodeApproved")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryAddressVotingRootNodeApprovedIterator{contract: _ContractRegistryAddressVoting.contract, event: "RootNodeApproved", logs: logs, sub: sub}, nil
}

// WatchRootNodeApproved is a free log subscription operation binding the contract event 0xdcd66ff3278394a103acd0febedb4f0cfae077df25e5b1e05b6b214f3669dd30.
//
// Solidity: event RootNodeApproved(uint256 _proposalId, address _rootNode)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingFilterer) WatchRootNodeApproved(opts *bind.WatchOpts, sink chan<- *ContractRegistryAddressVotingRootNodeApproved) (event.Subscription, error) {

	logs, sub, err := _ContractRegistryAddressVoting.contract.WatchLogs(opts, "RootNodeApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryAddressVotingRootNodeApproved)
				if err := _ContractRegistryAddressVoting.contract.UnpackLog(event, "RootNodeApproved", log); err != nil {
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
// Solidity: event RootNodeApproved(uint256 _proposalId, address _rootNode)
func (_ContractRegistryAddressVoting *ContractRegistryAddressVotingFilterer) ParseRootNodeApproved(log types.Log) (*ContractRegistryAddressVotingRootNodeApproved, error) {
	event := new(ContractRegistryAddressVotingRootNodeApproved)
	if err := _ContractRegistryAddressVoting.contract.UnpackLog(event, "RootNodeApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
