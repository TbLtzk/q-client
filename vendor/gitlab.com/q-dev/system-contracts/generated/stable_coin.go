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
	Bin: "0x60806040523480156200001157600080fd5b50604051620017193803806200171983398101604081905262000034916200033e565b83838181816003908051906020019062000050929190620000da565b50805162000066906004906020840190620000da565b50505050506200007c81620000b860201b60201c565b600680546001600160a01b0319166001600160a01b0387161790558151620000ac90600790602085019062000169565b505050505050620004c4565b600580546001600160a01b0319166001600160a01b0392909216919091179055565b828054620000e89062000487565b90600052602060002090601f0160209004810192826200010c576000855562000157565b82601f106200012757805160ff191683800117855562000157565b8280016001018555821562000157579182015b82811115620001575782518255916020019190600101906200013a565b5062000165929150620001c9565b5090565b828054828255906000526020600020908101928215620001bb579160200282015b82811115620001bb5782518051620001aa918491602090910190620000da565b50916020019190600101906200018a565b5062000165929150620001e0565b5b80821115620001655760008155600101620001ca565b8082111562000165576000620001f7828262000201565b50600101620001e0565b5080546200020f9062000487565b6000825580601f1062000220575050565b601f016020900490600052602060002090810190620002409190620001c9565b50565b80516001600160a01b03811681146200025b57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620002a157620002a162000260565b604052919050565b600082601f830112620002bb57600080fd5b81516001600160401b03811115620002d757620002d762000260565b6020620002ed601f8301601f1916820162000276565b82815285828487010111156200030257600080fd5b60005b838110156200032257858101830151828201840152820162000305565b83811115620003345760008385840101525b5095945050505050565b600080600080600060a086880312156200035757600080fd5b620003628662000243565b602087810151919650906001600160401b03808211156200038257600080fd5b620003908a838b01620002a9565b96506040890151915080821115620003a757600080fd5b620003b58a838b01620002a9565b95506060890151915080821115620003cc57600080fd5b818901915089601f830112620003e157600080fd5b815181811115620003f657620003f662000260565b8060051b6200040785820162000276565b918252838101850191858101908d8411156200042257600080fd5b86860192505b838310156200046357825185811115620004425760008081fd5b620004528f89838a0101620002a9565b835250918601919086019062000428565b809850505050505050506200047b6080870162000243565b90509295509295909350565b600181811c908216806200049c57607f821691505b60208210811415620004be57634e487b7160e01b600052602260045260246000fd5b50919050565b61124580620004d46000396000f3fe608060405234801561001057600080fd5b50600436106100f65760003560e01c8063486ff0cd11610092578063486ff0cd146101be578063572b6c05146101df57806370a08231146101f257806379cc67901461021b5780637da0a8771461022e57806395d89b4114610249578063a457c2d714610251578063a9059cbb14610264578063dd62ed3e1461027757600080fd5b806306fdde03146100fb578063095ea7b31461011957806318160ddd1461013c57806323b872dd1461014e578063313ce5671461016157806339509351146101705780634000aea01461018357806340c10f191461019657806342966c68146101a9575b600080fd5b6101036102b0565b6040516101109190610e1f565b60405180910390f35b61012c610127366004610e47565b610342565b6040519015158152602001610110565b6002545b604051908152602001610110565b61012c61015c366004610e73565b61035f565b60405160128152602001610110565b61012c61017e366004610e47565b610437565b61012c610191366004610eca565b61048b565b61012c6101a4366004610e47565b610594565b6101bc6101b7366004610f97565b610724565b005b6040805180820190915260058152640322e322e360dc1b6020820152610103565b61012c6101ed366004610fb0565b610738565b610140610200366004610fb0565b6001600160a01b031660009081526020819052604090205490565b6101bc610229366004610e47565b61074c565b6005546040516001600160a01b039091168152602001610110565b6101036107db565b61012c61025f366004610e47565b6107ea565b61012c610272366004610e47565b610897565b610140610285366004610fcd565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6060600380546102bf90611006565b80601f01602080910402602001604051908101604052809291908181526020018280546102eb90611006565b80156103385780601f1061030d57610100808354040283529160200191610338565b820191906000526020600020905b81548152906001019060200180831161031b57829003601f168201915b5050505050905090565b600061035661034f6108ab565b84846108ba565b50600192915050565b600061036c8484846109de565b6001600160a01b03841660009081526001602052604081208161038d6108ab565b6001600160a01b03166001600160a01b03168152602001908152602001600020549050828110156104165760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b61042a856104226108ab565b8584036108ba565b60019150505b9392505050565b60006103566104446108ab565b8484600160006104526108ab565b6001600160a01b03908116825260208083019390935260409182016000908120918b16815292529020546104869190611057565b6108ba565b6000806104988585610897565b9050806104a9576000915050610430565b846001600160a01b03166104bb6108ab565b6001600160a01b03167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1686866040516104f592919061106f565b60405180910390a3846001600160a01b03811663c0ee0b8a6105156108ab565b87876040518463ffffffff1660e01b815260040161053593929190611090565b602060405180830381600087803b15801561054f57600080fd5b505af1158015610563573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061058791906110c0565b5060019695505050505050565b60008060005b60075481101561068a576105ac6108ab565b6001600160a01b0316600660009054906101000a90046001600160a01b03166001600160a01b0316633fb90271600784815481106105ec576105ec6110e2565b906000526020600020016040518263ffffffff1660e01b815260040161061291906110f8565b60206040518083038186803b15801561062a57600080fd5b505afa15801561063e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066291906111a0565b6001600160a01b0316141561067a576001915061068a565b610683816111bd565b905061059a565b50806107105760405162461bcd60e51b815260206004820152604960248201527f5b5145432d3032303030305d2d5065726d697373696f6e2064656e696564202d60448201527f206f6e6c792074686520656c696769626c6520636f6e74726163747320686176606482015268329030b1b1b2b9b99760b91b608482015260a40161040d565b61071a8484610b9b565b5060019392505050565b61073561072f6108ab565b82610c68565b50565b6005546001600160a01b0391821691161490565b600061075a836102856108ab565b9050818110156107b85760405162461bcd60e51b8152602060048201526024808201527f45524332303a206275726e20616d6f756e74206578636565647320616c6c6f77604482015263616e636560e01b606482015260840161040d565b6107cc836107c46108ab565b8484036108ba565b6107d68383610c68565b505050565b6060600480546102bf90611006565b600080600160006107f96108ab565b6001600160a01b03908116825260208083019390935260409182016000908120918816815292529020549050828110156108835760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b606482015260840161040d565b61071a61088e6108ab565b858584036108ba565b60006103566108a46108ab565b84846109de565b60006108b5610da4565b905090565b6001600160a01b03831661091c5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161040d565b6001600160a01b03821661097d5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161040d565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b038316610a425760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161040d565b6001600160a01b038216610aa45760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161040d565b6001600160a01b03831660009081526020819052604090205481811015610b1c5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161040d565b6001600160a01b03808516600090815260208190526040808220858503905591851681529081208054849290610b53908490611057565b92505081905550826001600160a01b0316846001600160a01b03166000805160206111f083398151915284604051610b8d91815260200190565b60405180910390a350505050565b6001600160a01b038216610bf15760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161040d565b8060026000828254610c039190611057565b90915550506001600160a01b03821660009081526020819052604081208054839290610c30908490611057565b90915550506040518181526001600160a01b038316906000906000805160206111f08339815191529060200160405180910390a35050565b6001600160a01b038216610cc85760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b606482015260840161040d565b6001600160a01b03821660009081526020819052604090205481811015610d3c5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b606482015260840161040d565b6001600160a01b0383166000908152602081905260408120838303905560028054849290610d6b9084906111d8565b90915550506040518281526000906001600160a01b038516906000805160206111f08339815191529060200160405180910390a3505050565b600060143610801590610dbb5750610dbb33610738565b15610dcd575060131936013560601c90565b503390565b6000815180845260005b81811015610df857602081850181015186830182015201610ddc565b81811115610e0a576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006104306020830184610dd2565b6001600160a01b038116811461073557600080fd5b60008060408385031215610e5a57600080fd5b8235610e6581610e32565b946020939093013593505050565b600080600060608486031215610e8857600080fd5b8335610e9381610e32565b92506020840135610ea381610e32565b929592945050506040919091013590565b634e487b7160e01b600052604160045260246000fd5b600080600060608486031215610edf57600080fd5b8335610eea81610e32565b925060208401359150604084013567ffffffffffffffff80821115610f0e57600080fd5b818601915086601f830112610f2257600080fd5b813581811115610f3457610f34610eb4565b604051601f8201601f19908116603f01168101908382118183101715610f5c57610f5c610eb4565b81604052828152896020848701011115610f7557600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b600060208284031215610fa957600080fd5b5035919050565b600060208284031215610fc257600080fd5b813561043081610e32565b60008060408385031215610fe057600080fd5b8235610feb81610e32565b91506020830135610ffb81610e32565b809150509250929050565b600181811c9082168061101a57607f821691505b6020821081141561103b57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b6000821982111561106a5761106a611041565b500190565b8281526040602082015260006110886040830184610dd2565b949350505050565b60018060a01b03841681528260208201526060604082015260006110b76060830184610dd2565b95945050505050565b6000602082840312156110d257600080fd5b8151801515811461043057600080fd5b634e487b7160e01b600052603260045260246000fd5b600060208083526000845481600182811c91508083168061111a57607f831692505b85831081141561113857634e487b7160e01b85526022600452602485fd5b878601838152602001818015611155576001811461116657611191565b60ff19861682528782019650611191565b60008b81526020902060005b8681101561118b57815484820152908501908901611172565b83019750505b50949998505050505050505050565b6000602082840312156111b257600080fd5b815161043081610e32565b60006000198214156111d1576111d1611041565b5060010190565b6000828210156111ea576111ea611041565b50039056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa26469706673582212208a8c1b96e91fa835c7ae49319eb908f21c9972b8179b0ba4126b6954e6d5425c64736f6c63430008090033",
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
