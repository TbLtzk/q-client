package core

import (
	"gitlab.com/q-dev/q-client/accounts"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/event"
)

// TODO move this to types?
type Backend interface {
	AccountManager() *accounts.Manager
	BlockProcessor() *BlockProcessor
	ChainManager() *ChainManager
	TxPool() *TxPool
	BlockDb() common.Database
	StateDb() common.Database
	EventMux() *event.TypeMux
}
