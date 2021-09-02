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
	CurrentDebt              *big.Int
	DebtThreshold            *big.Int
	IsDebtAuctionPossible    bool
	CurrentSurplus           *big.Int
	SurplusThreshold         *big.Int
	IsSurplusAuctionPossible bool
}

// SystemBalanceMetaData contains all meta data concerning the SystemBalance contract.
var SystemBalanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_stc\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_eligibleContractKeys\",\"type\":\"string[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtAmount\",\"type\":\"uint256\"}],\"name\":\"increaseDebt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lot\",\"type\":\"uint256\"}],\"name\":\"transferSurplusAuctionAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferAccruedInterestAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performNetting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalanceDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isDebtAuctionPossible\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"currentSurplus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"surplusThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSurplusAuctionPossible\",\"type\":\"bool\"}],\"internalType\":\"structSystemBalance.SystemBalanceDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSurplus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611fb7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100835760003560e01c806312065fe01461008857806314a6bf0f146100a6578063152353d9146100ae5780632383b074146100ce5780632b7c7b11146100d65780633d03f797146100e957806364eb4369146100f1578063943d552e14610106578063effd85ad1461011b575b600080fd5b61009061012e565b60405161009d9190611c4b565b60405180910390f35b610090610142565b6100c16100bc366004611be3565b610148565b60405161009d9190611c40565b6100906102ab565b6100c16100e4366004611be3565b610335565b6100c161043b565b6100f96104ee565b60405161009d9190611ec6565b610119610114366004611af3565b610adc565b005b6100c1610129366004611be3565b610bcb565b600060035461013b6102ab565b0390505b90565b60035490565b600080546002805460408051602061010060018516150260001901909316849004601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb90271936101f693919290918301828280156101ec5780601f106101c1576101008083540402835291602001916101ec565b820191906000526020600020905b8154815290600101906020018083116101cf57829003601f168201915b5050505050610d15565b6040518263ffffffff1660e01b81526004016102129190611c54565b60206040518083038186803b15801561022a57600080fd5b505afa15801561023e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102629190611ad7565b6001600160a01b0316336001600160a01b03161461029b5760405162461bcd60e51b815260040161029290611d9c565b60405180910390fd5b6102a53383610ece565b92915050565b60006102b5610f5e565b6001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016102e09190611c13565b60206040518083038186803b1580156102f857600080fd5b505afa15801561030c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103309190611bfb565b905090565b60008060005b600154811015610404576000546001805433926201000090046001600160a01b031691633fb90271918590811061036e57fe5b906000526020600020016040518263ffffffff1660e01b81526004016103949190611ca7565b60206040518083038186803b1580156103ac57600080fd5b505afa1580156103c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e49190611ad7565b6001600160a01b031614156103fc5760019150610404565b60010161033b565b50806104225760405162461bcd60e51b815260040161029290611e57565b60035461042f9084611078565b60035550600192915050565b6000806104466102ab565b9050600060035482111561045d5750600354610460565b50805b806104705760009250505061013f565b60035461047d90826110d0565b600355610488610f5e565b6001600160a01b03166342966c68826040518263ffffffff1660e01b81526004016104b39190611c4b565b600060405180830381600087803b1580156104cd57600080fd5b505af11580156104e1573d6000803e3d6000fd5b5050505060019250505090565b6104f66118da565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061052a90600401611e15565b60206040518083038186803b15801561054257600080fd5b505afa158015610556573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057a9190611ad7565b90506105846118da565b60035481526105916102ab565b606082015260028054604080516020601f600019610100600187161502019094168590049384018190048102820181019092528281526001600160a01b0386169363498bff00936106379383018282801561062d5780601f106106025761010080835404028352916020019161062d565b820191906000526020600020905b81548152906001019060200180831161061057829003601f168201915b5050505050611112565b6040518263ffffffff1660e01b81526004016106539190611c54565b60206040518083038186803b15801561066b57600080fd5b505afa15801561067f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a39190611bfb565b6020808301919091526002805460408051601f600019610100600186161502019093168490049283018590048502810185019091528181526001600160a01b0386169363498bff009361074d9390918301828280156107435780601f1061071857610100808354040283529160200191610743565b820191906000526020600020905b81548152906001019060200180831161072657829003601f168201915b5050505050611143565b6040518263ffffffff1660e01b81526004016107699190611c54565b60206040518083038186803b15801561078157600080fd5b505afa158015610795573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107b99190611bfb565b6080820152606081015181511061098357606081015181516000916107de91906110d0565b90508160200151811015801561097a57506000546002805460408051602061010060018516150260001901909316849004601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb902719361089c93919290918301828280156108925780601f1061086757610100808354040283529160200191610892565b820191906000526020600020905b81548152906001019060200180831161087557829003601f168201915b5050505050611177565b6040518263ffffffff1660e01b81526004016108b89190611c54565b60206040518083038186803b1580156108d057600080fd5b505afa1580156108e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109089190611ad7565b6001600160a01b031663500ff7e66040518163ffffffff1660e01b815260040160206040518083038186803b15801561094057600080fd5b505afa158015610954573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109789190611bc3565b155b15156040830152505b8051606082015110610ad657805160608201516000916109a391906110d0565b60028054604080516020601f600019610100600187161502019094168590049384018190048102820181019092528281529394506000936001600160a01b0388169363498bff0093610a4a93830182828015610a405780601f10610a1557610100808354040283529160200191610a40565b820191906000526020600020905b815481529060010190602001808311610a2357829003601f168201915b505050505061132d565b6040518263ffffffff1660e01b8152600401610a669190611c54565b60206040518083038186803b158015610a7e57600080fd5b505afa158015610a92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab69190611bfb565b905082608001518210158015610acc5750808210155b151560a084015250505b91505090565b600054610100900460ff1680610af55750610af561135b565b80610b03575060005460ff16155b610b3e5760405162461bcd60e51b815260040180806020018281038252602e815260200180611f54602e913960400191505060405180910390fd5b600054610100900460ff16158015610b69576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038716021790558151610b9e906001906020850190611914565b508251610bb2906002906020860190611971565b508015610bc5576000805461ff00191690555b50505050565b600080546002805460408051602061010060018516150260001901909316849004601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb9027193610c799391929091830182828015610c6f5780601f10610c4457610100808354040283529160200191610c6f565b820191906000526020600020905b815481529060010190602001808311610c5257829003601f168201915b5050505050611361565b6040518263ffffffff1660e01b8152600401610c959190611c54565b60206040518083038186803b158015610cad57600080fd5b505afa158015610cc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce59190611ad7565b6001600160a01b0316336001600160a01b03161461029b5760405162461bcd60e51b815260040161029290611d31565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b60208310610d9b5780518252601f199092019160209182019101610d7c565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b60208310610de35780518252601f199092019160209182019101610dc4565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310610e2b5780518252601f199092019160209182019101610e0c565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310610e735780518252601f199092019160209182019101610e54565b5181516020939093036101000a60001901801990911692169190911790527339bcb9ba32b6a9bab938363ab9a0bab1ba34b7b760611b92019182525060408051808303600b1901815260149092019052979650505050505050565b6000610ed8610f5e565b6001600160a01b031663a9059cbb84846040518363ffffffff1660e01b8152600401610f05929190611c27565b602060405180830381600087803b158015610f1f57600080fd5b505af1158015610f33573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f579190611bc3565b9392505050565b600080546002805460408051602061010060018516150260001901909316849004601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb902719361100c93919290918301828280156110025780601f10610fd757610100808354040283529160200191611002565b820191906000526020600020905b815481529060010190602001808311610fe557829003601f168201915b505050505061150c565b6040518263ffffffff1660e01b81526004016110289190611c54565b60206040518083038186803b15801561104057600080fd5b505afa158015611054573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103309190611ad7565b600082820183811015610f57576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b6000610f5783836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506116b5565b60606102a5826040518060400160405280600d81526020016c1919589d151a1c995cda1bdb19609a1b81525061174c565b60606102a5826040518060400160405280601081526020016f1cdd5c9c1b1d5cd51a1c995cda1bdb1960821b81525061174c565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106111fd5780518252601f1990920191602091820191016111de565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106112455780518252601f199092019160209182019101611226565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b6020831061128d5780518252601f19909201916020918201910161126e565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106112d55780518252601f1990920191602091820191016112b6565b5181516020939093036101000a60001901801990911692169190911790527039bcb9ba32b6a232b13a20bab1ba34b7b760791b92019182525060408051808303600e1901815260119092019052979650505050505050565b60606102a5826040518060400160405280600a8152602001691cdd5c9c1b1d5cd31bdd60b21b81525061174c565b303b1590565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106113e75780518252601f1990920191602091820191016113c8565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b6020831061142f5780518252601f199092019160209182019101611410565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106114775780518252601f199092019160209182019101611458565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106114bf5780518252601f1990920191602091820191016114a0565b5181516020939093036101000a600019018019909116921691909117905265736176696e6760d01b9201918252506040805180830360191901815260069092019052979650505050505050565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106115925780518252601f199092019160209182019101611573565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106115da5780518252601f1990920191602091820191016115bb565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106116225780518252601f199092019160209182019101611603565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061166a5780518252601f19909201916020918201910161164b565b5181516020939093036101000a60001901801990911692169190911790526331b7b4b760e11b92019182525060408051808303601b1901815260049092019052979650505050505050565b600081848411156117445760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156117095781810151838201526020016116f1565b50505050905090810190601f1680156117365780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60606040518060400160405280600e81526020016d33b7bb32b93732b21722a822291760911b81525083604051806040016040528060018152602001605f60f81b815250846040516020018085805190602001908083835b602083106117c35780518252601f1990920191602091820191016117a4565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b6020831061180b5780518252601f1990920191602091820191016117ec565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106118535780518252601f199092019160209182019101611834565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061189b5780518252601f19909201916020918201910161187c565b6001836020036101000a038019825116818451168082178552505050505050905001945050505050604051602081830303815290604052905092915050565b6040518060c00160405280600081526020016000815260200160001515815260200160008152602001600081526020016000151581525090565b828054828255906000526020600020908101928215611961579160200282015b828111156119615782518051611951918491602090910190611971565b5091602001919060010190611934565b5061196d9291506119f9565b5090565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826119a757600085556119ed565b82601f106119c057805160ff19168380011785556119ed565b828001600101855582156119ed579182015b828111156119ed5782518255916020019190600101906119d2565b5061196d929150611a16565b8082111561196d576000611a0d8282611a2b565b506001016119f9565b5b8082111561196d5760008155600101611a17565b50805460018160011615610100020316600290046000825580601f10611a515750611a6f565b601f016020900490600052602060002090810190611a6f9190611a16565b50565b600082601f830112611a82578081fd5b813567ffffffffffffffff811115611a9657fe5b611aa9601f8201601f1916602001611f0e565b818152846020838601011115611abd578283fd5b816020850160208301379081016020019190915292915050565b600060208284031215611ae8578081fd5b8151610f5781611f3e565b600080600060608486031215611b07578182fd5b8335611b1281611f3e565b925060208481013567ffffffffffffffff80821115611b2f578485fd5b611b3b88838901611a72565b94506040870135915080821115611b50578384fd5b818701915087601f830112611b63578384fd5b813581811115611b6f57fe5b611b7c8485830201611f0e565b8181528481019250838501865b83811015611bb257611ba08c888435890101611a72565b85529386019390860190600101611b89565b505080955050505050509250925092565b600060208284031215611bd4578081fd5b81518015158114610f57578182fd5b600060208284031215611bf4578081fd5b5035919050565b600060208284031215611c0c578081fd5b5051919050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b90815260200190565b6000602080835283518082850152825b81811015611c8057858101830151858201604001528201611c64565b81811115611c915783604083870101525b50601f01601f1916929092016040019392505050565b60006020808301818452828554600180821660008114611cce5760018114611cec57611d24565b60028304607f16855260ff1983166040890152606088019350611d24565b60028304808652611cfc8a611f32565b885b82811015611d1a5781548b820160400152908401908801611cfe565b8a01604001955050505b5091979650505050505050565b60208082526045908201527f5b5145432d3032333030315d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c792074686520536176696e6720636f6e74726163742068617320616360608201526431b2b9b99760d91b608082015260a00190565b60208082526053908201527f5b5145432d3032333030305d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c79207468652053797374656d537572706c757341756374696f6e206360608201527237b73a3930b1ba103430b99030b1b1b2b9b99760691b608082015260a00190565b60208082526022908201527f676f7665726e616e63652e657870657274732e455044522e706172616d657465604082015261727360f01b606082015260800190565b60208082526049908201527f5b5145432d3032333030325d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c792074686520656c696769626c6520636f6e74726163747320686176606082015268329030b1b1b2b9b99760b91b608082015260a00190565b600060c0820190508251825260208301516020830152604083015115156040830152606083015160608301526080830151608083015260a0830151151560a083015292915050565b60405181810167ffffffffffffffff81118282101715611f2a57fe5b604052919050565b60009081526020902090565b6001600160a01b0381168114611a6f57600080fdfe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a26469706673582212208307fa0c3fdfcc60f075a460f3acba052cc8e3eb1cf5d08c8d6def7795d6f5d464736f6c63430007060033",
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
// Solidity: function getBalanceDetails() view returns((uint256,uint256,bool,uint256,uint256,bool))
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
// Solidity: function getBalanceDetails() view returns((uint256,uint256,bool,uint256,uint256,bool))
func (_SystemBalance *SystemBalanceSession) GetBalanceDetails() (SystemBalanceSystemBalanceDetails, error) {
	return _SystemBalance.Contract.GetBalanceDetails(&_SystemBalance.CallOpts)
}

