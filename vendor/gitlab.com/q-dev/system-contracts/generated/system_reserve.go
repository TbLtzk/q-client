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
	Bin: "0x608060405234801561001057600080fd5b5061112f806100206000396000f3fe6080604052600436106100645760003560e01c806328011170146100705780632d250a5b1461009b5780632e1a7d4d146100bd57806391f7cfb9146100ea57806394e6f56f146100ff5780639d2f83f014610121578063cdb88ad1146101365761006b565b3661006b57005b600080fd5b34801561007c57600080fd5b50610085610156565b604051610092919061106f565b60405180910390f35b3480156100a757600080fd5b506100b061015c565b6040516100929190610c2e565b3480156100c957600080fd5b506100dd6100d8366004610be7565b610234565b6040516100929190610cc5565b3480156100f657600080fd5b5061008561041c565b34801561010b57600080fd5b5061011f61011a366004610aba565b610422565b005b34801561012d57600080fd5b506100dd610514565b34801561014257600080fd5b5061011f610151366004610baf565b61051d565b60035481565b60606001805480602002602001604051908101604052809291908181526020016000905b8282101561022b5760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156102175780601f106101ec57610100808354040283529160200191610217565b820191906000526020600020905b8154815290600101906020018083116101fa57829003601f168201915b505050505081526020019060010190610180565b50505050905090565b60008060005b600154811015610303576000546001805433926201000090046001600160a01b031691633fb90271918590811061026d57fe5b906000526020600020016040518263ffffffff1660e01b81526004016102939190610cd0565b60206040518083038186803b1580156102ab57600080fd5b505afa1580156102bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e39190610a97565b6001600160a01b031614156102fb5760019150610303565b60010161023a565b508061032a5760405162461bcd60e51b815260040161032190610f8f565b60405180910390fd5b60025460ff161561034d5760405162461bcd60e51b815260040161032190610f04565b61035561074e565b82471080610361575082155b1561036f5760009150610416565b8260045410156103915760405162461bcd60e51b815260040161032190610eab565b600480548490039055604051600090339085906103ad90610c17565b60006040518083038185875af1925050503d80600081146103ea576040519150601f19603f3d011682016040523d82523d6000602084013e6103ef565b606091505b50509050806104105760405162461bcd60e51b815260040161032190610d5a565b60019250505b50919050565b60045481565b600054610100900460ff168061043b575061043b610773565b80610449575060005460ff16155b6104845760405162461bcd60e51b815260040180806020018281038252602e8152602001806110cc602e913960400191505060405180910390fd5b600054610100900460ff161580156104af576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b0386160217905581516104e4906001906020850190610939565b506104ed610784565b42016003556104fa610886565b600455801561050f576000805461ff00191690555b505050565b60025460ff1681565b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061055090600401610df3565b60206040518083038186803b15801561056857600080fd5b505afa15801561057c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105a09190610a97565b6001600160a01b031663a230c524336040518263ffffffff1660e01b81526004016105cb9190610c1a565b60206040518083038186803b1580156105e357600080fd5b505afa1580156105f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061b9190610bcb565b8061071f5750600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061065490600401610ffa565b60206040518083038186803b15801561066c57600080fd5b505afa158015610680573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a49190610a97565b6001600160a01b031663a230c524336040518263ffffffff1660e01b81526004016106cf9190610c1a565b60206040518083038186803b1580156106e757600080fd5b505afa1580156106fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071f9190610bcb565b61073b5760405162461bcd60e51b815260040161032190610e36565b6002805460ff1916911515919091179055565b42600354101561077157610760610784565b420160035561076d610886565b6004555b565b600061077e30610933565b15905090565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906107b890600401610db0565b60206040518083038186803b1580156107d057600080fd5b505afa1580156107e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108089190610a97565b6001600160a01b031663498bff006040518163ffffffff1660e01b815260040161083190610f58565b60206040518083038186803b15801561084957600080fd5b505afa15801561085d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108819190610bff565b905090565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906108ba90600401610db0565b60206040518083038186803b1580156108d257600080fd5b505afa1580156108e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061090a9190610a97565b6001600160a01b031663498bff006040518163ffffffff1660e01b815260040161083190611028565b3b151590565b828054828255906000526020600020908101928215610986579160200282015b828111156109865782518051610976918491602090910190610996565b5091602001919060010190610959565b50610992929150610a1e565b5090565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826109cc5760008555610a12565b82601f106109e557805160ff1916838001178555610a12565b82800160010185558215610a12579182015b82811115610a125782518255916020019190600101906109f7565b50610992929150610a3b565b80821115610992576000610a328282610a50565b50600101610a1e565b5b808211156109925760008155600101610a3c565b50805460018160011615610100020316600290046000825580601f10610a765750610a94565b601f016020900490600052602060002090810190610a949190610a3b565b50565b600060208284031215610aa8578081fd5b8151610ab3816110a8565b9392505050565b6000806040808486031215610acd578182fd5b8335610ad8816110a8565b925060208481013567ffffffffffffffff80821115610af5578485fd5b818701915087601f830112610b08578485fd5b813581811115610b1457fe5b610b218485830201611078565b81815284810190848601885b84811015610b9d57813587018d603f820112610b47578a8bfd5b8881013587811115610b5557fe5b610b67601f8201601f19168b01611078565b8181528f8c838501011115610b7a578c8dfd5b818c84018c8301379081018a018c90528552509287019290870190600101610b2d565b50989b909a5098505050505050505050565b600060208284031215610bc0578081fd5b8135610ab3816110bd565b600060208284031215610bdc578081fd5b8151610ab3816110bd565b600060208284031215610bf8578081fd5b5035919050565b600060208284031215610c10578081fd5b5051919050565b90565b6001600160a01b0391909116815260200190565b6000602080830181845280855180835260408601915060408482028701019250838701855b82811015610cb857878503603f1901845281518051808752885b81811015610c88578281018901518882018a01528801610c6d565b81811115610c98578989838a0101525b50601f01601f191695909501860194509285019290850190600101610c53565b5092979650505050505050565b901515815260200190565b60006020808301818452828554600180821660008114610cf75760018114610d1557610d4d565b60028304607f16855260ff1983166040890152606088019350610d4d565b60028304808652610d258a61109c565b885b82811015610d435781548b820160400152908401908801610d27565b8a01604001955050505b5091979650505050505050565b60208082526036908201527f5b5145432d3031383030345d2d5472616e73666572206f66207468652077697460408201527534323930bbb0b61030b6b7bab73a103330b4b632b21760511b606082015260800190565b60208082526023908201527f676f7665726e616e63652e657870657274732e45505146492e706172616d657460408201526265727360e81b606082015260800190565b60208082526023908201527f676f7665726e616e63652e657870657274732e45505146492e6d656d6265727360408201526206869760ec1b606082015260800190565b6020808252604f908201527f5b5145432d3031383030315d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c7920726f6f74206e6f64657320616e6420455051464920657870657260608201526e3a39903430bb329030b1b1b2b9b99760891b608082015260a00190565b60208082526039908201527f5b5145432d3031383030335d2d496e73756666696369656e742066756e64732060408201527830bb30b4b630b13632903337b9103bb4ba34323930bbb0b61760391b606082015260800190565b60208082526034908201527f5b5145432d3031383030325d2d5468652073797374656d20726573657276652060408201527331b7b73a3930b1ba1034b9903830bab9b2b2171760611b606082015260800190565b6020808252601f908201527f676f7665726e65642e45505146492e72657365727665436f6f6c446f776e5000604082015260600190565b60208082526045908201527f5b5145432d3031383030305d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c7920656c696769626c6520636f6e747261637473206861766520616360608201526431b2b9b99760d91b608082015260a00190565b602080825260149082015273676f7665726e616e63652e726f6f744e6f64657360601b604082015260600190565b60208082526027908201527f676f7665726e65642e45505146492e72657365727665436f6f6c446f776e54686040820152661c995cda1bdb1960ca1b606082015260800190565b90815260200190565b60405181810167ffffffffffffffff8111828210171561109457fe5b604052919050565b60009081526020902090565b6001600160a01b0381168114610a9457600080fd5b8015158114610a9457600080fdfe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a2646970667358221220f8448ae8504e32d7780d0746cc6ce81b6aede11b2487611b7c5e718cb1808ca764736f6c63430007060033",
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
