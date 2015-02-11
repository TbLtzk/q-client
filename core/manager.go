package core

import (
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/ethutil"
	"gitlab.com/q-dev/q-client/event"
	"gitlab.com/q-dev/q-client/p2p"
)

type EthManager interface {
	BlockProcessor() *BlockProcessor
	ChainManager() *ChainManager
	TxPool() *TxPool
	PeerCount() int
	IsMining() bool
	IsListening() bool
	Peers() []*p2p.Peer
	KeyManager() *crypto.KeyManager
	ClientIdentity() p2p.ClientIdentity
	Db() ethutil.Database
	EventMux() *event.TypeMux
}
