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

// TimeLockEntry is an auto generated low-level Go binding around an user-defined struct.
type TimeLockEntry struct {
	Amount       *big.Int
	ReleaseStart *big.Int
	ReleaseEnd   *big.Int
}

// VestingMetaData contains all meta data concerning the Vesting contract.
var VestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseEnd\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTimeLockEntry\",\"name\":\"_timelock\",\"type\":\"tuple\"}],\"name\":\"NewTimeLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseEnd\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTimeLockEntry\",\"name\":\"_timelock\",\"type\":\"tuple\"}],\"name\":\"Purged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_releaseStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_releaseEnd\",\"type\":\"uint256\"}],\"name\":\"depositOnBehalfOf\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinimumBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getTimeLocks\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseEnd\",\"type\":\"uint256\"}],\"internalType\":\"structTimeLockEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"purgeTimeLocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"announceUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLockInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingUnlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingUnlockTime\",\"type\":\"uint256\"}],\"internalType\":\"structVotingLockInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"withdrawableBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061135c806100206000396000f3fe6080604052600436106100975760003560e01c80632e1a7d4d1461009c5780633aa6c4f9146100be5780636198e339146100f157806370a082311461011157806373983f4314610131578063a1c9e7c614610144578063afb60cff14610164578063b0426a3f14610184578063c4d66de8146101cc578063dd467064146101ec578063e7522fc01461020c578063f0e810241461022c575b600080fd5b3480156100a857600080fd5b506100bc6100b7366004611041565b610259565b005b3480156100ca57600080fd5b506100de6100d936600461106f565b610328565b6040519081526020015b60405180910390f35b3480156100fd57600080fd5b506100bc61010c366004611041565b610377565b34801561011d57600080fd5b506100de61012c36600461106f565b610384565b6100bc61013f36600461108c565b61039f565b34801561015057600080fd5b506100bc61015f366004611041565b6105e1565b34801561017057600080fd5b506100bc61017f36600461106f565b610659565b34801561019057600080fd5b50610199610662565b6040516100e891908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b3480156101d857600080fd5b506100bc6101e736600461106f565b610677565b3480156101f857600080fd5b506100bc610207366004611041565b610746565b34801561021857600080fd5b506100de6102273660046110c1565b610750565b34801561023857600080fd5b5061024c61024736600461106f565b610830565b6040516100e891906110ed565b3361026481836108c3565b506000816001600160a01b03168360405160006040518083038185875af1925050503d80600081146102b2576040519150601f19603f3d011682016040523d82523d6000602084013e6102b7565b606091505b50509050806103235760405162461bcd60e51b815260206004820152602d60248201527f5b5145432d3033313030335d2d4661696c656420746f2077697468647261772060448201526c333937b6903b32b9ba34b7339760991b60648201526084015b60405180910390fd5b505050565b60008061033483610a6b565b905060006103428442610750565b9050808260000151111561036d57815161035b85610384565b6103659190611165565b949350505050565b8061035b85610384565b6103813382610b06565b50565b6001600160a01b031660009081526002602052604090205490565b8082106104145760405162461bcd60e51b815260206004820152603f60248201527f5b5145432d3033323030315d2d54696d65206c6f636b2072656c65617365206560448201527f6e64206d7573742062652061667465722072656c656173652073746172742e00606482015260840161031a565b678ac7230489e800003410156104885760405162461bcd60e51b815260206004820152603360248201527f5b5145432d3033323030325d2d476976656e20616d6f756e742069732062656c60448201527237bb903232b837b9b4ba1036b4b734b6bab69760691b606482015260840161031a565b61049183610b78565b6001600160a01b038316600090815260016020526040902054600a1161051c5760405162461bcd60e51b815260206004820152603a60248201527f5b5145432d3033323030335d2d4661696c656420746f206465706f736974206160448201527936b7bab73a16103a37b79036b0b73c903a34b6b2b637b1b5b99760311b606482015260840161031a565b428111156105d75761054860405180606001604052806000815260200160008152602001600081525090565b348152602080820184815260408084018581526001600160a01b03881660008181526001808752848220805480830182559083529690912087516003909702019586559351938501939093555160029093019290925590517f42dc784443c8161a28f39bb05a48bc8bac980ab9115d5b9cfef307f00204bd7f906105cd90849061117c565b60405180910390a2505b6103238334610d36565b60006105eb610d6e565b604051633a5f594760e21b815233600482015260248101849052600060448201529091506001600160a01b0382169063e97d651c90606401600060405180830381600087803b15801561063d57600080fd5b505af1158015610651573d6000803e3d6000fd5b505050505050565b61038181610b78565b61066a610fd7565b610672610e20565b905090565b600054610100900460ff1680610690575060005460ff16155b6106f35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161031a565b600054610100900460ff16158015610715576000805461ffff19166101011790555b600380546001600160a01b0319166001600160a01b0384161790558015610742576000805461ff00191690555b5050565b6103813382610e31565b6001600160a01b0382166000908152600160205260408120546107755750600061082a565b6000805b6001600160a01b038516600090815260016020526040902054811015610826576001600160a01b03851660009081526001602052604081208054839081106107c3576107c361119d565b9060005260206000209060030201604051806060016040529081600082015481526020016001820154815260200160028201548152505090506108068186610f5e565b61081090846111b3565b925050808061081e906111cb565b915050610779565b5090505b92915050565b6001600160a01b0381166000908152600160209081526040808320805482518185028101850190935280835260609492939192909184015b828210156108b85783829060005260206000209060030201604051806060016040529081600082015481526020016001820154815260200160028201548152505081526020019060010190610868565b505050509050919050565b6000816108d25750600061082a565b816108dc84610384565b101561094c5760405162461bcd60e51b815260206004820152603960248201527f5b5145432d3033313030315d2d496e73756666696369656e742062616c616e6360448201527832903337b9103b32b9ba34b733903bb4ba34323930bbb0b61760391b606482015260840161031a565b6109568342610750565b8261096085610384565b61096a9190611165565b10156109da5760405162461bcd60e51b815260206004820152603960248201527f5b5145432d3033313030325d2d56657374696e67207769746864726177616c20604482015278383932bb32b73a32b210313c903a34b6b2b637b1b59439949760391b606482015260840161031a565b60006109e584610a6b565b90506000816040015182600001516109fd91906111b3565b9050600081610a0b87610384565b610a159190611165565b905084811015610a3257610a3286610a2d8388611165565b610b06565b6001600160a01b03861660009081526002602052604081208054879290610a5a908490611165565b909155506001979650505050505050565b610a73610fd7565b6000610a7d610d6e565b6040516218724160e61b81523060048201526001600160a01b0385811660248301529192509082169063061c90409060440160806040518083038186803b158015610ac757600080fd5b505afa158015610adb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aff91906111e6565b9392505050565b6000610b10610d6e565b604051637eee288d60e01b81529091506001600160a01b03821690637eee288d90610b41908690869060040161125a565b600060405180830381600087803b158015610b5b57600080fd5b505af1158015610b6f573d6000803e3d6000fd5b50505050505050565b6001600160a01b038116600090815260016020908152604080832080548251818502810185019093528083529192909190849084015b82821015610bfe5783829060005260206000209060030201604051806060016040529081600082015481526020016001820154815260200160028201548152505081526020019060010190610bae565b5050506001600160a01b0384166000908152600160205260408120929350610c27929150610fff565b60005b815181101561032357818181518110610c4557610c4561119d565b602002602001015160400151421015610cc9576001600160a01b03831660009081526001602052604090208251839083908110610c8457610c8461119d565b60209081029190910181015182546001818101855560009485529383902082516003909202019081559181015192820192909255604090910151600290910155610d24565b826001600160a01b03167f646189d6820ad36442dfb379131756840be21178303b28f1f607ca4e45762e93838381518110610d0657610d0661119d565b6020026020010151604051610d1b919061117c565b60405180910390a25b80610d2e816111cb565b915050610c2a565b6001600160a01b03821660009081526002602052604081208054839290610d5e9084906111b3565b9091555061074290508282610e31565b600354604080518082018252601c81527f676f7665726e616e63652e766f74696e6757656967687450726f78790000000060208201529051633fb9027160e01b81526000926001600160a01b031691633fb9027191610dd09190600401611273565b60206040518083038186803b158015610de857600080fd5b505afa158015610dfc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061067291906112c8565b610e28610fd7565b61067233610a6b565b6000610e3b610d6e565b90506000610e4884610a6b565b9050600081604001518260000151610e6091906111b3565b90508381610e6d87610384565b610e779190611165565b1015610ef75760405162461bcd60e51b815260206004820152604360248201527f5b5145432d3033313030345d2d546865206c6f636b20616d6f756e74206d757360448201527f74206e6f74206578636565642074686520617661696c61626c652062616c616e60648201526231b29760e91b608482015260a40161031a565b60405163282d3fdf60e01b81526001600160a01b0384169063282d3fdf90610f25908890889060040161125a565b600060405180830381600087803b158015610f3f57600080fd5b505af1158015610f53573d6000803e3d6000fd5b505050505050505050565b60008260200151821015610f745750815161082a565b82604001518210610f875750600061082a565b6000828460400151610f999190611165565b9050600084602001518560400151610fb19190611165565b905080828660000151610fc491906112e5565b610fce9190611304565b95945050505050565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b508054600082556003029060005260206000209081019061038191905b8082111561103d57600080825560018201819055600282015560030161101c565b5090565b60006020828403121561105357600080fd5b5035919050565b6001600160a01b038116811461038157600080fd5b60006020828403121561108157600080fd5b8135610aff8161105a565b6000806000606084860312156110a157600080fd5b83356110ac8161105a565b95602085013595506040909401359392505050565b600080604083850312156110d457600080fd5b82356110df8161105a565b946020939093013593505050565b6020808252825182820181905260009190848201906040850190845b81811015611143576111308385518051825260208082015190830152604090810151910152565b9284019260609290920191600101611109565b50909695505050505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156111775761117761114f565b500390565b8151815260208083015190820152604080830151908201526060810161082a565b634e487b7160e01b600052603260045260246000fd5b600082198211156111c6576111c661114f565b500190565b60006000198214156111df576111df61114f565b5060010190565b6000608082840312156111f857600080fd5b6040516080810181811067ffffffffffffffff8211171561122957634e487b7160e01b600052604160045260246000fd5b8060405250825181526020830151602082015260408301516040820152606083015160608201528091505092915050565b6001600160a01b03929092168252602082015260400190565b600060208083528351808285015260005b818110156112a057858101830151858201604001528201611284565b818111156112b2576000604083870101525b50601f01601f1916929092016040019392505050565b6000602082840312156112da57600080fd5b8151610aff8161105a565b60008160001904831182151516156112ff576112ff61114f565b500290565b60008261132157634e487b7160e01b600052601260045260246000fd5b50049056fea26469706673582212208266a334831257af3b69e752d9a5e7190f10c5fe623029fd30dc7cba72fb2d3764736f6c63430008090033",
}

