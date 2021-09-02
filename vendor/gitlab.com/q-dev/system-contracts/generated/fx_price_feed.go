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

// FxPriceFeedMetaData contains all meta data concerning the FxPriceFeed contract.
var FxPriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pair\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_decimal\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_maintainersList\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_baseTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quoteTokenAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimalPlaces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pair\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tstamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maintainer\",\"type\":\"address\"}],\"name\":\"setMaintainer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leaveMaintainers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exchangeRate\",\"type\":\"uint256\"}],\"name\":\"setExchangeRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaintainers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001c2438038062001c24833981810160405260a08110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b506040818152602083015192018051929491939192846401000000008211156200010f57600080fd5b9083019060208201858111156200012557600080fd5b82518660208202830111640100000000821117156200014357600080fd5b82525081516020918201928201910280838360005b838110156200017257818101518382015260200162000158565b505050509190910160409081526020830151920151919350909150600090506200019b6200043f565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3508451620002245760405162461bcd60e51b815260040180806020018281038252603581526020018062001bb76035913960400191505060405180910390fd5b60008411620002655760405162461bcd60e51b815260040180806020018281038252603881526020018062001bec6038913960400191505060405180910390fd5b6001600160a01b038216620002ac5760405162461bcd60e51b815260040180806020018281038252604381526020018062001ad16043913960600191505060405180910390fd5b6001600160a01b038116620002f35760405162461bcd60e51b815260040180806020018281038252604481526020018062001a8d6044913960600191505060405180910390fd5b620002fe8262000443565b6200033b5760405162461bcd60e51b815260040180806020018281038252605181526020018062001b666051913960600191505060405180910390fd5b620003468162000443565b620003835760405162461bcd60e51b815260040180806020018281038252605281526020018062001b146052913960600191505060405180910390fd5b8260405162000392906200044f565b6020808252825181830152825182916040830191858201910280838360005b83811015620003cb578181015183820152602001620003b1565b5050505090500192505050604051809103906000f080158015620003f3573d6000803e3d6000fd5b50600180546001600160a01b03199081166001600160a01b03938416179091556005959095556003805486169382169390931790925560048054909416911617909155506200045d9050565b3390565b3b63ffffffff16151590565b610cb08062000ddd83390190565b610970806200046d6000396000f3fe608060405234801561001057600080fd5b50600436106100af5760003560e01c80630181390c146100b457806313ea5d29146100ce5780633ba0b9a9146100f6578063715018a6146100fe5780637acede8b146101065780638da5cb5b1461012a578063a094600e14610132578063a8aa1b311461013a578063d17f4225146101b7578063d52040331461020f578063db068e0e14610217578063e1725c9214610248578063f2fde38b14610250575b600080fd5b6100bc610276565b60408051918252519081900360200190f35b6100f4600480360360208110156100e457600080fd5b50356001600160a01b031661027c565b005b6100bc610397565b6100f461039d565b61010e610451565b604080516001600160a01b039092168252519081900360200190f35b61010e610460565b61010e61046f565b61014261047e565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561017c578181015183820152602001610164565b50505050905090810190601f1680156101a95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101bf610509565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101fb5781810151838201526020016101e3565b505050509050019250505060405180910390f35b6100f461061a565b6102346004803603602081101561022d57600080fd5b5035610680565b604080519115158252519081900360200190f35b6100bc61077e565b6100f46004803603602081101561026657600080fd5b50356001600160a01b0316610784565b60065481565b60015460408051630bb7c8fd60e31b815233600482015290516001600160a01b0390921691635dbe47e891602480820192602092909190829003018186803b1580156102c757600080fd5b505afa1580156102db573d6000803e3d6000fd5b505050506040513d60208110156102f157600080fd5b505161032e5760405162461bcd60e51b815260040180806020018281038252603e8152602001806108b9603e913960400191505060405180910390fd5b60015460408051639e9405c360e01b81526001600160a01b03848116600483015291519190921691639e9405c391602480830192600092919082900301818387803b15801561037c57600080fd5b505af1158015610390573d6000803e3d6000fd5b5050505050565b60075481565b6103a561088e565b6000546001600160a01b03908116911614610407576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6004546001600160a01b031681565b6000546001600160a01b031690565b6003546001600160a01b031681565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156105015780601f106104d657610100808354040283529160200191610501565b820191906000526020600020905b8154815290600101906020018083116104e457829003601f168201915b505050505081565b600154604080516351cfd60960e11b815290516060926001600160a01b03169163a39fac12916004808301926000929190829003018186803b15801561054e57600080fd5b505afa158015610562573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561058b57600080fd5b8101908080516040519392919084600160201b8211156105aa57600080fd5b9083019060208201858111156105bf57600080fd5b82518660208202830111600160201b821117156105db57600080fd5b82525081516020918201928201910280838360005b838110156106085781810151838201526020016105f0565b50505050905001604052505050905090565b60015460408051636196c02d60e11b815233600482015290516001600160a01b039092169163c32d805a9160248082019260009290919082900301818387803b15801561066657600080fd5b505af115801561067a573d6000803e3d6000fd5b50505050565b60015460408051630bb7c8fd60e31b815233600482015290516000926001600160a01b031691635dbe47e8916024808301926020929190829003018186803b1580156106cb57600080fd5b505afa1580156106df573d6000803e3d6000fd5b505050506040513d60208110156106f557600080fd5b50516107325760405162461bcd60e51b815260040180806020018281038252603e8152602001806108b9603e913960400191505060405180910390fd5b600082116107715760405162461bcd60e51b81526004018080602001828103825260448152602001806108f76044913960600191505060405180910390fd5b5060075542600655600190565b60055481565b61078c61088e565b6000546001600160a01b039081169116146107ee576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b0381166108335760405162461bcd60e51b81526004018080602001828103825260268152602001806108936026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573735b5145432d3031393030305d2d5065726d697373696f6e2064656e696564202d206f6e6c79206d61696e7461696e6572732068617665206163636573732e5b5145432d3031393030375d2d496e76616c69642065786368616e676520726174652c206661696c656420746f20736574207468652065786368616e676520726174652ea2646970667358221220ed238da104fa7dbed28c8dd626cb196acb5bbd546ade08866298e403f7b7592364736f6c63430007060033608060405234801561001057600080fd5b50604051610cb0380380610cb08339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b5050505090500160405250505060006100cf6101b560201b60201c565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060005b81518110156101ae57600282828151811061013357fe5b60209081029190910181015182546001808201855560009485529284200180546001600160a01b0319166001600160a01b03909216919091179055600254845190929085908590811061018257fe5b6020908102919091018101516001600160a01b031682528101919091526040016000205560010161011c565b50506101b9565b3390565b610ae8806101c86000396000f3fe608060405234801561001057600080fd5b50600436106100995760003560e01c80630a3b0a4f1461009e57806329092d0e146100d857806352efea6e146100fe5780635dbe47e814610106578063715018a61461012c5780638da5cb5b14610136578063949d225d1461015a5780639e9405c314610174578063a39fac121461019a578063c32d805a146101f2578063f2fde38b14610218575b600080fd5b6100c4600480360360208110156100b457600080fd5b50356001600160a01b031661023e565b604080519115158252519081900360200190f35b6100c4600480360360208110156100ee57600080fd5b50356001600160a01b03166102d0565b6100c4610372565b6100c46004803603602081101561011c57600080fd5b50356001600160a01b031661042a565b610134610447565b005b61013e6104e9565b604080516001600160a01b039092168252519081900360200190f35b6101626104f8565b60408051918252519081900360200190f35b6101346004803603602081101561018a57600080fd5b50356001600160a01b03166104fe565b6101a261059d565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101de5781810151838201526020016101c6565b505050509050019250505060405180910390f35b6101346004803603602081101561020857600080fd5b50356001600160a01b03166105ff565b6101346004803603602081101561022e57600080fd5b50356001600160a01b031661069b565b6000610248610793565b6000546001600160a01b03908116911614610298576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b6001600160a01b038216600090815260016020526040902054156102be575060006102cb565b6102c782610797565b5060015b919050565b60006102da610793565b6000546001600160a01b0390811691161461032a576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b6001600160a01b038216600090815260016020526040902054801580610351575060025481115b156103605760009150506102cb565b610369836107fe565b50600192915050565b600061037c610793565b6000546001600160a01b039081169116146103cc576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b60005b6002548110156104175760016000600283815481106103ea57fe5b60009182526020808320909101546001600160a01b031683528201929092526040018120556001016103cf565b5061042460026000610986565b50600190565b6001600160a01b0316600090815260016020526040902054151590565b61044f610793565b6000546001600160a01b0390811691161461049f576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b60025490565b610506610793565b6000546001600160a01b03908116911614610556576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b61055f8161023e565b61059a5760405162461bcd60e51b8152600401808060200182810382526071815260200180610a226071913960800191505060405180910390fd5b50565b606060028054806020026020016040519081016040528092919081815260200182805480156105f557602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116105d7575b5050505050905090565b610607610793565b6000546001600160a01b03908116911614610657576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b610660816102d0565b61059a5760405162461bcd60e51b81526004018080602001828103825260438152602001806109df6043913960600191505060405180910390fd5b6106a3610793565b6000546001600160a01b039081169116146106f3576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b6001600160a01b0381166107385760405162461bcd60e51b81526004018080602001828103825260268152602001806109b96026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b60028054600180820183557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace90910180546001600160a01b0319166001600160a01b0385169081179091559154600092835260209190915260409091205561059a81610938565b6001600160a01b038116600090815260016020908152604080832081519283019091525481526002805491926000198301929091908390811061083d57fe5b60009182526020909120015483516001600160a01b0390911691506000190182146108c15782516001600160a01b03821660009081526001602052604090205582516002805483926000190190811061089257fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b60028054806108cc57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038616825260019052604081205561091084610938565b836001600160a01b0316816001600160a01b0316146109325761093281610938565b50505050565b6001600160a01b03811660009081526001602052604090205460025481111561095d57fe5b6109668261042a565b1561097a576000811161097557fe5b610982565b801561098257fe5b5050565b508054600082559060005260206000209081019061059a91905b808211156109b457600081556001016109a0565b509056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573735b5145432d3033353030305d2d4661696c656420746f2072656d6f76652074686520616464726573732066726f6d2074686520616464726573732073746f726167652e5b5145432d3033353030325d2d54686520616464726573732068617320616c7265616479206265656e20616464656420746f207468652073746f726167652c206661696c656420746f2061646420746865206164647265737320746f2074686520616464726573732073746f726167652e4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a26469706673582212203421bef274aa9bcf497e4b6a0bb1e7be50a80a19ba4c0e2f760c17b87dc2d41c64736f6c634300070600335b5145432d3031393030345d2d496e76616c69642071756f746520746f6b656e20616464726573732c2074686520696e697469616c697a6174696f6e206661696c65642e5b5145432d3031393030335d2d496e76616c6964206261736520746f6b656e20616464726573732c2074686520696e697469616c697a6174696f6e206661696c65642e5b5145432d3031393030365d2d5468652071756f746520746f6b656e2061646472657373206973206e6f74206120636f6e74726163742c2074686520696e697469616c697a6174696f6e206661696c65642e5b5145432d3031393030355d2d546865206261736520746f6b656e2061646472657373206973206e6f74206120636f6e74726163742c2074686520696e697469616c697a6174696f6e206661696c65642e5b5145432d3031393030315d2d496e76616c696420706169722c2074686520696e697469616c697a6174696f6e206661696c65642e5b5145432d3031393030325d2d496e76616c696420646563696d616c2c2074686520696e697469616c697a6174696f6e206661696c65642e",
}

// FxPriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use FxPriceFeedMetaData.ABI instead.
var FxPriceFeedABI = FxPriceFeedMetaData.ABI

// FxPriceFeedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FxPriceFeedMetaData.Bin instead.
var FxPriceFeedBin = FxPriceFeedMetaData.Bin

// DeployFxPriceFeed deploys a new Ethereum contract, binding an instance of FxPriceFeed to it.
func DeployFxPriceFeed(auth *bind.TransactOpts, backend bind.ContractBackend, _pair string, _decimal *big.Int, _maintainersList []common.Address, _baseTokenAddr common.Address, _quoteTokenAddr common.Address) (common.Address, *types.Transaction, *FxPriceFeed, error) {
	parsed, err := FxPriceFeedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FxPriceFeedBin), backend, _pair, _decimal, _maintainersList, _baseTokenAddr, _quoteTokenAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FxPriceFeed{FxPriceFeedCaller: FxPriceFeedCaller{contract: contract}, FxPriceFeedTransactor: FxPriceFeedTransactor{contract: contract}, FxPriceFeedFilterer: FxPriceFeedFilterer{contract: contract}}, nil
}

// FxPriceFeed is an auto generated Go binding around an Ethereum contract.
type FxPriceFeed struct {
	FxPriceFeedCaller     // Read-only binding to the contract
	FxPriceFeedTransactor // Write-only binding to the contract
	FxPriceFeedFilterer   // Log filterer for contract events
}

// FxPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type FxPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FxPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FxPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FxPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FxPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FxPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FxPriceFeedSession struct {
	Contract     *FxPriceFeed      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FxPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FxPriceFeedCallerSession struct {
	Contract *FxPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FxPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FxPriceFeedTransactorSession struct {
	Contract     *FxPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FxPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type FxPriceFeedRaw struct {
	Contract *FxPriceFeed // Generic contract binding to access the raw methods on
}

// FxPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FxPriceFeedCallerRaw struct {
	Contract *FxPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// FxPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FxPriceFeedTransactorRaw struct {
	Contract *FxPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFxPriceFeed creates a new instance of FxPriceFeed, bound to a specific deployed contract.
func NewFxPriceFeed(address common.Address, backend bind.ContractBackend) (*FxPriceFeed, error) {
	contract, err := bindFxPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FxPriceFeed{FxPriceFeedCaller: FxPriceFeedCaller{contract: contract}, FxPriceFeedTransactor: FxPriceFeedTransactor{contract: contract}, FxPriceFeedFilterer: FxPriceFeedFilterer{contract: contract}}, nil
}

// NewFxPriceFeedCaller creates a new read-only instance of FxPriceFeed, bound to a specific deployed contract.
func NewFxPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*FxPriceFeedCaller, error) {
	contract, err := bindFxPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FxPriceFeedCaller{contract: contract}, nil
}

// NewFxPriceFeedTransactor creates a new write-only instance of FxPriceFeed, bound to a specific deployed contract.
func NewFxPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*FxPriceFeedTransactor, error) {
	contract, err := bindFxPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FxPriceFeedTransactor{contract: contract}, nil
}

