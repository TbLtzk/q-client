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

// SystemBalanceSystemBalanceDetails is an auto generated low-level Go binding around an user-defined struct.
type SystemBalanceSystemBalanceDetails struct {
	IsDebtAuctionPossible    bool
	IsSurplusAuctionPossible bool
	CurrentDebt              *big.Int
	DebtThreshold            *big.Int
	CurrentSurplus           *big.Int
	SurplusThreshold         *big.Int
}

// SystemBalanceMetaData contains all meta data concerning the SystemBalance contract.
var SystemBalanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_stc\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtAmount\",\"type\":\"uint256\"}],\"name\":\"increaseDebt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lot\",\"type\":\"uint256\"}],\"name\":\"transferSurplusAuctionAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferAccruedInterestAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performNetting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalanceDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isDebtAuctionPossible\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSurplusAuctionPossible\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"currentDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentSurplus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"surplusThreshold\",\"type\":\"uint256\"}],\"internalType\":\"structSystemBalance.SystemBalanceDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSurplus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506120ee806100206000396000f3fe608060405234801561001057600080fd5b50600436106100835760003560e01c806312065fe01461008857806314a6bf0f146100a6578063152353d9146100ae5780632383b074146100ce5780632b7c7b11146100d65780633d03f797146100e957806364eb4369146100f1578063effd85ad14610106578063f399e22e14610119575b600080fd5b61009061012e565b60405161009d9190611e25565b60405180910390f35b610090610142565b6100c16100bc366004611dbd565b610148565b60405161009d9190611e1a565b6100906102ad565b6100c16100e4366004611dbd565b610337565b6100c16105cb565b6100f961067e565b60405161009d919061202a565b6100c1610114366004611dbd565b610c7c565b61012c610127366004611cf3565b610d93565b005b600060025461013b6102ad565b0390505b90565b60025490565b6000805460018054604080516020600261010085871615026000190190941693909304601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb90271936101f893919290918301828280156101ee5780601f106101c3576101008083540402835291602001916101ee565b820191906000526020600020905b8154815290600101906020018083116101d157829003601f168201915b5050505050610e6d565b6040518263ffffffff1660e01b81526004016102149190611e2e565b60206040518083038186803b15801561022c57600080fd5b505afa158015610240573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102649190611cd7565b6001600160a01b0316336001600160a01b03161461029d5760405162461bcd60e51b815260040161029490611f6f565b60405180910390fd5b6102a73383611026565b92915050565b60006102b76110b6565b6001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016102e29190611ded565b60206040518083038186803b1580156102fa57600080fd5b505afa15801561030e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103329190611dd5565b905090565b6000805460018054604080516020600261010085871615026000190190941693909304601f810184900484028201840190925281815285946201000090046001600160a01b031693633fb90271936103e593908301828280156103db5780601f106103b0576101008083540402835291602001916103db565b820191906000526020600020905b8154815290600101906020018083116103be57829003601f168201915b50505050506111d2565b6040518263ffffffff1660e01b81526004016104019190611e2e565b60206040518083038186803b15801561041957600080fd5b505afa15801561042d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104519190611cd7565b6000805460018054604080516020600261010085871615026000190190941693909304601f81018490048402820184019092528181526001600160a01b039687163314975094956201000090940490931693633fb902719361050a9391928301828280156105005780601f106104d557610100808354040283529160200191610500565b820191906000526020600020905b8154815290600101906020018083116104e357829003601f168201915b5050505050611389565b6040518263ffffffff1660e01b81526004016105269190611e2e565b60206040518083038186803b15801561053e57600080fd5b505afa158015610552573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105769190611cd7565b6001600160a01b0316336001600160a01b031614905081806105955750805b6105b15760405162461bcd60e51b815260040161029490611eec565b6002546105be9085611534565b6002555060019392505050565b6000806105d66102ad565b905060006002548211156105ed57506002546105f0565b50805b806106005760009250505061013f565b60025461060d908261158c565b6002556106186110b6565b6001600160a01b03166342966c68826040518263ffffffff1660e01b81526004016106439190611e25565b600060405180830381600087803b15801561065d57600080fd5b505af1158015610671573d6000803e3d6000fd5b5050505060019250505090565b610686611bfc565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906106ba90600401611fe8565b60206040518083038186803b1580156106d257600080fd5b505afa1580156106e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070a9190611cd7565b9050610714611bfc565b60025460408201526107246102ad565b608082015260018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281526001600160a01b0386169363498bff00936107cc938301828280156107c25780601f10610797576101008083540402835291602001916107c2565b820191906000526020600020905b8154815290600101906020018083116107a557829003601f168201915b50505050506115ce565b6040518263ffffffff1660e01b81526004016107e89190611e2e565b60206040518083038186803b15801561080057600080fd5b505afa158015610814573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108389190611dd5565b606082015260018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281526001600160a01b0386169363498bff00936108e0938301828280156108d65780601f106108ab576101008083540402835291602001916108d6565b820191906000526020600020905b8154815290600101906020018083116108b957829003601f168201915b50505050506115ff565b6040518263ffffffff1660e01b81526004016108fc9190611e2e565b60206040518083038186803b15801561091457600080fd5b505afa158015610928573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061094c9190611dd5565b60a08201526080810151604082015110610b2457600061097d8260800151836040015161158c90919063ffffffff16565b9050816060015181118015610b1a575060005460018054604080516020600261010085871615026000190190941693909304601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb9027193610a3c9391929091830182828015610a325780601f10610a0757610100808354040283529160200191610a32565b820191906000526020600020905b815481529060010190602001808311610a1557829003601f168201915b5050505050611633565b6040518263ffffffff1660e01b8152600401610a589190611e2e565b60206040518083038186803b158015610a7057600080fd5b505afa158015610a84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa89190611cd7565b6001600160a01b031663500ff7e66040518163ffffffff1660e01b815260040160206040518083038186803b158015610ae057600080fd5b505afa158015610af4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b189190611d9d565b155b1515825250610c76565b6000610b418260400151836080015161158c90919063ffffffff16565b60018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281529394506000936001600160a01b0388169363498bff0093610bea93830182828015610be05780601f10610bb557610100808354040283529160200191610be0565b820191906000526020600020905b815481529060010190602001808311610bc357829003601f168201915b50505050506117e9565b6040518263ffffffff1660e01b8152600401610c069190611e2e565b60206040518083038186803b158015610c1e57600080fd5b505afa158015610c32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c569190611dd5565b90508260a001518210158015610c6c5750808210155b1515602084015250505b91505090565b6000805460018054604080516020600261010085871615026000190190941693909304601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb9027193610cf793919290918301828280156105005780601f106104d557610100808354040283529160200191610500565b6040518263ffffffff1660e01b8152600401610d139190611e2e565b60206040518083038186803b158015610d2b57600080fd5b505afa158015610d3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d639190611cd7565b6001600160a01b0316336001600160a01b03161461029d5760405162461bcd60e51b815260040161029490611e81565b600054610100900460ff1680610dac5750610dac611817565b80610dba575060005460ff16155b610df55760405162461bcd60e51b815260040180806020018281038252602e81526020018061208b602e913960400191505060405180910390fd5b600054610100900460ff16158015610e20576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038616021790558151610e55906001906020850190611c36565b508015610e68576000805461ff00191690555b505050565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b60208310610ef35780518252601f199092019160209182019101610ed4565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b60208310610f3b5780518252601f199092019160209182019101610f1c565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310610f835780518252601f199092019160209182019101610f64565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310610fcb5780518252601f199092019160209182019101610fac565b5181516020939093036101000a60001901801990911692169190911790527339bcb9ba32b6a9bab938363ab9a0bab1ba34b7b760611b92019182525060408051808303600b1901815260149092019052979650505050505050565b60006110306110b6565b6001600160a01b031663a9059cbb84846040518363ffffffff1660e01b815260040161105d929190611e01565b602060405180830381600087803b15801561107757600080fd5b505af115801561108b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110af9190611d9d565b9392505050565b6000805460018054604080516020600261010085871615026000190190941693909304601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb9027193611166939192909183018282801561115c5780601f106111315761010080835404028352916020019161115c565b820191906000526020600020905b81548152906001019060200180831161113f57829003601f168201915b5050505050611828565b6040518263ffffffff1660e01b81526004016111829190611e2e565b60206040518083038186803b15801561119a57600080fd5b505afa1580156111ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103329190611cd7565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106112585780518252601f199092019160209182019101611239565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106112a05780518252601f199092019160209182019101611281565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106112e85780518252601f1990920191602091820191016112c9565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106113305780518252601f199092019160209182019101611311565b5181516020939093036101000a6000190180199091169216919091179052713634b8bab4b230ba34b7b720bab1ba34b7b760711b92019182525060408051808303600d1901815260129092019052979650505050505050565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b6020831061140f5780518252601f1990920191602091820191016113f0565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106114575780518252601f199092019160209182019101611438565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b6020831061149f5780518252601f199092019160209182019101611480565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106114e75780518252601f1990920191602091820191016114c8565b5181516020939093036101000a600019018019909116921691909117905265736176696e6760d01b9201918252506040805180830360191901815260069092019052979650505050505050565b6000828201838110156110af576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b60006110af83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506119d1565b60606102a7826040518060400160405280600d81526020016c1919589d151a1c995cda1bdb19609a1b815250611a68565b60606102a7826040518060400160405280601081526020016f1cdd5c9c1b1d5cd51a1c995cda1bdb1960821b815250611a68565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106116b95780518252601f19909201916020918201910161169a565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106117015780518252601f1990920191602091820191016116e2565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106117495780518252601f19909201916020918201910161172a565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106117915780518252601f199092019160209182019101611772565b5181516020939093036101000a60001901801990911692169190911790527039bcb9ba32b6a232b13a20bab1ba34b7b760791b92019182525060408051808303600e1901815260119092019052979650505050505050565b60606102a7826040518060400160405280600a8152602001691cdd5c9c1b1d5cd31bdd60b21b815250611a68565b600061182230611bf6565b15905090565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106118ae5780518252601f19909201916020918201910161188f565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106118f65780518252601f1990920191602091820191016118d7565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b6020831061193e5780518252601f19909201916020918201910161191f565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106119865780518252601f199092019160209182019101611967565b5181516020939093036101000a60001901801990911692169190911790526331b7b4b760e11b92019182525060408051808303601b1901815260049092019052979650505050505050565b60008184841115611a605760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611a25578181015183820152602001611a0d565b50505050905090810190601f168015611a525780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60606040518060400160405280600e81526020016d33b7bb32b93732b21722a822291760911b81525083604051806040016040528060018152602001605f60f81b815250846040516020018085805190602001908083835b60208310611adf5780518252601f199092019160209182019101611ac0565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b60208310611b275780518252601f199092019160209182019101611b08565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310611b6f5780518252601f199092019160209182019101611b50565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310611bb75780518252601f199092019160209182019101611b98565b6001836020036101000a038019825116818451168082178552505050505050905001945050505050604051602081830303815290604052905092915050565b3b151590565b6040518060c00160405280600015158152602001600015158152602001600081526020016000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282611c6c5760008555611cb2565b82601f10611c8557805160ff1916838001178555611cb2565b82800160010185558215611cb2579182015b82811115611cb2578251825591602001919060010190611c97565b50611cbe929150611cc2565b5090565b5b80821115611cbe5760008155600101611cc3565b600060208284031215611ce8578081fd5b81516110af81612072565b60008060408385031215611d05578081fd5b8235611d1081612072565b915060208381013567ffffffffffffffff80821115611d2d578384fd5b818601915086601f830112611d40578384fd5b813581811115611d4c57fe5b604051601f8201601f1916810185018381118282101715611d6957fe5b6040528181528382018501891015611d7f578586fd5b81858501868301378585838301015280955050505050509250929050565b600060208284031215611dae578081fd5b815180151581146110af578182fd5b600060208284031215611dce578081fd5b5035919050565b600060208284031215611de6578081fd5b5051919050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b90815260200190565b6000602080835283518082850152825b81811015611e5a57858101830151858201604001528201611e3e565b81811115611e6b5783604083870101525b50601f01601f1916929092016040019392505050565b60208082526045908201527f5b5145432d3032333030315d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c792074686520536176696e6720636f6e74726163742068617320616360608201526431b2b9b99760d91b608082015260a00190565b6020808252605f908201527f5b5145432d3032333030325d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c7920746865204c69717569646174696f6e2061756374696f6e20616e60608201527f6420536176696e6720636f6e7472616374732068617665206163636573732e00608082015260a00190565b60208082526053908201527f5b5145432d3032333030305d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c79207468652053797374656d537572706c757341756374696f6e206360608201527237b73a3930b1ba103430b99030b1b1b2b9b99760691b608082015260a00190565b60208082526022908201527f676f7665726e616e63652e657870657274732e455044522e706172616d657465604082015261727360f01b606082015260800190565b600060c08201905082511515825260208301511515602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015292915050565b6001600160a01b038116811461208757600080fd5b5056fe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a2646970667358221220f030dc5ab4cd13eca64d96f5b16740e1af6b4f25675d505404a3f97717eabc6864736f6c63430007060033",
}

// SystemBalanceABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemBalanceMetaData.ABI instead.
var SystemBalanceABI = SystemBalanceMetaData.ABI

// SystemBalanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemBalanceMetaData.Bin instead.
var SystemBalanceBin = SystemBalanceMetaData.Bin

// DeploySystemBalance deploys a new Ethereum contract, binding an instance of SystemBalance to it.
func DeploySystemBalance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemBalance, error) {
	parsed, err := SystemBalanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemBalanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemBalance{SystemBalanceCaller: SystemBalanceCaller{contract: contract}, SystemBalanceTransactor: SystemBalanceTransactor{contract: contract}, SystemBalanceFilterer: SystemBalanceFilterer{contract: contract}}, nil
}

// SystemBalance is an auto generated Go binding around an Ethereum contract.
type SystemBalance struct {
	SystemBalanceCaller     // Read-only binding to the contract
	SystemBalanceTransactor // Write-only binding to the contract
	SystemBalanceFilterer   // Log filterer for contract events
}

// SystemBalanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemBalanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemBalanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemBalanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemBalanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemBalanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemBalanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemBalanceSession struct {
	Contract     *SystemBalance    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemBalanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemBalanceCallerSession struct {
	Contract *SystemBalanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SystemBalanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemBalanceTransactorSession struct {
	Contract     *SystemBalanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SystemBalanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemBalanceRaw struct {
	Contract *SystemBalance // Generic contract binding to access the raw methods on
}

// SystemBalanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemBalanceCallerRaw struct {
	Contract *SystemBalanceCaller // Generic read-only contract binding to access the raw methods on
}

// SystemBalanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemBalanceTransactorRaw struct {
	Contract *SystemBalanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemBalance creates a new instance of SystemBalance, bound to a specific deployed contract.
func NewSystemBalance(address common.Address, backend bind.ContractBackend) (*SystemBalance, error) {
	contract, err := bindSystemBalance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemBalance{SystemBalanceCaller: SystemBalanceCaller{contract: contract}, SystemBalanceTransactor: SystemBalanceTransactor{contract: contract}, SystemBalanceFilterer: SystemBalanceFilterer{contract: contract}}, nil
}

// NewSystemBalanceCaller creates a new read-only instance of SystemBalance, bound to a specific deployed contract.
func NewSystemBalanceCaller(address common.Address, caller bind.ContractCaller) (*SystemBalanceCaller, error) {
	contract, err := bindSystemBalance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemBalanceCaller{contract: contract}, nil
}

// NewSystemBalanceTransactor creates a new write-only instance of SystemBalance, bound to a specific deployed contract.
func NewSystemBalanceTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemBalanceTransactor, error) {
	contract, err := bindSystemBalance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemBalanceTransactor{contract: contract}, nil
}

