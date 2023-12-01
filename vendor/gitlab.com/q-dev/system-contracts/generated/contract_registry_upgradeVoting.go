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

// ContractRegistryUpgradeVotingUpgradeProposal is an auto generated low-level Go binding around an user-defined struct.
type ContractRegistryUpgradeVotingUpgradeProposal struct {
	Executed          bool
	VotingStartTime   *big.Int
	VotingExpiredTime *big.Int
	Proxy             common.Address
	Implementation    common.Address
}

// ContractRegistryUpgradeVotingMetaData contains all meta data concerning the ContractRegistryUpgradeVoting contract.
var ContractRegistryUpgradeVotingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_rootNode\",\"type\":\"address\"}],\"name\":\"RootNodeApproved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingExpiredTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"internalType\":\"structContractRegistryUpgradeVoting.UpgradeProposal\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getProposalStats\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentMajority\",\"type\":\"uint256\"}],\"internalType\":\"structARootNodeApprovalVoting.ProposalStats\",\"name\":\"_stats\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumARootNodeApprovalVoting.ProposalStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voteCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061159a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a45760003560e01c80630a9f46ad146100a9578063307a064f146100c55780634fc8a20d146100f35780635277b4ae146101135780635c622a0e14610151578063715018a6146101715780638b34b5341461017b5780638da5cb5b1461018e578063b759f954146101a3578063c4d66de8146101b6578063c7f758a8146101c9578063f2fde38b1461022a575b600080fd5b6100b260035481565b6040519081526020015b60405180910390f35b6100d86100d33660046111e1565b61023d565b604080518251815260209283015192810192909252016100bc565b6100b26101013660046111e1565b60026020526000908152604090205481565b61014161012136600461120f565b600160209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016100bc565b61016461015f3660046111e1565b61027a565b6040516100bc919061123f565b6101796103a3565b005b6100b2610189366004611267565b6103e7565b610196610449565b6040516100bc9190611295565b6101796101b13660046111e1565b610458565b6101796101c43660046112a9565b610817565b6101dc6101d73660046111e1565b6108a7565b6040516100bc919081511515815260208083015190820152604080830151908201526060808301516001600160a01b0390811691830191909152608092830151169181019190915260a00190565b6101796102383660046112a9565b6109e3565b604080518082019091526000808252602082015261025a82610a83565b602080830191909152600092835260049052604090912060030154815290565b6000600354821061028d57506000919050565b6000828152600460208181526040808420815160a081018352815460ff16151581526001820154938101939093526002810154918301919091526003810154606083015291820180549192916080840191906102e8906112c6565b80601f0160208091040260200160405190810160405280929190818152602001828054610314906112c6565b80156103615780601f1061033657610100808354040283529160200191610361565b820191906000526020600020905b81548152906001019060200180831161034457829003601f168201915b505050505081525050905080600001511561037f5750600392915050565b80604001514210156103945750600192915050565b50600292915050565b50919050565b336103ac610449565b6001600160a01b0316146103db5760405162461bcd60e51b81526004016103d2906112fb565b60405180910390fd5b6103e56000610b31565b565b6000336103f2610449565b6001600160a01b0316146104185760405162461bcd60e51b81526004016103d2906112fb565b610442838360405160200161042e929190611330565b604051602081830303815290604052610b83565b9392505050565b6038546001600160a01b031690565b80600354811061047a5760405162461bcd60e51b81526004016103d29061134a565b33610483610d87565b6001600160a01b031663a230c524826040518263ffffffff1660e01b81526004016104ae9190611295565b60206040518083038186803b1580156104c657600080fd5b505afa1580156104da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104fe919061138e565b6105705760405162461bcd60e51b815260206004820152603d60248201527f5b5145432d3033393030305d2d5065726d697373696f6e2064656e696564202d60448201527f206f6e6c7920726f6f74206e6f6465732068617665206163636573732e00000060648201526084016103d2565b6000838152600460208181526040808420815160a081018352815460ff16151581526001820154938101939093526002810154918301919091526003810154606083015291820180549192916080840191906105cb906112c6565b80601f01602080910402602001604051908101604052809291908181526020018280546105f7906112c6565b80156106445780601f1061061957610100808354040283529160200191610644565b820191906000526020600020905b81548152906001019060200180831161062757829003601f168201915b50505050508152505090508060000151156106a15760405162461bcd60e51b815260206004820152601e60248201527f5b5145432d3033393030315d2d416c72656164792065786563757465642e000060448201526064016103d2565b806040015142106107025760405162461bcd60e51b815260206004820152602560248201527f5b5145432d3033393030325d2d566f74696e672074696d65206861732065787060448201526434b932b21760d91b60648201526084016103d2565b6000848152600160209081526040808320338452909152812054819060ff166107a15760008681526001602081815260408084203385528252808420805460ff191690931790925588835260029052812080549161075f836113c6565b9091555050604080518781523360208201527fdcd66ff3278394a103acd0febedb4f0cfae077df25e5b1e05b6b214f3669dd30910160405180910390a1600191505b6107aa86610e34565b905081806107b55750805b61080f5760405162461bcd60e51b815260206004820152602560248201527f5b5145432d3033393030345d2d53656e6465722068617320616c7265616479206044820152641d9bdd195960da1b60648201526084016103d2565b505050505050565b600554610100900460ff1680610830575060055460ff16155b61084c5760405162461bcd60e51b81526004016103d2906113e1565b600554610100900460ff1615801561086e576005805461ffff19166101011790555b600080546001600160a01b0319166001600160a01b038416179055610891610f36565b80156108a3576005805461ff00191690555b5050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091528160035481106108f45760405162461bcd60e51b81526004016103d29061134a565b600083815260046020818152604092839020805460ff161515865260018101549186019190915260028101549285019290925281018054610934906112c6565b80601f0160208091040260200160405190810160405280929190818152602001828054610960906112c6565b80156109ad5780601f10610982576101008083540402835291602001916109ad565b820191906000526020600020905b81548152906001019060200180831161099057829003601f168201915b50505050508060200190518101906109c5919061142f565b6001600160a01b039081166080860152166060840152509092915050565b336109ec610449565b6001600160a01b031614610a125760405162461bcd60e51b81526004016103d2906112fb565b6001600160a01b038116610a775760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103d2565b610a8081610b31565b50565b600080610a8e610d87565b6001600160a01b031663de8fa4316040518163ffffffff1660e01b815260040160206040518083038186803b158015610ac657600080fd5b505afa158015610ada573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610afe919061145e565b9050806b033b2e3c9fd0803ce8000000600085815260026020526040902054610b279190611477565b6104429190611496565b603880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6003805460009182919082610b97836113c6565b9190505590506000610ba7610fb1565b60405162498bff60e81b815260206004820152601f60248201527f636f6e737469747574696f6e2e70726f706f73616c457865637574696f6e500060448201526001600160a01b03919091169063498bff009060640160206040518083038186803b158015610c1557600080fd5b505afa158015610c29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c4d919061145e565b6000838152600460208181526040909220875193945092610c749291840191880190611148565b50610c7d610fb1565b60405162498bff60e81b815260206004820152602260248201527f636f6e737469747574696f6e2e766f74696e672e656d6751557064617465524d60448201526120a560f11b60648201526001600160a01b03919091169063498bff009060840160206040518083038186803b158015610cf657600080fd5b505afa158015610d0a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2e919061145e565b60038201554260018201819055610d469083906114b8565b60028201556040518381527fc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe079060200160405180910390a150909392505050565b600080546040805180820182526014815273676f7665726e616e63652e726f6f744e6f64657360601b60208201529051633fb9027160e01b81526001600160a01b0390921691633fb9027191610ddf916004016114d0565b60206040518083038186803b158015610df757600080fd5b505afa158015610e0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e2f9190611525565b905090565b60008181526004602052604081206003810154610e5084610a83565b1115610f2c57805460ff19166001178155600481018054610ef89190610e75906112c6565b80601f0160208091040260200160405190810160405280929190818152602001828054610ea1906112c6565b8015610eee5780601f10610ec357610100808354040283529160200191610eee565b820191906000526020600020905b815481529060010190602001808311610ed157829003601f168201915b5050505050610ffc565b6040518381527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f9060200160405180910390a15b5460ff1692915050565b600554610100900460ff1680610f4f575060055460ff16155b610f6b5760405162461bcd60e51b81526004016103d2906113e1565b600554610100900460ff16158015610f8d576005805461ffff19166101011790555b610f9561107e565b610f9d6110e8565b8015610a80576005805461ff001916905550565b60008054604080516060810190915260228082526001600160a01b0390921691633fb90271919061154360208301396040518263ffffffff1660e01b8152600401610ddf91906114d0565b60008082806020019051810190611013919061142f565b6000546040516287452360e01b81529294509092506001600160a01b03169062874523906110479085908590600401611330565b600060405180830381600087803b15801561106157600080fd5b505af1158015611075573d6000803e3d6000fd5b50505050505050565b600554610100900460ff1680611097575060055460ff16155b6110b35760405162461bcd60e51b81526004016103d2906113e1565b600554610100900460ff16158015610f9d576005805461ffff19166101011790558015610a80576005805461ff001916905550565b600554610100900460ff1680611101575060055460ff16155b61111d5760405162461bcd60e51b81526004016103d2906113e1565b600554610100900460ff1615801561113f576005805461ffff19166101011790555b610f9d33610b31565b828054611154906112c6565b90600052602060002090601f01602090048101928261117657600085556111bc565b82601f1061118f57805160ff19168380011785556111bc565b828001600101855582156111bc579182015b828111156111bc5782518255916020019190600101906111a1565b506111c89291506111cc565b5090565b5b808211156111c857600081556001016111cd565b6000602082840312156111f357600080fd5b5035919050565b6001600160a01b0381168114610a8057600080fd5b6000806040838503121561122257600080fd5b823591506020830135611234816111fa565b809150509250929050565b602081016004831061126157634e487b7160e01b600052602160045260246000fd5b91905290565b6000806040838503121561127a57600080fd5b8235611285816111fa565b91506020830135611234816111fa565b6001600160a01b0391909116815260200190565b6000602082840312156112bb57600080fd5b8135610442816111fa565b600181811c908216806112da57607f821691505b6020821081141561039d57634e487b7160e01b600052602260045260246000fd5b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6001600160a01b0392831681529116602082015260400190565b60208082526024908201527f5b5145432d3033393030355d2d50726f706f73616c206861766e277420696e696040820152633a32b21760e11b606082015260800190565b6000602082840312156113a057600080fd5b8151801515811461044257600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156113da576113da6113b0565b5060010190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6000806040838503121561144257600080fd5b825161144d816111fa565b6020840151909250611234816111fa565b60006020828403121561147057600080fd5b5051919050565b6000816000190483118215151615611491576114916113b0565b500290565b6000826114b357634e487b7160e01b600052601260045260246000fd5b500490565b600082198211156114cb576114cb6113b0565b500190565b600060208083528351808285015260005b818110156114fd578581018301518582016040015282016114e1565b8181111561150f576000604083870101525b50601f01601f1916929092016040019392505050565b60006020828403121561153757600080fd5b8151610442816111fa56fe676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273a2646970667358221220f7b4b427535b5cfae336f37558a5e137ef6c5298635a2401253e1034b0bc4c4e64736f6c63430008090033",
}

// ContractRegistryUpgradeVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractRegistryUpgradeVotingMetaData.ABI instead.
var ContractRegistryUpgradeVotingABI = ContractRegistryUpgradeVotingMetaData.ABI

// ContractRegistryUpgradeVotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractRegistryUpgradeVotingMetaData.Bin instead.
var ContractRegistryUpgradeVotingBin = ContractRegistryUpgradeVotingMetaData.Bin

// DeployContractRegistryUpgradeVoting deploys a new Ethereum contract, binding an instance of ContractRegistryUpgradeVoting to it.
func DeployContractRegistryUpgradeVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractRegistryUpgradeVoting, error) {
	parsed, err := ContractRegistryUpgradeVotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractRegistryUpgradeVotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractRegistryUpgradeVoting{ContractRegistryUpgradeVotingCaller: ContractRegistryUpgradeVotingCaller{contract: contract}, ContractRegistryUpgradeVotingTransactor: ContractRegistryUpgradeVotingTransactor{contract: contract}, ContractRegistryUpgradeVotingFilterer: ContractRegistryUpgradeVotingFilterer{contract: contract}}, nil
}

// ContractRegistryUpgradeVoting is an auto generated Go binding around an Ethereum contract.
type ContractRegistryUpgradeVoting struct {
	ContractRegistryUpgradeVotingCaller     // Read-only binding to the contract
	ContractRegistryUpgradeVotingTransactor // Write-only binding to the contract
	ContractRegistryUpgradeVotingFilterer   // Log filterer for contract events
}

// ContractRegistryUpgradeVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractRegistryUpgradeVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryUpgradeVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractRegistryUpgradeVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryUpgradeVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractRegistryUpgradeVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryUpgradeVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractRegistryUpgradeVotingSession struct {
	Contract     *ContractRegistryUpgradeVoting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ContractRegistryUpgradeVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractRegistryUpgradeVotingCallerSession struct {
	Contract *ContractRegistryUpgradeVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// ContractRegistryUpgradeVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractRegistryUpgradeVotingTransactorSession struct {
	Contract     *ContractRegistryUpgradeVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// ContractRegistryUpgradeVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRegistryUpgradeVotingRaw struct {
	Contract *ContractRegistryUpgradeVoting // Generic contract binding to access the raw methods on
}

// ContractRegistryUpgradeVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractRegistryUpgradeVotingCallerRaw struct {
	Contract *ContractRegistryUpgradeVotingCaller // Generic read-only contract binding to access the raw methods on
}

// ContractRegistryUpgradeVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractRegistryUpgradeVotingTransactorRaw struct {
	Contract *ContractRegistryUpgradeVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractRegistryUpgradeVoting creates a new instance of ContractRegistryUpgradeVoting, bound to a specific deployed contract.
func NewContractRegistryUpgradeVoting(address common.Address, backend bind.ContractBackend) (*ContractRegistryUpgradeVoting, error) {
	contract, err := bindContractRegistryUpgradeVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryUpgradeVoting{ContractRegistryUpgradeVotingCaller: ContractRegistryUpgradeVotingCaller{contract: contract}, ContractRegistryUpgradeVotingTransactor: ContractRegistryUpgradeVotingTransactor{contract: contract}, ContractRegistryUpgradeVotingFilterer: ContractRegistryUpgradeVotingFilterer{contract: contract}}, nil
}

