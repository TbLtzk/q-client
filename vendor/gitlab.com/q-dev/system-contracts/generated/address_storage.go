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

// AddressStorageMetaData contains all meta data concerning the AddressStorage contract.
var AddressStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrList\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"mustRemove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"mustAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clear\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610d47806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a45760003560e01c80630a3b0a4f146100a957806329092d0e146100d157806352efea6e146100e45780635dbe47e8146100ec578063715018a6146100ff5780638da5cb5b14610109578063949d225d146101295780639e9405c31461013a578063a224cee71461014d578063a39fac1214610160578063c32d805a14610175578063f2fde38b14610188575b600080fd5b6100bc6100b7366004610aba565b61019b565b60405190151581526020015b60405180910390f35b6100bc6100df366004610aba565b6101fd565b6100bc61024c565b6100bc6100fa366004610aba565b6102eb565b610107610308565b005b610111610343565b6040516001600160a01b0390911681526020016100c8565b6066546040519081526020016100c8565b610107610148366004610aba565b610352565b61010761015b366004610af2565b610410565b610168610534565b6040516100c89190610bb7565b610107610183366004610aba565b610596565b610107610196366004610aba565b61061d565b6000336101a6610343565b6001600160a01b0316146101d55760405162461bcd60e51b81526004016101cc90610c04565b60405180910390fd5b6101de826102eb565b156101eb57506000919050565b6101f4826106ba565b5060015b919050565b600033610208610343565b6001600160a01b03161461022e5760405162461bcd60e51b81526004016101cc90610c04565b610237826102eb565b61024357506000919050565b6101f48261071d565b600033610257610343565b6001600160a01b03161461027d5760405162461bcd60e51b81526004016101cc90610c04565b60005b6066548110156102d85760656000606683815481106102a1576102a1610c39565b60009182526020808320909101546001600160a01b03168352820192909252604001812055806102d081610c65565b915050610280565b506102e560666000610a71565b50600190565b6001600160a01b0316600090815260656020526040902054151590565b33610311610343565b6001600160a01b0316146103375760405162461bcd60e51b81526004016101cc90610c04565b6103416000610883565b565b6033546001600160a01b031690565b61035b8161019b565b61040d5760405162461bcd60e51b815260206004820152607160248201527f5b5145432d3033353030325d2d54686520616464726573732068617320616c7260448201527f65616479206265656e20616464656420746f207468652073746f726167652c2060648201527f6661696c656420746f2061646420746865206164647265737320746f207468656084820152701030b2323932b9b99039ba37b930b3b29760791b60a482015260c4016101cc565b50565b600054610100900460ff1680610429575060005460ff16155b6104455760405162461bcd60e51b81526004016101cc90610c80565b600054610100900460ff16158015610467576000805461ffff19166101011790555b60005b825181101561051557606683828151811061048757610487610c39565b6020908102919091018101518254600181018455600093845291832090910180546001600160a01b0319166001600160a01b03909216919091179055606654845190916065918690859081106104df576104df610c39565b6020908102919091018101516001600160a01b03168252810191909152604001600020558061050d81610c65565b91505061046a565b5061051e6108d5565b8015610530576000805461ff00191690555b5050565b6060606680548060200260200160405190810160405280929190818152602001828054801561058c57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161056e575b5050505050905090565b61059f816101fd565b61040d5760405162461bcd60e51b815260206004820152604360248201527f5b5145432d3033353030305d2d4661696c656420746f2072656d6f766520746860448201527f6520616464726573732066726f6d2074686520616464726573732073746f726160648201526233b29760e91b608482015260a4016101cc565b33610626610343565b6001600160a01b03161461064c5760405162461bcd60e51b81526004016101cc90610c04565b6001600160a01b0381166106b15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101cc565b61040d81610883565b606680546001810182557f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943540180546001600160a01b0319166001600160a01b03841690811790915590546000918252606560205260409091205561040d81610950565b6001600160a01b0381166000908152606560209081526040808320815192830190915254815260665490919061075590600190610cce565b905060006066828154811061076c5761076c610c39565b60009182526020909120015483516001600160a01b03909116915061079390600190610cce565b82146108065782516001600160a01b038216600090815260656020526040902055825181906066906107c790600190610cce565b815481106107d7576107d7610c39565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b606680548061081757610817610ce5565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038616825260659052604081205561085b84610950565b836001600160a01b0316816001600160a01b03161461087d5761087d81610950565b50505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16806108ee575060005460ff16155b61090a5760405162461bcd60e51b81526004016101cc90610c80565b600054610100900460ff1615801561092c576000805461ffff19166101011790555b6109346109a7565b61093c610a11565b801561040d576000805461ff001916905550565b6001600160a01b03811660009081526065602052604090205460665481111561097b5761097b610cfb565b610984826102eb565b15610999576000811161053057610530610cfb565b801561053057610530610cfb565b600054610100900460ff16806109c0575060005460ff16155b6109dc5760405162461bcd60e51b81526004016101cc90610c80565b600054610100900460ff1615801561093c576000805461ffff1916610101179055801561040d576000805461ff001916905550565b600054610100900460ff1680610a2a575060005460ff16155b610a465760405162461bcd60e51b81526004016101cc90610c80565b600054610100900460ff16158015610a68576000805461ffff19166101011790555b61093c33610883565b508054600082559060005260206000209081019061040d91905b80821115610a9f5760008155600101610a8b565b5090565b80356001600160a01b03811681146101f857600080fd5b600060208284031215610acc57600080fd5b610ad582610aa3565b9392505050565b634e487b7160e01b600052604160045260246000fd5b60006020808385031215610b0557600080fd5b823567ffffffffffffffff80821115610b1d57600080fd5b818501915085601f830112610b3157600080fd5b813581811115610b4357610b43610adc565b8060051b604051601f19603f83011681018181108582111715610b6857610b68610adc565b604052918252848201925083810185019188831115610b8657600080fd5b938501935b82851015610bab57610b9c85610aa3565b84529385019392850192610b8b565b98975050505050505050565b6020808252825182820181905260009190848201906040850190845b81811015610bf85783516001600160a01b031683529284019291840191600101610bd3565b50909695505050505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415610c7957610c79610c4f565b5060010190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b600082821015610ce057610ce0610c4f565b500390565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052600160045260246000fdfea2646970667358221220977078b179f0297982ef3002bc7834252cf2dcba52f17a39a6edce4e4ceafb3b64736f6c63430008090033",
}

// AddressStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressStorageMetaData.ABI instead.
var AddressStorageABI = AddressStorageMetaData.ABI

// AddressStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressStorageMetaData.Bin instead.
var AddressStorageBin = AddressStorageMetaData.Bin

// DeployAddressStorage deploys a new Ethereum contract, binding an instance of AddressStorage to it.
func DeployAddressStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressStorage, error) {
	parsed, err := AddressStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressStorage{AddressStorageCaller: AddressStorageCaller{contract: contract}, AddressStorageTransactor: AddressStorageTransactor{contract: contract}, AddressStorageFilterer: AddressStorageFilterer{contract: contract}}, nil
}

// AddressStorage is an auto generated Go binding around an Ethereum contract.
type AddressStorage struct {
	AddressStorageCaller     // Read-only binding to the contract
	AddressStorageTransactor // Write-only binding to the contract
	AddressStorageFilterer   // Log filterer for contract events
}

// AddressStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressStorageSession struct {
	Contract     *AddressStorage   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressStorageCallerSession struct {
	Contract *AddressStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AddressStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressStorageTransactorSession struct {
	Contract     *AddressStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AddressStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressStorageRaw struct {
	Contract *AddressStorage // Generic contract binding to access the raw methods on
}

// AddressStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressStorageCallerRaw struct {
	Contract *AddressStorageCaller // Generic read-only contract binding to access the raw methods on
}

// AddressStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressStorageTransactorRaw struct {
	Contract *AddressStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressStorage creates a new instance of AddressStorage, bound to a specific deployed contract.
func NewAddressStorage(address common.Address, backend bind.ContractBackend) (*AddressStorage, error) {
	contract, err := bindAddressStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressStorage{AddressStorageCaller: AddressStorageCaller{contract: contract}, AddressStorageTransactor: AddressStorageTransactor{contract: contract}, AddressStorageFilterer: AddressStorageFilterer{contract: contract}}, nil
}

// NewAddressStorageCaller creates a new read-only instance of AddressStorage, bound to a specific deployed contract.
func NewAddressStorageCaller(address common.Address, caller bind.ContractCaller) (*AddressStorageCaller, error) {
	contract, err := bindAddressStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressStorageCaller{contract: contract}, nil
}

// NewAddressStorageTransactor creates a new write-only instance of AddressStorage, bound to a specific deployed contract.
func NewAddressStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressStorageTransactor, error) {
	contract, err := bindAddressStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressStorageTransactor{contract: contract}, nil
}

