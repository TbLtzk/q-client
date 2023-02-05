package governance

import (
	"math/big"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
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

func (a *GovernancePublicAPI) ActiveExclusionListPrettify() *ExclusionListPrettify {
	return newExclusionListPrettify(a.gov.RootManager.getActiveExclusionSet(), a.gov.RootManager.bc.CurrentBlock().Number().Int64())
}

func (a *GovernancePublicAPI) DesiredExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.getDesiredExclusionSet())
}

func (a *GovernancePublicAPI) DesiredExclusionListPrettify() *ExclusionListPrettify {
	return newExclusionListPrettify(a.gov.RootManager.getDesiredExclusionSet(), a.gov.RootManager.bc.CurrentBlock().Number().Int64())
}

func (a *GovernancePublicAPI) ProposedExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.getProposedExclusionSet())
}

func (a *GovernancePublicAPI) ProposedExclusionListPrettify() *ExclusionListPrettify {
	return newExclusionListPrettify(a.gov.RootManager.getProposedExclusionSet(), a.gov.RootManager.bc.CurrentBlock().Number().Int64())
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
	var res []common.RootNodeApproval
	if list != nil {
		res = list.Approvals
	}
	return &res, err
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

type ExclusionListPrettify struct {
	Hash       common.Hash                 `json:"hash"`
	Timestamp  string                      `json:"timestamp (created at)"`
	Signers    []common.Address            `json:"signers"`
	Validators map[common.Address][]string `json:"validators"`
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

func newExclusionListPrettify(set *exclusionSet, currentBlock int64) *ExclusionListPrettify {
	if set == nil {
		return nil
	}

	var signers []common.Address
	for addr := range set.signers {
		signers = append(signers, addr)
	}

	formattedBlockRanges := make(map[common.Address][]string)

	for address, blockRangesByAddress := range set.blockRanges {
		var stringBlockRanges []string
		var endBlockPlusState string
		for _, br := range blockRangesByAddress {
			switch {
			case br.EndAddress == 0:
				endBlockPlusState = " (active, everlasting)"
			case int64(br.EndAddress) < currentBlock && int64(br.StartAddress) < currentBlock:
				endBlockPlusState = " - #" + strconv.FormatInt(int64(br.EndAddress), 10) + " (expired)"
			case int64(br.EndAddress) > currentBlock && int64(br.StartAddress) > currentBlock:
				endBlockPlusState = " - #" + strconv.FormatInt(int64(br.EndAddress), 10) +
					" (scheduled at block #" + strconv.FormatInt(int64(br.StartAddress), 10) + ")"
			case int64(br.EndAddress) > currentBlock && int64(br.StartAddress) < currentBlock:
				endBlockPlusState = " - #" + strconv.FormatInt(int64(br.EndAddress), 10) +
					" (active, ends at block #" + strconv.FormatInt(int64(br.EndAddress), 10) + ")"
			}
			stringBlockRange := "#" + strconv.FormatInt(int64(br.StartAddress), 10) + endBlockPlusState
			stringBlockRanges = append(stringBlockRanges, stringBlockRange)
		}
		formattedBlockRanges[address] = stringBlockRanges
	}

	timeUnix := time.Unix(int64(set.timestamp), 0)
	return &ExclusionListPrettify{
		Timestamp:  strconv.FormatInt(int64(set.timestamp), 10) + " (" + timeUnix.String() + ")",
		Hash:       set.hash,
		Signers:    signers,
		Validators: formattedBlockRanges,
	}
}
