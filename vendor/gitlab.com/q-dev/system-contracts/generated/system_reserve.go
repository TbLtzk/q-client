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

// SystemReserveABI is the input ABI used to generate the binding from.
const SystemReserveABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"availableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coolDownPhase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_keys\",\"type\":\"string[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setPauseState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEligibleContractKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SystemReserveBin is the compiled bytecode used for deploying new contracts.
var SystemReserveBin = "0x608060405234801561001057600080fd5b506121e1806100206000396000f3fe6080604052600436106200006d5760003560e01c806328011170146200007a5780632d250a5b14620000aa5780632e1a7d4d14620000d157806391f7cfb9146200010557806394e6f56f146200011d5780639d2f83f01462000144578063cdb88ad1146200015c5762000075565b366200007557005b600080fd5b3480156200008757600080fd5b506200009262000181565b604051620000a19190620010b6565b60405180910390f35b348015620000b757600080fd5b50620000c262000187565b604051620000a1919062000ce9565b348015620000de57600080fd5b50620000f6620000f036600462000c72565b62000211565b604051620000a1919062000d4d565b3480156200011257600080fd5b5062000092620004ac565b3480156200012a57600080fd5b50620001426200013c36600462000a7a565b620004b2565b005b3480156200015157600080fd5b50620000f6620005f6565b3480156200016957600080fd5b50620001426200017b36600462000c34565b62000606565b60025481565b600154604080516308d9a5ad60e41b815290516060926001600160a01b031691638d9a5ad0916004808301926000929190829003018186803b158015620001cd57600080fd5b505afa158015620001e2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526200020c919081019062000b63565b905090565b60008054600154604080516308d9a5ad60e41b815290516001600160a01b036201000090940484169385931691638d9a5ad09160048083019286929190829003018186803b1580156200026357600080fd5b505afa15801562000278573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052620002a2919081019062000b63565b90506000805b82518110156200037557336001600160a01b0316846001600160a01b0316633fb90271858481518110620002d857fe5b60200260200101516040518263ffffffff1660e01b8152600401620002fe919062000d58565b60206040518083038186803b1580156200031757600080fd5b505afa1580156200032c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000352919062000a54565b6001600160a01b031614156200036c576001915062000375565b600101620002a8565b50806200039f5760405162461bcd60e51b8152600401620003969062000f41565b60405180910390fd5b600154600160a01b900460ff1615620003cc5760405162461bcd60e51b8152600401620003969062000ec8565b620003d662000858565b84471080620003e3575084155b15620003f35760009350620004a4565b846003541015620004185760405162461bcd60e51b8152600401620003969062000e6a565b60038054869003905560405160009033908790620004369062000cd2565b60006040518083038185875af1925050503d806000811462000475576040519150601f19603f3d011682016040523d82523d6000602084013e6200047a565b606091505b50509050806200049e5760405162461bcd60e51b8152600401620003969062000fd1565b60019450505b505050919050565b60035481565b600054610100900460ff1680620004ce5750620004ce62000882565b80620004dd575060005460ff16155b6200051a5760405162461bcd60e51b815260040180806020018281038252602e8152602001806200217e602e913960400191505060405180910390fd5b600054610100900460ff1615801562000546576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038616021790556040518290620005789062000a46565b62000584919062000ce9565b604051809103906000f080158015620005a1573d6000803e3d6000fd5b50600180546001600160a01b0319166001600160a01b0392909216919091179055620005cc62000888565b4201600255620005db62000991565b6003558015620005f1576000805461ff00191690555b505050565b600154600160a01b900460ff1681565b600054604051633fb9027160e01b8152620100009091046001600160a01b0316908190633fb90271906200063d9060040162000e27565b60206040518083038186803b1580156200065657600080fd5b505afa1580156200066b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000691919062000a54565b6001600160a01b031663a230c524336040518263ffffffff1660e01b8152600401620006be919062000cd5565b60206040518083038186803b158015620006d757600080fd5b505afa158015620006ec573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000712919062000c53565b806200081a5750604051633fb9027160e01b81526001600160a01b03821690633fb9027190620007459060040162001041565b60206040518083038186803b1580156200075e57600080fd5b505afa15801562000773573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000799919062000a54565b6001600160a01b031663a230c524336040518263ffffffff1660e01b8152600401620007c6919062000cd5565b60206040518083038186803b158015620007df57600080fd5b505afa158015620007f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200081a919062000c53565b620008395760405162461bcd60e51b8152600401620003969062000d6d565b5060018054911515600160a01b0260ff60a01b19909216919091179055565b42600254101562000880576200086d62000888565b42016002556200087c62000991565b6003555b565b303b1590565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190620008be9060040162000de4565b60206040518083038186803b158015620008d757600080fd5b505afa158015620008ec573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000912919062000a54565b6001600160a01b031663498bff006040518163ffffffff1660e01b81526004016200093d9062000f0a565b60206040518083038186803b1580156200095657600080fd5b505afa1580156200096b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200020c919062000c8b565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190620009c79060040162000de4565b60206040518083038186803b158015620009e057600080fd5b505afa158015620009f5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000a1b919062000a54565b6001600160a01b031663498bff006040518163ffffffff1660e01b81526004016200093d906200106f565b610fff806200117f83390190565b60006020828403121562000a66578081fd5b815162000a738162001156565b9392505050565b600080604080848603121562000a8e578182fd5b833562000a9b8162001156565b92506020848101356001600160401b0381111562000ab7578384fd5b8501601f8101871362000ac8578384fd5b803562000adf62000ad982620010e3565b620010bf565b81815283810190838501875b8481101562000b5257813586018c603f82011262000b0757898afd5b8781013562000b1a62000ad98262001101565b8181528e8b83850101111562000b2e578b8cfd5b818b84018b83013790810189018b9052855250928601929086019060010162000aeb565b50979a909950975050505050505050565b6000602080838503121562000b76578182fd5b82516001600160401b0381111562000b8c578283fd5b8301601f8101851362000b9d578283fd5b805162000bae62000ad982620010e3565b81815283810190838501865b8481101562000c2657815186018a603f82011262000bd6578889fd5b87810151604062000beb62000ad98362001101565b8281528d8284860101111562000bff578b8cfd5b62000c10838c830184870162001123565b8752505050928601929086019060010162000bba565b509098975050505050505050565b60006020828403121562000c46578081fd5b813562000a73816200116f565b60006020828403121562000c65578081fd5b815162000a73816200116f565b60006020828403121562000c84578081fd5b5035919050565b60006020828403121562000c9d578081fd5b5051919050565b6000815180845262000cbe81602086016020860162001123565b601f01601f19169290920160200192915050565b90565b6001600160a01b0391909116815260200190565b6000602080830181845280855180835260408601915060408482028701019250838701855b8281101562000d4057603f1988860301845262000d2d85835162000ca4565b9450928501929085019060010162000d0e565b5092979650505050505050565b901515815260200190565b60006020825262000a73602083018462000ca4565b60208082526051908201527f5b5145432d3031383030315d2d5468652063616c6c657220646f6573206e6f7460408201527f2068617665206163636573732c206f6e6c79207468652065707166692065787060608201527032b93a39903430bb329030b1b1b2b9b99760791b608082015260a00190565b60208082526023908201527f676f7665726e616e63652e657870657274732e45505146492e706172616d657460408201526265727360e81b606082015260800190565b60208082526023908201527f676f7665726e616e63652e657870657274732e45505146492e6d656d6265727360408201526206869760ec1b606082015260800190565b602080825260409082018190527f5b5145432d3031383030335d2d496e73756666696369656e742066756e647320908201527f746f2077697468647261772c206661696c656420746f2077697468647261772e606082015260800190565b60208082526022908201527f5b5145432d3031383030325d2d5468652073797374656d206973207061757365604082015261321760f11b606082015260800190565b6020808252601f908201527f676f7665726e65642e45505146492e72657365727665436f6f6c446f776e5000604082015260600190565b60208082526064908201527f5b5145432d3031383030305d2d5468652063616c6c657220646f6573206e6f7460408201527f206861766520656e6f756768207374616b6520746f20636f766572207468652060608201527f64656c65676174696f6e732c207374616b652064656c65676174696f6e2066616080820152631a5b195960e21b60a082015260c00190565b6020808252604a908201527f5b5145432d3031383030345d2d4661696c656420746f207472616e736665722060408201527f746865207769746864726177616c20616d6f756e742c206661696c656420746f606082015269103bb4ba34323930bb9760b11b608082015260a00190565b602080825260149082015273676f7665726e616e63652e726f6f744e6f64657360601b604082015260600190565b60208082526027908201527f676f7665726e65642e45505146492e72657365727665436f6f6c446f776e54686040820152661c995cda1bdb1960ca1b606082015260800190565b90815260200190565b6040518181016001600160401b0381118282101715620010db57fe5b604052919050565b60006001600160401b03821115620010f757fe5b5060209081020190565b60006001600160401b038211156200111557fe5b50601f01601f191660200190565b60005b838110156200114057818101518382015260200162001126565b8381111562001150576000848401525b50505050565b6001600160a01b03811681146200116c57600080fd5b50565b80151581146200116c57600080fdfe60806040523480156200001157600080fd5b5060405162000fff38038062000fff8339810160408190526200003491620001d8565b60006200004062000128565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060005b815181101562000120576002828281518110620000a657fe5b60209081029190910181015182546001810184556000938452928290208151620000da94919091019291909101906200012c565b506002805490506001838381518110620000f057fe5b6020026020010151604051620001079190620002ca565b908152604051908190036020019020556001016200008d565b50506200033f565b3390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620001645760008555620001af565b82601f106200017f57805160ff1916838001178555620001af565b82800160010185558215620001af579182015b82811115620001af57825182559160200191906001019062000192565b50620001bd929150620001c1565b5090565b5b80821115620001bd5760008155600101620001c2565b60006020808385031215620001eb578182fd5b82516001600160401b038082111562000202578384fd5b8185019150601f868184011262000217578485fd5b8251828111156200022457fe5b620002338586830201620002e8565b81815285810190858701885b84811015620002ba57815188018c603f8201126200025b578a8bfd5b898101516040898211156200026c57fe5b6200027f828a01601f19168d01620002e8565b8281528f8284860101111562000293578d8efd5b620002a4838e83018487016200030c565b875250505092880192908801906001016200023f565b50909a9950505050505050505050565b60008251620002de8184602087016200030c565b9190910192915050565b6040518181016001600160401b03811182821017156200030457fe5b604052919050565b60005b83811015620003295781810151838201526020016200030f565b8381111562000339576000848401525b50505050565b610cb0806200034f6000396000f3fe608060405234801561001057600080fd5b50600436106100995760003560e01c8063374155e51461009e5780636833d54f146100b3578063715018a6146100dc57806380599e4b146100e45780638d9a5ad0146100f75780638da5cb5b1461010c578063949d225d146101215780639b1574eb14610136578063b0c8f9dc14610149578063ea1bbe351461015c578063f2fde38b1461016f575b600080fd5b6100b16100ac3660046108b5565b610182565b005b6100c66100c13660046108b5565b6102e5565b6040516100d39190610a76565b60405180910390f35b6100b1610311565b6100c66100f23660046108b5565b6103c5565b6100ff61042a565b6040516100d39190610a16565b610114610502565b6040516100d39190610a02565b610129610511565b6040516100d39190610c1b565b6100b16101443660046108b5565b610517565b6100c66101573660046108b5565b6105bf565b61012961016a3660046108b5565b6105e5565b6100b161017d366004610887565b61060e565b60006001826040516101949190610976565b908152604080519182900360209081018320908301909152548082529091506101d85760405162461bcd60e51b81526004016101cf90610b18565b60405180910390fd5b600254815111156101fb5760405162461bcd60e51b81526004016101cf90610a94565b6002805482516000198201929091600191908490811061021757fe5b9060005260206000200160405161022e9190610992565b90815260405190819003602001902055600280548290811061024c57fe5b90600052602060002001600260018460000151038154811061026a57fe5b90600052602060002001908054600181600116156101000203166002900461029392919061071c565b50600280548061029f57fe5b6001900381819060005260206000200160006102bb91906107af565b90556001836040516102cd9190610976565b90815260405190819003602001902060009055505050565b6000806001836040516102f89190610976565b908152604051908190036020019020541190505b919050565b610319610718565b6000546001600160a01b0390811691161461037b576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60405163374155e560e01b8152600090309063374155e5906103eb908590600401610a81565b600060405180830381600087803b15801561040557600080fd5b505af1925050508015610416575060015b6104225750600061030c565b50600161030c565b60606002805480602002602001604051908101604052809291908181526020016000905b828210156104f95760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156104e55780601f106104ba576101008083540402835291602001916104e5565b820191906000526020600020905b8154815290600101906020018083116104c857829003601f168201915b50505050508152602001906001019061044e565b50505050905090565b6000546001600160a01b031690565b60025490565b6001816040516105279190610976565b90815260405190819003602001902054156105545760405162461bcd60e51b81526004016101cf90610b8d565b600280546001810182556000919091528151610597917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace019060208401906107f6565b506002546040516001906105ac908490610976565b9081526040519081900360200190205550565b604051639b1574eb60e01b81526000903090639b1574eb906103eb908590600401610a81565b60006001826040516105f79190610976565b908152604051908190036020019020549050919050565b610616610718565b6000546001600160a01b03908116911614610678576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b0381166106bd5760405162461bcd60e51b8152600401808060200182810382526026815260200180610c556026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282610752576000855561079f565b82601f10610763578054855561079f565b8280016001018555821561079f57600052602060002091601f016020900482015b8281111561079f578254825591600101919060010190610784565b506107ab929150610872565b5090565b50805460018160011615610100020316600290046000825580601f106107d557506107f3565b601f0160209004906000526020600020908101906107f39190610872565b50565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928261082c576000855561079f565b82601f1061084557805160ff191683800117855561079f565b8280016001018555821561079f579182015b8281111561079f578251825591602001919060010190610857565b5b808211156107ab5760008155600101610873565b600060208284031215610898578081fd5b81356001600160a01b03811681146108ae578182fd5b9392505050565b600060208083850312156108c7578182fd5b823567ffffffffffffffff808211156108de578384fd5b818501915085601f8301126108f1578384fd5b8135818111156108fd57fe5b604051601f8201601f191681018501838111828210171561091a57fe5b6040528181528382018501881015610930578586fd5b818585018683013790810190930193909352509392505050565b60008151808452610962816020860160208601610c24565b601f01601f19169290920160200192915050565b60008251610988818460208701610c24565b9190910192915050565b60008083546001808216600081146109b157600181146109c8576109f7565b60ff198316865260028304607f16860193506109f7565b600283048786526020808720875b838110156109ef5781548a8201529085019082016109d6565b505050860193505b509195945050505050565b6001600160a01b0391909116815260200190565b6000602080830181845280855180835260408601915060408482028701019250838701855b82811015610a6957603f19888603018452610a5785835161094a565b94509285019290850190600101610a3b565b5092979650505050505050565b901515815260200190565b6000602082526108ae602083018461094a565b602080825260609082018190527f496e76616c696420696e6465782076616c756520666f7220746865207374726960408301527f6e672073746f726167652c206661696c656420746f2072656d6f766520746865908201527f20737472696e672066726f6d2074686520737472696e672073746f726167652e608082015260a00190565b6020808252604f908201527f54686520737472696e6720646f6573206e6f742065786973742c206661696c6560408201527f6420746f2072656d6f76652074686520737472696e672066726f6d207468652060608201526e39ba3934b7339039ba37b930b3b29760891b608082015260a00190565b60208082526062908201527f54686520737472696e672068617320616c7265616479206265656e206164646560408201527f6420746f207468652073746f726167652c206661696c656420746f206164642060608201527f746865206164647265737320746f2074686520737472696e672073746f726167608082015261329760f11b60a082015260c00190565b90815260200190565b60005b83811015610c3f578181015183820152602001610c27565b83811115610c4e576000848401525b5050505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373a2646970667358221220c9bba98d34519760ee09990f1b1baf2cfa169406aae7695a283a68be52dde19c64736f6c63430007060033496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a264697066735822122049d822c677a71217821f17014e9f1ad24a1538a3381a57b590f22fdd726f405e64736f6c63430007060033"

