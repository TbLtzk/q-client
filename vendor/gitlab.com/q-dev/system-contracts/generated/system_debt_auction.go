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

// SystemDebtAuctionAuctionInfo is an auto generated low-level Go binding around an user-defined struct.
type SystemDebtAuctionAuctionInfo struct {
	Status     uint8
	Bidder     common.Address
	HighestBid *big.Int
	EndTime    *big.Int
	ReserveLot *big.Int
}

// SystemDebtAuctionABI is the input ABI used to generate the binding from.
const SystemDebtAuctionABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSystemDebtAuction.AuctionStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLot\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSystemDebtAuction.AuctionInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"enumSystemDebtAuction.AuctionStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentAuctionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_stc\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasActiveAuction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"startAuction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getRaisingBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SystemDebtAuctionBin is the compiled bytecode used for deploying new contracts.
var SystemDebtAuctionBin = "0x608060405234801561001057600080fd5b50612605806100206000396000f3fe60806040526004361061006f5760003560e01c8063065de74c1461007b5780630c0b86ca146100b1578063454a2ab3146100d3578063500ff7e6146100f3578063571a26a0146101085780636146195414610139578063a21659201461014e578063f399e22e1461016e57610076565b3661007657005b600080fd5b34801561008757600080fd5b5061009b610096366004611ee2565b610190565b6040516100a89190611f71565b60405180910390f35b3480156100bd57600080fd5b506100c6610a7a565b6040516100a8919061255f565b3480156100df57600080fd5b5061009b6100ee366004611ee2565b610a80565b3480156100ff57600080fd5b5061009b610f07565b34801561011457600080fd5b50610128610123366004611ee2565b610f33565b6040516100a8959493929190611f7c565b34801561014557600080fd5b5061009b610f6e565b34801561015a57600080fd5b506100c6610169366004611ee2565b611432565b34801561017a57600080fd5b5061018e610189366004611e18565b611465565b005b60008054604051633fb9027160e01b8152620100009091046001600160a01b03169082908290633fb90271906101c89060040161214b565b60206040518083038186803b1580156101e057600080fd5b505afa1580156101f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102189190611dfc565b60018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281529394506000936001600160a01b03871693633fb90271936102c1938301828280156102b75780601f1061028c576101008083540402835291602001916102b7565b820191906000526020600020905b81548152906001019060200180831161029a57829003601f168201915b505050505061153f565b6040518263ffffffff1660e01b81526004016102dd9190611fb6565b60206040518083038186803b1580156102f557600080fd5b505afa158015610309573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061032d9190611dfc565b90506000836001600160a01b0316633fb902716040518163ffffffff1660e01b815260040161035b9061208c565b60206040518083038186803b15801561037357600080fd5b505afa158015610387573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ab9190611dfc565b6003546000818152600260208190526040808320815160a0810190925280549596509394929390929091839160ff16908111156103e457fe5b60028111156103ef57fe5b8152815461010090046001600160a01b031660208201526001808301546040830152600283015460608301526003909201546080909101529091508151600281111561043757fe5b141561045e5760405162461bcd60e51b815260040161045590612420565b60405180910390fd5b836001600160a01b0316633d03f7976040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561049957600080fd5b505af11580156104ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d19190611ec2565b5060018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281526001600160a01b0389169363498bff00936105759383018282801561056b5780601f106105405761010080835404028352916020019161056b565b820191906000526020600020905b81548152906001019060200180831161054e57829003601f168201915b50505050506116f1565b6040518263ffffffff1660e01b81526004016105919190611fb6565b60206040518083038186803b1580156105a957600080fd5b505afa1580156105bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e19190611efa565b846001600160a01b03166314a6bf0f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561061a57600080fd5b505afa15801561062e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106529190611efa565b1161066f5760405162461bcd60e51b8152600401610455906122a6565b60018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281526001600160a01b038a1693633fb9027193610712938301828280156107085780601f106106dd57610100808354040283529160200191610708565b820191906000526020600020905b8154815290600101906020018083116106eb57829003601f168201915b5050505050611722565b6040518263ffffffff1660e01b815260040161072e9190611fb6565b60206040518083038186803b15801561074657600080fd5b505afa15801561075a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077e9190611dfc565b6001600160a01b03166323b872dd33308b6040518463ffffffff1660e01b81526004016107ad93929190611f34565b602060405180830381600087803b1580156107c757600080fd5b505af11580156107db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ff9190611ec2565b5060405162498bff60e81b81526000906001600160a01b0387169063498bff009061082c9060040161248e565b60206040518083038186803b15801561084457600080fd5b505afa158015610858573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087c9190611efa565b604051632e1a7d4d60e01b81529091506001600160a01b03851690632e1a7d4d906108ab90849060040161255f565b602060405180830381600087803b1580156108c557600080fd5b505af11580156108d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fd9190611ec2565b6109195760405162461bcd60e51b815260040161045590612009565b6001825233602083015260408083018a9052608083018290525162498bff60e81b81526109b4906001600160a01b0388169063498bff009061095d90600401612117565b60206040518083038186803b15801561097557600080fd5b505afa158015610989573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ad9190611efa565b42906118cb565b60608301526000838152600260208190526040909120835181548593839160ff19169060019084908111156109e557fe5b0217905550602082015181546001600160a01b0390911661010002610100600160a81b031990911617815560408083015160018301556060830151600283015560809092015160039091015551339084907fd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f049247590610a63908d9061255f565b60405180910390a350600198975050505050505050565b60035481565b60008054600354808352600260208190526040808520815160a081019092528054620100009095046001600160a01b0316948693839160ff1690811115610ac357fe5b6002811115610ace57fe5b8152815461010090046001600160a01b0316602082015260018083015460408301526002830154606083015260039092015460809091015290915081516002811115610b1657fe5b14610b335760405162461bcd60e51b8152600401610455906124c0565b4281606001511015610b575760405162461bcd60e51b8152600401610455906120c3565b610b6082611432565b851015610b7f5760405162461bcd60e51b81526004016104559061235b565b60018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281526001600160a01b03871693633fb9027193610bed938301828280156107085780601f106106dd57610100808354040283529160200191610708565b6040518263ffffffff1660e01b8152600401610c099190611fb6565b60206040518083038186803b158015610c2157600080fd5b505afa158015610c35573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c599190611dfc565b6001600160a01b031663a9059cbb826020015183604001516040518363ffffffff1660e01b8152600401610c8e929190611f58565b602060405180830381600087803b158015610ca857600080fd5b505af1158015610cbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce09190611ec2565b5060018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281526001600160a01b03871693633fb9027193610d4f938301828280156107085780601f106106dd57610100808354040283529160200191610708565b6040518263ffffffff1660e01b8152600401610d6b9190611fb6565b60206040518083038186803b158015610d8357600080fd5b505afa158015610d97573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dbb9190611dfc565b6001600160a01b03166323b872dd3330886040518463ffffffff1660e01b8152600401610dea93929190611f34565b602060405180830381600087803b158015610e0457600080fd5b505af1158015610e18573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3c9190611ec2565b50336020808301919091526040808301879052600084815260029283905220825181548493839160ff1916906001908490811115610e7657fe5b0217905550602082015181546001600160a01b0390911661010002610100600160a81b031990911617815560408083015160018301556060830151600283015560809092015160039091015551339083907fdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca319990610ef490899061255f565b60405180910390a3506001949350505050565b6000600160035460009081526002602081905260409091205460ff1690811115610f2d57fe5b14905090565b6002602081905260009182526040909120805460018201549282015460039092015460ff8216936101009092046001600160a01b0316929085565b60035460008054828252600260208190526040808420815160a081019092528054949594620100009094046001600160a01b0316938693839160ff1690811115610fb457fe5b6002811115610fbf57fe5b81526020016000820160019054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152602001600182015481526020016002820154815260200160038201548152505090506000826001600160a01b0316633fb9027161109160018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102b75780601f1061028c576101008083540402835291602001916102b7565b6040518263ffffffff1660e01b81526004016110ad9190611fb6565b60206040518083038186803b1580156110c557600080fd5b505afa1580156110d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110fd9190611dfc565b60018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281529394506000936001600160a01b03881693633fb9027193611171938301828280156107085780601f106106dd57610100808354040283529160200191610708565b6040518263ffffffff1660e01b815260040161118d9190611fb6565b60206040518083038186803b1580156111a557600080fd5b505afa1580156111b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111dd9190611dfc565b9050826060015142116112025760405162461bcd60e51b81526004016104559061218d565b60018351600281111561121157fe5b1461122e5760405162461bcd60e51b8152600401610455906121fc565b604080840151905163a9059cbb60e01b81526001600160a01b0383169163a9059cbb9161125f918691600401611f58565b602060405180830381600087803b15801561127957600080fd5b505af115801561128d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112b19190611ec2565b50600380546000908152600260208190526040808320805460ff191690921790915582546001019092559051633fb9027160e01b81526001600160a01b03861690633fb9027190611304906004016123eb565b60206040518083038186803b15801561131c57600080fd5b505afa158015611330573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113549190611dfc565b60808501516020860151604051632377789f60e21b81529293506001600160a01b03841692638ddde27c929161138c91600401611f20565b6020604051808303818588803b1580156113a557600080fd5b505af11580156113b9573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906113de9190611ec2565b5083602001516001600160a01b0316867f342522ce238e8878f82d1a020742d0d67f988f004878a3b9842b52b02d278d648660405161141d9190612516565b60405180910390a36001965050505050505090565b60006003548211156114565760405162461bcd60e51b815260040161045590612313565b61145f8261192a565b92915050565b600054610100900460ff168061147e575061147e611a82565b8061148c575060005460ff16155b6114c75760405162461bcd60e51b815260040180806020018281038252602e815260200180612581602e913960400191505060405180910390fd5b600054610100900460ff161580156114f2576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038616021790558151611527906001906020850190611d5b565b50801561153a576000805461ff00191690555b505050565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106115c55780518252601f1990920191602091820191016115a6565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b6020831061160d5780518252601f1990920191602091820191016115ee565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106116555780518252601f199092019160209182019101611636565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061169d5780518252601f19909201916020918201910161167e565b5181516020939093036101000a60001901801990911692169190911790526c73797374656d42616c616e636560981b92019182525060408051808303601219018152600d9092019052979650505050505050565b606061145f826040518060400160405280600d81526020016c1919589d151a1c995cda1bdb19609a1b815250611a88565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106117a85780518252601f199092019160209182019101611789565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106117f05780518252601f1990920191602091820191016117d1565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106118385780518252601f199092019160209182019101611819565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106118805780518252601f199092019160209182019101611861565b5181516020939093036101000a60001901801990911692169190911790526331b7b4b760e11b92019182525060408051808303601b1901815260049092019052979650505050505050565b600082820183811015611923576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b9392505050565b60008054604051633fb9027160e01b815282916201000090046001600160a01b031690633fb902719061195f9060040161214b565b60206040518083038186803b15801561197757600080fd5b505afa15801561198b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119af9190611dfc565b6001600160a01b031663498bff006040518163ffffffff1660e01b81526004016119d890612265565b60206040518083038186803b1580156119f057600080fd5b505afa158015611a04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a289190611efa565b600084815260026020526040812060010154919250611a6282611a5c611a4c611c16565b611a568688611c26565b90611c7f565b906118cb565b905081811415611a7a57611a778160016118cb565b90505b949350505050565b303b1590565b60606040518060400160405280600e81526020016d33b7bb32b93732b21722a822291760911b81525083604051806040016040528060018152602001605f60f81b815250846040516020018085805190602001908083835b60208310611aff5780518252601f199092019160209182019101611ae0565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b60208310611b475780518252601f199092019160209182019101611b28565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310611b8f5780518252601f199092019160209182019101611b70565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310611bd75780518252601f199092019160209182019101611bb8565b6001836020036101000a038019825116818451168082178552505050505050905001945050505050604051602081830303815290604052905092915050565b6b033b2e3c9fd0803ce800000090565b600082611c355750600061145f565b82820282848281611c4257fe5b04146119235760405162461bcd60e51b81526004018080602001828103825260218152602001806125af6021913960400191505060405180910390fd5b600061192383836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b81525060008183611d455760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611d0a578181015183820152602001611cf2565b50505050905090810190601f168015611d375780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838581611d5157fe5b0495945050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282611d915760008555611dd7565b82601f10611daa57805160ff1916838001178555611dd7565b82800160010185558215611dd7579182015b82811115611dd7578251825591602001919060010190611dbc565b50611de3929150611de7565b5090565b5b80821115611de35760008155600101611de8565b600060208284031215611e0d578081fd5b815161192381612568565b60008060408385031215611e2a578081fd5b8235611e3581612568565b915060208381013567ffffffffffffffff80821115611e52578384fd5b818601915086601f830112611e65578384fd5b813581811115611e7157fe5b604051601f8201601f1916810185018381118282101715611e8e57fe5b6040528181528382018501891015611ea4578586fd5b81858501868301378585838301015280955050505050509250929050565b600060208284031215611ed3578081fd5b81518015158114611923578182fd5b600060208284031215611ef3578081fd5b5035919050565b600060208284031215611f0b578081fd5b5051919050565b60038110611f1c57fe5b9052565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b60a08101611f8a8288611f12565b6001600160a01b0395909516602082015260408101939093526060830191909152608090910152919050565b6000602080835283518082850152825b81811015611fe257858101830151858201604001528201611fc6565b81811115611ff35783604083870101525b50601f01601f1916929092016040019392505050565b6020808252605d908201527f5b5145432d3032363030325d2d4661696c656420746f2077697468647261772060408201527f66726f6d207468652053797374656d5265736572766520636f6e74726163742c60608201527f206661696c656420746f207374617274207468652061756374696f6e2e000000608082015260a00190565b6020808252601c908201527f746f6b656e65636f6e6f6d6963732e73797374656d5265736572766500000000604082015260600190565b60208082526034908201527f5b5145432d3032363030345d2d5468652061756374696f6e2069732066696e6960408201527339b432b216103330b4b632b2103a37903134b21760611b606082015260800190565b6020808252601a90820152790676f7665726e65642e455044522e6465627441756374696f6e560341b604082015260600190565b60208082526022908201527f676f7665726e616e63652e657870657274732e455044522e706172616d657465604082015261727360f01b606082015260800190565b60208082526049908201527f5b5145432d3032363030365d2d5468652061756374696f6e206973206e6f742060408201527f66696e6973686564207965742c206661696c656420746f207472616e73666572606082015268103a3432903637ba1760b91b608082015260a00190565b60208082526043908201527f5b5145432d3032363030375d2d5468652061756374696f6e206973206e6f742060408201527f6163746976652c206661696c656420746f207472616e7366657220746865206c60608201526237ba1760e91b608082015260a00190565b60208082526021908201527f676f7665726e65642e455044522e61756374696f6e4d696e496e6372656d656e6040820152601d60fa1b606082015260800190565b60208082526047908201527f5b5145432d3032363030315d2d5468652073797374656d20646562742069732060408201527f746f6f20736d616c6c2c206661696c656420746f2073746172742074686520616060820152663ab1ba34b7b71760c91b608082015260a00190565b60208082526028908201527f5b5145432d3032363030395d2d5468652061756374696f6e20646f6573206e6f6040820152673a1032bc34b9ba1760c11b606082015260800190565b60208082526064908201527f5b5145432d3032363030355d2d5468652062696420616d6f756e74206d75737460408201527f206578636565642074686520686967686573742062696420627920746865206d60608201527f696e696d756d20696e6372656d656e742070657263656e74616765206f72206d60808201526337b9329760e11b60a082015260c00190565b6020808252601b908201527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b604082015260600190565b60208082526048908201527f5b5145432d3032363030305d2d5468652061756374696f6e20697320616c726560408201527f616479206163746976652c206661696c656420746f207374617274207468652060608201526730bab1ba34b7b71760c11b608082015260a00190565b60208082526018908201527719dbdd995c9b99590b915411148b9c995cd95c9d99531bdd60421b604082015260600190565b60208082526036908201527f5b5145432d3032363030335d2d5468652061756374696f6e206973206e6f742060408201527530b1ba34bb3296103330b4b632b2103a37903134b21760511b606082015260800190565b600060a082019050612529828451611f12565b60018060a01b03602084015116602083015260408301516040830152606083015160608301526080830151608083015292915050565b90815260200190565b6001600160a01b038116811461257d57600080fd5b5056fe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a2646970667358221220d200e945236b5f005d18b3efbd4058ea1cd2735aa2806fd793e068814954365664736f6c63430007060033"