// NewFxPriceFeedFilterer creates a new log filterer instance of FxPriceFeed, bound to a specific deployed contract.
func NewFxPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*FxPriceFeedFilterer, error) {
	contract, err := bindFxPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FxPriceFeedFilterer{contract: contract}, nil
}

// bindFxPriceFeed binds a generic wrapper to an already deployed contract.
func bindFxPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FxPriceFeedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FxPriceFeed *FxPriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FxPriceFeed.Contract.FxPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FxPriceFeed *FxPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.FxPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FxPriceFeed *FxPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.FxPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FxPriceFeed *FxPriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FxPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FxPriceFeed *FxPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FxPriceFeed *FxPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// BaseTokenAddr is a free data retrieval call binding the contract method 0xa094600e.
//
// Solidity: function baseTokenAddr() view returns(address)
func (_FxPriceFeed *FxPriceFeedCaller) BaseTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FxPriceFeed.contract.Call(opts, &out, "baseTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseTokenAddr is a free data retrieval call binding the contract method 0xa094600e.
//
// Solidity: function baseTokenAddr() view returns(address)
func (_FxPriceFeed *FxPriceFeedSession) BaseTokenAddr() (common.Address, error) {
	return _FxPriceFeed.Contract.BaseTokenAddr(&_FxPriceFeed.CallOpts)
}

// BaseTokenAddr is a free data retrieval call binding the contract method 0xa094600e.
//
// Solidity: function baseTokenAddr() view returns(address)
func (_FxPriceFeed *FxPriceFeedCallerSession) BaseTokenAddr() (common.Address, error) {
	return _FxPriceFeed.Contract.BaseTokenAddr(&_FxPriceFeed.CallOpts)
}

// DecimalPlaces is a free data retrieval call binding the contract method 0xe1725c92.
//
// Solidity: function decimalPlaces() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedCaller) DecimalPlaces(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FxPriceFeed.contract.Call(opts, &out, "decimalPlaces")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalPlaces is a free data retrieval call binding the contract method 0xe1725c92.
//
// Solidity: function decimalPlaces() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedSession) DecimalPlaces() (*big.Int, error) {
	return _FxPriceFeed.Contract.DecimalPlaces(&_FxPriceFeed.CallOpts)
}