// VestingABI is the input ABI used to generate the binding from.
// Deprecated: Use VestingMetaData.ABI instead.
var VestingABI = VestingMetaData.ABI

// VestingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VestingMetaData.Bin instead.
var VestingBin = VestingMetaData.Bin

// DeployVesting deploys a new Ethereum contract, binding an instance of Vesting to it.
func DeployVesting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vesting, error) {
	parsed, err := VestingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VestingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vesting{VestingCaller: VestingCaller{contract: contract}, VestingTransactor: VestingTransactor{contract: contract}, VestingFilterer: VestingFilterer{contract: contract}}, nil
}

// Vesting is an auto generated Go binding around an Ethereum contract.
type Vesting struct {
	VestingCaller     // Read-only binding to the contract
	VestingTransactor // Write-only binding to the contract
	VestingFilterer   // Log filterer for contract events
}

// VestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VestingSession struct {
	Contract     *Vesting          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VestingCallerSession struct {
	Contract *VestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VestingTransactorSession struct {
	Contract     *VestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VestingRaw struct {
	Contract *Vesting // Generic contract binding to access the raw methods on
}

// VestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VestingCallerRaw struct {
	Contract *VestingCaller // Generic read-only contract binding to access the raw methods on
}

// VestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VestingTransactorRaw struct {
	Contract *VestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVesting creates a new instance of Vesting, bound to a specific deployed contract.
func NewVesting(address common.Address, backend bind.ContractBackend) (*Vesting, error) {
	contract, err := bindVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vesting{VestingCaller: VestingCaller{contract: contract}, VestingTransactor: VestingTransactor{contract: contract}, VestingFilterer: VestingFilterer{contract: contract}}, nil
}

// NewVestingCaller creates a new read-only instance of Vesting, bound to a specific deployed contract.
func NewVestingCaller(address common.Address, caller bind.ContractCaller) (*VestingCaller, error) {
	contract, err := bindVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VestingCaller{contract: contract}, nil
}

// NewVestingTransactor creates a new write-only instance of Vesting, bound to a specific deployed contract.
func NewVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*VestingTransactor, error) {
	contract, err := bindVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VestingTransactor{contract: contract}, nil
}