// DeploySystemDebtAuction deploys a new Ethereum contract, binding an instance of SystemDebtAuction to it.
func DeploySystemDebtAuction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemDebtAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemDebtAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SystemDebtAuctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemDebtAuction{SystemDebtAuctionCaller: SystemDebtAuctionCaller{contract: contract}, SystemDebtAuctionTransactor: SystemDebtAuctionTransactor{contract: contract}, SystemDebtAuctionFilterer: SystemDebtAuctionFilterer{contract: contract}}, nil
}

// SystemDebtAuction is an auto generated Go binding around an Ethereum contract.
type SystemDebtAuction struct {
	SystemDebtAuctionCaller     // Read-only binding to the contract
	SystemDebtAuctionTransactor // Write-only binding to the contract
	SystemDebtAuctionFilterer   // Log filterer for contract events
}

// SystemDebtAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemDebtAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemDebtAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemDebtAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemDebtAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemDebtAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemDebtAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemDebtAuctionSession struct {
	Contract     *SystemDebtAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SystemDebtAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemDebtAuctionCallerSession struct {
	Contract *SystemDebtAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SystemDebtAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemDebtAuctionTransactorSession struct {
	Contract     *SystemDebtAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SystemDebtAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemDebtAuctionRaw struct {
	Contract *SystemDebtAuction // Generic contract binding to access the raw methods on
}

// SystemDebtAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemDebtAuctionCallerRaw struct {
	Contract *SystemDebtAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SystemDebtAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemDebtAuctionTransactorRaw struct {
	Contract *SystemDebtAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemDebtAuction creates a new instance of SystemDebtAuction, bound to a specific deployed contract.
func NewSystemDebtAuction(address common.Address, backend bind.ContractBackend) (*SystemDebtAuction, error) {
	contract, err := bindSystemDebtAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuction{SystemDebtAuctionCaller: SystemDebtAuctionCaller{contract: contract}, SystemDebtAuctionTransactor: SystemDebtAuctionTransactor{contract: contract}, SystemDebtAuctionFilterer: SystemDebtAuctionFilterer{contract: contract}}, nil
}

// NewSystemDebtAuctionCaller creates a new read-only instance of SystemDebtAuction, bound to a specific deployed contract.
func NewSystemDebtAuctionCaller(address common.Address, caller bind.ContractCaller) (*SystemDebtAuctionCaller, error) {
	contract, err := bindSystemDebtAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuctionCaller{contract: contract}, nil
}

// NewSystemDebtAuctionTransactor creates a new write-only instance of SystemDebtAuction, bound to a specific deployed contract.
func NewSystemDebtAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemDebtAuctionTransactor, error) {
	contract, err := bindSystemDebtAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuctionTransactor{contract: contract}, nil
}

// NewSystemDebtAuctionFilterer creates a new log filterer instance of SystemDebtAuction, bound to a specific deployed contract.
func NewSystemDebtAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemDebtAuctionFilterer, error) {
	contract, err := bindSystemDebtAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuctionFilterer{contract: contract}, nil
}

// bindSystemDebtAuction binds a generic wrapper to an already deployed contract.
func bindSystemDebtAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemDebtAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemDebtAuction *SystemDebtAuctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemDebtAuction.Contract.SystemDebtAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemDebtAuction *SystemDebtAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.SystemDebtAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemDebtAuction *SystemDebtAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.SystemDebtAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemDebtAuction *SystemDebtAuctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemDebtAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemDebtAuction *SystemDebtAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemDebtAuction *SystemDebtAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.contract.Transact(opts, method, params...)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(uint8 status, address bidder, uint256 highestBid, uint256 endTime, uint256 reserveLot)
func (_SystemDebtAuction *SystemDebtAuctionCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status     uint8
	Bidder     common.Address
	HighestBid *big.Int
	EndTime    *big.Int
	ReserveLot *big.Int
}, error) {
	ret := new(struct {
		Status     uint8
		Bidder     common.Address
		HighestBid *big.Int
		EndTime    *big.Int
		ReserveLot *big.Int
	})
	out := ret
	err := _SystemDebtAuction.contract.Call(opts, out, "auctions", arg0)
	return *ret, err
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(uint8 status, address bidder, uint256 highestBid, uint256 endTime, uint256 reserveLot)
func (_SystemDebtAuction *SystemDebtAuctionSession) Auctions(arg0 *big.Int) (struct {
	Status     uint8
	Bidder     common.Address
	HighestBid *big.Int
	EndTime    *big.Int
	ReserveLot *big.Int
}, error) {
	return _SystemDebtAuction.Contract.Auctions(&_SystemDebtAuction.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(uint8 status, address bidder, uint256 highestBid, uint256 endTime, uint256 reserveLot)
func (_SystemDebtAuction *SystemDebtAuctionCallerSession) Auctions(arg0 *big.Int) (struct {
	Status     uint8
	Bidder     common.Address
	HighestBid *big.Int
	EndTime    *big.Int
	ReserveLot *big.Int
}, error) {
	return _SystemDebtAuction.Contract.Auctions(&_SystemDebtAuction.CallOpts, arg0)
}

// CurrentAuctionId is a free data retrieval call binding the contract method 0x0c0b86ca.
//
// Solidity: function currentAuctionId() view returns(uint256)
func (_SystemDebtAuction *SystemDebtAuctionCaller) CurrentAuctionId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemDebtAuction.contract.Call(opts, out, "currentAuctionId")
	return *ret0, err
}

// CurrentAuctionId is a free data retrieval call binding the contract method 0x0c0b86ca.
//
// Solidity: function currentAuctionId() view returns(uint256)
func (_SystemDebtAuction *SystemDebtAuctionSession) CurrentAuctionId() (*big.Int, error) {
	return _SystemDebtAuction.Contract.CurrentAuctionId(&_SystemDebtAuction.CallOpts)
}

// CurrentAuctionId is a free data retrieval call binding the contract method 0x0c0b86ca.
//
// Solidity: function currentAuctionId() view returns(uint256)
func (_SystemDebtAuction *SystemDebtAuctionCallerSession) CurrentAuctionId() (*big.Int, error) {
	return _SystemDebtAuction.Contract.CurrentAuctionId(&_SystemDebtAuction.CallOpts)
}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemDebtAuction *SystemDebtAuctionCaller) GetRaisingBid(opts *bind.CallOpts, _auctionId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemDebtAuction.contract.Call(opts, out, "getRaisingBid", _auctionId)
	return *ret0, err
}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemDebtAuction *SystemDebtAuctionSession) GetRaisingBid(_auctionId *big.Int) (*big.Int, error) {
	return _SystemDebtAuction.Contract.GetRaisingBid(&_SystemDebtAuction.CallOpts, _auctionId)
}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemDebtAuction *SystemDebtAuctionCallerSession) GetRaisingBid(_auctionId *big.Int) (*big.Int, error) {
	return _SystemDebtAuction.Contract.GetRaisingBid(&_SystemDebtAuction.CallOpts, _auctionId)
}