// DecimalPlaces is a free data retrieval call binding the contract method 0xe1725c92.
//
// Solidity: function decimalPlaces() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedCallerSession) DecimalPlaces() (*big.Int, error) {
	return _FxPriceFeed.Contract.DecimalPlaces(&_FxPriceFeed.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FxPriceFeed.contract.Call(opts, &out, "exchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedSession) ExchangeRate() (*big.Int, error) {
	return _FxPriceFeed.Contract.ExchangeRate(&_FxPriceFeed.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedCallerSession) ExchangeRate() (*big.Int, error) {
	return _FxPriceFeed.Contract.ExchangeRate(&_FxPriceFeed.CallOpts)
}

// GetMaintainers is a free data retrieval call binding the contract method 0xd17f4225.
//
// Solidity: function getMaintainers() view returns(address[])
func (_FxPriceFeed *FxPriceFeedCaller) GetMaintainers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FxPriceFeed.contract.Call(opts, &out, "getMaintainers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMaintainers is a free data retrieval call binding the contract method 0xd17f4225.
//
// Solidity: function getMaintainers() view returns(address[])
func (_FxPriceFeed *FxPriceFeedSession) GetMaintainers() ([]common.Address, error) {
	return _FxPriceFeed.Contract.GetMaintainers(&_FxPriceFeed.CallOpts)
}

// GetMaintainers is a free data retrieval call binding the contract method 0xd17f4225.
//
// Solidity: function getMaintainers() view returns(address[])
func (_FxPriceFeed *FxPriceFeedCallerSession) GetMaintainers() ([]common.Address, error) {
	return _FxPriceFeed.Contract.GetMaintainers(&_FxPriceFeed.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FxPriceFeed *FxPriceFeedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FxPriceFeed.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FxPriceFeed *FxPriceFeedSession) Owner() (common.Address, error) {
	return _FxPriceFeed.Contract.Owner(&_FxPriceFeed.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FxPriceFeed *FxPriceFeedCallerSession) Owner() (common.Address, error) {
	return _FxPriceFeed.Contract.Owner(&_FxPriceFeed.CallOpts)
}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(string)
func (_FxPriceFeed *FxPriceFeedCaller) Pair(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FxPriceFeed.contract.Call(opts, &out, "pair")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(string)
func (_FxPriceFeed *FxPriceFeedSession) Pair() (string, error) {
	return _FxPriceFeed.Contract.Pair(&_FxPriceFeed.CallOpts)
}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(string)
func (_FxPriceFeed *FxPriceFeedCallerSession) Pair() (string, error) {
	return _FxPriceFeed.Contract.Pair(&_FxPriceFeed.CallOpts)
}

// QuoteTokenAddr is a free data retrieval call binding the contract method 0x7acede8b.
//
// Solidity: function quoteTokenAddr() view returns(address)
func (_FxPriceFeed *FxPriceFeedCaller) QuoteTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FxPriceFeed.contract.Call(opts, &out, "quoteTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteTokenAddr is a free data retrieval call binding the contract method 0x7acede8b.
//
// Solidity: function quoteTokenAddr() view returns(address)
func (_FxPriceFeed *FxPriceFeedSession) QuoteTokenAddr() (common.Address, error) {
	return _FxPriceFeed.Contract.QuoteTokenAddr(&_FxPriceFeed.CallOpts)
}

// QuoteTokenAddr is a free data retrieval call binding the contract method 0x7acede8b.
//
// Solidity: function quoteTokenAddr() view returns(address)
func (_FxPriceFeed *FxPriceFeedCallerSession) QuoteTokenAddr() (common.Address, error) {
	return _FxPriceFeed.Contract.QuoteTokenAddr(&_FxPriceFeed.CallOpts)
}

// Tstamp is a free data retrieval call binding the contract method 0x0181390c.
//
// Solidity: function tstamp() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedCaller) Tstamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FxPriceFeed.contract.Call(opts, &out, "tstamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tstamp is a free data retrieval call binding the contract method 0x0181390c.
//
// Solidity: function tstamp() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedSession) Tstamp() (*big.Int, error) {
	return _FxPriceFeed.Contract.Tstamp(&_FxPriceFeed.CallOpts)
}

// Tstamp is a free data retrieval call binding the contract method 0x0181390c.
//
// Solidity: function tstamp() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedCallerSession) Tstamp() (*big.Int, error) {
	return _FxPriceFeed.Contract.Tstamp(&_FxPriceFeed.CallOpts)
}

// LeaveMaintainers is a paid mutator transaction binding the contract method 0xd5204033.
//
// Solidity: function leaveMaintainers() returns()
func (_FxPriceFeed *FxPriceFeedTransactor) LeaveMaintainers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FxPriceFeed.contract.Transact(opts, "leaveMaintainers")
}

// LeaveMaintainers is a paid mutator transaction binding the contract method 0xd5204033.
//
// Solidity: function leaveMaintainers() returns()
func (_FxPriceFeed *FxPriceFeedSession) LeaveMaintainers() (*types.Transaction, error) {
	return _FxPriceFeed.Contract.LeaveMaintainers(&_FxPriceFeed.TransactOpts)
}

// LeaveMaintainers is a paid mutator transaction binding the contract method 0xd5204033.
//
// Solidity: function leaveMaintainers() returns()
func (_FxPriceFeed *FxPriceFeedTransactorSession) LeaveMaintainers() (*types.Transaction, error) {
	return _FxPriceFeed.Contract.LeaveMaintainers(&_FxPriceFeed.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FxPriceFeed *FxPriceFeedTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FxPriceFeed.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FxPriceFeed *FxPriceFeedSession) RenounceOwnership() (*types.Transaction, error) {
	return _FxPriceFeed.Contract.RenounceOwnership(&_FxPriceFeed.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FxPriceFeed *FxPriceFeedTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FxPriceFeed.Contract.RenounceOwnership(&_FxPriceFeed.TransactOpts)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 _exchangeRate) returns(bool)
func (_FxPriceFeed *FxPriceFeedTransactor) SetExchangeRate(opts *bind.TransactOpts, _exchangeRate *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.contract.Transact(opts, "setExchangeRate", _exchangeRate)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 _exchangeRate) returns(bool)
func (_FxPriceFeed *FxPriceFeedSession) SetExchangeRate(_exchangeRate *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetExchangeRate(&_FxPriceFeed.TransactOpts, _exchangeRate)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 _exchangeRate) returns(bool)
func (_FxPriceFeed *FxPriceFeedTransactorSession) SetExchangeRate(_exchangeRate *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetExchangeRate(&_FxPriceFeed.TransactOpts, _exchangeRate)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address _maintainer) returns()
func (_FxPriceFeed *FxPriceFeedTransactor) SetMaintainer(opts *bind.TransactOpts, _maintainer common.Address) (*types.Transaction, error) {
	return _FxPriceFeed.contract.Transact(opts, "setMaintainer", _maintainer)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address _maintainer) returns()
func (_FxPriceFeed *FxPriceFeedSession) SetMaintainer(_maintainer common.Address) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetMaintainer(&_FxPriceFeed.TransactOpts, _maintainer)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address _maintainer) returns()
func (_FxPriceFeed *FxPriceFeedTransactorSession) SetMaintainer(_maintainer common.Address) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetMaintainer(&_FxPriceFeed.TransactOpts, _maintainer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FxPriceFeed *FxPriceFeedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FxPriceFeed.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FxPriceFeed *FxPriceFeedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.TransferOwnership(&_FxPriceFeed.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FxPriceFeed *FxPriceFeedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.TransferOwnership(&_FxPriceFeed.TransactOpts, newOwner)
}

// FxPriceFeedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FxPriceFeed contract.
type FxPriceFeedOwnershipTransferredIterator struct {
	Event *FxPriceFeedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FxPriceFeedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FxPriceFeedOwnershipTransferred)
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
		it.Event = new(FxPriceFeedOwnershipTransferred)
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
func (it *FxPriceFeedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FxPriceFeedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FxPriceFeedOwnershipTransferred represents a OwnershipTransferred event raised by the FxPriceFeed contract.
type FxPriceFeedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FxPriceFeed *FxPriceFeedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FxPriceFeedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FxPriceFeed.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FxPriceFeedOwnershipTransferredIterator{contract: _FxPriceFeed.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FxPriceFeed *FxPriceFeedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FxPriceFeedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FxPriceFeed.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FxPriceFeedOwnershipTransferred)
				if err := _FxPriceFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FxPriceFeed *FxPriceFeedFilterer) ParseOwnershipTransferred(log types.Log) (*FxPriceFeedOwnershipTransferred, error) {
	event := new(FxPriceFeedOwnershipTransferred)
	if err := _FxPriceFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