// NewVestingFilterer creates a new log filterer instance of Vesting, bound to a specific deployed contract.
func NewVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*VestingFilterer, error) {
	contract, err := bindVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VestingFilterer{contract: contract}, nil
}

// bindVesting binds a generic wrapper to an already deployed contract.
func bindVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.VestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Vesting *VestingCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Vesting *VestingSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Vesting.Contract.BalanceOf(&_Vesting.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Vesting *VestingCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Vesting.Contract.BalanceOf(&_Vesting.CallOpts, _account)
}

// GetLockInfo is a free data retrieval call binding the contract method 0xb0426a3f.
//
// Solidity: function getLockInfo() view returns((uint256,uint256,uint256,uint256))
func (_Vesting *VestingCaller) GetLockInfo(opts *bind.CallOpts) (VotingLockInfo, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "getLockInfo")

	if err != nil {
		return *new(VotingLockInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(VotingLockInfo)).(*VotingLockInfo)

	return out0, err

}

// GetLockInfo is a free data retrieval call binding the contract method 0xb0426a3f.
//
// Solidity: function getLockInfo() view returns((uint256,uint256,uint256,uint256))
func (_Vesting *VestingSession) GetLockInfo() (VotingLockInfo, error) {
	return _Vesting.Contract.GetLockInfo(&_Vesting.CallOpts)
}

