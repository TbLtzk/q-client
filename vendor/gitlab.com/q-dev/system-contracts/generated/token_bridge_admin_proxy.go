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
	Bin: "0x608060405234801561001057600080fd5b50610e12806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063485cc955146100ea578063715018a6146100ff5780638da5cb5b14610107578063b40ba22714610125578063f2fde38b1461012d575b61005f610140565b6001600160a01b0316610070610144565b6001600160a01b0316146100b9576040805162461bcd60e51b81526020600482018190526024820152600080516020610d9d833981519152604482015290519081900360640190fd5b6066546001600160a01b0316366000803760008036600080855af13d6000803e8080156100e5573d6000f35b3d6000fd5b6100fd6100f8366004610bcb565b610153565b005b6100fd61022f565b61010f610144565b60405161011c9190610cf0565b60405180910390f35b6100fd6102c9565b6100fd61013b366004610b8c565b6102e1565b3390565b6033546001600160a01b031690565b600054610100900460ff168061016c575061016c6103d2565b8061017a575060005460ff16155b6101b55760405162461bcd60e51b815260040180806020018281038252602e815260200180610d6f602e913960400191505060405180910390fd5b600054610100900460ff161580156101e0576000805460ff1961ff0019909116610100171660011790555b606580546001600160a01b038086166001600160a01b03199283161790925560668054928516929091169190911790556102186103e3565b801561022a576000805461ff00191690555b505050565b610237610140565b6001600160a01b0316610248610144565b6001600160a01b031614610291576040805162461bcd60e51b81526020600482018190526024820152600080516020610d9d833981519152604482015290519081900360640190fd5b6033546040516000916001600160a01b031690600080516020610dbd833981519152908390a3603380546001600160a01b0319169055565b60006102d3610494565b90506102de816106c5565b50565b6102e9610140565b6001600160a01b03166102fa610144565b6001600160a01b031614610343576040805162461bcd60e51b81526020600482018190526024820152600080516020610d9d833981519152604482015290519081900360640190fd5b6001600160a01b0381166103885760405162461bcd60e51b8152600401808060200182810382526026815260200180610d496026913960400191505060405180910390fd5b6033546040516001600160a01b03808416921690600080516020610dbd83398151915290600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b60006103dd306109f0565b15905090565b600054610100900460ff16806103fc57506103fc6103d2565b8061040a575060005460ff16155b6104455760405162461bcd60e51b815260040180806020018281038252602e815260200180610d6f602e913960400191505060405180910390fd5b600054610100900460ff16158015610470576000805460ff1961ff0019909116610100171660011790555b6104786109fa565b610480610a9a565b80156102de576000805461ff001916905550565b606554604051633fb9027160e01b81526060916000916001600160a01b0390911690633fb90271906104c890600401610d04565b60206040518083038186803b1580156104e057600080fd5b505afa1580156104f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105189190610baf565b6001600160a01b0316633d6206286040518163ffffffff1660e01b815260040160006040518083038186803b15801561055057600080fd5b505afa158015610564573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261058c9190810190610c03565b90506000815190506000606660009054906101000a90046001600160a01b03166001600160a01b0316630f43a6776040518163ffffffff1660e01b815260040160206040518083038186803b1580156105e457600080fd5b505afa1580156105f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061c9190610cd8565b9050825181101561062b578091505b60008267ffffffffffffffff8111801561064457600080fd5b5060405190808252806020026020018201604052801561066e578160200160208202803683370190505b50905060005b838110156106bc5784818151811061068857fe5b602002602001015182828151811061069c57fe5b6001600160a01b0390921660209283029190910190910152600101610674565b50935050505090565b60665460408051635890ef7960e01b815290516000926001600160a01b031691635890ef799160048083019286929190829003018186803b15801561070957600080fd5b505afa15801561071d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107459190810190610c03565b90506000606660009054906101000a90046001600160a01b03166001600160a01b0316630f43a6776040518163ffffffff1660e01b815260040160206040518083038186803b15801561079757600080fd5b505afa1580156107ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107cf9190610cd8565b825190915060005b83518110156108c15760008482815181106107ee57fe5b60200260200101519050600080600090505b87518110156108475787818151811061081557fe5b60200260200101516001600160a01b0316836001600160a01b0316141561083f5760019150610847565b600101610800565b50806108b7576066546040516340a141ff60e01b8152600019909501946001600160a01b03909116906340a141ff90610884908590600401610cf0565b600060405180830381600087803b15801561089e57600080fd5b505af11580156108b2573d6000803e3d6000fd5b505050505b50506001016107d7565b5060005b8451811080156108d457508282105b156109e95760008582815181106108e757fe5b602090810291909101015160665460405163facd743b60e01b81529192506001600160a01b03169063facd743b90610923908490600401610cf0565b60206040518083038186803b15801561093b57600080fd5b505afa15801561094f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109739190610cb8565b6109e057606654604051632691c64760e11b81526001909401936001600160a01b0390911690634d238c8e906109ad908490600401610cf0565b600060405180830381600087803b1580156109c757600080fd5b505af11580156109db573d6000803e3d6000fd5b505050505b506001016108c5565b5050505050565b803b15155b919050565b600054610100900460ff1680610a135750610a136103d2565b80610a21575060005460ff16155b610a5c5760405162461bcd60e51b815260040180806020018281038252602e815260200180610d6f602e913960400191505060405180910390fd5b600054610100900460ff16158015610480576000805460ff1961ff00199091166101001716600117905580156102de576000805461ff001916905550565b600054610100900460ff1680610ab35750610ab36103d2565b80610ac1575060005460ff16155b610afc5760405162461bcd60e51b815260040180806020018281038252602e815260200180610d6f602e913960400191505060405180910390fd5b600054610100900460ff16158015610b27576000805460ff1961ff0019909116610100171660011790555b6000610b31610140565b603380546001600160a01b0319166001600160a01b03831690811790915560405191925090600090600080516020610dbd833981519152908290a35080156102de576000805461ff001916905550565b80516109f581610d33565b600060208284031215610b9d578081fd5b8135610ba881610d33565b9392505050565b600060208284031215610bc0578081fd5b8151610ba881610d33565b60008060408385031215610bdd578081fd5b8235610be881610d33565b91506020830135610bf881610d33565b809150509250929050565b60006020808385031215610c15578182fd5b825167ffffffffffffffff80821115610c2c578384fd5b818501915085601f830112610c3f578384fd5b815181811115610c4b57fe5b83810260405185828201018181108582111715610c6457fe5b604052828152858101935084860182860187018a1015610c82578788fd5b8795505b83861015610cab57610c9781610b81565b855260019590950194938601938601610c86565b5098975050505050505050565b600060208284031215610cc9578081fd5b81518015158114610ba8578182fd5b600060208284031215610ce9578081fd5b5051919050565b6001600160a01b0391909116815260200190565b602080825260159082015274676f7665726e616e63652e76616c696461746f727360581b604082015260600190565b6001600160a01b03811681146102de57600080fdfe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a65644f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a26469706673582212200dbcba2b27251b00b4e5409377ac42259d240d2af707638c36f2bf10fdc13b1264736f6c63430007060033",
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
