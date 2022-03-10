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

// TokenBridgeAdminProxyMetaData contains all meta data concerning the TokenBridgeAdminProxy contract.
var TokenBridgeAdminProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeValidators\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateTokenbridgeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610d65806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063485cc955146100c0578063715018a6146100d55780638da5cb5b146100dd578063b40ba227146100fb578063f2fde38b14610103575b33610060610116565b6001600160a01b03161461008f5760405162461bcd60e51b815260040161008690610a2d565b60405180910390fd5b6066546001600160a01b0316366000803760008036600080855af13d6000803e8080156100bb573d6000f35b3d6000fd5b6100d36100ce366004610a77565b610125565b005b6100d36101cb565b6100e5610116565b6040516100f29190610ab0565b60405180910390f35b6100d3610206565b6100d3610111366004610ac4565b61021e565b6033546001600160a01b031690565b600054610100900460ff168061013e575060005460ff16155b61015a5760405162461bcd60e51b815260040161008690610ae8565b600054610100900460ff1615801561017c576000805461ffff19166101011790555b606580546001600160a01b038086166001600160a01b03199283161790925560668054928516929091169190911790556101b46102bb565b80156101c6576000805461ff00191690555b505050565b336101d4610116565b6001600160a01b0316146101fa5760405162461bcd60e51b815260040161008690610a2d565b6102046000610336565b565b6000610210610388565b905061021b816106b3565b50565b33610227610116565b6001600160a01b03161461024d5760405162461bcd60e51b815260040161008690610a2d565b6001600160a01b0381166102b25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610086565b61021b81610336565b600054610100900460ff16806102d4575060005460ff16155b6102f05760405162461bcd60e51b815260040161008690610ae8565b600054610100900460ff16158015610312576000805461ffff19166101011790555b61031a610963565b6103226109cd565b801561021b576000805461ff001916905550565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6065546040805180820182526015815274676f7665726e616e63652e76616c696461746f727360581b60208201529051633fb9027160e01b81526060926000926001600160a01b0390911691633fb90271916103e691600401610b36565b60206040518083038186803b1580156103fe57600080fd5b505afa158015610412573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104369190610b9b565b6001600160a01b0316633d6206286040518163ffffffff1660e01b815260040160006040518083038186803b15801561046e57600080fd5b505afa158015610482573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104aa9190810190610bce565b80516065546040805160608101909152602280825293945091926000926001600160a01b039092169163bf40fac191610d0e60208301396040518263ffffffff1660e01b81526004016104fd9190610b36565b60206040518083038186803b15801561051557600080fd5b505afa158015610529573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061054d9190610b9b565b60405162498bff60e81b815260206004820152601d60248201527f676f7665726e65642e455044522e6d6178544276616c696461746f727300000060448201526001600160a01b03919091169063498bff009060640160206040518083038186803b1580156105bb57600080fd5b505afa1580156105cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105f39190610c93565b90508251811015610602578091505b60008267ffffffffffffffff81111561061d5761061d610bb8565b604051908082528060200260200182016040528015610646578160200160208202803683370190505b50905060005b838110156106aa5784818151811061066657610666610cac565b602002602001015182828151811061068057610680610cac565b6001600160a01b0390921660209283029190910190910152806106a281610cc2565b91505061064c565b50949350505050565b60665460408051635890ef7960e01b815290516000926001600160a01b031691635890ef799160048083019286929190829003018186803b1580156106f757600080fd5b505afa15801561070b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107339190810190610bce565b905060005b815181101561083c57600082828151811061075557610755610cac565b60200260200101519050600080600090505b85518110156107be5785818151811061078257610782610cac565b60200260200101516001600160a01b0316836001600160a01b031614156107ac57600191506107be565b806107b681610cc2565b915050610767565b5080610827576066546040516340a141ff60e01b81526001600160a01b03909116906340a141ff906107f4908590600401610ab0565b600060405180830381600087803b15801561080e57600080fd5b505af1158015610822573d6000803e3d6000fd5b505050505b5050808061083490610cc2565b915050610738565b5060005b82518110156101c657600083828151811061085d5761085d610cac565b602090810291909101015160665460405163facd743b60e01b81529192506001600160a01b03169063facd743b90610899908490600401610ab0565b60206040518083038186803b1580156108b157600080fd5b505afa1580156108c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e99190610ceb565b61095057606654604051632691c64760e11b81526001600160a01b0390911690634d238c8e9061091d908490600401610ab0565b600060405180830381600087803b15801561093757600080fd5b505af115801561094b573d6000803e3d6000fd5b505050505b508061095b81610cc2565b915050610840565b600054610100900460ff168061097c575060005460ff16155b6109985760405162461bcd60e51b815260040161008690610ae8565b600054610100900460ff16158015610322576000805461ffff1916610101179055801561021b576000805461ff001916905550565b600054610100900460ff16806109e6575060005460ff16155b610a025760405162461bcd60e51b815260040161008690610ae8565b600054610100900460ff16158015610a24576000805461ffff19166101011790555b61032233610336565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6001600160a01b038116811461021b57600080fd5b60008060408385031215610a8a57600080fd5b8235610a9581610a62565b91506020830135610aa581610a62565b809150509250929050565b6001600160a01b0391909116815260200190565b600060208284031215610ad657600080fd5b8135610ae181610a62565b9392505050565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b600060208083528351808285015260005b81811015610b6357858101830151858201604001528201610b47565b81811115610b75576000604083870101525b50601f01601f1916929092016040019392505050565b8051610b9681610a62565b919050565b600060208284031215610bad57600080fd5b8151610ae181610a62565b634e487b7160e01b600052604160045260246000fd5b60006020808385031215610be157600080fd5b825167ffffffffffffffff80821115610bf957600080fd5b818501915085601f830112610c0d57600080fd5b815181811115610c1f57610c1f610bb8565b8060051b604051601f19603f83011681018181108582111715610c4457610c44610bb8565b604052918252848201925083810185019188831115610c6257600080fd5b938501935b82851015610c8757610c7885610b8b565b84529385019392850192610c67565b98975050505050505050565b600060208284031215610ca557600080fd5b5051919050565b634e487b7160e01b600052603260045260246000fd5b6000600019821415610ce457634e487b7160e01b600052601160045260246000fd5b5060010190565b600060208284031215610cfd57600080fd5b81518015158114610ae157600080fdfe676f7665726e616e63652e657870657274732e455044522e706172616d6574657273a2646970667358221220656d70652d0675eda45b7108bcee5cdb0f18fc5f73b5835cef4f15c28d6c791f64736f6c63430008090033",
}

// TokenBridgeAdminProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenBridgeAdminProxyMetaData.ABI instead.
var TokenBridgeAdminProxyABI = TokenBridgeAdminProxyMetaData.ABI

// TokenBridgeAdminProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenBridgeAdminProxyMetaData.Bin instead.
var TokenBridgeAdminProxyBin = TokenBridgeAdminProxyMetaData.Bin

// DeployTokenBridgeAdminProxy deploys a new Ethereum contract, binding an instance of TokenBridgeAdminProxy to it.
func DeployTokenBridgeAdminProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenBridgeAdminProxy, error) {
	parsed, err := TokenBridgeAdminProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenBridgeAdminProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenBridgeAdminProxy{TokenBridgeAdminProxyCaller: TokenBridgeAdminProxyCaller{contract: contract}, TokenBridgeAdminProxyTransactor: TokenBridgeAdminProxyTransactor{contract: contract}, TokenBridgeAdminProxyFilterer: TokenBridgeAdminProxyFilterer{contract: contract}}, nil
}

// TokenBridgeAdminProxy is an auto generated Go binding around an Ethereum contract.
type TokenBridgeAdminProxy struct {
	TokenBridgeAdminProxyCaller     // Read-only binding to the contract
	TokenBridgeAdminProxyTransactor // Write-only binding to the contract
	TokenBridgeAdminProxyFilterer   // Log filterer for contract events
}

// TokenBridgeAdminProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenBridgeAdminProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBridgeAdminProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenBridgeAdminProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBridgeAdminProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenBridgeAdminProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBridgeAdminProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenBridgeAdminProxySession struct {
	Contract     *TokenBridgeAdminProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TokenBridgeAdminProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenBridgeAdminProxyCallerSession struct {
	Contract *TokenBridgeAdminProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TokenBridgeAdminProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenBridgeAdminProxyTransactorSession struct {
	Contract     *TokenBridgeAdminProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TokenBridgeAdminProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenBridgeAdminProxyRaw struct {
	Contract *TokenBridgeAdminProxy // Generic contract binding to access the raw methods on
}

// TokenBridgeAdminProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenBridgeAdminProxyCallerRaw struct {
	Contract *TokenBridgeAdminProxyCaller // Generic read-only contract binding to access the raw methods on
}

// TokenBridgeAdminProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenBridgeAdminProxyTransactorRaw struct {
	Contract *TokenBridgeAdminProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenBridgeAdminProxy creates a new instance of TokenBridgeAdminProxy, bound to a specific deployed contract.
func NewTokenBridgeAdminProxy(address common.Address, backend bind.ContractBackend) (*TokenBridgeAdminProxy, error) {
	contract, err := bindTokenBridgeAdminProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeAdminProxy{TokenBridgeAdminProxyCaller: TokenBridgeAdminProxyCaller{contract: contract}, TokenBridgeAdminProxyTransactor: TokenBridgeAdminProxyTransactor{contract: contract}, TokenBridgeAdminProxyFilterer: TokenBridgeAdminProxyFilterer{contract: contract}}, nil
}

// NewTokenBridgeAdminProxyCaller creates a new read-only instance of TokenBridgeAdminProxy, bound to a specific deployed contract.
func NewTokenBridgeAdminProxyCaller(address common.Address, caller bind.ContractCaller) (*TokenBridgeAdminProxyCaller, error) {
	contract, err := bindTokenBridgeAdminProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeAdminProxyCaller{contract: contract}, nil
}

// NewTokenBridgeAdminProxyTransactor creates a new write-only instance of TokenBridgeAdminProxy, bound to a specific deployed contract.
func NewTokenBridgeAdminProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenBridgeAdminProxyTransactor, error) {
	contract, err := bindTokenBridgeAdminProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeAdminProxyTransactor{contract: contract}, nil
}

// NewTokenBridgeAdminProxyFilterer creates a new log filterer instance of TokenBridgeAdminProxy, bound to a specific deployed contract.
func NewTokenBridgeAdminProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenBridgeAdminProxyFilterer, error) {
	contract, err := bindTokenBridgeAdminProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeAdminProxyFilterer{contract: contract}, nil
}

// bindTokenBridgeAdminProxy binds a generic wrapper to an already deployed contract.
func bindTokenBridgeAdminProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenBridgeAdminProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenBridgeAdminProxy.Contract.TokenBridgeAdminProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.TokenBridgeAdminProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.TokenBridgeAdminProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenBridgeAdminProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenBridgeAdminProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxySession) Owner() (common.Address, error) {
	return _TokenBridgeAdminProxy.Contract.Owner(&_TokenBridgeAdminProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyCallerSession) Owner() (common.Address, error) {
	return _TokenBridgeAdminProxy.Contract.Owner(&_TokenBridgeAdminProxy.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _bridgeValidators) returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _bridgeValidators common.Address) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.contract.Transact(opts, "initialize", _registry, _bridgeValidators)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _bridgeValidators) returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxySession) Initialize(_registry common.Address, _bridgeValidators common.Address) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.Initialize(&_TokenBridgeAdminProxy.TransactOpts, _registry, _bridgeValidators)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _bridgeValidators) returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyTransactorSession) Initialize(_registry common.Address, _bridgeValidators common.Address) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.Initialize(&_TokenBridgeAdminProxy.TransactOpts, _registry, _bridgeValidators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.RenounceOwnership(&_TokenBridgeAdminProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.RenounceOwnership(&_TokenBridgeAdminProxy.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.TransferOwnership(&_TokenBridgeAdminProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.TransferOwnership(&_TokenBridgeAdminProxy.TransactOpts, newOwner)
}

// UpdateTokenbridgeValidators is a paid mutator transaction binding the contract method 0xb40ba227.
//
// Solidity: function updateTokenbridgeValidators() returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyTransactor) UpdateTokenbridgeValidators(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.contract.Transact(opts, "updateTokenbridgeValidators")
}

// UpdateTokenbridgeValidators is a paid mutator transaction binding the contract method 0xb40ba227.
//
// Solidity: function updateTokenbridgeValidators() returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxySession) UpdateTokenbridgeValidators() (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.UpdateTokenbridgeValidators(&_TokenBridgeAdminProxy.TransactOpts)
}

// UpdateTokenbridgeValidators is a paid mutator transaction binding the contract method 0xb40ba227.
//
// Solidity: function updateTokenbridgeValidators() returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyTransactorSession) UpdateTokenbridgeValidators() (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.UpdateTokenbridgeValidators(&_TokenBridgeAdminProxy.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.Fallback(&_TokenBridgeAdminProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TokenBridgeAdminProxy.Contract.Fallback(&_TokenBridgeAdminProxy.TransactOpts, calldata)
}

// TokenBridgeAdminProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenBridgeAdminProxy contract.
type TokenBridgeAdminProxyOwnershipTransferredIterator struct {
	Event *TokenBridgeAdminProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenBridgeAdminProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeAdminProxyOwnershipTransferred)
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
		it.Event = new(TokenBridgeAdminProxyOwnershipTransferred)
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
func (it *TokenBridgeAdminProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeAdminProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeAdminProxyOwnershipTransferred represents a OwnershipTransferred event raised by the TokenBridgeAdminProxy contract.
type TokenBridgeAdminProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenBridgeAdminProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenBridgeAdminProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeAdminProxyOwnershipTransferredIterator{contract: _TokenBridgeAdminProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenBridgeAdminProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenBridgeAdminProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeAdminProxyOwnershipTransferred)
				if err := _TokenBridgeAdminProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyFilterer) ParseOwnershipTransferred(log types.Log) (*TokenBridgeAdminProxyOwnershipTransferred, error) {
	event := new(TokenBridgeAdminProxyOwnershipTransferred)
	if err := _TokenBridgeAdminProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeAdminProxyValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the TokenBridgeAdminProxy contract.
type TokenBridgeAdminProxyValidatorAddedIterator struct {
	Event *TokenBridgeAdminProxyValidatorAdded // Event containing the contract specifics and raw log

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
func (it *TokenBridgeAdminProxyValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeAdminProxyValidatorAdded)
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
		it.Event = new(TokenBridgeAdminProxyValidatorAdded)
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
func (it *TokenBridgeAdminProxyValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeAdminProxyValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeAdminProxyValidatorAdded represents a ValidatorAdded event raised by the TokenBridgeAdminProxy contract.
type TokenBridgeAdminProxyValidatorAdded struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed _validator)
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyFilterer) FilterValidatorAdded(opts *bind.FilterOpts, _validator []common.Address) (*TokenBridgeAdminProxyValidatorAddedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _TokenBridgeAdminProxy.contract.FilterLogs(opts, "ValidatorAdded", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeAdminProxyValidatorAddedIterator{contract: _TokenBridgeAdminProxy.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed _validator)
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *TokenBridgeAdminProxyValidatorAdded, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _TokenBridgeAdminProxy.contract.WatchLogs(opts, "ValidatorAdded", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeAdminProxyValidatorAdded)
				if err := _TokenBridgeAdminProxy.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed _validator)
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyFilterer) ParseValidatorAdded(log types.Log) (*TokenBridgeAdminProxyValidatorAdded, error) {
	event := new(TokenBridgeAdminProxyValidatorAdded)
	if err := _TokenBridgeAdminProxy.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeAdminProxyValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the TokenBridgeAdminProxy contract.
type TokenBridgeAdminProxyValidatorRemovedIterator struct {
	Event *TokenBridgeAdminProxyValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *TokenBridgeAdminProxyValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeAdminProxyValidatorRemoved)
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
		it.Event = new(TokenBridgeAdminProxyValidatorRemoved)
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
func (it *TokenBridgeAdminProxyValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeAdminProxyValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeAdminProxyValidatorRemoved represents a ValidatorRemoved event raised by the TokenBridgeAdminProxy contract.
type TokenBridgeAdminProxyValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed _validator)
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, _validator []common.Address) (*TokenBridgeAdminProxyValidatorRemovedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _TokenBridgeAdminProxy.contract.FilterLogs(opts, "ValidatorRemoved", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeAdminProxyValidatorRemovedIterator{contract: _TokenBridgeAdminProxy.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed _validator)
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *TokenBridgeAdminProxyValidatorRemoved, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _TokenBridgeAdminProxy.contract.WatchLogs(opts, "ValidatorRemoved", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeAdminProxyValidatorRemoved)
				if err := _TokenBridgeAdminProxy.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed _validator)
func (_TokenBridgeAdminProxy *TokenBridgeAdminProxyFilterer) ParseValidatorRemoved(log types.Log) (*TokenBridgeAdminProxyValidatorRemoved, error) {
	event := new(TokenBridgeAdminProxyValidatorRemoved)
	if err := _TokenBridgeAdminProxy.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