// DeploySystemReserve deploys a new Ethereum contract, binding an instance of SystemReserve to it.
func DeploySystemReserve(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemReserve, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemReserveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SystemReserveBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemReserve{SystemReserveCaller: SystemReserveCaller{contract: contract}, SystemReserveTransactor: SystemReserveTransactor{contract: contract}, SystemReserveFilterer: SystemReserveFilterer{contract: contract}}, nil
}

// SystemReserve is an auto generated Go binding around an Ethereum contract.
type SystemReserve struct {
	SystemReserveCaller     // Read-only binding to the contract
	SystemReserveTransactor // Write-only binding to the contract
	SystemReserveFilterer   // Log filterer for contract events
}

// SystemReserveCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemReserveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemReserveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemReserveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemReserveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemReserveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemReserveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemReserveSession struct {
	Contract     *SystemReserve    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemReserveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemReserveCallerSession struct {
	Contract *SystemReserveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SystemReserveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemReserveTransactorSession struct {
	Contract     *SystemReserveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SystemReserveRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemReserveRaw struct {
	Contract *SystemReserve // Generic contract binding to access the raw methods on
}

// SystemReserveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemReserveCallerRaw struct {
	Contract *SystemReserveCaller // Generic read-only contract binding to access the raw methods on
}

// SystemReserveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemReserveTransactorRaw struct {
	Contract *SystemReserveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemReserve creates a new instance of SystemReserve, bound to a specific deployed contract.
func NewSystemReserve(address common.Address, backend bind.ContractBackend) (*SystemReserve, error) {
	contract, err := bindSystemReserve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemReserve{SystemReserveCaller: SystemReserveCaller{contract: contract}, SystemReserveTransactor: SystemReserveTransactor{contract: contract}, SystemReserveFilterer: SystemReserveFilterer{contract: contract}}, nil
}

// NewSystemReserveCaller creates a new read-only instance of SystemReserve, bound to a specific deployed contract.
func NewSystemReserveCaller(address common.Address, caller bind.ContractCaller) (*SystemReserveCaller, error) {
	contract, err := bindSystemReserve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemReserveCaller{contract: contract}, nil
}

// NewSystemReserveTransactor creates a new write-only instance of SystemReserve, bound to a specific deployed contract.
func NewSystemReserveTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemReserveTransactor, error) {
	contract, err := bindSystemReserve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemReserveTransactor{contract: contract}, nil
}

