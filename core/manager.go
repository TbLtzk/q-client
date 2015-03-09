package core

import (
	"gitlab.com/q-dev/q-client/ethutil"
	"gitlab.com/q-dev/q-client/event"
	"gitlab.com/q-dev/q-client/p2p"
)

type Backend interface {
	BlockProcessor() *BlockProcessor
	ChainManager() *ChainManager
	TxPool() *TxPool
	PeerCount() int
	IsListening() bool
	Peers() []*p2p.Peer
	BlockDb() ethutil.Database
	StateDb() ethutil.Database
	EventMux() *event.TypeMux
}