// GetLockInfo is a free data retrieval call binding the contract method 0xb0426a3f.
//
// Solidity: function getLockInfo() view returns((uint256,uint256,uint256,uint256))
func (_Vesting *VestingCallerSession) GetLockInfo() (VotingLockInfo, error) {
	return _Vesting.Contract.GetLockInfo(&_Vesting.CallOpts)
}

// GetMinimumBalance is a free data retrieval call binding the contract method 0xe7522fc0.
//
// Solidity: function getMinimumBalance(address _account, uint256 _timestamp) view returns(uint256)
func (_Vesting *VestingCaller) GetMinimumBalance(opts *bind.CallOpts, _account common.Address, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "getMinimumBalance", _account, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumBalance is a free data retrieval call binding the contract method 0xe7522fc0.
//
// Solidity: function getMinimumBalance(address _account, uint256 _timestamp) view returns(uint256)
func (_Vesting *VestingSession) GetMinimumBalance(_account common.Address, _timestamp *big.Int) (*big.Int, error) {
	return _Vesting.Contract.GetMinimumBalance(&_Vesting.CallOpts, _account, _timestamp)
}

// GetMinimumBalance is a free data retrieval call binding the contract method 0xe7522fc0.
//
// Solidity: function getMinimumBalance(address _account, uint256 _timestamp) view returns(uint256)
func (_Vesting *VestingCallerSession) GetMinimumBalance(_account common.Address, _timestamp *big.Int) (*big.Int, error) {
	return _Vesting.Contract.GetMinimumBalance(&_Vesting.CallOpts, _account, _timestamp)
}

// GetTimeLocks is a free data retrieval call binding the contract method 0xf0e81024.
//
// Solidity: function getTimeLocks(address _account) view returns((uint256,uint256,uint256)[])
func (_Vesting *VestingCaller) GetTimeLocks(opts *bind.CallOpts, _account common.Address) ([]TimeLockEntry, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "getTimeLocks", _account)

	if err != nil {
		return *new([]TimeLockEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]TimeLockEntry)).(*[]TimeLockEntry)

	return out0, err

}

