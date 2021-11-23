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

// SystemSurplusAuctionAuctionInfo is an auto generated low-level Go binding around an user-defined struct.
type SystemSurplusAuctionAuctionInfo struct {
	IsExecuted bool
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
}

// SystemSurplusAuctionMetaData contains all meta data concerning the SystemSurplusAuction contract.
var SystemSurplusAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isExecuted\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSystemSurplusAuction.AuctionInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isExecuted\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_stc\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getRaisingBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611f42806100206000396000f3fe6080604052600436106100605760003560e01c8063454a2ab314610065578063571a26a01461008e5780636b64c769146100bf578063a2165920146100d4578063b830620a146100f4578063f399e22e14610109578063fe0d94c11461012b575b600080fd5b61007861007336600461194a565b61014b565b60405161008591906119aa565b60405180910390f35b34801561009a57600080fd5b506100ae6100a936600461194a565b610386565b6040516100859594939291906119b5565b6100c76103c1565b6040516100859190611e03565b3480156100e057600080fd5b506100c76100ef36600461194a565b610a34565b34801561010057600080fd5b506100c7610a68565b34801561011557600080fd5b50610129610124366004611880565b610a6e565b005b34801561013757600080fd5b5061007861014636600461194a565b610b48565b60008160025481106101785760405162461bcd60e51b815260040161016f90611baf565b60405180910390fd5b6000838152600360208190526040909120908101544211156101ac5760405162461bcd60e51b815260040161016f90611ced565b6101b584610ef2565b3410156101d45760405162461bcd60e51b815260040161016f90611a36565b60008054604080518082018252601b81527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb9027191610238916004016119e3565b60206040518083038186803b15801561025057600080fd5b505afa158015610264573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102889190611864565b600183018054845433610100908102610100600160a81b0319831617875534909355604051632377789f60e21b815293945090926001600160a01b0392909104821691841690638ddde27c9084906102e490859060040161197d565b6020604051808303818588803b1580156102fd57600080fd5b505af1158015610311573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250810190610336919061192a565b50336001600160a01b0316877fdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199346040516103719190611e03565b60405180910390a35060019695505050505050565b6003602081905260009182526040909120805460018201546002830154929093015460ff8216936101009092046001600160a01b0316929085565b6000346103e05760405162461bcd60e51b815260040161016f90611ac6565b6000805460018054604080516020600261010085871615026000190190941693909304601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb902719361049093919290918301828280156104865780601f1061045b57610100808354040283529160200191610486565b820191906000526020600020905b81548152906001019060200180831161046957829003601f168201915b5050505050610fd0565b6040518263ffffffff1660e01b81526004016104ac91906119e3565b60206040518083038186803b1580156104c457600080fd5b505afa1580156104d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104fc9190611864565b9050806001600160a01b0316633d03f7976040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561053957600080fd5b505af115801561054d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610571919061192a565b50600061057c611182565b60018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281529394506000936001600160a01b0386169363498bff00936106259383018282801561061b5780601f106105f05761010080835404028352916020019161061b565b820191906000526020600020905b8154815290600101906020018083116105fe57829003601f168201915b5050505050611231565b6040518263ffffffff1660e01b815260040161064191906119e3565b60206040518083038186803b15801561065957600080fd5b505afa15801561066d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106919190611962565b60018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281529394506000936001600160a01b0387169363498bff009361073a938301828280156107305780601f1061070557610100808354040283529160200191610730565b820191906000526020600020905b81548152906001019060200180831161071357829003601f168201915b5050505050611265565b6040518263ffffffff1660e01b815260040161075691906119e3565b60206040518083038186803b15801561076e57600080fd5b505afa158015610782573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107a69190611962565b90506000846001600160a01b0316632383b0746040518163ffffffff1660e01b815260040160206040518083038186803b1580156107e357600080fd5b505afa1580156107f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081b9190611962565b90508181101561083d5760405162461bcd60e51b815260040161016f90611c38565b8281101561085d5760405162461bcd60e51b815260040161016f90611b20565b60405163152353d960e01b81526001600160a01b0386169063152353d990610889908690600401611e03565b602060405180830381600087803b1580156108a357600080fd5b505af11580156108b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108db919061192a565b5060028054600181019091556108ef611789565b33602082015234604080830191909152606082018690525162498bff60e81b8152610988906001600160a01b0388169063498bff009061093190600401611b78565b60206040518083038186803b15801561094957600080fd5b505afa15801561095d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109819190611962565b4290611299565b6080820190815260008381526003602081815260409283902085518154928701516001600160a01b031661010002610100600160a81b031991151560ff1990941693909317169190911781558285015160018201556060850151600282015592519201919091555133907fd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f049247590610a219085903490611e50565b60405180910390a2509550505050505090565b6000816002548110610a585760405162461bcd60e51b815260040161016f90611baf565b610a6183610ef2565b9392505050565b60025481565b600054610100900460ff1680610a875750610a876112f1565b80610a95575060005460ff16155b610ad05760405162461bcd60e51b815260040180806020018281038252602e815260200180611ebe602e913960400191505060405180910390fd5b600054610100900460ff16158015610afb576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038616021790558151610b309060019060208501906117c3565b508015610b43576000805461ff00191690555b505050565b6000816002548110610b6c5760405162461bcd60e51b815260040161016f90611baf565b6000838152600360208190526040909120908101544211610b9f5760405162461bcd60e51b815260040161016f90611ca4565b805460ff1615610bc15760405162461bcd60e51b815260040161016f90611d41565b6000805460018054604080516020600261010085871615026000190190941693909304601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb9027193610c719391929091830182828015610c675780601f10610c3c57610100808354040283529160200191610c67565b820191906000526020600020905b815481529060010190602001808311610c4a57829003601f168201915b5050505050611302565b6040518263ffffffff1660e01b8152600401610c8d91906119e3565b60206040518083038186803b158015610ca557600080fd5b505afa158015610cb9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cdd9190611864565b8254600284015460405163a9059cbb60e01b81529293506001600160a01b038085169363a9059cbb93610d1d936101009091049092169190600401611991565b602060405180830381600087803b158015610d3757600080fd5b505af1158015610d4b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d6f919061192a565b508154600160ff199091161782556000805460408051606081019091526025808252620100009092046001600160a01b031691633fb902719190611e7760208301396040518263ffffffff1660e01b8152600401610dcd91906119e3565b60206040518083038186803b158015610de557600080fd5b505afa158015610df9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1d9190611864565b6001600160a01b03168360010154604051610e379061197a565b60006040518083038185875af1925050503d8060008114610e74576040519150601f19603f3d011682016040523d82523d6000602084013e610e79565b606091505b5050905080610e9a5760405162461bcd60e51b815260040161016f90611d94565b82546040516101009091046001600160a01b0316907f2510a5567d82839751607a8a4f3d96799802809b132d2364298542dc215ce49b90610ede9089908790611e0c565b60405180910390a250600195945050505050565b600080610efd611182565b6001600160a01b031663498bff006040518163ffffffff1660e01b8152600401610f2690611bf7565b60206040518083038186803b158015610f3e57600080fd5b505afa158015610f52573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f769190611962565b600084815260036020526040812060010154919250610fb082610faa610f9a6114ab565b610fa486886114bb565b90611514565b90611299565b905081811415610fc857610fc5816001611299565b90505b949350505050565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106110565780518252601f199092019160209182019101611037565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b6020831061109e5780518252601f19909201916020918201910161107f565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106110e65780518252601f1990920191602091820191016110c7565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061112e5780518252601f19909201916020918201910161110f565b5181516020939093036101000a60001901801990911692169190911790526c73797374656d42616c616e636560981b92019182525060408051808303601219018152600d9092019052979650505050505050565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060228152602001611e9c602291396040518263ffffffff1660e01b81526004016111dc91906119e3565b60206040518083038186803b1580156111f457600080fd5b505afa158015611208573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061122c9190611864565b905090565b606061125f826040518060400160405280600a8152602001691cdd5c9c1b1d5cd31bdd60b21b815250611553565b92915050565b606061125f826040518060400160405280601081526020016f1cdd5c9c1b1d5cd51a1c995cda1bdb1960821b815250611553565b600082820183811015610a61576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b60006112fc306116e1565b15905090565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106113885780518252601f199092019160209182019101611369565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106113d05780518252601f1990920191602091820191016113b1565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106114185780518252601f1990920191602091820191016113f9565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106114605780518252601f199092019160209182019101611441565b5181516020939093036101000a60001901801990911692169190911790526331b7b4b760e11b92019182525060408051808303601b1901815260049092019052979650505050505050565b6b033b2e3c9fd0803ce800000090565b6000826114ca5750600061125f565b828202828482816114d757fe5b0414610a615760405162461bcd60e51b8152600401808060200182810382526021815260200180611eec6021913960400191505060405180910390fd5b6000610a6183836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b8152506116e7565b60606040518060400160405280600e81526020016d33b7bb32b93732b21722a822291760911b81525083604051806040016040528060018152602001605f60f81b815250846040516020018085805190602001908083835b602083106115ca5780518252601f1990920191602091820191016115ab565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106116125780518252601f1990920191602091820191016115f3565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b6020831061165a5780518252601f19909201916020918201910161163b565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106116a25780518252601f199092019160209182019101611683565b6001836020036101000a038019825116818451168082178552505050505050905001945050505050604051602081830303815290604052905092915050565b3b151590565b600081836117735760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611738578181015183820152602001611720565b50505050905090810190601f1680156117655780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161177f57fe5b0495945050505050565b6040518060a0016040528060001515815260200160006001600160a01b031681526020016000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826117f9576000855561183f565b82601f1061181257805160ff191683800117855561183f565b8280016001018555821561183f579182015b8281111561183f578251825591602001919060010190611824565b5061184b92915061184f565b5090565b5b8082111561184b5760008155600101611850565b600060208284031215611875578081fd5b8151610a6181611e5e565b60008060408385031215611892578081fd5b823561189d81611e5e565b915060208381013567ffffffffffffffff808211156118ba578384fd5b818601915086601f8301126118cd578384fd5b8135818111156118d957fe5b604051601f8201601f19168101850183811182821017156118f657fe5b604052818152838201850189101561190c578586fd5b81858501868301378585838301015280955050505050509250929050565b60006020828403121561193b578081fd5b81518015158114610a61578182fd5b60006020828403121561195b578081fd5b5035919050565b600060208284031215611973578081fd5b5051919050565b90565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b94151585526001600160a01b0393909316602085015260408401919091526060830152608082015260a00190565b6000602080835283518082850152825b81811015611a0f578581018301518582016040015282016119f3565b81811115611a205783604083870101525b50601f01601f1916929092016040019392505050565b60208082526064908201527f5b5145432d3032353030365d2d5468652062696420616d6f756e74206d75737460408201527f206578636565642074686520686967686573742062696420627920746865206d60608201527f696e696d756d20696e6372656d656e742070657263656e74616765206f72206d60808201526337b9329760e11b60a082015260c00190565b6020808252603a908201527f5b5145432d3032353030305d2d496e76616c69642062696420616d6f756e742c604082015279103330b4b632b21039ba30b93a103a34329030bab1ba34b7b71760311b606082015260800190565b60208082526038908201527f5b5145432d3032353030335d2d4e6f7420656e6f75676820737572706c7573206040820152773a37903334b636103a34329030bab1ba34b7b7103637ba1760411b606082015260800190565b6020808252601d908201527f676f7665726e65642e455044522e737572706c757341756374696f6e50000000604082015260600190565b60208082526028908201527f5b5145432d3032353030345d2d5468652061756374696f6e20646f6573206e6f6040820152673a1032bc34b9ba1760c11b606082015260800190565b60208082526021908201527f676f7665726e65642e455044522e61756374696f6e4d696e496e6372656d656e6040820152601d60fa1b606082015260800190565b60208082526046908201527f5b5145432d3032353030325d2d537572706c7573206166746572206e6574746960408201527f6e672069732062656c6f7720737572706c75732061756374696f6e207468726560608201526539b437b6321760d11b608082015260a00190565b60208082526029908201527f5b5145432d3032353030385d2d5468652061756374696f6e206973206e6f74206040820152683334b734b9b432b21760b91b606082015260800190565b60208082526034908201527f5b5145432d3032353030355d2d5468652061756374696f6e2069732066696e6960408201527339b432b216103330b4b632b2103a37903134b21760611b606082015260800190565b60208082526033908201527f5b5145432d3032353030395d2d5468652061756374696f6e2068617320616c7260408201527232b0b23c903132b2b71032bc32b1baba32b21760691b606082015260800190565b60208082526049908201527f5145432d3032353031302d4661696c656420746f207472616e7366657220746860408201527f6520686967686573742062696420666f7220636f6d6d756e69747920646973746060820152683934b13aba34b7b71760b91b608082015260a00190565b90815260200190565b918252805460ff81161515602084015260081c6001600160a01b0316604083015260018101546060830152600281015460808301526003015460a082015260c00190565b918252602082015260400190565b6001600160a01b0381168114611e7357600080fd5b5056fe746f6b656e65636f6e6f6d6963732e64656661756c74416c6c6f636174696f6e50726f7879676f7665726e616e63652e657870657274732e455044522e706172616d6574657273496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a264697066735822122068b78d8910009aaf54cd947c2e6c42762bed7761b477a950bbf2a45457e091fc64736f6c63430007060033",
}

// SystemSurplusAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemSurplusAuctionMetaData.ABI instead.
var SystemSurplusAuctionABI = SystemSurplusAuctionMetaData.ABI

// SystemSurplusAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemSurplusAuctionMetaData.Bin instead.
var SystemSurplusAuctionBin = SystemSurplusAuctionMetaData.Bin

// DeploySystemSurplusAuction deploys a new Ethereum contract, binding an instance of SystemSurplusAuction to it.
func DeploySystemSurplusAuction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemSurplusAuction, error) {
	parsed, err := SystemSurplusAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemSurplusAuctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemSurplusAuction{SystemSurplusAuctionCaller: SystemSurplusAuctionCaller{contract: contract}, SystemSurplusAuctionTransactor: SystemSurplusAuctionTransactor{contract: contract}, SystemSurplusAuctionFilterer: SystemSurplusAuctionFilterer{contract: contract}}, nil
}

// SystemSurplusAuction is an auto generated Go binding around an Ethereum contract.
type SystemSurplusAuction struct {
	SystemSurplusAuctionCaller     // Read-only binding to the contract
	SystemSurplusAuctionTransactor // Write-only binding to the contract
	SystemSurplusAuctionFilterer   // Log filterer for contract events
}

// SystemSurplusAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemSurplusAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSurplusAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemSurplusAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSurplusAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemSurplusAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSurplusAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemSurplusAuctionSession struct {
	Contract     *SystemSurplusAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SystemSurplusAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemSurplusAuctionCallerSession struct {
	Contract *SystemSurplusAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SystemSurplusAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemSurplusAuctionTransactorSession struct {
	Contract     *SystemSurplusAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SystemSurplusAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemSurplusAuctionRaw struct {
	Contract *SystemSurplusAuction // Generic contract binding to access the raw methods on
}

// SystemSurplusAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemSurplusAuctionCallerRaw struct {
	Contract *SystemSurplusAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SystemSurplusAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemSurplusAuctionTransactorRaw struct {
	Contract *SystemSurplusAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemSurplusAuction creates a new instance of SystemSurplusAuction, bound to a specific deployed contract.
func NewSystemSurplusAuction(address common.Address, backend bind.ContractBackend) (*SystemSurplusAuction, error) {
	contract, err := bindSystemSurplusAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuction{SystemSurplusAuctionCaller: SystemSurplusAuctionCaller{contract: contract}, SystemSurplusAuctionTransactor: SystemSurplusAuctionTransactor{contract: contract}, SystemSurplusAuctionFilterer: SystemSurplusAuctionFilterer{contract: contract}}, nil
}

// NewSystemSurplusAuctionCaller creates a new read-only instance of SystemSurplusAuction, bound to a specific deployed contract.
func NewSystemSurplusAuctionCaller(address common.Address, caller bind.ContractCaller) (*SystemSurplusAuctionCaller, error) {
	contract, err := bindSystemSurplusAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionCaller{contract: contract}, nil
}

// NewSystemSurplusAuctionTransactor creates a new write-only instance of SystemSurplusAuction, bound to a specific deployed contract.
func NewSystemSurplusAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemSurplusAuctionTransactor, error) {
	contract, err := bindSystemSurplusAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionTransactor{contract: contract}, nil
}

// NewSystemSurplusAuctionFilterer creates a new log filterer instance of SystemSurplusAuction, bound to a specific deployed contract.
func NewSystemSurplusAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemSurplusAuctionFilterer, error) {
	contract, err := bindSystemSurplusAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionFilterer{contract: contract}, nil
}

// bindSystemSurplusAuction binds a generic wrapper to an already deployed contract.
func bindSystemSurplusAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemSurplusAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemSurplusAuction *SystemSurplusAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemSurplusAuction.Contract.SystemSurplusAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemSurplusAuction *SystemSurplusAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.SystemSurplusAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemSurplusAuction *SystemSurplusAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.SystemSurplusAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemSurplusAuction *SystemSurplusAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemSurplusAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.contract.Transact(opts, method, params...)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(bool isExecuted, address bidder, uint256 highestBid, uint256 lot, uint256 endTime)
func (_SystemSurplusAuction *SystemSurplusAuctionCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsExecuted bool
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
}, error) {
	var out []interface{}
	err := _SystemSurplusAuction.contract.Call(opts, &out, "auctions", arg0)

	outstruct := new(struct {
		IsExecuted bool
		Bidder     common.Address
		HighestBid *big.Int
		Lot        *big.Int
		EndTime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsExecuted = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Bidder = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.HighestBid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Lot = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(bool isExecuted, address bidder, uint256 highestBid, uint256 lot, uint256 endTime)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) Auctions(arg0 *big.Int) (struct {
	IsExecuted bool
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
}, error) {
	return _SystemSurplusAuction.Contract.Auctions(&_SystemSurplusAuction.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(bool isExecuted, address bidder, uint256 highestBid, uint256 lot, uint256 endTime)
func (_SystemSurplusAuction *SystemSurplusAuctionCallerSession) Auctions(arg0 *big.Int) (struct {
	IsExecuted bool
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
}, error) {
	return _SystemSurplusAuction.Contract.Auctions(&_SystemSurplusAuction.CallOpts, arg0)
}

// AuctionsCount is a free data retrieval call binding the contract method 0xb830620a.
//
// Solidity: function auctionsCount() view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionCaller) AuctionsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemSurplusAuction.contract.Call(opts, &out, "auctionsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionsCount is a free data retrieval call binding the contract method 0xb830620a.
//
// Solidity: function auctionsCount() view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) AuctionsCount() (*big.Int, error) {
	return _SystemSurplusAuction.Contract.AuctionsCount(&_SystemSurplusAuction.CallOpts)
}

// AuctionsCount is a free data retrieval call binding the contract method 0xb830620a.
//
// Solidity: function auctionsCount() view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionCallerSession) AuctionsCount() (*big.Int, error) {
	return _SystemSurplusAuction.Contract.AuctionsCount(&_SystemSurplusAuction.CallOpts)
}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionCaller) GetRaisingBid(opts *bind.CallOpts, _auctionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SystemSurplusAuction.contract.Call(opts, &out, "getRaisingBid", _auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) GetRaisingBid(_auctionId *big.Int) (*big.Int, error) {
	return _SystemSurplusAuction.Contract.GetRaisingBid(&_SystemSurplusAuction.CallOpts, _auctionId)
}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionCallerSession) GetRaisingBid(_auctionId *big.Int) (*big.Int, error) {
	return _SystemSurplusAuction.Contract.GetRaisingBid(&_SystemSurplusAuction.CallOpts, _auctionId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _auctionId) payable returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactor) Bid(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.contract.Transact(opts, "bid", _auctionId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _auctionId) payable returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) Bid(_auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Bid(&_SystemSurplusAuction.TransactOpts, _auctionId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _auctionId) payable returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorSession) Bid(_auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Bid(&_SystemSurplusAuction.TransactOpts, _auctionId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _auctionId) returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactor) Execute(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.contract.Transact(opts, "execute", _auctionId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _auctionId) returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) Execute(_auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Execute(&_SystemSurplusAuction.TransactOpts, _auctionId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _auctionId) returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorSession) Execute(_auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Execute(&_SystemSurplusAuction.TransactOpts, _auctionId)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemSurplusAuction *SystemSurplusAuctionTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemSurplusAuction.contract.Transact(opts, "initialize", _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemSurplusAuction *SystemSurplusAuctionSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Initialize(&_SystemSurplusAuction.TransactOpts, _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Initialize(&_SystemSurplusAuction.TransactOpts, _registry, _stc)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() payable returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactor) StartAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemSurplusAuction.contract.Transact(opts, "startAuction")
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() payable returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) StartAuction() (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.StartAuction(&_SystemSurplusAuction.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() payable returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorSession) StartAuction() (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.StartAuction(&_SystemSurplusAuction.TransactOpts)
}

// SystemSurplusAuctionAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionAuctionStartedIterator struct {
	Event *SystemSurplusAuctionAuctionStarted // Event containing the contract specifics and raw log

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
func (it *SystemSurplusAuctionAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemSurplusAuctionAuctionStarted)
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
		it.Event = new(SystemSurplusAuctionAuctionStarted)
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
func (it *SystemSurplusAuctionAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemSurplusAuctionAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemSurplusAuctionAuctionStarted represents a AuctionStarted event raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionAuctionStarted struct {
	AuctionId *big.Int
	Bidder    common.Address
	Bid       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0xd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475.
//
// Solidity: event AuctionStarted(uint256 _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) FilterAuctionStarted(opts *bind.FilterOpts, _bidder []common.Address) (*SystemSurplusAuctionAuctionStartedIterator, error) {

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.FilterLogs(opts, "AuctionStarted", _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionAuctionStartedIterator{contract: _SystemSurplusAuction.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0xd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475.
//
// Solidity: event AuctionStarted(uint256 _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *SystemSurplusAuctionAuctionStarted, _bidder []common.Address) (event.Subscription, error) {

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.WatchLogs(opts, "AuctionStarted", _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemSurplusAuctionAuctionStarted)
				if err := _SystemSurplusAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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

// ParseAuctionStarted is a log parse operation binding the contract event 0xd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475.
//
// Solidity: event AuctionStarted(uint256 _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) ParseAuctionStarted(log types.Log) (*SystemSurplusAuctionAuctionStarted, error) {
	event := new(SystemSurplusAuctionAuctionStarted)
	if err := _SystemSurplusAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemSurplusAuctionBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionBidIterator struct {
	Event *SystemSurplusAuctionBid // Event containing the contract specifics and raw log

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
func (it *SystemSurplusAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemSurplusAuctionBid)
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
		it.Event = new(SystemSurplusAuctionBid)
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
func (it *SystemSurplusAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemSurplusAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemSurplusAuctionBid represents a Bid event raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionBid struct {
	AuctionId *big.Int
	Bidder    common.Address
	Bid       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199.
//
// Solidity: event Bid(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) FilterBid(opts *bind.FilterOpts, _auctionId []*big.Int, _bidder []common.Address) (*SystemSurplusAuctionBidIterator, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.FilterLogs(opts, "Bid", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionBidIterator{contract: _SystemSurplusAuction.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199.
//
// Solidity: event Bid(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *SystemSurplusAuctionBid, _auctionId []*big.Int, _bidder []common.Address) (event.Subscription, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.WatchLogs(opts, "Bid", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemSurplusAuctionBid)
				if err := _SystemSurplusAuction.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0xdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199.
//
// Solidity: event Bid(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) ParseBid(log types.Log) (*SystemSurplusAuctionBid, error) {
	event := new(SystemSurplusAuctionBid)
	if err := _SystemSurplusAuction.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemSurplusAuctionExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionExecutedIterator struct {
	Event *SystemSurplusAuctionExecuted // Event containing the contract specifics and raw log

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
func (it *SystemSurplusAuctionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemSurplusAuctionExecuted)
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
		it.Event = new(SystemSurplusAuctionExecuted)
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
func (it *SystemSurplusAuctionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemSurplusAuctionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemSurplusAuctionExecuted represents a Executed event raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionExecuted struct {
	AuctionId *big.Int
	Bidder    common.Address
	Info      SystemSurplusAuctionAuctionInfo
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x2510a5567d82839751607a8a4f3d96799802809b132d2364298542dc215ce49b.
//
// Solidity: event Executed(uint256 _auctionId, address indexed _bidder, (bool,address,uint256,uint256,uint256) _info)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) FilterExecuted(opts *bind.FilterOpts, _bidder []common.Address) (*SystemSurplusAuctionExecutedIterator, error) {

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.FilterLogs(opts, "Executed", _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionExecutedIterator{contract: _SystemSurplusAuction.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x2510a5567d82839751607a8a4f3d96799802809b132d2364298542dc215ce49b.
//
// Solidity: event Executed(uint256 _auctionId, address indexed _bidder, (bool,address,uint256,uint256,uint256) _info)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *SystemSurplusAuctionExecuted, _bidder []common.Address) (event.Subscription, error) {

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.WatchLogs(opts, "Executed", _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemSurplusAuctionExecuted)
				if err := _SystemSurplusAuction.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x2510a5567d82839751607a8a4f3d96799802809b132d2364298542dc215ce49b.
//
// Solidity: event Executed(uint256 _auctionId, address indexed _bidder, (bool,address,uint256,uint256,uint256) _info)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) ParseExecuted(log types.Log) (*SystemSurplusAuctionExecuted, error) {
	event := new(SystemSurplusAuctionExecuted)
	if err := _SystemSurplusAuction.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
