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

// EPDRMembershipABI is the input ABI used to generate the binding from.
const EPDRMembershipABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"addMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expertLimitKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expertVotingKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"experts\",\"outputs\":[{\"internalType\":\"contractAddressStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"isMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"removeMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldExp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newExp\",\"type\":\"address\"}],\"name\":\"swapMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_experts\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EPDRMembershipBin is the compiled bytecode used for deploying new contracts.
var EPDRMembershipBin = "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b6121938061007d6000396000f3fe60806040523480156200001157600080fd5b5060043610620000ca5760003560e01c80630b1ca49a14620000cf5780633e02181414620000e85780635313263a146200010a578063715018a614620001145780638da5cb5b146200011e578063946d920414620001375780639eab5253146200014e578063a230c5241462000167578063b278982d146200018d578063b295a00e14620001a4578063c683682514620001bd578063ca6d56dc14620001c7578063de8fa43114620001de578063f2fde38b14620001e8575b600080fd5b620000e6620000e036600462000f23565b620001ff565b005b620000f26200034d565b604051620001019190620011b1565b60405180910390f35b620000f2620003df565b620000e66200043d565b62000128620004f4565b60405162000101919062001143565b620000e66200014836600462000fa5565b62000503565b620001586200058b565b60405162000101919062001157565b6200017e6200017836600462000f23565b62000615565b604051620001019190620011a6565b620000e66200019e36600462000f68565b620006a2565b620001ae62000870565b60405162000101919062001451565b62000128620009a0565b620000e6620001d836600462000f23565b620009af565b620001ae62000b98565b620000e6620001f936600462000f23565b62000c19565b600254604051633fb9027160e01b81526001600160a01b0390911690633fb90271906200023190600490810162001207565b60206040518083038186803b1580156200024a57600080fd5b505afa1580156200025f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000285919062000f49565b6001600160a01b0316336001600160a01b031614620002c15760405162461bcd60e51b8152600401620002b890620013e0565b60405180910390fd5b600154604051631484968760e11b81526001600160a01b03909116906329092d0e90620002f390849060040162001143565b602060405180830381600087803b1580156200030e57600080fd5b505af115801562000323573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000349919062001108565b5050565b6003805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015620003d75780601f10620003ab57610100808354040283529160200191620003d7565b820191906000526020600020905b815481529060010190602001808311620003b957829003601f168201915b505050505081565b6004805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015620003d75780601f10620003ab57610100808354040283529160200191620003d7565b6200044762000d28565b6000546001600160a01b03908116911614620004aa576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b6200050f828262000d2c565b60408051808201909152601d8082527f636f6e737469747574696f6e2e455044522e6d61784e457870657274730000006020909201918252620005559160039162000e69565b5060405180606001604052806028815260200162002136602891398051620005869160049160209091019062000e69565b505050565b600154604080516351cfd60960e11b815290516060926001600160a01b03169163a39fac12916004808301926000929190829003018186803b158015620005d157600080fd5b505afa158015620005e6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000610919081019062001065565b905090565b600154604051630bb7c8fd60e31b81526000916001600160a01b031690635dbe47e8906200064890859060040162001143565b60206040518083038186803b1580156200066157600080fd5b505afa15801562000676573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200069c919062001108565b92915050565b600254604051633fb9027160e01b81526001600160a01b0390911690633fb9027190620006d490600490810162001207565b60206040518083038186803b158015620006ed57600080fd5b505afa15801562000702573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000728919062000f49565b6001600160a01b0316336001600160a01b0316146200075b5760405162461bcd60e51b8152600401620002b890620013e0565b600154604051631484968760e11b81526001600160a01b039091169081906329092d0e906200078f90869060040162001143565b602060405180830381600087803b158015620007aa57600080fd5b505af1158015620007bf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620007e5919062001108565b50604051630a3b0a4f60e01b81526001600160a01b03821690630a3b0a4f906200081490859060040162001143565b602060405180830381600087803b1580156200082f57600080fd5b505af115801562000844573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200086a919062001108565b50505050565b600254604051633fb9027160e01b815260009182916001600160a01b0390911690633fb9027190620008a5906004016200139e565b60206040518083038186803b158015620008be57600080fd5b505afa158015620008d3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620008f9919062000f49565b6001600160a01b031663498bff0060036040518263ffffffff1660e01b815260040162000927919062001207565b60206040518083038186803b1580156200094057600080fd5b505afa15801562000955573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200097b91906200112a565b905060008111620006105760405162461bcd60e51b8152600401620002b89062001299565b6001546001600160a01b031681565b600254604051633fb9027160e01b81526001600160a01b0390911690633fb9027190620009e190600490810162001207565b60206040518083038186803b158015620009fa57600080fd5b505afa15801562000a0f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000a35919062000f49565b6001600160a01b0316336001600160a01b03161462000a685760405162461bcd60e51b8152600401620002b890620013e0565b6001546001600160a01b031662000a7e62000870565b816001600160a01b031663949d225d6040518163ffffffff1660e01b815260040160206040518083038186803b15801562000ab857600080fd5b505afa15801562000acd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000af391906200112a565b111562000b145760405162461bcd60e51b8152600401620002b89062001326565b604051630a3b0a4f60e01b81526001600160a01b03821690630a3b0a4f9062000b4290859060040162001143565b602060405180830381600087803b15801562000b5d57600080fd5b505af115801562000b72573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000586919062001108565b6001546040805163949d225d60e01b815290516000926001600160a01b03169163949d225d916004808301926020929190829003018186803b15801562000bde57600080fd5b505afa15801562000bf3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200061091906200112a565b62000c2362000d28565b6000546001600160a01b0390811691161462000c86576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b03811662000ccd5760405162461bcd60e51b8152600401808060200182810382526026815260200180620020e26026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b600054600160a81b900460ff168062000d4a575062000d4a62000e63565b8062000d605750600054600160a01b900460ff16155b62000d9d5760405162461bcd60e51b815260040180806020018281038252602e81526020018062002108602e913960400191505060405180910390fd5b600054600160a81b900460ff1615801562000dd5576000805460ff60a01b1960ff60a81b19909116600160a81b1716600160a01b1790555b600280546001600160a01b0319166001600160a01b038516179055604051829062000e009062000efe565b62000e0c919062001157565b604051809103906000f08015801562000e29573d6000803e3d6000fd5b50600180546001600160a01b0319166001600160a01b0392909216919091179055801562000586576000805460ff60a81b19169055505050565b303b1590565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928262000ea1576000855562000eec565b82601f1062000ebc57805160ff191683800117855562000eec565b8280016001018555821562000eec579182015b8281111562000eec57825182559160200191906001019062000ecf565b5062000efa92915062000f0c565b5090565b610c2080620014c283390190565b5b8082111562000efa576000815560010162000f0d565b60006020828403121562000f35578081fd5b813562000f4281620014a8565b9392505050565b60006020828403121562000f5b578081fd5b815162000f4281620014a8565b6000806040838503121562000f7b578081fd5b823562000f8881620014a8565b9150602083013562000f9a81620014a8565b809150509250929050565b6000806040838503121562000fb8578182fd5b823562000fc581620014a8565b91506020838101356001600160401b0381111562000fe1578283fd5b8401601f8101861362000ff2578283fd5b80356200100962001003826200147e565b6200145a565b81815283810190838501858402850186018a101562001026578687fd5b8694505b83851015620010555780356200104081620014a8565b8352600194909401939185019185016200102a565b5080955050505050509250929050565b6000602080838503121562001078578182fd5b82516001600160401b038111156200108e578283fd5b8301601f810185136200109f578283fd5b8051620010b062001003826200147e565b8181528381019083850185840285018601891015620010cd578687fd5b8694505b83851015620010fc578051620010e781620014a8565b835260019490940193918501918501620010d1565b50979650505050505050565b6000602082840312156200111a578081fd5b8151801515811462000f42578182fd5b6000602082840312156200113c578081fd5b5051919050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b818110156200119a5783516001600160a01b03168352928401929184019160010162001173565b50909695505050505050565b901515815260200190565b6000602080835283518082850152825b81811015620011df57858101830151858201604001528201620011c1565b81811115620011f15783604083870101525b50601f01601f1916929092016040019392505050565b6000602080830181845282855460018082166000811462001231576001811462001250576200128c565b60028304607f16855260ff19831660408901526060880193506200128c565b60028304808652620012628a6200149c565b885b82811015620012825781548b82016040015290840190880162001264565b8a01604001955050505b5091979650505050505050565b60208082526061908201527f5b5145432d3030373030315d2d5468657265206973206e6f206578706572742060408201527f636f756e74206c696d69742c206661696c656420746f2067657420746865206c60608201527f696d6974206f662074686520636f756e74206f662074686520657870657274736080820152601760f91b60a082015260c00190565b60208082526052908201527f5b5145432d3030373030305d2d546865206c696d6974206f662074686520636f60408201527f756e74206f66206578706572747320697320726561636865642c206661696c6560608201527132103a379030903732bb9032bc3832b93a1760711b608082015260a00190565b60208082526022908201527f676f7665726e616e63652e636f6e737469747574696f6e2e706172616d657465604082015261727360f01b606082015260800190565b6020808252604b908201527f5b5145432d3030373030325d2d5468652063616c6c657220646f6573206e6f7460408201527f2068617665206163636573732c206f6e6c79207468652065787065727473206860608201526a30bb329030b1b1b2b9b99760a91b608082015260a00190565b90815260200190565b6040518181016001600160401b03811182821017156200147657fe5b604052919050565b60006001600160401b038211156200149257fe5b5060209081020190565b60009081526020902090565b6001600160a01b0381168114620014be57600080fd5b5056fe608060405234801561001057600080fd5b50604051610c20380380610c208339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b5050505090500160405250505060006100cf6101b560201b60201c565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060005b81518110156101ae57600282828151811061013357fe5b60209081029190910181015182546001808201855560009485529284200180546001600160a01b0319166001600160a01b03909216919091179055600254845190929085908590811061018257fe5b6020908102919091018101516001600160a01b031682528101919091526040016000205560010161011c565b50506101b9565b3390565b610a58806101c86000396000f3fe608060405234801561001057600080fd5b506004361061008e5760003560e01c80630a3b0a4f1461009357806329092d0e146100cd5780635dbe47e8146100f3578063715018a6146101195780638da5cb5b14610123578063949d225d146101475780639e9405c314610161578063a39fac1214610187578063c32d805a146101df578063f2fde38b14610205575b600080fd5b6100b9600480360360208110156100a957600080fd5b50356001600160a01b031661022b565b604080519115158252519081900360200190f35b6100b9600480360360208110156100e357600080fd5b50356001600160a01b03166102bd565b6100b96004803603602081101561010957600080fd5b50356001600160a01b031661035f565b61012161037c565b005b61012b61041e565b604080516001600160a01b039092168252519081900360200190f35b61014f61042d565b60408051918252519081900360200190f35b6101216004803603602081101561017757600080fd5b50356001600160a01b0316610433565b61018f6104ec565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101cb5781810151838201526020016101b3565b505050509050019250505060405180910390f35b610121600480360360208110156101f557600080fd5b50356001600160a01b031661054e565b6101216004803603602081101561021b57600080fd5b50356001600160a01b0316610649565b6000610235610741565b6000546001600160a01b03908116911614610285576040805162461bcd60e51b81526020600482018190526024820152600080516020610a03833981519152604482015290519081900360640190fd5b6001600160a01b038216600090815260016020526040902054156102ab575060006102b8565b6102b482610745565b5060015b919050565b60006102c7610741565b6000546001600160a01b03908116911614610317576040805162461bcd60e51b81526020600482018190526024820152600080516020610a03833981519152604482015290519081900360640190fd5b6001600160a01b03821660009081526001602052604090205480158061033e575060025481115b1561034d5760009150506102b8565b610356836107a6565b50600192915050565b6001600160a01b0316600090815260016020526040902054151590565b610384610741565b6000546001600160a01b039081169116146103d4576040805162461bcd60e51b81526020600482018190526024820152600080516020610a03833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b60025490565b61043b610741565b6000546001600160a01b0390811691161461048b576040805162461bcd60e51b81526020600482018190526024820152600080516020610a03833981519152604482015290519081900360640190fd5b6001600160a01b038116600090815260016020526040902054156104e05760405162461bcd60e51b815260040180806020018281038252606481526020018061094d6064913960800191505060405180910390fd5b6104e981610745565b50565b6060600280548060200260200160405190810160405280929190818152602001828054801561054457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610526575b5050505050905090565b610556610741565b6000546001600160a01b039081169116146105a6576040805162461bcd60e51b81526020600482018190526024820152600080516020610a03833981519152604482015290519081900360640190fd5b6001600160a01b038116600090815260016020526040902054806105fb5760405162461bcd60e51b81526004018080602001828103825260528152602001806109b16052913960600191505060405180910390fd5b60025481111561063c5760405162461bcd60e51b81526004018080602001828103825260638152602001806108ea6063913960800191505060405180910390fd5b610645826107a6565b5050565b610651610741565b6000546001600160a01b039081169116146106a1576040805162461bcd60e51b81526020600482018190526024820152600080516020610a03833981519152604482015290519081900360640190fd5b6001600160a01b0381166106e65760405162461bcd60e51b81526004018080602001828103825260268152602001806108c46026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b60028054600181810183557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace90910180546001600160a01b039094166001600160a01b03199094168417905590546000928352602091909152604090912055565b6001600160a01b03811660009081526001602081815260408084208151928301909152548082526002805492946000198401949293909190859081106107e857fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902055600280548290811061081b57fe5b6000918252602090912001548251600280546001600160a01b039093169290916000190190811061084857fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600280548061088157fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0394909416815260019093525050604081205556fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e76616c696420696e6465782076616c756520666f722074686520616464726573732073746f726167652c206661696c656420746f2072656d6f76652074686520616464726573732066726f6d2074686520616464726573732073746f726167652e54686520616464726573732068617320616c7265616479206265656e20616464656420746f207468652073746f726167652c206661696c656420746f2061646420746865206164647265737320746f2074686520616464726573732073746f726167652e546865206164647265737320646f6573206e6f742065786973742c206661696c656420746f2072656d6f76652074686520616464726573732066726f6d2074686520616464726573732073746f726167652e4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a264697066735822122083626d47adb3491349ee9a4368a52afe61478e893a6e65791394413c25ec40ea64736f6c634300070600334f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564676f7665726e616e63652e657870657274732e455044522e6d656d62657273686970566f74696e67a26469706673582212208ef6218023cf4de373274cd9c9cd0ed5f03b59e0e0d473a734f7aef3639f70d564736f6c63430007060033"

