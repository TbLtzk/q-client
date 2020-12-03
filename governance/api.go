package governance

import (
	"gitlab.com/q-dev/q-client/common"
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
func (a *GovernanceAPI) ActiveRootList() common.RootList {
	return a.gov.RootManager.current.makeList()
}
