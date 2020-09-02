package governance

import (
	"github.com/pkg/errors"
	"gitlab.com/q-dev/go-ethereum/accounts/keystore"
	"gitlab.com/q-dev/go-ethereum/log"
	"gitlab.com/q-dev/go-ethereum/node"
	"gitlab.com/q-dev/go-ethereum/p2p"
	"gitlab.com/q-dev/go-ethereum/rpc"
)

// Config of governance svc.
type Config struct {
	InstanceDir string `toml:"-"`
}

// Governance service is responsible
// for 2nd layer functionality.
type Governance struct {
	rootMgr *RootManager
	handler *handler
}

// New Governance service.
func New(ctx *node.ServiceContext, cfg *Config) (*Governance, error) {
	ks := ctx.AccountManager.Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	rootMgr, err := NewRootManager(ks, cfg.InstanceDir)
	if err != nil {
		return nil, err
	}

	handler := newHandler(rootMgr)
	return &Governance{
		rootMgr: rootMgr,
		handler: handler,
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
	return nil
}

// Start Governance service.
func (g *Governance) Start(srv *p2p.Server) error {
	if err := g.rootMgr.run(); err != nil {
		return errors.Wrap(err, "failed to start root manager")
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
