// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ParametersABI is the input ABI used to generate the binding from.
const ParametersABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinbaseAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCoinbaseAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDecimal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQIDHoldersCoinbaseSubsidyShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQIDHoldersNativeAppsFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQIDHoldersTransactionFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQTokenHoldersCoinbaseSubsidyShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQTokenHoldersNativeAppsFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQTokenHoldersTransactionFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootNodesCoinbaseSubsidyShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootNodesNativeAppsFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootNodesTransactionFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorNodesCoinbaseSubsidyShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorNodesNativeAppsFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorNodesTransactionFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_coinbaseAddress\",\"type\":\"address\"}],\"name\":\"setCoinbaseAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setQIDHoldersCoinbaseSubsidyShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setQIDHoldersNativeAppsFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setQIDHoldersTransactionFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setQTokenHoldersCoinbaseSubsidyShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setQTokenHoldersNativeAppsFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setQTokenHoldersTransactionFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setRootNodesCoinbaseSubsidyShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setRootNodesNativeAppsFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setRootNodesTransactionFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setValidatorNodesCoinbaseSubsidyShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setValidatorNodesNativeAppsFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setValidatorNodesTransactionFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"setVersionHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ParametersFuncSigs maps the 4-byte function signature to its string representation.
var ParametersFuncSigs = map[string]string{
	"325292ba": "getCoinbaseAddress()",
	"34ce10c4": "getDecimal()",
	"fc72b8d3": "getQIDHoldersCoinbaseSubsidyShare()",
	"4aa027b3": "getQIDHoldersNativeAppsFeeShare()",
	"cbb636f8": "getQIDHoldersTransactionFeeShare()",
	"be54279e": "getQTokenHoldersCoinbaseSubsidyShare()",
	"99a8cf6a": "getQTokenHoldersNativeAppsFeeShare()",
	"03dfb900": "getQTokenHoldersTransactionFeeShare()",
	"3e344f6d": "getRootNodesCoinbaseSubsidyShare()",
	"ed219949": "getRootNodesNativeAppsFeeShare()",
	"20527af1": "getRootNodesTransactionFeeShare()",
	"539c907a": "getValidatorNodesCoinbaseSubsidyShare()",
	"e54ede97": "getValidatorNodesNativeAppsFeeShare()",
	"834b7b88": "getValidatorNodesTransactionFeeShare()",
	"c6590930": "getVersionHash()",
	"82ba3321": "setCoinbaseAddress(address)",
	"a3ffe985": "setQIDHoldersCoinbaseSubsidyShare(uint256)",
	"319089dc": "setQIDHoldersNativeAppsFeeShare(uint256)",
	"5c23da02": "setQIDHoldersTransactionFeeShare(uint256)",
	"fe361df7": "setQTokenHoldersCoinbaseSubsidyShare(uint256)",
	"8a278645": "setQTokenHoldersNativeAppsFeeShare(uint256)",
	"81f7f4c8": "setQTokenHoldersTransactionFeeShare(uint256)",
	"8cf5b300": "setRootNodesCoinbaseSubsidyShare(uint256)",
	"f443f349": "setRootNodesNativeAppsFeeShare(uint256)",
	"f6e3a1d7": "setRootNodesTransactionFeeShare(uint256)",
	"d2407598": "setValidatorNodesCoinbaseSubsidyShare(uint256)",
	"d0795a79": "setValidatorNodesNativeAppsFeeShare(uint256)",
	"891fefbb": "setValidatorNodesTransactionFeeShare(uint256)",
	"94c38179": "setVersionHash(string)",
}

