// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractRegistryPair is an auto generated low-level Go binding around an user-defined struct.
type ContractRegistryPair struct {
	Key  string
	Addr common.Address
}

// ContractRegistryABI is the input ABI used to generate the binding from.
const ContractRegistryABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"keys\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_maintainersList\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_keys\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_keys\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"setAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maintainer\",\"type\":\"address\"}],\"name\":\"setMaintainer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leaveMaintainers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"mustGetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContracts\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structContractRegistry.Pair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractRegistryBin is the compiled bytecode used for deploying new contracts.
var ContractRegistryBin = "0x608060405234801561001057600080fd5b506121d6806100206000396000f3fe60806040523480156200001157600080fd5b5060043610620000b15760003560e01c806287452314620000b65780630cb6aaf114620000cf57806313ea5d2914620000fe57806315ac72ca146200011557806336476dbf146200013b5780633fb9027114620001525780636833d54f14620001695780637d69a892146200018f5780639b2ea4bd14620001a6578063bf40fac114620001bd578063c3a2a93a14620001d4578063d520403314620001ed575b600080fd5b620000cd620000c736600462000fc8565b620001f7565b005b620000e6620000e0366004620011f6565b6200030b565b604051620000f59190620013b8565b60405180910390f35b620000cd6200010f36600462000f83565b620003b9565b6200012c6200012636600462000f83565b620004c7565b604051620000f59190620012ca565b620000cd6200014c36600462001005565b620005ea565b6200012c620001633660046200111b565b620006fc565b620001806200017a366004620011b8565b6200077b565b604051620000f59190620013ad565b620000cd620001a036600462001092565b620007bb565b620000cd620001b73660046200115e565b620008f0565b6200012c620001ce3660046200111b565b620009dd565b620001de62000a14565b604051620000f591906200132d565b620000cd62000bf5565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e89062000229903390600401620012ca565b60206040518083038186803b1580156200024257600080fd5b505afa15801562000257573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200027d9190620010f9565b620002a55760405162461bcd60e51b81526004016200029c906200141e565b60405180910390fd5b604051631b2ce7f360e11b81526001600160a01b03831690633659cfe690620002d3908490600401620012ca565b600060405180830381600087803b158015620002ee57600080fd5b505af115801562000303573d6000803e3d6000fd5b505050505050565b600281815481106200031c57600080fd5b600091825260209182902001805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815293509091830182828015620003b15780601f106200038557610100808354040283529160200191620003b1565b820191906000526020600020905b8154815290600101906020018083116200039357829003601f168201915b505050505081565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e890620003eb903390600401620012ca565b60206040518083038186803b1580156200040457600080fd5b505afa15801562000419573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200043f9190620010f9565b6200045e5760405162461bcd60e51b81526004016200029c906200141e565b600354604051639e9405c360e01b81526001600160a01b0390911690639e9405c39062000490908490600401620012ca565b600060405180830381600087803b158015620004ab57600080fd5b505af1158015620004c0573d6000803e3d6000fd5b5050505050565b600354604051630bb7c8fd60e31b81526000916001600160a01b031690635dbe47e890620004fa903390600401620012ca565b60206040518083038186803b1580156200051357600080fd5b505afa15801562000528573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200054e9190620010f9565b6200056d5760405162461bcd60e51b81526004016200029c906200141e565b816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015620005a957600080fd5b505af1158015620005be573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620005e4919062000fa9565b92915050565b600054610100900460ff16806200060657506200060662000c57565b8062000615575060005460ff16155b620006525760405162461bcd60e51b815260040180806020018281038252602e81526020018062002173602e913960400191505060405180910390fd5b600054610100900460ff161580156200067e576000805460ff1961ff0019909116610100171660011790555b836040516200068d9062000d23565b620006999190620012de565b604051809103906000f080158015620006b6573d6000803e3d6000fd5b50600380546001600160a01b0319166001600160a01b0392909216919091179055620006e38383620007bb565b8015620006f6576000805461ff00191690555b50505050565b60008060018484604051620007139291906200123d565b90815260405160209181900382018120546001600160a01b0316925082151591620007439187918791016200126b565b60405160208183030381529060405290620007735760405162461bcd60e51b81526004016200029c9190620013b8565b509392505050565b6000806001600160a01b03166001836040516200079991906200124d565b908152604051908190036020019020546001600160a01b031614159050919050565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e890620007ed903390600401620012ca565b60206040518083038186803b1580156200080657600080fd5b505afa1580156200081b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620008419190620010f9565b620008605760405162461bcd60e51b81526004016200029c906200141e565b60006040518060800160405280604d815260200162002126604d9139905081518351148190620008a55760405162461bcd60e51b81526004016200029c9190620013b8565b5060005b8351811015620006f657620008e7848281518110620008c457fe5b6020026020010151848381518110620008d957fe5b602002602001015162000c5d565b600101620008a9565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e89062000922903390600401620012ca565b60206040518083038186803b1580156200093b57600080fd5b505afa15801562000950573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620009769190620010f9565b620009955760405162461bcd60e51b81526004016200029c906200141e565b620009d883838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525085925062000c5d915050565b505050565b600060018383604051620009f39291906200123d565b908152604051908190036020019020546001600160a01b0316905092915050565b606060006002805480602002602001604051908101604052809291908181526020016000905b8282101562000aeb5760008481526020908190208301805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529283018282801562000ad65780601f1062000aaa5761010080835404028352916020019162000ad6565b820191906000526020600020905b81548152906001019060200180831162000ab857829003601f168201915b50505050508152602001906001019062000a3a565b505050509050600081516001600160401b038111801562000b0b57600080fd5b5060405190808252806020026020018201604052801562000b4957816020015b62000b3562000d31565b81526020019060019003908162000b2b5790505b50905060005b82518160ff16101562000bee576040518060400160405280848360ff168151811062000b7757fe5b602002602001015181526020016001858460ff168151811062000b9657fe5b602002602001015160405162000bad91906200124d565b908152604051908190036020019020546001600160a01b031690528251839060ff841690811062000bda57fe5b602090810291909101015260010162000b4f565b5091505090565b600354604051636196c02d60e11b81526001600160a01b039091169063c32d805a9062000c27903390600401620012ca565b600060405180830381600087803b15801562000c4257600080fd5b505af1158015620006f6573d6000803e3d6000fd5b303b1590565b6001600160a01b03811662000c865760405162461bcd60e51b81526004016200029c90620013cd565b62000c91826200077b565b62000cdd5760028054600181018255600091909152825162000cdb917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0190602085019062000d49565b505b8060018360405162000cf091906200124d565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b03199092169190911790555050565b610c20806200150683390190565b60408051808201909152606081526000602082015290565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928262000d81576000855562000dcc565b82601f1062000d9c57805160ff191683800117855562000dcc565b8280016001018555821562000dcc579182015b8281111562000dcc57825182559160200191906001019062000daf565b5062000dda92915062000dde565b5090565b5b8082111562000dda576000815560010162000ddf565b600082601f83011262000e06578081fd5b8135602062000e1f62000e19836200149f565b6200147b565b828152818101908583018385028701840188101562000e3c578586fd5b855b8581101562000e6757813562000e5481620014ec565b8452928401929084019060010162000e3e565b5090979650505050505050565b600082601f83011262000e85578081fd5b8135602062000e9862000e19836200149f565b82815281810190858301855b8581101562000e675762000ebe898684358b010162000f1a565b8452928401929084019060010162000ea4565b60008083601f84011262000ee3578182fd5b5081356001600160401b0381111562000efa578182fd5b60208301915083602082850101111562000f1357600080fd5b9250929050565b600082601f83011262000f2b578081fd5b81356001600160401b0381111562000f3f57fe5b62000f54601f8201601f19166020016200147b565b81815284602083860101111562000f69578283fd5b816020850160208301379081016020019190915292915050565b60006020828403121562000f95578081fd5b813562000fa281620014ec565b9392505050565b60006020828403121562000fbb578081fd5b815162000fa281620014ec565b6000806040838503121562000fdb578081fd5b823562000fe881620014ec565b9150602083013562000ffa81620014ec565b809150509250929050565b6000806000606084860312156200101a578081fd5b83356001600160401b038082111562001031578283fd5b6200103f8783880162000df5565b9450602086013591508082111562001055578283fd5b620010638783880162000e74565b9350604086013591508082111562001079578283fd5b50620010888682870162000df5565b9150509250925092565b60008060408385031215620010a5578182fd5b82356001600160401b0380821115620010bc578384fd5b620010ca8683870162000e74565b93506020850135915080821115620010e0578283fd5b50620010ef8582860162000df5565b9150509250929050565b6000602082840312156200110b578081fd5b8151801515811462000fa2578182fd5b600080602083850312156200112e578182fd5b82356001600160401b0381111562001144578283fd5b620011528582860162000ed1565b90969095509350505050565b60008060006040848603121562001173578283fd5b83356001600160401b0381111562001189578384fd5b620011978682870162000ed1565b9094509250506020840135620011ad81620014ec565b809150509250925092565b600060208284031215620011ca578081fd5b81356001600160401b03811115620011e0578182fd5b620011ee8482850162000f1a565b949350505050565b60006020828403121562001208578081fd5b5035919050565b6000815180845262001229816020860160208601620014bd565b601f01601f19169290920160200192915050565b6000828483379101908152919050565b6000825162001261818460208701620014bd565b9190910192915050565b60006302a3432960e51b825282846004840137507f206b657920646f6573206e6f742065786973742c206661696c656420746f2067910160048101919091526e32ba103a34329030b2323932b9b99760891b6024820152603301919050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b81811015620013215783516001600160a01b031683529284019291840191600101620012fa565b50909695505050505050565b60208082528251828201819052600091906040908185019080840286018301878501865b838110156200139f57888303603f190185528151805187855262001378888601826200120f565b918901516001600160a01b0316948901949094529487019492509086019060010162001351565b509098975050505050505050565b901515815260200190565b60006020825262000fa260208301846200120f565b60208082526031908201527f496e76616c696420616464726573732076616c75652c206661696c656420746f6040820152701039b2ba103a34329030b2323932b9b99760791b606082015260800190565b6020808252603e908201527f5468652063616c6c657220646f6573206e6f742068617665206163636573732c60408201527f206f6e6c79206d61696e7461696e6572732068617665206163636573732e0000606082015260800190565b6040518181016001600160401b03811182821017156200149757fe5b604052919050565b60006001600160401b03821115620014b357fe5b5060209081020190565b60005b83811015620014da578181015183820152602001620014c0565b83811115620006f65750506000910152565b6001600160a01b03811681146200150257600080fd5b5056fe608060405234801561001057600080fd5b50604051610c20380380610c208339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b5050505090500160405250505060006100cf6101b560201b60201c565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060005b81518110156101ae57600282828151811061013357fe5b60209081029190910181015182546001808201855560009485529284200180546001600160a01b0319166001600160a01b03909216919091179055600254845190929085908590811061018257fe5b6020908102919091018101516001600160a01b031682528101919091526040016000205560010161011c565b50506101b9565b3390565b610a58806101c86000396000f3fe608060405234801561001057600080fd5b506004361061008e5760003560e01c80630a3b0a4f1461009357806329092d0e146100cd5780635dbe47e8146100f3578063715018a6146101195780638da5cb5b14610123578063949d225d146101475780639e9405c314610161578063a39fac1214610187578063c32d805a146101df578063f2fde38b14610205575b600080fd5b6100b9600480360360208110156100a957600080fd5b50356001600160a01b031661022b565b604080519115158252519081900360200190f35b6100b9600480360360208110156100e357600080fd5b50356001600160a01b03166102bd565b6100b96004803603602081101561010957600080fd5b50356001600160a01b031661035f565b61012161037c565b005b61012b61041e565b604080516001600160a01b039092168252519081900360200190f35b61014f61042d565b60408051918252519081900360200190f35b6101216004803603602081101561017757600080fd5b50356001600160a01b0316610433565b61018f6104ec565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101cb5781810151838201526020016101b3565b505050509050019250505060405180910390f35b610121600480360360208110156101f557600080fd5b50356001600160a01b031661054e565b6101216004803603602081101561021b57600080fd5b50356001600160a01b0316610649565b6000610235610741565b6000546001600160a01b03908116911614610285576040805162461bcd60e51b81526020600482018190526024820152600080516020610a03833981519152604482015290519081900360640190fd5b6001600160a01b038216600090815260016020526040902054156102ab575060006102b8565b6102b482610745565b5060015b919050565b60006102c7610741565b6000546001600160a01b03908116911614610317576040805162461bcd60e51b81526020600482018190526024820152600080516020610a03833981519152604482015290519081900360640190fd5b6001600160a01b03821660009081526001602052604090205480158061033e575060025481115b1561034d5760009150506102b8565b610356836107a6565b50600192915050565b6001600160a01b0316600090815260016020526040902054151590565b610384610741565b6000546001600160a01b039081169116146103d4576040805162461bcd60e51b81526020600482018190526024820152600080516020610a03833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b60025490565b61043b610741565b6000546001600160a01b0390811691161461048b576040805162461bcd60e51b81526020600482018190526024820152600080516020610a03833981519152604482015290519081900360640190fd5b6001600160a01b038116600090815260016020526040902054156104e05760405162461bcd60e51b815260040180806020018281038252606481526020018061094d6064913960800191505060405180910390fd5b6104e981610745565b50565b6060600280548060200260200160405190810160405280929190818152602001828054801561054457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610526575b5050505050905090565b610556610741565b6000546001600160a01b039081169116146105a6576040805162461bcd60e51b81526020600482018190526024820152600080516020610a03833981519152604482015290519081900360640190fd5b6001600160a01b038116600090815260016020526040902054806105fb5760405162461bcd60e51b81526004018080602001828103825260528152602001806109b16052913960600191505060405180910390fd5b60025481111561063c5760405162461bcd60e51b81526004018080602001828103825260638152602001806108ea6063913960800191505060405180910390fd5b610645826107a6565b5050565b610651610741565b6000546001600160a01b039081169116146106a1576040805162461bcd60e51b81526020600482018190526024820152600080516020610a03833981519152604482015290519081900360640190fd5b6001600160a01b0381166106e65760405162461bcd60e51b81526004018080602001828103825260268152602001806108c46026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b60028054600181810183557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace90910180546001600160a01b039094166001600160a01b03199094168417905590546000928352602091909152604090912055565b6001600160a01b03811660009081526001602081815260408084208151928301909152548082526002805492946000198401949293909190859081106107e857fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902055600280548290811061081b57fe5b6000918252602090912001548251600280546001600160a01b039093169290916000190190811061084857fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600280548061088157fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0394909416815260019093525050604081205556fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e76616c696420696e6465782076616c756520666f722074686520616464726573732073746f726167652c206661696c656420746f2072656d6f76652074686520616464726573732066726f6d2074686520616464726573732073746f726167652e54686520616464726573732068617320616c7265616479206265656e20616464656420746f207468652073746f726167652c206661696c656420746f2061646420746865206164647265737320746f2074686520616464726573732073746f726167652e546865206164647265737320646f6573206e6f742065786973742c206661696c656420746f2072656d6f76652074686520616464726573732066726f6d2074686520616464726573732073746f726167652e4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a264697066735822122083626d47adb3491349ee9a4368a52afe61478e893a6e65791394413c25ec40ea64736f6c63430007060033546865206e756d626572206f66206b65797320616e64206164647265737365732073686f756c64206265207468652073616d652c206661696c656420746f20736574206164647265737365732e496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a26469706673582212203a5a83aa23c33b57eb0f802672da69505c83a5f8bb0256ca6a11c72d88c67fe564736f6c63430007060033"

