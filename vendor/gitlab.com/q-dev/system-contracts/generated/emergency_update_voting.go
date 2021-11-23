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

// EmergencyUpdateVotingMetaData contains all meta data concerning the EmergencyUpdateVoting contract.
var EmergencyUpdateVotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIVoting.BaseProposal\",\"name\":\"_proposal\",\"type\":\"tuple\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIVoting.VotingOption\",\"name\":\"_votingOption\",\"type\":\"uint8\"}],\"name\":\"UserVoted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_remark\",\"type\":\"string\"}],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"voteFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"voteAgainst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVotesFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVotesAgainst\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"internalType\":\"structIVoting.BaseProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalStats\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVetoPercentage\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumIVoting.ProposalStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611dbe806100206000396000f3fe608060405234801561001057600080fd5b50600436106100af5760003560e01c8063013cf08b146100b4578063307a064f146100e057806349c2a1a6146101005780635277b4ae146101205780635c622a0e14610140578063750e443a1461016057806386a5053514610175578063b6f61f6614610188578063c4d66de81461019b578063c7f758a8146101ae578063d40b65eb146101ce578063da35c664146101e1578063fe0d94c1146101e9575b600080fd5b6100c76100c23660046117eb565b6101fc565b6040516100d794939291906119a7565b60405180910390f35b6100f36100ee3660046117eb565b61032b565b6040516100d79190611c9b565b61011361010e366004611756565b6105a6565b6040516100d79190611cd5565b61013361012e36600461181b565b610a17565b6040516100d79190611967565b61015361014e3660046117eb565b610a37565b6040516100d79190611972565b61017361016e3660046117eb565b610cf0565b005b6101736101833660046117eb565b610dd3565b6101136101963660046117eb565b610eb2565b6101736101a93660046116fe565b610ec7565b6101c16101bc3660046117eb565b610f8b565b6040516100d79190611c88565b6101136101dc3660046117eb565b611115565b61011361112a565b6101736101f73660046117eb565b611130565b60016020818152600092835260409283902080548451600294821615610100026000190190911693909304601f810183900483028401830190945283835292839183018282801561028e5780601f106102635761010080835404028352916020019161028e565b820191906000526020600020905b81548152906001019060200180831161027157829003601f168201915b50506040805161010081018252600187015481526002870154602080830191909152600388015482840152600488015460608084019190915260058901546080840152600689015460a0840152600789015460c0840152600889015460e08401528351908101845260098901548152600a89015491810191909152600b88015492810192909252600c9096015494959490935060ff169150859050565b610333611594565b81600061033f82610a37565b600781111561034a57fe5b14156103715760405162461bcd60e51b815260040161036890611a42565b60405180910390fd5b60008381526001602081815260408084208151815460029581161561010002600019011694909404601f8101849004909302840160a0908101909252608084018381529092849284919084018282801561040c5780601f106103e15761010080835404028352916020019161040c565b820191906000526020600020905b8154815290600101906020018083116103ef57829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602082810191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015491810191909152600b8501548184015291830191909152600c9092015460ff16151591015290506104bc611594565b602080830180516080015183525160a001516040808401919091528301519081015190516000916104ed91906111ac565b905060006104f961120d565b6001600160a01b031663de8fa4316040518163ffffffff1660e01b815260040160206040518083038186803b15801561053157600080fd5b505afa158015610545573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105699190611803565b9050600061057783836112bf565b6020850181905260408601515190915060009061059490856112bf565b60608601525092979650505050505050565b60006105b061120d565b6001600160a01b031663a230c524336040518263ffffffff1660e01b81526004016105db9190611953565b60206040518083038186803b1580156105f357600080fd5b505afa158015610607573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061062b9190611736565b6106475760405162461bcd60e51b815260040161036890611bf4565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060228152602001611d67602291396040518263ffffffff1660e01b81526004016106a19190611994565b60206040518083038186803b1580156106b957600080fd5b505afa1580156106cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f1919061171a565b90506106fb6115c3565b838152602081015142905260405162498bff60e81b81526001600160a01b0383169063498bff009061072f90600401611b64565b60206040518083038186803b15801561074757600080fd5b505afa15801561075b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077f9190611803565b602080830151429290920191015260405162498bff60e81b81526001600160a01b0383169063498bff00906107b690600401611ae1565b60206040518083038186803b1580156107ce57600080fd5b505afa1580156107e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108069190611803565b60208201516080015260405162498bff60e81b81526001600160a01b0383169063498bff009061083890600401611b22565b60206040518083038186803b15801561085057600080fd5b505afa158015610864573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108889190611803565b602082015160a0015260405162498bff60e81b81526001600160a01b0383169063498bff00906108ba90600401611c51565b60206040518083038186803b1580156108d257600080fd5b505afa1580156108e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061090a9190611803565b602080830151606001919091526003546000908152600182526040902082518051849361093b9284929101906115f7565b506020828101518051600184015580820151600284015560408082015160038086019190915560608084015160048701556080840151600587015560a0840151600687015560c0840151600787015560e0909301516008860155818601518051600987015593840151600a86015592810151600b850155930151600c909201805460ff1916921515929092179091555490517fa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5916109fa918490611cde565b60405180910390a15050600380546001810190915590505b919050565b600260209081526000928352604080842090915290825290205460ff1681565b60008181526001602081815260408084208151815460029581161561010002600019011694909404601f8101849004909302840160a090810190925260808401838152859493919284928491840182828015610ad45780601f10610aa957610100808354040283529160200191610ad4565b820191906000526020600020905b815481529060010190602001808311610ab757829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602082810191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015481830152600b8601548185015292840192909252600c9093015460ff1615159290910191909152818101510151909150610b93576000915050610a12565b806060015115610ba7576005915050610a12565b806020015160200151421015610bc1576001915050610a12565b604081015160208101519051600091610bda91906111ac565b90506000610be661120d565b6001600160a01b031663de8fa4316040518163ffffffff1660e01b815260040160206040518083038186803b158015610c1e57600080fd5b505afa158015610c32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c569190611803565b90506000610c6483836112bf565b9050836020015160800151811015610c83576002945050505050610a12565b6000610c97856040015160000151856112bf565b9050846020015160a00151811015610cb757600295505050505050610a12565b6020808601516060810151910151610cce916111ac565b421115610ce357600795505050505050610a12565b5060049695505050505050565b610cf861120d565b6001600160a01b031663a230c524336040518263ffffffff1660e01b8152600401610d239190611953565b60206040518083038186803b158015610d3b57600080fd5b505afa158015610d4f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d739190611736565b610d8f5760405162461bcd60e51b815260040161036890611bf4565b806000610d9b82610a37565b6007811115610da657fe5b1415610dc45760405162461bcd60e51b815260040161036890611a42565b610dcf8260026112ea565b5050565b610ddb61120d565b6001600160a01b031663a230c524336040518263ffffffff1660e01b8152600401610e069190611953565b60206040518083038186803b158015610e1e57600080fd5b505afa158015610e32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e569190611736565b610e725760405162461bcd60e51b815260040161036890611bf4565b806000610e7e82610a37565b6007811115610e8957fe5b1415610ea75760405162461bcd60e51b815260040161036890611a42565b610dcf8260016112ea565b6000908152600160205260409020600a015490565b600054610100900460ff1680610ee05750610ee0611433565b80610eee575060005460ff16155b610f295760405162461bcd60e51b815260040180806020018281038252602e815260200180611d18602e913960400191505060405180910390fd5b600054610100900460ff16158015610f54576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610dcf576000805461ff00191690555050565b610f936115c3565b816000610f9f82610a37565b6007811115610faa57fe5b1415610fc85760405162461bcd60e51b815260040161036890611a42565b6000838152600160208181526040928390208351815460029481161561010002600019011693909304601f8101839004909202830160a09081019094526080830182815292939092849290918491908401828280156110685780601f1061103d57610100808354040283529160200191611068565b820191906000526020600020905b81548152906001019060200180831161104b57829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602080830191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015491810191909152600b8501548184015291830191909152600c9092015460ff1615159101529392505050565b60009081526001602052604090206009015490565b60035481565b600461113b82610a37565b600781111561114657fe5b146111635760405162461bcd60e51b8152600401610368906119eb565b6000818152600160208190526040808320600c01805460ff19169092179091555182917f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f91a250565b600082820183811015611204576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b90505b92915050565b600080546040805180820182526014815273676f7665726e616e63652e726f6f744e6f64657360601b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb902719161126a91600401611994565b60206040518083038186803b15801561128257600080fd5b505afa158015611296573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112ba919061171a565b905090565b6000816112ce57506000611207565b611204826112e46112dd611444565b8690611454565b906114ad565b60016112f583610a37565b600781111561130057fe5b1461131d5760405162461bcd60e51b815260040161036890611a87565b600082815260026020908152604080832033845290915290205460ff16156113575760405162461bcd60e51b815260040161036890611b99565b600181600281111561136557fe5b14156113a057600082815260016020819052604090912060090154611389916111ac565b6000838152600160205260409020600901556113d1565b6000828152600160208190526040909120600a01546113be916111ac565b6000838152600160205260409020600a01555b600082815260026020908152604080832033845290915290819020805460ff191660011790555182907fe5a35d72c96d3a6952cdc06084bbe1e9ed2f42834f2f463c89110ab5777b84c590611427908490611986565b60405180910390a25050565b600061143e306114ec565b15905090565b6b033b2e3c9fd0803ce800000090565b60008261146357506000611207565b8282028284828161147057fe5b04146112045760405162461bcd60e51b8152600401808060200182810382526021815260200180611d466021913960400191505060405180910390fd5b600061120483836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b8152506114f2565b3b151590565b6000818361157e5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561154357818101518382015260200161152b565b50505050905090810190601f1680156115705780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161158a57fe5b0495945050505050565b6040518060a0016040528060008152602001600081526020016000815260200160008152602001600081525090565b6040518060800160405280606081526020016115dd611683565b81526020016115ea6116c8565b8152600060209091015290565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928261162d5760008555611673565b82601f1061164657805160ff1916838001178555611673565b82800160010185558215611673579182015b82811115611673578251825591602001919060010190611658565b5061167f9291506116e9565b5090565b60405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b60405180606001604052806000815260200160008152602001600081525090565b5b8082111561167f57600081556001016116ea565b60006020828403121561170f578081fd5b813561120481611cff565b60006020828403121561172b578081fd5b815161120481611cff565b600060208284031215611747578081fd5b81518015158114611204578182fd5b60006020808385031215611768578182fd5b823567ffffffffffffffff8082111561177f578384fd5b818501915085601f830112611792578384fd5b81358181111561179e57fe5b604051601f8201601f19168101850183811182821017156117bb57fe5b60405281815283820185018810156117d1578586fd5b818585018683013790810190930193909352509392505050565b6000602082840312156117fc578081fd5b5035919050565b600060208284031215611814578081fd5b5051919050565b6000806040838503121561182d578081fd5b82359150602083013561183f81611cff565b809150509250929050565b60008151808452815b8181101561186f57602081850181015186830182015201611853565b818111156118805782602083870101525b50601f01601f19169290920160200192915050565b60006101a082518185526118ab8286018261184a565b91505060208301516118c06020860182611905565b5060408301516118d46101208601826118ec565b50606083015115156101808501528091505092915050565b8051825260208082015190830152604090810151910152565b805182526020810151602083015260408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c083015260e081015160e08301525050565b6001600160a01b0391909116815260200190565b901515815260200190565b602081016008831061198057fe5b91905290565b602081016003831061198057fe5b600060208252611204602083018461184a565b60006101a08083526119bb8184018861184a565b9150506119cb6020830186611905565b6119d96101208301856118ec565b82151561018083015295945050505050565b60208082526037908201527f5b5145432d3033303030335d2d50726f706f73616c206d75737420626520504160408201527629a9a2a2103132b337b9329032bc31b2b1baba34b7339760491b606082015260800190565b60208082526025908201527f5b5145432d3033303030355d2d50726f706f73616c20646f6573206e6f7420656040820152643c34b9ba1760d91b606082015260800190565b6020808252603a908201527f5b5145432d3033303030305d2d566f74696e67206973206f6e6c7920706f737360408201527934b136329037b7102822a72224a72390383937b837b9b0b6399760311b606082015260800190565b60208082526021908201527f636f6e737469747574696f6e2e766f74696e672e656d675155706461746551526040820152604d60f81b606082015260800190565b60208082526022908201527f636f6e737469747574696f6e2e766f74696e672e656d6751557064617465524d60408201526120a560f11b606082015260800190565b6020808252818101527f636f6e737469747574696f6e2e766f74696e672e656d67515570646174655650604082015260600190565b6020808252603b908201527f5b5145432d3033303030325d2d5468652063616c6c65722068617320616c726560408201527a30b23c903b37ba32b2103337b9103a343290383937b837b9b0b61760291b606082015260800190565b6020808252603d908201527f5b5145432d3033303030345d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c7920726f6f74206e6f6465732068617665206163636573732e000000606082015260800190565b6020808252601f908201527f636f6e737469747574696f6e2e70726f706f73616c457865637574696f6e5000604082015260600190565b6000602082526112046020830184611895565b600060a082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015292915050565b90815260200190565b600083825260406020830152611cf76040830184611895565b949350505050565b6001600160a01b0381168114611d1457600080fd5b5056fe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273a264697066735822122027512bee354a95c7395409de80b4d2c6e85a7a2bfe5bee945ab6738f02500e2764736f6c63430007060033",
}

// EmergencyUpdateVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use EmergencyUpdateVotingMetaData.ABI instead.
var EmergencyUpdateVotingABI = EmergencyUpdateVotingMetaData.ABI

// EmergencyUpdateVotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EmergencyUpdateVotingMetaData.Bin instead.
var EmergencyUpdateVotingBin = EmergencyUpdateVotingMetaData.Bin

// DeployEmergencyUpdateVoting deploys a new Ethereum contract, binding an instance of EmergencyUpdateVoting to it.
func DeployEmergencyUpdateVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EmergencyUpdateVoting, error) {
	parsed, err := EmergencyUpdateVotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EmergencyUpdateVotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EmergencyUpdateVoting{EmergencyUpdateVotingCaller: EmergencyUpdateVotingCaller{contract: contract}, EmergencyUpdateVotingTransactor: EmergencyUpdateVotingTransactor{contract: contract}, EmergencyUpdateVotingFilterer: EmergencyUpdateVotingFilterer{contract: contract}}, nil
}

// EmergencyUpdateVoting is an auto generated Go binding around an Ethereum contract.
type EmergencyUpdateVoting struct {
	EmergencyUpdateVotingCaller     // Read-only binding to the contract
	EmergencyUpdateVotingTransactor // Write-only binding to the contract
	EmergencyUpdateVotingFilterer   // Log filterer for contract events
}

// EmergencyUpdateVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type EmergencyUpdateVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmergencyUpdateVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EmergencyUpdateVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmergencyUpdateVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EmergencyUpdateVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmergencyUpdateVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EmergencyUpdateVotingSession struct {
	Contract     *EmergencyUpdateVoting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EmergencyUpdateVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EmergencyUpdateVotingCallerSession struct {
	Contract *EmergencyUpdateVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// EmergencyUpdateVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EmergencyUpdateVotingTransactorSession struct {
	Contract     *EmergencyUpdateVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// EmergencyUpdateVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type EmergencyUpdateVotingRaw struct {
	Contract *EmergencyUpdateVoting // Generic contract binding to access the raw methods on
}

// EmergencyUpdateVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EmergencyUpdateVotingCallerRaw struct {
	Contract *EmergencyUpdateVotingCaller // Generic read-only contract binding to access the raw methods on
}

// EmergencyUpdateVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EmergencyUpdateVotingTransactorRaw struct {
	Contract *EmergencyUpdateVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEmergencyUpdateVoting creates a new instance of EmergencyUpdateVoting, bound to a specific deployed contract.
func NewEmergencyUpdateVoting(address common.Address, backend bind.ContractBackend) (*EmergencyUpdateVoting, error) {
	contract, err := bindEmergencyUpdateVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVoting{EmergencyUpdateVotingCaller: EmergencyUpdateVotingCaller{contract: contract}, EmergencyUpdateVotingTransactor: EmergencyUpdateVotingTransactor{contract: contract}, EmergencyUpdateVotingFilterer: EmergencyUpdateVotingFilterer{contract: contract}}, nil
}