// ParametersBin is the compiled bytecode used for deploying new contracts.
var ParametersBin = "0x608060405261271060025534801561001657600080fd5b506040516109293803806109298339818101604052602081101561003957600080fd5b505160408051808201909152601d8082527f5120436f6e737469747574696f6e20506172616d65746572204c697374000000602090920191825261007f916000916100f1565b506040518060600160405280604081526020016108e96040913980516100ad916001916020909101906100f1565b5050610d056003819055600481905560058190556006819055600781905560088190556009819055600a819055600b819055600c819055600d819055600e5561018c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061013257805160ff191683800117855561015f565b8280016001018555821561015f579182015b8281111561015f578251825591602001919060010190610144565b5061016b92915061016f565b5090565b61018991905b8082111561016b5760008155600101610175565b90565b61074e8061019b6000396000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80638cf5b30011610104578063d0795a79116100a2578063f443f34911610071578063f443f349146104ba578063f6e3a1d7146104d7578063fc72b8d3146104f4578063fe361df7146104fc576101cf565b8063d0795a7914610470578063d24075981461048d578063e54ede97146104aa578063ed219949146104b2576101cf565b8063a3ffe985116100de578063a3ffe985146103c6578063be54279e146103e3578063c6590930146103eb578063cbb636f814610468576101cf565b80638cf5b300146102fb57806394c381791461031857806399a8cf6a146103be576101cf565b8063539c907a1161017157806382ba33211161014b57806382ba332114610293578063834b7b88146102b9578063891fefbb146102c15780638a278645146102de576101cf565b8063539c907a146102515780635c23da021461025957806381f7f4c814610276576101cf565b8063325292ba116101ad578063325292ba1461021557806334ce10c4146102395780633e344f6d146102415780634aa027b314610249576101cf565b806303dfb900146101d457806320527af1146101ee578063319089dc146101f6575b600080fd5b6101dc610519565b60408051918252519081900360200190f35b6101dc610520565b6102136004803603602081101561020c57600080fd5b5035610526565b005b61021d61052b565b604080516001600160a01b039092168252519081900360200190f35b6101dc61053a565b6101dc610540565b6101dc610546565b6101dc61054c565b6102136004803603602081101561026f57600080fd5b5035610552565b6102136004803603602081101561028c57600080fd5b5035610557565b610213600480360360208110156102a957600080fd5b50356001600160a01b031661055c565b6101dc61057e565b610213600480360360208110156102d757600080fd5b5035610584565b610213600480360360208110156102f457600080fd5b5035610589565b6102136004803603602081101561031157600080fd5b503561058e565b6102136004803603602081101561032e57600080fd5b81019060208101813564010000000081111561034957600080fd5b82018360208201111561035b57600080fd5b8035906020019184600183028401116401000000008311171561037d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610593945050505050565b6101dc6105aa565b610213600480360360208110156103dc57600080fd5b50356105b0565b6101dc6105b5565b6103f36105bb565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561042d578181015183820152602001610415565b50505050905090810190601f16801561045a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101dc610650565b6102136004803603602081101561048657600080fd5b5035610656565b610213600480360360208110156104a357600080fd5b503561065b565b6101dc610660565b6101dc610666565b610213600480360360208110156104d057600080fd5b503561066c565b610213600480360360208110156104ed57600080fd5b5035610671565b6101dc610676565b6102136004803603602081101561051257600080fd5b503561067c565b6003545b90565b600c5490565b600755565b600f546001600160a01b031690565b60025490565b600e5490565b60075490565b600b5490565b600655565b600355565b600f80546001600160a01b0319166001600160a01b0392909216919091179055565b60095490565b600955565b600455565b600e55565b80516105a6906001906020840190610681565b5050565b60045490565b600855565b60055490565b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156106465780601f1061061b57610100808354040283529160200191610646565b820191906000526020600020905b81548152906001019060200180831161062957829003601f168201915b5050505050905090565b60065490565b600a55565b600b55565b600a5490565b600d5490565b600d55565b600c55565b60085490565b600555565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106c257805160ff19168380011785556106ef565b828001600101855582156106ef579182015b828111156106ef5782518255916020019190600101906106d4565b506106fb9291506106ff565b5090565b61051d91905b808211156106fb576000815560010161070556fea265627a7a723158208c1861c6db89ad42f662e3eeff5a9084e42b4ac7cf8989e325772246b033718a64736f6c6343000511003230424142344244433743353045383746353132373341314344393544393538383530383334393739374333383433324637383839373036393346343042463837"