// DeployContractRegistry deploys a new Ethereum contract, binding an instance of ContractRegistry to it.
func DeployContractRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractRegistry{ContractRegistryCaller: ContractRegistryCaller{contract: contract}, ContractRegistryTransactor: ContractRegistryTransactor{contract: contract}, ContractRegistryFilterer: ContractRegistryFilterer{contract: contract}}, nil
}

// ContractRegistry is an auto generated Go binding around an Ethereum contract.
type ContractRegistry struct {
	ContractRegistryCaller     // Read-only binding to the contract
	ContractRegistryTransactor // Write-only binding to the contract
	ContractRegistryFilterer   // Log filterer for contract events
}

// ContractRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractRegistrySession struct {
	Contract     *ContractRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractRegistryCallerSession struct {
	Contract *ContractRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ContractRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractRegistryTransactorSession struct {
	Contract     *ContractRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRegistryRaw struct {
	Contract *ContractRegistry // Generic contract binding to access the raw methods on
}

// ContractRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractRegistryCallerRaw struct {
	Contract *ContractRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractRegistryTransactorRaw struct {
	Contract *ContractRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractRegistry creates a new instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistry(address common.Address, backend bind.ContractBackend) (*ContractRegistry, error) {
	contract, err := bindContractRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractRegistry{ContractRegistryCaller: ContractRegistryCaller{contract: contract}, ContractRegistryTransactor: ContractRegistryTransactor{contract: contract}, ContractRegistryFilterer: ContractRegistryFilterer{contract: contract}}, nil
}

// NewContractRegistryCaller creates a new read-only instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractRegistryCaller, error) {
	contract, err := bindContractRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCaller{contract: contract}, nil
}

// NewContractRegistryTransactor creates a new write-only instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractRegistryTransactor, error) {
	contract, err := bindContractRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryTransactor{contract: contract}, nil
}

// NewContractRegistryFilterer creates a new log filterer instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractRegistryFilterer, error) {
	contract, err := bindContractRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryFilterer{contract: contract}, nil
}

