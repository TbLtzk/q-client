package governance

import (
	"gitlab.com/q-dev/q-client/accounts/keystore"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/node"
	"gitlab.com/q-dev/q-client/p2p"
	"gitlab.com/q-dev/q-client/rpc"
)

// Config of governance svc.
type Config struct {
	RootList common.RootList `toml:"-"`
}

// Governance service is responsible
// for 2nd layer functionality.
type Governance struct {
	RootManager *RootManager

	handler *handler
}

// New Governance service.
func New(stack *node.Node, cfg *Config) (*Governance, error) {
	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	rootMgr, err := newRootManager(ks, stack.InstanceDir(), cfg)
	if err != nil {
		return nil, err
	}

	handler := newHandler(rootMgr)
	return &Governance{
		RootManager: rootMgr,
		handler:     handler,
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