// NewSystemReserveFilterer creates a new log filterer instance of SystemReserve, bound to a specific deployed contract.
func NewSystemReserveFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemReserveFilterer, error) {
	contract, err := bindSystemReserve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemReserveFilterer{contract: contract}, nil
}

// bindSystemReserve binds a generic wrapper to an already deployed contract.
func bindSystemReserve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemReserveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemReserve *SystemReserveRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemReserve.Contract.SystemReserveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemReserve *SystemReserveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemReserve.Contract.SystemReserveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemReserve *SystemReserveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemReserve.Contract.SystemReserveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemReserve *SystemReserveCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemReserve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemReserve *SystemReserveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemReserve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemReserve *SystemReserveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemReserve.Contract.contract.Transact(opts, method, params...)
}

// AvailableAmount is a free data retrieval call binding the contract method 0x91f7cfb9.
//
// Solidity: function availableAmount() view returns(uint256)
func (_SystemReserve *SystemReserveCaller) AvailableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemReserve.contract.Call(opts, out, "availableAmount")
	return *ret0, err
}

// AvailableAmount is a free data retrieval call binding the contract method 0x91f7cfb9.
//
// Solidity: function availableAmount() view returns(uint256)
func (_SystemReserve *SystemReserveSession) AvailableAmount() (*big.Int, error) {
	return _SystemReserve.Contract.AvailableAmount(&_SystemReserve.CallOpts)
}

