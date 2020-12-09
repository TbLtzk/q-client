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
func (a *GovernanceAPI) ActiveRootList() RootList {
	return newRootList(a.gov.RootManager.currentRootSet())
}

func (a *GovernanceAPI) DesiredRootList() RootList {
	set := a.gov.RootManager.targetRootSet()
	if set == nil {
		return RootList{}
	}

	return newRootList(set)
}

func (a *GovernanceAPI) ProposedRootList() RootList {
	set := a.gov.RootManager.proposedRootSet()
	if set == nil {
		return RootList{}
	}

	return newRootList(set)
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

type RootList struct {
	Timestamp uint64           `json:"timestamp"`
	Nodes     []common.Address `json:"nodes"`
	Hash      common.Hash      `json:"hash"`

	Signers []common.Address `json:"signers"`
}

func newRootList(set *rootSet) RootList {
	var signers []common.Address
	for addr := range set.signers {
		signers = append(signers, addr)
	}

	return RootList{
		Timestamp: set.timestamp,
		Hash:      set.hash,
		Nodes:     set.rootAddresses,
		Signers:   signers,
	}
}
