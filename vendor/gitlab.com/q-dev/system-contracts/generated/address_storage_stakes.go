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

// AddressStorageStakesMetaData contains all meta data concerning the AddressStorageStakes contract.
var AddressStorageStakesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"AddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"AddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"DelegatedStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"StakeDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"StakeIncreased\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addrList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addrStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedStake\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_delegatedStake\",\"type\":\"uint256\"}],\"name\":\"setDelegated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b610e578061007d6000396000f3fe608060405234801561001057600080fd5b50600436106100995760003560e01c806326ffee081461009e5780635dbe47e8146100de5780636467b40414610104578063715018a614610130578063804025641461013a5780638da5cb5b14610173578063949d225d1461017b578063a39fac1214610195578063d9034503146101ed578063f2fde38b14610231578063f5d82b6b14610257575b600080fd5b6100ca600480360360408110156100b457600080fd5b506001600160a01b038135169060200135610283565b604080519115158252519081900360200190f35b6100ca600480360360208110156100f457600080fd5b50356001600160a01b0316610413565b6100ca6004803603604081101561011a57600080fd5b506001600160a01b038135169060200135610458565b610138610548565b005b6101576004803603602081101561015057600080fd5b50356105ea565b604080516001600160a01b039092168252519081900360200190f35b610157610614565b610183610623565b60408051918252519081900360200190f35b61019d610629565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101d95781810151838201526020016101c1565b505050509050019250505060405180910390f35b6102136004803603602081101561020357600080fd5b50356001600160a01b031661068b565b60408051938452602084019290925282820152519081900360600190f35b6101386004803603602081101561024757600080fd5b50356001600160a01b03166106ab565b6100ca6004803603604081101561026d57600080fd5b506001600160a01b0381351690602001356107a3565b600061028d61088f565b6000546001600160a01b039081169116146102dd576040805162461bcd60e51b81526020600482018190526024820152600080516020610e02833981519152604482015290519081900360640190fd5b6001600160a01b03831660009081526001602081815260409283902083516060810185528154808252938201549281019290925260020154928101929092526103575760405162461bcd60e51b8152600401808060200182810382526046815260200180610dbc6046913960600191505060405180910390fd5b828160200151101561039a5760405162461bcd60e51b8152600401808060200182810382526045815260200180610d516045913960600191505060405180910390fd5b60208101516103a99084610893565b6001600160a01b038516600081815260016020818152604092839020909101939093558051868152905191927f700865370ffb2a65a2b0242e6a64b21ac907ed5ecd46c9cffc729c177b2b1c6992918290030190a2610407846108dc565b60019150505b92915050565b6001600160a01b03811660009081526001602081905260408220015415158061040d5750506001600160a01b0316600090815260016020526040902060020154151590565b600061046261088f565b6000546001600160a01b039081169116146104b2576040805162461bcd60e51b81526020600482018190526024820152600080516020610e02833981519152604482015290519081900360640190fd5b6001600160a01b0383166000908152600160205260409020541580156104d6575081155b156104e35750600161040d565b6001600160a01b038316600081815260016020908152604091829020600201859055815185815291517fc85c322a22d4a5cd3d9611287b2abf5bc9e6426d3ed5dc8d77017a7fe8fdc9109281900390910190a261053f836108dc565b50600192915050565b61055061088f565b6000546001600160a01b039081169116146105a0576040805162461bcd60e51b81526020600482018190526024820152600080516020610e02833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b600281815481106105fa57600080fd5b6000918252602090912001546001600160a01b0316905081565b6000546001600160a01b031690565b60025490565b6060600280548060200260200160405190810160405280929190818152602001828054801561068157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610663575b5050505050905090565b600160208190526000918252604090912080549181015460029091015483565b6106b361088f565b6000546001600160a01b03908116911614610703576040805162461bcd60e51b81526020600482018190526024820152600080516020610e02833981519152604482015290519081900360640190fd5b6001600160a01b0381166107485760405162461bcd60e51b8152600401808060200182810382526026815260200180610d966026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60006107ad61088f565b6000546001600160a01b039081169116146107fd576040805162461bcd60e51b81526020600482018190526024820152600080516020610e02833981519152604482015290519081900360640190fd5b8161080a5750600061040d565b6001600160a01b03831660009081526001602081905260409091200154610831908361093e565b6001600160a01b038416600081815260016020818152604092839020909101939093558051858152905191927f8b0ed825817a2e696c9a931715af4609fc60e1701f09c89ee7645130e937eb2d92918290030190a261053f836108dc565b3390565b60006108d583836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610996565b9392505050565b6001600160a01b0381166000908152600160208190526040909120015415801561091f57506001600160a01b038116600090815260016020526040902060020154155b156109325761092d81610a2d565b61093b565b61093b81610bbc565b50565b6000828201838110156108d5576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b60008184841115610a255760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156109ea5781810151838201526020016109d2565b50505050905090810190601f168015610a175780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6001600160a01b0381166000908152600160205260409020546002546000199182019101808214610af557816001016001600060028481548110610a6d57fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020556002805482908110610aa057fe5b600091825260209091200154600280546001600160a01b039092169184908110610ac657fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b6002805480610b0057fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03851682526001908190526040822082815590810182905560020155610b5283610c6f565b6040516001600160a01b038416907f24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb90600090a2808214610bb757610bb760028381548110610b9d57fe5b6000918252602090912001546001600160a01b0316610c6f565b505050565b6001600160a01b03811660009081526001602052604090205461093b576002805460018101825560009182527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180546001600160a01b0319166001600160a01b03841690811790915560405190917fa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f91a26002546001600160a01b03821660009081526001602052604090205561093b815b6001600160a01b038116600090815260016020526040902054600254811115610c9457fe5b610c9d82610413565b15610cfb5760008111610cac57fe5b6001600160a01b03821660009081526001602081905260409091200154151580610cf057506001600160a01b03821660009081526001602052604090206002015415155b610cf657fe5b610d4c565b8015610d0357fe5b6001600160a01b03821660009081526001602081905260409091200154158015610d4657506001600160a01b038216600090815260016020526040902060020154155b610d4c57fe5b505056fe5b5145432d3033363030325d2d546865207374616b6520746f2072656d6f76652069732067726561746572207468616e2074686520617661696c61626c65207374616b652e4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573735b5145432d3033363030305d2d546865206164647265737320646f6573206e6f742065786973742c206661696c656420746f20646563726561736520746865207374616b652e4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a264697066735822122031a1be38cf8db3c3e95e011d0097a175af169c12490018534fbf555751ba2dc064736f6c63430007060033",
}

// AddressStorageStakesABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressStorageStakesMetaData.ABI instead.
var AddressStorageStakesABI = AddressStorageStakesMetaData.ABI

// AddressStorageStakesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressStorageStakesMetaData.Bin instead.
var AddressStorageStakesBin = AddressStorageStakesMetaData.Bin

// DeployAddressStorageStakes deploys a new Ethereum contract, binding an instance of AddressStorageStakes to it.
func DeployAddressStorageStakes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressStorageStakes, error) {
	parsed, err := AddressStorageStakesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressStorageStakesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressStorageStakes{AddressStorageStakesCaller: AddressStorageStakesCaller{contract: contract}, AddressStorageStakesTransactor: AddressStorageStakesTransactor{contract: contract}, AddressStorageStakesFilterer: AddressStorageStakesFilterer{contract: contract}}, nil
}

// AddressStorageStakes is an auto generated Go binding around an Ethereum contract.
type AddressStorageStakes struct {
	AddressStorageStakesCaller     // Read-only binding to the contract
	AddressStorageStakesTransactor // Write-only binding to the contract
	AddressStorageStakesFilterer   // Log filterer for contract events
}

// AddressStorageStakesCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressStorageStakesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageStakesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressStorageStakesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageStakesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressStorageStakesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageStakesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressStorageStakesSession struct {
	Contract     *AddressStorageStakes // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AddressStorageStakesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressStorageStakesCallerSession struct {
	Contract *AddressStorageStakesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AddressStorageStakesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressStorageStakesTransactorSession struct {
	Contract     *AddressStorageStakesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AddressStorageStakesRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressStorageStakesRaw struct {
	Contract *AddressStorageStakes // Generic contract binding to access the raw methods on
}

// AddressStorageStakesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressStorageStakesCallerRaw struct {
	Contract *AddressStorageStakesCaller // Generic read-only contract binding to access the raw methods on
}

// AddressStorageStakesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressStorageStakesTransactorRaw struct {
	Contract *AddressStorageStakesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressStorageStakes creates a new instance of AddressStorageStakes, bound to a specific deployed contract.
func NewAddressStorageStakes(address common.Address, backend bind.ContractBackend) (*AddressStorageStakes, error) {
	contract, err := bindAddressStorageStakes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakes{AddressStorageStakesCaller: AddressStorageStakesCaller{contract: contract}, AddressStorageStakesTransactor: AddressStorageStakesTransactor{contract: contract}, AddressStorageStakesFilterer: AddressStorageStakesFilterer{contract: contract}}, nil
}

// NewAddressStorageStakesCaller creates a new read-only instance of AddressStorageStakes, bound to a specific deployed contract.
func NewAddressStorageStakesCaller(address common.Address, caller bind.ContractCaller) (*AddressStorageStakesCaller, error) {
	contract, err := bindAddressStorageStakes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesCaller{contract: contract}, nil
}

// NewAddressStorageStakesTransactor creates a new write-only instance of AddressStorageStakes, bound to a specific deployed contract.
func NewAddressStorageStakesTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressStorageStakesTransactor, error) {
	contract, err := bindAddressStorageStakes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesTransactor{contract: contract}, nil
}

// NewAddressStorageStakesFilterer creates a new log filterer instance of AddressStorageStakes, bound to a specific deployed contract.
func NewAddressStorageStakesFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressStorageStakesFilterer, error) {
	contract, err := bindAddressStorageStakes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesFilterer{contract: contract}, nil
}

// bindAddressStorageStakes binds a generic wrapper to an already deployed contract.
func bindAddressStorageStakes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressStorageStakesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressStorageStakes *AddressStorageStakesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressStorageStakes.Contract.AddressStorageStakesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressStorageStakes *AddressStorageStakesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.AddressStorageStakesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressStorageStakes *AddressStorageStakesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.AddressStorageStakesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressStorageStakes *AddressStorageStakesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressStorageStakes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressStorageStakes *AddressStorageStakesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressStorageStakes *AddressStorageStakesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.contract.Transact(opts, method, params...)
}

