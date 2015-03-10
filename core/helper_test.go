package core

import (
	"container/list"
	"fmt"

	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/ethdb"
	"gitlab.com/q-dev/q-client/ethutil"
	"gitlab.com/q-dev/q-client/event"
)

// Implement our EthTest Manager
type TestManager struct {
	// stateManager *StateManager
	eventMux *event.TypeMux

	db         ethutil.Database
	txPool     *TxPool
	blockChain *ChainManager
	Blocks     []*types.Block
}

func (s *TestManager) IsListening() bool {
	return false
}

func (s *TestManager) IsMining() bool {
	return false
}

func (s *TestManager) PeerCount() int {
	return 0
}

func (s *TestManager) Peers() *list.List {
	return list.New()
}

func (s *TestManager) ChainManager() *ChainManager {
	return s.blockChain
}

func (tm *TestManager) TxPool() *TxPool {
	return tm.txPool
}

// func (tm *TestManager) StateManager() *StateManager {
// 	return tm.stateManager
// }

func (tm *TestManager) EventMux() *event.TypeMux {
	return tm.eventMux
}
func (tm *TestManager) KeyManager() *crypto.KeyManager {
	return nil
}

func (tm *TestManager) Db() ethutil.Database {
	return tm.db
}

func NewTestManager() *TestManager {
	db, err := ethdb.NewMemDatabase()
	if err != nil {
		fmt.Println("Could not create mem-db, failing")
		return nil
	}

	testManager := &TestManager{}
	testManager.eventMux = new(event.TypeMux)
	testManager.db = db
	// testManager.txPool = NewTxPool(testManager)
	// testManager.blockChain = NewChainManager(testManager)
	// testManager.stateManager = NewStateManager(testManager)

	// Start the tx pool
	testManager.txPool.Start()

	return testManager
}