// DeployParameters deploys a new Ethereum contract, binding an instance of Parameters to it.
func DeployParameters(auth *bind.TransactOpts, backend bind.ContractBackend, coinbaseAddress common.Address) (common.Address, *types.Transaction, *Parameters, error) {
	parsed, err := abi.JSON(strings.NewReader(ParametersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParametersBin), backend, coinbaseAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Parameters{ParametersCaller: ParametersCaller{contract: contract}, ParametersTransactor: ParametersTransactor{contract: contract}, ParametersFilterer: ParametersFilterer{contract: contract}}, nil
}

// Parameters is an auto generated Go binding around an Ethereum contract.
type Parameters struct {
	ParametersCaller     // Read-only binding to the contract
	ParametersTransactor // Write-only binding to the contract
	ParametersFilterer   // Log filterer for contract events
}

// ParametersCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParametersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParametersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParametersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParametersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParametersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParametersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParametersSession struct {
	Contract     *Parameters       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParametersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParametersCallerSession struct {
	Contract *ParametersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ParametersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParametersTransactorSession struct {
	Contract     *ParametersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ParametersRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParametersRaw struct {
	Contract *Parameters // Generic contract binding to access the raw methods on
}

// ParametersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParametersCallerRaw struct {
	Contract *ParametersCaller // Generic read-only contract binding to access the raw methods on
}

// ParametersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParametersTransactorRaw struct {
	Contract *ParametersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParameters creates a new instance of Parameters, bound to a specific deployed contract.
func NewParameters(address common.Address, backend bind.ContractBackend) (*Parameters, error) {
	contract, err := bindParameters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Parameters{ParametersCaller: ParametersCaller{contract: contract}, ParametersTransactor: ParametersTransactor{contract: contract}, ParametersFilterer: ParametersFilterer{contract: contract}}, nil
}

// NewParametersCaller creates a new read-only instance of Parameters, bound to a specific deployed contract.
func NewParametersCaller(address common.Address, caller bind.ContractCaller) (*ParametersCaller, error) {
	contract, err := bindParameters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParametersCaller{contract: contract}, nil
}

// NewParametersTransactor creates a new write-only instance of Parameters, bound to a specific deployed contract.
func NewParametersTransactor(address common.Address, transactor bind.ContractTransactor) (*ParametersTransactor, error) {
	contract, err := bindParameters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParametersTransactor{contract: contract}, nil
}

// NewParametersFilterer creates a new log filterer instance of Parameters, bound to a specific deployed contract.
func NewParametersFilterer(address common.Address, filterer bind.ContractFilterer) (*ParametersFilterer, error) {
	contract, err := bindParameters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParametersFilterer{contract: contract}, nil
}

// bindParameters binds a generic wrapper to an already deployed contract.
func bindParameters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParametersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parameters *ParametersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Parameters.Contract.ParametersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parameters *ParametersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parameters.Contract.ParametersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parameters *ParametersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parameters.Contract.ParametersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parameters *ParametersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Parameters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parameters *ParametersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parameters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parameters *ParametersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parameters.Contract.contract.Transact(opts, method, params...)
}

// GetCoinbaseAddress is a free data retrieval call binding the contract method 0x325292ba.
//
// Solidity: function getCoinbaseAddress() view returns(address)
func (_Parameters *ParametersCaller) GetCoinbaseAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getCoinbaseAddress")
	return *ret0, err
}

// GetCoinbaseAddress is a free data retrieval call binding the contract method 0x325292ba.
//
// Solidity: function getCoinbaseAddress() view returns(address)
func (_Parameters *ParametersSession) GetCoinbaseAddress() (common.Address, error) {
	return _Parameters.Contract.GetCoinbaseAddress(&_Parameters.CallOpts)
}