// DeployEPDRMembership deploys a new Ethereum contract, binding an instance of EPDRMembership to it.
func DeployEPDRMembership(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EPDRMembership, error) {
	parsed, err := abi.JSON(strings.NewReader(EPDRMembershipABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EPDRMembershipBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EPDRMembership{EPDRMembershipCaller: EPDRMembershipCaller{contract: contract}, EPDRMembershipTransactor: EPDRMembershipTransactor{contract: contract}, EPDRMembershipFilterer: EPDRMembershipFilterer{contract: contract}}, nil
}

// EPDRMembership is an auto generated Go binding around an Ethereum contract.
type EPDRMembership struct {
	EPDRMembershipCaller     // Read-only binding to the contract
	EPDRMembershipTransactor // Write-only binding to the contract
	EPDRMembershipFilterer   // Log filterer for contract events
}

// EPDRMembershipCaller is an auto generated read-only Go binding around an Ethereum contract.
type EPDRMembershipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPDRMembershipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EPDRMembershipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPDRMembershipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EPDRMembershipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPDRMembershipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EPDRMembershipSession struct {
	Contract     *EPDRMembership   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EPDRMembershipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EPDRMembershipCallerSession struct {
	Contract *EPDRMembershipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EPDRMembershipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EPDRMembershipTransactorSession struct {
	Contract     *EPDRMembershipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EPDRMembershipRaw is an auto generated low-level Go binding around an Ethereum contract.
type EPDRMembershipRaw struct {
	Contract *EPDRMembership // Generic contract binding to access the raw methods on
}

// EPDRMembershipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EPDRMembershipCallerRaw struct {
	Contract *EPDRMembershipCaller // Generic read-only contract binding to access the raw methods on
}

// EPDRMembershipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EPDRMembershipTransactorRaw struct {
	Contract *EPDRMembershipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEPDRMembership creates a new instance of EPDRMembership, bound to a specific deployed contract.
func NewEPDRMembership(address common.Address, backend bind.ContractBackend) (*EPDRMembership, error) {
	contract, err := bindEPDRMembership(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EPDRMembership{EPDRMembershipCaller: EPDRMembershipCaller{contract: contract}, EPDRMembershipTransactor: EPDRMembershipTransactor{contract: contract}, EPDRMembershipFilterer: EPDRMembershipFilterer{contract: contract}}, nil
}

// NewEPDRMembershipCaller creates a new read-only instance of EPDRMembership, bound to a specific deployed contract.
func NewEPDRMembershipCaller(address common.Address, caller bind.ContractCaller) (*EPDRMembershipCaller, error) {
	contract, err := bindEPDRMembership(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EPDRMembershipCaller{contract: contract}, nil
}

// NewEPDRMembershipTransactor creates a new write-only instance of EPDRMembership, bound to a specific deployed contract.
func NewEPDRMembershipTransactor(address common.Address, transactor bind.ContractTransactor) (*EPDRMembershipTransactor, error) {
	contract, err := bindEPDRMembership(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EPDRMembershipTransactor{contract: contract}, nil
}

// NewEPDRMembershipFilterer creates a new log filterer instance of EPDRMembership, bound to a specific deployed contract.
func NewEPDRMembershipFilterer(address common.Address, filterer bind.ContractFilterer) (*EPDRMembershipFilterer, error) {
	contract, err := bindEPDRMembership(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EPDRMembershipFilterer{contract: contract}, nil
}

// bindEPDRMembership binds a generic wrapper to an already deployed contract.
func bindEPDRMembership(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EPDRMembershipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EPDRMembership *EPDRMembershipRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EPDRMembership.Contract.EPDRMembershipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EPDRMembership *EPDRMembershipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPDRMembership.Contract.EPDRMembershipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EPDRMembership *EPDRMembershipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EPDRMembership.Contract.EPDRMembershipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EPDRMembership *EPDRMembershipCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EPDRMembership.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EPDRMembership *EPDRMembershipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPDRMembership.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EPDRMembership *EPDRMembershipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EPDRMembership.Contract.contract.Transact(opts, method, params...)
}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPDRMembership *EPDRMembershipCaller) ExpertLimitKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EPDRMembership.contract.Call(opts, out, "expertLimitKey")
	return *ret0, err
}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPDRMembership *EPDRMembershipSession) ExpertLimitKey() (string, error) {
	return _EPDRMembership.Contract.ExpertLimitKey(&_EPDRMembership.CallOpts)
}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPDRMembership *EPDRMembershipCallerSession) ExpertLimitKey() (string, error) {
	return _EPDRMembership.Contract.ExpertLimitKey(&_EPDRMembership.CallOpts)
}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPDRMembership *EPDRMembershipCaller) ExpertVotingKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EPDRMembership.contract.Call(opts, out, "expertVotingKey")
	return *ret0, err
}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPDRMembership *EPDRMembershipSession) ExpertVotingKey() (string, error) {
	return _EPDRMembership.Contract.ExpertVotingKey(&_EPDRMembership.CallOpts)
}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPDRMembership *EPDRMembershipCallerSession) ExpertVotingKey() (string, error) {
	return _EPDRMembership.Contract.ExpertVotingKey(&_EPDRMembership.CallOpts)
}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPDRMembership *EPDRMembershipCaller) Experts(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EPDRMembership.contract.Call(opts, out, "experts")
	return *ret0, err
}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPDRMembership *EPDRMembershipSession) Experts() (common.Address, error) {
	return _EPDRMembership.Contract.Experts(&_EPDRMembership.CallOpts)
}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPDRMembership *EPDRMembershipCallerSession) Experts() (common.Address, error) {
	return _EPDRMembership.Contract.Experts(&_EPDRMembership.CallOpts)
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPDRMembership *EPDRMembershipCaller) GetLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EPDRMembership.contract.Call(opts, out, "getLimit")
	return *ret0, err
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPDRMembership *EPDRMembershipSession) GetLimit() (*big.Int, error) {
	return _EPDRMembership.Contract.GetLimit(&_EPDRMembership.CallOpts)
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPDRMembership *EPDRMembershipCallerSession) GetLimit() (*big.Int, error) {
	return _EPDRMembership.Contract.GetLimit(&_EPDRMembership.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPDRMembership *EPDRMembershipCaller) GetMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _EPDRMembership.contract.Call(opts, out, "getMembers")
	return *ret0, err
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPDRMembership *EPDRMembershipSession) GetMembers() ([]common.Address, error) {
	return _EPDRMembership.Contract.GetMembers(&_EPDRMembership.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPDRMembership *EPDRMembershipCallerSession) GetMembers() ([]common.Address, error) {
	return _EPDRMembership.Contract.GetMembers(&_EPDRMembership.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPDRMembership *EPDRMembershipCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EPDRMembership.contract.Call(opts, out, "getSize")
	return *ret0, err
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPDRMembership *EPDRMembershipSession) GetSize() (*big.Int, error) {
	return _EPDRMembership.Contract.GetSize(&_EPDRMembership.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPDRMembership *EPDRMembershipCallerSession) GetSize() (*big.Int, error) {
	return _EPDRMembership.Contract.GetSize(&_EPDRMembership.CallOpts)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPDRMembership *EPDRMembershipCaller) IsMember(opts *bind.CallOpts, _exp common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EPDRMembership.contract.Call(opts, out, "isMember", _exp)
	return *ret0, err
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPDRMembership *EPDRMembershipSession) IsMember(_exp common.Address) (bool, error) {
	return _EPDRMembership.Contract.IsMember(&_EPDRMembership.CallOpts, _exp)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPDRMembership *EPDRMembershipCallerSession) IsMember(_exp common.Address) (bool, error) {
	return _EPDRMembership.Contract.IsMember(&_EPDRMembership.CallOpts, _exp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPDRMembership *EPDRMembershipCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EPDRMembership.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPDRMembership *EPDRMembershipSession) Owner() (common.Address, error) {
	return _EPDRMembership.Contract.Owner(&_EPDRMembership.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPDRMembership *EPDRMembershipCallerSession) Owner() (common.Address, error) {
	return _EPDRMembership.Contract.Owner(&_EPDRMembership.CallOpts)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipTransactor) AddMember(opts *bind.TransactOpts, _exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "addMember", _exp)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipSession) AddMember(_exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.AddMember(&_EPDRMembership.TransactOpts, _exp)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) AddMember(_exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.AddMember(&_EPDRMembership.TransactOpts, _exp)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPDRMembership *EPDRMembershipTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "initialize", _registry, _experts)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPDRMembership *EPDRMembershipSession) Initialize(_registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.Initialize(&_EPDRMembership.TransactOpts, _registry, _experts)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) Initialize(_registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.Initialize(&_EPDRMembership.TransactOpts, _registry, _experts)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipTransactor) RemoveMember(opts *bind.TransactOpts, _exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "removeMember", _exp)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipSession) RemoveMember(_exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.RemoveMember(&_EPDRMembership.TransactOpts, _exp)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) RemoveMember(_exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.RemoveMember(&_EPDRMembership.TransactOpts, _exp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPDRMembership *EPDRMembershipTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPDRMembership *EPDRMembershipSession) RenounceOwnership() (*types.Transaction, error) {
	return _EPDRMembership.Contract.RenounceOwnership(&_EPDRMembership.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EPDRMembership.Contract.RenounceOwnership(&_EPDRMembership.TransactOpts)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPDRMembership *EPDRMembershipTransactor) SwapMember(opts *bind.TransactOpts, _oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "swapMember", _oldExp, _newExp)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPDRMembership *EPDRMembershipSession) SwapMember(_oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.SwapMember(&_EPDRMembership.TransactOpts, _oldExp, _newExp)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) SwapMember(_oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.SwapMember(&_EPDRMembership.TransactOpts, _oldExp, _newExp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPDRMembership *EPDRMembershipTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPDRMembership *EPDRMembershipSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.TransferOwnership(&_EPDRMembership.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.TransferOwnership(&_EPDRMembership.TransactOpts, newOwner)
}

// EPDRMembershipOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EPDRMembership contract.
type EPDRMembershipOwnershipTransferredIterator struct {
	Event *EPDRMembershipOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EPDRMembershipOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EPDRMembershipOwnershipTransferred)
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
		it.Event = new(EPDRMembershipOwnershipTransferred)
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
func (it *EPDRMembershipOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EPDRMembershipOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EPDRMembershipOwnershipTransferred represents a OwnershipTransferred event raised by the EPDRMembership contract.
type EPDRMembershipOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EPDRMembership *EPDRMembershipFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EPDRMembershipOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EPDRMembership.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EPDRMembershipOwnershipTransferredIterator{contract: _EPDRMembership.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EPDRMembership *EPDRMembershipFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EPDRMembershipOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EPDRMembership.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EPDRMembershipOwnershipTransferred)
				if err := _EPDRMembership.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EPDRMembership *EPDRMembershipFilterer) ParseOwnershipTransferred(log types.Log) (*EPDRMembershipOwnershipTransferred, error) {
	event := new(EPDRMembershipOwnershipTransferred)
	if err := _EPDRMembership.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
