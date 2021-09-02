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
	Bin: "0x608060405234801561001057600080fd5b5061111e806100206000396000f3fe6080604052600436106100645760003560e01c806328011170146100705780632d250a5b1461009b5780632e1a7d4d146100bd57806391f7cfb9146100ea57806394e6f56f146100ff5780639d2f83f014610121578063cdb88ad1146101365761006b565b3661006b57005b600080fd5b34801561007c57600080fd5b50610085610156565b604051610092919061105e565b60405180910390f35b3480156100a757600080fd5b506100b061015c565b6040516100929190610c1d565b3480156100c957600080fd5b506100dd6100d8366004610bd6565b610234565b6040516100929190610cb4565b3480156100f657600080fd5b5061008561041c565b34801561010b57600080fd5b5061011f61011a366004610aa9565b610422565b005b34801561012d57600080fd5b506100dd610514565b34801561014257600080fd5b5061011f610151366004610b9e565b61051d565b60035481565b60606001805480602002602001604051908101604052809291908181526020016000905b8282101561022b5760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156102175780601f106101ec57610100808354040283529160200191610217565b820191906000526020600020905b8154815290600101906020018083116101fa57829003601f168201915b505050505081526020019060010190610180565b50505050905090565b60008060005b600154811015610303576000546001805433926201000090046001600160a01b031691633fb90271918590811061026d57fe5b906000526020600020016040518263ffffffff1660e01b81526004016102939190610cbf565b60206040518083038186803b1580156102ab57600080fd5b505afa1580156102bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e39190610a86565b6001600160a01b031614156102fb5760019150610303565b60010161023a565b508061032a5760405162461bcd60e51b815260040161032190610f7e565b60405180910390fd5b60025460ff161561034d5760405162461bcd60e51b815260040161032190610ef3565b61035561074e565b82471080610361575082155b1561036f5760009150610416565b8260045410156103915760405162461bcd60e51b815260040161032190610e9a565b600480548490039055604051600090339085906103ad90610c06565b60006040518083038185875af1925050503d80600081146103ea576040519150601f19603f3d011682016040523d82523d6000602084013e6103ef565b606091505b50509050806104105760405162461bcd60e51b815260040161032190610d49565b60019250505b50919050565b60045481565b600054610100900460ff168061043b575061043b610773565b80610449575060005460ff16155b6104845760405162461bcd60e51b815260040180806020018281038252602e8152602001806110bb602e913960400191505060405180910390fd5b600054610100900460ff161580156104af576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b0386160217905581516104e4906001906020850190610928565b506104ed610779565b42016003556104fa61087b565b600455801561050f576000805461ff00191690555b505050565b60025460ff1681565b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061055090600401610de2565b60206040518083038186803b15801561056857600080fd5b505afa15801561057c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105a09190610a86565b6001600160a01b031663a230c524336040518263ffffffff1660e01b81526004016105cb9190610c09565b60206040518083038186803b1580156105e357600080fd5b505afa1580156105f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061b9190610bba565b8061071f5750600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061065490600401610fe9565b60206040518083038186803b15801561066c57600080fd5b505afa158015610680573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a49190610a86565b6001600160a01b031663a230c524336040518263ffffffff1660e01b81526004016106cf9190610c09565b60206040518083038186803b1580156106e757600080fd5b505afa1580156106fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071f9190610bba565b61073b5760405162461bcd60e51b815260040161032190610e25565b6002805460ff1916911515919091179055565b42600354101561077157610760610779565b420160035561076d61087b565b6004555b565b303b1590565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906107ad90600401610d9f565b60206040518083038186803b1580156107c557600080fd5b505afa1580156107d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107fd9190610a86565b6001600160a01b031663498bff006040518163ffffffff1660e01b815260040161082690610f47565b60206040518083038186803b15801561083e57600080fd5b505afa158015610852573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108769190610bee565b905090565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906108af90600401610d9f565b60206040518083038186803b1580156108c757600080fd5b505afa1580156108db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ff9190610a86565b6001600160a01b031663498bff006040518163ffffffff1660e01b815260040161082690611017565b828054828255906000526020600020908101928215610975579160200282015b828111156109755782518051610965918491602090910190610985565b5091602001919060010190610948565b50610981929150610a0d565b5090565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826109bb5760008555610a01565b82601f106109d457805160ff1916838001178555610a01565b82800160010185558215610a01579182015b82811115610a015782518255916020019190600101906109e6565b50610981929150610a2a565b80821115610981576000610a218282610a3f565b50600101610a0d565b5b808211156109815760008155600101610a2b565b50805460018160011615610100020316600290046000825580601f10610a655750610a83565b601f016020900490600052602060002090810190610a839190610a2a565b50565b600060208284031215610a97578081fd5b8151610aa281611097565b9392505050565b6000806040808486031215610abc578182fd5b8335610ac781611097565b925060208481013567ffffffffffffffff80821115610ae4578485fd5b818701915087601f830112610af7578485fd5b813581811115610b0357fe5b610b108485830201611067565b81815284810190848601885b84811015610b8c57813587018d603f820112610b36578a8bfd5b8881013587811115610b4457fe5b610b56601f8201601f19168b01611067565b8181528f8c838501011115610b69578c8dfd5b818c84018c8301379081018a018c90528552509287019290870190600101610b1c565b50989b909a5098505050505050505050565b600060208284031215610baf578081fd5b8135610aa2816110ac565b600060208284031215610bcb578081fd5b8151610aa2816110ac565b600060208284031215610be7578081fd5b5035919050565b600060208284031215610bff578081fd5b5051919050565b90565b6001600160a01b0391909116815260200190565b6000602080830181845280855180835260408601915060408482028701019250838701855b82811015610ca757878503603f1901845281518051808752885b81811015610c77578281018901518882018a01528801610c5c565b81811115610c87578989838a0101525b50601f01601f191695909501860194509285019290850190600101610c42565b5092979650505050505050565b901515815260200190565b60006020808301818452828554600180821660008114610ce65760018114610d0457610d3c565b60028304607f16855260ff1983166040890152606088019350610d3c565b60028304808652610d148a61108b565b885b82811015610d325781548b820160400152908401908801610d16565b8a01604001955050505b5091979650505050505050565b60208082526036908201527f5b5145432d3031383030345d2d5472616e73666572206f66207468652077697460408201527534323930bbb0b61030b6b7bab73a103330b4b632b21760511b606082015260800190565b60208082526023908201527f676f7665726e616e63652e657870657274732e45505146492e706172616d657460408201526265727360e81b606082015260800190565b60208082526023908201527f676f7665726e616e63652e657870657274732e45505146492e6d656d6265727360408201526206869760ec1b606082015260800190565b6020808252604f908201527f5b5145432d3031383030315d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c7920726f6f74206e6f64657320616e6420455051464920657870657260608201526e3a39903430bb329030b1b1b2b9b99760891b608082015260a00190565b60208082526039908201527f5b5145432d3031383030335d2d496e73756666696369656e742066756e64732060408201527830bb30b4b630b13632903337b9103bb4ba34323930bbb0b61760391b606082015260800190565b60208082526034908201527f5b5145432d3031383030325d2d5468652073797374656d20726573657276652060408201527331b7b73a3930b1ba1034b9903830bab9b2b2171760611b606082015260800190565b6020808252601f908201527f676f7665726e65642e45505146492e72657365727665436f6f6c446f776e5000604082015260600190565b60208082526045908201527f5b5145432d3031383030305d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c7920656c696769626c6520636f6e747261637473206861766520616360608201526431b2b9b99760d91b608082015260a00190565b602080825260149082015273676f7665726e616e63652e726f6f744e6f64657360601b604082015260600190565b60208082526027908201527f676f7665726e65642e45505146492e72657365727665436f6f6c446f776e54686040820152661c995cda1bdb1960ca1b606082015260800190565b90815260200190565b60405181810167ffffffffffffffff8111828210171561108357fe5b604052919050565b60009081526020902090565b6001600160a01b0381168114610a8357600080fd5b8015158114610a8357600080fdfe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a2646970667358221220fef964ebb8a6e0b1cb3c693e68e51cfddc743e1c5d41181f6782bc95cfbc44d664736f6c63430007060033",
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