// GetTimeLocks is a free data retrieval call binding the contract method 0xf0e81024.
//
// Solidity: function getTimeLocks(address _account) view returns((uint256,uint256,uint256)[])
func (_Vesting *VestingSession) GetTimeLocks(_account common.Address) ([]TimeLockEntry, error) {
	return _Vesting.Contract.GetTimeLocks(&_Vesting.CallOpts, _account)
}

// GetTimeLocks is a free data retrieval call binding the contract method 0xf0e81024.
//
// Solidity: function getTimeLocks(address _account) view returns((uint256,uint256,uint256)[])
func (_Vesting *VestingCallerSession) GetTimeLocks(_account common.Address) ([]TimeLockEntry, error) {
	return _Vesting.Contract.GetTimeLocks(&_Vesting.CallOpts, _account)
}

// WithdrawableBalanceOf is a free data retrieval call binding the contract method 0x3aa6c4f9.
//
// Solidity: function withdrawableBalanceOf(address _account) view returns(uint256)
func (_Vesting *VestingCaller) WithdrawableBalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "withdrawableBalanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableBalanceOf is a free data retrieval call binding the contract method 0x3aa6c4f9.
//
// Solidity: function withdrawableBalanceOf(address _account) view returns(uint256)
func (_Vesting *VestingSession) WithdrawableBalanceOf(_account common.Address) (*big.Int, error) {
	return _Vesting.Contract.WithdrawableBalanceOf(&_Vesting.CallOpts, _account)
}

// WithdrawableBalanceOf is a free data retrieval call binding the contract method 0x3aa6c4f9.
//
// Solidity: function withdrawableBalanceOf(address _account) view returns(uint256)
func (_Vesting *VestingCallerSession) WithdrawableBalanceOf(_account common.Address) (*big.Int, error) {
	return _Vesting.Contract.WithdrawableBalanceOf(&_Vesting.CallOpts, _account)
}

// AnnounceUnlock is a paid mutator transaction binding the contract method 0xa1c9e7c6.
//
// Solidity: function announceUnlock(uint256 _amount) returns()
func (_Vesting *VestingTransactor) AnnounceUnlock(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "announceUnlock", _amount)
}

// AnnounceUnlock is a paid mutator transaction binding the contract method 0xa1c9e7c6.
//
// Solidity: function announceUnlock(uint256 _amount) returns()
func (_Vesting *VestingSession) AnnounceUnlock(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.AnnounceUnlock(&_Vesting.TransactOpts, _amount)
}

