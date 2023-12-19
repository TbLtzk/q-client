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

// SystemBalanceSystemBalanceDetails is an auto generated low-level Go binding around an user-defined struct.
type SystemBalanceSystemBalanceDetails struct {
	IsDebtAuctionPossible    bool
	IsSurplusAuctionPossible bool
	CurrentDebt              *big.Int
	DebtThreshold            *big.Int
	CurrentSurplus           *big.Int
	SurplusThreshold         *big.Int
}

// SystemBalanceMetaData contains all meta data concerning the SystemBalance contract.
var SystemBalanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalanceDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isDebtAuctionPossible\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSurplusAuctionPossible\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"currentDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentSurplus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"surplusThreshold\",\"type\":\"uint256\"}],\"internalType\":\"structSystemBalance.SystemBalanceDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSurplus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtAmount\",\"type\":\"uint256\"}],\"name\":\"increaseDebt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_stc\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performNetting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferAccruedInterestAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lot\",\"type\":\"uint256\"}],\"name\":\"transferSurplusAuctionAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611a7a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100835760003560e01c806312065fe01461008857806314a6bf0f146100a3578063152353d9146100ab5780632383b074146100ce5780632b7c7b11146100d65780633d03f797146100e957806364eb4369146100f1578063effd85ad14610149578063f399e22e1461015c575b600080fd5b610090610171565b6040519081526020015b60405180910390f35b600254610090565b6100be6100b936600461155a565b610184565b604051901515815260200161009a565b61009061035f565b6100be6100e436600461155a565b6103e7565b6100be610719565b6100f96107cd565b60405161009a9190600060c08201905082511515825260208301511515602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015292915050565b6100be61015736600461155a565b610ded565b61016f61016a3660046115a1565b610f22565b005b600060025461017e61035f565b03905090565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb9027161023c600180546101b990611665565b80601f01602080910402602001604051908101604052809291908181526020018280546101e590611665565b80156102325780601f1061020757610100808354040283529160200191610232565b820191906000526020600020905b81548152906001019060200180831161021557829003601f168201915b505050505061100d565b6040518263ffffffff1660e01b815260040161025891906116d0565b60206040518083038186803b15801561027057600080fd5b505afa158015610284573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102a89190611703565b6001600160a01b0316336001600160a01b03161461034f5760405162461bcd60e51b815260206004820152605360248201527f5b5145432d3032333030305d2d5065726d697373696f6e2064656e696564202d60448201527f206f6e6c79207468652053797374656d537572706c757341756374696f6e206360648201527237b73a3930b1ba103430b99030b1b1b2b9b99760691b608482015260a4015b60405180910390fd5b6103593383611078565b92915050565b600061036961110d565b6040516370a0823160e01b81523060048201526001600160a01b0391909116906370a082319060240160206040518083038186803b1580156103aa57600080fd5b505afa1580156103be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e29190611720565b905090565b600080600060029054906101000a90046001600160a01b03166001600160a01b0316633fb902716104a16001805461041e90611665565b80601f016020809104026020016040519081016040528092919081815260200182805461044a90611665565b80156104975780601f1061046c57610100808354040283529160200191610497565b820191906000526020600020905b81548152906001019060200180831161047a57829003601f168201915b5050505050611231565b6040518263ffffffff1660e01b81526004016104bd91906116d0565b60206040518083038186803b1580156104d557600080fd5b505afa1580156104e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061050d9190611703565b6001600160a01b0316336001600160a01b031614905060008060029054906101000a90046001600160a01b03166001600160a01b0316633fb902716105db6001805461055890611665565b80601f016020809104026020016040519081016040528092919081815260200182805461058490611665565b80156105d15780601f106105a6576101008083540402835291602001916105d1565b820191906000526020600020905b8154815290600101906020018083116105b457829003601f168201915b5050505050611286565b6040518263ffffffff1660e01b81526004016105f791906116d0565b60206040518083038186803b15801561060f57600080fd5b505afa158015610623573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106479190611703565b6001600160a01b0316336001600160a01b031614905081806106665750805b6106fe5760405162461bcd60e51b815260206004820152605f60248201527f5b5145432d3032333030325d2d5065726d697373696f6e2064656e696564202d60448201527f206f6e6c7920746865204c69717569646174696f6e2061756374696f6e20616e60648201527f6420536176696e6720636f6e7472616374732068617665206163636573732e00608482015260a401610346565b8360025461070c919061174f565b6002555060019392505050565b60008061072461035f565b9050600060025482111561073b575060025461073e565b50805b8061074c5760009250505090565b8060025461075a9190611767565b60025561076561110d565b6001600160a01b03166342966c68826040518263ffffffff1660e01b815260040161079291815260200190565b600060405180830381600087803b1580156107ac57600080fd5b505af11580156107c0573d6000803e3d6000fd5b5050505060019250505090565b6107d5611487565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060228152602001611a23602291396040518263ffffffff1660e01b815260040161082f91906116d0565b60206040518083038186803b15801561084757600080fd5b505afa15801561085b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087f9190611703565b9050610889611487565b600254604082015261089961035f565b816080018181525050816001600160a01b031663498bff00610944600180546108c190611665565b80601f01602080910402602001604051908101604052809291908181526020018280546108ed90611665565b801561093a5780601f1061090f5761010080835404028352916020019161093a565b820191906000526020600020905b81548152906001019060200180831161091d57829003601f168201915b50505050506112db565b6040518263ffffffff1660e01b815260040161096091906116d0565b60206040518083038186803b15801561097857600080fd5b505afa15801561098c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b09190611720565b816060018181525050816001600160a01b031663498bff00610a5b600180546109d890611665565b80601f0160208091040260200160405190810160405280929190818152602001828054610a0490611665565b8015610a515780601f10610a2657610100808354040283529160200191610a51565b820191906000526020600020905b815481529060010190602001808311610a3457829003601f168201915b505050505061130c565b6040518263ffffffff1660e01b8152600401610a7791906116d0565b60206040518083038186803b158015610a8f57600080fd5b505afa158015610aa3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac79190611720565b60a08201526080810151604082015110610ca057600081608001518260400151610af19190611767565b9050816060015181118015610c965750600060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271610bb860018054610b3590611665565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6190611665565b8015610bae5780601f10610b8357610100808354040283529160200191610bae565b820191906000526020600020905b815481529060010190602001808311610b9157829003601f168201915b5050505050611340565b6040518263ffffffff1660e01b8152600401610bd491906116d0565b60206040518083038186803b158015610bec57600080fd5b505afa158015610c00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c249190611703565b6001600160a01b031663500ff7e66040518163ffffffff1660e01b815260040160206040518083038186803b158015610c5c57600080fd5b505afa158015610c70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c94919061177e565b155b1515825250610359565b600081604001518260800151610cb69190611767565b90506000836001600160a01b031663498bff00610d5c60018054610cd990611665565b80601f0160208091040260200160405190810160405280929190818152602001828054610d0590611665565b8015610d525780601f10610d2757610100808354040283529160200191610d52565b820191906000526020600020905b815481529060010190602001808311610d3557829003601f168201915b5050505050611395565b6040518263ffffffff1660e01b8152600401610d7891906116d0565b60206040518083038186803b158015610d9057600080fd5b505afa158015610da4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc89190611720565b90508260a001518210158015610dde5750808210155b15156020840152505092915050565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271610e226001805461055890611665565b6040518263ffffffff1660e01b8152600401610e3e91906116d0565b60206040518083038186803b158015610e5657600080fd5b505afa158015610e6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8e9190611703565b6001600160a01b0316336001600160a01b03161461034f5760405162461bcd60e51b815260206004820152604560248201527f5b5145432d3032333030315d2d5065726d697373696f6e2064656e696564202d60448201527f206f6e6c792074686520536176696e6720636f6e74726163742068617320616360648201526431b2b9b99760d91b608482015260a401610346565b600054610100900460ff1680610f3b575060005460ff16155b610f9e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610346565b600054610100900460ff16158015610fc0576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b038616021790558151610ff59060019060208501906114c1565b508015611008576000805461ff00191690555b505050565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b8284018190528551808701875291825281840152935160609461106294938792016117a0565b6040516020818303038152906040529050919050565b600061108261110d565b60405163a9059cbb60e01b81526001600160a01b03858116600483015260248201859052919091169063a9059cbb90604401602060405180830381600087803b1580156110ce57600080fd5b505af11580156110e2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611106919061177e565b9392505050565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb902716111c56001805461114290611665565b80601f016020809104026020016040519081016040528092919081815260200182805461116e90611665565b80156111bb5780601f10611190576101008083540402835291602001916111bb565b820191906000526020600020905b81548152906001019060200180831161119e57829003601f168201915b50505050506113c3565b6040518263ffffffff1660e01b81526004016111e191906116d0565b60206040518083038186803b1580156111f957600080fd5b505afa15801561120d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e29190611703565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b828401819052855180870187529182528184015293516060946110629493879201611816565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b82840181905285518087018752918252818401529351606094611062949387920161188a565b6060610359826040518060400160405280600d81526020016c1919589d151a1c995cda1bdb19609a1b815250611418565b6060610359826040518060400160405280601081526020016f1cdd5c9c1b1d5cd51a1c995cda1bdb1960821b815250611418565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b8284018190528551808701875291825281840152935160609461106294938792016118f2565b6060610359826040518060400160405280600a8152602001691cdd5c9c1b1d5cd31bdd60b21b815250611418565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b828401819052855180870187529182528184015293516060946110629493879201611965565b60606040518060400160405280600e81526020016d33b7bb32b93732b21722a822291760911b81525083604051806040016040528060018152602001605f60f81b8152508460405160200161147094939291906119cb565b604051602081830303815290604052905092915050565b6040518060c00160405280600015158152602001600015158152602001600081526020016000815260200160008152602001600081525090565b8280546114cd90611665565b90600052602060002090601f0160209004810192826114ef5760008555611535565b82601f1061150857805160ff1916838001178555611535565b82800160010185558215611535579182015b8281111561153557825182559160200191906001019061151a565b50611541929150611545565b5090565b5b808211156115415760008155600101611546565b60006020828403121561156c57600080fd5b5035919050565b6001600160a01b038116811461158857600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156115b457600080fd5b82356115bf81611573565b9150602083013567ffffffffffffffff808211156115dc57600080fd5b818501915085601f8301126115f057600080fd5b8135818111156116025761160261158b565b604051601f8201601f19908116603f0116810190838211818310171561162a5761162a61158b565b8160405282815288602084870101111561164357600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600181811c9082168061167957607f821691505b6020821081141561169a57634e487b7160e01b600052602260045260246000fd5b50919050565b60005b838110156116bb5781810151838201526020016116a3565b838111156116ca576000848401525b50505050565b60208152600082518060208401526116ef8160408501602087016116a0565b601f01601f19169190910160400192915050565b60006020828403121561171557600080fd5b815161110681611573565b60006020828403121561173257600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b6000821982111561176257611762611739565b500190565b60008282101561177957611779611739565b500390565b60006020828403121561179057600080fd5b8151801515811461110657600080fd5b600085516117b2818460208a016116a0565b8551908301906117c6818360208a016116a0565b85519101906117d98183602089016116a0565b84519101906117ec8183602088016116a0565b7339bcb9ba32b6a9bab938363ab9a0bab1ba34b7b760611b91019081526014019695505050505050565b60008551611828818460208a016116a0565b85519083019061183c818360208a016116a0565b855191019061184f8183602089016116a0565b84519101906118628183602088016116a0565b713634b8bab4b230ba34b7b720bab1ba34b7b760711b91019081526012019695505050505050565b6000855161189c818460208a016116a0565b8551908301906118b0818360208a016116a0565b85519101906118c38183602089016116a0565b84519101906118d68183602088016116a0565b65736176696e6760d01b91019081526006019695505050505050565b60008551611904818460208a016116a0565b855190830190611918818360208a016116a0565b855191019061192b8183602089016116a0565b845191019061193e8183602088016116a0565b7039bcb9ba32b6a232b13a20bab1ba34b7b760791b91019081526011019695505050505050565b60008551611977818460208a016116a0565b85519083019061198b818360208a016116a0565b855191019061199e8183602089016116a0565b84519101906119b18183602088016116a0565b6331b7b4b760e11b91019081526004019695505050505050565b600085516119dd818460208a016116a0565b8551908301906119f1818360208a016116a0565b8551910190611a048183602089016116a0565b8451910190611a178183602088016116a0565b01969550505050505056fe676f7665726e616e63652e657870657274732e455044522e706172616d6574657273a2646970667358221220eab8cf594047de4c1aa96c079a3de9cbd8e3985de794e0eada9651b6bdb30bee64736f6c63430008090033",
}

// SystemBalanceABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemBalanceMetaData.ABI instead.
var SystemBalanceABI = SystemBalanceMetaData.ABI

// SystemBalanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemBalanceMetaData.Bin instead.
var SystemBalanceBin = SystemBalanceMetaData.Bin

// DeploySystemBalance deploys a new Ethereum contract, binding an instance of SystemBalance to it.
func DeploySystemBalance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemBalance, error) {
	parsed, err := SystemBalanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemBalanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemBalance{SystemBalanceCaller: SystemBalanceCaller{contract: contract}, SystemBalanceTransactor: SystemBalanceTransactor{contract: contract}, SystemBalanceFilterer: SystemBalanceFilterer{contract: contract}}, nil
}

// SystemBalance is an auto generated Go binding around an Ethereum contract.
type SystemBalance struct {
	SystemBalanceCaller     // Read-only binding to the contract
	SystemBalanceTransactor // Write-only binding to the contract
	SystemBalanceFilterer   // Log filterer for contract events
}

// SystemBalanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemBalanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemBalanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemBalanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemBalanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemBalanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemBalanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemBalanceSession struct {
	Contract     *SystemBalance    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemBalanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemBalanceCallerSession struct {
	Contract *SystemBalanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SystemBalanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemBalanceTransactorSession struct {
	Contract     *SystemBalanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SystemBalanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemBalanceRaw struct {
	Contract *SystemBalance // Generic contract binding to access the raw methods on
}

// SystemBalanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemBalanceCallerRaw struct {
	Contract *SystemBalanceCaller // Generic read-only contract binding to access the raw methods on
}

// SystemBalanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemBalanceTransactorRaw struct {
	Contract *SystemBalanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemBalance creates a new instance of SystemBalance, bound to a specific deployed contract.
func NewSystemBalance(address common.Address, backend bind.ContractBackend) (*SystemBalance, error) {
	contract, err := bindSystemBalance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemBalance{SystemBalanceCaller: SystemBalanceCaller{contract: contract}, SystemBalanceTransactor: SystemBalanceTransactor{contract: contract}, SystemBalanceFilterer: SystemBalanceFilterer{contract: contract}}, nil
}

// NewSystemBalanceCaller creates a new read-only instance of SystemBalance, bound to a specific deployed contract.
func NewSystemBalanceCaller(address common.Address, caller bind.ContractCaller) (*SystemBalanceCaller, error) {
	contract, err := bindSystemBalance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemBalanceCaller{contract: contract}, nil
}

// NewSystemBalanceTransactor creates a new write-only instance of SystemBalance, bound to a specific deployed contract.
func NewSystemBalanceTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemBalanceTransactor, error) {
	contract, err := bindSystemBalance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemBalanceTransactor{contract: contract}, nil
}

