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

// ContractRegistryPair is an auto generated low-level Go binding around an user-defined struct.
type ContractRegistryPair struct {
	Key  string
	Addr common.Address
}

// ContractRegistryMetaData contains all meta data concerning the ContractRegistry contract.
var ContractRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"keys\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_maintainersList\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_keys\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maintainer\",\"type\":\"address\"}],\"name\":\"setMaintainer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leaveMaintainers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"mustGetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContracts\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structContractRegistry.Pair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_keys\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"setAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612448806100206000396000f3fe60806040523480156200001157600080fd5b5060043610620000b15760003560e01c806287452314620000b65780630cb6aaf114620000cf57806313ea5d2914620000fe57806315ac72ca146200011557806336476dbf146200013b5780633fb9027114620001525780636833d54f14620001695780637d69a892146200018f5780639b2ea4bd14620001a6578063bf40fac114620001bd578063c3a2a93a14620001d4578063d520403314620001ed575b600080fd5b620000cd620000c7366004620010d3565b620001f7565b005b620000e6620000e036600462001301565b6200030b565b604051620000f59190620014e9565b60405180910390f35b620000cd6200010f3660046200108e565b620003b9565b6200012c620001263660046200108e565b620004c7565b604051620000f59190620013fb565b620000cd6200014c36600462001110565b620005ea565b6200012c6200016336600462001226565b620006fc565b620001806200017a366004620012c3565b6200077b565b604051620000f59190620014de565b620000cd620001a03660046200119d565b620007bb565b620000cd620001b736600462001269565b620008d3565b6200012c620001ce36600462001226565b620009bb565b620001de620009f2565b604051620000f591906200145e565b620000cd62000bd3565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e89062000229903390600401620013fb565b60206040518083038186803b1580156200024257600080fd5b505afa15801562000257573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200027d919062001204565b620002a55760405162461bcd60e51b81526004016200029c906200164d565b60405180910390fd5b604051631b2ce7f360e11b81526001600160a01b03831690633659cfe690620002d3908490600401620013fb565b600060405180830381600087803b158015620002ee57600080fd5b505af115801562000303573d6000803e3d6000fd5b505050505050565b600281815481106200031c57600080fd5b600091825260209182902001805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815293509091830182828015620003b15780601f106200038557610100808354040283529160200191620003b1565b820191906000526020600020905b8154815290600101906020018083116200039357829003601f168201915b505050505081565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e890620003eb903390600401620013fb565b60206040518083038186803b1580156200040457600080fd5b505afa15801562000419573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200043f919062001204565b6200045e5760405162461bcd60e51b81526004016200029c906200164d565b600354604051639e9405c360e01b81526001600160a01b0390911690639e9405c39062000490908490600401620013fb565b600060405180830381600087803b158015620004ab57600080fd5b505af1158015620004c0573d6000803e3d6000fd5b5050505050565b600354604051630bb7c8fd60e31b81526000916001600160a01b031690635dbe47e890620004fa903390600401620013fb565b60206040518083038186803b1580156200051357600080fd5b505afa15801562000528573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200054e919062001204565b6200056d5760405162461bcd60e51b81526004016200029c906200164d565b816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015620005a957600080fd5b505af1158015620005be573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620005e49190620010b4565b92915050565b600054610100900460ff16806200060657506200060662000d62565b8062000615575060005460ff16155b620006525760405162461bcd60e51b815260040180806020018281038252602e815260200180620023e5602e913960400191505060405180910390fd5b600054610100900460ff161580156200067e576000805460ff1961ff0019909116610100171660011790555b836040516200068d9062000e2e565b6200069991906200140f565b604051809103906000f080158015620006b6573d6000803e3d6000fd5b50600380546001600160a01b0319166001600160a01b0392909216919091179055620006e38383620007bb565b8015620006f6576000805461ff00191690555b50505050565b600080600184846040516200071392919062001361565b90815260405160209181900382018120546001600160a01b0316925082151591620007439187918791016200138f565b60405160208183030381529060405290620007735760405162461bcd60e51b81526004016200029c9190620014e9565b509392505050565b6000806001600160a01b031660018360405162000799919062001371565b908152604051908190036020019020546001600160a01b031614159050919050565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e890620007ed903390600401620013fb565b60206040518083038186803b1580156200080657600080fd5b505afa1580156200081b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000841919062001204565b620008605760405162461bcd60e51b81526004016200029c906200164d565b8051825114620008845760405162461bcd60e51b81526004016200029c90620014fe565b60005b8251811015620008ce57620008c5838281518110620008a257fe5b6020026020010151838381518110620008b757fe5b602002602001015162000d68565b60010162000887565b505050565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e89062000905903390600401620013fb565b60206040518083038186803b1580156200091e57600080fd5b505afa15801562000933573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000959919062001204565b620009785760405162461bcd60e51b81526004016200029c906200164d565b620008ce83838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525085925062000d68915050565b600060018383604051620009d192919062001361565b908152604051908190036020019020546001600160a01b0316905092915050565b606060006002805480602002602001604051908101604052809291908181526020016000905b8282101562000ac95760008481526020908190208301805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529283018282801562000ab45780601f1062000a885761010080835404028352916020019162000ab4565b820191906000526020600020905b81548152906001019060200180831162000a9657829003601f168201915b50505050508152602001906001019062000a18565b505050509050600081516001600160401b038111801562000ae957600080fd5b5060405190808252806020026020018201604052801562000b2757816020015b62000b1362000e3c565b81526020019060019003908162000b095790505b50905060005b82518160ff16101562000bcc576040518060400160405280848360ff168151811062000b5557fe5b602002602001015181526020016001858460ff168151811062000b7457fe5b602002602001015160405162000b8b919062001371565b908152604051908190036020019020546001600160a01b031690528251839060ff841690811062000bb857fe5b602090810291909101015260010162000b2d565b5091505090565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e89062000c05903390600401620013fb565b60206040518083038186803b15801562000c1e57600080fd5b505afa15801562000c33573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000c59919062001204565b1562000d00576003546040805163949d225d60e01b815290516001926001600160a01b03169163949d225d916004808301926020929190829003018186803b15801562000ca557600080fd5b505afa15801562000cba573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000ce091906200131a565b1162000d005760405162461bcd60e51b81526004016200029c90620015db565b600354604051636196c02d60e11b81526001600160a01b039091169063c32d805a9062000d32903390600401620013fb565b600060405180830381600087803b15801562000d4d57600080fd5b505af1158015620006f6573d6000803e3d6000fd5b303b1590565b6001600160a01b03811662000d915760405162461bcd60e51b81526004016200029c906200157e565b62000d9c826200077b565b62000de85760028054600181018255600091909152825162000de6917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0190602085019062000e54565b505b8060018360405162000dfb919062001371565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b03199092169190911790555050565b610cb0806200173583390190565b60408051808201909152606081526000602082015290565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928262000e8c576000855562000ed7565b82601f1062000ea757805160ff191683800117855562000ed7565b8280016001018555821562000ed7579182015b8281111562000ed757825182559160200191906001019062000eba565b5062000ee592915062000ee9565b5090565b5b8082111562000ee5576000815560010162000eea565b600082601f83011262000f11578081fd5b8135602062000f2a62000f2483620016ce565b620016aa565b828152818101908583018385028701840188101562000f47578586fd5b855b8581101562000f7257813562000f5f816200171b565b8452928401929084019060010162000f49565b5090979650505050505050565b600082601f83011262000f90578081fd5b8135602062000fa362000f2483620016ce565b82815281810190858301855b8581101562000f725762000fc9898684358b010162001025565b8452928401929084019060010162000faf565b60008083601f84011262000fee578182fd5b5081356001600160401b0381111562001005578182fd5b6020830191508360208285010111156200101e57600080fd5b9250929050565b600082601f83011262001036578081fd5b81356001600160401b038111156200104a57fe5b6200105f601f8201601f1916602001620016aa565b81815284602083860101111562001074578283fd5b816020850160208301379081016020019190915292915050565b600060208284031215620010a0578081fd5b8135620010ad816200171b565b9392505050565b600060208284031215620010c6578081fd5b8151620010ad816200171b565b60008060408385031215620010e6578081fd5b8235620010f3816200171b565b9150602083013562001105816200171b565b809150509250929050565b60008060006060848603121562001125578081fd5b83356001600160401b03808211156200113c578283fd5b6200114a8783880162000f00565b9450602086013591508082111562001160578283fd5b6200116e8783880162000f7f565b9350604086013591508082111562001184578283fd5b50620011938682870162000f00565b9150509250925092565b60008060408385031215620011b0578182fd5b82356001600160401b0380821115620011c7578384fd5b620011d58683870162000f7f565b93506020850135915080821115620011eb578283fd5b50620011fa8582860162000f00565b9150509250929050565b60006020828403121562001216578081fd5b81518015158114620010ad578182fd5b6000806020838503121562001239578182fd5b82356001600160401b038111156200124f578283fd5b6200125d8582860162000fdc565b90969095509350505050565b6000806000604084860312156200127e578283fd5b83356001600160401b0381111562001294578384fd5b620012a28682870162000fdc565b9094509250506020840135620012b8816200171b565b809150509250925092565b600060208284031215620012d5578081fd5b81356001600160401b03811115620012eb578182fd5b620012f98482850162001025565b949350505050565b60006020828403121562001313578081fd5b5035919050565b6000602082840312156200132c578081fd5b5051919050565b600081518084526200134d816020860160208601620016ec565b601f01601f19169290920160200192915050565b6000828483379101908152919050565b6000825162001385818460208701620016ec565b9190910192915050565b60007002da8a2a19698199a181819ae96aa34329607d1b825282846011840137507f206b657920646f6573206e6f742065786973742c206661696c656420746f2067910160118101919091526e32ba103a34329030b2323932b9b99760891b6031820152604001919050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b81811015620014525783516001600160a01b0316835292840192918401916001016200142b565b50909695505050505050565b60208082528251828201819052600091906040908185019080840286018301878501865b83811015620014d057888303603f1901855281518051878552620014a98886018262001333565b918901516001600160a01b0316948901949094529487019492509086019060010162001482565b509098975050505050505050565b901515815260200190565b600060208252620010ad602083018462001333565b6020808252605a908201527f5b5145432d3033343030315d2d546865206e756d626572206f66206b6579732060408201527f616e64206164647265737365732073686f756c64206265207468652073616d6560608201527916103330b4b632b2103a379039b2ba1030b2323932b9b9b2b99760311b608082015260a00190565b6020808252603e908201527f5b5145432d3033343030305d2d496e76616c696420616464726573732076616c60408201527f75652c206661696c656420746f207365742074686520616464726573732e0000606082015260800190565b6020808252604c908201527f5b5145432d3033343030325d2d43616e6e6f74206c6561766520746865206d6160408201527f696e7461696e657273206c6973742c20796f752061726520746865206f6e6c7960608201526b1036b0b4b73a30b4b732b91760a11b608082015260a00190565b6020808252603e908201527f5b5145432d3033343030345d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c79206d61696e7461696e6572732068617665206163636573732e0000606082015260800190565b6040518181016001600160401b0381118282101715620016c657fe5b604052919050565b60006001600160401b03821115620016e257fe5b5060209081020190565b60005b8381101562001709578181015183820152602001620016ef565b83811115620006f65750506000910152565b6001600160a01b03811681146200173157600080fd5b5056fe608060405234801561001057600080fd5b50604051610cb0380380610cb08339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b5050505090500160405250505060006100cf6101b560201b60201c565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060005b81518110156101ae57600282828151811061013357fe5b60209081029190910181015182546001808201855560009485529284200180546001600160a01b0319166001600160a01b03909216919091179055600254845190929085908590811061018257fe5b6020908102919091018101516001600160a01b031682528101919091526040016000205560010161011c565b50506101b9565b3390565b610ae8806101c86000396000f3fe608060405234801561001057600080fd5b50600436106100995760003560e01c80630a3b0a4f1461009e57806329092d0e146100d857806352efea6e146100fe5780635dbe47e814610106578063715018a61461012c5780638da5cb5b14610136578063949d225d1461015a5780639e9405c314610174578063a39fac121461019a578063c32d805a146101f2578063f2fde38b14610218575b600080fd5b6100c4600480360360208110156100b457600080fd5b50356001600160a01b031661023e565b604080519115158252519081900360200190f35b6100c4600480360360208110156100ee57600080fd5b50356001600160a01b03166102d0565b6100c4610372565b6100c46004803603602081101561011c57600080fd5b50356001600160a01b031661042a565b610134610447565b005b61013e6104e9565b604080516001600160a01b039092168252519081900360200190f35b6101626104f8565b60408051918252519081900360200190f35b6101346004803603602081101561018a57600080fd5b50356001600160a01b03166104fe565b6101a261059d565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101de5781810151838201526020016101c6565b505050509050019250505060405180910390f35b6101346004803603602081101561020857600080fd5b50356001600160a01b03166105ff565b6101346004803603602081101561022e57600080fd5b50356001600160a01b031661069b565b6000610248610793565b6000546001600160a01b03908116911614610298576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b6001600160a01b038216600090815260016020526040902054156102be575060006102cb565b6102c782610797565b5060015b919050565b60006102da610793565b6000546001600160a01b0390811691161461032a576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b6001600160a01b038216600090815260016020526040902054801580610351575060025481115b156103605760009150506102cb565b610369836107fe565b50600192915050565b600061037c610793565b6000546001600160a01b039081169116146103cc576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b60005b6002548110156104175760016000600283815481106103ea57fe5b60009182526020808320909101546001600160a01b031683528201929092526040018120556001016103cf565b5061042460026000610986565b50600190565b6001600160a01b0316600090815260016020526040902054151590565b61044f610793565b6000546001600160a01b0390811691161461049f576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b60025490565b610506610793565b6000546001600160a01b03908116911614610556576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b61055f8161023e565b61059a5760405162461bcd60e51b8152600401808060200182810382526071815260200180610a226071913960800191505060405180910390fd5b50565b606060028054806020026020016040519081016040528092919081815260200182805480156105f557602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116105d7575b5050505050905090565b610607610793565b6000546001600160a01b03908116911614610657576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b610660816102d0565b61059a5760405162461bcd60e51b81526004018080602001828103825260438152602001806109df6043913960600191505060405180910390fd5b6106a3610793565b6000546001600160a01b039081169116146106f3576040805162461bcd60e51b81526020600482018190526024820152600080516020610a93833981519152604482015290519081900360640190fd5b6001600160a01b0381166107385760405162461bcd60e51b81526004018080602001828103825260268152602001806109b96026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b60028054600180820183557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace90910180546001600160a01b0319166001600160a01b0385169081179091559154600092835260209190915260409091205561059a81610938565b6001600160a01b038116600090815260016020908152604080832081519283019091525481526002805491926000198301929091908390811061083d57fe5b60009182526020909120015483516001600160a01b0390911691506000190182146108c15782516001600160a01b03821660009081526001602052604090205582516002805483926000190190811061089257fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b60028054806108cc57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038616825260019052604081205561091084610938565b836001600160a01b0316816001600160a01b0316146109325761093281610938565b50505050565b6001600160a01b03811660009081526001602052604090205460025481111561095d57fe5b6109668261042a565b1561097a576000811161097557fe5b610982565b801561098257fe5b5050565b508054600082559060005260206000209081019061059a91905b808211156109b457600081556001016109a0565b509056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573735b5145432d3033353030305d2d4661696c656420746f2072656d6f76652074686520616464726573732066726f6d2074686520616464726573732073746f726167652e5b5145432d3033353030325d2d54686520616464726573732068617320616c7265616479206265656e20616464656420746f207468652073746f726167652c206661696c656420746f2061646420746865206164647265737320746f2074686520616464726573732073746f726167652e4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a26469706673582212203421bef274aa9bcf497e4b6a0bb1e7be50a80a19ba4c0e2f760c17b87dc2d41c64736f6c63430007060033496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a2646970667358221220389a51e4c3c5ce481ce3cfc2e008860459f9912cd8a49674b584a3eea840897864736f6c63430007060033",
}

// ContractRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractRegistryMetaData.ABI instead.
var ContractRegistryABI = ContractRegistryMetaData.ABI

// ContractRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractRegistryMetaData.Bin instead.
var ContractRegistryBin = ContractRegistryMetaData.Bin

// DeployContractRegistry deploys a new Ethereum contract, binding an instance of ContractRegistry to it.
func DeployContractRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractRegistry, error) {
	parsed, err := ContractRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractRegistryBin), backend)
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
func (_ContractRegistry *ContractRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_ContractRegistry *ContractRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
	var out []interface{}
	err := _ContractRegistry.contract.Call(opts, &out, "contains", _key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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
	var out []interface{}
	err := _ContractRegistry.contract.Call(opts, &out, "getAddress", _key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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
	var out []interface{}
	err := _ContractRegistry.contract.Call(opts, &out, "getContracts")

	if err != nil {
		return *new([]ContractRegistryPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]ContractRegistryPair)).(*[]ContractRegistryPair)

	return out0, err

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
	var out []interface{}
	err := _ContractRegistry.contract.Call(opts, &out, "keys", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

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
	var out []interface{}
	err := _ContractRegistry.contract.Call(opts, &out, "mustGetAddress", _key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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