// bindContractRegistry binds a generic wrapper to an already deployed contract.
func bindContractRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistry *ContractRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractRegistry.Contract.ContractRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistry *ContractRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistry.Contract.ContractRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistry *ContractRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistry.Contract.ContractRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistry *ContractRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistry *ContractRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistry *ContractRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistry.Contract.contract.Transact(opts, method, params...)
}

// Contains is a free data retrieval call binding the contract method 0x6833d54f.
//
// Solidity: function contains(string _key) view returns(bool)
func (_ContractRegistry *ContractRegistryCaller) Contains(opts *bind.CallOpts, _key string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "contains", _key)
	return *ret0, err
}

// Contains is a free data retrieval call binding the contract method 0x6833d54f.
//
// Solidity: function contains(string _key) view returns(bool)
func (_ContractRegistry *ContractRegistrySession) Contains(_key string) (bool, error) {
	return _ContractRegistry.Contract.Contains(&_ContractRegistry.CallOpts, _key)
}

// Contains is a free data retrieval call binding the contract method 0x6833d54f.
//
// Solidity: function contains(string _key) view returns(bool)
func (_ContractRegistry *ContractRegistryCallerSession) Contains(_key string) (bool, error) {
	return _ContractRegistry.Contract.Contains(&_ContractRegistry.CallOpts, _key)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string _key) view returns(address)