// NewEmergencyUpdateVotingCaller creates a new read-only instance of EmergencyUpdateVoting, bound to a specific deployed contract.
func NewEmergencyUpdateVotingCaller(address common.Address, caller bind.ContractCaller) (*EmergencyUpdateVotingCaller, error) {
	contract, err := bindEmergencyUpdateVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVotingCaller{contract: contract}, nil
}

// NewEmergencyUpdateVotingTransactor creates a new write-only instance of EmergencyUpdateVoting, bound to a specific deployed contract.
func NewEmergencyUpdateVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*EmergencyUpdateVotingTransactor, error) {
	contract, err := bindEmergencyUpdateVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVotingTransactor{contract: contract}, nil
}

// NewEmergencyUpdateVotingFilterer creates a new log filterer instance of EmergencyUpdateVoting, bound to a specific deployed contract.
func NewEmergencyUpdateVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*EmergencyUpdateVotingFilterer, error) {
	contract, err := bindEmergencyUpdateVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVotingFilterer{contract: contract}, nil
}

// bindEmergencyUpdateVoting binds a generic wrapper to an already deployed contract.
func bindEmergencyUpdateVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EmergencyUpdateVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmergencyUpdateVoting *EmergencyUpdateVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmergencyUpdateVoting.Contract.EmergencyUpdateVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmergencyUpdateVoting *EmergencyUpdateVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.EmergencyUpdateVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmergencyUpdateVoting *EmergencyUpdateVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.EmergencyUpdateVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmergencyUpdateVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.contract.Transact(opts, method, params...)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) GetProposal(opts *bind.CallOpts, _proposalId *big.Int) (IVotingBaseProposal, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "getProposal", _proposalId)

	if err != nil {
		return *new(IVotingBaseProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(IVotingBaseProposal)).(*IVotingBaseProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) GetProposal(_proposalId *big.Int) (IVotingBaseProposal, error) {
	return _EmergencyUpdateVoting.Contract.GetProposal(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) GetProposal(_proposalId *big.Int) (IVotingBaseProposal, error) {
	return _EmergencyUpdateVoting.Contract.GetProposal(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) GetProposalStats(opts *bind.CallOpts, _proposalId *big.Int) (IVotingVotingStats, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "getProposalStats", _proposalId)

	if err != nil {
		return *new(IVotingVotingStats), err
	}

	out0 := *abi.ConvertType(out[0], new(IVotingVotingStats)).(*IVotingVotingStats)

	return out0, err

}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) GetProposalStats(_proposalId *big.Int) (IVotingVotingStats, error) {
	return _EmergencyUpdateVoting.Contract.GetProposalStats(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) GetProposalStats(_proposalId *big.Int) (IVotingVotingStats, error) {
	return _EmergencyUpdateVoting.Contract.GetProposalStats(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _proposalId) view returns(uint8)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) GetStatus(opts *bind.CallOpts, _proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "getStatus", _proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _proposalId) view returns(uint8)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) GetStatus(_proposalId *big.Int) (uint8, error) {
	return _EmergencyUpdateVoting.Contract.GetStatus(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _proposalId) view returns(uint8)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) GetStatus(_proposalId *big.Int) (uint8, error) {
	return _EmergencyUpdateVoting.Contract.GetStatus(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetVotesAgainst is a free data retrieval call binding the contract method 0xb6f61f66.
//
// Solidity: function getVotesAgainst(uint256 _proposalId) view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) GetVotesAgainst(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "getVotesAgainst", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesAgainst is a free data retrieval call binding the contract method 0xb6f61f66.
//
// Solidity: function getVotesAgainst(uint256 _proposalId) view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) GetVotesAgainst(_proposalId *big.Int) (*big.Int, error) {
	return _EmergencyUpdateVoting.Contract.GetVotesAgainst(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetVotesAgainst is a free data retrieval call binding the contract method 0xb6f61f66.
//
// Solidity: function getVotesAgainst(uint256 _proposalId) view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) GetVotesAgainst(_proposalId *big.Int) (*big.Int, error) {
	return _EmergencyUpdateVoting.Contract.GetVotesAgainst(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetVotesFor is a free data retrieval call binding the contract method 0xd40b65eb.
//
// Solidity: function getVotesFor(uint256 _proposalId) view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) GetVotesFor(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "getVotesFor", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesFor is a free data retrieval call binding the contract method 0xd40b65eb.
//
// Solidity: function getVotesFor(uint256 _proposalId) view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) GetVotesFor(_proposalId *big.Int) (*big.Int, error) {
	return _EmergencyUpdateVoting.Contract.GetVotesFor(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetVotesFor is a free data retrieval call binding the contract method 0xd40b65eb.
//
// Solidity: function getVotesFor(uint256 _proposalId) view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) GetVotesFor(_proposalId *big.Int) (*big.Int, error) {
	return _EmergencyUpdateVoting.Contract.GetVotesFor(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) ProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "proposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) ProposalCount() (*big.Int, error) {
	return _EmergencyUpdateVoting.Contract.ProposalCount(&_EmergencyUpdateVoting.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) ProposalCount() (*big.Int, error) {
	return _EmergencyUpdateVoting.Contract.ProposalCount(&_EmergencyUpdateVoting.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(string remark, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params, (uint256,uint256,uint256) counters, bool executed)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Remark   string
	Params   IVotingVotingParams
	Counters IVotingVotingCounters
	Executed bool
}, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "proposals", arg0)

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
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) Proposals(arg0 *big.Int) (struct {
	Remark   string
	Params   IVotingVotingParams
	Counters IVotingVotingCounters
	Executed bool
}, error) {
	return _EmergencyUpdateVoting.Contract.Proposals(&_EmergencyUpdateVoting.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(string remark, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params, (uint256,uint256,uint256) counters, bool executed)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) Proposals(arg0 *big.Int) (struct {
	Remark   string
	Params   IVotingVotingParams
	Counters IVotingVotingCounters
	Executed bool
}, error) {
	return _EmergencyUpdateVoting.Contract.Proposals(&_EmergencyUpdateVoting.CallOpts, arg0)
}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) Voted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "voted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) Voted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _EmergencyUpdateVoting.Contract.Voted(&_EmergencyUpdateVoting.CallOpts, arg0, arg1)
}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) Voted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _EmergencyUpdateVoting.Contract.Voted(&_EmergencyUpdateVoting.CallOpts, arg0, arg1)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _remark) returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactor) CreateProposal(opts *bind.TransactOpts, _remark string) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.contract.Transact(opts, "createProposal", _remark)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _remark) returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) CreateProposal(_remark string) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.CreateProposal(&_EmergencyUpdateVoting.TransactOpts, _remark)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _remark) returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorSession) CreateProposal(_remark string) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.CreateProposal(&_EmergencyUpdateVoting.TransactOpts, _remark)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactor) Execute(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.contract.Transact(opts, "execute", _proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) Execute(_proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.Execute(&_EmergencyUpdateVoting.TransactOpts, _proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorSession) Execute(_proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.Execute(&_EmergencyUpdateVoting.TransactOpts, _proposalId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.Initialize(&_EmergencyUpdateVoting.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.Initialize(&_EmergencyUpdateVoting.TransactOpts, _registry)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactor) VoteAgainst(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.contract.Transact(opts, "voteAgainst", _proposalId)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) VoteAgainst(_proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.VoteAgainst(&_EmergencyUpdateVoting.TransactOpts, _proposalId)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorSession) VoteAgainst(_proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.VoteAgainst(&_EmergencyUpdateVoting.TransactOpts, _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactor) VoteFor(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.contract.Transact(opts, "voteFor", _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) VoteFor(_proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.VoteFor(&_EmergencyUpdateVoting.TransactOpts, _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorSession) VoteFor(_proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.VoteFor(&_EmergencyUpdateVoting.TransactOpts, _proposalId)
}

// EmergencyUpdateVotingProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the EmergencyUpdateVoting contract.
type EmergencyUpdateVotingProposalCreatedIterator struct {
	Event *EmergencyUpdateVotingProposalCreated // Event containing the contract specifics and raw log

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
func (it *EmergencyUpdateVotingProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmergencyUpdateVotingProposalCreated)
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
		it.Event = new(EmergencyUpdateVotingProposalCreated)
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
func (it *EmergencyUpdateVotingProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmergencyUpdateVotingProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmergencyUpdateVotingProposalCreated represents a ProposalCreated event raised by the EmergencyUpdateVoting contract.
type EmergencyUpdateVotingProposalCreated struct {
	Id       *big.Int
	Proposal IVotingBaseProposal
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5.
//
// Solidity: event ProposalCreated(uint256 _id, (string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) _proposal)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*EmergencyUpdateVotingProposalCreatedIterator, error) {

	logs, sub, err := _EmergencyUpdateVoting.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVotingProposalCreatedIterator{contract: _EmergencyUpdateVoting.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5.
//
// Solidity: event ProposalCreated(uint256 _id, (string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) _proposal)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *EmergencyUpdateVotingProposalCreated) (event.Subscription, error) {

	logs, sub, err := _EmergencyUpdateVoting.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmergencyUpdateVotingProposalCreated)
				if err := _EmergencyUpdateVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) ParseProposalCreated(log types.Log) (*EmergencyUpdateVotingProposalCreated, error) {
	event := new(EmergencyUpdateVotingProposalCreated)
	if err := _EmergencyUpdateVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmergencyUpdateVotingProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the EmergencyUpdateVoting contract.
type EmergencyUpdateVotingProposalExecutedIterator struct {
	Event *EmergencyUpdateVotingProposalExecuted // Event containing the contract specifics and raw log

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
func (it *EmergencyUpdateVotingProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmergencyUpdateVotingProposalExecuted)
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
		it.Event = new(EmergencyUpdateVotingProposalExecuted)
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
func (it *EmergencyUpdateVotingProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmergencyUpdateVotingProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmergencyUpdateVotingProposalExecuted represents a ProposalExecuted event raised by the EmergencyUpdateVoting contract.
type EmergencyUpdateVotingProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed _proposalId)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) FilterProposalExecuted(opts *bind.FilterOpts, _proposalId []*big.Int) (*EmergencyUpdateVotingProposalExecutedIterator, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _EmergencyUpdateVoting.contract.FilterLogs(opts, "ProposalExecuted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVotingProposalExecutedIterator{contract: _EmergencyUpdateVoting.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed _proposalId)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *EmergencyUpdateVotingProposalExecuted, _proposalId []*big.Int) (event.Subscription, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _EmergencyUpdateVoting.contract.WatchLogs(opts, "ProposalExecuted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmergencyUpdateVotingProposalExecuted)
				if err := _EmergencyUpdateVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) ParseProposalExecuted(log types.Log) (*EmergencyUpdateVotingProposalExecuted, error) {
	event := new(EmergencyUpdateVotingProposalExecuted)
	if err := _EmergencyUpdateVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmergencyUpdateVotingUserVotedIterator is returned from FilterUserVoted and is used to iterate over the raw logs and unpacked data for UserVoted events raised by the EmergencyUpdateVoting contract.
type EmergencyUpdateVotingUserVotedIterator struct {
	Event *EmergencyUpdateVotingUserVoted // Event containing the contract specifics and raw log

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
func (it *EmergencyUpdateVotingUserVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmergencyUpdateVotingUserVoted)
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
		it.Event = new(EmergencyUpdateVotingUserVoted)
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
func (it *EmergencyUpdateVotingUserVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmergencyUpdateVotingUserVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmergencyUpdateVotingUserVoted represents a UserVoted event raised by the EmergencyUpdateVoting contract.
type EmergencyUpdateVotingUserVoted struct {
	ProposalId   *big.Int
	VotingOption uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserVoted is a free log retrieval operation binding the contract event 0xe5a35d72c96d3a6952cdc06084bbe1e9ed2f42834f2f463c89110ab5777b84c5.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) FilterUserVoted(opts *bind.FilterOpts, _proposalId []*big.Int) (*EmergencyUpdateVotingUserVotedIterator, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _EmergencyUpdateVoting.contract.FilterLogs(opts, "UserVoted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVotingUserVotedIterator{contract: _EmergencyUpdateVoting.contract, event: "UserVoted", logs: logs, sub: sub}, nil
}

// WatchUserVoted is a free log subscription operation binding the contract event 0xe5a35d72c96d3a6952cdc06084bbe1e9ed2f42834f2f463c89110ab5777b84c5.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) WatchUserVoted(opts *bind.WatchOpts, sink chan<- *EmergencyUpdateVotingUserVoted, _proposalId []*big.Int) (event.Subscription, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _EmergencyUpdateVoting.contract.WatchLogs(opts, "UserVoted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmergencyUpdateVotingUserVoted)
				if err := _EmergencyUpdateVoting.contract.UnpackLog(event, "UserVoted", log); err != nil {
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

// ParseUserVoted is a log parse operation binding the contract event 0xe5a35d72c96d3a6952cdc06084bbe1e9ed2f42834f2f463c89110ab5777b84c5.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) ParseUserVoted(log types.Log) (*EmergencyUpdateVotingUserVoted, error) {
	event := new(EmergencyUpdateVotingUserVoted)
	if err := _EmergencyUpdateVoting.contract.UnpackLog(event, "UserVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
