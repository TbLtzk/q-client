package core

import (
	"gitlab.com/q-dev/q-client/accounts"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/eth/downloader"
	"gitlab.com/q-dev/q-client/event"
	"gitlab.com/q-dev/q-client/p2p"
)

// TODO move this to types?
type Backend interface {
	AccountManager() *accounts.Manager
	BlockProcessor() *BlockProcessor
	ChainManager() *ChainManager
	TxPool() *TxPool
	PeerCount() int
	IsListening() bool
	Peers() []*p2p.Peer
	BlockDb() common.Database
	StateDb() common.Database
	EventMux() *event.TypeMux
	Downloader() *downloader.Downloader
}
