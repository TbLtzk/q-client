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

// ForeignChainTokenBridgeAdminProxyMetaData contains all meta data concerning the ForeignChainTokenBridgeAdminProxy contract.
var ForeignChainTokenBridgeAdminProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeValidators\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_newValidatorsList\",\"type\":\"address[]\"}],\"name\":\"updateTokenbridgeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610b84806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063278dcd62146100ea578063715018a6146100ff5780638da5cb5b14610107578063c4d66de814610125578063f2fde38b14610138575b61005f61014b565b6001600160a01b031661007061014f565b6001600160a01b0316146100b9576040805162461bcd60e51b81526020600482018190526024820152600080516020610b0f833981519152604482015290519081900360640190fd5b6065546001600160a01b0316366000803760008036600080855af13d6000803e8080156100e5573d6000f35b3d6000fd5b6100fd6100f8366004610908565b61015e565b005b6100fd610444565b61010f61014f565b60405161011c9190610a51565b60405180910390f35b6100fd6101333660046108e5565b6104de565b6100fd6101463660046108e5565b6105a4565b3390565b6033546001600160a01b031690565b61016661014b565b6001600160a01b031661017761014f565b6001600160a01b0316146101c0576040805162461bcd60e51b81526020600482018190526024820152600080516020610b0f833981519152604482015290519081900360640190fd5b60655460408051635890ef7960e01b815290516000926001600160a01b031691635890ef799160048083019286929190829003018186803b15801561020457600080fd5b505afa158015610218573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261024091908101906109a5565b905060005b815181101561032857600082828151811061025c57fe5b60200260200101519050600080600090505b85518110156102b55785818151811061028357fe5b60200260200101516001600160a01b0316836001600160a01b031614156102ad57600191506102b5565b60010161026e565b508061031e576065546040516340a141ff60e01b81526001600160a01b03909116906340a141ff906102eb908590600401610a51565b600060405180830381600087803b15801561030557600080fd5b505af1158015610319573d6000803e3d6000fd5b505050505b5050600101610245565b5060005b825181101561043f57600083828151811061034357fe5b602090810291909101015160655460405163facd743b60e01b81529192506001600160a01b03169063facd743b9061037f908490600401610a51565b60206040518083038186803b15801561039757600080fd5b505afa1580156103ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103cf9190610a31565b61043657606554604051632691c64760e11b81526001600160a01b0390911690634d238c8e90610403908490600401610a51565b600060405180830381600087803b15801561041d57600080fd5b505af1158015610431573d6000803e3d6000fd5b505050505b5060010161032c565b505050565b61044c61014b565b6001600160a01b031661045d61014f565b6001600160a01b0316146104a6576040805162461bcd60e51b81526020600482018190526024820152600080516020610b0f833981519152604482015290519081900360640190fd5b6033546040516000916001600160a01b031690600080516020610b2f833981519152908390a3603380546001600160a01b0319169055565b600054610100900460ff16806104f757506104f7610695565b80610505575060005460ff16155b6105405760405162461bcd60e51b815260040180806020018281038252602e815260200180610ae1602e913960400191505060405180910390fd5b600054610100900460ff1615801561056b576000805460ff1961ff0019909116610100171660011790555b606580546001600160a01b0319166001600160a01b03841617905561058e6106a6565b80156105a0576000805461ff00191690555b5050565b6105ac61014b565b6001600160a01b03166105bd61014f565b6001600160a01b031614610606576040805162461bcd60e51b81526020600482018190526024820152600080516020610b0f833981519152604482015290519081900360640190fd5b6001600160a01b03811661064b5760405162461bcd60e51b8152600401808060200182810382526026815260200180610abb6026913960400191505060405180910390fd5b6033546040516001600160a01b03808416921690600080516020610b2f83398151915290600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b60006106a030610758565b15905090565b600054610100900460ff16806106bf57506106bf610695565b806106cd575060005460ff16155b6107085760405162461bcd60e51b815260040180806020018281038252602e815260200180610ae1602e913960400191505060405180910390fd5b600054610100900460ff16158015610733576000805460ff1961ff0019909116610100171660011790555b61073b61075e565b6107436107fe565b8015610755576000805461ff00191690555b50565b3b151590565b600054610100900460ff16806107775750610777610695565b80610785575060005460ff16155b6107c05760405162461bcd60e51b815260040180806020018281038252602e815260200180610ae1602e913960400191505060405180910390fd5b600054610100900460ff16158015610743576000805460ff1961ff0019909116610100171660011790558015610755576000805461ff001916905550565b600054610100900460ff16806108175750610817610695565b80610825575060005460ff16155b6108605760405162461bcd60e51b815260040180806020018281038252602e815260200180610ae1602e913960400191505060405180910390fd5b600054610100900460ff1615801561088b576000805460ff1961ff0019909116610100171660011790555b600061089561014b565b603380546001600160a01b0319166001600160a01b03831690811790915560405191925090600090600080516020610b2f833981519152908290a3508015610755576000805461ff001916905550565b6000602082840312156108f6578081fd5b813561090181610aa5565b9392505050565b6000602080838503121561091a578182fd5b82356001600160401b0381111561092f578283fd5b8301601f8101851361093f578283fd5b803561095261094d82610a88565b610a65565b818152838101908385018584028501860189101561096e578687fd5b8694505b8385101561099957803561098581610aa5565b835260019490940193918501918501610972565b50979650505050505050565b600060208083850312156109b7578182fd5b82516001600160401b038111156109cc578283fd5b8301601f810185136109dc578283fd5b80516109ea61094d82610a88565b8181528381019083850185840285018601891015610a06578687fd5b8694505b83851015610999578051610a1d81610aa5565b835260019490940193918501918501610a0a565b600060208284031215610a42578081fd5b81518015158114610901578182fd5b6001600160a01b0391909116815260200190565b6040518181016001600160401b0381118282101715610a8057fe5b604052919050565b60006001600160401b03821115610a9b57fe5b5060209081020190565b6001600160a01b038116811461075557600080fdfe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a65644f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a264697066735822122066f2ba5741036e007fdcd2cdb4cd78323a05df44da53edf9c83cb12a6b7595df64736f6c63430007060033",
}

// ForeignChainTokenBridgeAdminProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ForeignChainTokenBridgeAdminProxyMetaData.ABI instead.
var ForeignChainTokenBridgeAdminProxyABI = ForeignChainTokenBridgeAdminProxyMetaData.ABI

// ForeignChainTokenBridgeAdminProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ForeignChainTokenBridgeAdminProxyMetaData.Bin instead.
var ForeignChainTokenBridgeAdminProxyBin = ForeignChainTokenBridgeAdminProxyMetaData.Bin

// DeployForeignChainTokenBridgeAdminProxy deploys a new Ethereum contract, binding an instance of ForeignChainTokenBridgeAdminProxy to it.
func DeployForeignChainTokenBridgeAdminProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ForeignChainTokenBridgeAdminProxy, error) {
	parsed, err := ForeignChainTokenBridgeAdminProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ForeignChainTokenBridgeAdminProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ForeignChainTokenBridgeAdminProxy{ForeignChainTokenBridgeAdminProxyCaller: ForeignChainTokenBridgeAdminProxyCaller{contract: contract}, ForeignChainTokenBridgeAdminProxyTransactor: ForeignChainTokenBridgeAdminProxyTransactor{contract: contract}, ForeignChainTokenBridgeAdminProxyFilterer: ForeignChainTokenBridgeAdminProxyFilterer{contract: contract}}, nil
}

// ForeignChainTokenBridgeAdminProxy is an auto generated Go binding around an Ethereum contract.
type ForeignChainTokenBridgeAdminProxy struct {
	ForeignChainTokenBridgeAdminProxyCaller     // Read-only binding to the contract
	ForeignChainTokenBridgeAdminProxyTransactor // Write-only binding to the contract
	ForeignChainTokenBridgeAdminProxyFilterer   // Log filterer for contract events
}

// ForeignChainTokenBridgeAdminProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForeignChainTokenBridgeAdminProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForeignChainTokenBridgeAdminProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForeignChainTokenBridgeAdminProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForeignChainTokenBridgeAdminProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForeignChainTokenBridgeAdminProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForeignChainTokenBridgeAdminProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForeignChainTokenBridgeAdminProxySession struct {
	Contract     *ForeignChainTokenBridgeAdminProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                      // Call options to use throughout this session
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ForeignChainTokenBridgeAdminProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForeignChainTokenBridgeAdminProxyCallerSession struct {
	Contract *ForeignChainTokenBridgeAdminProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                            // Call options to use throughout this session
}

// ForeignChainTokenBridgeAdminProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForeignChainTokenBridgeAdminProxyTransactorSession struct {
	Contract     *ForeignChainTokenBridgeAdminProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                            // Transaction auth options to use throughout this session
}

// ForeignChainTokenBridgeAdminProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForeignChainTokenBridgeAdminProxyRaw struct {
	Contract *ForeignChainTokenBridgeAdminProxy // Generic contract binding to access the raw methods on
}

// ForeignChainTokenBridgeAdminProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForeignChainTokenBridgeAdminProxyCallerRaw struct {
	Contract *ForeignChainTokenBridgeAdminProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ForeignChainTokenBridgeAdminProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForeignChainTokenBridgeAdminProxyTransactorRaw struct {
	Contract *ForeignChainTokenBridgeAdminProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForeignChainTokenBridgeAdminProxy creates a new instance of ForeignChainTokenBridgeAdminProxy, bound to a specific deployed contract.
func NewForeignChainTokenBridgeAdminProxy(address common.Address, backend bind.ContractBackend) (*ForeignChainTokenBridgeAdminProxy, error) {
	contract, err := bindForeignChainTokenBridgeAdminProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxy{ForeignChainTokenBridgeAdminProxyCaller: ForeignChainTokenBridgeAdminProxyCaller{contract: contract}, ForeignChainTokenBridgeAdminProxyTransactor: ForeignChainTokenBridgeAdminProxyTransactor{contract: contract}, ForeignChainTokenBridgeAdminProxyFilterer: ForeignChainTokenBridgeAdminProxyFilterer{contract: contract}}, nil
}

// NewForeignChainTokenBridgeAdminProxyCaller creates a new read-only instance of ForeignChainTokenBridgeAdminProxy, bound to a specific deployed contract.
func NewForeignChainTokenBridgeAdminProxyCaller(address common.Address, caller bind.ContractCaller) (*ForeignChainTokenBridgeAdminProxyCaller, error) {
	contract, err := bindForeignChainTokenBridgeAdminProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxyCaller{contract: contract}, nil
}

// NewForeignChainTokenBridgeAdminProxyTransactor creates a new write-only instance of ForeignChainTokenBridgeAdminProxy, bound to a specific deployed contract.
func NewForeignChainTokenBridgeAdminProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ForeignChainTokenBridgeAdminProxyTransactor, error) {
	contract, err := bindForeignChainTokenBridgeAdminProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxyTransactor{contract: contract}, nil
}

// NewForeignChainTokenBridgeAdminProxyFilterer creates a new log filterer instance of ForeignChainTokenBridgeAdminProxy, bound to a specific deployed contract.
func NewForeignChainTokenBridgeAdminProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ForeignChainTokenBridgeAdminProxyFilterer, error) {
	contract, err := bindForeignChainTokenBridgeAdminProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxyFilterer{contract: contract}, nil
}

