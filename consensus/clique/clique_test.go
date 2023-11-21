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
	"errors"
	"math/big"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/consensus"
	"gitlab.com/q-dev/q-client/contracts"
	"gitlab.com/q-dev/q-client/contracts/mocks"
	"gitlab.com/q-dev/q-client/core"
	"gitlab.com/q-dev/q-client/core/rawdb"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/core/vm"
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/ethdb"
	"gitlab.com/q-dev/q-client/log"
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
		engine = New(params.AllCliqueProtocolChanges.Clique, db, &consensus.NoopExclusionSetProvider{}, reg)
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

func TestForkChoice(t *testing.T) {
	var (
		db                = rawdb.NewMemoryDatabase()
		reg               = contracts.NewTestModeRegistry()
		engine            = New(params.AllCliqueProtocolChanges.Clique, db, &consensus.NoopExclusionSetProvider{}, reg)
		signersCount      = 40
		commonBlocksCount = 80
		keys              = make(map[common.Address]*ecdsa.PrivateKey)
		addresses         = make([]common.Address, 0, signersCount)
	)
	engine.fakeDiff = true

	for len(addresses) != signersCount {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		keys[addr] = key
		addresses = append(addresses, addr)
	}
	sort.Sort(signersAscending(addresses))

	genSpec := &core.Genesis{
		ExtraData: make([]byte, extraVanity+common.AddressLength*signersCount+extraSeal),
		Alloc: map[common.Address]core.GenesisAccount{
			addresses[0]: {Balance: big.NewInt(10000000000000000)},
		},
		BaseFee:    big.NewInt(params.InitialBaseFee),
		Difficulty: big.NewInt(0),
	}
	for i := 0; i < signersCount; i++ {
		copy(genSpec.ExtraData[extraVanity+i*common.AddressLength:], addresses[i][:])
	}
	genesis := genSpec.MustCommit(db)

	var chain *core.BlockChain

	shouldPreserve := func(header *types.Header, externalHeader *types.Header) bool {
		rHeader, err := engine.ChooseBlockWithMostRecentSigner(chain, header, externalHeader)
		if err != nil {
			log.Warn("Failed to retrieve recent signer list for preserve check for local header", "number", header.Number.Uint64(), "hash", header.Hash(), "err", err)
			return false
		}

		return rHeader == header
	}

	//Generate chain common for both cases and both subchains
	commonBlocks := generateChain(t, params.AllCliqueProtocolChanges, genesis, engine, db, commonBlocksCount, addresses, keys, reg, false)

	//foreign chain is canonical
	chain, _ = core.NewBlockChain(db, nil, params.AllCliqueProtocolChanges, engine, vm.Config{}, shouldPreserve, nil)
	if _, err := chain.InsertChain(commonBlocks); err != nil {
		t.Fatalf("failed to insert commonBlocks blocks: %v", err)
	}
	localBlocksSigners1 := addresses[10:15]
	localBlocksCase1 := generateChain(t, params.AllCliqueProtocolChanges, commonBlocks[len(commonBlocks)-1], engine, db, 5, localBlocksSigners1, keys, reg, true)
	if _, err := chain.InsertChain(localBlocksCase1); err != nil {
		t.Fatalf("failed to insert local blocks: %v", err)
	}

	foreignBlocksSigners1 := addresses[10:13]
	foreignBlocksSigners1 = append(foreignBlocksSigners1, addresses[5])
	foreignBlocksSigners1 = append(foreignBlocksSigners1, addresses[15]) //same signer as in local chain
	foreignBlocksCase1 := generateChain(t, params.AllCliqueProtocolChanges, commonBlocks[len(commonBlocks)-1], engine, db, 5, foreignBlocksSigners1, keys, reg, true)
	if _, err := chain.InsertChain(foreignBlocksCase1); err != nil {
		t.Fatalf("failed to insert foreign blocks: %v", err)
	}

	var minHash common.Hash
	if foreignBlocksCase1[len(foreignBlocksCase1)-1].Hash().Big().Cmp(localBlocksCase1[len(localBlocksCase1)-1].Hash().Big()) < 0 {
		minHash = foreignBlocksCase1[len(foreignBlocksCase1)-1].Hash()
	} else {
		minHash = localBlocksCase1[len(localBlocksCase1)-1].Hash()
	}
	if chain.CurrentBlock().Header().Hash() != minHash {
		t.Errorf("Rule 4 failed")
	}
}

