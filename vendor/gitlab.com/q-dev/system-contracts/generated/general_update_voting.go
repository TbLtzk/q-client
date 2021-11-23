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

// GeneralUpdateVotingMetaData contains all meta data concerning the GeneralUpdateVoting contract.
var GeneralUpdateVotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIVoting.BaseProposal\",\"name\":\"_proposal\",\"type\":\"tuple\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"QuorumReached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIVoting.VotingOption\",\"name\":\"_votingOption\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"UserVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"VetoOccurred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasRootVetoed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasUserVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_remark\",\"type\":\"string\"}],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"voteFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"voteAgainst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"veto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVotesFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVotesAgainst\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVetosNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVetosPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"internalType\":\"structIVoting.BaseProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalStats\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVetoPercentage\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVotingWeightInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"hasAlreadyVoted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canVote\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ownWeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"votingAgent\",\"type\":\"address\"},{\"internalType\":\"enumDelegationStatus\",\"name\":\"delegationStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"}],\"internalType\":\"structBaseVotingWeightInfo\",\"name\":\"base\",\"type\":\"tuple\"}],\"internalType\":\"structIQthVoting.VotingWeightInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumIVoting.ProposalStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612477806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f65760003560e01c8063bb1d689311610092578063bb1d6893146101f5578063c4d66de814610208578063c7f758a81461021b578063d40b65eb1461023b578063da35c6641461024e578063dc296ae114610256578063e8d2e44214610276578063f99b395414610289578063fe0d94c11461029c576100f6565b8063013cf08b146100fb5780631d28dec714610127578063307a064f1461013c57806349c2a1a61461015c5780635c622a0e1461017c578063750e443a1461019c57806386a50535146101af578063ad0ccf4d146101c2578063b6f61f66146101e2575b600080fd5b61010e610109366004611ca8565b6102af565b60405161011e9493929190611e86565b60405180910390f35b61013a610135366004611ca8565b6103de565b005b61014f61014a366004611ca8565b610565565b60405161011e91906122d1565b61016f61016a366004611ba2565b61079d565b60405161011e9190612364565b61018f61018a366004611ca8565b610bfd565b60405161011e9190611e48565b61013a6101aa366004611ca8565b610eae565b61013a6101bd366004611ca8565b610ef2565b6101d56101d0366004611ca8565b610f32565b60405161011e919061230b565b61016f6101f0366004611ca8565b611055565b61016f610203366004611ca8565b61106a565b61013a610216366004611b4a565b6110b7565b61022e610229366004611ca8565b61117b565b60405161011e91906122be565b61016f610249366004611ca8565b611305565b61016f61131a565b610269610264366004611cd8565b611320565b60405161011e9190611e3d565b610269610284366004611cd8565b611340565b61016f610297366004611ca8565b611360565b61013a6102aa366004611ca8565b6113a6565b60016020818152600092835260409283902080548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529283918301828280156103415780601f1061031657610100808354040283529160200191610341565b820191906000526020600020905b81548152906001019060200180831161032457829003601f168201915b50506040805161010081018252600187015481526002870154602080830191909152600388015482840152600488015460608084019190915260058901546080840152600689015460a0840152600789015460c0840152600889015460e08401528351908101845260098901548152600a89015491810191909152600b88015492810192909252600c9096015494959490935060ff169150859050565b6103e6611422565b6001600160a01b031663a230c524336040518263ffffffff1660e01b81526004016104119190611e10565b60206040518083038186803b15801561042957600080fd5b505afa15801561043d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104619190611b82565b6104865760405162461bcd60e51b815260040161047d906120c4565b60405180910390fd5b80600061049282610bfd565b600781111561049d57fe5b14156104bb5760405162461bcd60e51b815260040161047d90611eca565b60036104c683610bfd565b60078111156104d157fe5b146104ee5760405162461bcd60e51b815260040161047d90611f6a565b600082815260036020908152604080832033845290915290205460ff16156105285760405162461bcd60e51b815260040161047d9061221d565b5060008181526003602090815260408083203384528252808320805460ff19166001908117909155938352908390529020600b0180549091019055565b61056d611994565b81600061057982610bfd565b600781111561058457fe5b14156105a25760405162461bcd60e51b815260040161047d90611eca565b60008381526001602081815260408084208151815460029581161561010002600019011694909404601f8101849004909302840160a0908101909252608084018381529092849284919084018282801561063d5780601f106106125761010080835404028352916020019161063d565b820191906000526020600020905b81548152906001019060200180831161062057829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602082810191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015491810191909152600b8501548184015291830191909152600c9092015460ff16151591015290506106ed611994565b602080830180516080015183525160a0015160408084019190915283015190810151905160009161071e91906114d4565b9050600061074e670de0b6b3a76400006107486402540be40061074243600f611535565b906114d4565b90611535565b9050600061075c838361158e565b60208501819052604086015151909150600090610779908561158e565b60608601819052905061078b896115b3565b60808601525092979650505050505050565b600080546040805160608101909152602280825283926201000090046001600160a01b031691633fb902719161242060208301396040518263ffffffff1660e01b81526004016107ed9190611e73565b60206040518083038186803b15801561080557600080fd5b505afa158015610819573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083d9190611b66565b90506108476119c3565b838152602081015142905260405162498bff60e81b81526001600160a01b0383169063498bff009061087b9060040161203a565b60206040518083038186803b15801561089357600080fd5b505afa1580156108a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108cb9190611cc0565b602080830151429290920191015260405162498bff60e81b81526001600160a01b0383169063498bff00906109029060040161217b565b60206040518083038186803b15801561091a57600080fd5b505afa15801561092e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109529190611cc0565b60208201516080015260405162498bff60e81b81526001600160a01b0383169063498bff009061098490600401612289565b60206040518083038186803b15801561099c57600080fd5b505afa1580156109b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d49190611cc0565b602082015160a0015260405162498bff60e81b8152610a68906001600160a01b0384169063498bff0090610a0a906004016121e6565b60206040518083038186803b158015610a2257600080fd5b505afa158015610a36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5a9190611cc0565b6020808401510151906114d4565b60208201516040908101919091525162498bff60e81b81526001600160a01b0383169063498bff0090610a9d906004016121af565b60206040518083038186803b158015610ab557600080fd5b505afa158015610ac9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aed9190611cc0565b6020808301516060019190915260045460009081526001825260409020825180518493610b1e9284929101906119f7565b506020828101518051600184015580820151600284015560408082015160038501556060808301516004808701919091556080840151600587015560a0840151600687015560c0840151600787015560e0909301516008860155818601518051600987015593840151600a86015592810151600b8501559190930151600c909201805460ff191692151592909217909155905490517fa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e591610be091849061236d565b60405180910390a15050600480546001810190915590505b919050565b60008181526001602081815260408084208151815460029581161561010002600019011694909404601f8101849004909302840160a090810190925260808401838152859493919284928491840182828015610c9a5780601f10610c6f57610100808354040283529160200191610c9a565b820191906000526020600020905b815481529060010190602001808311610c7d57829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602082810191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015481830152600b8601548185015292840192909252600c9093015460ff1615159290910191909152818101510151909150610d59576000915050610bf8565b806060015115610d6d576005915050610bf8565b806020015160200151421015610d87576001915050610bf8565b604081015160208101519051600091610da091906114d4565b90506000610dc4670de0b6b3a76400006107486402540be40061074243600f611535565b90506000610dd2838361158e565b9050836020015160800151811015610df1576002945050505050610bf8565b6000610e058560400151600001518561158e565b9050846020015160a00151811015610e2557600295505050505050610bf8565b610e386002610e32611653565b90611663565b610e41886115b3565b1115610e5557600295505050505050610bf8565b846020015160400151421015610e7357600395505050505050610bf8565b60208501516060810151604090910151610e8c916114d4565b421115610ea157600795505050505050610bf8565b5060049695505050505050565b806000610eba82610bfd565b6007811115610ec557fe5b1415610ee35760405162461bcd60e51b815260040161047d90611eca565b610eee8260026116a2565b5050565b806000610efe82610bfd565b6007811115610f0957fe5b1415610f275760405162461bcd60e51b815260040161047d90611eca565b610eee8260016116a2565b610f3a611a83565b816000610f4682610bfd565b6007811115610f5157fe5b1415610f6f5760405162461bcd60e51b815260040161047d90611eca565b6000610f79611875565b9050610f83611a83565b600085815260016020526040908190206002015490516310874b5b60e01b81526001600160a01b038416916310874b5b91610fc2913391600401611e24565b60806040518083038186803b158015610fda57600080fd5b505afa158015610fee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110129190611c37565b60408083019182526000878152600260209081528282203383529052205460ff161515825251511580159061104657508051155b15156020820152949350505050565b6000908152600160205260409020600a015490565b6000818161107782610bfd565b600781111561108257fe5b14156110a05760405162461bcd60e51b815260040161047d90611eca565b50506000908152600160205260409020600b015490565b600054610100900460ff16806110d057506110d06118db565b806110de575060005460ff16155b6111195760405162461bcd60e51b815260040180806020018281038252602e8152602001806123d1602e913960400191505060405180910390fd5b600054610100900460ff16158015611144576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610eee576000805461ff00191690555050565b6111836119c3565b81600061118f82610bfd565b600781111561119a57fe5b14156111b85760405162461bcd60e51b815260040161047d90611eca565b6000838152600160208181526040928390208351815460029481161561010002600019011693909304601f8101839004909202830160a09081019094526080830182815292939092849290918491908401828280156112585780601f1061122d57610100808354040283529160200191611258565b820191906000526020600020905b81548152906001019060200180831161123b57829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602080830191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015491810191909152600b8501548184015291830191909152600c9092015460ff1615159101529392505050565b60009081526001602052604090206009015490565b60045481565b600260209081526000928352604080842090915290825290205460ff1681565b600360209081526000928352604080842090915290825290205460ff1681565b6000818161136d82610bfd565b600781111561137857fe5b14156113965760405162461bcd60e51b815260040161047d90611eca565b61139f836115b3565b9392505050565b60046113b182610bfd565b60078111156113bc57fe5b146113d95760405162461bcd60e51b815260040161047d9061206d565b6000818152600160208190526040808320600c01805460ff19169092179091555182917f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f91a250565b600080546040805180820182526014815273676f7665726e616e63652e726f6f744e6f64657360601b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb902719161147f91600401611e73565b60206040518083038186803b15801561149757600080fd5b505afa1580156114ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114cf9190611b66565b905090565b60008282018381101561152c576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b90505b92915050565b6000826115445750600061152f565b8282028284828161155157fe5b041461152c5760405162461bcd60e51b81526004018080602001828103825260218152602001806123ff6021913960400191505060405180910390fd5b60008161159d5750600061152f565b61152c82610e326115ac611653565b8690611535565b600061152f6115c0611422565b6001600160a01b031663de8fa4316040518163ffffffff1660e01b815260040160206040518083038186803b1580156115f857600080fd5b505afa15801561160c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116309190611cc0565b610e3261163b611653565b6000868152600160205260409020600b015490611535565b6b033b2e3c9fd0803ce800000090565b600061152c83836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b8152506118ec565b60016116ad83610bfd565b60078111156116b857fe5b146116d55760405162461bcd60e51b815260040161047d90612121565b600082815260026020908152604080832033845290915290205460ff161561170f5760405162461bcd60e51b815260040161047d90611f0f565b60008281526002602090815260408083203384528252808320805460ff19166001908117909155858452909152812090611747611875565b6002830154604051630f7f3f2360e31b81529192506000916001600160a01b03841691637bf9f9189161177e913391600401611e24565b602060405180830381600087803b15801561179857600080fd5b505af11580156117ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117d09190611cc0565b9050600081116117f25760405162461bcd60e51b815260040161047d90611fc5565b600184600281111561180057fe5b141561181f57600983015461181590826114d4565b6009840155611834565b600a83015461182e90826114d4565b600a8401555b847f5ac937fb2a69c6ddee38a23a1b04bbe8a7edb77cdc9bbfe2f9e26dd5a53166d48583604051611866929190611e5c565b60405180910390a25050505050565b60008054604080518082018252601c81527f676f7665726e616e63652e766f74696e6757656967687450726f78790000000060208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb902719161147f91600401611e73565b60006118e63061198e565b15905090565b600081836119785760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561193d578181015183820152602001611925565b50505050905090810190601f16801561196a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161198457fe5b0495945050505050565b3b151590565b6040518060a0016040528060008152602001600081526020016000815260200160008152602001600081525090565b6040518060800160405280606081526020016119dd611aa7565b81526020016119ea611aec565b8152600060209091015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282611a2d5760008555611a73565b82601f10611a4657805160ff1916838001178555611a73565b82800160010185558215611a73579182015b82811115611a73578251825591602001919060010190611a58565b50611a7f929150611b0d565b5090565b6040805160608101825260008082526020820152908101611aa2611b22565b905290565b60405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b60405180606001604052806000815260200160008152602001600081525090565b5b80821115611a7f5760008155600101611b0e565b6040805160808101825260008082526020820181905290918201908152602001600081525090565b600060208284031215611b5b578081fd5b813561152c8161239b565b600060208284031215611b77578081fd5b815161152c8161239b565b600060208284031215611b93578081fd5b8151801515811461152c578182fd5b60006020808385031215611bb4578182fd5b823567ffffffffffffffff80821115611bcb578384fd5b818501915085601f830112611bde578384fd5b813581811115611bea57fe5b604051601f8201601f1916810185018381118282101715611c0757fe5b6040528181528382018501881015611c1d578586fd5b818585018683013790810190930193909352509392505050565b600060808284031215611c48578081fd5b6040516080810181811067ffffffffffffffff82111715611c6557fe5b604052825181526020830151611c7a8161239b565b6020820152604083015160038110611c90578283fd5b60408201526060928301519281019290925250919050565b600060208284031215611cb9578081fd5b5035919050565b600060208284031215611cd1578081fd5b5051919050565b60008060408385031215611cea578081fd5b823591506020830135611cfc8161239b565b809150509250929050565b60008151808452815b81811015611d2c57602081850181015186830182015201611d10565b81811115611d3d5782602083870101525b50601f01601f19169290920160200192915050565b60006101a08251818552611d6882860182611d07565b9150506020830151611d7d6020860182611dc2565b506040830151611d91610120860182611da9565b50606083015115156101808501528091505092915050565b8051825260208082015190830152604090810151910152565b805182526020810151602083015260408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c083015260e081015160e08301525050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b6020810160088310611e5657fe5b91905290565b60408101611e698461238e565b9281526020015290565b60006020825261152c6020830184611d07565b60006101a0808352611e9a81840188611d07565b915050611eaa6020830186611dc2565b611eb8610120830185611da9565b82151561018083015295945050505050565b60208082526025908201527f5b5145432d3032393030365d2d50726f706f73616c20646f6573206e6f7420656040820152643c34b9ba1760d91b606082015260800190565b6020808252603b908201527f5b5145432d3032393030325d2d5468652063616c6c65722068617320616c726560408201527a30b23c903b37ba32b2103337b9103a343290383937b837b9b0b61760291b606082015260800190565b6020808252603b908201527f5b5145432d3032393030345d2d50726f706f73616c20737461747573206d757360408201527a3a1031329030b1b1b2b83a32b216103b32ba37903330b4b632b21760291b606082015260800190565b6020808252604f908201527f5b5145432d3032393030335d2d54686520746f74616c20766f74696e6720776560408201527f69676874206d7573742062652067726561746572207468616e207a65726f2c2060608201526e3330b4b632b2103a37903b37ba329760891b608082015260a00190565b60208082526025908201526000805160206123b183398151915260408201526406e737456560dc1b606082015260800190565b60208082526037908201527f5b5145432d3032393030305d2d50726f706f73616c206d75737420626520504160408201527629a9a2a2103132b337b9329032bc31b2b1baba34b7339760491b606082015260800190565b6020808252603d908201527f5b5145432d3032393030375d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c7920726f6f74206e6f6465732068617665206163636573732e000000606082015260800190565b6020808252603a908201527f5b5145432d3032393030315d2d566f74696e67206973206f6e6c7920706f737360408201527934b136329037b7102822a72224a72390383937b837b9b0b6399760311b606082015260800190565b60208082526026908201526000805160206123b18339815191526040820152656e737451524d60d01b606082015260800190565b6020808252601f908201527f636f6e737469747574696f6e2e70726f706f73616c457865637574696f6e5000604082015260600190565b60208082526029908201526000805160206123b183398151915260408201526806e7374524e56414c560bc1b606082015260800190565b60208082526046908201527f5b5145432d3032393030355d2d5468652073656e6465722068617320616c726560408201527f616479207665746f656420746869732070726f706f73616c2c207665746f206660608201526530b4b632b21760d11b608082015260a00190565b60208082526027908201526000805160206123b18339815191526040820152663739ba2926a0a560c91b606082015260800190565b60006020825261152c6020830184611d52565b600060a082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015292915050565b815115158152602080830151151581830152604080840151805182850152918201516001600160a01b0316606084015281015160c08301919061234d8161238e565b60808401526060015160a090920191909152919050565b90815260200190565b6000838252604060208301526123866040830184611d52565b949350505050565b6003811061239857fe5b50565b6001600160a01b038116811461239857600080fdfe636f6e737469747574696f6e2e766f74696e672e6368616e6765516e6f74436f496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273a2646970667358221220b29d0b87cfde2dd1229fc3bf8aa8e17654e9672f7ac18d2f30e54b2106afa6be64736f6c63430007060033",
}

// GeneralUpdateVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use GeneralUpdateVotingMetaData.ABI instead.
var GeneralUpdateVotingABI = GeneralUpdateVotingMetaData.ABI

// GeneralUpdateVotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GeneralUpdateVotingMetaData.Bin instead.
var GeneralUpdateVotingBin = GeneralUpdateVotingMetaData.Bin

// DeployGeneralUpdateVoting deploys a new Ethereum contract, binding an instance of GeneralUpdateVoting to it.
func DeployGeneralUpdateVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GeneralUpdateVoting, error) {
	parsed, err := GeneralUpdateVotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GeneralUpdateVotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GeneralUpdateVoting{GeneralUpdateVotingCaller: GeneralUpdateVotingCaller{contract: contract}, GeneralUpdateVotingTransactor: GeneralUpdateVotingTransactor{contract: contract}, GeneralUpdateVotingFilterer: GeneralUpdateVotingFilterer{contract: contract}}, nil
}

// GeneralUpdateVoting is an auto generated Go binding around an Ethereum contract.
type GeneralUpdateVoting struct {
	GeneralUpdateVotingCaller     // Read-only binding to the contract
	GeneralUpdateVotingTransactor // Write-only binding to the contract
	GeneralUpdateVotingFilterer   // Log filterer for contract events
}

// GeneralUpdateVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type GeneralUpdateVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneralUpdateVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GeneralUpdateVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneralUpdateVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GeneralUpdateVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneralUpdateVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GeneralUpdateVotingSession struct {
	Contract     *GeneralUpdateVoting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GeneralUpdateVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GeneralUpdateVotingCallerSession struct {
	Contract *GeneralUpdateVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// GeneralUpdateVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GeneralUpdateVotingTransactorSession struct {
	Contract     *GeneralUpdateVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// GeneralUpdateVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type GeneralUpdateVotingRaw struct {
	Contract *GeneralUpdateVoting // Generic contract binding to access the raw methods on
}

// GeneralUpdateVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GeneralUpdateVotingCallerRaw struct {
	Contract *GeneralUpdateVotingCaller // Generic read-only contract binding to access the raw methods on
}

// GeneralUpdateVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GeneralUpdateVotingTransactorRaw struct {
	Contract *GeneralUpdateVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGeneralUpdateVoting creates a new instance of GeneralUpdateVoting, bound to a specific deployed contract.
func NewGeneralUpdateVoting(address common.Address, backend bind.ContractBackend) (*GeneralUpdateVoting, error) {
	contract, err := bindGeneralUpdateVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVoting{GeneralUpdateVotingCaller: GeneralUpdateVotingCaller{contract: contract}, GeneralUpdateVotingTransactor: GeneralUpdateVotingTransactor{contract: contract}, GeneralUpdateVotingFilterer: GeneralUpdateVotingFilterer{contract: contract}}, nil
}