// AvailableAmount is a free data retrieval call binding the contract method 0x91f7cfb9.
//
// Solidity: function availableAmount() view returns(uint256)
func (_SystemReserve *SystemReserveCallerSession) AvailableAmount() (*big.Int, error) {
	return _SystemReserve.Contract.AvailableAmount(&_SystemReserve.CallOpts)
}

// CoolDownPhase is a free data retrieval call binding the contract method 0x28011170.
//
// Solidity: function coolDownPhase() view returns(uint256)
func (_SystemReserve *SystemReserveCaller) CoolDownPhase(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemReserve.contract.Call(opts, out, "coolDownPhase")
	return *ret0, err
}

// CoolDownPhase is a free data retrieval call binding the contract method 0x28011170.
//
// Solidity: function coolDownPhase() view returns(uint256)
func (_SystemReserve *SystemReserveSession) CoolDownPhase() (*big.Int, error) {
	return _SystemReserve.Contract.CoolDownPhase(&_SystemReserve.CallOpts)
}

// CoolDownPhase is a free data retrieval call binding the contract method 0x28011170.
//
// Solidity: function coolDownPhase() view returns(uint256)
func (_SystemReserve *SystemReserveCallerSession) CoolDownPhase() (*big.Int, error) {
	return _SystemReserve.Contract.CoolDownPhase(&_SystemReserve.CallOpts)
}