// NewContractRegistryUpgradeVotingCaller creates a new read-only instance of ContractRegistryUpgradeVoting, bound to a specific deployed contract.
func NewContractRegistryUpgradeVotingCaller(address common.Address, caller bind.ContractCaller) (*ContractRegistryUpgradeVotingCaller, error) {
	contract, err := bindContractRegistryUpgradeVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryUpgradeVotingCaller{contract: contract}, nil
}

// NewContractRegistryUpgradeVotingTransactor creates a new write-only instance of ContractRegistryUpgradeVoting, bound to a specific deployed contract.
func NewContractRegistryUpgradeVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractRegistryUpgradeVotingTransactor, error) {
	contract, err := bindContractRegistryUpgradeVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryUpgradeVotingTransactor{contract: contract}, nil
}

// NewContractRegistryUpgradeVotingFilterer creates a new log filterer instance of ContractRegistryUpgradeVoting, bound to a specific deployed contract.
func NewContractRegistryUpgradeVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractRegistryUpgradeVotingFilterer, error) {
	contract, err := bindContractRegistryUpgradeVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryUpgradeVotingFilterer{contract: contract}, nil
}

// bindContractRegistryUpgradeVoting binds a generic wrapper to an already deployed contract.
func bindContractRegistryUpgradeVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractRegistryUpgradeVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractRegistryUpgradeVoting.Contract.ContractRegistryUpgradeVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.ContractRegistryUpgradeVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.ContractRegistryUpgradeVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractRegistryUpgradeVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.contract.Transact(opts, method, params...)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _id) view returns((bool,uint256,uint256,address,address) proposal)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCaller) GetProposal(opts *bind.CallOpts, _id *big.Int) (ContractRegistryUpgradeVotingUpgradeProposal, error) {
	var out []interface{}
	err := _ContractRegistryUpgradeVoting.contract.Call(opts, &out, "getProposal", _id)

	if err != nil {
		return *new(ContractRegistryUpgradeVotingUpgradeProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(ContractRegistryUpgradeVotingUpgradeProposal)).(*ContractRegistryUpgradeVotingUpgradeProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _id) view returns((bool,uint256,uint256,address,address) proposal)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingSession) GetProposal(_id *big.Int) (ContractRegistryUpgradeVotingUpgradeProposal, error) {
	return _ContractRegistryUpgradeVoting.Contract.GetProposal(&_ContractRegistryUpgradeVoting.CallOpts, _id)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _id) view returns((bool,uint256,uint256,address,address) proposal)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCallerSession) GetProposal(_id *big.Int) (ContractRegistryUpgradeVotingUpgradeProposal, error) {
	return _ContractRegistryUpgradeVoting.Contract.GetProposal(&_ContractRegistryUpgradeVoting.CallOpts, _id)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _id) view returns((uint256,uint256) _stats)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCaller) GetProposalStats(opts *bind.CallOpts, _id *big.Int) (ARootNodeApprovalVotingProposalStats, error) {
	var out []interface{}
	err := _ContractRegistryUpgradeVoting.contract.Call(opts, &out, "getProposalStats", _id)

	if err != nil {
		return *new(ARootNodeApprovalVotingProposalStats), err
	}

	out0 := *abi.ConvertType(out[0], new(ARootNodeApprovalVotingProposalStats)).(*ARootNodeApprovalVotingProposalStats)

	return out0, err

}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _id) view returns((uint256,uint256) _stats)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingSession) GetProposalStats(_id *big.Int) (ARootNodeApprovalVotingProposalStats, error) {
	return _ContractRegistryUpgradeVoting.Contract.GetProposalStats(&_ContractRegistryUpgradeVoting.CallOpts, _id)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _id) view returns((uint256,uint256) _stats)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCallerSession) GetProposalStats(_id *big.Int) (ARootNodeApprovalVotingProposalStats, error) {
	return _ContractRegistryUpgradeVoting.Contract.GetProposalStats(&_ContractRegistryUpgradeVoting.CallOpts, _id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _id) view returns(uint8)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCaller) GetStatus(opts *bind.CallOpts, _id *big.Int) (uint8, error) {
	var out []interface{}
	err := _ContractRegistryUpgradeVoting.contract.Call(opts, &out, "getStatus", _id)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _id) view returns(uint8)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingSession) GetStatus(_id *big.Int) (uint8, error) {
	return _ContractRegistryUpgradeVoting.Contract.GetStatus(&_ContractRegistryUpgradeVoting.CallOpts, _id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _id) view returns(uint8)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCallerSession) GetStatus(_id *big.Int) (uint8, error) {
	return _ContractRegistryUpgradeVoting.Contract.GetStatus(&_ContractRegistryUpgradeVoting.CallOpts, _id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryUpgradeVoting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingSession) Owner() (common.Address, error) {
	return _ContractRegistryUpgradeVoting.Contract.Owner(&_ContractRegistryUpgradeVoting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCallerSession) Owner() (common.Address, error) {
	return _ContractRegistryUpgradeVoting.Contract.Owner(&_ContractRegistryUpgradeVoting.CallOpts)
}

// ProposalsCount is a free data retrieval call binding the contract method 0x0a9f46ad.
//
// Solidity: function proposalsCount() view returns(uint256)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCaller) ProposalsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryUpgradeVoting.contract.Call(opts, &out, "proposalsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalsCount is a free data retrieval call binding the contract method 0x0a9f46ad.
//
// Solidity: function proposalsCount() view returns(uint256)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingSession) ProposalsCount() (*big.Int, error) {
	return _ContractRegistryUpgradeVoting.Contract.ProposalsCount(&_ContractRegistryUpgradeVoting.CallOpts)
}

// ProposalsCount is a free data retrieval call binding the contract method 0x0a9f46ad.
//
// Solidity: function proposalsCount() view returns(uint256)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCallerSession) ProposalsCount() (*big.Int, error) {
	return _ContractRegistryUpgradeVoting.Contract.ProposalsCount(&_ContractRegistryUpgradeVoting.CallOpts)
}

