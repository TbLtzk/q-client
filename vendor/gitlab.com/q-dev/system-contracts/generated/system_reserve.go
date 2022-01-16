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

// SystemReserveMetaData contains all meta data concerning the SystemReserve contract.
var SystemReserveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"availableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coolDownPhase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_keys\",\"type\":\"string[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setPauseState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEligibleContractKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061118d806100206000396000f3fe6080604052600436106100645760003560e01c806328011170146100705780632d250a5b146100995780632e1a7d4d146100bb57806391f7cfb9146100eb57806394e6f56f146101015780639d2f83f014610123578063cdb88ad11461013d57600080fd5b3661006b57005b600080fd5b34801561007c57600080fd5b5061008660035481565b6040519081526020015b60405180910390f35b3480156100a557600080fd5b506100ae61015d565b6040516100909190610d0e565b3480156100c757600080fd5b506100db6100d6366004610d70565b610236565b6040519015158152602001610090565b3480156100f757600080fd5b5061008660045481565b34801561010d57600080fd5b5061012161011c366004610de5565b61057a565b005b34801561012f57600080fd5b506002546100db9060ff1681565b34801561014957600080fd5b50610121610158366004610f20565b610685565b60606001805480602002602001604051908101604052809291908181526020016000905b8282101561022d5783829060005260206000200180546101a090610f44565b80601f01602080910402602001604051908101604052809291908181526020018280546101cc90610f44565b80156102195780601f106101ee57610100808354040283529160200191610219565b820191906000526020600020905b8154815290600101906020018083116101fc57829003601f168201915b505050505081526020019060010190610181565b50505050905090565b60008060005b600154811015610313576000546001805433926201000090046001600160a01b031691633fb90271918590811061027557610275610f79565b906000526020600020016040518263ffffffff1660e01b815260040161029b9190610f8f565b60206040518083038186803b1580156102b357600080fd5b505afa1580156102c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102eb9190611037565b6001600160a01b031614156103035760019150610313565b61030c8161106a565b905061023c565b508061039a5760405162461bcd60e51b815260206004820152604560248201527f5b5145432d3031383030305d2d5065726d697373696f6e2064656e696564202d60448201527f206f6e6c7920656c696769626c6520636f6e747261637473206861766520616360648201526431b2b9b99760d91b608482015260a4015b60405180910390fd5b60025460ff161561040a5760405162461bcd60e51b815260206004820152603460248201527f5b5145432d3031383030325d2d5468652073797374656d20726573657276652060448201527331b7b73a3930b1ba1034b9903830bab9b2b2171760611b6064820152608401610391565b610412610973565b8247108061041e575082155b1561042c5760009150610574565b8260045410156104a05760405162461bcd60e51b815260206004820152603960248201527f5b5145432d3031383030335d2d496e73756666696369656e742066756e64732060448201527830bb30b4b630b13632903337b9103bb4ba34323930bbb0b61760391b6064820152608401610391565b82600460008282546104b29190611085565b9091555050604051600090339085908381818185875af1925050503d80600081146104f9576040519150601f19603f3d011682016040523d82523d6000602084013e6104fe565b606091505b505090508061056e5760405162461bcd60e51b815260206004820152603660248201527f5b5145432d3031383030345d2d5472616e73666572206f66207468652077697460448201527534323930bbb0b61030b6b7bab73a103330b4b632b21760511b6064820152608401610391565b60019250505b50919050565b600054610100900460ff1680610593575060005460ff16155b6105f65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610391565b600054610100900460ff16158015610618576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b03861602179055815161064d906001906020850190610b75565b506106566109a0565b610660904261109c565b60035561066b610a56565b6004558015610680576000805461ff00191690555b505050565b600060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060238152602001611135602391396040518263ffffffff1660e01b81526004016106de91906110b4565b60206040518083038186803b1580156106f657600080fd5b505afa15801561070a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072e9190611037565b6001600160a01b031663a230c524336040518263ffffffff1660e01b815260040161075991906110c7565b60206040518083038186803b15801561077157600080fd5b505afa158015610785573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107a991906110db565b806108d657506000546040805180820182526014815273676f7665726e616e63652e726f6f744e6f64657360601b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb902719161080b916004016110b4565b60206040518083038186803b15801561082357600080fd5b505afa158015610837573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085b9190611037565b6001600160a01b031663a230c524336040518263ffffffff1660e01b815260040161088691906110c7565b60206040518083038186803b15801561089e57600080fd5b505afa1580156108b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d691906110db565b6109605760405162461bcd60e51b815260206004820152604f60248201527f5b5145432d3031383030315d2d5065726d697373696f6e2064656e696564202d60448201527f206f6e6c7920726f6f74206e6f64657320616e6420455051464920657870657260648201526e3a39903430bb329030b1b1b2b9b99760891b608482015260a401610391565b6002805460ff1916911515919091179055565b42600354101561099e576109856109a0565b61098f904261109c565b60035561099a610a56565b6004555b565b60006109aa610acb565b60405162498bff60e81b815260206004820152601f60248201527f676f7665726e65642e45505146492e72657365727665436f6f6c446f776e500060448201526001600160a01b03919091169063498bff00906064015b60206040518083038186803b158015610a1957600080fd5b505afa158015610a2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5191906110f8565b905090565b6000610a60610acb565b60405162498bff60e81b815260206004820152602760248201527f676f7665726e65642e45505146492e72657365727665436f6f6c446f776e54686044820152661c995cda1bdb1960ca1b60648201526001600160a01b03919091169063498bff0090608401610a01565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060238152602001611112602391396040518263ffffffff1660e01b8152600401610b2591906110b4565b60206040518083038186803b158015610b3d57600080fd5b505afa158015610b51573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a519190611037565b828054828255906000526020600020908101928215610bc2579160200282015b82811115610bc25782518051610bb2918491602090910190610bd2565b5091602001919060010190610b95565b50610bce929150610c52565b5090565b828054610bde90610f44565b90600052602060002090601f016020900481019282610c005760008555610c46565b82601f10610c1957805160ff1916838001178555610c46565b82800160010185558215610c46579182015b82811115610c46578251825591602001919060010190610c2b565b50610bce929150610c6f565b80821115610bce576000610c668282610c84565b50600101610c52565b5b80821115610bce5760008155600101610c70565b508054610c9090610f44565b6000825580601f10610ca0575050565b601f016020900490600052602060002090810190610cbe9190610c6f565b50565b6000815180845260005b81811015610ce757602081850181015186830182015201610ccb565b81811115610cf9576000602083870101525b50601f01601f19169290920160200192915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610d6357603f19888603018452610d51858351610cc1565b94509285019290850190600101610d35565b5092979650505050505050565b600060208284031215610d8257600080fd5b5035919050565b6001600160a01b0381168114610cbe57600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610ddd57610ddd610d9e565b604052919050565b6000806040808486031215610df957600080fd5b8335610e0481610d89565b925060208481013567ffffffffffffffff80821115610e2257600080fd5b8187019150601f8881840112610e3757600080fd5b823582811115610e4957610e49610d9e565b8060051b610e58868201610db4565b918252848101860191868101908c841115610e7257600080fd5b87870192505b83831015610eff57823586811115610e905760008081fd5b8701603f81018e13610ea25760008081fd5b8881013587811115610eb657610eb6610d9e565b610ec7818801601f19168b01610db4565b8181528f8c838501011115610edc5760008081fd5b818c84018c83013760009181018b01919091528352509187019190870190610e78565b8099505050505050505050509250929050565b8015158114610cbe57600080fd5b600060208284031215610f3257600080fd5b8135610f3d81610f12565b9392505050565b600181811c90821680610f5857607f821691505b6020821081141561057457634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600060208083526000845481600182811c915080831680610fb157607f831692505b858310811415610fcf57634e487b7160e01b85526022600452602485fd5b878601838152602001818015610fec5760018114610ffd57611028565b60ff19861682528782019650611028565b60008b81526020902060005b8681101561102257815484820152908501908901611009565b83019750505b50949998505050505050505050565b60006020828403121561104957600080fd5b8151610f3d81610d89565b634e487b7160e01b600052601160045260246000fd5b600060001982141561107e5761107e611054565b5060010190565b60008282101561109757611097611054565b500390565b600082198211156110af576110af611054565b500190565b602081526000610f3d6020830184610cc1565b6001600160a01b0391909116815260200190565b6000602082840312156110ed57600080fd5b8151610f3d81610f12565b60006020828403121561110a57600080fd5b505191905056fe676f7665726e616e63652e657870657274732e45505146492e706172616d6574657273676f7665726e616e63652e657870657274732e45505146492e6d656d62657273686970a26469706673582212206ea7337f46e8bf1749d52134dc939a833b54aec71b10366ee0547845fec79ac764736f6c63430008090033",
}

// SystemReserveABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemReserveMetaData.ABI instead.
var SystemReserveABI = SystemReserveMetaData.ABI

// SystemReserveBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemReserveMetaData.Bin instead.
var SystemReserveBin = SystemReserveMetaData.Bin

// DeploySystemReserve deploys a new Ethereum contract, binding an instance of SystemReserve to it.
func DeploySystemReserve(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemReserve, error) {
	parsed, err := SystemReserveMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemReserveBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemReserve{SystemReserveCaller: SystemReserveCaller{contract: contract}, SystemReserveTransactor: SystemReserveTransactor{contract: contract}, SystemReserveFilterer: SystemReserveFilterer{contract: contract}}, nil
}

// SystemReserve is an auto generated Go binding around an Ethereum contract.
type SystemReserve struct {
	SystemReserveCaller     // Read-only binding to the contract
	SystemReserveTransactor // Write-only binding to the contract
	SystemReserveFilterer   // Log filterer for contract events
}

// SystemReserveCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemReserveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemReserveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemReserveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemReserveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemReserveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemReserveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemReserveSession struct {
	Contract     *SystemReserve    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemReserveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemReserveCallerSession struct {
	Contract *SystemReserveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SystemReserveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemReserveTransactorSession struct {
	Contract     *SystemReserveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SystemReserveRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemReserveRaw struct {
	Contract *SystemReserve // Generic contract binding to access the raw methods on
}

// SystemReserveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemReserveCallerRaw struct {
	Contract *SystemReserveCaller // Generic read-only contract binding to access the raw methods on
}

// SystemReserveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemReserveTransactorRaw struct {
	Contract *SystemReserveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemReserve creates a new instance of SystemReserve, bound to a specific deployed contract.
func NewSystemReserve(address common.Address, backend bind.ContractBackend) (*SystemReserve, error) {
	contract, err := bindSystemReserve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemReserve{SystemReserveCaller: SystemReserveCaller{contract: contract}, SystemReserveTransactor: SystemReserveTransactor{contract: contract}, SystemReserveFilterer: SystemReserveFilterer{contract: contract}}, nil
}

// NewSystemReserveCaller creates a new read-only instance of SystemReserve, bound to a specific deployed contract.
func NewSystemReserveCaller(address common.Address, caller bind.ContractCaller) (*SystemReserveCaller, error) {
	contract, err := bindSystemReserve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemReserveCaller{contract: contract}, nil
}

// NewSystemReserveTransactor creates a new write-only instance of SystemReserve, bound to a specific deployed contract.
func NewSystemReserveTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemReserveTransactor, error) {
	contract, err := bindSystemReserve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemReserveTransactor{contract: contract}, nil
}