// AnnounceUnlock is a paid mutator transaction binding the contract method 0xa1c9e7c6.
//
// Solidity: function announceUnlock(uint256 _amount) returns()
func (_Vesting *VestingTransactorSession) AnnounceUnlock(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.AnnounceUnlock(&_Vesting.TransactOpts, _amount)
}

// DepositOnBehalfOf is a paid mutator transaction binding the contract method 0x73983f43.
//
// Solidity: function depositOnBehalfOf(address _account, uint256 _releaseStart, uint256 _releaseEnd) payable returns()
func (_Vesting *VestingTransactor) DepositOnBehalfOf(opts *bind.TransactOpts, _account common.Address, _releaseStart *big.Int, _releaseEnd *big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "depositOnBehalfOf", _account, _releaseStart, _releaseEnd)
}

// DepositOnBehalfOf is a paid mutator transaction binding the contract method 0x73983f43.
//
// Solidity: function depositOnBehalfOf(address _account, uint256 _releaseStart, uint256 _releaseEnd) payable returns()
func (_Vesting *VestingSession) DepositOnBehalfOf(_account common.Address, _releaseStart *big.Int, _releaseEnd *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.DepositOnBehalfOf(&_Vesting.TransactOpts, _account, _releaseStart, _releaseEnd)
}