// Generates blocks with the required signer and difficulty in the middle
func generateChain(t *testing.T, config *params.ChainConfig, parent *types.Block, engine *Clique, db ethdb.Database, numBlocks int, addresses []common.Address, keys map[common.Address]*ecdsa.PrivateKey, reg *contracts.Registry, side bool) []*types.Block {
	current := 1
	blocks, _ := core.GenerateChain(config, parent, engine, db, numBlocks, nil)

	for i, block := range blocks {
		header := block.Header()
		if i > 0 {
			header.ParentHash = blocks[i-1].Hash()
		} else {
			header.ParentHash = parent.Hash()
		}

		signer := addresses[current]

		difficulty := diffInTurn

		key := keys[signer]
		header.Extra = make([]byte, extraVanity+extraSeal)
		header.Difficulty = difficulty
		header.Coinbase = reg.RewardReceiver()

		sig, _ := crypto.Sign(SealHash(header).Bytes(), key)
		copy(header.Extra[len(header.Extra)-extraSeal:], sig)
		blocks[i] = block.WithSeal(header)

		current++
		if current == len(addresses) {
			current = 0
		}
	}

	return blocks
}

// nolint:unused,deadcode
func testBlockChoiceRule4(t *testing.T) {
	var (
		db           = rawdb.NewMemoryDatabase()
		reg          = contracts.NewTestModeRegistry()
		engine       = New(params.AllCliqueProtocolChanges.Clique, db, &consensus.NoopExclusionSetProvider{}, reg)
		signersCount = 3
		blocksCount  = 7
		keys         = make(map[common.Address]*ecdsa.PrivateKey)
		addrs        = make([]common.Address, 0, signersCount)
		signer       = new(types.HomesteadSigner)
	)

	for len(addrs) != signersCount {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		keys[addr] = key
		addrs = append(addrs, addr)
	}

	sort.Sort(signersAscending(addrs))

	genSpec := &core.Genesis{
		ExtraData: make([]byte, extraVanity+common.AddressLength*signersCount+extraSeal),
		Alloc: map[common.Address]core.GenesisAccount{
			addrs[0]: {Balance: big.NewInt(10000000000000000)},
		},
		BaseFee:    big.NewInt(params.InitialBaseFee),
		Difficulty: big.NewInt(0),
	}
	for i := 0; i < signersCount; i++ {
		copy(genSpec.ExtraData[extraVanity+i*common.AddressLength:], addrs[i][:])
	}
	genesis := genSpec.MustCommit(db)

	chain, _ := core.NewBlockChain(db, nil, params.AllCliqueProtocolChanges, engine, vm.Config{}, nil, nil)
	defer chain.Stop()

	// Generate two chains. The side chain's last block hash will differ
	canon, _ := core.GenerateChain(params.AllCliqueProtocolChanges, genesis, engine, db, blocksCount, nil)
	sideChain, _ := core.GenerateChain(params.AllCliqueProtocolChanges, genesis, engine, db, blocksCount, func(i int, block *core.BlockGen) {
		if i == blocksCount-1 {
			addr := addrs[0]
			tx, err := types.SignTx(types.NewTransaction(block.TxNonce(addr), common.Address{0x00}, new(big.Int), params.TxGas, block.BaseFee(), nil), signer, keys[addr])
			if err != nil {
				panic(err)
			}
			block.AddTxWithChain(chain, tx)
		}
	})

	for i, block := range canon {
		header := block.Header()
		if i > 0 {
			header.ParentHash = canon[i-1].Hash()
		}
		header.Extra = make([]byte, extraVanity+extraSeal)
		header.Difficulty = diffInTurn
		header.Coinbase = reg.RewardReceiver()

		sig, _ := crypto.Sign(SealHash(header).Bytes(), keys[addrs[block.NumberU64()%uint64(signersCount)]])
		copy(header.Extra[len(header.Extra)-extraSeal:], sig)
		canon[i] = block.WithSeal(header)
	}

	if _, err := chain.InsertChain(canon); err != nil {
		t.Fatalf("failed to insert initial blocks: %v", err)
	}

	for i, block := range sideChain {
		header := block.Header()
		if i > 0 {
			header.ParentHash = sideChain[i-1].Hash()
		}
		header.Extra = make([]byte, extraVanity+extraSeal)
		header.Difficulty = diffInTurn
		header.Coinbase = reg.RewardReceiver()

		sig, _ := crypto.Sign(SealHash(header).Bytes(), keys[addrs[block.NumberU64()%uint64(signersCount)]])
		copy(header.Extra[len(header.Extra)-extraSeal:], sig)
		sideChain[i] = block.WithSeal(header)
	}

	if _, err := chain.InsertChain(sideChain); err != nil {
		t.Fatalf("failed to insert sidechain: %v", err)
	}

	// Compare hashes of the last blocks and choose the accepted chain
	canonHash := canon[len(canon)-1].Hash()
	sideHash := sideChain[len(sideChain)-1].Hash()

	winnerHash := canonHash
	if new(big.Int).SetBytes(canonHash.Bytes()).Cmp(new(big.Int).SetBytes(sideHash.Bytes())) > 0 {
		winnerHash = sideHash
	}
	assert.Equal(t, chain.CurrentBlock().Hash(), winnerHash)
}