// GetBalanceDetails is a free data retrieval call binding the contract method 0x64eb4369.
//
// Solidity: function getBalanceDetails() view returns((uint256,uint256,bool,uint256,uint256,bool))
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

// Initialize is a paid mutator transaction binding the contract method 0x943d552e.
//
// Solidity: function initialize(address _registry, string _stc, string[] _eligibleContractKeys) returns()
func (_SystemBalance *SystemBalanceTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _stc string, _eligibleContractKeys []string) (*types.Transaction, error) {
	return _SystemBalance.contract.Transact(opts, "initialize", _registry, _stc, _eligibleContractKeys)
}

// Initialize is a paid mutator transaction binding the contract method 0x943d552e.
//
// Solidity: function initialize(address _registry, string _stc, string[] _eligibleContractKeys) returns()
func (_SystemBalance *SystemBalanceSession) Initialize(_registry common.Address, _stc string, _eligibleContractKeys []string) (*types.Transaction, error) {
	return _SystemBalance.Contract.Initialize(&_SystemBalance.TransactOpts, _registry, _stc, _eligibleContractKeys)
}

// Initialize is a paid mutator transaction binding the contract method 0x943d552e.
//
// Solidity: function initialize(address _registry, string _stc, string[] _eligibleContractKeys) returns()
func (_SystemBalance *SystemBalanceTransactorSession) Initialize(_registry common.Address, _stc string, _eligibleContractKeys []string) (*types.Transaction, error) {
	return _SystemBalance.Contract.Initialize(&_SystemBalance.TransactOpts, _registry, _stc, _eligibleContractKeys)
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
