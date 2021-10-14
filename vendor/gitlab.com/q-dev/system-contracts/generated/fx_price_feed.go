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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pair\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_decimal\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_maintainersList\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_baseTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quoteTokenAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimalPlaces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futurePricingTolerance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pair\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maintainer\",\"type\":\"address\"}],\"name\":\"setMaintainer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leaveMaintainers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exchangeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricingTime\",\"type\":\"uint256\"}],\"name\":\"setExchangeRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tolerance\",\"type\":\"uint256\"}],\"name\":\"setFuturePricingTolerance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaintainers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620023e1380380620023e1833981810160405260a08110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b506040818152602083015192018051929491939192846401000000008211156200010f57600080fd5b9083019060208201858111156200012557600080fd5b82518660208202830111640100000000821117156200014357600080fd5b82525081516020918201928201910280838360005b838110156200017257818101518382015260200162000158565b505050509190910160409081526020830151920151919350909150600090506200019b620004ab565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3508451620002245760405162461bcd60e51b8152600401808060200182810382526035815260200180620023746035913960400191505060405180910390fd5b60008411620002655760405162461bcd60e51b8152600401808060200182810382526038815260200180620023a96038913960400191505060405180910390fd5b6001600160a01b038216620002ac5760405162461bcd60e51b81526004018080602001828103825260438152602001806200228e6043913960600191505060405180910390fd5b6001600160a01b038116620002f35760405162461bcd60e51b81526004018080602001828103825260448152602001806200224a6044913960600191505060405180910390fd5b620002fe82620004af565b6200033b5760405162461bcd60e51b8152600401808060200182810382526051815260200180620023236051913960600191505060405180910390fd5b6200034681620004af565b620003835760405162461bcd60e51b8152600401808060200182810382526052815260200180620022d16052913960600191505060405180910390fd5b6040516200039190620004bb565b604051809103906000f080158015620003ae573d6000803e3d6000fd5b50600180546001600160a01b0319166001600160a01b03928316179081905560405163a224cee760e01b8152602060048201818152875160248401528751939094169363a224cee79388939192839260440191858201910280838360005b83811015620004265781810151838201526020016200040c565b5050505090500192505050600060405180830381600087803b1580156200044c57600080fd5b505af115801562000461573d6000803e3d6000fd5b505086516200047a9250600291506020880190620004c9565b5060c0939093526001600160601b0319606091821b811660805292901b90911660a052505061070860065562000575565b3390565b3b63ffffffff16151590565b610f6480620012e683390190565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826200050157600085556200054c565b82601f106200051c57805160ff19168380011785556200054c565b828001600101855582156200054c579182015b828111156200054c5782518255916020019190600101906200052f565b506200055a9291506200055e565b5090565b5b808211156200055a57600081556001016200055f565b60805160601c60a05160601c60c051610d3e620005a8600039806109065250806105685250806105a75250610d3e6000f3fe608060405234801561001057600080fd5b50600436106100e05760003560e01c8063a094600e11610087578063a094600e14610188578063a8aa1b3114610190578063d17f42251461020d578063d520403314610265578063dc5550901461026d578063e1725c9214610275578063f2fde38b1461027d578063f55ecf06146102a3576100e0565b806313ea5d29146100e557806337c5f2bb1461010d5780633ba0b9a91461012a578063715018a6146101445780637acede8b1461014c5780637f1ba22814610170578063871ebd47146101785780638da5cb5b14610180575b600080fd5b61010b600480360360208110156100fb57600080fd5b50356001600160a01b03166102da565b005b61010b6004803603602081101561012357600080fd5b50356103f5565b6101326104ac565b60408051918252519081900360200190f35b61010b6104b2565b610154610566565b604080516001600160a01b039092168252519081900360200190f35b61013261058a565b610132610590565b610154610596565b6101546105a5565b6101986105c9565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101d25781810151838201526020016101ba565b50505050905090810190601f1680156101ff5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610215610654565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610251578181015183820152602001610239565b505050509050019250505060405180910390f35b61010b610765565b6101326108fe565b610132610904565b61010b6004803603602081101561029357600080fd5b50356001600160a01b0316610928565b6102c6600480360360408110156102b957600080fd5b5080359060200135610a32565b604080519115158252519081900360200190f35b60015460408051630bb7c8fd60e31b815233600482015290516001600160a01b0390921691635dbe47e891602480820192602092909190829003018186803b15801561032557600080fd5b505afa158015610339573d6000803e3d6000fd5b505050506040513d602081101561034f57600080fd5b505161038c5760405162461bcd60e51b815260040180806020018281038252603e815260200180610c63603e913960400191505060405180910390fd5b60015460408051639e9405c360e01b81526001600160a01b03848116600483015291519190921691639e9405c391602480830192600092919082900301818387803b1580156103da57600080fd5b505af11580156103ee573d6000803e3d6000fd5b5050505050565b60015460408051630bb7c8fd60e31b815233600482015290516001600160a01b0390921691635dbe47e891602480820192602092909190829003018186803b15801561044057600080fd5b505afa158015610454573d6000803e3d6000fd5b505050506040513d602081101561046a57600080fd5b50516104a75760405162461bcd60e51b815260040180806020018281038252603e815260200180610c63603e913960400191505060405180910390fd5b600655565b60055481565b6104ba610bb9565b6000546001600160a01b0390811691161461051c576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b7f000000000000000000000000000000000000000000000000000000000000000081565b60065481565b60045481565b6000546001600160a01b031690565b7f000000000000000000000000000000000000000000000000000000000000000081565b6002805460408051602060018416156101000260001901909316849004601f8101849004840282018401909252818152929183018282801561064c5780601f106106215761010080835404028352916020019161064c565b820191906000526020600020905b81548152906001019060200180831161062f57829003601f168201915b505050505081565b600154604080516351cfd60960e11b815290516060926001600160a01b03169163a39fac12916004808301926000929190829003018186803b15801561069957600080fd5b505afa1580156106ad573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156106d657600080fd5b8101908080516040519392919084600160201b8211156106f557600080fd5b90830190602082018581111561070a57600080fd5b82518660208202830111600160201b8211171561072657600080fd5b82525081516020918201928201910280838360005b8381101561075357818101518382015260200161073b565b50505050905001604052505050905090565b60015460408051630bb7c8fd60e31b815233600482015290516001600160a01b0390921691635dbe47e891602480820192602092909190829003018186803b1580156107b057600080fd5b505afa1580156107c4573d6000803e3d6000fd5b505050506040513d60208110156107da57600080fd5b5051156108985760018060009054906101000a90046001600160a01b03166001600160a01b031663949d225d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561083057600080fd5b505afa158015610844573d6000803e3d6000fd5b505050506040513d602081101561085a57600080fd5b5051116108985760405162461bcd60e51b815260040180806020018281038252604c815260200180610be4604c913960600191505060405180910390fd5b60015460408051636196c02d60e11b815233600482015290516001600160a01b039092169163c32d805a9160248082019260009290919082900301818387803b1580156108e457600080fd5b505af11580156108f8573d6000803e3d6000fd5b50505050565b60035481565b7f000000000000000000000000000000000000000000000000000000000000000081565b610930610bb9565b6000546001600160a01b03908116911614610992576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b0381166109d75760405162461bcd60e51b8152600401808060200182810382526026815260200180610bbe6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60015460408051630bb7c8fd60e31b815233600482015290516000926001600160a01b031691635dbe47e8916024808301926020929190829003018186803b158015610a7d57600080fd5b505afa158015610a91573d6000803e3d6000fd5b505050506040513d6020811015610aa757600080fd5b5051610ae45760405162461bcd60e51b815260040180806020018281038252603e815260200180610c63603e913960400191505060405180910390fd5b60008311610b235760405162461bcd60e51b8152600401808060200182810382526044815260200180610cc56044913960600191505060405180910390fd5b6004548211610b635760405162461bcd60e51b8152600401808060200182810382526033815260200180610c306033913960400191505060405180910390fd5b6006544201821115610ba65760405162461bcd60e51b8152600401808060200182810382526024815260200180610ca16024913960400191505060405180910390fd5b5060059190915542600355600455600190565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573735b5145432d3031393030385d2d43616e6e6f74206c6561766520746865206d61696e7461696e657273206c6973742c20796f752061726520746865206f6e6c79206d61696e7461696e65722e5b5145432d3031393030395d2d50726963696e672074696d65206973206c657373206f7220657175616c2063757272656e742e5b5145432d3031393030305d2d5065726d697373696f6e2064656e696564202d206f6e6c79206d61696e7461696e6572732068617665206163636573732e5b5145432d3031393031305d2d50726963696e672074696d6520696e206675747572652e5b5145432d3031393030375d2d496e76616c69642065786368616e676520726174652c206661696c656420746f20736574207468652065786368616e676520726174652ea2646970667358221220f4eb46a1f26b4a732b7920e499f063bd792f57dfa452ab6f30b0a28a676a4d7764736f6c63430007060033608060405234801561001057600080fd5b50610f44806100206000396000f3fe608060405234801561001057600080fd5b50600436106100af5760003560e01c80630a3b0a4f146100b457806329092d0e146100ee57806341c0e1b51461011457806352efea6e1461011e5780635dbe47e814610126578063715018a61461014c5780638da5cb5b14610154578063949d225d146101785780639e9405c314610192578063a224cee7146101b8578063a39fac1214610259578063c32d805a146102b1578063f2fde38b146102d7575b600080fd5b6100da600480360360208110156100ca57600080fd5b50356001600160a01b03166102fd565b604080519115158252519081900360200190f35b6100da6004803603602081101561010457600080fd5b50356001600160a01b0316610399565b61011c610445565b005b6100da6104aa565b6100da6004803603602081101561013c57600080fd5b50356001600160a01b031661056c565b61011c610589565b61015c610623565b604080516001600160a01b039092168252519081900360200190f35b610180610632565b60408051918252519081900360200190f35b61011c600480360360208110156101a857600080fd5b50356001600160a01b0316610638565b61011c600480360360208110156101ce57600080fd5b810190602081018135600160201b8111156101e857600080fd5b8201836020820111156101fa57600080fd5b803590602001918460208302840111600160201b8311171561021b57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061067f945050505050565b6102616107c3565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561029d578181015183820152602001610285565b505050509050019250505060405180910390f35b61011c600480360360208110156102c757600080fd5b50356001600160a01b0316610825565b61011c600480360360208110156102ed57600080fd5b50356001600160a01b0316610869565b600061030761095a565b6001600160a01b0316610318610623565b6001600160a01b031614610361576040805162461bcd60e51b81526020600482018190526024820152600080516020610ecf833981519152604482015290519081900360640190fd5b6001600160a01b0382166000908152606560205260409020541561038757506000610394565b6103908261095e565b5060015b919050565b60006103a361095a565b6001600160a01b03166103b4610623565b6001600160a01b0316146103fd576040805162461bcd60e51b81526020600482018190526024820152600080516020610ecf833981519152604482015290519081900360640190fd5b6001600160a01b038216600090815260656020526040902054801580610424575060665481115b15610433576000915050610394565b61043c836109c1565b50600192915050565b61044d61095a565b6001600160a01b031661045e610623565b6001600160a01b0316146104a7576040805162461bcd60e51b81526020600482018190526024820152600080516020610ecf833981519152604482015290519081900360640190fd5b33ff5b60006104b461095a565b6001600160a01b03166104c5610623565b6001600160a01b03161461050e576040805162461bcd60e51b81526020600482018190526024820152600080516020610ecf833981519152604482015290519081900360640190fd5b60005b60665481101561055957606560006066838154811061052c57fe5b60009182526020808320909101546001600160a01b03168352820192909252604001812055600101610511565b5061056660666000610d94565b50600190565b6001600160a01b0316600090815260656020526040902054151590565b61059161095a565b6001600160a01b03166105a2610623565b6001600160a01b0316146105eb576040805162461bcd60e51b81526020600482018190526024820152600080516020610ecf833981519152604482015290519081900360640190fd5b6033546040516000916001600160a01b031690600080516020610eef833981519152908390a3603380546001600160a01b0319169055565b6033546001600160a01b031690565b60665490565b610641816102fd565b61067c5760405162461bcd60e51b8152600401808060200182810382526071815260200180610e306071913960800191505060405180910390fd5b50565b600054610100900460ff16806106985750610698610afb565b806106a6575060005460ff16155b6106e15760405162461bcd60e51b815260040180806020018281038252602e815260200180610ea1602e913960400191505060405180910390fd5b600054610100900460ff1615801561070c576000805460ff1961ff0019909116610100171660011790555b60005b82518110156107a457606683828151811061072657fe5b6020908102919091018101518254600181018455600093845291832090910180546001600160a01b0319166001600160a01b039092169190911790556066548451909160659186908590811061077857fe5b6020908102919091018101516001600160a01b031682528101919091526040016000205560010161070f565b506107ad610b0c565b80156107bf576000805461ff00191690555b5050565b6060606680548060200260200160405190810160405280929190818152602001828054801561081b57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116107fd575b5050505050905090565b61082e81610399565b61067c5760405162461bcd60e51b8152600401808060200182810382526043815260200180610ded6043913960600191505060405180910390fd5b61087161095a565b6001600160a01b0316610882610623565b6001600160a01b0316146108cb576040805162461bcd60e51b81526020600482018190526024820152600080516020610ecf833981519152604482015290519081900360640190fd5b6001600160a01b0381166109105760405162461bcd60e51b8152600401808060200182810382526026815260200180610dc76026913960400191505060405180910390fd5b6033546040516001600160a01b03808416921690600080516020610eef83398151915290600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b606680546001810182557f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943540180546001600160a01b0319166001600160a01b03841690811790915590546000918252606560205260409091205561067c81610bbd565b6001600160a01b0381166000908152606560209081526040808320815192830190915254815260668054919260001983019290919083908110610a0057fe5b60009182526020909120015483516001600160a01b039091169150600019018214610a845782516001600160a01b038216600090815260656020526040902055825160668054839260001901908110610a5557fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b6066805480610a8f57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0386168252606590526040812055610ad384610bbd565b836001600160a01b0316816001600160a01b031614610af557610af581610bbd565b50505050565b6000610b0630610c07565b15905090565b600054610100900460ff1680610b255750610b25610afb565b80610b33575060005460ff16155b610b6e5760405162461bcd60e51b815260040180806020018281038252602e815260200180610ea1602e913960400191505060405180910390fd5b600054610100900460ff16158015610b99576000805460ff1961ff0019909116610100171660011790555b610ba1610c0d565b610ba9610cad565b801561067c576000805461ff001916905550565b6001600160a01b038116600090815260656020526040902054606654811115610be257fe5b610beb8261056c565b15610bff5760008111610bfa57fe5b6107bf565b80156107bf57fe5b3b151590565b600054610100900460ff1680610c265750610c26610afb565b80610c34575060005460ff16155b610c6f5760405162461bcd60e51b815260040180806020018281038252602e815260200180610ea1602e913960400191505060405180910390fd5b600054610100900460ff16158015610ba9576000805460ff1961ff001990911661010017166001179055801561067c576000805461ff001916905550565b600054610100900460ff1680610cc65750610cc6610afb565b80610cd4575060005460ff16155b610d0f5760405162461bcd60e51b815260040180806020018281038252602e815260200180610ea1602e913960400191505060405180910390fd5b600054610100900460ff16158015610d3a576000805460ff1961ff0019909116610100171660011790555b6000610d4461095a565b603380546001600160a01b0319166001600160a01b03831690811790915560405191925090600090600080516020610eef833981519152908290a350801561067c576000805461ff001916905550565b508054600082559060005260206000209081019061067c91905b80821115610dc25760008155600101610dae565b509056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573735b5145432d3033353030305d2d4661696c656420746f2072656d6f76652074686520616464726573732066726f6d2074686520616464726573732073746f726167652e5b5145432d3033353030325d2d54686520616464726573732068617320616c7265616479206265656e20616464656420746f207468652073746f726167652c206661696c656420746f2061646420746865206164647265737320746f2074686520616464726573732073746f726167652e496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a65644f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a2646970667358221220240f308dc1c73d3733a6dd3668ac2504982e6d0d7b04d5fdbdcdedb7de9451a864736f6c634300070600335b5145432d3031393030345d2d496e76616c69642071756f746520746f6b656e20616464726573732c2074686520696e697469616c697a6174696f6e206661696c65642e5b5145432d3031393030335d2d496e76616c6964206261736520746f6b656e20616464726573732c2074686520696e697469616c697a6174696f6e206661696c65642e5b5145432d3031393030365d2d5468652071756f746520746f6b656e2061646472657373206973206e6f74206120636f6e74726163742c2074686520696e697469616c697a6174696f6e206661696c65642e5b5145432d3031393030355d2d546865206261736520746f6b656e2061646472657373206973206e6f74206120636f6e74726163742c2074686520696e697469616c697a6174696f6e206661696c65642e5b5145432d3031393030315d2d496e76616c696420706169722c2074686520696e697469616c697a6174696f6e206661696c65642e5b5145432d3031393030325d2d496e76616c696420646563696d616c2c2074686520696e697469616c697a6174696f6e206661696c65642e",
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
// Solidity: function setExchangeRate(uint256 _exchangeRate, uint256 _pricingTime) returns(bool)
func (_FxPriceFeed *FxPriceFeedTransactor) SetExchangeRate(opts *bind.TransactOpts, _exchangeRate *big.Int, _pricingTime *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.contract.Transact(opts, "setExchangeRate", _exchangeRate, _pricingTime)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xf55ecf06.
//
// Solidity: function setExchangeRate(uint256 _exchangeRate, uint256 _pricingTime) returns(bool)
func (_FxPriceFeed *FxPriceFeedSession) SetExchangeRate(_exchangeRate *big.Int, _pricingTime *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetExchangeRate(&_FxPriceFeed.TransactOpts, _exchangeRate, _pricingTime)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xf55ecf06.
//
// Solidity: function setExchangeRate(uint256 _exchangeRate, uint256 _pricingTime) returns(bool)
func (_FxPriceFeed *FxPriceFeedTransactorSession) SetExchangeRate(_exchangeRate *big.Int, _pricingTime *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetExchangeRate(&_FxPriceFeed.TransactOpts, _exchangeRate, _pricingTime)
}

// SetFuturePricingTolerance is a paid mutator transaction binding the contract method 0x37c5f2bb.
//
// Solidity: function setFuturePricingTolerance(uint256 _tolerance) returns()
func (_FxPriceFeed *FxPriceFeedTransactor) SetFuturePricingTolerance(opts *bind.TransactOpts, _tolerance *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.contract.Transact(opts, "setFuturePricingTolerance", _tolerance)
}

// SetFuturePricingTolerance is a paid mutator transaction binding the contract method 0x37c5f2bb.
//
// Solidity: function setFuturePricingTolerance(uint256 _tolerance) returns()
func (_FxPriceFeed *FxPriceFeedSession) SetFuturePricingTolerance(_tolerance *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetFuturePricingTolerance(&_FxPriceFeed.TransactOpts, _tolerance)
}

// SetFuturePricingTolerance is a paid mutator transaction binding the contract method 0x37c5f2bb.
//
// Solidity: function setFuturePricingTolerance(uint256 _tolerance) returns()
func (_FxPriceFeed *FxPriceFeedTransactorSession) SetFuturePricingTolerance(_tolerance *big.Int) (*types.Transaction, error) {
	return _FxPriceFeed.Contract.SetFuturePricingTolerance(&_FxPriceFeed.TransactOpts, _tolerance)
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
