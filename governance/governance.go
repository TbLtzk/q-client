package governance

import (
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/node"
	"gitlab.com/q-dev/q-client/p2p"
	"gitlab.com/q-dev/q-client/rpc"
)

// Governance service is responsible
// for 2nd layer functionality.
type Governance struct {
	RootManager         *RootManager
	ConstitutionManager *ConstitutionManager

	handler *handler
}

// New Governance service.
func New(stack *node.Node, rm *RootManager) (*Governance, error) {

	cm, errCm := NewConstitutionManager(stack.ConstitutionDir(), rm.db, rm)
	if errCm != nil {
		log.Error("Can't create ConstitutionManager: %v", errCm)
	}

	handler := newHandler(rm, cm)

	return &Governance{
		RootManager:         rm,
		ConstitutionManager: cm,
		handler:             handler,
	}, nil
}

// Protocols supported by governance.
func (g *Governance) Protocols() []p2p.Protocol {
	var protocols []p2p.Protocol
	for _, version := range ProtocolVersions {
		protocols = append(protocols, g.handler.makeProtocol(version))
	}

	return protocols
}

// APIs provided by governance.
func (g *Governance) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "gov",
			Version:   "1.0",
			Service:   NewGovernanceAPI(g),
			Public:    false,
		},
	}
}

// Start Governance service.
func (g *Governance) Start() error {
	if g.RootManager.isRootNode() {
		log.Info("Node belongs to the current root node set")
	}

	g.handler.run()

	return nil
}

// Stop Governance service.
func (g *Governance) Stop() error {
	g.handler.stop()
	log.Info("governance svc is down")
	return nil
}