// bindForeignChainTokenBridgeAdminProxy binds a generic wrapper to an already deployed contract.
func bindForeignChainTokenBridgeAdminProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ForeignChainTokenBridgeAdminProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ForeignChainTokenBridgeAdminProxy.Contract.ForeignChainTokenBridgeAdminProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.ForeignChainTokenBridgeAdminProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.ForeignChainTokenBridgeAdminProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ForeignChainTokenBridgeAdminProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ForeignChainTokenBridgeAdminProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxySession) Owner() (common.Address, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.Owner(&_ForeignChainTokenBridgeAdminProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyCallerSession) Owner() (common.Address, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.Owner(&_ForeignChainTokenBridgeAdminProxy.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridgeValidators) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactor) Initialize(opts *bind.TransactOpts, _bridgeValidators common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.contract.Transact(opts, "initialize", _bridgeValidators)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridgeValidators) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxySession) Initialize(_bridgeValidators common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.Initialize(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, _bridgeValidators)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridgeValidators) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorSession) Initialize(_bridgeValidators common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.Initialize(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, _bridgeValidators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.RenounceOwnership(&_ForeignChainTokenBridgeAdminProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.RenounceOwnership(&_ForeignChainTokenBridgeAdminProxy.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.TransferOwnership(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.TransferOwnership(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, newOwner)
}

// UpdateTokenbridgeValidators is a paid mutator transaction binding the contract method 0x278dcd62.
//
// Solidity: function updateTokenbridgeValidators(address[] _newValidatorsList) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactor) UpdateTokenbridgeValidators(opts *bind.TransactOpts, _newValidatorsList []common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.contract.Transact(opts, "updateTokenbridgeValidators", _newValidatorsList)
}

// UpdateTokenbridgeValidators is a paid mutator transaction binding the contract method 0x278dcd62.
//
// Solidity: function updateTokenbridgeValidators(address[] _newValidatorsList) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxySession) UpdateTokenbridgeValidators(_newValidatorsList []common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.UpdateTokenbridgeValidators(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, _newValidatorsList)
}

// UpdateTokenbridgeValidators is a paid mutator transaction binding the contract method 0x278dcd62.
//
// Solidity: function updateTokenbridgeValidators(address[] _newValidatorsList) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorSession) UpdateTokenbridgeValidators(_newValidatorsList []common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.UpdateTokenbridgeValidators(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, _newValidatorsList)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.Fallback(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.Fallback(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, calldata)
}

// ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ForeignChainTokenBridgeAdminProxy contract.
type ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator struct {
	Event *ForeignChainTokenBridgeAdminProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForeignChainTokenBridgeAdminProxyOwnershipTransferred)
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
		it.Event = new(ForeignChainTokenBridgeAdminProxyOwnershipTransferred)
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
func (it *ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForeignChainTokenBridgeAdminProxyOwnershipTransferred represents a OwnershipTransferred event raised by the ForeignChainTokenBridgeAdminProxy contract.
type ForeignChainTokenBridgeAdminProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ForeignChainTokenBridgeAdminProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator{contract: _ForeignChainTokenBridgeAdminProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ForeignChainTokenBridgeAdminProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ForeignChainTokenBridgeAdminProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForeignChainTokenBridgeAdminProxyOwnershipTransferred)
				if err := _ForeignChainTokenBridgeAdminProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) ParseOwnershipTransferred(log types.Log) (*ForeignChainTokenBridgeAdminProxyOwnershipTransferred, error) {
	event := new(ForeignChainTokenBridgeAdminProxyOwnershipTransferred)
	if err := _ForeignChainTokenBridgeAdminProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForeignChainTokenBridgeAdminProxyValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the ForeignChainTokenBridgeAdminProxy contract.
type ForeignChainTokenBridgeAdminProxyValidatorAddedIterator struct {
	Event *ForeignChainTokenBridgeAdminProxyValidatorAdded // Event containing the contract specifics and raw log

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
func (it *ForeignChainTokenBridgeAdminProxyValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForeignChainTokenBridgeAdminProxyValidatorAdded)
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
		it.Event = new(ForeignChainTokenBridgeAdminProxyValidatorAdded)
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
func (it *ForeignChainTokenBridgeAdminProxyValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForeignChainTokenBridgeAdminProxyValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForeignChainTokenBridgeAdminProxyValidatorAdded represents a ValidatorAdded event raised by the ForeignChainTokenBridgeAdminProxy contract.
type ForeignChainTokenBridgeAdminProxyValidatorAdded struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed _validator)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) FilterValidatorAdded(opts *bind.FilterOpts, _validator []common.Address) (*ForeignChainTokenBridgeAdminProxyValidatorAddedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _ForeignChainTokenBridgeAdminProxy.contract.FilterLogs(opts, "ValidatorAdded", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxyValidatorAddedIterator{contract: _ForeignChainTokenBridgeAdminProxy.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed _validator)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *ForeignChainTokenBridgeAdminProxyValidatorAdded, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _ForeignChainTokenBridgeAdminProxy.contract.WatchLogs(opts, "ValidatorAdded", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForeignChainTokenBridgeAdminProxyValidatorAdded)
				if err := _ForeignChainTokenBridgeAdminProxy.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) ParseValidatorAdded(log types.Log) (*ForeignChainTokenBridgeAdminProxyValidatorAdded, error) {
	event := new(ForeignChainTokenBridgeAdminProxyValidatorAdded)
	if err := _ForeignChainTokenBridgeAdminProxy.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the ForeignChainTokenBridgeAdminProxy contract.
type ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator struct {
	Event *ForeignChainTokenBridgeAdminProxyValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForeignChainTokenBridgeAdminProxyValidatorRemoved)
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
		it.Event = new(ForeignChainTokenBridgeAdminProxyValidatorRemoved)
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
func (it *ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForeignChainTokenBridgeAdminProxyValidatorRemoved represents a ValidatorRemoved event raised by the ForeignChainTokenBridgeAdminProxy contract.
type ForeignChainTokenBridgeAdminProxyValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed _validator)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, _validator []common.Address) (*ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _ForeignChainTokenBridgeAdminProxy.contract.FilterLogs(opts, "ValidatorRemoved", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator{contract: _ForeignChainTokenBridgeAdminProxy.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed _validator)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *ForeignChainTokenBridgeAdminProxyValidatorRemoved, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _ForeignChainTokenBridgeAdminProxy.contract.WatchLogs(opts, "ValidatorRemoved", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForeignChainTokenBridgeAdminProxyValidatorRemoved)
				if err := _ForeignChainTokenBridgeAdminProxy.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) ParseValidatorRemoved(log types.Log) (*ForeignChainTokenBridgeAdminProxyValidatorRemoved, error) {
	event := new(ForeignChainTokenBridgeAdminProxyValidatorRemoved)
	if err := _ForeignChainTokenBridgeAdminProxy.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
