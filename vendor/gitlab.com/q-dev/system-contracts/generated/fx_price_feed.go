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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pair_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimal_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"maintainersList_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"baseTokenAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteTokenAddr_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimalPlaces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futurePricingTolerance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaintainers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leaveMaintainers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pair\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exchangeRate_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricingTime_\",\"type\":\"uint256\"}],\"name\":\"setExchangeRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tolerance_\",\"type\":\"uint256\"}],\"name\":\"setFuturePricingTolerance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maintainer_\",\"type\":\"address\"}],\"name\":\"setMaintainer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620023873803806200238783398101604081905262000034916200064f565b6200003f3362000454565b8451620000b95760405162461bcd60e51b815260206004820152603560248201527f5b5145432d3031393030315d2d496e76616c696420706169722c20746865206960448201527f6e697469616c697a6174696f6e206661696c65642e000000000000000000000060648201526084015b60405180910390fd5b60008411620001315760405162461bcd60e51b815260206004820152603860248201527f5b5145432d3031393030325d2d496e76616c696420646563696d616c2c20746860448201527f6520696e697469616c697a6174696f6e206661696c65642e00000000000000006064820152608401620000b0565b6001600160a01b038216620001bb5760405162461bcd60e51b815260206004820152604360248201527f5b5145432d3031393030335d2d496e76616c6964206261736520746f6b656e2060448201527f616464726573732c2074686520696e697469616c697a6174696f6e206661696c60648201526232b21760e91b608482015260a401620000b0565b6001600160a01b038116620002475760405162461bcd60e51b8152602060048201526044602482018190527f5b5145432d3031393030345d2d496e76616c69642071756f746520746f6b656e908201527f20616464726573732c2074686520696e697469616c697a6174696f6e206661696064820152633632b21760e11b608482015260a401620000b0565b63ffffffff823b16620002dd5760405162461bcd60e51b815260206004820152605160248201527f5b5145432d3031393030355d2d546865206261736520746f6b656e206164647260448201527f657373206973206e6f74206120636f6e74726163742c2074686520696e69746960648201527030b634bd30ba34b7b7103330b4b632b21760791b608482015260a401620000b0565b63ffffffff813b16620003745760405162461bcd60e51b815260206004820152605260248201527f5b5145432d3031393030365d2d5468652071756f746520746f6b656e2061646460448201527f72657373206973206e6f74206120636f6e74726163742c2074686520696e697460648201527134b0b634bd30ba34b7b7103330b4b632b21760711b608482015260a401620000b0565b6040516200038290620004a4565b604051809103906000f0801580156200039f573d6000803e3d6000fd5b50600180546001600160a01b0319166001600160a01b0392909216918217905560405163a224cee760e01b815263a224cee790620003e290869060040162000763565b600060405180830381600087803b158015620003fd57600080fd5b505af115801562000412573d6000803e3d6000fd5b505086516200042b9250600291506020880190620004b2565b5060c0939093526001600160a01b0390811660805290911660a0525050610708600655620007ef565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610db780620015d083390190565b828054620004c090620007b2565b90600052602060002090601f016020900481019282620004e457600085556200052f565b82601f10620004ff57805160ff19168380011785556200052f565b828001600101855582156200052f579182015b828111156200052f57825182559160200191906001019062000512565b506200053d92915062000541565b5090565b5b808211156200053d576000815560010162000542565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171562000599576200059962000558565b604052919050565b80516001600160a01b0381168114620005b957600080fd5b919050565b600082601f830112620005d057600080fd5b815160206001600160401b03821115620005ee57620005ee62000558565b8160051b620005ff8282016200056e565b92835284810182019282810190878511156200061a57600080fd5b83870192505b8483101562000644576200063483620005a1565b8252918301919083019062000620565b979650505050505050565b600080600080600060a086880312156200066857600080fd5b85516001600160401b03808211156200068057600080fd5b818801915088601f8301126200069557600080fd5b815181811115620006aa57620006aa62000558565b6020620006c0601f8301601f191682016200056e565b8281528b82848701011115620006d557600080fd5b60005b83811015620006f5578581018301518282018401528201620006d8565b83811115620007075760008385840101525b50908a015160408b01519199509750925050808211156200072757600080fd5b506200073688828901620005be565b9350506200074760608701620005a1565b91506200075760808701620005a1565b90509295509295909350565b6020808252825182820181905260009190848201906040850190845b81811015620007a65783516001600160a01b0316835292840192918401916001016200077f565b50909695505050505050565b600181811c90821680620007c757607f821691505b60208210811415620007e957634e487b7160e01b600052602260045260246000fd5b50919050565b60805160a05160c051610db16200081f60003960006101e601526000610136015260006101840152610db16000f3fe608060405234801561001057600080fd5b50600436106100e05760003560e01c8063a094600e11610087578063a094600e1461017f578063a8aa1b31146101a6578063d17f4225146101bb578063d5204033146101d0578063dc555090146101d8578063e1725c92146101e1578063f2fde38b14610208578063f55ecf061461021b57600080fd5b806313ea5d29146100e557806337c5f2bb146100fa5780633ba0b9a91461010d578063715018a6146101295780637acede8b146101315780637f1ba22814610165578063871ebd471461016e5780638da5cb5b14610177575b600080fd5b6100f86100f3366004610a4d565b61023e565b005b6100f8610108366004610a71565b610348565b61011660055481565b6040519081526020015b60405180910390f35b6100f86103e9565b6101587f000000000000000000000000000000000000000000000000000000000000000081565b6040516101209190610a8a565b61011660065481565b61011660045481565b610158610424565b6101587f000000000000000000000000000000000000000000000000000000000000000081565b6101ae610433565b6040516101209190610a9e565b6101c36104c1565b6040516101209190610af3565b6100f8610547565b61011660035481565b6101167f000000000000000000000000000000000000000000000000000000000000000081565b6100f8610216366004610a4d565b61073f565b61022e610229366004610b40565b6107df565b6040519015158152602001610120565b600154604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e89061026e903390600401610a8a565b60206040518083038186803b15801561028657600080fd5b505afa15801561029a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102be9190610b62565b6102e35760405162461bcd60e51b81526004016102da90610b84565b60405180910390fd5b600154604051639e9405c360e01b81526001600160a01b0390911690639e9405c390610313908490600401610a8a565b600060405180830381600087803b15801561032d57600080fd5b505af1158015610341573d6000803e3d6000fd5b5050505050565b600154604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e890610378903390600401610a8a565b60206040518083038186803b15801561039057600080fd5b505afa1580156103a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103c89190610b62565b6103e45760405162461bcd60e51b81526004016102da90610b84565b600655565b336103f2610424565b6001600160a01b0316146104185760405162461bcd60e51b81526004016102da90610be1565b61042260006109e8565b565b6000546001600160a01b031690565b6002805461044090610c16565b80601f016020809104026020016040519081016040528092919081815260200182805461046c90610c16565b80156104b95780601f1061048e576101008083540402835291602001916104b9565b820191906000526020600020905b81548152906001019060200180831161049c57829003601f168201915b505050505081565b600154604080516351cfd60960e11b815290516060926001600160a01b03169163a39fac12916004808301926000929190829003018186803b15801561050657600080fd5b505afa15801561051a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526105429190810190610c77565b905090565b600154604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e890610577903390600401610a8a565b60206040518083038186803b15801561058f57600080fd5b505afa1580156105a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c79190610b62565b156106db5760018060009054906101000a90046001600160a01b03166001600160a01b031663949d225d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561061b57600080fd5b505afa15801561062f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106539190610d3c565b116106db5760405162461bcd60e51b815260206004820152604c60248201527f5b5145432d3031393030385d2d43616e6e6f74206c6561766520746865206d6160448201527f696e7461696e657273206c6973742c20796f752061726520746865206f6e6c7960648201526b1036b0b4b73a30b4b732b91760a11b608482015260a4016102da565b600154604051636196c02d60e11b81526001600160a01b039091169063c32d805a9061070b903390600401610a8a565b600060405180830381600087803b15801561072557600080fd5b505af1158015610739573d6000803e3d6000fd5b50505050565b33610748610424565b6001600160a01b03161461076e5760405162461bcd60e51b81526004016102da90610be1565b6001600160a01b0381166107d35760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102da565b6107dc816109e8565b50565b600154604051630bb7c8fd60e31b81526000916001600160a01b031690635dbe47e890610810903390600401610a8a565b60206040518083038186803b15801561082857600080fd5b505afa15801561083c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108609190610b62565b61087c5760405162461bcd60e51b81526004016102da90610b84565b600083116109005760405162461bcd60e51b8152602060048201526044602482018190527f5b5145432d3031393030375d2d496e76616c69642065786368616e6765207261908201527f74652c206661696c656420746f20736574207468652065786368616e6765207260648201526330ba329760e11b608482015260a4016102da565b600454821161096d5760405162461bcd60e51b815260206004820152603360248201527f5b5145432d3031393030395d2d50726963696e672074696d65206973206c6573604482015272399037b91032b8bab0b61031bab93932b73a1760691b60648201526084016102da565b60065461097a9042610d55565b8211156109d55760405162461bcd60e51b8152602060048201526024808201527f5b5145432d3031393031305d2d50726963696e672074696d6520696e206675746044820152633ab9329760e11b60648201526084016102da565b5060059190915542600355600455600190565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811681146107dc57600080fd5b600060208284031215610a5f57600080fd5b8135610a6a81610a38565b9392505050565b600060208284031215610a8357600080fd5b5035919050565b6001600160a01b0391909116815260200190565b600060208083528351808285015260005b81811015610acb57858101830151858201604001528201610aaf565b81811115610add576000604083870101525b50601f01601f1916929092016040019392505050565b6020808252825182820181905260009190848201906040850190845b81811015610b345783516001600160a01b031683529284019291840191600101610b0f565b50909695505050505050565b60008060408385031215610b5357600080fd5b50508035926020909101359150565b600060208284031215610b7457600080fd5b81518015158114610a6a57600080fd5b6020808252603e908201527f5b5145432d3031393030305d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c79206d61696e7461696e6572732068617665206163636573732e0000606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b600181811c90821680610c2a57607f821691505b60208210811415610c4b57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b8051610c7281610a38565b919050565b60006020808385031215610c8a57600080fd5b825167ffffffffffffffff80821115610ca257600080fd5b818501915085601f830112610cb657600080fd5b815181811115610cc857610cc8610c51565b8060051b604051601f19603f83011681018181108582111715610ced57610ced610c51565b604052918252848201925083810185019188831115610d0b57600080fd5b938501935b82851015610d3057610d2185610c67565b84529385019392850192610d10565b98975050505050505050565b600060208284031215610d4e57600080fd5b5051919050565b60008219821115610d7657634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220b2cbfb7be8a76e349c2017776788a8fcba3493f2af7d1f386a0d3d134438aca164736f6c63430008090033608060405234801561001057600080fd5b50610d97806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a45760003560e01c80630a3b0a4f146100a957806329092d0e146100d157806352efea6e146100e45780635dbe47e8146100ec578063715018a6146100ff5780638da5cb5b14610109578063949d225d146101295780639e9405c31461013a578063a224cee71461014d578063a39fac1214610160578063c32d805a14610175578063f2fde38b14610188575b600080fd5b6100bc6100b7366004610b0a565b61019b565b60405190151581526020015b60405180910390f35b6100bc6100df366004610b0a565b6101fd565b6100bc61024c565b6100bc6100fa366004610b0a565b6102eb565b610107610308565b005b610111610343565b6040516001600160a01b0390911681526020016100c8565b6066546040519081526020016100c8565b610107610148366004610b0a565b610352565b61010761015b366004610b42565b610410565b610168610584565b6040516100c89190610c07565b610107610183366004610b0a565b6105e6565b610107610196366004610b0a565b61066d565b6000336101a6610343565b6001600160a01b0316146101d55760405162461bcd60e51b81526004016101cc90610c54565b60405180910390fd5b6101de826102eb565b156101eb57506000919050565b6101f48261070a565b5060015b919050565b600033610208610343565b6001600160a01b03161461022e5760405162461bcd60e51b81526004016101cc90610c54565b610237826102eb565b61024357506000919050565b6101f48261076d565b600033610257610343565b6001600160a01b03161461027d5760405162461bcd60e51b81526004016101cc90610c54565b60005b6066548110156102d85760656000606683815481106102a1576102a1610c89565b60009182526020808320909101546001600160a01b03168352820192909252604001812055806102d081610cb5565b915050610280565b506102e560666000610ac1565b50600190565b6001600160a01b0316600090815260656020526040902054151590565b33610311610343565b6001600160a01b0316146103375760405162461bcd60e51b81526004016101cc90610c54565b61034160006108d3565b565b6033546001600160a01b031690565b61035b8161019b565b61040d5760405162461bcd60e51b815260206004820152607160248201527f5b5145432d3033353030325d2d54686520616464726573732068617320616c7260448201527f65616479206265656e20616464656420746f207468652073746f726167652c2060648201527f6661696c656420746f2061646420746865206164647265737320746f207468656084820152701030b2323932b9b99039ba37b930b3b29760791b60a482015260c4016101cc565b50565b600054610100900460ff1680610429575060005460ff16155b6104455760405162461bcd60e51b81526004016101cc90610cd0565b600054610100900460ff16158015610467576000805461ffff19166101011790555b60005b82518110156105655760006065600085848151811061048b5761048b610c89565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206000015411156104c257610553565b60668382815181106104d6576104d6610c89565b6020908102919091018101518254600181018455600093845291832090910180546001600160a01b0319166001600160a01b039092169190911790556066548451909160659186908590811061052e5761052e610c89565b6020908102919091018101516001600160a01b03168252810191909152604001600020555b8061055d81610cb5565b91505061046a565b5061056e610925565b8015610580576000805461ff00191690555b5050565b606060668054806020026020016040519081016040528092919081815260200182805480156105dc57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116105be575b5050505050905090565b6105ef816101fd565b61040d5760405162461bcd60e51b815260206004820152604360248201527f5b5145432d3033353030305d2d4661696c656420746f2072656d6f766520746860448201527f6520616464726573732066726f6d2074686520616464726573732073746f726160648201526233b29760e91b608482015260a4016101cc565b33610676610343565b6001600160a01b03161461069c5760405162461bcd60e51b81526004016101cc90610c54565b6001600160a01b0381166107015760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101cc565b61040d816108d3565b606680546001810182557f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943540180546001600160a01b0319166001600160a01b03841690811790915590546000918252606560205260409091205561040d816109a0565b6001600160a01b038116600090815260656020908152604080832081519283019091525481526066549091906107a590600190610d1e565b90506000606682815481106107bc576107bc610c89565b60009182526020909120015483516001600160a01b0390911691506107e390600190610d1e565b82146108565782516001600160a01b0382166000908152606560205260409020558251819060669061081790600190610d1e565b8154811061082757610827610c89565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b606680548061086757610867610d35565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526065905260408120556108ab846109a0565b836001600160a01b0316816001600160a01b0316146108cd576108cd816109a0565b50505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff168061093e575060005460ff16155b61095a5760405162461bcd60e51b81526004016101cc90610cd0565b600054610100900460ff1615801561097c576000805461ffff19166101011790555b6109846109f7565b61098c610a61565b801561040d576000805461ff001916905550565b6001600160a01b0381166000908152606560205260409020546066548111156109cb576109cb610d4b565b6109d4826102eb565b156109e9576000811161058057610580610d4b565b801561058057610580610d4b565b600054610100900460ff1680610a10575060005460ff16155b610a2c5760405162461bcd60e51b81526004016101cc90610cd0565b600054610100900460ff1615801561098c576000805461ffff1916610101179055801561040d576000805461ff001916905550565b600054610100900460ff1680610a7a575060005460ff16155b610a965760405162461bcd60e51b81526004016101cc90610cd0565b600054610100900460ff16158015610ab8576000805461ffff19166101011790555b61098c336108d3565b508054600082559060005260206000209081019061040d91905b80821115610aef5760008155600101610adb565b5090565b80356001600160a01b03811681146101f857600080fd5b600060208284031215610b1c57600080fd5b610b2582610af3565b9392505050565b634e487b7160e01b600052604160045260246000fd5b60006020808385031215610b5557600080fd5b823567ffffffffffffffff80821115610b6d57600080fd5b818501915085601f830112610b8157600080fd5b813581811115610b9357610b93610b2c565b8060051b604051601f19603f83011681018181108582111715610bb857610bb8610b2c565b604052918252848201925083810185019188831115610bd657600080fd5b938501935b82851015610bfb57610bec85610af3565b84529385019392850192610bdb565b98975050505050505050565b6020808252825182820181905260009190848201906040850190845b81811015610c485783516001600160a01b031683529284019291840191600101610c23565b50909695505050505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415610cc957610cc9610c9f565b5060010190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b600082821015610d3057610d30610c9f565b500390565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052600160045260246000fdfea2646970667358221220b5288929432e31abe76024cbbc8c3ca26bca79cb8f80ddf9b8f50d8c257f0d6e64736f6c63430008090033",
}

// FxPriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use FxPriceFeedMetaData.ABI instead.
var FxPriceFeedABI = FxPriceFeedMetaData.ABI

// FxPriceFeedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FxPriceFeedMetaData.Bin instead.
var FxPriceFeedBin = FxPriceFeedMetaData.Bin

// DeployFxPriceFeed deploys a new Ethereum contract, binding an instance of FxPriceFeed to it.
func DeployFxPriceFeed(auth *bind.TransactOpts, backend bind.ContractBackend, pair_ string, decimal_ *big.Int, maintainersList_ []common.Address, baseTokenAddr_ common.Address, quoteTokenAddr_ common.Address) (common.Address, *types.Transaction, *FxPriceFeed, error) {
	parsed, err := FxPriceFeedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FxPriceFeedBin), backend, pair_, decimal_, maintainersList_, baseTokenAddr_, quoteTokenAddr_)
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

// FuturePricingTolerance is a free data retrieval call binding the contract method 0x7f1ba228.
//
// Solidity: function futurePricingTolerance() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedCaller) FuturePricingTolerance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FxPriceFeed.contract.Call(opts, &out, "futurePricingTolerance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FuturePricingTolerance is a free data retrieval call binding the contract method 0x7f1ba228.
//
// Solidity: function futurePricingTolerance() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedSession) FuturePricingTolerance() (*big.Int, error) {
	return _FxPriceFeed.Contract.FuturePricingTolerance(&_FxPriceFeed.CallOpts)
}