// GetCoinbaseAddress is a free data retrieval call binding the contract method 0x325292ba.
//
// Solidity: function getCoinbaseAddress() view returns(address)
func (_Parameters *ParametersCallerSession) GetCoinbaseAddress() (common.Address, error) {
	return _Parameters.Contract.GetCoinbaseAddress(&_Parameters.CallOpts)
}

// GetDecimal is a free data retrieval call binding the contract method 0x34ce10c4.
//
// Solidity: function getDecimal() view returns(uint256)
func (_Parameters *ParametersCaller) GetDecimal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getDecimal")
	return *ret0, err
}

// GetDecimal is a free data retrieval call binding the contract method 0x34ce10c4.
//
// Solidity: function getDecimal() view returns(uint256)
func (_Parameters *ParametersSession) GetDecimal() (*big.Int, error) {
	return _Parameters.Contract.GetDecimal(&_Parameters.CallOpts)
}

// GetDecimal is a free data retrieval call binding the contract method 0x34ce10c4.
//
// Solidity: function getDecimal() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetDecimal() (*big.Int, error) {
	return _Parameters.Contract.GetDecimal(&_Parameters.CallOpts)
}

// GetQIDHoldersCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0xfc72b8d3.
//
// Solidity: function getQIDHoldersCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetQIDHoldersCoinbaseSubsidyShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getQIDHoldersCoinbaseSubsidyShare")
	return *ret0, err
}

// GetQIDHoldersCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0xfc72b8d3.
//
// Solidity: function getQIDHoldersCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersSession) GetQIDHoldersCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetQIDHoldersCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetQIDHoldersCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0xfc72b8d3.
//
// Solidity: function getQIDHoldersCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetQIDHoldersCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetQIDHoldersCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetQIDHoldersNativeAppsFeeShare is a free data retrieval call binding the contract method 0x4aa027b3.
//
// Solidity: function getQIDHoldersNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetQIDHoldersNativeAppsFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getQIDHoldersNativeAppsFeeShare")
	return *ret0, err
}

// GetQIDHoldersNativeAppsFeeShare is a free data retrieval call binding the contract method 0x4aa027b3.
//
// Solidity: function getQIDHoldersNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetQIDHoldersNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQIDHoldersNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetQIDHoldersNativeAppsFeeShare is a free data retrieval call binding the contract method 0x4aa027b3.
//
// Solidity: function getQIDHoldersNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetQIDHoldersNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQIDHoldersNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetQIDHoldersTransactionFeeShare is a free data retrieval call binding the contract method 0xcbb636f8.
//
// Solidity: function getQIDHoldersTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetQIDHoldersTransactionFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getQIDHoldersTransactionFeeShare")
	return *ret0, err
}

// GetQIDHoldersTransactionFeeShare is a free data retrieval call binding the contract method 0xcbb636f8.
//
// Solidity: function getQIDHoldersTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetQIDHoldersTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQIDHoldersTransactionFeeShare(&_Parameters.CallOpts)
}

// GetQIDHoldersTransactionFeeShare is a free data retrieval call binding the contract method 0xcbb636f8.
//
// Solidity: function getQIDHoldersTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetQIDHoldersTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQIDHoldersTransactionFeeShare(&_Parameters.CallOpts)
}

// GetQTokenHoldersCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0xbe54279e.
//
// Solidity: function getQTokenHoldersCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetQTokenHoldersCoinbaseSubsidyShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getQTokenHoldersCoinbaseSubsidyShare")
	return *ret0, err
}

// GetQTokenHoldersCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0xbe54279e.
//
// Solidity: function getQTokenHoldersCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersSession) GetQTokenHoldersCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetQTokenHoldersCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetQTokenHoldersCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0xbe54279e.
//
// Solidity: function getQTokenHoldersCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetQTokenHoldersCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetQTokenHoldersCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetQTokenHoldersNativeAppsFeeShare is a free data retrieval call binding the contract method 0x99a8cf6a.
//
// Solidity: function getQTokenHoldersNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetQTokenHoldersNativeAppsFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getQTokenHoldersNativeAppsFeeShare")
	return *ret0, err
}