// NewGeneralUpdateVotingCaller creates a new read-only instance of GeneralUpdateVoting, bound to a specific deployed contract.
func NewGeneralUpdateVotingCaller(address common.Address, caller bind.ContractCaller) (*GeneralUpdateVotingCaller, error) {
	contract, err := bindGeneralUpdateVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingCaller{contract: contract}, nil
}

// NewGeneralUpdateVotingTransactor creates a new write-only instance of GeneralUpdateVoting, bound to a specific deployed contract.
func NewGeneralUpdateVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*GeneralUpdateVotingTransactor, error) {
	contract, err := bindGeneralUpdateVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingTransactor{contract: contract}, nil
}

// NewGeneralUpdateVotingFilterer creates a new log filterer instance of GeneralUpdateVoting, bound to a specific deployed contract.
func NewGeneralUpdateVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*GeneralUpdateVotingFilterer, error) {
	contract, err := bindGeneralUpdateVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingFilterer{contract: contract}, nil
}

// bindGeneralUpdateVoting binds a generic wrapper to an already deployed contract.
func bindGeneralUpdateVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GeneralUpdateVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneralUpdateVoting *GeneralUpdateVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GeneralUpdateVoting.Contract.GeneralUpdateVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneralUpdateVoting *GeneralUpdateVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.GeneralUpdateVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneralUpdateVoting *GeneralUpdateVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.GeneralUpdateVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GeneralUpdateVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.contract.Transact(opts, method, params...)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetProposal(opts *bind.CallOpts, _proposalId *big.Int) (IVotingBaseProposal, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getProposal", _proposalId)

	if err != nil {
		return *new(IVotingBaseProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(IVotingBaseProposal)).(*IVotingBaseProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetProposal(_proposalId *big.Int) (IVotingBaseProposal, error) {
	return _GeneralUpdateVoting.Contract.GetProposal(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetProposal(_proposalId *big.Int) (IVotingBaseProposal, error) {
	return _GeneralUpdateVoting.Contract.GetProposal(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetProposalStats(opts *bind.CallOpts, _proposalId *big.Int) (IVotingVotingStats, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getProposalStats", _proposalId)

	if err != nil {
		return *new(IVotingVotingStats), err
	}

	out0 := *abi.ConvertType(out[0], new(IVotingVotingStats)).(*IVotingVotingStats)

	return out0, err

}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetProposalStats(_proposalId *big.Int) (IVotingVotingStats, error) {
	return _GeneralUpdateVoting.Contract.GetProposalStats(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetProposalStats(_proposalId *big.Int) (IVotingVotingStats, error) {
	return _GeneralUpdateVoting.Contract.GetProposalStats(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _proposalId) view returns(uint8)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetStatus(opts *bind.CallOpts, _proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getStatus", _proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _proposalId) view returns(uint8)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetStatus(_proposalId *big.Int) (uint8, error) {
	return _GeneralUpdateVoting.Contract.GetStatus(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _proposalId) view returns(uint8)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetStatus(_proposalId *big.Int) (uint8, error) {
	return _GeneralUpdateVoting.Contract.GetStatus(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVetosNumber is a free data retrieval call binding the contract method 0xbb1d6893.
//
// Solidity: function getVetosNumber(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetVetosNumber(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getVetosNumber", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVetosNumber is a free data retrieval call binding the contract method 0xbb1d6893.
//
// Solidity: function getVetosNumber(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetVetosNumber(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVetosNumber(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVetosNumber is a free data retrieval call binding the contract method 0xbb1d6893.
//
// Solidity: function getVetosNumber(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetVetosNumber(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVetosNumber(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVetosPercentage is a free data retrieval call binding the contract method 0xf99b3954.
//
// Solidity: function getVetosPercentage(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetVetosPercentage(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getVetosPercentage", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVetosPercentage is a free data retrieval call binding the contract method 0xf99b3954.
//
// Solidity: function getVetosPercentage(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetVetosPercentage(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVetosPercentage(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVetosPercentage is a free data retrieval call binding the contract method 0xf99b3954.
//
// Solidity: function getVetosPercentage(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetVetosPercentage(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVetosPercentage(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVotesAgainst is a free data retrieval call binding the contract method 0xb6f61f66.
//
// Solidity: function getVotesAgainst(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetVotesAgainst(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getVotesAgainst", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesAgainst is a free data retrieval call binding the contract method 0xb6f61f66.
//
// Solidity: function getVotesAgainst(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetVotesAgainst(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVotesAgainst(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVotesAgainst is a free data retrieval call binding the contract method 0xb6f61f66.
//
// Solidity: function getVotesAgainst(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetVotesAgainst(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVotesAgainst(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVotesFor is a free data retrieval call binding the contract method 0xd40b65eb.
//
// Solidity: function getVotesFor(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetVotesFor(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getVotesFor", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesFor is a free data retrieval call binding the contract method 0xd40b65eb.
//
// Solidity: function getVotesFor(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetVotesFor(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVotesFor(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVotesFor is a free data retrieval call binding the contract method 0xd40b65eb.
//
// Solidity: function getVotesFor(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetVotesFor(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVotesFor(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVotingWeightInfo is a free data retrieval call binding the contract method 0xad0ccf4d.
//
// Solidity: function getVotingWeightInfo(uint256 _proposalId) view returns((bool,bool,(uint256,address,uint8,uint256)))
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetVotingWeightInfo(opts *bind.CallOpts, _proposalId *big.Int) (IQthVotingVotingWeightInfo, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getVotingWeightInfo", _proposalId)

	if err != nil {
		return *new(IQthVotingVotingWeightInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IQthVotingVotingWeightInfo)).(*IQthVotingVotingWeightInfo)

	return out0, err

}

// GetVotingWeightInfo is a free data retrieval call binding the contract method 0xad0ccf4d.
//
// Solidity: function getVotingWeightInfo(uint256 _proposalId) view returns((bool,bool,(uint256,address,uint8,uint256)))
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetVotingWeightInfo(_proposalId *big.Int) (IQthVotingVotingWeightInfo, error) {
	return _GeneralUpdateVoting.Contract.GetVotingWeightInfo(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVotingWeightInfo is a free data retrieval call binding the contract method 0xad0ccf4d.
//
// Solidity: function getVotingWeightInfo(uint256 _proposalId) view returns((bool,bool,(uint256,address,uint8,uint256)))
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetVotingWeightInfo(_proposalId *big.Int) (IQthVotingVotingWeightInfo, error) {
	return _GeneralUpdateVoting.Contract.GetVotingWeightInfo(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// HasRootVetoed is a free data retrieval call binding the contract method 0xe8d2e442.
//
// Solidity: function hasRootVetoed(uint256 , address ) view returns(bool)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) HasRootVetoed(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "hasRootVetoed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRootVetoed is a free data retrieval call binding the contract method 0xe8d2e442.
//
// Solidity: function hasRootVetoed(uint256 , address ) view returns(bool)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) HasRootVetoed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _GeneralUpdateVoting.Contract.HasRootVetoed(&_GeneralUpdateVoting.CallOpts, arg0, arg1)
}

// HasRootVetoed is a free data retrieval call binding the contract method 0xe8d2e442.
//
// Solidity: function hasRootVetoed(uint256 , address ) view returns(bool)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) HasRootVetoed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _GeneralUpdateVoting.Contract.HasRootVetoed(&_GeneralUpdateVoting.CallOpts, arg0, arg1)
}

// HasUserVoted is a free data retrieval call binding the contract method 0xdc296ae1.
//
// Solidity: function hasUserVoted(uint256 , address ) view returns(bool)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) HasUserVoted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "hasUserVoted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserVoted is a free data retrieval call binding the contract method 0xdc296ae1.
//
// Solidity: function hasUserVoted(uint256 , address ) view returns(bool)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) HasUserVoted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _GeneralUpdateVoting.Contract.HasUserVoted(&_GeneralUpdateVoting.CallOpts, arg0, arg1)
}

// HasUserVoted is a free data retrieval call binding the contract method 0xdc296ae1.
//
// Solidity: function hasUserVoted(uint256 , address ) view returns(bool)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) HasUserVoted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _GeneralUpdateVoting.Contract.HasUserVoted(&_GeneralUpdateVoting.CallOpts, arg0, arg1)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) ProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "proposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) ProposalCount() (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.ProposalCount(&_GeneralUpdateVoting.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) ProposalCount() (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.ProposalCount(&_GeneralUpdateVoting.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(string remark, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params, (uint256,uint256,uint256) counters, bool executed)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Remark   string
	Params   IVotingVotingParams
	Counters IVotingVotingCounters
	Executed bool
}, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Remark   string
		Params   IVotingVotingParams
		Counters IVotingVotingCounters
		Executed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Remark = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Params = *abi.ConvertType(out[1], new(IVotingVotingParams)).(*IVotingVotingParams)
	outstruct.Counters = *abi.ConvertType(out[2], new(IVotingVotingCounters)).(*IVotingVotingCounters)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(string remark, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params, (uint256,uint256,uint256) counters, bool executed)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) Proposals(arg0 *big.Int) (struct {
	Remark   string
	Params   IVotingVotingParams
	Counters IVotingVotingCounters
	Executed bool
}, error) {
	return _GeneralUpdateVoting.Contract.Proposals(&_GeneralUpdateVoting.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(string remark, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params, (uint256,uint256,uint256) counters, bool executed)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) Proposals(arg0 *big.Int) (struct {
	Remark   string
	Params   IVotingVotingParams
	Counters IVotingVotingCounters
	Executed bool
}, error) {
	return _GeneralUpdateVoting.Contract.Proposals(&_GeneralUpdateVoting.CallOpts, arg0)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _remark) returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactor) CreateProposal(opts *bind.TransactOpts, _remark string) (*types.Transaction, error) {
	return _GeneralUpdateVoting.contract.Transact(opts, "createProposal", _remark)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _remark) returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) CreateProposal(_remark string) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.CreateProposal(&_GeneralUpdateVoting.TransactOpts, _remark)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _remark) returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorSession) CreateProposal(_remark string) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.CreateProposal(&_GeneralUpdateVoting.TransactOpts, _remark)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactor) Execute(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.contract.Transact(opts, "execute", _proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) Execute(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.Execute(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorSession) Execute(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.Execute(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _GeneralUpdateVoting.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.Initialize(&_GeneralUpdateVoting.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.Initialize(&_GeneralUpdateVoting.TransactOpts, _registry)
}

// Veto is a paid mutator transaction binding the contract method 0x1d28dec7.
//
// Solidity: function veto(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactor) Veto(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.contract.Transact(opts, "veto", _proposalId)
}

// Veto is a paid mutator transaction binding the contract method 0x1d28dec7.
//
// Solidity: function veto(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) Veto(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.Veto(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// Veto is a paid mutator transaction binding the contract method 0x1d28dec7.
//
// Solidity: function veto(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorSession) Veto(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.Veto(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactor) VoteAgainst(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.contract.Transact(opts, "voteAgainst", _proposalId)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) VoteAgainst(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.VoteAgainst(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorSession) VoteAgainst(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.VoteAgainst(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactor) VoteFor(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.contract.Transact(opts, "voteFor", _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) VoteFor(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.VoteFor(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorSession) VoteFor(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.VoteFor(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// GeneralUpdateVotingProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingProposalCreatedIterator struct {
	Event *GeneralUpdateVotingProposalCreated // Event containing the contract specifics and raw log

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
func (it *GeneralUpdateVotingProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralUpdateVotingProposalCreated)
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
		it.Event = new(GeneralUpdateVotingProposalCreated)
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
func (it *GeneralUpdateVotingProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralUpdateVotingProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralUpdateVotingProposalCreated represents a ProposalCreated event raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingProposalCreated struct {
	Id       *big.Int
	Proposal IVotingBaseProposal
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5.
//
// Solidity: event ProposalCreated(uint256 _id, (string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) _proposal)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*GeneralUpdateVotingProposalCreatedIterator, error) {

	logs, sub, err := _GeneralUpdateVoting.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingProposalCreatedIterator{contract: _GeneralUpdateVoting.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5.
//
// Solidity: event ProposalCreated(uint256 _id, (string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) _proposal)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *GeneralUpdateVotingProposalCreated) (event.Subscription, error) {

	logs, sub, err := _GeneralUpdateVoting.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralUpdateVotingProposalCreated)
				if err := _GeneralUpdateVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0xa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5.
//
// Solidity: event ProposalCreated(uint256 _id, (string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) _proposal)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) ParseProposalCreated(log types.Log) (*GeneralUpdateVotingProposalCreated, error) {
	event := new(GeneralUpdateVotingProposalCreated)
	if err := _GeneralUpdateVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GeneralUpdateVotingProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingProposalExecutedIterator struct {
	Event *GeneralUpdateVotingProposalExecuted // Event containing the contract specifics and raw log

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
func (it *GeneralUpdateVotingProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralUpdateVotingProposalExecuted)
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
		it.Event = new(GeneralUpdateVotingProposalExecuted)
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
func (it *GeneralUpdateVotingProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralUpdateVotingProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralUpdateVotingProposalExecuted represents a ProposalExecuted event raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed _proposalId)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) FilterProposalExecuted(opts *bind.FilterOpts, _proposalId []*big.Int) (*GeneralUpdateVotingProposalExecutedIterator, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _GeneralUpdateVoting.contract.FilterLogs(opts, "ProposalExecuted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingProposalExecutedIterator{contract: _GeneralUpdateVoting.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed _proposalId)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GeneralUpdateVotingProposalExecuted, _proposalId []*big.Int) (event.Subscription, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _GeneralUpdateVoting.contract.WatchLogs(opts, "ProposalExecuted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralUpdateVotingProposalExecuted)
				if err := _GeneralUpdateVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed _proposalId)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) ParseProposalExecuted(log types.Log) (*GeneralUpdateVotingProposalExecuted, error) {
	event := new(GeneralUpdateVotingProposalExecuted)
	if err := _GeneralUpdateVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GeneralUpdateVotingQuorumReachedIterator is returned from FilterQuorumReached and is used to iterate over the raw logs and unpacked data for QuorumReached events raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingQuorumReachedIterator struct {
	Event *GeneralUpdateVotingQuorumReached // Event containing the contract specifics and raw log

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
func (it *GeneralUpdateVotingQuorumReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralUpdateVotingQuorumReached)
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
		it.Event = new(GeneralUpdateVotingQuorumReached)
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
func (it *GeneralUpdateVotingQuorumReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralUpdateVotingQuorumReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralUpdateVotingQuorumReached represents a QuorumReached event raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingQuorumReached struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterQuorumReached is a free log retrieval operation binding the contract event 0x878536ebf930768ad5274a079ba36028b128aeca4b9212fece414176c39e30f7.
//
// Solidity: event QuorumReached(uint256 id)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) FilterQuorumReached(opts *bind.FilterOpts) (*GeneralUpdateVotingQuorumReachedIterator, error) {

	logs, sub, err := _GeneralUpdateVoting.contract.FilterLogs(opts, "QuorumReached")
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingQuorumReachedIterator{contract: _GeneralUpdateVoting.contract, event: "QuorumReached", logs: logs, sub: sub}, nil
}

// WatchQuorumReached is a free log subscription operation binding the contract event 0x878536ebf930768ad5274a079ba36028b128aeca4b9212fece414176c39e30f7.
//
// Solidity: event QuorumReached(uint256 id)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) WatchQuorumReached(opts *bind.WatchOpts, sink chan<- *GeneralUpdateVotingQuorumReached) (event.Subscription, error) {

	logs, sub, err := _GeneralUpdateVoting.contract.WatchLogs(opts, "QuorumReached")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralUpdateVotingQuorumReached)
				if err := _GeneralUpdateVoting.contract.UnpackLog(event, "QuorumReached", log); err != nil {
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

// ParseQuorumReached is a log parse operation binding the contract event 0x878536ebf930768ad5274a079ba36028b128aeca4b9212fece414176c39e30f7.
//
// Solidity: event QuorumReached(uint256 id)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) ParseQuorumReached(log types.Log) (*GeneralUpdateVotingQuorumReached, error) {
	event := new(GeneralUpdateVotingQuorumReached)
	if err := _GeneralUpdateVoting.contract.UnpackLog(event, "QuorumReached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GeneralUpdateVotingUserVotedIterator is returned from FilterUserVoted and is used to iterate over the raw logs and unpacked data for UserVoted events raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingUserVotedIterator struct {
	Event *GeneralUpdateVotingUserVoted // Event containing the contract specifics and raw log

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
func (it *GeneralUpdateVotingUserVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralUpdateVotingUserVoted)
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
		it.Event = new(GeneralUpdateVotingUserVoted)
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
func (it *GeneralUpdateVotingUserVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralUpdateVotingUserVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralUpdateVotingUserVoted represents a UserVoted event raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingUserVoted struct {
	ProposalId   *big.Int
	VotingOption uint8
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserVoted is a free log retrieval operation binding the contract event 0x5ac937fb2a69c6ddee38a23a1b04bbe8a7edb77cdc9bbfe2f9e26dd5a53166d4.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption, uint256 _amount)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) FilterUserVoted(opts *bind.FilterOpts, _proposalId []*big.Int) (*GeneralUpdateVotingUserVotedIterator, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _GeneralUpdateVoting.contract.FilterLogs(opts, "UserVoted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingUserVotedIterator{contract: _GeneralUpdateVoting.contract, event: "UserVoted", logs: logs, sub: sub}, nil
}

// WatchUserVoted is a free log subscription operation binding the contract event 0x5ac937fb2a69c6ddee38a23a1b04bbe8a7edb77cdc9bbfe2f9e26dd5a53166d4.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption, uint256 _amount)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) WatchUserVoted(opts *bind.WatchOpts, sink chan<- *GeneralUpdateVotingUserVoted, _proposalId []*big.Int) (event.Subscription, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _GeneralUpdateVoting.contract.WatchLogs(opts, "UserVoted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralUpdateVotingUserVoted)
				if err := _GeneralUpdateVoting.contract.UnpackLog(event, "UserVoted", log); err != nil {
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

// ParseUserVoted is a log parse operation binding the contract event 0x5ac937fb2a69c6ddee38a23a1b04bbe8a7edb77cdc9bbfe2f9e26dd5a53166d4.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption, uint256 _amount)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) ParseUserVoted(log types.Log) (*GeneralUpdateVotingUserVoted, error) {
	event := new(GeneralUpdateVotingUserVoted)
	if err := _GeneralUpdateVoting.contract.UnpackLog(event, "UserVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GeneralUpdateVotingVetoOccurredIterator is returned from FilterVetoOccurred and is used to iterate over the raw logs and unpacked data for VetoOccurred events raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingVetoOccurredIterator struct {
	Event *GeneralUpdateVotingVetoOccurred // Event containing the contract specifics and raw log

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
func (it *GeneralUpdateVotingVetoOccurredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralUpdateVotingVetoOccurred)
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
		it.Event = new(GeneralUpdateVotingVetoOccurred)
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
func (it *GeneralUpdateVotingVetoOccurredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralUpdateVotingVetoOccurredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralUpdateVotingVetoOccurred represents a VetoOccurred event raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingVetoOccurred struct {
	Id     *big.Int
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVetoOccurred is a free log retrieval operation binding the contract event 0x11e347c8ff2734bf22b4aaead8d3a24eb006d62ee60ab6bbb7adf2827a8e2204.
//
// Solidity: event VetoOccurred(uint256 indexed id, address indexed sender)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) FilterVetoOccurred(opts *bind.FilterOpts, id []*big.Int, sender []common.Address) (*GeneralUpdateVotingVetoOccurredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GeneralUpdateVoting.contract.FilterLogs(opts, "VetoOccurred", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingVetoOccurredIterator{contract: _GeneralUpdateVoting.contract, event: "VetoOccurred", logs: logs, sub: sub}, nil
}

// WatchVetoOccurred is a free log subscription operation binding the contract event 0x11e347c8ff2734bf22b4aaead8d3a24eb006d62ee60ab6bbb7adf2827a8e2204.
//
// Solidity: event VetoOccurred(uint256 indexed id, address indexed sender)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) WatchVetoOccurred(opts *bind.WatchOpts, sink chan<- *GeneralUpdateVotingVetoOccurred, id []*big.Int, sender []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GeneralUpdateVoting.contract.WatchLogs(opts, "VetoOccurred", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralUpdateVotingVetoOccurred)
				if err := _GeneralUpdateVoting.contract.UnpackLog(event, "VetoOccurred", log); err != nil {
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

// ParseVetoOccurred is a log parse operation binding the contract event 0x11e347c8ff2734bf22b4aaead8d3a24eb006d62ee60ab6bbb7adf2827a8e2204.
//
// Solidity: event VetoOccurred(uint256 indexed id, address indexed sender)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) ParseVetoOccurred(log types.Log) (*GeneralUpdateVotingVetoOccurred, error) {
	event := new(GeneralUpdateVotingVetoOccurred)
	if err := _GeneralUpdateVoting.contract.UnpackLog(event, "VetoOccurred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