// NewSystemReserveFilterer creates a new log filterer instance of SystemReserve, bound to a specific deployed contract.
func NewSystemReserveFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemReserveFilterer, error) {
	contract, err := bindSystemReserve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemReserveFilterer{contract: contract}, nil
}

// bindSystemReserve binds a generic wrapper to an already deployed contract.
func bindSystemReserve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemReserveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemReserve *SystemReserveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemReserve.Contract.SystemReserveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemReserve *SystemReserveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemReserve.Contract.SystemReserveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemReserve *SystemReserveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemReserve.Contract.SystemReserveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemReserve *SystemReserveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemReserve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemReserve *SystemReserveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemReserve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemReserve *SystemReserveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemReserve.Contract.contract.Transact(opts, method, params...)
}

// AvailableAmount is a free data retrieval call binding the contract method 0x91f7cfb9.
//
// Solidity: function availableAmount() view returns(uint256)
func (_SystemReserve *SystemReserveCaller) AvailableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemReserve.contract.Call(opts, &out, "availableAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableAmount is a free data retrieval call binding the contract method 0x91f7cfb9.
//
// Solidity: function availableAmount() view returns(uint256)
func (_SystemReserve *SystemReserveSession) AvailableAmount() (*big.Int, error) {
	return _SystemReserve.Contract.AvailableAmount(&_SystemReserve.CallOpts)
}

// AvailableAmount is a free data retrieval call binding the contract method 0x91f7cfb9.
//
// Solidity: function availableAmount() view returns(uint256)
func (_SystemReserve *SystemReserveCallerSession) AvailableAmount() (*big.Int, error) {
	return _SystemReserve.Contract.AvailableAmount(&_SystemReserve.CallOpts)
}

// CoolDownPhase is a free data retrieval call binding the contract method 0x28011170.
//
// Solidity: function coolDownPhase() view returns(uint256)
func (_SystemReserve *SystemReserveCaller) CoolDownPhase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemReserve.contract.Call(opts, &out, "coolDownPhase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoolDownPhase is a free data retrieval call binding the contract method 0x28011170.
//
// Solidity: function coolDownPhase() view returns(uint256)
func (_SystemReserve *SystemReserveSession) CoolDownPhase() (*big.Int, error) {
	return _SystemReserve.Contract.CoolDownPhase(&_SystemReserve.CallOpts)
}

// CoolDownPhase is a free data retrieval call binding the contract method 0x28011170.
//
// Solidity: function coolDownPhase() view returns(uint256)
func (_SystemReserve *SystemReserveCallerSession) CoolDownPhase() (*big.Int, error) {
	return _SystemReserve.Contract.CoolDownPhase(&_SystemReserve.CallOpts)
}

// GetEligibleContractKeys is a free data retrieval call binding the contract method 0x2d250a5b.
//
// Solidity: function getEligibleContractKeys() view returns(string[])
func (_SystemReserve *SystemReserveCaller) GetEligibleContractKeys(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _SystemReserve.contract.Call(opts, &out, "getEligibleContractKeys")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetEligibleContractKeys is a free data retrieval call binding the contract method 0x2d250a5b.
//
// Solidity: function getEligibleContractKeys() view returns(string[])
func (_SystemReserve *SystemReserveSession) GetEligibleContractKeys() ([]string, error) {
	return _SystemReserve.Contract.GetEligibleContractKeys(&_SystemReserve.CallOpts)
}

// GetEligibleContractKeys is a free data retrieval call binding the contract method 0x2d250a5b.
//
// Solidity: function getEligibleContractKeys() view returns(string[])
func (_SystemReserve *SystemReserveCallerSession) GetEligibleContractKeys() ([]string, error) {
	return _SystemReserve.Contract.GetEligibleContractKeys(&_SystemReserve.CallOpts)
}

// SystemPaused is a free data retrieval call binding the contract method 0x9d2f83f0.
//
// Solidity: function systemPaused() view returns(bool)
func (_SystemReserve *SystemReserveCaller) SystemPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SystemReserve.contract.Call(opts, &out, "systemPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SystemPaused is a free data retrieval call binding the contract method 0x9d2f83f0.
//
// Solidity: function systemPaused() view returns(bool)
func (_SystemReserve *SystemReserveSession) SystemPaused() (bool, error) {
	return _SystemReserve.Contract.SystemPaused(&_SystemReserve.CallOpts)
}

// SystemPaused is a free data retrieval call binding the contract method 0x9d2f83f0.
//
// Solidity: function systemPaused() view returns(bool)
func (_SystemReserve *SystemReserveCallerSession) SystemPaused() (bool, error) {
	return _SystemReserve.Contract.SystemPaused(&_SystemReserve.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x94e6f56f.
//
// Solidity: function initialize(address _registry, string[] _keys) returns()
func (_SystemReserve *SystemReserveTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _keys []string) (*types.Transaction, error) {
	return _SystemReserve.contract.Transact(opts, "initialize", _registry, _keys)
}

// Initialize is a paid mutator transaction binding the contract method 0x94e6f56f.
//
// Solidity: function initialize(address _registry, string[] _keys) returns()
func (_SystemReserve *SystemReserveSession) Initialize(_registry common.Address, _keys []string) (*types.Transaction, error) {
	return _SystemReserve.Contract.Initialize(&_SystemReserve.TransactOpts, _registry, _keys)
}

// Initialize is a paid mutator transaction binding the contract method 0x94e6f56f.
//
// Solidity: function initialize(address _registry, string[] _keys) returns()
func (_SystemReserve *SystemReserveTransactorSession) Initialize(_registry common.Address, _keys []string) (*types.Transaction, error) {
	return _SystemReserve.Contract.Initialize(&_SystemReserve.TransactOpts, _registry, _keys)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xcdb88ad1.
//
// Solidity: function setPauseState(bool _state) returns()
func (_SystemReserve *SystemReserveTransactor) SetPauseState(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _SystemReserve.contract.Transact(opts, "setPauseState", _state)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xcdb88ad1.
//
// Solidity: function setPauseState(bool _state) returns()
func (_SystemReserve *SystemReserveSession) SetPauseState(_state bool) (*types.Transaction, error) {
	return _SystemReserve.Contract.SetPauseState(&_SystemReserve.TransactOpts, _state)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xcdb88ad1.
//
// Solidity: function setPauseState(bool _state) returns()
func (_SystemReserve *SystemReserveTransactorSession) SetPauseState(_state bool) (*types.Transaction, error) {
	return _SystemReserve.Contract.SetPauseState(&_SystemReserve.TransactOpts, _state)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_SystemReserve *SystemReserveTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SystemReserve.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_SystemReserve *SystemReserveSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _SystemReserve.Contract.Withdraw(&_SystemReserve.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_SystemReserve *SystemReserveTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _SystemReserve.Contract.Withdraw(&_SystemReserve.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemReserve *SystemReserveTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemReserve.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemReserve *SystemReserveSession) Receive() (*types.Transaction, error) {
	return _SystemReserve.Contract.Receive(&_SystemReserve.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemReserve *SystemReserveTransactorSession) Receive() (*types.Transaction, error) {
	return _SystemReserve.Contract.Receive(&_SystemReserve.TransactOpts)
}