// GetQTokenHoldersNativeAppsFeeShare is a free data retrieval call binding the contract method 0x99a8cf6a.
//
// Solidity: function getQTokenHoldersNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetQTokenHoldersNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQTokenHoldersNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetQTokenHoldersNativeAppsFeeShare is a free data retrieval call binding the contract method 0x99a8cf6a.
//
// Solidity: function getQTokenHoldersNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetQTokenHoldersNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQTokenHoldersNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetQTokenHoldersTransactionFeeShare is a free data retrieval call binding the contract method 0x03dfb900.
//
// Solidity: function getQTokenHoldersTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetQTokenHoldersTransactionFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getQTokenHoldersTransactionFeeShare")
	return *ret0, err
}

// GetQTokenHoldersTransactionFeeShare is a free data retrieval call binding the contract method 0x03dfb900.
//
// Solidity: function getQTokenHoldersTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetQTokenHoldersTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQTokenHoldersTransactionFeeShare(&_Parameters.CallOpts)
}

// GetQTokenHoldersTransactionFeeShare is a free data retrieval call binding the contract method 0x03dfb900.
//
// Solidity: function getQTokenHoldersTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetQTokenHoldersTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQTokenHoldersTransactionFeeShare(&_Parameters.CallOpts)
}

// GetRootNodesCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0x3e344f6d.
//
// Solidity: function getRootNodesCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetRootNodesCoinbaseSubsidyShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getRootNodesCoinbaseSubsidyShare")
	return *ret0, err
}

// GetRootNodesCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0x3e344f6d.
//
// Solidity: function getRootNodesCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersSession) GetRootNodesCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetRootNodesCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetRootNodesCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0x3e344f6d.
//
// Solidity: function getRootNodesCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetRootNodesCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetRootNodesCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetRootNodesNativeAppsFeeShare is a free data retrieval call binding the contract method 0xed219949.
//
// Solidity: function getRootNodesNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetRootNodesNativeAppsFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getRootNodesNativeAppsFeeShare")
	return *ret0, err
}

// GetRootNodesNativeAppsFeeShare is a free data retrieval call binding the contract method 0xed219949.
//
// Solidity: function getRootNodesNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetRootNodesNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetRootNodesNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetRootNodesNativeAppsFeeShare is a free data retrieval call binding the contract method 0xed219949.
//
// Solidity: function getRootNodesNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetRootNodesNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetRootNodesNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetRootNodesTransactionFeeShare is a free data retrieval call binding the contract method 0x20527af1.
//
// Solidity: function getRootNodesTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetRootNodesTransactionFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getRootNodesTransactionFeeShare")
	return *ret0, err
}

// GetRootNodesTransactionFeeShare is a free data retrieval call binding the contract method 0x20527af1.
//
// Solidity: function getRootNodesTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetRootNodesTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetRootNodesTransactionFeeShare(&_Parameters.CallOpts)
}

// GetRootNodesTransactionFeeShare is a free data retrieval call binding the contract method 0x20527af1.
//
// Solidity: function getRootNodesTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetRootNodesTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetRootNodesTransactionFeeShare(&_Parameters.CallOpts)
}

// GetValidatorNodesCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0x539c907a.
//
// Solidity: function getValidatorNodesCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetValidatorNodesCoinbaseSubsidyShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getValidatorNodesCoinbaseSubsidyShare")
	return *ret0, err
}

// GetValidatorNodesCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0x539c907a.
//
// Solidity: function getValidatorNodesCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersSession) GetValidatorNodesCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetValidatorNodesCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetValidatorNodesCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0x539c907a.
//
// Solidity: function getValidatorNodesCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetValidatorNodesCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetValidatorNodesCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetValidatorNodesNativeAppsFeeShare is a free data retrieval call binding the contract method 0xe54ede97.
//
// Solidity: function getValidatorNodesNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetValidatorNodesNativeAppsFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getValidatorNodesNativeAppsFeeShare")
	return *ret0, err
}

