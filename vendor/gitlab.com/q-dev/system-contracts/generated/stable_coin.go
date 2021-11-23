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

// StableCoinMetaData contains all meta data concerning the StableCoin contract.
var StableCoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_eligibleContracts\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recepient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionRecipient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001774380380620017748339810160408190526200003491620002f2565b83838181816003908051906020019062000050929190620000d4565b50805162000066906004906020840190620000d4565b505060058054601260ff1990911617610100600160a81b0319166101006001600160a01b038781169190910291909117909155600680546001600160a01b031916918a1691909117905550508251620000c89150600790602085019062000169565b5050505050506200042e565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826200010c576000855562000157565b82601f106200012757805160ff191683800117855562000157565b8280016001018555821562000157579182015b82811115620001575782518255916020019190600101906200013a565b5062000165929150620001c9565b5090565b828054828255906000526020600020908101928215620001bb579160200282015b82811115620001bb5782518051620001aa918491602090910190620000d4565b50916020019190600101906200018a565b5062000165929150620001e0565b5b80821115620001655760008155600101620001ca565b8082111562000165576000620001f7828262000201565b50600101620001e0565b50805460018160011615610100020316600290046000825580601f1062000229575062000249565b601f016020900490600052602060002090810190620002499190620001c9565b50565b80516001600160a01b03811681146200026457600080fd5b919050565b600082601f8301126200027a578081fd5b81516001600160401b038111156200028e57fe5b6020620002a4601f8301601f191682016200040a565b8281528582848701011115620002b8578384fd5b835b83811015620002d7578581018301518282018401528201620002ba565b83811115620002e857848385840101525b5095945050505050565b600080600080600060a086880312156200030a578081fd5b62000315866200024c565b602087810151919650906001600160401b038082111562000334578384fd5b620003428a838b0162000269565b9650604089015191508082111562000358578384fd5b620003668a838b0162000269565b955060608901519150808211156200037c578384fd5b818901915089601f83011262000390578384fd5b8151818111156200039d57fe5b620003ac84858302016200040a565b8181528481019250838501865b83811015620003e657620003d38e88845189010162000269565b85529386019390860190600101620003b9565b50508096505050505050620003fe608087016200024c565b90509295509295909350565b6040518181016001600160401b03811182821017156200042657fe5b604052919050565b611336806200043e6000396000f3fe608060405234801561001057600080fd5b50600436106100f65760003560e01c8063486ff0cd11610092578063486ff0cd146101c4578063572b6c05146101cc57806370a08231146101df57806379cc6790146101f25780637da0a8771461020557806395d89b411461021a578063a457c2d714610222578063a9059cbb14610235578063dd62ed3e14610248576100f6565b806306fdde03146100fb578063095ea7b31461011957806318160ddd1461013957806323b872dd1461014e578063313ce5671461016157806339509351146101765780634000aea01461018957806340c10f191461019c57806342966c68146101af575b600080fd5b61010361025b565b6040516101109190610ff4565b60405180910390f35b61012c610127366004610edf565b6102f2565b6040516101109190610fe9565b61014161030f565b6040516101109190611140565b61012c61015c366004610e9f565b610315565b61016961039d565b6040516101109190611149565b61012c610184366004610edf565b6103a6565b61012c610197366004610f0a565b6103f4565b61012c6101aa366004610edf565b6105cf565b6101c26101bd366004610fbd565b6106e8565b005b6101036106fc565b61012c6101da366004610e2f565b61071b565b6101416101ed366004610e2f565b610734565b6101c2610200366004610edf565b61074f565b61020d6107a4565b6040516101109190610fd5565b6101036107b8565b61012c610230366004610edf565b610819565b61012c610243366004610edf565b610881565b610141610256366004610e67565b610895565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156102e75780601f106102bc576101008083540402835291602001916102e7565b820191906000526020600020905b8154815290600101906020018083116102ca57829003601f168201915b505050505090505b90565b60006103066102ff6108c0565b84846108cf565b50600192915050565b60025490565b60006103228484846109bb565b6103928461032e6108c0565b61038d85604051806060016040528060288152602001611206602891396001600160a01b038a1660009081526001602052604081209061036c6108c0565b6001600160a01b031681526020810191909152604001600020549190610b04565b6108cf565b5060015b9392505050565b60055460ff1690565b60006103066103b36108c0565b8461038d85600160006103c46108c0565b6001600160a01b03908116825260208083019390935260409182016000908120918c168152925290205490610b9b565b6000806104018585610881565b905080610412576000915050610396565b846001600160a01b03166104246108c0565b6001600160a01b03167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1686866040518083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561049357818101518382015260200161047b565b50505050905090810190601f1680156104c05780820380516001836020036101000a031916815260200191505b50935050505060405180910390a3846001600160a01b03811663c0ee0b8a6104e66108c0565b87876040518463ffffffff1660e01b815260040180846001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610549578181015183820152602001610531565b50505050905090810190601f1680156105765780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561059757600080fd5b505af11580156105ab573d6000803e3d6000fd5b505050506040513d60208110156105c157600080fd5b506001979650505050505050565b60008060005b6007548110156106b7576105e76108c0565b6001600160a01b0316600660009054906101000a90046001600160a01b03166001600160a01b0316633fb902716007848154811061062157fe5b906000526020600020016040518263ffffffff1660e01b81526004016106479190611047565b60206040518083038186803b15801561065f57600080fd5b505afa158015610673573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106979190610e4b565b6001600160a01b031614156106af57600191506106b7565b6001016105d5565b50806106de5760405162461bcd60e51b81526004016106d5906110d1565b60405180910390fd5b6103928484610bf3565b6106f96106f36108c0565b82610cd1565b50565b6040805180820190915260058152640322e322e360dc1b602082015290565b60055461010090046001600160a01b0390811691161490565b6001600160a01b031660009081526020819052604090205490565b60006107818260405180606001604052806024815260200161122e6024913961077a866102566108c0565b9190610b04565b90506107958361078f6108c0565b836108cf565b61079f8383610cd1565b505050565b60055461010090046001600160a01b031681565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156102e75780601f106102bc576101008083540402835291602001916102e7565b60006103066108266108c0565b8461038d856040518060600160405280602581526020016112dc60259139600160006108506108c0565b6001600160a01b03908116825260208083019390935260409182016000908120918d16815292529020549190610b04565b600061030661088e6108c0565b84846109bb565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b60006108ca610dbb565b905090565b6001600160a01b0383166109145760405162461bcd60e51b81526004018080602001828103825260248152602001806112b86024913960400191505060405180910390fd5b6001600160a01b0382166109595760405162461bcd60e51b81526004018080602001828103825260228152602001806111be6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316610a005760405162461bcd60e51b81526004018080602001828103825260258152602001806112936025913960400191505060405180910390fd5b6001600160a01b038216610a455760405162461bcd60e51b81526004018080602001828103825260238152602001806111796023913960400191505060405180910390fd5b610a5083838361079f565b610a8d816040518060600160405280602681526020016111e0602691396001600160a01b0386166000908152602081905260409020549190610b04565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610abc9082610b9b565b6001600160a01b0380841660008181526020818152604091829020949094558051858152905191939287169260008051602061125283398151915292918290030190a3505050565b60008184841115610b935760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610b58578181015183820152602001610b40565b50505050905090810190601f168015610b855780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610396576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b6001600160a01b038216610c4e576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b610c5a6000838361079f565b600254610c679082610b9b565b6002556001600160a01b038216600090815260208190526040902054610c8d9082610b9b565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391926000805160206112528339815191529281900390910190a35050565b6001600160a01b038216610d165760405162461bcd60e51b81526004018080602001828103825260218152602001806112726021913960400191505060405180910390fd5b610d228260008361079f565b610d5f8160405180606001604052806022815260200161119c602291396001600160a01b0385166000908152602081905260409020549190610b04565b6001600160a01b038316600090815260208190526040902055600254610d859082610ded565b6002556040805182815290516000916001600160a01b038516916000805160206112528339815191529181900360200190a35050565b600060143610801590610dd25750610dd23361071b565b15610de6575060131936013560601c6102ef565b50336102ef565b600061039683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610b04565b600060208284031215610e40578081fd5b813561039681611163565b600060208284031215610e5c578081fd5b815161039681611163565b60008060408385031215610e79578081fd5b8235610e8481611163565b91506020830135610e9481611163565b809150509250929050565b600080600060608486031215610eb3578081fd5b8335610ebe81611163565b92506020840135610ece81611163565b929592945050506040919091013590565b60008060408385031215610ef1578182fd5b8235610efc81611163565b946020939093013593505050565b600080600060608486031215610f1e578283fd5b8335610f2981611163565b92506020848101359250604085013567ffffffffffffffff80821115610f4d578384fd5b818701915087601f830112610f60578384fd5b813581811115610f6c57fe5b604051601f8201601f1916810185018381118282101715610f8957fe5b60405281815283820185018a1015610f9f578586fd5b81858501868301378585838301015280955050505050509250925092565b600060208284031215610fce578081fd5b5035919050565b6001600160a01b0391909116815260200190565b901515815260200190565b6000602080835283518082850152825b8181101561102057858101830151858201604001528201611004565b818111156110315783604083870101525b50601f01601f1916929092016040019392505050565b6000602080830181845282855460018082166000811461106e576001811461108c576110c4565b60028304607f16855260ff19831660408901526060880193506110c4565b6002830480865261109c8a611157565b885b828110156110ba5781548b82016040015290840190880161109e565b8a01604001955050505b5091979650505050505050565b60208082526049908201527f5b5145432d3032303030305d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c792074686520656c696769626c6520636f6e74726163747320686176606082015268329030b1b1b2b9b99760b91b608082015260a00190565b90815260200190565b60ff91909116815260200190565b60009081526020902090565b6001600160a01b03811681146106f957600080fdfe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e6365ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef45524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220adc04b0c6fefa4d8f37f82f61d1984addf4f44c804e1896f7165bceb10fb292d64736f6c63430007060033",
}

// StableCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use StableCoinMetaData.ABI instead.
var StableCoinABI = StableCoinMetaData.ABI

// StableCoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StableCoinMetaData.Bin instead.
var StableCoinBin = StableCoinMetaData.Bin

// DeployStableCoin deploys a new Ethereum contract, binding an instance of StableCoin to it.
func DeployStableCoin(auth *bind.TransactOpts, backend bind.ContractBackend, _registry common.Address, _name string, _symbol string, _eligibleContracts []string, _forwarder common.Address) (common.Address, *types.Transaction, *StableCoin, error) {
	parsed, err := StableCoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StableCoinBin), backend, _registry, _name, _symbol, _eligibleContracts, _forwarder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StableCoin{StableCoinCaller: StableCoinCaller{contract: contract}, StableCoinTransactor: StableCoinTransactor{contract: contract}, StableCoinFilterer: StableCoinFilterer{contract: contract}}, nil
}

// StableCoin is an auto generated Go binding around an Ethereum contract.
type StableCoin struct {
	StableCoinCaller     // Read-only binding to the contract
	StableCoinTransactor // Write-only binding to the contract
	StableCoinFilterer   // Log filterer for contract events
}

// StableCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type StableCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StableCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StableCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StableCoinSession struct {
	Contract     *StableCoin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StableCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StableCoinCallerSession struct {
	Contract *StableCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StableCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StableCoinTransactorSession struct {
	Contract     *StableCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StableCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type StableCoinRaw struct {
	Contract *StableCoin // Generic contract binding to access the raw methods on
}

// StableCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StableCoinCallerRaw struct {
	Contract *StableCoinCaller // Generic read-only contract binding to access the raw methods on
}

// StableCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StableCoinTransactorRaw struct {
	Contract *StableCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStableCoin creates a new instance of StableCoin, bound to a specific deployed contract.
func NewStableCoin(address common.Address, backend bind.ContractBackend) (*StableCoin, error) {
	contract, err := bindStableCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StableCoin{StableCoinCaller: StableCoinCaller{contract: contract}, StableCoinTransactor: StableCoinTransactor{contract: contract}, StableCoinFilterer: StableCoinFilterer{contract: contract}}, nil
}

// NewStableCoinCaller creates a new read-only instance of StableCoin, bound to a specific deployed contract.
func NewStableCoinCaller(address common.Address, caller bind.ContractCaller) (*StableCoinCaller, error) {
	contract, err := bindStableCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StableCoinCaller{contract: contract}, nil
}

// NewStableCoinTransactor creates a new write-only instance of StableCoin, bound to a specific deployed contract.
func NewStableCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*StableCoinTransactor, error) {
	contract, err := bindStableCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StableCoinTransactor{contract: contract}, nil
}

