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
	return newRootList(a.gov.RootManager.currentRootSet())
}

func (a *GovernanceAPI) DesiredRootList() *RootList {
	return newRootList(a.gov.RootManager.targetRootSet())
}

func (a *GovernanceAPI) ProposedRootList() *RootList {
	return newRootList(a.gov.RootManager.proposedRootSet())
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

	rm.lock.Lock()
	defer rm.lock.Unlock()

	if set.timestamp <= rm.current.timestamp || (rm.target != nil && set.timestamp <= rm.target.timestamp) {
		return common.Hash{}, errors.New("obsolete root list")
	}

	if rm.signRootSet(set) {
		log.Info("signed desired root list")
	}

	rm.target = set
	rm.targetRootFeed.Send(set)

	rm.db.saveDesiredRootSet(set)

	return set.hash, nil
}

func (a *GovernanceAPI) AcceptProposedRootList() error {
	rm := a.gov.RootManager

	rm.lock.Lock()
	defer rm.lock.Unlock()

	if rm.proposed == nil {
		return errors.New("proposed list is empty")
	}

	if rm.target != nil && rm.proposed.timestamp <= rm.target.timestamp {
		return errors.New("proposed list is obsolete")
	}

	if rm.signRootSet(rm.proposed) {
		log.Info("Signed proposed root list", "hash", rm.proposed.hash.Hex())
	}

	rm.target = rm.proposed
	rm.targetRootFeed.Send(rm.target)

	// very edge case but still possible
	if rm.current.isAcceptable(rm.proposed) {
		rm.upgradeRootSet(rm.proposed)
	}

	rm.proposed = nil
	rm.db.deleteProposedRootSet()
	return nil
}

func (a *GovernanceAPI) ActiveExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.currentExclusionSet())
}

func (a *GovernanceAPI) DesiredExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.desiredExclusionSet())
}

func (a *GovernanceAPI) ProposedExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.proposedExclusionSet())
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

	olderThanCurrent := rm.exclusionSet != nil && set.timestamp <= rm.exclusionSet.timestamp
	olderThanDesired := rm.targetExSet != nil && set.timestamp <= rm.targetExSet.timestamp
	if olderThanCurrent || olderThanDesired {
		return common.Hash{}, errors.New("obsolete exclusion list")
	}

	rm.signExclusionSet(set)

	rm.targetExSet = set
	rm.targetExListFeed.Send(set.copy())

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

	if rm.targetExSet != nil && rm.proposedExSet.timestamp <= rm.targetExSet.timestamp {
		return errors.New("proposed list is obsolete")
	}

	rm.signExclusionSet(rm.proposedExSet)

	rm.targetExSet = rm.proposedExSet
	rm.targetExListFeed.Send(rm.targetExSet.copy())

	// very edge case but still possible
	if rm.currentRootSet().isEnoughExSetSignatures(rm.proposedExSet) {
		rm.upgradeExclusionSet(rm.proposedExSet)
	}

	rm.proposedExSet = nil
	rm.db.deleteProposedExclusionSet()
	return nil
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