// AddrList is a free data retrieval call binding the contract method 0x80402564.
//
// Solidity: function addrList(uint256 ) view returns(address)
func (_AddressStorageStakes *AddressStorageStakesCaller) AddrList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AddressStorageStakes.contract.Call(opts, &out, "addrList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddrList is a free data retrieval call binding the contract method 0x80402564.
//
// Solidity: function addrList(uint256 ) view returns(address)
func (_AddressStorageStakes *AddressStorageStakesSession) AddrList(arg0 *big.Int) (common.Address, error) {
	return _AddressStorageStakes.Contract.AddrList(&_AddressStorageStakes.CallOpts, arg0)
}

// AddrList is a free data retrieval call binding the contract method 0x80402564.
//
// Solidity: function addrList(uint256 ) view returns(address)
func (_AddressStorageStakes *AddressStorageStakesCallerSession) AddrList(arg0 *big.Int) (common.Address, error) {
	return _AddressStorageStakes.Contract.AddrList(&_AddressStorageStakes.CallOpts, arg0)
}

// AddrStake is a free data retrieval call binding the contract method 0xd9034503.
//
// Solidity: function addrStake(address ) view returns(uint256 listIndex, uint256 stake, uint256 delegatedStake)
func (_AddressStorageStakes *AddressStorageStakesCaller) AddrStake(opts *bind.CallOpts, arg0 common.Address) (struct {
	ListIndex      *big.Int
	Stake          *big.Int
	DelegatedStake *big.Int
}, error) {
	var out []interface{}
	err := _AddressStorageStakes.contract.Call(opts, &out, "addrStake", arg0)

	outstruct := new(struct {
		ListIndex      *big.Int
		Stake          *big.Int
		DelegatedStake *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListIndex = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Stake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DelegatedStake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AddrStake is a free data retrieval call binding the contract method 0xd9034503.
//
// Solidity: function addrStake(address ) view returns(uint256 listIndex, uint256 stake, uint256 delegatedStake)
func (_AddressStorageStakes *AddressStorageStakesSession) AddrStake(arg0 common.Address) (struct {
	ListIndex      *big.Int
	Stake          *big.Int
	DelegatedStake *big.Int
}, error) {
	return _AddressStorageStakes.Contract.AddrStake(&_AddressStorageStakes.CallOpts, arg0)
}

// AddrStake is a free data retrieval call binding the contract method 0xd9034503.
//
// Solidity: function addrStake(address ) view returns(uint256 listIndex, uint256 stake, uint256 delegatedStake)
func (_AddressStorageStakes *AddressStorageStakesCallerSession) AddrStake(arg0 common.Address) (struct {
	ListIndex      *big.Int
	Stake          *big.Int
	DelegatedStake *big.Int
}, error) {
	return _AddressStorageStakes.Contract.AddrStake(&_AddressStorageStakes.CallOpts, arg0)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorageStakes *AddressStorageStakesCaller) Contains(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _AddressStorageStakes.contract.Call(opts, &out, "contains", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorageStakes *AddressStorageStakesSession) Contains(_addr common.Address) (bool, error) {
	return _AddressStorageStakes.Contract.Contains(&_AddressStorageStakes.CallOpts, _addr)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorageStakes *AddressStorageStakesCallerSession) Contains(_addr common.Address) (bool, error) {
	return _AddressStorageStakes.Contract.Contains(&_AddressStorageStakes.CallOpts, _addr)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorageStakes *AddressStorageStakesCaller) GetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AddressStorageStakes.contract.Call(opts, &out, "getAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorageStakes *AddressStorageStakesSession) GetAddresses() ([]common.Address, error) {
	return _AddressStorageStakes.Contract.GetAddresses(&_AddressStorageStakes.CallOpts)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorageStakes *AddressStorageStakesCallerSession) GetAddresses() ([]common.Address, error) {
	return _AddressStorageStakes.Contract.GetAddresses(&_AddressStorageStakes.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorageStakes *AddressStorageStakesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressStorageStakes.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorageStakes *AddressStorageStakesSession) Owner() (common.Address, error) {
	return _AddressStorageStakes.Contract.Owner(&_AddressStorageStakes.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorageStakes *AddressStorageStakesCallerSession) Owner() (common.Address, error) {
	return _AddressStorageStakes.Contract.Owner(&_AddressStorageStakes.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_AddressStorageStakes *AddressStorageStakesCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AddressStorageStakes.contract.Call(opts, &out, "size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_AddressStorageStakes *AddressStorageStakesSession) Size() (*big.Int, error) {
	return _AddressStorageStakes.Contract.Size(&_AddressStorageStakes.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_AddressStorageStakes *AddressStorageStakesCallerSession) Size() (*big.Int, error) {
	return _AddressStorageStakes.Contract.Size(&_AddressStorageStakes.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesTransactor) Add(opts *bind.TransactOpts, _addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.contract.Transact(opts, "add", _addr, _stake)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesSession) Add(_addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.Add(&_AddressStorageStakes.TransactOpts, _addr, _stake)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesTransactorSession) Add(_addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.Add(&_AddressStorageStakes.TransactOpts, _addr, _stake)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorageStakes *AddressStorageStakesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorageStakes.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorageStakes *AddressStorageStakesSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.RenounceOwnership(&_AddressStorageStakes.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorageStakes *AddressStorageStakesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.RenounceOwnership(&_AddressStorageStakes.TransactOpts)
}

// SetDelegated is a paid mutator transaction binding the contract method 0x6467b404.
//
// Solidity: function setDelegated(address _addr, uint256 _delegatedStake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesTransactor) SetDelegated(opts *bind.TransactOpts, _addr common.Address, _delegatedStake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.contract.Transact(opts, "setDelegated", _addr, _delegatedStake)
}

// SetDelegated is a paid mutator transaction binding the contract method 0x6467b404.
//
// Solidity: function setDelegated(address _addr, uint256 _delegatedStake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesSession) SetDelegated(_addr common.Address, _delegatedStake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.SetDelegated(&_AddressStorageStakes.TransactOpts, _addr, _delegatedStake)
}

// SetDelegated is a paid mutator transaction binding the contract method 0x6467b404.
//
// Solidity: function setDelegated(address _addr, uint256 _delegatedStake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesTransactorSession) SetDelegated(_addr common.Address, _delegatedStake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.SetDelegated(&_AddressStorageStakes.TransactOpts, _addr, _delegatedStake)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesTransactor) Sub(opts *bind.TransactOpts, _addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.contract.Transact(opts, "sub", _addr, _stake)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesSession) Sub(_addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.Sub(&_AddressStorageStakes.TransactOpts, _addr, _stake)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesTransactorSession) Sub(_addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.Sub(&_AddressStorageStakes.TransactOpts, _addr, _stake)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorageStakes *AddressStorageStakesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorageStakes.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorageStakes *AddressStorageStakesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.TransferOwnership(&_AddressStorageStakes.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorageStakes *AddressStorageStakesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.TransferOwnership(&_AddressStorageStakes.TransactOpts, newOwner)
}

// AddressStorageStakesAddressAddedIterator is returned from FilterAddressAdded and is used to iterate over the raw logs and unpacked data for AddressAdded events raised by the AddressStorageStakes contract.
type AddressStorageStakesAddressAddedIterator struct {
	Event *AddressStorageStakesAddressAdded // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesAddressAdded)
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
		it.Event = new(AddressStorageStakesAddressAdded)
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
func (it *AddressStorageStakesAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesAddressAdded represents a AddressAdded event raised by the AddressStorageStakes contract.
type AddressStorageStakesAddressAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddressAdded is a free log retrieval operation binding the contract event 0xa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f.
//
// Solidity: event AddressAdded(address indexed _addr)
func (_AddressStorageStakes *AddressStorageStakesFilterer) FilterAddressAdded(opts *bind.FilterOpts, _addr []common.Address) (*AddressStorageStakesAddressAddedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.FilterLogs(opts, "AddressAdded", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesAddressAddedIterator{contract: _AddressStorageStakes.contract, event: "AddressAdded", logs: logs, sub: sub}, nil
}

// WatchAddressAdded is a free log subscription operation binding the contract event 0xa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f.
//
// Solidity: event AddressAdded(address indexed _addr)
func (_AddressStorageStakes *AddressStorageStakesFilterer) WatchAddressAdded(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesAddressAdded, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.WatchLogs(opts, "AddressAdded", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesAddressAdded)
				if err := _AddressStorageStakes.contract.UnpackLog(event, "AddressAdded", log); err != nil {
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

// ParseAddressAdded is a log parse operation binding the contract event 0xa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f.
//
// Solidity: event AddressAdded(address indexed _addr)
func (_AddressStorageStakes *AddressStorageStakesFilterer) ParseAddressAdded(log types.Log) (*AddressStorageStakesAddressAdded, error) {
	event := new(AddressStorageStakesAddressAdded)
	if err := _AddressStorageStakes.contract.UnpackLog(event, "AddressAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageStakesAddressRemovedIterator is returned from FilterAddressRemoved and is used to iterate over the raw logs and unpacked data for AddressRemoved events raised by the AddressStorageStakes contract.
type AddressStorageStakesAddressRemovedIterator struct {
	Event *AddressStorageStakesAddressRemoved // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesAddressRemoved)
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
		it.Event = new(AddressStorageStakesAddressRemoved)
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
func (it *AddressStorageStakesAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesAddressRemoved represents a AddressRemoved event raised by the AddressStorageStakes contract.
type AddressStorageStakesAddressRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddressRemoved is a free log retrieval operation binding the contract event 0x24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb.
//
// Solidity: event AddressRemoved(address indexed _addr)
func (_AddressStorageStakes *AddressStorageStakesFilterer) FilterAddressRemoved(opts *bind.FilterOpts, _addr []common.Address) (*AddressStorageStakesAddressRemovedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.FilterLogs(opts, "AddressRemoved", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesAddressRemovedIterator{contract: _AddressStorageStakes.contract, event: "AddressRemoved", logs: logs, sub: sub}, nil
}

// WatchAddressRemoved is a free log subscription operation binding the contract event 0x24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb.
//
// Solidity: event AddressRemoved(address indexed _addr)
func (_AddressStorageStakes *AddressStorageStakesFilterer) WatchAddressRemoved(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesAddressRemoved, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.WatchLogs(opts, "AddressRemoved", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesAddressRemoved)
				if err := _AddressStorageStakes.contract.UnpackLog(event, "AddressRemoved", log); err != nil {
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

// ParseAddressRemoved is a log parse operation binding the contract event 0x24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb.
//
// Solidity: event AddressRemoved(address indexed _addr)
func (_AddressStorageStakes *AddressStorageStakesFilterer) ParseAddressRemoved(log types.Log) (*AddressStorageStakesAddressRemoved, error) {
	event := new(AddressStorageStakesAddressRemoved)
	if err := _AddressStorageStakes.contract.UnpackLog(event, "AddressRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageStakesDelegatedStakeUpdatedIterator is returned from FilterDelegatedStakeUpdated and is used to iterate over the raw logs and unpacked data for DelegatedStakeUpdated events raised by the AddressStorageStakes contract.
type AddressStorageStakesDelegatedStakeUpdatedIterator struct {
	Event *AddressStorageStakesDelegatedStakeUpdated // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesDelegatedStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesDelegatedStakeUpdated)
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
		it.Event = new(AddressStorageStakesDelegatedStakeUpdated)
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
func (it *AddressStorageStakesDelegatedStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesDelegatedStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesDelegatedStakeUpdated represents a DelegatedStakeUpdated event raised by the AddressStorageStakes contract.
type AddressStorageStakesDelegatedStakeUpdated struct {
	Addr  common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDelegatedStakeUpdated is a free log retrieval operation binding the contract event 0xc85c322a22d4a5cd3d9611287b2abf5bc9e6426d3ed5dc8d77017a7fe8fdc910.
//
// Solidity: event DelegatedStakeUpdated(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) FilterDelegatedStakeUpdated(opts *bind.FilterOpts, _addr []common.Address) (*AddressStorageStakesDelegatedStakeUpdatedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.FilterLogs(opts, "DelegatedStakeUpdated", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesDelegatedStakeUpdatedIterator{contract: _AddressStorageStakes.contract, event: "DelegatedStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegatedStakeUpdated is a free log subscription operation binding the contract event 0xc85c322a22d4a5cd3d9611287b2abf5bc9e6426d3ed5dc8d77017a7fe8fdc910.
//
// Solidity: event DelegatedStakeUpdated(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) WatchDelegatedStakeUpdated(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesDelegatedStakeUpdated, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.WatchLogs(opts, "DelegatedStakeUpdated", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesDelegatedStakeUpdated)
				if err := _AddressStorageStakes.contract.UnpackLog(event, "DelegatedStakeUpdated", log); err != nil {
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

// ParseDelegatedStakeUpdated is a log parse operation binding the contract event 0xc85c322a22d4a5cd3d9611287b2abf5bc9e6426d3ed5dc8d77017a7fe8fdc910.
//
// Solidity: event DelegatedStakeUpdated(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) ParseDelegatedStakeUpdated(log types.Log) (*AddressStorageStakesDelegatedStakeUpdated, error) {
	event := new(AddressStorageStakesDelegatedStakeUpdated)
	if err := _AddressStorageStakes.contract.UnpackLog(event, "DelegatedStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageStakesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddressStorageStakes contract.
type AddressStorageStakesOwnershipTransferredIterator struct {
	Event *AddressStorageStakesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesOwnershipTransferred)
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
		it.Event = new(AddressStorageStakesOwnershipTransferred)
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
func (it *AddressStorageStakesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesOwnershipTransferred represents a OwnershipTransferred event raised by the AddressStorageStakes contract.
type AddressStorageStakesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressStorageStakes *AddressStorageStakesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddressStorageStakesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesOwnershipTransferredIterator{contract: _AddressStorageStakes.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressStorageStakes *AddressStorageStakesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesOwnershipTransferred)
				if err := _AddressStorageStakes.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AddressStorageStakes *AddressStorageStakesFilterer) ParseOwnershipTransferred(log types.Log) (*AddressStorageStakesOwnershipTransferred, error) {
	event := new(AddressStorageStakesOwnershipTransferred)
	if err := _AddressStorageStakes.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageStakesStakeDecreasedIterator is returned from FilterStakeDecreased and is used to iterate over the raw logs and unpacked data for StakeDecreased events raised by the AddressStorageStakes contract.
type AddressStorageStakesStakeDecreasedIterator struct {
	Event *AddressStorageStakesStakeDecreased // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesStakeDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesStakeDecreased)
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
		it.Event = new(AddressStorageStakesStakeDecreased)
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
func (it *AddressStorageStakesStakeDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesStakeDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesStakeDecreased represents a StakeDecreased event raised by the AddressStorageStakes contract.
type AddressStorageStakesStakeDecreased struct {
	Addr  common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStakeDecreased is a free log retrieval operation binding the contract event 0x700865370ffb2a65a2b0242e6a64b21ac907ed5ecd46c9cffc729c177b2b1c69.
//
// Solidity: event StakeDecreased(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) FilterStakeDecreased(opts *bind.FilterOpts, _addr []common.Address) (*AddressStorageStakesStakeDecreasedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.FilterLogs(opts, "StakeDecreased", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesStakeDecreasedIterator{contract: _AddressStorageStakes.contract, event: "StakeDecreased", logs: logs, sub: sub}, nil
}

// WatchStakeDecreased is a free log subscription operation binding the contract event 0x700865370ffb2a65a2b0242e6a64b21ac907ed5ecd46c9cffc729c177b2b1c69.
//
// Solidity: event StakeDecreased(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) WatchStakeDecreased(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesStakeDecreased, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.WatchLogs(opts, "StakeDecreased", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesStakeDecreased)
				if err := _AddressStorageStakes.contract.UnpackLog(event, "StakeDecreased", log); err != nil {
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

// ParseStakeDecreased is a log parse operation binding the contract event 0x700865370ffb2a65a2b0242e6a64b21ac907ed5ecd46c9cffc729c177b2b1c69.
//
// Solidity: event StakeDecreased(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) ParseStakeDecreased(log types.Log) (*AddressStorageStakesStakeDecreased, error) {
	event := new(AddressStorageStakesStakeDecreased)
	if err := _AddressStorageStakes.contract.UnpackLog(event, "StakeDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageStakesStakeIncreasedIterator is returned from FilterStakeIncreased and is used to iterate over the raw logs and unpacked data for StakeIncreased events raised by the AddressStorageStakes contract.
type AddressStorageStakesStakeIncreasedIterator struct {
	Event *AddressStorageStakesStakeIncreased // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesStakeIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesStakeIncreased)
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
		it.Event = new(AddressStorageStakesStakeIncreased)
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
func (it *AddressStorageStakesStakeIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesStakeIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesStakeIncreased represents a StakeIncreased event raised by the AddressStorageStakes contract.
type AddressStorageStakesStakeIncreased struct {
	Addr  common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStakeIncreased is a free log retrieval operation binding the contract event 0x8b0ed825817a2e696c9a931715af4609fc60e1701f09c89ee7645130e937eb2d.
//
// Solidity: event StakeIncreased(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) FilterStakeIncreased(opts *bind.FilterOpts, _addr []common.Address) (*AddressStorageStakesStakeIncreasedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.FilterLogs(opts, "StakeIncreased", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesStakeIncreasedIterator{contract: _AddressStorageStakes.contract, event: "StakeIncreased", logs: logs, sub: sub}, nil
}

// WatchStakeIncreased is a free log subscription operation binding the contract event 0x8b0ed825817a2e696c9a931715af4609fc60e1701f09c89ee7645130e937eb2d.
//
// Solidity: event StakeIncreased(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) WatchStakeIncreased(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesStakeIncreased, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.WatchLogs(opts, "StakeIncreased", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesStakeIncreased)
				if err := _AddressStorageStakes.contract.UnpackLog(event, "StakeIncreased", log); err != nil {
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

// ParseStakeIncreased is a log parse operation binding the contract event 0x8b0ed825817a2e696c9a931715af4609fc60e1701f09c89ee7645130e937eb2d.
//
// Solidity: event StakeIncreased(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) ParseStakeIncreased(log types.Log) (*AddressStorageStakesStakeIncreased, error) {
	event := new(AddressStorageStakesStakeIncreased)
	if err := _AddressStorageStakes.contract.UnpackLog(event, "StakeIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