// GetValidatorNodesNativeAppsFeeShare is a free data retrieval call binding the contract method 0xe54ede97.
//
// Solidity: function getValidatorNodesNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetValidatorNodesNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetValidatorNodesNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetValidatorNodesNativeAppsFeeShare is a free data retrieval call binding the contract method 0xe54ede97.
//
// Solidity: function getValidatorNodesNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetValidatorNodesNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetValidatorNodesNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetValidatorNodesTransactionFeeShare is a free data retrieval call binding the contract method 0x834b7b88.
//
// Solidity: function getValidatorNodesTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetValidatorNodesTransactionFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getValidatorNodesTransactionFeeShare")
	return *ret0, err
}

// GetValidatorNodesTransactionFeeShare is a free data retrieval call binding the contract method 0x834b7b88.
//
// Solidity: function getValidatorNodesTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetValidatorNodesTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetValidatorNodesTransactionFeeShare(&_Parameters.CallOpts)
}

// GetValidatorNodesTransactionFeeShare is a free data retrieval call binding the contract method 0x834b7b88.
//
// Solidity: function getValidatorNodesTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetValidatorNodesTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetValidatorNodesTransactionFeeShare(&_Parameters.CallOpts)
}

// GetVersionHash is a free data retrieval call binding the contract method 0xc6590930.
//
// Solidity: function getVersionHash() view returns(string)
func (_Parameters *ParametersCaller) GetVersionHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getVersionHash")
	return *ret0, err
}

// GetVersionHash is a free data retrieval call binding the contract method 0xc6590930.
//
// Solidity: function getVersionHash() view returns(string)
func (_Parameters *ParametersSession) GetVersionHash() (string, error) {
	return _Parameters.Contract.GetVersionHash(&_Parameters.CallOpts)
}

// GetVersionHash is a free data retrieval call binding the contract method 0xc6590930.
//
// Solidity: function getVersionHash() view returns(string)
func (_Parameters *ParametersCallerSession) GetVersionHash() (string, error) {
	return _Parameters.Contract.GetVersionHash(&_Parameters.CallOpts)
}

// SetCoinbaseAddress is a paid mutator transaction binding the contract method 0x82ba3321.
//
// Solidity: function setCoinbaseAddress(address _coinbaseAddress) returns()
func (_Parameters *ParametersTransactor) SetCoinbaseAddress(opts *bind.TransactOpts, _coinbaseAddress common.Address) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setCoinbaseAddress", _coinbaseAddress)
}

// SetCoinbaseAddress is a paid mutator transaction binding the contract method 0x82ba3321.
//
// Solidity: function setCoinbaseAddress(address _coinbaseAddress) returns()
func (_Parameters *ParametersSession) SetCoinbaseAddress(_coinbaseAddress common.Address) (*types.Transaction, error) {
	return _Parameters.Contract.SetCoinbaseAddress(&_Parameters.TransactOpts, _coinbaseAddress)
}

// SetCoinbaseAddress is a paid mutator transaction binding the contract method 0x82ba3321.
//
// Solidity: function setCoinbaseAddress(address _coinbaseAddress) returns()
func (_Parameters *ParametersTransactorSession) SetCoinbaseAddress(_coinbaseAddress common.Address) (*types.Transaction, error) {
	return _Parameters.Contract.SetCoinbaseAddress(&_Parameters.TransactOpts, _coinbaseAddress)
}

// SetQIDHoldersCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xa3ffe985.
//
// Solidity: function setQIDHoldersCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetQIDHoldersCoinbaseSubsidyShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setQIDHoldersCoinbaseSubsidyShare", feeShare)
}

// SetQIDHoldersCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xa3ffe985.
//
// Solidity: function setQIDHoldersCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetQIDHoldersCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQIDHoldersCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetQIDHoldersCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xa3ffe985.
//
// Solidity: function setQIDHoldersCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetQIDHoldersCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQIDHoldersCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetQIDHoldersNativeAppsFeeShare is a paid mutator transaction binding the contract method 0x319089dc.
//
// Solidity: function setQIDHoldersNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetQIDHoldersNativeAppsFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setQIDHoldersNativeAppsFeeShare", feeShare)
}

// SetQIDHoldersNativeAppsFeeShare is a paid mutator transaction binding the contract method 0x319089dc.
//
// Solidity: function setQIDHoldersNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetQIDHoldersNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQIDHoldersNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQIDHoldersNativeAppsFeeShare is a paid mutator transaction binding the contract method 0x319089dc.
//
// Solidity: function setQIDHoldersNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetQIDHoldersNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQIDHoldersNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQIDHoldersTransactionFeeShare is a paid mutator transaction binding the contract method 0x5c23da02.
//
// Solidity: function setQIDHoldersTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetQIDHoldersTransactionFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setQIDHoldersTransactionFeeShare", feeShare)
}

// SetQIDHoldersTransactionFeeShare is a paid mutator transaction binding the contract method 0x5c23da02.
//
// Solidity: function setQIDHoldersTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetQIDHoldersTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQIDHoldersTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQIDHoldersTransactionFeeShare is a paid mutator transaction binding the contract method 0x5c23da02.
//
// Solidity: function setQIDHoldersTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetQIDHoldersTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQIDHoldersTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQTokenHoldersCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xfe361df7.
//
// Solidity: function setQTokenHoldersCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetQTokenHoldersCoinbaseSubsidyShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setQTokenHoldersCoinbaseSubsidyShare", feeShare)
}

// SetQTokenHoldersCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xfe361df7.
//
// Solidity: function setQTokenHoldersCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetQTokenHoldersCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQTokenHoldersCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetQTokenHoldersCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xfe361df7.
//
// Solidity: function setQTokenHoldersCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetQTokenHoldersCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQTokenHoldersCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetQTokenHoldersNativeAppsFeeShare is a paid mutator transaction binding the contract method 0x8a278645.
//
// Solidity: function setQTokenHoldersNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetQTokenHoldersNativeAppsFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setQTokenHoldersNativeAppsFeeShare", feeShare)
}

// SetQTokenHoldersNativeAppsFeeShare is a paid mutator transaction binding the contract method 0x8a278645.
//
// Solidity: function setQTokenHoldersNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetQTokenHoldersNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQTokenHoldersNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQTokenHoldersNativeAppsFeeShare is a paid mutator transaction binding the contract method 0x8a278645.
//
// Solidity: function setQTokenHoldersNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetQTokenHoldersNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQTokenHoldersNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQTokenHoldersTransactionFeeShare is a paid mutator transaction binding the contract method 0x81f7f4c8.
//
// Solidity: function setQTokenHoldersTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetQTokenHoldersTransactionFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setQTokenHoldersTransactionFeeShare", feeShare)
}

// SetQTokenHoldersTransactionFeeShare is a paid mutator transaction binding the contract method 0x81f7f4c8.
//
// Solidity: function setQTokenHoldersTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetQTokenHoldersTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQTokenHoldersTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQTokenHoldersTransactionFeeShare is a paid mutator transaction binding the contract method 0x81f7f4c8.
//
// Solidity: function setQTokenHoldersTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetQTokenHoldersTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQTokenHoldersTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetRootNodesCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0x8cf5b300.
//
// Solidity: function setRootNodesCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetRootNodesCoinbaseSubsidyShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setRootNodesCoinbaseSubsidyShare", feeShare)
}

// SetRootNodesCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0x8cf5b300.
//
// Solidity: function setRootNodesCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetRootNodesCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetRootNodesCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetRootNodesCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0x8cf5b300.
//
// Solidity: function setRootNodesCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetRootNodesCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetRootNodesCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetRootNodesNativeAppsFeeShare is a paid mutator transaction binding the contract method 0xf443f349.
//
// Solidity: function setRootNodesNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetRootNodesNativeAppsFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setRootNodesNativeAppsFeeShare", feeShare)
}