func (_ContractRegistry *ContractRegistryCaller) GetAddress(opts *bind.CallOpts, _key string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "getAddress", _key)
	return *ret0, err
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string _key) view returns(address)
func (_ContractRegistry *ContractRegistrySession) GetAddress(_key string) (common.Address, error) {
	return _ContractRegistry.Contract.GetAddress(&_ContractRegistry.CallOpts, _key)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string _key) view returns(address)
func (_ContractRegistry *ContractRegistryCallerSession) GetAddress(_key string) (common.Address, error) {
	return _ContractRegistry.Contract.GetAddress(&_ContractRegistry.CallOpts, _key)
}

// GetContracts is a free data retrieval call binding the contract method 0xc3a2a93a.
//
// Solidity: function getContracts() view returns((string,address)[])
func (_ContractRegistry *ContractRegistryCaller) GetContracts(opts *bind.CallOpts) ([]ContractRegistryPair, error) {
	var (
		ret0 = new([]ContractRegistryPair)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "getContracts")
	return *ret0, err
}

// GetContracts is a free data retrieval call binding the contract method 0xc3a2a93a.
//
// Solidity: function getContracts() view returns((string,address)[])
func (_ContractRegistry *ContractRegistrySession) GetContracts() ([]ContractRegistryPair, error) {
	return _ContractRegistry.Contract.GetContracts(&_ContractRegistry.CallOpts)
}

