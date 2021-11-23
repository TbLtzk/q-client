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
	Bin: "0x608060405234801561001057600080fd5b50610e92806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a45760003560e01c80630a3b0a4f146100a957806329092d0e146100e357806352efea6e146101095780635dbe47e814610111578063715018a6146101375780638da5cb5b14610141578063949d225d146101655780639e9405c31461017f578063a224cee7146101a5578063a39fac1214610246578063c32d805a1461029e578063f2fde38b146102c4575b600080fd5b6100cf600480360360208110156100bf57600080fd5b50356001600160a01b03166102ea565b604080519115158252519081900360200190f35b6100cf600480360360208110156100f957600080fd5b50356001600160a01b0316610376565b6100cf6103f8565b6100cf6004803603602081101561012757600080fd5b50356001600160a01b03166104ba565b61013f6104d7565b005b610149610571565b604080516001600160a01b039092168252519081900360200190f35b61016d610580565b60408051918252519081900360200190f35b61013f6004803603602081101561019557600080fd5b50356001600160a01b0316610586565b61013f600480360360208110156101bb57600080fd5b810190602081018135600160201b8111156101d557600080fd5b8201836020820111156101e757600080fd5b803590602001918460208302840111600160201b8311171561020857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506105cd945050505050565b61024e610711565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561028a578181015183820152602001610272565b505050509050019250505060405180910390f35b61013f600480360360208110156102b457600080fd5b50356001600160a01b0316610773565b61013f600480360360208110156102da57600080fd5b50356001600160a01b03166107b7565b60006102f46108a8565b6001600160a01b0316610305610571565b6001600160a01b03161461034e576040805162461bcd60e51b81526020600482018190526024820152600080516020610e1d833981519152604482015290519081900360640190fd5b610357826104ba565b1561036457506000610371565b61036d826108ac565b5060015b919050565b60006103806108a8565b6001600160a01b0316610391610571565b6001600160a01b0316146103da576040805162461bcd60e51b81526020600482018190526024820152600080516020610e1d833981519152604482015290519081900360640190fd5b6103e3826104ba565b6103ef57506000610371565b61036d8261090f565b60006104026108a8565b6001600160a01b0316610413610571565b6001600160a01b03161461045c576040805162461bcd60e51b81526020600482018190526024820152600080516020610e1d833981519152604482015290519081900360640190fd5b60005b6066548110156104a757606560006066838154811061047a57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400181205560010161045f565b506104b460666000610ce2565b50600190565b6001600160a01b0316600090815260656020526040902054151590565b6104df6108a8565b6001600160a01b03166104f0610571565b6001600160a01b031614610539576040805162461bcd60e51b81526020600482018190526024820152600080516020610e1d833981519152604482015290519081900360640190fd5b6033546040516000916001600160a01b031690600080516020610e3d833981519152908390a3603380546001600160a01b0319169055565b6033546001600160a01b031690565b60665490565b61058f816102ea565b6105ca5760405162461bcd60e51b8152600401808060200182810382526071815260200180610d7e6071913960800191505060405180910390fd5b50565b600054610100900460ff16806105e657506105e6610a49565b806105f4575060005460ff16155b61062f5760405162461bcd60e51b815260040180806020018281038252602e815260200180610def602e913960400191505060405180910390fd5b600054610100900460ff1615801561065a576000805460ff1961ff0019909116610100171660011790555b60005b82518110156106f257606683828151811061067457fe5b6020908102919091018101518254600181018455600093845291832090910180546001600160a01b0319166001600160a01b03909216919091179055606654845190916065918690859081106106c657fe5b6020908102919091018101516001600160a01b031682528101919091526040016000205560010161065d565b506106fb610a5a565b801561070d576000805461ff00191690555b5050565b6060606680548060200260200160405190810160405280929190818152602001828054801561076957602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161074b575b5050505050905090565b61077c81610376565b6105ca5760405162461bcd60e51b8152600401808060200182810382526043815260200180610d3b6043913960600191505060405180910390fd5b6107bf6108a8565b6001600160a01b03166107d0610571565b6001600160a01b031614610819576040805162461bcd60e51b81526020600482018190526024820152600080516020610e1d833981519152604482015290519081900360640190fd5b6001600160a01b03811661085e5760405162461bcd60e51b8152600401808060200182810382526026815260200180610d156026913960400191505060405180910390fd5b6033546040516001600160a01b03808416921690600080516020610e3d83398151915290600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b606680546001810182557f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943540180546001600160a01b0319166001600160a01b0384169081179091559054600091825260656020526040909120556105ca81610b0b565b6001600160a01b038116600090815260656020908152604080832081519283019091525481526066805491926000198301929091908390811061094e57fe5b60009182526020909120015483516001600160a01b0390911691506000190182146109d25782516001600160a01b0382166000908152606560205260409020558251606680548392600019019081106109a357fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b60668054806109dd57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0386168252606590526040812055610a2184610b0b565b836001600160a01b0316816001600160a01b031614610a4357610a4381610b0b565b50505050565b6000610a5430610b55565b15905090565b600054610100900460ff1680610a735750610a73610a49565b80610a81575060005460ff16155b610abc5760405162461bcd60e51b815260040180806020018281038252602e815260200180610def602e913960400191505060405180910390fd5b600054610100900460ff16158015610ae7576000805460ff1961ff0019909116610100171660011790555b610aef610b5b565b610af7610bfb565b80156105ca576000805461ff001916905550565b6001600160a01b038116600090815260656020526040902054606654811115610b3057fe5b610b39826104ba565b15610b4d5760008111610b4857fe5b61070d565b801561070d57fe5b3b151590565b600054610100900460ff1680610b745750610b74610a49565b80610b82575060005460ff16155b610bbd5760405162461bcd60e51b815260040180806020018281038252602e815260200180610def602e913960400191505060405180910390fd5b600054610100900460ff16158015610af7576000805460ff1961ff00199091166101001716600117905580156105ca576000805461ff001916905550565b600054610100900460ff1680610c145750610c14610a49565b80610c22575060005460ff16155b610c5d5760405162461bcd60e51b815260040180806020018281038252602e815260200180610def602e913960400191505060405180910390fd5b600054610100900460ff16158015610c88576000805460ff1961ff0019909116610100171660011790555b6000610c926108a8565b603380546001600160a01b0319166001600160a01b03831690811790915560405191925090600090600080516020610e3d833981519152908290a35080156105ca576000805461ff001916905550565b50805460008255906000526020600020908101906105ca91905b80821115610d105760008155600101610cfc565b509056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573735b5145432d3033353030305d2d4661696c656420746f2072656d6f76652074686520616464726573732066726f6d2074686520616464726573732073746f726167652e5b5145432d3033353030325d2d54686520616464726573732068617320616c7265616479206265656e20616464656420746f207468652073746f726167652c206661696c656420746f2061646420746865206164647265737320746f2074686520616464726573732073746f726167652e496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a65644f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a2646970667358221220d7b4baffc7f477c8d47723baeffc50476e57d425fa2469934cb5b72c449aac4364736f6c63430007060033",
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
