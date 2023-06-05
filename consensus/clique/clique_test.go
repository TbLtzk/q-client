// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package clique

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/contracts"
	"gitlab.com/q-dev/q-client/core"
	"gitlab.com/q-dev/q-client/core/rawdb"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/core/vm"
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/params"
)

// This test case is a repro of an annoying bug that took us forever to catch.
// In Clique PoA networks (Rinkeby, Görli, etc), consecutive blocks might have
// the same state root (no block subsidy, empty block). If a node crashes, the
// chain ends up losing the recent state and needs to regenerate it from blocks
// already in the database. The bug was that processing the block *prior* to an
// empty one **also completes** the empty one, ending up in a known-block error.
func TestReimportMirroredState(t *testing.T) {
	// Initialize a Clique chain with a single signer
	var (
		db     = rawdb.NewMemoryDatabase()
		key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
		addr   = crypto.PubkeyToAddress(key.PublicKey)
		reg    = contracts.NewTestModeRegistry()
		engine = New(params.AllCliqueProtocolChanges.Clique, db, &NoopExclusionSetProvider{}, reg)
		signer = new(types.HomesteadSigner)
	)
	genspec := &core.Genesis{
		ExtraData: make([]byte, extraVanity+common.AddressLength+extraSeal),
		Alloc: map[common.Address]core.GenesisAccount{
			addr: {Balance: big.NewInt(10000000000000000)},
		},
		BaseFee: big.NewInt(params.InitialBaseFee),
	}
	copy(genspec.ExtraData[extraVanity:], addr[:])
	genesis := genspec.MustCommit(db)

	// Generate a batch of blocks, each properly signed
	chain, _ := core.NewBlockChain(db, nil, params.AllCliqueProtocolChanges, engine, vm.Config{}, nil, nil)
	defer chain.Stop()

	blocks, _ := core.GenerateChain(params.AllCliqueProtocolChanges, genesis, engine, db, 3, func(i int, block *core.BlockGen) {
		// The chain maker doesn't have access to a chain, so the difficulty will be
		// lets unset (nil). Set it here to the correct value.
		block.SetDifficulty(diffInTurn)

		// We want to simulate an empty middle block, having the same state as the
		// first one. The last is needs a state change again to force a reorg.
		if i != 1 {
			tx, err := types.SignTx(types.NewTransaction(block.TxNonce(addr), common.Address{0x00}, new(big.Int), params.TxGas, block.BaseFee(), nil), signer, key)
			if err != nil {
				panic(err)
			}
			block.AddTxWithChain(chain, tx)
		}
	})
	for i, block := range blocks {
		header := block.Header()
		if i > 0 {
			header.ParentHash = blocks[i-1].Hash()
		}
		header.Extra = make([]byte, extraVanity+extraSeal)
		header.Difficulty = diffInTurn
		header.Coinbase = reg.RewardReceiver()

		sig, _ := crypto.Sign(SealHash(header).Bytes(), key)
		copy(header.Extra[len(header.Extra)-extraSeal:], sig)
		blocks[i] = block.WithSeal(header)
	}
	// Insert the first two blocks and make sure the chain is valid
	db = rawdb.NewMemoryDatabase()
	genspec.MustCommit(db)

	chain, _ = core.NewBlockChain(db, nil, params.AllCliqueProtocolChanges, engine, vm.Config{}, nil, nil)
	defer chain.Stop()

	if _, err := chain.InsertChain(blocks[:2]); err != nil {
		t.Fatalf("failed to insert initial blocks: %v", err)
	}
	if head := chain.CurrentBlock().NumberU64(); head != 2 {
		t.Fatalf("chain head mismatch: have %d, want %d", head, 2)
	}

	// Simulate a crash by creating a new chain on top of the database, without
	// flushing the dirty states out. Insert the last block, triggering a sidechain
	// reimport.
	chain, _ = core.NewBlockChain(db, nil, params.AllCliqueProtocolChanges, engine, vm.Config{}, nil, nil)
	defer chain.Stop()

	if _, err := chain.InsertChain(blocks[2:]); err != nil {
		t.Fatalf("failed to insert final block: %v", err)
	}
	if head := chain.CurrentBlock().NumberU64(); head != 3 {
		t.Fatalf("chain head mismatch: have %d, want %d", head, 3)
	}
}

func TestSealHash(t *testing.T) {
	have := SealHash(&types.Header{
		Difficulty: new(big.Int),
		Number:     new(big.Int),
		Extra:      make([]byte, 32+65),
		BaseFee:    new(big.Int),
	})
	want := common.HexToHash("0xbd3d1fa43fbc4c5bfcc91b179ec92e2861df3654de60468beb908ff805359e8f")
	if have != want {
		t.Errorf("have %x, want %x", have, want)
	}
}