// NewSystemBalanceFilterer creates a new log filterer instance of SystemBalance, bound to a specific deployed contract.
func NewSystemBalanceFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemBalanceFilterer, error) {
	contract, err := bindSystemBalance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemBalanceFilterer{contract: contract}, nil
}

// bindSystemBalance binds a generic wrapper to an already deployed contract.
func bindSystemBalance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemBalanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemBalance *SystemBalanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemBalance.Contract.SystemBalanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemBalance *SystemBalanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemBalance.Contract.SystemBalanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemBalance *SystemBalanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemBalance.Contract.SystemBalanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemBalance *SystemBalanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemBalance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemBalance *SystemBalanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemBalance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemBalance *SystemBalanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemBalance.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(int256)
func (_SystemBalance *SystemBalanceCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemBalance.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(int256)
func (_SystemBalance *SystemBalanceSession) GetBalance() (*big.Int, error) {
	return _SystemBalance.Contract.GetBalance(&_SystemBalance.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(int256)
func (_SystemBalance *SystemBalanceCallerSession) GetBalance() (*big.Int, error) {
	return _SystemBalance.Contract.GetBalance(&_SystemBalance.CallOpts)
}

// GetBalanceDetails is a free data retrieval call binding the contract method 0x64eb4369.
//
// Solidity: function getBalanceDetails() view returns((bool,bool,uint256,uint256,uint256,uint256))
func (_SystemBalance *SystemBalanceCaller) GetBalanceDetails(opts *bind.CallOpts) (SystemBalanceSystemBalanceDetails, error) {
	var out []interface{}
	err := _SystemBalance.contract.Call(opts, &out, "getBalanceDetails")

	if err != nil {
		return *new(SystemBalanceSystemBalanceDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(SystemBalanceSystemBalanceDetails)).(*SystemBalanceSystemBalanceDetails)

	return out0, err

}

// GetBalanceDetails is a free data retrieval call binding the contract method 0x64eb4369.
//
// Solidity: function getBalanceDetails() view returns((bool,bool,uint256,uint256,uint256,uint256))
func (_SystemBalance *SystemBalanceSession) GetBalanceDetails() (SystemBalanceSystemBalanceDetails, error) {
	return _SystemBalance.Contract.GetBalanceDetails(&_SystemBalance.CallOpts)
}

// GetBalanceDetails is a free data retrieval call binding the contract method 0x64eb4369.
//
// Solidity: function getBalanceDetails() view returns((bool,bool,uint256,uint256,uint256,uint256))
func (_SystemBalance *SystemBalanceCallerSession) GetBalanceDetails() (SystemBalanceSystemBalanceDetails, error) {
	return _SystemBalance.Contract.GetBalanceDetails(&_SystemBalance.CallOpts)
}

// GetDebt is a free data retrieval call binding the contract method 0x14a6bf0f.
//
// Solidity: function getDebt() view returns(uint256)
func (_SystemBalance *SystemBalanceCaller) GetDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemBalance.contract.Call(opts, &out, "getDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDebt is a free data retrieval call binding the contract method 0x14a6bf0f.
//
// Solidity: function getDebt() view returns(uint256)
func (_SystemBalance *SystemBalanceSession) GetDebt() (*big.Int, error) {
	return _SystemBalance.Contract.GetDebt(&_SystemBalance.CallOpts)
}

// GetDebt is a free data retrieval call binding the contract method 0x14a6bf0f.
//
// Solidity: function getDebt() view returns(uint256)
func (_SystemBalance *SystemBalanceCallerSession) GetDebt() (*big.Int, error) {
	return _SystemBalance.Contract.GetDebt(&_SystemBalance.CallOpts)
}

// GetSurplus is a free data retrieval call binding the contract method 0x2383b074.
//
// Solidity: function getSurplus() view returns(uint256)
func (_SystemBalance *SystemBalanceCaller) GetSurplus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemBalance.contract.Call(opts, &out, "getSurplus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSurplus is a free data retrieval call binding the contract method 0x2383b074.
//
// Solidity: function getSurplus() view returns(uint256)
func (_SystemBalance *SystemBalanceSession) GetSurplus() (*big.Int, error) {
	return _SystemBalance.Contract.GetSurplus(&_SystemBalance.CallOpts)
}

// GetSurplus is a free data retrieval call binding the contract method 0x2383b074.
//
// Solidity: function getSurplus() view returns(uint256)
func (_SystemBalance *SystemBalanceCallerSession) GetSurplus() (*big.Int, error) {
	return _SystemBalance.Contract.GetSurplus(&_SystemBalance.CallOpts)
}

// IncreaseDebt is a paid mutator transaction binding the contract method 0x2b7c7b11.
//
// Solidity: function increaseDebt(uint256 _debtAmount) returns(bool)
func (_SystemBalance *SystemBalanceTransactor) IncreaseDebt(opts *bind.TransactOpts, _debtAmount *big.Int) (*types.Transaction, error) {
	return _SystemBalance.contract.Transact(opts, "increaseDebt", _debtAmount)
}

// IncreaseDebt is a paid mutator transaction binding the contract method 0x2b7c7b11.
//
// Solidity: function increaseDebt(uint256 _debtAmount) returns(bool)
func (_SystemBalance *SystemBalanceSession) IncreaseDebt(_debtAmount *big.Int) (*types.Transaction, error) {
	return _SystemBalance.Contract.IncreaseDebt(&_SystemBalance.TransactOpts, _debtAmount)
}

// IncreaseDebt is a paid mutator transaction binding the contract method 0x2b7c7b11.
//
// Solidity: function increaseDebt(uint256 _debtAmount) returns(bool)
func (_SystemBalance *SystemBalanceTransactorSession) IncreaseDebt(_debtAmount *big.Int) (*types.Transaction, error) {
	return _SystemBalance.Contract.IncreaseDebt(&_SystemBalance.TransactOpts, _debtAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemBalance *SystemBalanceTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemBalance.contract.Transact(opts, "initialize", _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemBalance *SystemBalanceSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemBalance.Contract.Initialize(&_SystemBalance.TransactOpts, _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemBalance *SystemBalanceTransactorSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemBalance.Contract.Initialize(&_SystemBalance.TransactOpts, _registry, _stc)
}

// PerformNetting is a paid mutator transaction binding the contract method 0x3d03f797.
//
// Solidity: function performNetting() returns(bool)
func (_SystemBalance *SystemBalanceTransactor) PerformNetting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemBalance.contract.Transact(opts, "performNetting")
}

// PerformNetting is a paid mutator transaction binding the contract method 0x3d03f797.
//
// Solidity: function performNetting() returns(bool)
func (_SystemBalance *SystemBalanceSession) PerformNetting() (*types.Transaction, error) {
	return _SystemBalance.Contract.PerformNetting(&_SystemBalance.TransactOpts)
}

// PerformNetting is a paid mutator transaction binding the contract method 0x3d03f797.
//
// Solidity: function performNetting() returns(bool)
func (_SystemBalance *SystemBalanceTransactorSession) PerformNetting() (*types.Transaction, error) {
	return _SystemBalance.Contract.PerformNetting(&_SystemBalance.TransactOpts)
}

// TransferAccruedInterestAmount is a paid mutator transaction binding the contract method 0xeffd85ad.
//
// Solidity: function transferAccruedInterestAmount(uint256 _amount) returns(bool)
func (_SystemBalance *SystemBalanceTransactor) TransferAccruedInterestAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SystemBalance.contract.Transact(opts, "transferAccruedInterestAmount", _amount)
}

// TransferAccruedInterestAmount is a paid mutator transaction binding the contract method 0xeffd85ad.
//
// Solidity: function transferAccruedInterestAmount(uint256 _amount) returns(bool)
func (_SystemBalance *SystemBalanceSession) TransferAccruedInterestAmount(_amount *big.Int) (*types.Transaction, error) {
	return _SystemBalance.Contract.TransferAccruedInterestAmount(&_SystemBalance.TransactOpts, _amount)
}

// TransferAccruedInterestAmount is a paid mutator transaction binding the contract method 0xeffd85ad.
//
// Solidity: function transferAccruedInterestAmount(uint256 _amount) returns(bool)
func (_SystemBalance *SystemBalanceTransactorSession) TransferAccruedInterestAmount(_amount *big.Int) (*types.Transaction, error) {
	return _SystemBalance.Contract.TransferAccruedInterestAmount(&_SystemBalance.TransactOpts, _amount)
}

// TransferSurplusAuctionAmount is a paid mutator transaction binding the contract method 0x152353d9.
//
// Solidity: function transferSurplusAuctionAmount(uint256 _lot) returns(bool)
func (_SystemBalance *SystemBalanceTransactor) TransferSurplusAuctionAmount(opts *bind.TransactOpts, _lot *big.Int) (*types.Transaction, error) {
	return _SystemBalance.contract.Transact(opts, "transferSurplusAuctionAmount", _lot)
}

// TransferSurplusAuctionAmount is a paid mutator transaction binding the contract method 0x152353d9.
//
// Solidity: function transferSurplusAuctionAmount(uint256 _lot) returns(bool)
func (_SystemBalance *SystemBalanceSession) TransferSurplusAuctionAmount(_lot *big.Int) (*types.Transaction, error) {
	return _SystemBalance.Contract.TransferSurplusAuctionAmount(&_SystemBalance.TransactOpts, _lot)
}

// TransferSurplusAuctionAmount is a paid mutator transaction binding the contract method 0x152353d9.
//
// Solidity: function transferSurplusAuctionAmount(uint256 _lot) returns(bool)
func (_SystemBalance *SystemBalanceTransactorSession) TransferSurplusAuctionAmount(_lot *big.Int) (*types.Transaction, error) {
	return _SystemBalance.Contract.TransferSurplusAuctionAmount(&_SystemBalance.TransactOpts, _lot)
}