func TestFallbackToDefaultSigners(t *testing.T) {
	var (
		db                   = rawdb.NewMemoryDatabase()
		reg                  = contracts.NewTestModeRegistry()
		exclusionSetProvider = new(consensus.NoopExclusionSetProvider)
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
	exclusionSetProvider.Set = make(map[common.Address][]common.BlockRange)
	currentBlock := chain.CurrentBlock()
	for _, signer := range addrs {
		exclusionSetProvider.Set[signer] = []common.BlockRange{
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

// TestReorgExclusionSet checks the behavior of a consensus engine during a blockchain reorganization.
// The test involves two separate blockchain instances with their own databases and exclusion set providers.
// It simulates a scenario where blocks are added to both chains, and then an attempt is made to insert
// a sidechain into the primary chain, but the chains have different sets of excluded validators
func TestReorgExclusionSet(t *testing.T) {
	var (
		db1                   = rawdb.NewMemoryDatabase()
		db2                   = rawdb.NewMemoryDatabase()
		exclusionSetProvider1 = new(consensus.NoopExclusionSetProvider)
		exclusionSetProvider2 = new(consensus.NoopExclusionSetProvider)
		reg1                  = contracts.NewTestModeRegistry()
		reg2                  = contracts.NewTestModeRegistry()
		engine1               = New(params.AllCliqueProtocolChanges.Clique, db1, exclusionSetProvider1, reg1)
		engine2               = New(params.AllCliqueProtocolChanges.Clique, db2, exclusionSetProvider2, reg2)
		signersCount          = 3
		keys                  = make(map[common.Address]*ecdsa.PrivateKey)
		addresses             = make([]common.Address, 0, signersCount)
		commonBlocksCount     = 80
	)

	for len(addresses) != signersCount {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		keys[addr] = key
		addresses = append(addresses, addr)
	}
	sort.Sort(signersAscending(addresses))

	genSpec := &core.Genesis{
		ExtraData:  make([]byte, extraVanity+common.AddressLength*signersCount+extraSeal),
		Alloc:      map[common.Address]core.GenesisAccount{},
		BaseFee:    big.NewInt(params.InitialBaseFee),
		Difficulty: big.NewInt(0),
	}
	for i := 0; i < signersCount; i++ {
		copy(genSpec.ExtraData[extraVanity+i*common.AddressLength:], addresses[i][:])
	}

	genesis1 := genSpec.MustCommit(db1)
	genesis2 := genSpec.MustCommit(db2)

	commonBlocks1 := generateChain(t, params.AllCliqueProtocolChanges, genesis1, engine1, db1, commonBlocksCount,
		addresses, keys, reg1, false)
	commonBlocks2 := generateChain(t, params.AllCliqueProtocolChanges, genesis2, engine2, db2, commonBlocksCount,
		addresses, keys, reg2, false)

	// Insert common blocks
	chain1, _ := core.NewBlockChain(db1, nil, params.AllCliqueProtocolChanges, engine1, vm.Config{},
		nil, nil)
	defer chain1.Stop()
	if _, err := chain1.InsertChain(commonBlocks1); err != nil {
		t.Fatalf("failed to insert commonBlocks blocks: %v", err)
	}
	chain2, _ := core.NewBlockChain(db2, nil, params.AllCliqueProtocolChanges, engine2, vm.Config{},
		nil, nil)
	defer chain2.Stop()
	if _, err := chain2.InsertChain(commonBlocks2); err != nil {
		t.Fatalf("failed to insert commonBlocks blocks: %v", err)
	}

	blocksToGenerate := signersCount * 3
	sideChain := generateChain(t, params.AllCliqueProtocolChanges, chain2.CurrentBlock(), engine2, db2,
		blocksToGenerate, addresses, keys, reg2, true)

	// Modify exclusion set of the first blockchain
	exclusionSetProvider1.Set = map[common.Address][]common.BlockRange{
		addresses[0]: {
			{
				StartAddress: uint64(commonBlocksCount + 1),
				EndAddress:   uint64(commonBlocksCount + blocksToGenerate + 1),
			},
		},
	}
	engine1.exclusionSetProvider = exclusionSetProvider1

	_, err := chain1.InsertChain(sideChain)
	if err == nil {
		t.Fatalf("shouldn't validate blocks")
	}
	if !errors.Is(err, errUnauthorizedSigner) && !errors.Is(err, errWrongDifficulty) {
		t.Fatalf("failed to insert sidechain: %v", err)
	}
}

// TestReorgRegistry examines the behavior of a consensus engine in a scenario involving blockchain
// reorganization due to changes in the validator registry. This test case creates two blockchain
// instances with separate databases and exclusion set providers.
// It simulates a situation where both chains share common blocks but then diverge as new addresses
// are added to one of the chains' validator registries.
func TestReorgRegistry(t *testing.T) {
	var (
		db1                   = rawdb.NewMemoryDatabase()
		db2                   = rawdb.NewMemoryDatabase()
		exclusionSetProvider1 = new(consensus.NoopExclusionSetProvider)
		exclusionSetProvider2 = new(consensus.NoopExclusionSetProvider)
		signersCount          = 5
		keys                  = make(map[common.Address]*ecdsa.PrivateKey)
		addresses1            = make([]common.Address, 0, signersCount)
		addresses2            = make([]common.Address, 0, signersCount)
		commonBlocksCount     = 99
	)

	for len(addresses1) != signersCount {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		keys[addr] = key
		addresses1 = append(addresses1, addr)
		addresses2 = append(addresses2, addr)
	}
	sort.Sort(signersAscending(addresses1))
	sort.Sort(signersAscending(addresses2))

	reg1 := contracts.NewTestModeRegistryWithMocks(
		mocks.NewMockValidators(common.HexToAddress("0x1"), addresses1, nil),
		nil,
	)
	reg2 := contracts.NewTestModeRegistryWithMocks(
		mocks.NewMockValidators(common.HexToAddress("0x1"), addresses2, nil),
		nil,
	)
	engine1 := New(params.AllCliqueProtocolChanges.Clique, db1, exclusionSetProvider1, reg1)
	engine2 := New(params.AllCliqueProtocolChanges.Clique, db2, exclusionSetProvider2, reg2)

	genSpec := &core.Genesis{
		ExtraData:  make([]byte, extraVanity+common.AddressLength*signersCount+extraSeal),
		Alloc:      map[common.Address]core.GenesisAccount{},
		BaseFee:    big.NewInt(params.InitialBaseFee),
		Difficulty: big.NewInt(0),
	}
	for i := 0; i < signersCount; i++ {
		copy(genSpec.ExtraData[extraVanity+i*common.AddressLength:], addresses1[i][:])
	}

	genesis1 := genSpec.MustCommit(db1)
	genesis2 := genSpec.MustCommit(db2)

	commonBlocks1 := generateChain(t, params.AllCliqueProtocolChanges, genesis1, engine1, db1, commonBlocksCount,
		addresses1, keys, reg1, false)
	commonBlocks2 := generateChain(t, params.AllCliqueProtocolChanges, genesis2, engine2, db2, commonBlocksCount,
		addresses2, keys, reg2, false)

	// Insert common blocks
	chain1, _ := core.NewBlockChain(db1, nil, params.AllCliqueProtocolChanges, engine1, vm.Config{},
		nil, nil)
	defer chain1.Stop()
	if _, err := chain1.InsertChain(commonBlocks1); err != nil {
		t.Fatalf("failed to insert commonBlocks blocks: %v", err)
	}
	chain2, _ := core.NewBlockChain(db2, nil, params.AllCliqueProtocolChanges, engine2, vm.Config{},
		nil, nil)
	defer chain2.Stop()
	if _, err := chain2.InsertChain(commonBlocks2); err != nil {
		t.Fatalf("failed to insert commonBlocks blocks: %v", err)
	}

	// Add new address to the validators list of the second node
	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)
	keys[addr] = key
	addresses2 = append(addresses2, addr)
	sort.Sort(signersAscending(addresses2))

	reg2 = contracts.NewTestModeRegistryWithMocks(
		mocks.NewMockValidators(common.HexToAddress("0x1"), addresses2, nil),
		nil,
	)

	engine2 = New(params.AllCliqueProtocolChanges.Clique, db2, exclusionSetProvider2, reg2)

	// Generate side chain with different validators list
	blocksToGenerate := signersCount * 3
	sideChain := generateChain(t, params.AllCliqueProtocolChanges, chain2.CurrentBlock(), engine2, db2,
		blocksToGenerate, addresses2, keys, reg2, true)

	_, err := chain1.InsertChain(sideChain)
	if err == nil {
		t.Fatalf("shouldn't validate blocks")
	}
	if !errors.Is(err, errUnauthorizedSigner) && !errors.Is(err, errWrongDifficulty) {
		t.Fatalf("failed to insert sidechain: %v", err)
	}
}