// VoteCount is a free data retrieval call binding the contract method 0x4fc8a20d.
//
// Solidity: function voteCount(uint256 ) view returns(uint256)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCaller) VoteCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryUpgradeVoting.contract.Call(opts, &out, "voteCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteCount is a free data retrieval call binding the contract method 0x4fc8a20d.
//
// Solidity: function voteCount(uint256 ) view returns(uint256)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingSession) VoteCount(arg0 *big.Int) (*big.Int, error) {
	return _ContractRegistryUpgradeVoting.Contract.VoteCount(&_ContractRegistryUpgradeVoting.CallOpts, arg0)
}

// VoteCount is a free data retrieval call binding the contract method 0x4fc8a20d.
//
// Solidity: function voteCount(uint256 ) view returns(uint256)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCallerSession) VoteCount(arg0 *big.Int) (*big.Int, error) {
	return _ContractRegistryUpgradeVoting.Contract.VoteCount(&_ContractRegistryUpgradeVoting.CallOpts, arg0)
}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCaller) Voted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractRegistryUpgradeVoting.contract.Call(opts, &out, "voted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingSession) Voted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ContractRegistryUpgradeVoting.Contract.Voted(&_ContractRegistryUpgradeVoting.CallOpts, arg0, arg1)
}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingCallerSession) Voted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ContractRegistryUpgradeVoting.Contract.Voted(&_ContractRegistryUpgradeVoting.CallOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(uint256 _id) returns()
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingTransactor) Approve(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.contract.Transact(opts, "approve", _id)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(uint256 _id) returns()
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingSession) Approve(_id *big.Int) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.Approve(&_ContractRegistryUpgradeVoting.TransactOpts, _id)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(uint256 _id) returns()
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingTransactorSession) Approve(_id *big.Int) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.Approve(&_ContractRegistryUpgradeVoting.TransactOpts, _id)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x8b34b534.
//
// Solidity: function createProposal(address _proxy, address _implementation) returns(uint256)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingTransactor) CreateProposal(opts *bind.TransactOpts, _proxy common.Address, _implementation common.Address) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.contract.Transact(opts, "createProposal", _proxy, _implementation)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x8b34b534.
//
// Solidity: function createProposal(address _proxy, address _implementation) returns(uint256)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingSession) CreateProposal(_proxy common.Address, _implementation common.Address) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.CreateProposal(&_ContractRegistryUpgradeVoting.TransactOpts, _proxy, _implementation)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x8b34b534.
//
// Solidity: function createProposal(address _proxy, address _implementation) returns(uint256)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingTransactorSession) CreateProposal(_proxy common.Address, _implementation common.Address) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.CreateProposal(&_ContractRegistryUpgradeVoting.TransactOpts, _proxy, _implementation)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.Initialize(&_ContractRegistryUpgradeVoting.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.Initialize(&_ContractRegistryUpgradeVoting.TransactOpts, _registry)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.RenounceOwnership(&_ContractRegistryUpgradeVoting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.RenounceOwnership(&_ContractRegistryUpgradeVoting.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.TransferOwnership(&_ContractRegistryUpgradeVoting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistryUpgradeVoting.Contract.TransferOwnership(&_ContractRegistryUpgradeVoting.TransactOpts, newOwner)
}

// ContractRegistryUpgradeVotingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractRegistryUpgradeVoting contract.
type ContractRegistryUpgradeVotingOwnershipTransferredIterator struct {
	Event *ContractRegistryUpgradeVotingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractRegistryUpgradeVotingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryUpgradeVotingOwnershipTransferred)
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
		it.Event = new(ContractRegistryUpgradeVotingOwnershipTransferred)
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
func (it *ContractRegistryUpgradeVotingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryUpgradeVotingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryUpgradeVotingOwnershipTransferred represents a OwnershipTransferred event raised by the ContractRegistryUpgradeVoting contract.
type ContractRegistryUpgradeVotingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractRegistryUpgradeVotingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractRegistryUpgradeVoting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryUpgradeVotingOwnershipTransferredIterator{contract: _ContractRegistryUpgradeVoting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractRegistryUpgradeVotingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractRegistryUpgradeVoting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryUpgradeVotingOwnershipTransferred)
				if err := _ContractRegistryUpgradeVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingFilterer) ParseOwnershipTransferred(log types.Log) (*ContractRegistryUpgradeVotingOwnershipTransferred, error) {
	event := new(ContractRegistryUpgradeVotingOwnershipTransferred)
	if err := _ContractRegistryUpgradeVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryUpgradeVotingProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the ContractRegistryUpgradeVoting contract.
type ContractRegistryUpgradeVotingProposalCreatedIterator struct {
	Event *ContractRegistryUpgradeVotingProposalCreated // Event containing the contract specifics and raw log

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
func (it *ContractRegistryUpgradeVotingProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryUpgradeVotingProposalCreated)
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
		it.Event = new(ContractRegistryUpgradeVotingProposalCreated)
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
func (it *ContractRegistryUpgradeVotingProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryUpgradeVotingProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryUpgradeVotingProposalCreated represents a ProposalCreated event raised by the ContractRegistryUpgradeVoting contract.
type ContractRegistryUpgradeVotingProposalCreated struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 _proposalId)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*ContractRegistryUpgradeVotingProposalCreatedIterator, error) {

	logs, sub, err := _ContractRegistryUpgradeVoting.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryUpgradeVotingProposalCreatedIterator{contract: _ContractRegistryUpgradeVoting.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 _proposalId)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *ContractRegistryUpgradeVotingProposalCreated) (event.Subscription, error) {

	logs, sub, err := _ContractRegistryUpgradeVoting.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryUpgradeVotingProposalCreated)
				if err := _ContractRegistryUpgradeVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingFilterer) ParseProposalCreated(log types.Log) (*ContractRegistryUpgradeVotingProposalCreated, error) {
	event := new(ContractRegistryUpgradeVotingProposalCreated)
	if err := _ContractRegistryUpgradeVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryUpgradeVotingProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the ContractRegistryUpgradeVoting contract.
type ContractRegistryUpgradeVotingProposalExecutedIterator struct {
	Event *ContractRegistryUpgradeVotingProposalExecuted // Event containing the contract specifics and raw log

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
func (it *ContractRegistryUpgradeVotingProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryUpgradeVotingProposalExecuted)
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
		it.Event = new(ContractRegistryUpgradeVotingProposalExecuted)
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
func (it *ContractRegistryUpgradeVotingProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryUpgradeVotingProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryUpgradeVotingProposalExecuted represents a ProposalExecuted event raised by the ContractRegistryUpgradeVoting contract.
type ContractRegistryUpgradeVotingProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 _proposalId)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*ContractRegistryUpgradeVotingProposalExecutedIterator, error) {

	logs, sub, err := _ContractRegistryUpgradeVoting.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryUpgradeVotingProposalExecutedIterator{contract: _ContractRegistryUpgradeVoting.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 _proposalId)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *ContractRegistryUpgradeVotingProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _ContractRegistryUpgradeVoting.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryUpgradeVotingProposalExecuted)
				if err := _ContractRegistryUpgradeVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingFilterer) ParseProposalExecuted(log types.Log) (*ContractRegistryUpgradeVotingProposalExecuted, error) {
	event := new(ContractRegistryUpgradeVotingProposalExecuted)
	if err := _ContractRegistryUpgradeVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryUpgradeVotingRootNodeApprovedIterator is returned from FilterRootNodeApproved and is used to iterate over the raw logs and unpacked data for RootNodeApproved events raised by the ContractRegistryUpgradeVoting contract.
type ContractRegistryUpgradeVotingRootNodeApprovedIterator struct {
	Event *ContractRegistryUpgradeVotingRootNodeApproved // Event containing the contract specifics and raw log

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
func (it *ContractRegistryUpgradeVotingRootNodeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryUpgradeVotingRootNodeApproved)
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
		it.Event = new(ContractRegistryUpgradeVotingRootNodeApproved)
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
func (it *ContractRegistryUpgradeVotingRootNodeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryUpgradeVotingRootNodeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryUpgradeVotingRootNodeApproved represents a RootNodeApproved event raised by the ContractRegistryUpgradeVoting contract.
type ContractRegistryUpgradeVotingRootNodeApproved struct {
	ProposalId *big.Int
	RootNode   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRootNodeApproved is a free log retrieval operation binding the contract event 0xdcd66ff3278394a103acd0febedb4f0cfae077df25e5b1e05b6b214f3669dd30.
//
// Solidity: event RootNodeApproved(uint256 _proposalId, address _rootNode)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingFilterer) FilterRootNodeApproved(opts *bind.FilterOpts) (*ContractRegistryUpgradeVotingRootNodeApprovedIterator, error) {

	logs, sub, err := _ContractRegistryUpgradeVoting.contract.FilterLogs(opts, "RootNodeApproved")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryUpgradeVotingRootNodeApprovedIterator{contract: _ContractRegistryUpgradeVoting.contract, event: "RootNodeApproved", logs: logs, sub: sub}, nil
}

// WatchRootNodeApproved is a free log subscription operation binding the contract event 0xdcd66ff3278394a103acd0febedb4f0cfae077df25e5b1e05b6b214f3669dd30.
//
// Solidity: event RootNodeApproved(uint256 _proposalId, address _rootNode)
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingFilterer) WatchRootNodeApproved(opts *bind.WatchOpts, sink chan<- *ContractRegistryUpgradeVotingRootNodeApproved) (event.Subscription, error) {

	logs, sub, err := _ContractRegistryUpgradeVoting.contract.WatchLogs(opts, "RootNodeApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryUpgradeVotingRootNodeApproved)
				if err := _ContractRegistryUpgradeVoting.contract.UnpackLog(event, "RootNodeApproved", log); err != nil {
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
func (_ContractRegistryUpgradeVoting *ContractRegistryUpgradeVotingFilterer) ParseRootNodeApproved(log types.Log) (*ContractRegistryUpgradeVotingRootNodeApproved, error) {
	event := new(ContractRegistryUpgradeVotingRootNodeApproved)
	if err := _ContractRegistryUpgradeVoting.contract.UnpackLog(event, "RootNodeApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