// GetEligibleContractKeys is a free data retrieval call binding the contract method 0x2d250a5b.
//
// Solidity: function getEligibleContractKeys() view returns(string[])
func (_SystemReserve *SystemReserveCaller) GetEligibleContractKeys(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _SystemReserve.contract.Call(opts, out, "getEligibleContractKeys")
	return *ret0, err
}

// GetEligibleContractKeys is a free data retrieval call binding the contract method 0x2d250a5b.
//
// Solidity: function getEligibleContractKeys() view returns(string[])
func (_SystemReserve *SystemReserveSession) GetEligibleContractKeys() ([]string, error) {
	return _SystemReserve.Contract.GetEligibleContractKeys(&_SystemReserve.CallOpts)
}

// GetEligibleContractKeys is a free data retrieval call binding the contract method 0x2d250a5b.
//
// Solidity: function getEligibleContractKeys() view returns(string[])
func (_SystemReserve *SystemReserveCallerSession) GetEligibleContractKeys() ([]string, error) {
	return _SystemReserve.Contract.GetEligibleContractKeys(&_SystemReserve.CallOpts)
}

// SystemPaused is a free data retrieval call binding the contract method 0x9d2f83f0.
//
// Solidity: function systemPaused() view returns(bool)
func (_SystemReserve *SystemReserveCaller) SystemPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SystemReserve.contract.Call(opts, out, "systemPaused")
	return *ret0, err
}