// FuturePricingTolerance is a free data retrieval call binding the contract method 0x7f1ba228.
//
// Solidity: function futurePricingTolerance() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedCallerSession) FuturePricingTolerance() (*big.Int, error) {
	return _FxPriceFeed.Contract.FuturePricingTolerance(&_FxPriceFeed.CallOpts)
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

// PricingTime is a free data retrieval call binding the contract method 0x871ebd47.
//
// Solidity: function pricingTime() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedCaller) PricingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FxPriceFeed.contract.Call(opts, &out, "pricingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricingTime is a free data retrieval call binding the contract method 0x871ebd47.
//
// Solidity: function pricingTime() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedSession) PricingTime() (*big.Int, error) {
	return _FxPriceFeed.Contract.PricingTime(&_FxPriceFeed.CallOpts)
}

// PricingTime is a free data retrieval call binding the contract method 0x871ebd47.
//
// Solidity: function pricingTime() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedCallerSession) PricingTime() (*big.Int, error) {
	return _FxPriceFeed.Contract.PricingTime(&_FxPriceFeed.CallOpts)
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

// UpdateTime is a free data retrieval call binding the contract method 0xdc555090.
//
// Solidity: function updateTime() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedCaller) UpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FxPriceFeed.contract.Call(opts, &out, "updateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdateTime is a free data retrieval call binding the contract method 0xdc555090.
//
// Solidity: function updateTime() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedSession) UpdateTime() (*big.Int, error) {
	return _FxPriceFeed.Contract.UpdateTime(&_FxPriceFeed.CallOpts)
}