// NewStableCoinFilterer creates a new log filterer instance of StableCoin, bound to a specific deployed contract.
func NewStableCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*StableCoinFilterer, error) {
	contract, err := bindStableCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StableCoinFilterer{contract: contract}, nil
}

// bindStableCoin binds a generic wrapper to an already deployed contract.
func bindStableCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StableCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StableCoin *StableCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StableCoin.Contract.StableCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StableCoin *StableCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableCoin.Contract.StableCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StableCoin *StableCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StableCoin.Contract.StableCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StableCoin *StableCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StableCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StableCoin *StableCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StableCoin *StableCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StableCoin.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StableCoin *StableCoinCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StableCoin.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StableCoin *StableCoinSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StableCoin.Contract.Allowance(&_StableCoin.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StableCoin *StableCoinCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StableCoin.Contract.Allowance(&_StableCoin.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StableCoin *StableCoinCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StableCoin.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StableCoin *StableCoinSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StableCoin.Contract.BalanceOf(&_StableCoin.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StableCoin *StableCoinCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StableCoin.Contract.BalanceOf(&_StableCoin.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StableCoin *StableCoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StableCoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StableCoin *StableCoinSession) Decimals() (uint8, error) {
	return _StableCoin.Contract.Decimals(&_StableCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StableCoin *StableCoinCallerSession) Decimals() (uint8, error) {
	return _StableCoin.Contract.Decimals(&_StableCoin.CallOpts)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_StableCoin *StableCoinCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _StableCoin.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_StableCoin *StableCoinSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _StableCoin.Contract.IsTrustedForwarder(&_StableCoin.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_StableCoin *StableCoinCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _StableCoin.Contract.IsTrustedForwarder(&_StableCoin.CallOpts, forwarder)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StableCoin *StableCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StableCoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StableCoin *StableCoinSession) Name() (string, error) {
	return _StableCoin.Contract.Name(&_StableCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StableCoin *StableCoinCallerSession) Name() (string, error) {
	return _StableCoin.Contract.Name(&_StableCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StableCoin *StableCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StableCoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StableCoin *StableCoinSession) Symbol() (string, error) {
	return _StableCoin.Contract.Symbol(&_StableCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StableCoin *StableCoinCallerSession) Symbol() (string, error) {
	return _StableCoin.Contract.Symbol(&_StableCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StableCoin *StableCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StableCoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StableCoin *StableCoinSession) TotalSupply() (*big.Int, error) {
	return _StableCoin.Contract.TotalSupply(&_StableCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StableCoin *StableCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _StableCoin.Contract.TotalSupply(&_StableCoin.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_StableCoin *StableCoinCaller) TrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StableCoin.contract.Call(opts, &out, "trustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_StableCoin *StableCoinSession) TrustedForwarder() (common.Address, error) {
	return _StableCoin.Contract.TrustedForwarder(&_StableCoin.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_StableCoin *StableCoinCallerSession) TrustedForwarder() (common.Address, error) {
	return _StableCoin.Contract.TrustedForwarder(&_StableCoin.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() pure returns(string)
func (_StableCoin *StableCoinCaller) VersionRecipient(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StableCoin.contract.Call(opts, &out, "versionRecipient")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() pure returns(string)
func (_StableCoin *StableCoinSession) VersionRecipient() (string, error) {
	return _StableCoin.Contract.VersionRecipient(&_StableCoin.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() pure returns(string)
func (_StableCoin *StableCoinCallerSession) VersionRecipient() (string, error) {
	return _StableCoin.Contract.VersionRecipient(&_StableCoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StableCoin *StableCoinTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StableCoin *StableCoinSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.Approve(&_StableCoin.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StableCoin *StableCoinTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.Approve(&_StableCoin.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_StableCoin *StableCoinTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_StableCoin *StableCoinSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.Burn(&_StableCoin.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_StableCoin *StableCoinTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.Burn(&_StableCoin.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_StableCoin *StableCoinTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_StableCoin *StableCoinSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.BurnFrom(&_StableCoin.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_StableCoin *StableCoinTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.BurnFrom(&_StableCoin.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StableCoin *StableCoinTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StableCoin.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StableCoin *StableCoinSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.DecreaseAllowance(&_StableCoin.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StableCoin *StableCoinTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.DecreaseAllowance(&_StableCoin.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StableCoin *StableCoinTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StableCoin.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StableCoin *StableCoinSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.IncreaseAllowance(&_StableCoin.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StableCoin *StableCoinTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.IncreaseAllowance(&_StableCoin.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _recepient, uint256 _amount) returns(bool)
func (_StableCoin *StableCoinTransactor) Mint(opts *bind.TransactOpts, _recepient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.contract.Transact(opts, "mint", _recepient, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _recepient, uint256 _amount) returns(bool)
func (_StableCoin *StableCoinSession) Mint(_recepient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.Mint(&_StableCoin.TransactOpts, _recepient, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _recepient, uint256 _amount) returns(bool)
func (_StableCoin *StableCoinTransactorSession) Mint(_recepient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.Mint(&_StableCoin.TransactOpts, _recepient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StableCoin *StableCoinTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StableCoin *StableCoinSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.Transfer(&_StableCoin.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StableCoin *StableCoinTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.Transfer(&_StableCoin.TransactOpts, recipient, amount)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool)
func (_StableCoin *StableCoinTransactor) TransferAndCall(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _StableCoin.contract.Transact(opts, "transferAndCall", _to, _value, _data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool)
func (_StableCoin *StableCoinSession) TransferAndCall(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _StableCoin.Contract.TransferAndCall(&_StableCoin.TransactOpts, _to, _value, _data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool)
func (_StableCoin *StableCoinTransactorSession) TransferAndCall(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _StableCoin.Contract.TransferAndCall(&_StableCoin.TransactOpts, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StableCoin *StableCoinTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StableCoin *StableCoinSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.TransferFrom(&_StableCoin.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StableCoin *StableCoinTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCoin.Contract.TransferFrom(&_StableCoin.TransactOpts, sender, recipient, amount)
}

// StableCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StableCoin contract.
type StableCoinApprovalIterator struct {
	Event *StableCoinApproval // Event containing the contract specifics and raw log

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
func (it *StableCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableCoinApproval)
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
		it.Event = new(StableCoinApproval)
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
func (it *StableCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableCoinApproval represents a Approval event raised by the StableCoin contract.
type StableCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StableCoin *StableCoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StableCoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StableCoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StableCoinApprovalIterator{contract: _StableCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StableCoin *StableCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StableCoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StableCoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableCoinApproval)
				if err := _StableCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StableCoin *StableCoinFilterer) ParseApproval(log types.Log) (*StableCoinApproval, error) {
	event := new(StableCoinApproval)
	if err := _StableCoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StableCoin contract.
type StableCoinTransferIterator struct {
	Event *StableCoinTransfer // Event containing the contract specifics and raw log

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
func (it *StableCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableCoinTransfer)
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
		it.Event = new(StableCoinTransfer)
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
func (it *StableCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableCoinTransfer represents a Transfer event raised by the StableCoin contract.
type StableCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_StableCoin *StableCoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StableCoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StableCoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StableCoinTransferIterator{contract: _StableCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_StableCoin *StableCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StableCoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StableCoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableCoinTransfer)
				if err := _StableCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_StableCoin *StableCoinFilterer) ParseTransfer(log types.Log) (*StableCoinTransfer, error) {
	event := new(StableCoinTransfer)
	if err := _StableCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableCoinTransfer0Iterator is returned from FilterTransfer0 and is used to iterate over the raw logs and unpacked data for Transfer0 events raised by the StableCoin contract.
type StableCoinTransfer0Iterator struct {
	Event *StableCoinTransfer0 // Event containing the contract specifics and raw log

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
func (it *StableCoinTransfer0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableCoinTransfer0)
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
		it.Event = new(StableCoinTransfer0)
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
func (it *StableCoinTransfer0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableCoinTransfer0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableCoinTransfer0 represents a Transfer0 event raised by the StableCoin contract.
type StableCoinTransfer0 struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer0 is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StableCoin *StableCoinFilterer) FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StableCoinTransfer0Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StableCoin.contract.FilterLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StableCoinTransfer0Iterator{contract: _StableCoin.contract, event: "Transfer0", logs: logs, sub: sub}, nil
}

// WatchTransfer0 is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StableCoin *StableCoinFilterer) WatchTransfer0(opts *bind.WatchOpts, sink chan<- *StableCoinTransfer0, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StableCoin.contract.WatchLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableCoinTransfer0)
				if err := _StableCoin.contract.UnpackLog(event, "Transfer0", log); err != nil {
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

// ParseTransfer0 is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StableCoin *StableCoinFilterer) ParseTransfer0(log types.Log) (*StableCoinTransfer0, error) {
	event := new(StableCoinTransfer0)
	if err := _StableCoin.contract.UnpackLog(event, "Transfer0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