// SystemPaused is a free data retrieval call binding the contract method 0x9d2f83f0.
//
// Solidity: function systemPaused() view returns(bool)
func (_SystemReserve *SystemReserveSession) SystemPaused() (bool, error) {
	return _SystemReserve.Contract.SystemPaused(&_SystemReserve.CallOpts)
}

// SystemPaused is a free data retrieval call binding the contract method 0x9d2f83f0.
//
// Solidity: function systemPaused() view returns(bool)
func (_SystemReserve *SystemReserveCallerSession) SystemPaused() (bool, error) {
	return _SystemReserve.Contract.SystemPaused(&_SystemReserve.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x94e6f56f.
//
// Solidity: function initialize(address _registry, string[] _keys) returns()
func (_SystemReserve *SystemReserveTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _keys []string) (*types.Transaction, error) {
	return _SystemReserve.contract.Transact(opts, "initialize", _registry, _keys)
}

// Initialize is a paid mutator transaction binding the contract method 0x94e6f56f.
//
// Solidity: function initialize(address _registry, string[] _keys) returns()
func (_SystemReserve *SystemReserveSession) Initialize(_registry common.Address, _keys []string) (*types.Transaction, error) {
	return _SystemReserve.Contract.Initialize(&_SystemReserve.TransactOpts, _registry, _keys)
}

// Initialize is a paid mutator transaction binding the contract method 0x94e6f56f.
//
// Solidity: function initialize(address _registry, string[] _keys) returns()
func (_SystemReserve *SystemReserveTransactorSession) Initialize(_registry common.Address, _keys []string) (*types.Transaction, error) {
	return _SystemReserve.Contract.Initialize(&_SystemReserve.TransactOpts, _registry, _keys)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xcdb88ad1.
//
// Solidity: function setPauseState(bool _state) returns()
func (_SystemReserve *SystemReserveTransactor) SetPauseState(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _SystemReserve.contract.Transact(opts, "setPauseState", _state)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xcdb88ad1.
//
// Solidity: function setPauseState(bool _state) returns()
func (_SystemReserve *SystemReserveSession) SetPauseState(_state bool) (*types.Transaction, error) {
	return _SystemReserve.Contract.SetPauseState(&_SystemReserve.TransactOpts, _state)
}

// SetPauseState is a paid mutator transaction binding the contract method 0xcdb88ad1.
//
// Solidity: function setPauseState(bool _state) returns()
func (_SystemReserve *SystemReserveTransactorSession) SetPauseState(_state bool) (*types.Transaction, error) {
	return _SystemReserve.Contract.SetPauseState(&_SystemReserve.TransactOpts, _state)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_SystemReserve *SystemReserveTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SystemReserve.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_SystemReserve *SystemReserveSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _SystemReserve.Contract.Withdraw(&_SystemReserve.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_SystemReserve *SystemReserveTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _SystemReserve.Contract.Withdraw(&_SystemReserve.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemReserve *SystemReserveTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemReserve.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemReserve *SystemReserveSession) Receive() (*types.Transaction, error) {
	return _SystemReserve.Contract.Receive(&_SystemReserve.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemReserve *SystemReserveTransactorSession) Receive() (*types.Transaction, error) {
	return _SystemReserve.Contract.Receive(&_SystemReserve.TransactOpts)
}