// UpdateTime is a free data retrieval call binding the contract method 0xdc555090.
//
// Solidity: function updateTime() view returns(uint256)
func (_FxPriceFeed *FxPriceFeedCallerSession) UpdateTime() (*big.Int, error) {
	return _FxPriceFeed.Contract.UpdateTime(&_FxPriceFeed.CallOpts)
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

// SetExchangeRate is a paid mutator transaction binding the contract method 0xf55ecf06.
//
// Solidity: function setExchangeRate(uint256 exchangeRate_, uint256 pricingTime_) returns(bool)
func (_FxPriceFeed *FxPriceFeedTransactor) SetExchangeRate(opts *bind.TransactOpts, exchangeRate_ *big.Int, pricingTime_ *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.contract.Transact(opts, "setExchangeRate", exchangeRate_, pricingTime_)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xf55ecf06.
//
// Solidity: function setExchangeRate(uint256 exchangeRate_, uint256 pricingTime_) returns(bool)
func (_FxPriceFeed *FxPriceFeedSession) SetExchangeRate(exchangeRate_ *big.Int, pricingTime_ *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetExchangeRate(&_FxPriceFeed.TransactOpts, exchangeRate_, pricingTime_)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xf55ecf06.
//
// Solidity: function setExchangeRate(uint256 exchangeRate_, uint256 pricingTime_) returns(bool)
func (_FxPriceFeed *FxPriceFeedTransactorSession) SetExchangeRate(exchangeRate_ *big.Int, pricingTime_ *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetExchangeRate(&_FxPriceFeed.TransactOpts, exchangeRate_, pricingTime_)
}

// SetFuturePricingTolerance is a paid mutator transaction binding the contract method 0x37c5f2bb.
//
// Solidity: function setFuturePricingTolerance(uint256 tolerance_) returns()
func (_FxPriceFeed *FxPriceFeedTransactor) SetFuturePricingTolerance(opts *bind.TransactOpts, tolerance_ *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.contract.Transact(opts, "setFuturePricingTolerance", tolerance_)
}

// SetFuturePricingTolerance is a paid mutator transaction binding the contract method 0x37c5f2bb.
//
// Solidity: function setFuturePricingTolerance(uint256 tolerance_) returns()
func (_FxPriceFeed *FxPriceFeedSession) SetFuturePricingTolerance(tolerance_ *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetFuturePricingTolerance(&_FxPriceFeed.TransactOpts, tolerance_)
}

// SetFuturePricingTolerance is a paid mutator transaction binding the contract method 0x37c5f2bb.
//
// Solidity: function setFuturePricingTolerance(uint256 tolerance_) returns()
func (_FxPriceFeed *FxPriceFeedTransactorSession) SetFuturePricingTolerance(tolerance_ *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetFuturePricingTolerance(&_FxPriceFeed.TransactOpts, tolerance_)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address maintainer_) returns()
func (_FxPriceFeed *FxPriceFeedTransactor) SetMaintainer(opts *bind.TransactOpts, maintainer_ common.Address) (*types.Transaction, error) {
	return _FxPriceFeed.contract.Transact(opts, "setMaintainer", maintainer_)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address maintainer_) returns()
func (_FxPriceFeed *FxPriceFeedSession) SetMaintainer(maintainer_ common.Address) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetMaintainer(&_FxPriceFeed.TransactOpts, maintainer_)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address maintainer_) returns()
func (_FxPriceFeed *FxPriceFeedTransactorSession) SetMaintainer(maintainer_ common.Address) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetMaintainer(&_FxPriceFeed.TransactOpts, maintainer_)
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