// DepositOnBehalfOf is a paid mutator transaction binding the contract method 0x73983f43.
//
// Solidity: function depositOnBehalfOf(address _account, uint256 _releaseStart, uint256 _releaseEnd) payable returns()
func (_Vesting *VestingTransactorSession) DepositOnBehalfOf(_account common.Address, _releaseStart *big.Int, _releaseEnd *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.DepositOnBehalfOf(&_Vesting.TransactOpts, _account, _releaseStart, _releaseEnd)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_Vesting *VestingTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_Vesting *VestingSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.Initialize(&_Vesting.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_Vesting *VestingTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.Initialize(&_Vesting.TransactOpts, _registry)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 _amount) returns()
func (_Vesting *VestingTransactor) Lock(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "lock", _amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 _amount) returns()
func (_Vesting *VestingSession) Lock(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Lock(&_Vesting.TransactOpts, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 _amount) returns()
func (_Vesting *VestingTransactorSession) Lock(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Lock(&_Vesting.TransactOpts, _amount)
}

// PurgeTimeLocks is a paid mutator transaction binding the contract method 0xafb60cff.
//
// Solidity: function purgeTimeLocks(address _account) returns()
func (_Vesting *VestingTransactor) PurgeTimeLocks(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "purgeTimeLocks", _account)
}

// PurgeTimeLocks is a paid mutator transaction binding the contract method 0xafb60cff.
//
// Solidity: function purgeTimeLocks(address _account) returns()
func (_Vesting *VestingSession) PurgeTimeLocks(_account common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.PurgeTimeLocks(&_Vesting.TransactOpts, _account)
}

// PurgeTimeLocks is a paid mutator transaction binding the contract method 0xafb60cff.
//
// Solidity: function purgeTimeLocks(address _account) returns()
func (_Vesting *VestingTransactorSession) PurgeTimeLocks(_account common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.PurgeTimeLocks(&_Vesting.TransactOpts, _account)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 _amount) returns()
func (_Vesting *VestingTransactor) Unlock(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "unlock", _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 _amount) returns()
func (_Vesting *VestingSession) Unlock(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Unlock(&_Vesting.TransactOpts, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 _amount) returns()
func (_Vesting *VestingTransactorSession) Unlock(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Unlock(&_Vesting.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Vesting *VestingTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Vesting *VestingSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Withdraw(&_Vesting.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Vesting *VestingTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Withdraw(&_Vesting.TransactOpts, _amount)
}

// VestingNewTimeLockIterator is returned from FilterNewTimeLock and is used to iterate over the raw logs and unpacked data for NewTimeLock events raised by the Vesting contract.
type VestingNewTimeLockIterator struct {
	Event *VestingNewTimeLock // Event containing the contract specifics and raw log

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
func (it *VestingNewTimeLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingNewTimeLock)
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
		it.Event = new(VestingNewTimeLock)
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
func (it *VestingNewTimeLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingNewTimeLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingNewTimeLock represents a NewTimeLock event raised by the Vesting contract.
type VestingNewTimeLock struct {
	Account  common.Address
	Timelock TimeLockEntry
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewTimeLock is a free log retrieval operation binding the contract event 0x42dc784443c8161a28f39bb05a48bc8bac980ab9115d5b9cfef307f00204bd7f.
//
// Solidity: event NewTimeLock(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Vesting *VestingFilterer) FilterNewTimeLock(opts *bind.FilterOpts, _account []common.Address) (*VestingNewTimeLockIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "NewTimeLock", _accountRule)
	if err != nil {
		return nil, err
	}
	return &VestingNewTimeLockIterator{contract: _Vesting.contract, event: "NewTimeLock", logs: logs, sub: sub}, nil
}

// WatchNewTimeLock is a free log subscription operation binding the contract event 0x42dc784443c8161a28f39bb05a48bc8bac980ab9115d5b9cfef307f00204bd7f.
//
// Solidity: event NewTimeLock(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Vesting *VestingFilterer) WatchNewTimeLock(opts *bind.WatchOpts, sink chan<- *VestingNewTimeLock, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "NewTimeLock", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingNewTimeLock)
				if err := _Vesting.contract.UnpackLog(event, "NewTimeLock", log); err != nil {
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

// ParseNewTimeLock is a log parse operation binding the contract event 0x42dc784443c8161a28f39bb05a48bc8bac980ab9115d5b9cfef307f00204bd7f.
//
// Solidity: event NewTimeLock(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Vesting *VestingFilterer) ParseNewTimeLock(log types.Log) (*VestingNewTimeLock, error) {
	event := new(VestingNewTimeLock)
	if err := _Vesting.contract.UnpackLog(event, "NewTimeLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingPurgedIterator is returned from FilterPurged and is used to iterate over the raw logs and unpacked data for Purged events raised by the Vesting contract.
type VestingPurgedIterator struct {
	Event *VestingPurged // Event containing the contract specifics and raw log

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
func (it *VestingPurgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingPurged)
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
		it.Event = new(VestingPurged)
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
func (it *VestingPurgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingPurgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingPurged represents a Purged event raised by the Vesting contract.
type VestingPurged struct {
	Account  common.Address
	Timelock TimeLockEntry
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPurged is a free log retrieval operation binding the contract event 0x646189d6820ad36442dfb379131756840be21178303b28f1f607ca4e45762e93.
//
// Solidity: event Purged(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Vesting *VestingFilterer) FilterPurged(opts *bind.FilterOpts, _account []common.Address) (*VestingPurgedIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "Purged", _accountRule)
	if err != nil {
		return nil, err
	}
	return &VestingPurgedIterator{contract: _Vesting.contract, event: "Purged", logs: logs, sub: sub}, nil
}

// WatchPurged is a free log subscription operation binding the contract event 0x646189d6820ad36442dfb379131756840be21178303b28f1f607ca4e45762e93.
//
// Solidity: event Purged(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Vesting *VestingFilterer) WatchPurged(opts *bind.WatchOpts, sink chan<- *VestingPurged, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "Purged", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingPurged)
				if err := _Vesting.contract.UnpackLog(event, "Purged", log); err != nil {
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

// ParsePurged is a log parse operation binding the contract event 0x646189d6820ad36442dfb379131756840be21178303b28f1f607ca4e45762e93.
//
// Solidity: event Purged(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Vesting *VestingFilterer) ParsePurged(log types.Log) (*VestingPurged, error) {
	event := new(VestingPurged)
	if err := _Vesting.contract.UnpackLog(event, "Purged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