// GetContracts is a free data retrieval call binding the contract method 0xc3a2a93a.
//
// Solidity: function getContracts() view returns((string,address)[])
func (_ContractRegistry *ContractRegistryCallerSession) GetContracts() ([]ContractRegistryPair, error) {
	return _ContractRegistry.Contract.GetContracts(&_ContractRegistry.CallOpts)
}

// Keys is a free data retrieval call binding the contract method 0x0cb6aaf1.
//
// Solidity: function keys(uint256 ) view returns(string)
func (_ContractRegistry *ContractRegistryCaller) Keys(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "keys", arg0)
	return *ret0, err
}

// Keys is a free data retrieval call binding the contract method 0x0cb6aaf1.
//
// Solidity: function keys(uint256 ) view returns(string)
func (_ContractRegistry *ContractRegistrySession) Keys(arg0 *big.Int) (string, error) {
	return _ContractRegistry.Contract.Keys(&_ContractRegistry.CallOpts, arg0)
}

// Keys is a free data retrieval call binding the contract method 0x0cb6aaf1.
//
// Solidity: function keys(uint256 ) view returns(string)
func (_ContractRegistry *ContractRegistryCallerSession) Keys(arg0 *big.Int) (string, error) {
	return _ContractRegistry.Contract.Keys(&_ContractRegistry.CallOpts, arg0)
}

