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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beneficiaries\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"shares\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_beneficiaries\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_shares\",\"type\":\"string[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610d8a806100206000396000f3fe6080604052600436106100435760003560e01c806357a858fc1461004f578063abaa991614610085578063efe86a5e1461009c578063efeb5e58146100bc57600080fd5b3661004a57005b600080fd5b34801561005b57600080fd5b5061006f61006a3660046109e8565b6100dc565b60405161007c9190610a01565b60405180910390f35b34801561009157600080fd5b5061009a610188565b005b3480156100a857600080fd5b5061009a6100b7366004610ba5565b610700565b3480156100c857600080fd5b5061006f6100d73660046109e8565b61088c565b600281815481106100ec57600080fd5b90600052602060002001600091509050805461010790610c1a565b80601f016020809104026020016040519081016040528092919081815260200182805461013390610c1a565b80156101805780601f1061015557610100808354040283529160200191610180565b820191906000526020600020905b81548152906001019060200180831161016357829003601f168201915b505050505081565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060228152602001610d33602291396040518263ffffffff1660e01b81526004016101e29190610a01565b60206040518083038186803b1580156101fa57600080fd5b505afa15801561020e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102329190610c55565b905060006002805480602002602001604051908101604052809291908181526020016000905b8282101561030457838290600052602060002001805461027790610c1a565b80601f01602080910402602001604051908101604052809291908181526020018280546102a390610c1a565b80156102f05780601f106102c5576101008083540402835291602001916102f0565b820191906000526020600020905b8154815290600101906020018083116102d357829003601f168201915b505050505081526020019060010190610258565b505050509050600081516001600160401b0381111561032557610325610a6b565b60405190808252806020026020018201604052801561034e578160200160208202803683370190505b5090506000805b835181101561044a57846001600160a01b031663498bff0085838151811061037f5761037f610c79565b60200260200101516040518263ffffffff1660e01b81526004016103a39190610a01565b60206040518083038186803b1580156103bb57600080fd5b505afa1580156103cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f39190610c8f565b83828151811061040557610405610c79565b60200260200101818152505082818151811061042357610423610c79565b6020026020010151826104369190610cbe565b91508061044281610cd6565b915050610355565b5060006001805480602002602001604051908101604052809291908181526020016000905b8282101561051b57838290600052602060002001805461048e90610c1a565b80601f01602080910402602001604051908101604052809291908181526020018280546104ba90610c1a565b80156105075780601f106104dc57610100808354040283529160200191610507565b820191906000526020600020905b8154815290600101906020018083116104ea57829003601f168201915b50505050508152602001906001019061046f565b50505050905060008047905060005b83518110156106f65760008587838151811061054857610548610c79565b60200260200101518461055b9190610cf1565b6105659190610d10565b9050600060029054906101000a90046001600160a01b03166001600160a01b0316633fb9027186848151811061059d5761059d610c79565b60200260200101516040518263ffffffff1660e01b81526004016105c19190610a01565b60206040518083038186803b1580156105d957600080fd5b505afa1580156105ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106119190610c55565b93506000846001600160a01b03168260405160006040518083038185875af1925050503d8060008114610660576040519150601f19603f3d011682016040523d82523d6000602084013e610665565b606091505b50509050806106e15760405162461bcd60e51b815260206004820152603e60248201527f5b5145432d3031313030315d2d4661696c656420746f207472616e736665722060448201527f74686520616d6f756e742c20616c6c6f636174696f6e206661696c65642e000060648201526084015b60405180910390fd5b505080806106ee90610cd6565b91505061052a565b5050505050505050565b600054610100900460ff1680610719575060005460ff16155b61077c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106d8565b600054610100900460ff1615801561079e576000805461ffff19166101011790555b815183511461082a5760405162461bcd60e51b815260206004820152604c60248201527f5b5145432d3031313030305d2d4e756d626572206f662062656e65666963696160448201527f726965732073686f756c64206265207468652073616d65206173206e756d626560648201526b391037b31039b430b932b99760a11b608482015260a4016106d8565b825161083d90600190602086019061089c565b50815161085190600290602085019061089c565b506000805462010000600160b01b031916620100006001600160a01b038716021790558015610886576000805461ff00191690555b50505050565b600181815481106100ec57600080fd5b8280548282559060005260206000209081019282156108e9579160200282015b828111156108e957825180516108d99184916020909101906108f9565b50916020019190600101906108bc565b506108f5929150610979565b5090565b82805461090590610c1a565b90600052602060002090601f016020900481019282610927576000855561096d565b82601f1061094057805160ff191683800117855561096d565b8280016001018555821561096d579182015b8281111561096d578251825591602001919060010190610952565b506108f5929150610996565b808211156108f557600061098d82826109ab565b50600101610979565b5b808211156108f55760008155600101610997565b5080546109b790610c1a565b6000825580601f106109c7575050565b601f0160209004906000526020600020908101906109e59190610996565b50565b6000602082840312156109fa57600080fd5b5035919050565b600060208083528351808285015260005b81811015610a2e57858101830151858201604001528201610a12565b81811115610a40576000604083870101525b50601f01601f1916929092016040019392505050565b6001600160a01b03811681146109e557600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715610aa957610aa9610a6b565b604052919050565b6000601f8381840112610ac357600080fd5b823560206001600160401b0380831115610adf57610adf610a6b565b8260051b610aee838201610a81565b9384528681018301938381019089861115610b0857600080fd5b84890192505b85831015610b9857823584811115610b265760008081fd5b8901603f81018b13610b385760008081fd5b85810135604086821115610b4e57610b4e610a6b565b610b5f828b01601f19168901610a81565b8281528d82848601011115610b745760008081fd5b828285018a8301376000928101890192909252508352509184019190840190610b0e565b9998505050505050505050565b600080600060608486031215610bba57600080fd5b8335610bc581610a56565b925060208401356001600160401b0380821115610be157600080fd5b610bed87838801610ab1565b93506040860135915080821115610c0357600080fd5b50610c1086828701610ab1565b9150509250925092565b600181811c90821680610c2e57607f821691505b60208210811415610c4f57634e487b7160e01b600052602260045260246000fd5b50919050565b600060208284031215610c6757600080fd5b8151610c7281610a56565b9392505050565b634e487b7160e01b600052603260045260246000fd5b600060208284031215610ca157600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b60008219821115610cd157610cd1610ca8565b500190565b6000600019821415610cea57610cea610ca8565b5060010190565b6000816000190483118215151615610d0b57610d0b610ca8565b500290565b600082610d2d57634e487b7160e01b600052601260045260246000fd5b50049056fe676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273a2646970667358221220364ccb1fb84903a8ef7c05d0d1518bccd3f215b05a9e1284acc08db868e2ddec64736f6c63430008090033",
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