// HasActiveAuction is a free data retrieval call binding the contract method 0x500ff7e6.
//
// Solidity: function hasActiveAuction() view returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionCaller) HasActiveAuction(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SystemDebtAuction.contract.Call(opts, out, "hasActiveAuction")
	return *ret0, err
}

// HasActiveAuction is a free data retrieval call binding the contract method 0x500ff7e6.
//
// Solidity: function hasActiveAuction() view returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionSession) HasActiveAuction() (bool, error) {
	return _SystemDebtAuction.Contract.HasActiveAuction(&_SystemDebtAuction.CallOpts)
}

// HasActiveAuction is a free data retrieval call binding the contract method 0x500ff7e6.
//
// Solidity: function hasActiveAuction() view returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionCallerSession) HasActiveAuction() (bool, error) {
	return _SystemDebtAuction.Contract.HasActiveAuction(&_SystemDebtAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _bid) returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionTransactor) Bid(opts *bind.TransactOpts, _bid *big.Int) (*types.Transaction, error) {
	return _SystemDebtAuction.contract.Transact(opts, "bid", _bid)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _bid) returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionSession) Bid(_bid *big.Int) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Bid(&_SystemDebtAuction.TransactOpts, _bid)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _bid) returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionTransactorSession) Bid(_bid *big.Int) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Bid(&_SystemDebtAuction.TransactOpts, _bid)
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionTransactor) Execute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemDebtAuction.contract.Transact(opts, "execute")
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionSession) Execute() (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Execute(&_SystemDebtAuction.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionTransactorSession) Execute() (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Execute(&_SystemDebtAuction.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemDebtAuction *SystemDebtAuctionTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemDebtAuction.contract.Transact(opts, "initialize", _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemDebtAuction *SystemDebtAuctionSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Initialize(&_SystemDebtAuction.TransactOpts, _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemDebtAuction *SystemDebtAuctionTransactorSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Initialize(&_SystemDebtAuction.TransactOpts, _registry, _stc)
}

// StartAuction is a paid mutator transaction binding the contract method 0x065de74c.
//
// Solidity: function startAuction(uint256 _bid) returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionTransactor) StartAuction(opts *bind.TransactOpts, _bid *big.Int) (*types.Transaction, error) {
	return _SystemDebtAuction.contract.Transact(opts, "startAuction", _bid)
}

// StartAuction is a paid mutator transaction binding the contract method 0x065de74c.
//
// Solidity: function startAuction(uint256 _bid) returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionSession) StartAuction(_bid *big.Int) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.StartAuction(&_SystemDebtAuction.TransactOpts, _bid)
}

// StartAuction is a paid mutator transaction binding the contract method 0x065de74c.
//
// Solidity: function startAuction(uint256 _bid) returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionTransactorSession) StartAuction(_bid *big.Int) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.StartAuction(&_SystemDebtAuction.TransactOpts, _bid)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemDebtAuction *SystemDebtAuctionTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemDebtAuction.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemDebtAuction *SystemDebtAuctionSession) Receive() (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Receive(&_SystemDebtAuction.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemDebtAuction *SystemDebtAuctionTransactorSession) Receive() (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Receive(&_SystemDebtAuction.TransactOpts)
}

// SystemDebtAuctionAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the SystemDebtAuction contract.
type SystemDebtAuctionAuctionStartedIterator struct {
	Event *SystemDebtAuctionAuctionStarted // Event containing the contract specifics and raw log

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
func (it *SystemDebtAuctionAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemDebtAuctionAuctionStarted)
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
		it.Event = new(SystemDebtAuctionAuctionStarted)
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
func (it *SystemDebtAuctionAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemDebtAuctionAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemDebtAuctionAuctionStarted represents a AuctionStarted event raised by the SystemDebtAuction contract.
type SystemDebtAuctionAuctionStarted struct {
	AuctionId *big.Int
	Bidder    common.Address
	Bid       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0xd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475.
//
// Solidity: event AuctionStarted(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) FilterAuctionStarted(opts *bind.FilterOpts, _auctionId []*big.Int, _bidder []common.Address) (*SystemDebtAuctionAuctionStartedIterator, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemDebtAuction.contract.FilterLogs(opts, "AuctionStarted", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuctionAuctionStartedIterator{contract: _SystemDebtAuction.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0xd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475.
//
// Solidity: event AuctionStarted(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *SystemDebtAuctionAuctionStarted, _auctionId []*big.Int, _bidder []common.Address) (event.Subscription, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemDebtAuction.contract.WatchLogs(opts, "AuctionStarted", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemDebtAuctionAuctionStarted)
				if err := _SystemDebtAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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
// Solidity: event AuctionStarted(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) ParseAuctionStarted(log types.Log) (*SystemDebtAuctionAuctionStarted, error) {
	event := new(SystemDebtAuctionAuctionStarted)
	if err := _SystemDebtAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SystemDebtAuctionBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the SystemDebtAuction contract.
type SystemDebtAuctionBidIterator struct {
	Event *SystemDebtAuctionBid // Event containing the contract specifics and raw log

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
func (it *SystemDebtAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemDebtAuctionBid)
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
		it.Event = new(SystemDebtAuctionBid)
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
func (it *SystemDebtAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemDebtAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemDebtAuctionBid represents a Bid event raised by the SystemDebtAuction contract.
type SystemDebtAuctionBid struct {
	AuctionId *big.Int
	Bidder    common.Address
	Bid       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199.
//
// Solidity: event Bid(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) FilterBid(opts *bind.FilterOpts, _auctionId []*big.Int, _bidder []common.Address) (*SystemDebtAuctionBidIterator, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemDebtAuction.contract.FilterLogs(opts, "Bid", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuctionBidIterator{contract: _SystemDebtAuction.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199.
//
// Solidity: event Bid(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *SystemDebtAuctionBid, _auctionId []*big.Int, _bidder []common.Address) (event.Subscription, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemDebtAuction.contract.WatchLogs(opts, "Bid", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemDebtAuctionBid)
				if err := _SystemDebtAuction.contract.UnpackLog(event, "Bid", log); err != nil {
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
func (_SystemDebtAuction *SystemDebtAuctionFilterer) ParseBid(log types.Log) (*SystemDebtAuctionBid, error) {
	event := new(SystemDebtAuctionBid)
	if err := _SystemDebtAuction.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SystemDebtAuctionExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the SystemDebtAuction contract.
type SystemDebtAuctionExecutedIterator struct {
	Event *SystemDebtAuctionExecuted // Event containing the contract specifics and raw log

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
func (it *SystemDebtAuctionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemDebtAuctionExecuted)
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
		it.Event = new(SystemDebtAuctionExecuted)
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
func (it *SystemDebtAuctionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemDebtAuctionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemDebtAuctionExecuted represents a Executed event raised by the SystemDebtAuction contract.
type SystemDebtAuctionExecuted struct {
	AuctionId *big.Int
	Bidder    common.Address
	Info      SystemDebtAuctionAuctionInfo
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x342522ce238e8878f82d1a020742d0d67f988f004878a3b9842b52b02d278d64.
//
// Solidity: event Executed(uint256 indexed _auctionId, address indexed _bidder, (uint8,address,uint256,uint256,uint256) _info)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) FilterExecuted(opts *bind.FilterOpts, _auctionId []*big.Int, _bidder []common.Address) (*SystemDebtAuctionExecutedIterator, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemDebtAuction.contract.FilterLogs(opts, "Executed", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuctionExecutedIterator{contract: _SystemDebtAuction.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x342522ce238e8878f82d1a020742d0d67f988f004878a3b9842b52b02d278d64.
//
// Solidity: event Executed(uint256 indexed _auctionId, address indexed _bidder, (uint8,address,uint256,uint256,uint256) _info)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *SystemDebtAuctionExecuted, _auctionId []*big.Int, _bidder []common.Address) (event.Subscription, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemDebtAuction.contract.WatchLogs(opts, "Executed", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemDebtAuctionExecuted)
				if err := _SystemDebtAuction.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x342522ce238e8878f82d1a020742d0d67f988f004878a3b9842b52b02d278d64.
//
// Solidity: event Executed(uint256 indexed _auctionId, address indexed _bidder, (uint8,address,uint256,uint256,uint256) _info)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) ParseExecuted(log types.Log) (*SystemDebtAuctionExecuted, error) {
	event := new(SystemDebtAuctionExecuted)
	if err := _SystemDebtAuction.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	return event, nil
}
