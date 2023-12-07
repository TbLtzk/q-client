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

// DefaultAllocationProxyMetaData contains all meta data concerning the DefaultAllocationProxy contract.
var DefaultAllocationProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Allocated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beneficiaries\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_beneficiaries\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_shares\",\"type\":\"string[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"shares\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610dcc806100206000396000f3fe6080604052600436106100435760003560e01c806357a858fc1461004f578063abaa991614610085578063efe86a5e1461009c578063efeb5e58146100bc57600080fd5b3661004a57005b600080fd5b34801561005b57600080fd5b5061006f61006a366004610a2a565b6100dc565b60405161007c9190610a43565b60405180910390f35b34801561009157600080fd5b5061009a610188565b005b3480156100a857600080fd5b5061009a6100b7366004610be7565b610742565b3480156100c857600080fd5b5061006f6100d7366004610a2a565b6108ce565b600281815481106100ec57600080fd5b90600052602060002001600091509050805461010790610c5c565b80601f016020809104026020016040519081016040528092919081815260200182805461013390610c5c565b80156101805780601f1061015557610100808354040283529160200191610180565b820191906000526020600020905b81548152906001019060200180831161016357829003601f168201915b505050505081565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060228152602001610d75602291396040518263ffffffff1660e01b81526004016101e29190610a43565b60206040518083038186803b1580156101fa57600080fd5b505afa15801561020e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102329190610c97565b905060006002805480602002602001604051908101604052809291908181526020016000905b8282101561030457838290600052602060002001805461027790610c5c565b80601f01602080910402602001604051908101604052809291908181526020018280546102a390610c5c565b80156102f05780601f106102c5576101008083540402835291602001916102f0565b820191906000526020600020905b8154815290600101906020018083116102d357829003601f168201915b505050505081526020019060010190610258565b505050509050600081516001600160401b0381111561032557610325610aad565b60405190808252806020026020018201604052801561034e578160200160208202803683370190505b5090506000805b835181101561044a57846001600160a01b031663498bff0085838151811061037f5761037f610cbb565b60200260200101516040518263ffffffff1660e01b81526004016103a39190610a43565b60206040518083038186803b1580156103bb57600080fd5b505afa1580156103cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f39190610cd1565b83828151811061040557610405610cbb565b60200260200101818152505082818151811061042357610423610cbb565b6020026020010151826104369190610d00565b91508061044281610d18565b915050610355565b5060006001805480602002602001604051908101604052809291908181526020016000905b8282101561051b57838290600052602060002001805461048e90610c5c565b80601f01602080910402602001604051908101604052809291908181526020018280546104ba90610c5c565b80156105075780601f106104dc57610100808354040283529160200191610507565b820191906000526020600020905b8154815290600101906020018083116104ea57829003601f168201915b50505050508152602001906001019061046f565b505050509050600080600047905060005b84518110156107045760008688838151811061054a5761054a610cbb565b60200260200101518461055d9190610d33565b6105679190610d52565b90506105738185610d00565b9350600060029054906101000a90046001600160a01b03166001600160a01b0316633fb902718784815181106105ab576105ab610cbb565b60200260200101516040518263ffffffff1660e01b81526004016105cf9190610a43565b60206040518083038186803b1580156105e757600080fd5b505afa1580156105fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061f9190610c97565b94506000856001600160a01b03168260405160006040518083038185875af1925050503d806000811461066e576040519150601f19603f3d011682016040523d82523d6000602084013e610673565b606091505b50509050806106ef5760405162461bcd60e51b815260206004820152603e60248201527f5b5145432d3031313030315d2d4661696c656420746f207472616e736665722060448201527f74686520616d6f756e742c20616c6c6f636174696f6e206661696c65642e000060648201526084015b60405180910390fd5b505080806106fc90610d18565b91505061052c565b506040518281527f60a0c13c5aa0acfe5cd0c4f3ac6b9fe2742094a559183932b66e20108fc2be199060200160405180910390a15050505050505050565b600054610100900460ff168061075b575060005460ff16155b6107be5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106e6565b600054610100900460ff161580156107e0576000805461ffff19166101011790555b815183511461086c5760405162461bcd60e51b815260206004820152604c60248201527f5b5145432d3031313030305d2d4e756d626572206f662062656e65666963696160448201527f726965732073686f756c64206265207468652073616d65206173206e756d626560648201526b391037b31039b430b932b99760a11b608482015260a4016106e6565b825161087f9060019060208601906108de565b5081516108939060029060208501906108de565b506000805462010000600160b01b031916620100006001600160a01b0387160217905580156108c8576000805461ff00191690555b50505050565b600181815481106100ec57600080fd5b82805482825590600052602060002090810192821561092b579160200282015b8281111561092b578251805161091b91849160209091019061093b565b50916020019190600101906108fe565b506109379291506109bb565b5090565b82805461094790610c5c565b90600052602060002090601f01602090048101928261096957600085556109af565b82601f1061098257805160ff19168380011785556109af565b828001600101855582156109af579182015b828111156109af578251825591602001919060010190610994565b506109379291506109d8565b808211156109375760006109cf82826109ed565b506001016109bb565b5b8082111561093757600081556001016109d9565b5080546109f990610c5c565b6000825580601f10610a09575050565b601f016020900490600052602060002090810190610a2791906109d8565b50565b600060208284031215610a3c57600080fd5b5035919050565b600060208083528351808285015260005b81811015610a7057858101830151858201604001528201610a54565b81811115610a82576000604083870101525b50601f01601f1916929092016040019392505050565b6001600160a01b0381168114610a2757600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715610aeb57610aeb610aad565b604052919050565b6000601f8381840112610b0557600080fd5b823560206001600160401b0380831115610b2157610b21610aad565b8260051b610b30838201610ac3565b9384528681018301938381019089861115610b4a57600080fd5b84890192505b85831015610bda57823584811115610b685760008081fd5b8901603f81018b13610b7a5760008081fd5b85810135604086821115610b9057610b90610aad565b610ba1828b01601f19168901610ac3565b8281528d82848601011115610bb65760008081fd5b828285018a8301376000928101890192909252508352509184019190840190610b50565b9998505050505050505050565b600080600060608486031215610bfc57600080fd5b8335610c0781610a98565b925060208401356001600160401b0380821115610c2357600080fd5b610c2f87838801610af3565b93506040860135915080821115610c4557600080fd5b50610c5286828701610af3565b9150509250925092565b600181811c90821680610c7057607f821691505b60208210811415610c9157634e487b7160e01b600052602260045260246000fd5b50919050565b600060208284031215610ca957600080fd5b8151610cb481610a98565b9392505050565b634e487b7160e01b600052603260045260246000fd5b600060208284031215610ce357600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b60008219821115610d1357610d13610cea565b500190565b6000600019821415610d2c57610d2c610cea565b5060010190565b6000816000190483118215151615610d4d57610d4d610cea565b500290565b600082610d6f57634e487b7160e01b600052601260045260246000fd5b50049056fe676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273a26469706673582212202225f67c3879c79a6fc826f21d90ca609ec4181760eac6a6eb41a27ce5dd861364736f6c63430008090033",
}

// DefaultAllocationProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use DefaultAllocationProxyMetaData.ABI instead.
var DefaultAllocationProxyABI = DefaultAllocationProxyMetaData.ABI

// DefaultAllocationProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DefaultAllocationProxyMetaData.Bin instead.
var DefaultAllocationProxyBin = DefaultAllocationProxyMetaData.Bin

// DeployDefaultAllocationProxy deploys a new Ethereum contract, binding an instance of DefaultAllocationProxy to it.
func DeployDefaultAllocationProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DefaultAllocationProxy, error) {
	parsed, err := DefaultAllocationProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DefaultAllocationProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DefaultAllocationProxy{DefaultAllocationProxyCaller: DefaultAllocationProxyCaller{contract: contract}, DefaultAllocationProxyTransactor: DefaultAllocationProxyTransactor{contract: contract}, DefaultAllocationProxyFilterer: DefaultAllocationProxyFilterer{contract: contract}}, nil
}

// DefaultAllocationProxy is an auto generated Go binding around an Ethereum contract.
type DefaultAllocationProxy struct {
	DefaultAllocationProxyCaller     // Read-only binding to the contract
	DefaultAllocationProxyTransactor // Write-only binding to the contract
	DefaultAllocationProxyFilterer   // Log filterer for contract events
}

// DefaultAllocationProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefaultAllocationProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultAllocationProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefaultAllocationProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultAllocationProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefaultAllocationProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultAllocationProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefaultAllocationProxySession struct {
	Contract     *DefaultAllocationProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DefaultAllocationProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefaultAllocationProxyCallerSession struct {
	Contract *DefaultAllocationProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// DefaultAllocationProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefaultAllocationProxyTransactorSession struct {
	Contract     *DefaultAllocationProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// DefaultAllocationProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefaultAllocationProxyRaw struct {
	Contract *DefaultAllocationProxy // Generic contract binding to access the raw methods on
}

// DefaultAllocationProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefaultAllocationProxyCallerRaw struct {
	Contract *DefaultAllocationProxyCaller // Generic read-only contract binding to access the raw methods on
}

// DefaultAllocationProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefaultAllocationProxyTransactorRaw struct {
	Contract *DefaultAllocationProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefaultAllocationProxy creates a new instance of DefaultAllocationProxy, bound to a specific deployed contract.
func NewDefaultAllocationProxy(address common.Address, backend bind.ContractBackend) (*DefaultAllocationProxy, error) {
	contract, err := bindDefaultAllocationProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefaultAllocationProxy{DefaultAllocationProxyCaller: DefaultAllocationProxyCaller{contract: contract}, DefaultAllocationProxyTransactor: DefaultAllocationProxyTransactor{contract: contract}, DefaultAllocationProxyFilterer: DefaultAllocationProxyFilterer{contract: contract}}, nil
}

// NewDefaultAllocationProxyCaller creates a new read-only instance of DefaultAllocationProxy, bound to a specific deployed contract.
func NewDefaultAllocationProxyCaller(address common.Address, caller bind.ContractCaller) (*DefaultAllocationProxyCaller, error) {
	contract, err := bindDefaultAllocationProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefaultAllocationProxyCaller{contract: contract}, nil
}

// NewDefaultAllocationProxyTransactor creates a new write-only instance of DefaultAllocationProxy, bound to a specific deployed contract.
func NewDefaultAllocationProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*DefaultAllocationProxyTransactor, error) {
	contract, err := bindDefaultAllocationProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefaultAllocationProxyTransactor{contract: contract}, nil
}

// NewDefaultAllocationProxyFilterer creates a new log filterer instance of DefaultAllocationProxy, bound to a specific deployed contract.
func NewDefaultAllocationProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*DefaultAllocationProxyFilterer, error) {
	contract, err := bindDefaultAllocationProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefaultAllocationProxyFilterer{contract: contract}, nil
}

// bindDefaultAllocationProxy binds a generic wrapper to an already deployed contract.
func bindDefaultAllocationProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DefaultAllocationProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefaultAllocationProxy *DefaultAllocationProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefaultAllocationProxy.Contract.DefaultAllocationProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefaultAllocationProxy *DefaultAllocationProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.DefaultAllocationProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefaultAllocationProxy *DefaultAllocationProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.DefaultAllocationProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefaultAllocationProxy *DefaultAllocationProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefaultAllocationProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.contract.Transact(opts, method, params...)
}