func TestFallbackToDefaultSigners(t *testing.T) {
	var (
		db                   = rawdb.NewMemoryDatabase()
		reg                  = contracts.NewTestModeRegistry()
		exclusionSetProvider = new(NoopExclusionSetProvider)
		engine               = New(params.AllCliqueProtocolChanges.Clique, db, exclusionSetProvider, reg)
		genesisSignersCount  = 3
		keys                 = make([]*ecdsa.PrivateKey, 0, genesisSignersCount)
		addrs                = make([]common.Address, 0, genesisSignersCount)
	)
	engine.fakeDiff = true

	for len(keys) != genesisSignersCount {
		key, _ := crypto.GenerateKey()
		keys = append(keys, key)
		addrs = append(addrs, crypto.PubkeyToAddress(key.PublicKey))
	}

	genSpec := &core.Genesis{
		ExtraData: make([]byte, extraVanity+common.AddressLength*genesisSignersCount+extraSeal),
		Alloc: map[common.Address]core.GenesisAccount{
			addrs[0]: {Balance: big.NewInt(10000000000000000)},
		},
		BaseFee: big.NewInt(params.InitialBaseFee),
	}
	for i := 0; i < genesisSignersCount; i++ {
		copy(genSpec.ExtraData[extraVanity+i*common.AddressLength:], addrs[i][:])
	}
	genesis := genSpec.MustCommit(db)

	blocks, _ := core.GenerateChain(params.AllCliqueProtocolChanges, genesis, engine, db, genesisSignersCount*2, func(i int, block *core.BlockGen) {})

	for i, block := range blocks {
		header := block.Header()
		if i > 0 {
			header.ParentHash = blocks[i-1].Hash()
		}
		header.Extra = make([]byte, extraVanity+extraSeal)
		header.Difficulty = diffInTurn
		header.Coinbase = reg.RewardReceiver()

		sig, _ := crypto.Sign(SealHash(header).Bytes(), keys[i%genesisSignersCount])
		copy(header.Extra[len(header.Extra)-extraSeal:], sig)
		blocks[i] = block.WithSeal(header)
	}

	chain, _ := core.NewBlockChain(db, nil, params.AllCliqueProtocolChanges, engine, vm.Config{}, nil, nil)
	defer chain.Stop()

	if _, err := chain.InsertChain(blocks); err != nil {
		t.Fatalf("failed to insert initial blocks: %v", err)
	}

	// Exclude signers
	exclusionSetProvider.set = make(map[common.Address][]common.BlockRange)
	currentBlock := chain.CurrentBlock()
	for _, signer := range addrs {
		exclusionSetProvider.set[signer] = []common.BlockRange{
			{
				StartAddress: currentBlock.NumberU64(),
				EndAddress:   currentBlock.NumberU64() + 1000,
			},
		}
	}
	excludedSigners := engine.exclusionSetProvider.ExclusionSetValidators()
	assert.Equal(t, len(excludedSigners), genesisSignersCount)

	snap, err := engine.snapshot(chain, currentBlock.NumberU64(), currentBlock.Hash(), nil, false)
	if err != nil {
		t.Fatalf("faield to get snapshot of last block: %v", err)
	}

	assert.Equal(t, len(filterSigners(currentBlock.NumberU64(), snap.SignersList(), excludedSigners)), 0)

	blocks, _ = core.GenerateChain(params.AllCliqueProtocolChanges, chain.CurrentBlock(), engine, db, 1, func(i int, block *core.BlockGen) {})

	// Try to sign a new block with all validators banned
	block := blocks[0]
	header := block.Header()
	header.ParentHash = chain.CurrentBlock().Hash()
	header.Extra = make([]byte, extraVanity+extraSeal)
	header.Difficulty = diffInTurn
	header.Coinbase = reg.RewardReceiver()

	sig, _ := crypto.Sign(SealHash(header).Bytes(), keys[0])
	copy(header.Extra[len(header.Extra)-extraSeal:], sig)
	blocks[0] = block.WithSeal(header)

	if _, err := chain.InsertChain(blocks); err != nil {
		t.Fatalf("failed to insert the new block: %v", err)
	}

	snap, err = engine.snapshot(chain, chain.CurrentBlock().NumberU64(), chain.CurrentBlock().Hash(), nil, false)
	if err != nil {
		t.Fatalf("faield to get snapshot of last block: %v", err)
	}
	assert.Equal(t, len(snap.Signers), genesisSignersCount)
}