// SetRootNodesNativeAppsFeeShare is a paid mutator transaction binding the contract method 0xf443f349.
//
// Solidity: function setRootNodesNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetRootNodesNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetRootNodesNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetRootNodesNativeAppsFeeShare is a paid mutator transaction binding the contract method 0xf443f349.
//
// Solidity: function setRootNodesNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetRootNodesNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetRootNodesNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetRootNodesTransactionFeeShare is a paid mutator transaction binding the contract method 0xf6e3a1d7.
//
// Solidity: function setRootNodesTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetRootNodesTransactionFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setRootNodesTransactionFeeShare", feeShare)
}

// SetRootNodesTransactionFeeShare is a paid mutator transaction binding the contract method 0xf6e3a1d7.
//
// Solidity: function setRootNodesTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetRootNodesTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetRootNodesTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetRootNodesTransactionFeeShare is a paid mutator transaction binding the contract method 0xf6e3a1d7.
//
// Solidity: function setRootNodesTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetRootNodesTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetRootNodesTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetValidatorNodesCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xd2407598.
//
// Solidity: function setValidatorNodesCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetValidatorNodesCoinbaseSubsidyShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setValidatorNodesCoinbaseSubsidyShare", feeShare)
}

// SetValidatorNodesCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xd2407598.
//
// Solidity: function setValidatorNodesCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetValidatorNodesCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetValidatorNodesCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetValidatorNodesCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xd2407598.
//
// Solidity: function setValidatorNodesCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetValidatorNodesCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetValidatorNodesCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetValidatorNodesNativeAppsFeeShare is a paid mutator transaction binding the contract method 0xd0795a79.
//
// Solidity: function setValidatorNodesNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetValidatorNodesNativeAppsFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setValidatorNodesNativeAppsFeeShare", feeShare)
}

// SetValidatorNodesNativeAppsFeeShare is a paid mutator transaction binding the contract method 0xd0795a79.
//
// Solidity: function setValidatorNodesNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetValidatorNodesNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetValidatorNodesNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetValidatorNodesNativeAppsFeeShare is a paid mutator transaction binding the contract method 0xd0795a79.
//
// Solidity: function setValidatorNodesNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetValidatorNodesNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetValidatorNodesNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetValidatorNodesTransactionFeeShare is a paid mutator transaction binding the contract method 0x891fefbb.
//
// Solidity: function setValidatorNodesTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetValidatorNodesTransactionFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setValidatorNodesTransactionFeeShare", feeShare)
}

// SetValidatorNodesTransactionFeeShare is a paid mutator transaction binding the contract method 0x891fefbb.
//
// Solidity: function setValidatorNodesTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetValidatorNodesTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetValidatorNodesTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetValidatorNodesTransactionFeeShare is a paid mutator transaction binding the contract method 0x891fefbb.
//
// Solidity: function setValidatorNodesTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetValidatorNodesTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetValidatorNodesTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetVersionHash is a paid mutator transaction binding the contract method 0x94c38179.
//
// Solidity: function setVersionHash(string hash) returns()
func (_Parameters *ParametersTransactor) SetVersionHash(opts *bind.TransactOpts, hash string) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setVersionHash", hash)
}

// SetVersionHash is a paid mutator transaction binding the contract method 0x94c38179.
//
// Solidity: function setVersionHash(string hash) returns()
func (_Parameters *ParametersSession) SetVersionHash(hash string) (*types.Transaction, error) {
	return _Parameters.Contract.SetVersionHash(&_Parameters.TransactOpts, hash)
}

// SetVersionHash is a paid mutator transaction binding the contract method 0x94c38179.
//
// Solidity: function setVersionHash(string hash) returns()
func (_Parameters *ParametersTransactorSession) SetVersionHash(hash string) (*types.Transaction, error) {
	return _Parameters.Contract.SetVersionHash(&_Parameters.TransactOpts, hash)
}