// MustGetAddress is a free data retrieval call binding the contract method 0x3fb90271.
//
// Solidity: function mustGetAddress(string _key) view returns(address)
func (_ContractRegistry *ContractRegistryCaller) MustGetAddress(opts *bind.CallOpts, _key string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "mustGetAddress", _key)
	return *ret0, err
}

// MustGetAddress is a free data retrieval call binding the contract method 0x3fb90271.
//
// Solidity: function mustGetAddress(string _key) view returns(address)
func (_ContractRegistry *ContractRegistrySession) MustGetAddress(_key string) (common.Address, error) {
	return _ContractRegistry.Contract.MustGetAddress(&_ContractRegistry.CallOpts, _key)
}

// MustGetAddress is a free data retrieval call binding the contract method 0x3fb90271.
//
// Solidity: function mustGetAddress(string _key) view returns(address)
func (_ContractRegistry *ContractRegistryCallerSession) MustGetAddress(_key string) (common.Address, error) {
	return _ContractRegistry.Contract.MustGetAddress(&_ContractRegistry.CallOpts, _key)
}

// GetImplementation is a paid mutator transaction binding the contract method 0x15ac72ca.
//
// Solidity: function getImplementation(address _proxy) returns(address)
func (_ContractRegistry *ContractRegistryTransactor) GetImplementation(opts *bind.TransactOpts, _proxy common.Address) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "getImplementation", _proxy)
}

// GetImplementation is a paid mutator transaction binding the contract method 0x15ac72ca.
//
// Solidity: function getImplementation(address _proxy) returns(address)
func (_ContractRegistry *ContractRegistrySession) GetImplementation(_proxy common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.GetImplementation(&_ContractRegistry.TransactOpts, _proxy)
}

// GetImplementation is a paid mutator transaction binding the contract method 0x15ac72ca.
//
// Solidity: function getImplementation(address _proxy) returns(address)
func (_ContractRegistry *ContractRegistryTransactorSession) GetImplementation(_proxy common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.GetImplementation(&_ContractRegistry.TransactOpts, _proxy)
}

// Initialize is a paid mutator transaction binding the contract method 0x36476dbf.
//
// Solidity: function initialize(address[] _maintainersList, string[] _keys, address[] _addresses) returns()
func (_ContractRegistry *ContractRegistryTransactor) Initialize(opts *bind.TransactOpts, _maintainersList []common.Address, _keys []string, _addresses []common.Address) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "initialize", _maintainersList, _keys, _addresses)
}

// Initialize is a paid mutator transaction binding the contract method 0x36476dbf.
//
// Solidity: function initialize(address[] _maintainersList, string[] _keys, address[] _addresses) returns()
func (_ContractRegistry *ContractRegistrySession) Initialize(_maintainersList []common.Address, _keys []string, _addresses []common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.Initialize(&_ContractRegistry.TransactOpts, _maintainersList, _keys, _addresses)
}

// Initialize is a paid mutator transaction binding the contract method 0x36476dbf.
//
// Solidity: function initialize(address[] _maintainersList, string[] _keys, address[] _addresses) returns()
func (_ContractRegistry *ContractRegistryTransactorSession) Initialize(_maintainersList []common.Address, _keys []string, _addresses []common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.Initialize(&_ContractRegistry.TransactOpts, _maintainersList, _keys, _addresses)
}

