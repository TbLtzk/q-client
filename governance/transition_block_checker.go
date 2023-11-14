package governance

import (
	"errors"
	"math/big"
	"sync"
	"time"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/core"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/log"
)

const (
	checkerInterval   = 10 * time.Second
	unverifiedChanCap = 25
)

type transitionBlockChecker struct {
	bc *core.BlockChain

	transitionBlockVerifiedBlocks int64

	unverified   map[common.Hash]*types.Header
	unverifiedCh chan *types.Header

	handleCallback func(*types.Header)

	mu sync.RWMutex
}

func newTransitionBlockChecker(verifiedBlocks int64, handleCallback func(*types.Header)) *transitionBlockChecker {
	tbc := &transitionBlockChecker{
		transitionBlockVerifiedBlocks: verifiedBlocks,

		unverified:   make(map[common.Hash]*types.Header),
		unverifiedCh: make(chan *types.Header, unverifiedChanCap),

		handleCallback: handleCallback,
	}

	go tbc.run()

	return tbc
}

func (c *transitionBlockChecker) initBlockChain(bc *core.BlockChain) {
	c.bc = bc
}

func (c *transitionBlockChecker) checkTransitionBlockLater(header *types.Header) (bool, error) {
	if c.bc == nil {
		return false, errors.New("blockchain is not initialized")
	}

	// if current block number - header number < cfg.transitionBlockVerifiedBlocks
	// check the header later to make sure it is in canonical chain
	if new(big.Int).Sub(c.bc.CurrentBlock().Number(), header.Number).Cmp(big.NewInt(c.transitionBlockVerifiedBlocks)) > 0 {
		return false, nil
	}

	log.Debug("Added a new transition block to check later",
		"header hash", header.Hash(),
		"block number", header.Number.Uint64(),
	)

	c.unverifiedCh <- header

	return true, nil
}

func (c *transitionBlockChecker) run() {
	ticker := time.NewTicker(checkerInterval)

	for {
		select {
		case header := <-c.unverifiedCh:
			c.mu.RLock()
			c.unverified[header.Hash()] = header
			c.mu.RUnlock()

			log.Debug("Transition block added for future verification",
				"header hash", header.Hash(),
				"block number", header.Number.Uint64(),
			)
		case <-ticker.C:
			c.mu.Lock()

			for _, header := range c.unverified {
				go c.checkTransitionBlock(header)
			}

			c.mu.Unlock()
		}
	}
}

func (c *transitionBlockChecker) checkTransitionBlock(header *types.Header) {
	if c.bc == nil {
		log.Error("Transition block checker's bc is not initialized")
		return
	}

	// Skip if canonical chain isn't mature enough
	if new(big.Int).Sub(c.bc.CurrentBlock().Number(), header.Number).Cmp(big.NewInt(c.transitionBlockVerifiedBlocks)) < 0 {
		return
	}

	transitionBlockHash := header.Hash()
	c.removeByHash(transitionBlockHash)

	if !c.isCanonical(header) {
		log.Debug("Transition block missing, removing it from the checker",
			"header hash", transitionBlockHash,
			"block number", header.Number.Uint64(),
		)

		return
	}

	log.Debug("Chain has enough height, approving transition block",
		"header transitionBlockHash", transitionBlockHash,
		"block number", header.Number.Uint64(),
	)

	c.handleCallback(header)
}

func (c *transitionBlockChecker) isCanonical(header *types.Header) bool {
	transitionBlockHash := header.Hash()
	c.removeByHash(transitionBlockHash)

	canonicalBlockHash := c.bc.GetCanonicalHash(header.Number.Uint64())
	canonicalBlock := c.bc.GetBlock(canonicalBlockHash, header.Number.Uint64())

	if canonicalBlock.Header().Hash() != transitionBlockHash {
		return false
	}

	return true
}

func (c *transitionBlockChecker) removeByHash(hash common.Hash) {
	c.mu.Lock()
	delete(c.unverified, hash)
	c.mu.Unlock()
}
