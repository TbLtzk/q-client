package governance

import (
	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/log"
)

// GovernanceAPI.
type GovernanceAPI struct {
	gov *Governance
}

// NewGovernanceAPI.
func NewGovernanceAPI(back *Governance) *GovernanceAPI {
	return &GovernanceAPI{gov: back}
}

// ActiveRootList.
func (a *GovernanceAPI) ActiveRootList() *RootList {
	return newRootList(a.gov.RootManager.getActiveRootSet())
}

func (a *GovernanceAPI) DesiredRootList() *RootList {
	return newRootList(a.gov.RootManager.getDesiredRootSet())
}

func (a *GovernanceAPI) ProposedRootList() *RootList {
	return newRootList(a.gov.RootManager.getProposedRootSet())
}

func (a *GovernanceAPI) ProposeRootListUpdate(list common.RootList) (common.Hash, error) {
	set, err := newRootSet(&list)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "invalid root list")
	}

	rm := a.gov.RootManager
	if !rm.isMember(set.rootAddresses) {
		return common.Hash{}, errors.New("not a member of new list")
	}

	rm.rootLock.Lock()
	defer rm.rootLock.Unlock()

	if set.timestamp <= rm.active.timestamp || (rm.desired != nil && set.timestamp <= rm.desired.timestamp) {
		return common.Hash{}, errors.New("obsolete root list")
	}

	if rm.signRootSet(set) {
		log.Info("signed desired root list")
	}

	rm.desired = set
	rm.desiredRootFeed.Send(set)

	rm.db.saveDesiredRootSet(set)

	return set.hash, nil
}

func (a *GovernanceAPI) AcceptProposedRootList() error {
	rm := a.gov.RootManager

	rm.rootLock.Lock()
	defer rm.rootLock.Unlock()

	if rm.proposed == nil {
		return errors.New("proposed list is empty")
	}

	if rm.desired != nil && rm.proposed.timestamp <= rm.desired.timestamp {
		return errors.New("proposed list is obsolete")
	}

	if rm.signRootSet(rm.proposed) {
		log.Info("Signed proposed root list", "hash", rm.proposed.hash.Hex())
	}

	rm.desired = rm.proposed
	rm.desiredRootFeed.Send(rm.desired)

	// very edge case but still possible
	if rm.active.isAcceptable(rm.proposed) {
		rm.upgradeRootSet(rm.proposed)
	}

	rm.proposed = nil
	rm.db.deleteProposedRootSet()
	return nil
}

func (a *GovernanceAPI) DiffRootList(nameA, nameB string) ([]DiffEntry, error) {
	return a.gov.RootManager.diffRootListByName(nameA, nameB)
}

func (a *GovernanceAPI) ActiveExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.getActiveExclusionSet())
}

func (a *GovernanceAPI) DesiredExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.getDesiredExclusionSet())
}

func (a *GovernanceAPI) ProposedExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.getProposedExclusionSet())
}

func (a *GovernanceAPI) ProposeExclusionListUpdate(list common.ValidatorExclusionList) (common.Hash, error) {
	set, err := newExclusionSet(&list)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "invalid exclusion list")
	}

	rm := a.gov.RootManager
	if !rm.isRootNode() {
		return common.Hash{}, errors.New("not a root node")
	}

	rm.exLock.Lock()
	defer rm.exLock.Unlock()

	olderThanActive := rm.activeExSet != nil && set.timestamp <= rm.activeExSet.timestamp
	olderThanDesired := rm.desiredExSet != nil && set.timestamp <= rm.desiredExSet.timestamp
	if olderThanActive || olderThanDesired {
		return common.Hash{}, errors.New("obsolete exclusion list")
	}

	rm.signExclusionSet(set)

	rm.desiredExSet = set
	rm.desiredExFeed.Send(set.copy())

	rm.db.saveDesiredExclusionSet(set)
	return set.hash, nil
}

func (a *GovernanceAPI) AcceptProposedExclusionList() error {
	rm := a.gov.RootManager

	rm.exLock.Lock()
	defer rm.exLock.Unlock()

	if rm.proposedExSet == nil {
		return errors.New("proposed list is empty")
	}

	if rm.desiredExSet != nil && rm.proposedExSet.timestamp <= rm.desiredExSet.timestamp {
		return errors.New("proposed list is obsolete")
	}

	rm.signExclusionSet(rm.proposedExSet)

	rm.desiredExSet = rm.proposedExSet
	rm.desiredExFeed.Send(rm.desiredExSet.copy())

	// very edge case but still possible
	if rm.getActiveRootSet().isEnoughExSetSignatures(rm.proposedExSet) {
		rm.upgradeExclusionSet(rm.proposedExSet)
	}

	rm.proposedExSet = nil
	rm.db.deleteProposedExclusionSet()
	return nil
}

func (a *GovernanceAPI) DiffExclusionList(nameA, nameB string) ([]DiffEntry, error) {
	return a.gov.RootManager.diffExclusionListByName(nameA, nameB)
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