// NewAddressStorageFilterer creates a new log filterer instance of AddressStorage, bound to a specific deployed contract.
func NewAddressStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressStorageFilterer, error) {
	contract, err := bindAddressStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressStorageFilterer{contract: contract}, nil
}

// bindAddressStorage binds a generic wrapper to an already deployed contract.
func bindAddressStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressStorage *AddressStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressStorage.Contract.AddressStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressStorage *AddressStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorage.Contract.AddressStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressStorage *AddressStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressStorage.Contract.AddressStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressStorage *AddressStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressStorage *AddressStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressStorage *AddressStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressStorage.Contract.contract.Transact(opts, method, params...)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorage *AddressStorageCaller) Contains(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _AddressStorage.contract.Call(opts, &out, "contains", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorage *AddressStorageSession) Contains(_addr common.Address) (bool, error) {
	return _AddressStorage.Contract.Contains(&_AddressStorage.CallOpts, _addr)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorage *AddressStorageCallerSession) Contains(_addr common.Address) (bool, error) {
	return _AddressStorage.Contract.Contains(&_AddressStorage.CallOpts, _addr)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorage *AddressStorageCaller) GetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AddressStorage.contract.Call(opts, &out, "getAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorage *AddressStorageSession) GetAddresses() ([]common.Address, error) {
	return _AddressStorage.Contract.GetAddresses(&_AddressStorage.CallOpts)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorage *AddressStorageCallerSession) GetAddresses() ([]common.Address, error) {
	return _AddressStorage.Contract.GetAddresses(&_AddressStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorage *AddressStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorage *AddressStorageSession) Owner() (common.Address, error) {
	return _AddressStorage.Contract.Owner(&_AddressStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorage *AddressStorageCallerSession) Owner() (common.Address, error) {
	return _AddressStorage.Contract.Owner(&_AddressStorage.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_AddressStorage *AddressStorageCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AddressStorage.contract.Call(opts, &out, "size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_AddressStorage *AddressStorageSession) Size() (*big.Int, error) {
	return _AddressStorage.Contract.Size(&_AddressStorage.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_AddressStorage *AddressStorageCallerSession) Size() (*big.Int, error) {
	return _AddressStorage.Contract.Size(&_AddressStorage.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address _addr) returns(bool)
func (_AddressStorage *AddressStorageTransactor) Add(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "add", _addr)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address _addr) returns(bool)
func (_AddressStorage *AddressStorageSession) Add(_addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.Add(&_AddressStorage.TransactOpts, _addr)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address _addr) returns(bool)
func (_AddressStorage *AddressStorageTransactorSession) Add(_addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.Add(&_AddressStorage.TransactOpts, _addr)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_AddressStorage *AddressStorageTransactor) Clear(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "clear")
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_AddressStorage *AddressStorageSession) Clear() (*types.Transaction, error) {
	return _AddressStorage.Contract.Clear(&_AddressStorage.TransactOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_AddressStorage *AddressStorageTransactorSession) Clear() (*types.Transaction, error) {
	return _AddressStorage.Contract.Clear(&_AddressStorage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _addrList) returns()
func (_AddressStorage *AddressStorageTransactor) Initialize(opts *bind.TransactOpts, _addrList []common.Address) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "initialize", _addrList)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _addrList) returns()
func (_AddressStorage *AddressStorageSession) Initialize(_addrList []common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.Initialize(&_AddressStorage.TransactOpts, _addrList)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _addrList) returns()
func (_AddressStorage *AddressStorageTransactorSession) Initialize(_addrList []common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.Initialize(&_AddressStorage.TransactOpts, _addrList)
}

// MustAdd is a paid mutator transaction binding the contract method 0x9e9405c3.
//
// Solidity: function mustAdd(address _addr) returns()
func (_AddressStorage *AddressStorageTransactor) MustAdd(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "mustAdd", _addr)
}

// MustAdd is a paid mutator transaction binding the contract method 0x9e9405c3.
//
// Solidity: function mustAdd(address _addr) returns()
func (_AddressStorage *AddressStorageSession) MustAdd(_addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.MustAdd(&_AddressStorage.TransactOpts, _addr)
}

// MustAdd is a paid mutator transaction binding the contract method 0x9e9405c3.
//
// Solidity: function mustAdd(address _addr) returns()
func (_AddressStorage *AddressStorageTransactorSession) MustAdd(_addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.MustAdd(&_AddressStorage.TransactOpts, _addr)
}

// MustRemove is a paid mutator transaction binding the contract method 0xc32d805a.
//
// Solidity: function mustRemove(address _addr) returns()
func (_AddressStorage *AddressStorageTransactor) MustRemove(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "mustRemove", _addr)
}

// MustRemove is a paid mutator transaction binding the contract method 0xc32d805a.
//
// Solidity: function mustRemove(address _addr) returns()
func (_AddressStorage *AddressStorageSession) MustRemove(_addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.MustRemove(&_AddressStorage.TransactOpts, _addr)
}

// MustRemove is a paid mutator transaction binding the contract method 0xc32d805a.
//
// Solidity: function mustRemove(address _addr) returns()
func (_AddressStorage *AddressStorageTransactorSession) MustRemove(_addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.MustRemove(&_AddressStorage.TransactOpts, _addr)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address _addr) returns(bool)
func (_AddressStorage *AddressStorageTransactor) Remove(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "remove", _addr)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address _addr) returns(bool)
func (_AddressStorage *AddressStorageSession) Remove(_addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.Remove(&_AddressStorage.TransactOpts, _addr)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address _addr) returns(bool)
func (_AddressStorage *AddressStorageTransactorSession) Remove(_addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.Remove(&_AddressStorage.TransactOpts, _addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorage *AddressStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorage *AddressStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressStorage.Contract.RenounceOwnership(&_AddressStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorage *AddressStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressStorage.Contract.RenounceOwnership(&_AddressStorage.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorage *AddressStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorage *AddressStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.TransferOwnership(&_AddressStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorage *AddressStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.TransferOwnership(&_AddressStorage.TransactOpts, newOwner)
}

// AddressStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddressStorage contract.
type AddressStorageOwnershipTransferredIterator struct {
	Event *AddressStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AddressStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageOwnershipTransferred)
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
		it.Event = new(AddressStorageOwnershipTransferred)
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
func (it *AddressStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageOwnershipTransferred represents a OwnershipTransferred event raised by the AddressStorage contract.
type AddressStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressStorage *AddressStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddressStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageOwnershipTransferredIterator{contract: _AddressStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressStorage *AddressStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddressStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageOwnershipTransferred)
				if err := _AddressStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AddressStorage *AddressStorageFilterer) ParseOwnershipTransferred(log types.Log) (*AddressStorageOwnershipTransferred, error) {
	event := new(AddressStorageOwnershipTransferred)
	if err := _AddressStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