// NewSystemBalanceFilterer creates a new log filterer instance of SystemBalance, bound to a specific deployed contract.
func NewSystemBalanceFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemBalanceFilterer, error) {
	contract, err := bindSystemBalance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemBalanceFilterer{contract: contract}, nil
}

// bindSystemBalance binds a generic wrapper to an already deployed contract.
func bindSystemBalance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemBalanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemBalance *SystemBalanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemBalance.Contract.SystemBalanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemBalance *SystemBalanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemBalance.Contract.SystemBalanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemBalance *SystemBalanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemBalance.Contract.SystemBalanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemBalance *SystemBalanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemBalance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemBalance *SystemBalanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemBalance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemBalance *SystemBalanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemBalance.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(int256)
func (_SystemBalance *SystemBalanceCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemBalance.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(int256)
func (_SystemBalance *SystemBalanceSession) GetBalance() (*big.Int, error) {
	return _SystemBalance.Contract.GetBalance(&_SystemBalance.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(int256)
func (_SystemBalance *SystemBalanceCallerSession) GetBalance() (*big.Int, error) {
	return _SystemBalance.Contract.GetBalance(&_SystemBalance.CallOpts)
}

// GetBalanceDetails is a free data retrieval call binding the contract method 0x64eb4369.
//
// Solidity: function getBalanceDetails() view returns((bool,bool,uint256,uint256,uint256,uint256))
func (_SystemBalance *SystemBalanceCaller) GetBalanceDetails(opts *bind.CallOpts) (SystemBalanceSystemBalanceDetails, error) {
	var out []interface{}
	err := _SystemBalance.contract.Call(opts, &out, "getBalanceDetails")

	if err != nil {
		return *new(SystemBalanceSystemBalanceDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(SystemBalanceSystemBalanceDetails)).(*SystemBalanceSystemBalanceDetails)

	return out0, err

}

// GetBalanceDetails is a free data retrieval call binding the contract method 0x64eb4369.
//
// Solidity: function getBalanceDetails() view returns((bool,bool,uint256,uint256,uint256,uint256))
func (_SystemBalance *SystemBalanceSession) GetBalanceDetails() (SystemBalanceSystemBalanceDetails, error) {
	return _SystemBalance.Contract.GetBalanceDetails(&_SystemBalance.CallOpts)
}

// GetBalanceDetails is a free data retrieval call binding the contract method 0x64eb4369.
//
// Solidity: function getBalanceDetails() view returns((bool,bool,uint256,uint256,uint256,uint256))
func (_SystemBalance *SystemBalanceCallerSession) GetBalanceDetails() (SystemBalanceSystemBalanceDetails, error) {
	return _SystemBalance.Contract.GetBalanceDetails(&_SystemBalance.CallOpts)
}

// GetDebt is a free data retrieval call binding the contract method 0x14a6bf0f.
//
// Solidity: function getDebt() view returns(uint256)
func (_SystemBalance *SystemBalanceCaller) GetDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemBalance.contract.Call(opts, &out, "getDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDebt is a free data retrieval call binding the contract method 0x14a6bf0f.
//
// Solidity: function getDebt() view returns(uint256)
func (_SystemBalance *SystemBalanceSession) GetDebt() (*big.Int, error) {
	return _SystemBalance.Contract.GetDebt(&_SystemBalance.CallOpts)
}

// GetDebt is a free data retrieval call binding the contract method 0x14a6bf0f.
//
// Solidity: function getDebt() view returns(uint256)
func (_SystemBalance *SystemBalanceCallerSession) GetDebt() (*big.Int, error) {
	return _SystemBalance.Contract.GetDebt(&_SystemBalance.CallOpts)
}

// GetSurplus is a free data retrieval call binding the contract method 0x2383b074.
//
// Solidity: function getSurplus() view returns(uint256)
func (_SystemBalance *SystemBalanceCaller) GetSurplus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemBalance.contract.Call(opts, &out, "getSurplus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSurplus is a free data retrieval call binding the contract method 0x2383b074.
//
// Solidity: function getSurplus() view returns(uint256)
func (_SystemBalance *SystemBalanceSession) GetSurplus() (*big.Int, error) {
	return _SystemBalance.Contract.GetSurplus(&_SystemBalance.CallOpts)
}

// GetSurplus is a free data retrieval call binding the contract method 0x2383b074.
//
// Solidity: function getSurplus() view returns(uint256)
func (_SystemBalance *SystemBalanceCallerSession) GetSurplus() (*big.Int, error) {
	return _SystemBalance.Contract.GetSurplus(&_SystemBalance.CallOpts)
}

// IncreaseDebt is a paid mutator transaction binding the contract method 0x2b7c7b11.
//
// Solidity: function increaseDebt(uint256 _debtAmount) returns(bool)
func (_SystemBalance *SystemBalanceTransactor) IncreaseDebt(opts *bind.TransactOpts, _debtAmount *big.Int) (*types.Transaction, error) {
	return _SystemBalance.contract.Transact(opts, "increaseDebt", _debtAmount)
}

// IncreaseDebt is a paid mutator transaction binding the contract method 0x2b7c7b11.
//
// Solidity: function increaseDebt(uint256 _debtAmount) returns(bool)
func (_SystemBalance *SystemBalanceSession) IncreaseDebt(_debtAmount *big.Int) (*types.Transaction, error) {
	return _SystemBalance.Contract.IncreaseDebt(&_SystemBalance.TransactOpts, _debtAmount)
}

// IncreaseDebt is a paid mutator transaction binding the contract method 0x2b7c7b11.
//
// Solidity: function increaseDebt(uint256 _debtAmount) returns(bool)
func (_SystemBalance *SystemBalanceTransactorSession) IncreaseDebt(_debtAmount *big.Int) (*types.Transaction, error) {
	return _SystemBalance.Contract.IncreaseDebt(&_SystemBalance.TransactOpts, _debtAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemBalance *SystemBalanceTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemBalance.contract.Transact(opts, "initialize", _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemBalance *SystemBalanceSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemBalance.Contract.Initialize(&_SystemBalance.TransactOpts, _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemBalance *SystemBalanceTransactorSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemBalance.Contract.Initialize(&_SystemBalance.TransactOpts, _registry, _stc)
}

// PerformNetting is a paid mutator transaction binding the contract method 0x3d03f797.
//
// Solidity: function performNetting() returns(bool)
func (_SystemBalance *SystemBalanceTransactor) PerformNetting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemBalance.contract.Transact(opts, "performNetting")
}

// PerformNetting is a paid mutator transaction binding the contract method 0x3d03f797.
//
// Solidity: function performNetting() returns(bool)
func (_SystemBalance *SystemBalanceSession) PerformNetting() (*types.Transaction, error) {
	return _SystemBalance.Contract.PerformNetting(&_SystemBalance.TransactOpts)
}

// PerformNetting is a paid mutator transaction binding the contract method 0x3d03f797.
//
// Solidity: function performNetting() returns(bool)
func (_SystemBalance *SystemBalanceTransactorSession) PerformNetting() (*types.Transaction, error) {
	return _SystemBalance.Contract.PerformNetting(&_SystemBalance.TransactOpts)
}

// TransferAccruedInterestAmount is a paid mutator transaction binding the contract method 0xeffd85ad.
//
// Solidity: function transferAccruedInterestAmount(uint256 _amount) returns(bool)
func (_SystemBalance *SystemBalanceTransactor) TransferAccruedInterestAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SystemBalance.contract.Transact(opts, "transferAccruedInterestAmount", _amount)
}

// TransferAccruedInterestAmount is a paid mutator transaction binding the contract method 0xeffd85ad.
//
// Solidity: function transferAccruedInterestAmount(uint256 _amount) returns(bool)
func (_SystemBalance *SystemBalanceSession) TransferAccruedInterestAmount(_amount *big.Int) (*types.Transaction, error) {
	return _SystemBalance.Contract.TransferAccruedInterestAmount(&_SystemBalance.TransactOpts, _amount)
}

// TransferAccruedInterestAmount is a paid mutator transaction binding the contract method 0xeffd85ad.
//
// Solidity: function transferAccruedInterestAmount(uint256 _amount) returns(bool)
func (_SystemBalance *SystemBalanceTransactorSession) TransferAccruedInterestAmount(_amount *big.Int) (*types.Transaction, error) {
	return _SystemBalance.Contract.TransferAccruedInterestAmount(&_SystemBalance.TransactOpts, _amount)
}

// TransferSurplusAuctionAmount is a paid mutator transaction binding the contract method 0x152353d9.
//
// Solidity: function transferSurplusAuctionAmount(uint256 _lot) returns(bool)
func (_SystemBalance *SystemBalanceTransactor) TransferSurplusAuctionAmount(opts *bind.TransactOpts, _lot *big.Int) (*types.Transaction, error) {
	return _SystemBalance.contract.Transact(opts, "transferSurplusAuctionAmount", _lot)
}

// TransferSurplusAuctionAmount is a paid mutator transaction binding the contract method 0x152353d9.
//
// Solidity: function transferSurplusAuctionAmount(uint256 _lot) returns(bool)
func (_SystemBalance *SystemBalanceSession) TransferSurplusAuctionAmount(_lot *big.Int) (*types.Transaction, error) {
	return _SystemBalance.Contract.TransferSurplusAuctionAmount(&_SystemBalance.TransactOpts, _lot)
}

// TransferSurplusAuctionAmount is a paid mutator transaction binding the contract method 0x152353d9.
//
// Solidity: function transferSurplusAuctionAmount(uint256 _lot) returns(bool)
func (_SystemBalance *SystemBalanceTransactorSession) TransferSurplusAuctionAmount(_lot *big.Int) (*types.Transaction, error) {
	return _SystemBalance.Contract.TransferSurplusAuctionAmount(&_SystemBalance.TransactOpts, _lot)
}