// Beneficiaries is a free data retrieval call binding the contract method 0xefeb5e58.
//
// Solidity: function beneficiaries(uint256 ) view returns(string)
func (_DefaultAllocationProxy *DefaultAllocationProxyCaller) Beneficiaries(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _DefaultAllocationProxy.contract.Call(opts, &out, "beneficiaries", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Beneficiaries is a free data retrieval call binding the contract method 0xefeb5e58.
//
// Solidity: function beneficiaries(uint256 ) view returns(string)
func (_DefaultAllocationProxy *DefaultAllocationProxySession) Beneficiaries(arg0 *big.Int) (string, error) {
	return _DefaultAllocationProxy.Contract.Beneficiaries(&_DefaultAllocationProxy.CallOpts, arg0)
}

// Beneficiaries is a free data retrieval call binding the contract method 0xefeb5e58.
//
// Solidity: function beneficiaries(uint256 ) view returns(string)
func (_DefaultAllocationProxy *DefaultAllocationProxyCallerSession) Beneficiaries(arg0 *big.Int) (string, error) {
	return _DefaultAllocationProxy.Contract.Beneficiaries(&_DefaultAllocationProxy.CallOpts, arg0)
}

// Shares is a free data retrieval call binding the contract method 0x57a858fc.
//
// Solidity: function shares(uint256 ) view returns(string)
func (_DefaultAllocationProxy *DefaultAllocationProxyCaller) Shares(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _DefaultAllocationProxy.contract.Call(opts, &out, "shares", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0x57a858fc.
//
// Solidity: function shares(uint256 ) view returns(string)
func (_DefaultAllocationProxy *DefaultAllocationProxySession) Shares(arg0 *big.Int) (string, error) {
	return _DefaultAllocationProxy.Contract.Shares(&_DefaultAllocationProxy.CallOpts, arg0)
}

// Shares is a free data retrieval call binding the contract method 0x57a858fc.
//
// Solidity: function shares(uint256 ) view returns(string)
func (_DefaultAllocationProxy *DefaultAllocationProxyCallerSession) Shares(arg0 *big.Int) (string, error) {
	return _DefaultAllocationProxy.Contract.Shares(&_DefaultAllocationProxy.CallOpts, arg0)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactor) Allocate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultAllocationProxy.contract.Transact(opts, "allocate")
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_DefaultAllocationProxy *DefaultAllocationProxySession) Allocate() (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.Allocate(&_DefaultAllocationProxy.TransactOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactorSession) Allocate() (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.Allocate(&_DefaultAllocationProxy.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xefe86a5e.
//
// Solidity: function initialize(address _registry, string[] _beneficiaries, string[] _shares) returns()
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _beneficiaries []string, _shares []string) (*types.Transaction, error) {
	return _DefaultAllocationProxy.contract.Transact(opts, "initialize", _registry, _beneficiaries, _shares)
}

// Initialize is a paid mutator transaction binding the contract method 0xefe86a5e.
//
// Solidity: function initialize(address _registry, string[] _beneficiaries, string[] _shares) returns()
func (_DefaultAllocationProxy *DefaultAllocationProxySession) Initialize(_registry common.Address, _beneficiaries []string, _shares []string) (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.Initialize(&_DefaultAllocationProxy.TransactOpts, _registry, _beneficiaries, _shares)
}

// Initialize is a paid mutator transaction binding the contract method 0xefe86a5e.
//
// Solidity: function initialize(address _registry, string[] _beneficiaries, string[] _shares) returns()
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactorSession) Initialize(_registry common.Address, _beneficiaries []string, _shares []string) (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.Initialize(&_DefaultAllocationProxy.TransactOpts, _registry, _beneficiaries, _shares)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultAllocationProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefaultAllocationProxy *DefaultAllocationProxySession) Receive() (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.Receive(&_DefaultAllocationProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.Receive(&_DefaultAllocationProxy.TransactOpts)
}

// DefaultAllocationProxyAllocatedIterator is returned from FilterAllocated and is used to iterate over the raw logs and unpacked data for Allocated events raised by the DefaultAllocationProxy contract.
type DefaultAllocationProxyAllocatedIterator struct {
	Event *DefaultAllocationProxyAllocated // Event containing the contract specifics and raw log

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
func (it *DefaultAllocationProxyAllocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefaultAllocationProxyAllocated)
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
		it.Event = new(DefaultAllocationProxyAllocated)
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
func (it *DefaultAllocationProxyAllocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefaultAllocationProxyAllocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefaultAllocationProxyAllocated represents a Allocated event raised by the DefaultAllocationProxy contract.
type DefaultAllocationProxyAllocated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllocated is a free log retrieval operation binding the contract event 0x60a0c13c5aa0acfe5cd0c4f3ac6b9fe2742094a559183932b66e20108fc2be19.
//
// Solidity: event Allocated(uint256 amount)
func (_DefaultAllocationProxy *DefaultAllocationProxyFilterer) FilterAllocated(opts *bind.FilterOpts) (*DefaultAllocationProxyAllocatedIterator, error) {

	logs, sub, err := _DefaultAllocationProxy.contract.FilterLogs(opts, "Allocated")
	if err != nil {
		return nil, err
	}
	return &DefaultAllocationProxyAllocatedIterator{contract: _DefaultAllocationProxy.contract, event: "Allocated", logs: logs, sub: sub}, nil
}

// WatchAllocated is a free log subscription operation binding the contract event 0x60a0c13c5aa0acfe5cd0c4f3ac6b9fe2742094a559183932b66e20108fc2be19.
//
// Solidity: event Allocated(uint256 amount)
func (_DefaultAllocationProxy *DefaultAllocationProxyFilterer) WatchAllocated(opts *bind.WatchOpts, sink chan<- *DefaultAllocationProxyAllocated) (event.Subscription, error) {

	logs, sub, err := _DefaultAllocationProxy.contract.WatchLogs(opts, "Allocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefaultAllocationProxyAllocated)
				if err := _DefaultAllocationProxy.contract.UnpackLog(event, "Allocated", log); err != nil {
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

// ParseAllocated is a log parse operation binding the contract event 0x60a0c13c5aa0acfe5cd0c4f3ac6b9fe2742094a559183932b66e20108fc2be19.
//
// Solidity: event Allocated(uint256 amount)
func (_DefaultAllocationProxy *DefaultAllocationProxyFilterer) ParseAllocated(log types.Log) (*DefaultAllocationProxyAllocated, error) {
	event := new(DefaultAllocationProxyAllocated)
	if err := _DefaultAllocationProxy.contract.UnpackLog(event, "Allocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