// LeaveMaintainers is a paid mutator transaction binding the contract method 0xd5204033.
//
// Solidity: function leaveMaintainers() returns()
func (_ContractRegistry *ContractRegistryTransactor) LeaveMaintainers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "leaveMaintainers")
}

// LeaveMaintainers is a paid mutator transaction binding the contract method 0xd5204033.
//
// Solidity: function leaveMaintainers() returns()
func (_ContractRegistry *ContractRegistrySession) LeaveMaintainers() (*types.Transaction, error) {
	return _ContractRegistry.Contract.LeaveMaintainers(&_ContractRegistry.TransactOpts)
}

// LeaveMaintainers is a paid mutator transaction binding the contract method 0xd5204033.
//
// Solidity: function leaveMaintainers() returns()
func (_ContractRegistry *ContractRegistryTransactorSession) LeaveMaintainers() (*types.Transaction, error) {
	return _ContractRegistry.Contract.LeaveMaintainers(&_ContractRegistry.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _key, address _addr) returns()
func (_ContractRegistry *ContractRegistryTransactor) SetAddress(opts *bind.TransactOpts, _key string, _addr common.Address) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "setAddress", _key, _addr)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _key, address _addr) returns()
func (_ContractRegistry *ContractRegistrySession) SetAddress(_key string, _addr common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.SetAddress(&_ContractRegistry.TransactOpts, _key, _addr)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _key, address _addr) returns()
func (_ContractRegistry *ContractRegistryTransactorSession) SetAddress(_key string, _addr common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.SetAddress(&_ContractRegistry.TransactOpts, _key, _addr)
}

// SetAddresses is a paid mutator transaction binding the contract method 0x7d69a892.
//
// Solidity: function setAddresses(string[] _keys, address[] _addresses) returns()
func (_ContractRegistry *ContractRegistryTransactor) SetAddresses(opts *bind.TransactOpts, _keys []string, _addresses []common.Address) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "setAddresses", _keys, _addresses)
}

// SetAddresses is a paid mutator transaction binding the contract method 0x7d69a892.
//
// Solidity: function setAddresses(string[] _keys, address[] _addresses) returns()
func (_ContractRegistry *ContractRegistrySession) SetAddresses(_keys []string, _addresses []common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.SetAddresses(&_ContractRegistry.TransactOpts, _keys, _addresses)
}

// SetAddresses is a paid mutator transaction binding the contract method 0x7d69a892.
//
// Solidity: function setAddresses(string[] _keys, address[] _addresses) returns()
func (_ContractRegistry *ContractRegistryTransactorSession) SetAddresses(_keys []string, _addresses []common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.SetAddresses(&_ContractRegistry.TransactOpts, _keys, _addresses)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address _maintainer) returns()
func (_ContractRegistry *ContractRegistryTransactor) SetMaintainer(opts *bind.TransactOpts, _maintainer common.Address) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "setMaintainer", _maintainer)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address _maintainer) returns()
func (_ContractRegistry *ContractRegistrySession) SetMaintainer(_maintainer common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.SetMaintainer(&_ContractRegistry.TransactOpts, _maintainer)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address _maintainer) returns()
func (_ContractRegistry *ContractRegistryTransactorSession) SetMaintainer(_maintainer common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.SetMaintainer(&_ContractRegistry.TransactOpts, _maintainer)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x00874523.
//
// Solidity: function upgradeContract(address _proxy, address _newImplementation) returns()
func (_ContractRegistry *ContractRegistryTransactor) UpgradeContract(opts *bind.TransactOpts, _proxy common.Address, _newImplementation common.Address) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "upgradeContract", _proxy, _newImplementation)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x00874523.
//
// Solidity: function upgradeContract(address _proxy, address _newImplementation) returns()
func (_ContractRegistry *ContractRegistrySession) UpgradeContract(_proxy common.Address, _newImplementation common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.UpgradeContract(&_ContractRegistry.TransactOpts, _proxy, _newImplementation)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x00874523.
//
// Solidity: function upgradeContract(address _proxy, address _newImplementation) returns()
func (_ContractRegistry *ContractRegistryTransactorSession) UpgradeContract(_proxy common.Address, _newImplementation common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.UpgradeContract(&_ContractRegistry.TransactOpts, _proxy, _newImplementation)
}
