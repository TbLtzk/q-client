package governance

import (
	"math/big"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/log"
)

// GovernanceAPI. (shouldn't be opened via http/ws)
type GovernanceAPI struct {
	*GovernancePublicAPI
}

// GovernanceExtAPI
type GovernancePublicAPI struct {
	gov *Governance
}

// NewGovernanceAPI.
func NewGovernanceAPI(back *Governance, extApi *GovernancePublicAPI) *GovernanceAPI {
	return &GovernanceAPI{GovernancePublicAPI: extApi}
}

// NewGovernancePublicAPI
func NewGovernancePublicAPI(back *Governance) *GovernancePublicAPI {
	return &GovernancePublicAPI{gov: back}
}

func (a *GovernancePublicAPI) ActiveRootList() *RootList {
	return newRootList(a.gov.RootManager.getActiveRootSet(true))
}

func (a *GovernancePublicAPI) DesiredRootList() *RootList {
	return newRootList(a.gov.RootManager.getDesiredRootSet(true))
}

func (a *GovernancePublicAPI) ProposedRootList() *RootList {
	return newRootList(a.gov.RootManager.getProposedRootSet(true))
}

func (a *GovernancePublicAPI) OnchainRootList() *RootList {
	return newRootList(a.gov.RootManager.getOnchainRootSet(true))
}

func (a *GovernanceAPI) ProposeRootListUpdate(list common.RootList) (common.Hash, error) {
	set, err := newRootSet(&list)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "invalid root list")
	}
	set.updateAliases(a.gov.RootManager.getAliasesOfRoots(set.rootAddresses))

	set, err = a.gov.RootManager.proposeRootSet(set)
	if err != nil {
		return common.Hash{}, err
	}

	return set.hash, nil
}

func (a *GovernanceAPI) ProposeOnchainRootList() (common.Hash, error) {
	set := a.gov.RootManager.getOnchainRootSet(true)
	if set == nil {
		return common.Hash{}, errors.New("can't get on-cain root set")
	}

	set, err := a.gov.RootManager.proposeRootSet(set)
	if err != nil {
		return common.Hash{}, err
	}

	return set.hash, nil
}

func (a *GovernanceAPI) AcceptProposedRootList() error {
	return a.gov.RootManager.acceptProposedRootList(true)
}

func (a *GovernancePublicAPI) DiffRootList(nameA, nameB string) ([]DiffEntry, error) {
	return a.gov.RootManager.diffRootListByName(nameA, nameB, true)
}

func (a *GovernancePublicAPI) ActiveExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.getActiveExclusionSet())
}

func (a *GovernancePublicAPI) DesiredExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.getDesiredExclusionSet())
}

func (a *GovernancePublicAPI) ProposedExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.getProposedExclusionSet())
}

func (a *GovernanceAPI) ProposeExclusionListUpdate(list common.ValidatorExclusionList) (common.Hash, error) {
	set, err := newExclusionSet(&list)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "invalid exclusion list")
	}

	set, err = a.gov.RootManager.proposeExclusionSet(set)
	if err != nil {
		return common.Hash{}, err
	}

	return set.hash, nil
}

func (a *GovernanceAPI) AcceptProposedExclusionList() error {
	return a.gov.RootManager.acceptProposedExclusionList(true)
}

func (a *GovernancePublicAPI) DiffExclusionList(nameA, nameB string) ([]DiffEntry, error) {
	return a.gov.RootManager.diffExclusionListByName(nameA, nameB)
}

func (a *GovernanceAPI) GetRootNodeApprovals(blockNumber *big.Int, hash *common.Hash) (*[]common.RootNodeApproval, error) {
	list, err := a.gov.RootManager.getActiveApprovalList(blockNumber, hash)
	res := []common.RootNodeApproval{}
	if list != nil {
		res = list.Approvals
	}
	return &res, err
}

func (a *GovernanceAPI) GetLatestTransitionBlocks(amount int) (*[]TransitionBlocksWithApproval, error) {
	currentBlockNumber := a.gov.RootManager.bc.CurrentBlock().Number().Uint64()
	if uint64(amount) > currentBlockNumber/101 {
		log.Error("Not enough transition blocks!")
		return nil, nil
	}
	transitionBlockNumber := currentBlockNumber - currentBlockNumber%101
	transitionBlocks := make([]TransitionBlocksWithApproval, amount)
	for amount > 0 {
		transitionBlock := a.gov.RootManager.bc.GetHeaderByNumber(transitionBlockNumber)
		rNApprovals, err := a.GetRootNodeApprovals(transitionBlock.Number, nil)
		if err != nil {
			return nil, err
		}
		var approvals []Approval
		for _, approval := range *rNApprovals {
			if approval.Hash == transitionBlock.Hash() {
				approvals = append(approvals, Approval{
					Signer:    approval.Signer,
					Signature: approval.Signature,
				})
			}
		}
		transitionBlocks[amount-1] = TransitionBlocksWithApproval{
			BlockNumber: transitionBlock.Number,
			Hash:        transitionBlock.Hash(),
			Approvals:   approvals,
		}
		transitionBlockNumber = transitionBlockNumber - 101
		amount--
	}
	return &transitionBlocks, nil
}

type Approval struct {
	Signer    common.Address `json:"signer"`
	Signature []byte         `json:"signature"`
}

type TransitionBlocksWithApproval struct {
	BlockNumber *big.Int    `json:"blockNumber"`
	Hash        common.Hash `json:"hash"`
	Approvals   []Approval  `json:"approvals"`
}

type RootList struct {
	Timestamp uint64           `json:"timestamp"`
	Nodes     []common.Address `json:"nodes"`
	Hash      common.Hash      `json:"hash"`

	Signers []common.Address `json:"signers"`
}

func newRootList(set *rootSet) *RootList {
	if set == nil {
		return nil
	}

	var signers []common.Address
	for addr := range set.signers {
		signers = append(signers, addr)
	}

	return &RootList{
		Timestamp: set.timestamp,
		Hash:      set.hash,
		Nodes:     set.rootAddresses,
		Signers:   signers,
	}
}

type ExclusionList struct {
	Timestamp  uint64                     `json:"timestamp"`
	Validators []common.ExcludedValidator `json:"validators"`
	Hash       common.Hash                `json:"hash"`

	Signers []common.Address `json:"signers"`
}

func newExclusionList(set *exclusionSet) *ExclusionList {
	if set == nil {
		return nil
	}

	var signers []common.Address
	for addr := range set.signers {
		signers = append(signers, addr)
	}

	l := set.makeList()
	return &ExclusionList{
		Timestamp:  set.timestamp,
		Hash:       set.hash,
		Validators: l.Validators,
		Signers:    signers,
	}
}
