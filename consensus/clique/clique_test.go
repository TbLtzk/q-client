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
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	ethereum "gitlab.com/q-dev/q-client"
	"gitlab.com/q-dev/q-client/accounts/abi"
	"gitlab.com/q-dev/q-client/accounts/abi/bind"
	"gitlab.com/q-dev/q-client/accounts/abi/bind/backends"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/contracts"
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

func TestForkChoice(t *testing.T) {
	var (
		db                = rawdb.NewMemoryDatabase()
		reg               = contracts.NewTestModeRegistry()
		engine            = New(params.AllCliqueProtocolChanges.Clique, db, &NoopExclusionSetProvider{}, reg)
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
		log.Error(signer.String())

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
		engine       = New(params.AllCliqueProtocolChanges.Clique, db, &NoopExclusionSetProvider{}, reg)
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

func TestReorgExclusionSet(t *testing.T) {
	var (
		db1                   = rawdb.NewMemoryDatabase()
		db2                   = rawdb.NewMemoryDatabase()
		exclusionSetProvider1 = new(NoopExclusionSetProvider)
		exclusionSetProvider2 = new(NoopExclusionSetProvider)
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

	commonBlocks1 := generateChain(t, params.AllCliqueProtocolChanges, genesis1, engine1, db1, commonBlocksCount, addresses, keys, reg1, false)
	commonBlocks2 := generateChain(t, params.AllCliqueProtocolChanges, genesis2, engine2, db2, commonBlocksCount, addresses, keys, reg2, false)

	// Insert common blocks
	chain1, _ := core.NewBlockChain(db1, nil, params.AllCliqueProtocolChanges, engine1, vm.Config{}, nil, nil)
	defer chain1.Stop()
	if _, err := chain1.InsertChain(commonBlocks1); err != nil {
		t.Fatalf("failed to insert commonBlocks blocks: %v", err)
	}
	chain2, _ := core.NewBlockChain(db2, nil, params.AllCliqueProtocolChanges, engine2, vm.Config{}, nil, nil)
	defer chain2.Stop()
	if _, err := chain2.InsertChain(commonBlocks2); err != nil {
		t.Fatalf("failed to insert commonBlocks blocks: %v", err)
	}

	blocksToGenerate := signersCount * 3
	sideChain := generateChain(t, params.AllCliqueProtocolChanges, chain2.CurrentBlock(), engine2, db2, blocksToGenerate, addresses, keys, reg2, true)

	exclusionSetProvider1.set = map[common.Address][]common.BlockRange{
		addresses[0]: {{StartAddress: uint64(commonBlocksCount + 1), EndAddress: uint64(commonBlocksCount + blocksToGenerate + 1)}},
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

func TestReorgValidatorsState(t *testing.T) {
	const (
		contractsRegistryABI      = "[{\"type\":\"constructor\",\"stateMutability\":\"payable\",\"payable\":true,\"inputs\":[{\"type\":\"address\",\"name\":\"_logic\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_admin\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"_data\",\"internalType\":\"bytes\"}]},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousAdmin\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"address\",\"name\":\"newAdmin\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"type\":\"address\",\"name\":\"implementation\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"fallback\",\"stateMutability\":\"payable\",\"payable\":true},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"payable\":false,\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"admin\",\"inputs\":[],\"constant\":false},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"payable\":false,\"outputs\":[],\"name\":\"changeAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"newAdmin\",\"internalType\":\"address\"}],\"constant\":false},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"payable\":false,\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"implementation\",\"inputs\":[],\"constant\":false},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"payable\":false,\"outputs\":[],\"name\":\"upgradeTo\",\"inputs\":[{\"type\":\"address\",\"name\":\"newImplementation\",\"internalType\":\"address\"}],\"constant\":false},{\"type\":\"function\",\"stateMutability\":\"payable\",\"payable\":true,\"outputs\":[],\"name\":\"upgradeToAndCall\",\"inputs\":[{\"type\":\"address\",\"name\":\"newImplementation\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}],\"constant\":false}]"
		contractsRegistryCode     = "0x60806040526004361061004a5760003560e01c80633659cfe6146100545780634f1ef286146100a55780635c60da1b1461013e5780638f28397014610195578063f851a440146101e6575b61005261023d565b005b34801561006057600080fd5b506100a36004803603602081101561007757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610257565b005b61013c600480360360408110156100bb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156100f857600080fd5b82018360208201111561010a57600080fd5b8035906020019184600183028401116401000000008311171561012c57600080fd5b90919293919293905050506102ac565b005b34801561014a57600080fd5b506101536103ce565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101a157600080fd5b506101e4600480360360208110156101b857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610426565b005b3480156101f257600080fd5b506101fb61059f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102456105f7565b61025561025061068d565b6106be565b565b61025f6106e4565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102a05761029b81610715565b6102a9565b6102a861023d565b5b50565b6102b46106e4565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103c0576102f083610715565b60008373ffffffffffffffffffffffffffffffffffffffff168383604051808383808284378083019250505092505050600060405180830381855af49150503d806000811461035b576040519150601f19603f3d011682016040523d82523d6000602084013e610360565b606091505b50509050806103ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602381526020018061089e6023913960400191505060405180910390fd5b506103c9565b6103c861023d565b5b505050565b60006103d86106e4565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561041a5761041361068d565b9050610423565b61042261023d565b5b90565b61042e6106e4565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561059357600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156104e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001806108686036913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6105106106e4565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a161058e81610764565b61059c565b61059b61023d565b5b50565b60006105a96106e4565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156105eb576105e46106e4565b90506105f4565b6105f361023d565b5b90565b6105ff6106e4565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610683576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806108366032913960400191505060405180910390fd5b61068b610793565b565b6000807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b9050805491505090565b3660008037600080366000845af43d6000803e80600081146106df573d6000f35b3d6000fd5b6000807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b9050805491505090565b61071e81610795565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b90508181555050565b565b61079e81610822565b6107f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001806108c1603b913960400191505060405180910390fd5b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b90508181555050565b600080823b90506000811191505091905056fe43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e2066726f6d207468652070726f78792061646d696e43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20746865207a65726f206164647265737364656c6567617465642063616c6c206661696c6564206166746572207570677261646543616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a265627a7a723158200cff5bad2f803a2127039eba3a278e1d70096ece8708d9ca347834fd0b404ae864736f6c63430005110032"
		contractsRegistryImplABI  = "[{\"inputs\": [{\"internalType\": \"string\",\"name\": \"_key\",\"type\": \"string\"}],\"name\": \"contains\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"string\",\"name\": \"_key\",\"type\": \"string\"}],\"name\": \"getAddress\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"getContracts\",\"outputs\": [{\"components\": [{\"internalType\": \"string\",\"name\": \"key\",\"type\": \"string\"},{\"internalType\": \"address\",\"name\": \"addr\",\"type\": \"address\"}],\"internalType\": \"struct IContractRegistry.Pair[]\",\"name\": \"\",\"type\": \"tuple[]\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_proxy\",\"type\": \"address\"}],\"name\": \"getImplementation\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"getMaintainers\",\"outputs\": [{\"internalType\": \"address[]\",\"name\": \"\",\"type\": \"address[]\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address[]\",\"name\": \"_maintainersList\",\"type\": \"address[]\"},{\"internalType\": \"string[]\",\"name\": \"_keys\",\"type\": \"string[]\"},{\"internalType\": \"address[]\",\"name\": \"_addresses\",\"type\": \"address[]\"}],\"name\": \"initialize\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"leaveMaintainers\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"string\",\"name\": \"_key\",\"type\": \"string\"}],\"name\": \"mustGetAddress\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"string\",\"name\": \"_key\",\"type\": \"string\"}],\"name\": \"removeKey\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"string[]\",\"name\": \"_keys\",\"type\": \"string[]\"}],\"name\": \"removeKeys\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"string\",\"name\": \"_key\",\"type\": \"string\"},{\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"}],\"name\": \"setAddress\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"string[]\",\"name\": \"_keys\",\"type\": \"string[]\"},{\"internalType\": \"address[]\",\"name\": \"_addresses\",\"type\": \"address[]\"}],\"name\": \"setAddresses\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_maintainer\",\"type\": \"address\"}],\"name\": \"setMaintainer\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_proxy\",\"type\": \"address\"},{\"internalType\": \"address\",\"name\": \"_newImplementation\",\"type\": \"address\"}],\"name\": \"upgradeContract\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"}]"
		contractsRegistryImplCode = "0x60806040523480156200001157600080fd5b5060043610620000c95760003560e01c806287452314620000ce57806302fb7e7314620000e757806313ea5d2914620000fe57806315ac72ca146200011557806336476dbf14620001445780633fb90271146200015b5780636833d54f14620001725780637d69a892146200019a5780639b2ea4bd14620001b1578063bf40fac114620001c8578063c3a2a93a14620001df578063d17f422514620001f8578063d52040331462000211578063d5d33b55146200021b575b600080fd5b620000e5620000df3660046200139a565b62000232565b005b620000e5620000f836600462001561565b62000346565b620000e56200010f366004620015a1565b62000438565b6200012c62000126366004620015a1565b62000546565b6040516200013b9190620015c8565b60405180910390f35b620000e56200015536600462001649565b62000669565b6200012c6200016c36600462001725565b620008e4565b62000189620001833660046200176a565b62000963565b60405190151581526020016200013b565b620000e5620001ab366004620017a2565b620009a2565b620000e5620001c23660046200180c565b62000ad1565b6200012c620001d936600462001725565b62000bbe565b620001e962000bf5565b6040516200013b9190620018c5565b6200020262000dc0565b6040516200013b919062001948565b620000e562000e4a565b620000e56200022c36600462001725565b62001043565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e89062000264903390600401620015c8565b60206040518083038186803b1580156200027d57600080fd5b505afa15801562000292573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002b8919062001997565b620002e05760405162461bcd60e51b8152600401620002d790620019bb565b60405180910390fd5b604051631b2ce7f360e11b81526001600160a01b03831690633659cfe6906200030e908490600401620015c8565b600060405180830381600087803b1580156200032957600080fd5b505af11580156200033e573d6000803e3d6000fd5b505050505050565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e89062000378903390600401620015c8565b60206040518083038186803b1580156200039157600080fd5b505afa158015620003a6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003cc919062001997565b620003eb5760405162461bcd60e51b8152600401620002d790620019bb565b60005b815181101562000434576200041f82828151811062000411576200041162001a18565b602002602001015162001129565b806200042b8162001a44565b915050620003ee565b5050565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e8906200046a903390600401620015c8565b60206040518083038186803b1580156200048357600080fd5b505afa15801562000498573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620004be919062001997565b620004dd5760405162461bcd60e51b8152600401620002d790620019bb565b600354604051639e9405c360e01b81526001600160a01b0390911690639e9405c3906200050f908490600401620015c8565b600060405180830381600087803b1580156200052a57600080fd5b505af11580156200053f573d6000803e3d6000fd5b5050505050565b600354604051630bb7c8fd60e31b81526000916001600160a01b031690635dbe47e89062000579903390600401620015c8565b60206040518083038186803b1580156200059257600080fd5b505afa158015620005a7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620005cd919062001997565b620005ec5760405162461bcd60e51b8152600401620002d790620019bb565b816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156200062857600080fd5b505af11580156200063d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000663919062001a62565b92915050565b600054610100900460ff168062000683575060005460ff16155b620006e85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401620002d7565b600054610100900460ff161580156200070b576000805461ffff19166101011790555b82518251808214620007315760405162461bcd60e51b8152600401620002d79062001a82565b6040516200073f9062001368565b604051809103906000f0801580156200075c573d6000803e3d6000fd5b50600380546001600160a01b0319166001600160a01b0392909216918217905560405163a224cee760e01b815263a224cee7906200079f90899060040162001948565b600060405180830381600087803b158015620007ba57600080fd5b505af1158015620007cf573d6000803e3d6000fd5b5050505084604051620007e29062001376565b620007ee919062001b02565b604051809103906000f0801580156200080b573d6000803e3d6000fd5b50600280546001600160a01b0319166001600160a01b039290921691909117905560005b8451811015620008c8578481815181106200084e576200084e62001a18565b602002602001015160018783815181106200086d576200086d62001a18565b602002602001015160405162000884919062001b68565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b031990921691909117905580620008bf8162001a44565b9150506200082f565b5050508015620008de576000805461ff00191690555b50505050565b60008060018484604051620008fb92919062001b86565b90815260405160209181900382018120546001600160a01b03169250821515916200092b91879187910162001b96565b604051602081830303815290604052906200095b5760405162461bcd60e51b8152600401620002d7919062001bff565b509392505050565b6000806001600160a01b031660018360405162000981919062001b68565b908152604051908190036020019020546001600160a01b0316141592915050565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e890620009d4903390600401620015c8565b60206040518083038186803b158015620009ed57600080fd5b505afa15801562000a02573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000a28919062001997565b62000a475760405162461bcd60e51b8152600401620002d790620019bb565b8151815180821462000a6d5760405162461bcd60e51b8152600401620002d79062001a82565b60005b84518110156200053f5762000abe85828151811062000a935762000a9362001a18565b602002602001015185838151811062000ab05762000ab062001a18565b60200260200101516200120a565b62000ac98162001a44565b905062000a70565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e89062000b03903390600401620015c8565b60206040518083038186803b15801562000b1c57600080fd5b505afa15801562000b31573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000b57919062001997565b62000b765760405162461bcd60e51b8152600401620002d790620019bb565b62000bb983838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508592506200120a915050565b505050565b60006001838360405162000bd492919062001b86565b908152604051908190036020019020546001600160a01b0316905092915050565b60606000600260009054906101000a90046001600160a01b03166001600160a01b0316638d9a5ad06040518163ffffffff1660e01b815260040160006040518083038186803b15801562000c4857600080fd5b505afa15801562000c5d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000c87919081019062001c14565b9050600081516001600160401b0381111562000ca75762000ca7620013d8565b60405190808252806020026020018201604052801562000cef57816020015b60408051808201909152606081526000602082015281526020019060019003908162000cc65790505b50905060005b82518160ff16101562000db9576040518060400160405280848360ff168151811062000d255762000d2562001a18565b602002602001015181526020016001858460ff168151811062000d4c5762000d4c62001a18565b602002602001015160405162000d63919062001b68565b908152604051908190036020019020546001600160a01b031690528251839060ff841690811062000d985762000d9862001a18565b6020026020010181905250808062000db09062001d11565b91505062000cf5565b5092915050565b600354604080516351cfd60960e11b815290516060926001600160a01b03169163a39fac12916004808301926000929190829003018186803b15801562000e0657600080fd5b505afa15801562000e1b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000e45919081019062001d34565b905090565b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e89062000e7c903390600401620015c8565b60206040518083038186803b15801562000e9557600080fd5b505afa15801562000eaa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000ed0919062001997565b1562000fe1576003546040805163949d225d60e01b815290516001926001600160a01b03169163949d225d916004808301926020929190829003018186803b15801562000f1c57600080fd5b505afa15801562000f31573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000f57919062001dd8565b1162000fe15760405162461bcd60e51b815260206004820152604c60248201527f5b5145432d3033343030325d2d43616e6e6f74206c6561766520746865206d6160448201527f696e7461696e657273206c6973742c20796f752061726520746865206f6e6c7960648201526b1036b0b4b73a30b4b732b91760a11b608482015260a401620002d7565b600354604051636196c02d60e11b81526001600160a01b039091169063c32d805a9062001013903390600401620015c8565b600060405180830381600087803b1580156200102e57600080fd5b505af1158015620008de573d6000803e3d6000fd5b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e89062001075903390600401620015c8565b60206040518083038186803b1580156200108e57600080fd5b505afa158015620010a3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620010c9919062001997565b620010e85760405162461bcd60e51b8152600401620002d790620019bb565b6200043482828080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506200112992505050565b620011348162000963565b1562001207576002546040516380599e4b60e01b81526001600160a01b03909116906380599e4b906200116c90849060040162001bff565b602060405180830381600087803b1580156200118757600080fd5b505af11580156200119c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620011c2919062001997565b506000600182604051620011d7919062001b68565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b03199092169190911790555b50565b6001600160a01b038116620012885760405162461bcd60e51b815260206004820152603e60248201527f5b5145432d3033343030305d2d496e76616c696420616464726573732076616c60448201527f75652c206661696c656420746f207365742074686520616464726573732e00006064820152608401620002d7565b620012938262000963565b6200132257600254604051632c323e7760e21b81526001600160a01b039091169063b0c8f9dc90620012ca90859060040162001bff565b602060405180830381600087803b158015620012e557600080fd5b505af1158015620012fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062001320919062001997565b505b8060018360405162001335919062001b68565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b03199092169190911790555050565b610d678062001df383390190565b6114598062002b5a83390190565b6001600160a01b03811681146200120757600080fd5b60008060408385031215620013ae57600080fd5b8235620013bb8162001384565b91506020830135620013cd8162001384565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620014195762001419620013d8565b604052919050565b60006001600160401b038211156200143d576200143d620013d8565b5060051b60200190565b60006001600160401b03821115620014635762001463620013d8565b50601f01601f191660200190565b600082601f8301126200148357600080fd5b81356200149a620014948262001447565b620013ee565b818152846020838601011115620014b057600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112620014df57600080fd5b81356020620014f2620014948362001421565b82815260059290921b840181019181810190868411156200151257600080fd5b8286015b84811015620015565780356001600160401b03811115620015375760008081fd5b620015478986838b010162001471565b84525091830191830162001516565b509695505050505050565b6000602082840312156200157457600080fd5b81356001600160401b038111156200158b57600080fd5b6200159984828501620014cd565b949350505050565b600060208284031215620015b457600080fd5b8135620015c18162001384565b9392505050565b6001600160a01b0391909116815260200190565b600082601f830112620015ee57600080fd5b8135602062001601620014948362001421565b82815260059290921b840181019181810190868411156200162157600080fd5b8286015b84811015620015565780356200163b8162001384565b835291830191830162001625565b6000806000606084860312156200165f57600080fd5b83356001600160401b03808211156200167757600080fd5b6200168587838801620015dc565b945060208601359150808211156200169c57600080fd5b620016aa87838801620014cd565b93506040860135915080821115620016c157600080fd5b50620016d086828701620015dc565b9150509250925092565b60008083601f840112620016ed57600080fd5b5081356001600160401b038111156200170557600080fd5b6020830191508360208285010111156200171e57600080fd5b9250929050565b600080602083850312156200173957600080fd5b82356001600160401b038111156200175057600080fd5b6200175e85828601620016da565b90969095509350505050565b6000602082840312156200177d57600080fd5b81356001600160401b038111156200179457600080fd5b620015998482850162001471565b60008060408385031215620017b657600080fd5b82356001600160401b0380821115620017ce57600080fd5b620017dc86838701620014cd565b93506020850135915080821115620017f357600080fd5b506200180285828601620015dc565b9150509250929050565b6000806000604084860312156200182257600080fd5b83356001600160401b038111156200183957600080fd5b6200184786828701620016da565b90945092505060208401356200185d8162001384565b809150509250925092565b60005b83811015620018855781810151838201526020016200186b565b83811115620008de5750506000910152565b60008151808452620018b181602086016020860162001868565b601f01601f19169290920160200192915050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156200193a57888303603f1901855281518051878552620019138886018262001897565b918901516001600160a01b03169489019490945294870194925090860190600101620018ec565b509098975050505050505050565b6020808252825182820181905260009190848201906040850190845b818110156200198b5783516001600160a01b03168352928401929184019160010162001964565b50909695505050505050565b600060208284031215620019aa57600080fd5b81518015158114620015c157600080fd5b6020808252603e908201527f5b5145432d3033343030345d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c79206d61696e7461696e6572732068617665206163636573732e0000606082015260800190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141562001a5b5762001a5b62001a2e565b5060010190565b60006020828403121562001a7557600080fd5b8151620015c18162001384565b6020808252605a908201527f5b5145432d3033343030315d2d546865206e756d626572206f66206b6579732060408201527f616e64206164647265737365732073686f756c64206265207468652073616d6560608201527916103330b4b632b2103a379039b2ba1030b2323932b9b9b2b99760311b608082015260a00190565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101562001b5b57603f1988860301845262001b4885835162001897565b9450928501929085019060010162001b29565b5092979650505050505050565b6000825162001b7c81846020870162001868565b9190910192915050565b8183823760009101908152919050565b7002da8a2a19698199a181819ae96aa34329607d1b8152818360118301377f206b657920646f6573206e6f742065786973742c206661696c656420746f2067910160118101919091526e32ba103a34329030b2323932b9b99760891b6031820152604001919050565b602081526000620015c1602083018462001897565b6000602080838503121562001c2857600080fd5b82516001600160401b038082111562001c4057600080fd5b818501915085601f83011262001c5557600080fd5b815162001c66620014948262001421565b81815260059190911b8301840190848101908883111562001c8657600080fd5b8585015b8381101562001d045780518581111562001ca45760008081fd5b8601603f81018b1362001cb75760008081fd5b87810151604062001ccc620014948362001447565b8281528d8284860101111562001ce25760008081fd5b62001cf3838c830184870162001868565b865250505091860191860162001c8a565b5098975050505050505050565b600060ff821660ff81141562001d2b5762001d2b62001a2e565b60010192915050565b6000602080838503121562001d4857600080fd5b82516001600160401b0381111562001d5f57600080fd5b8301601f8101851362001d7157600080fd5b805162001d82620014948262001421565b81815260059190911b8201830190838101908783111562001da257600080fd5b928401925b8284101562001dcd57835162001dbd8162001384565b8252928401929084019062001da7565b979650505050505050565b60006020828403121562001deb57600080fd5b505191905056fe608060405234801561001057600080fd5b50610d47806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a45760003560e01c80630a3b0a4f146100a957806329092d0e146100d157806352efea6e146100e45780635dbe47e8146100ec578063715018a6146100ff5780638da5cb5b14610109578063949d225d146101295780639e9405c31461013a578063a224cee71461014d578063a39fac1214610160578063c32d805a14610175578063f2fde38b14610188575b600080fd5b6100bc6100b7366004610aba565b61019b565b60405190151581526020015b60405180910390f35b6100bc6100df366004610aba565b6101fd565b6100bc61024c565b6100bc6100fa366004610aba565b6102eb565b610107610308565b005b610111610343565b6040516001600160a01b0390911681526020016100c8565b6066546040519081526020016100c8565b610107610148366004610aba565b610352565b61010761015b366004610af2565b610410565b610168610534565b6040516100c89190610bb7565b610107610183366004610aba565b610596565b610107610196366004610aba565b61061d565b6000336101a6610343565b6001600160a01b0316146101d55760405162461bcd60e51b81526004016101cc90610c04565b60405180910390fd5b6101de826102eb565b156101eb57506000919050565b6101f4826106ba565b5060015b919050565b600033610208610343565b6001600160a01b03161461022e5760405162461bcd60e51b81526004016101cc90610c04565b610237826102eb565b61024357506000919050565b6101f48261071d565b600033610257610343565b6001600160a01b03161461027d5760405162461bcd60e51b81526004016101cc90610c04565b60005b6066548110156102d85760656000606683815481106102a1576102a1610c39565b60009182526020808320909101546001600160a01b03168352820192909252604001812055806102d081610c65565b915050610280565b506102e560666000610a71565b50600190565b6001600160a01b0316600090815260656020526040902054151590565b33610311610343565b6001600160a01b0316146103375760405162461bcd60e51b81526004016101cc90610c04565b6103416000610883565b565b6033546001600160a01b031690565b61035b8161019b565b61040d5760405162461bcd60e51b815260206004820152607160248201527f5b5145432d3033353030325d2d54686520616464726573732068617320616c7260448201527f65616479206265656e20616464656420746f207468652073746f726167652c2060648201527f6661696c656420746f2061646420746865206164647265737320746f207468656084820152701030b2323932b9b99039ba37b930b3b29760791b60a482015260c4016101cc565b50565b600054610100900460ff1680610429575060005460ff16155b6104455760405162461bcd60e51b81526004016101cc90610c80565b600054610100900460ff16158015610467576000805461ffff19166101011790555b60005b825181101561051557606683828151811061048757610487610c39565b6020908102919091018101518254600181018455600093845291832090910180546001600160a01b0319166001600160a01b03909216919091179055606654845190916065918690859081106104df576104df610c39565b6020908102919091018101516001600160a01b03168252810191909152604001600020558061050d81610c65565b91505061046a565b5061051e6108d5565b8015610530576000805461ff00191690555b5050565b6060606680548060200260200160405190810160405280929190818152602001828054801561058c57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161056e575b5050505050905090565b61059f816101fd565b61040d5760405162461bcd60e51b815260206004820152604360248201527f5b5145432d3033353030305d2d4661696c656420746f2072656d6f766520746860448201527f6520616464726573732066726f6d2074686520616464726573732073746f726160648201526233b29760e91b608482015260a4016101cc565b33610626610343565b6001600160a01b03161461064c5760405162461bcd60e51b81526004016101cc90610c04565b6001600160a01b0381166106b15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101cc565b61040d81610883565b606680546001810182557f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943540180546001600160a01b0319166001600160a01b03841690811790915590546000918252606560205260409091205561040d81610950565b6001600160a01b0381166000908152606560209081526040808320815192830190915254815260665490919061075590600190610cce565b905060006066828154811061076c5761076c610c39565b60009182526020909120015483516001600160a01b03909116915061079390600190610cce565b82146108065782516001600160a01b038216600090815260656020526040902055825181906066906107c790600190610cce565b815481106107d7576107d7610c39565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b606680548061081757610817610ce5565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038616825260659052604081205561085b84610950565b836001600160a01b0316816001600160a01b03161461087d5761087d81610950565b50505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16806108ee575060005460ff16155b61090a5760405162461bcd60e51b81526004016101cc90610c80565b600054610100900460ff1615801561092c576000805461ffff19166101011790555b6109346109a7565b61093c610a11565b801561040d576000805461ff001916905550565b6001600160a01b03811660009081526065602052604090205460665481111561097b5761097b610cfb565b610984826102eb565b15610999576000811161053057610530610cfb565b801561053057610530610cfb565b600054610100900460ff16806109c0575060005460ff16155b6109dc5760405162461bcd60e51b81526004016101cc90610c80565b600054610100900460ff1615801561093c576000805461ffff1916610101179055801561040d576000805461ff001916905550565b600054610100900460ff1680610a2a575060005460ff16155b610a465760405162461bcd60e51b81526004016101cc90610c80565b600054610100900460ff16158015610a68576000805461ffff19166101011790555b61093c33610883565b508054600082559060005260206000209081019061040d91905b80821115610a9f5760008155600101610a8b565b5090565b80356001600160a01b03811681146101f857600080fd5b600060208284031215610acc57600080fd5b610ad582610aa3565b9392505050565b634e487b7160e01b600052604160045260246000fd5b60006020808385031215610b0557600080fd5b823567ffffffffffffffff80821115610b1d57600080fd5b818501915085601f830112610b3157600080fd5b813581811115610b4357610b43610adc565b8060051b604051601f19603f83011681018181108582111715610b6857610b68610adc565b604052918252848201925083810185019188831115610b8657600080fd5b938501935b82851015610bab57610b9c85610aa3565b84529385019392850192610b8b565b98975050505050505050565b6020808252825182820181905260009190848201906040850190845b81811015610bf85783516001600160a01b031683529284019291840191600101610bd3565b50909695505050505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415610c7957610c79610c4f565b5060010190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b600082821015610ce057610ce0610c4f565b500390565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052600160045260246000fdfea2646970667358221220977078b179f0297982ef3002bc7834252cf2dcba52f17a39a6edce4e4ceafb3b64736f6c6343000809003360806040523480156200001157600080fd5b50604051620014593803806200145983398101604081905262000034916200026b565b6200003f33620000f9565b60005b8151811015620000f157600282828151811062000063576200006362000395565b6020908102919091018101518254600181018455600093845292829020815162000097949190910192919091019062000149565b506002805490506001838381518110620000b557620000b562000395565b6020026020010151604051620000cc9190620003ab565b9081526040519081900360200190205580620000e881620003c9565b91505062000042565b505062000430565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b8280546200015790620003f3565b90600052602060002090601f0160209004810192826200017b5760008555620001c6565b82601f106200019657805160ff1916838001178555620001c6565b82800160010185558215620001c6579182015b82811115620001c6578251825591602001919060010190620001a9565b50620001d4929150620001d8565b5090565b5b80821115620001d45760008155600101620001d9565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620002305762000230620001ef565b604052919050565b60005b83811015620002555781810151838201526020016200023b565b8381111562000265576000848401525b50505050565b600060208083850312156200027f57600080fd5b82516001600160401b03808211156200029757600080fd5b8185019150601f8681840112620002ad57600080fd5b825182811115620002c257620002c2620001ef565b8060051b620002d386820162000205565b918252848101860191868101908a841115620002ee57600080fd5b87870192505b8383101562000387578251868111156200030e5760008081fd5b8701603f81018c13620003215760008081fd5b888101516040888211156200033a576200033a620001ef565b6200034d828901601f19168c0162000205565b8281528e82848601011115620003635760008081fd5b62000374838d830184870162000238565b85525050509187019190870190620002f4565b9a9950505050505050505050565b634e487b7160e01b600052603260045260246000fd5b60008251620003bf81846020870162000238565b9190910192915050565b6000600019821415620003ec57634e487b7160e01b600052601160045260246000fd5b5060010190565b600181811c908216806200040857607f821691505b602082108114156200042a57634e487b7160e01b600052602260045260246000fd5b50919050565b61101980620004406000396000f3fe608060405234801561001057600080fd5b50600436106100995760003560e01c8063374155e51461009e57806352efea6e146100b35780636833d54f146100d0578063715018a6146100e357806380599e4b146100eb5780638d9a5ad0146100fe5780638da5cb5b14610113578063949d225d146101335780639b1574eb14610144578063b0c8f9dc14610157578063f2fde38b1461016a575b600080fd5b6100b16100ac366004610c12565b61017d565b005b6100bb61023f565b60405190151581526020015b60405180910390f35b6100bb6100de366004610c9a565b6102e1565b6100b161030b565b6100bb6100f9366004610c12565b610346565b61010661040b565b6040516100c79190610d77565b61011b6104e4565b6040516001600160a01b0390911681526020016100c7565b6002546040519081526020016100c7565b6100b1610152366004610c12565b6104f3565b6100bb610165366004610c12565b6105db565b6100b1610178366004610df1565b610662565b336101866104e4565b6001600160a01b0316146101b55760405162461bcd60e51b81526004016101ac90610e21565b60405180910390fd5b6101bf8282610346565b61023b5760405162461bcd60e51b815260206004820152604160248201527f5b5145432d3033383030305d2d4661696c656420746f2072656d6f766520746860448201527f6520737472696e672066726f6d2074686520737472696e672073746f726167656064820152601760f91b608482015260a4016101ac565b5050565b60003361024a6104e4565b6001600160a01b0316146102705760405162461bcd60e51b81526004016101ac90610e21565b60005b6002548110156102ce5760016002828154811061029257610292610e56565b906000526020600020016040516102a99190610ea7565b90815260405190819003602001902060009055806102c681610f59565b915050610273565b506102db60026000610a90565b50600190565b6000806001836040516102f49190610f74565b908152604051908190036020019020541192915050565b336103146104e4565b6001600160a01b03161461033a5760405162461bcd60e51b81526004016101ac90610e21565b6103446000610702565b565b6000336103516104e4565b6001600160a01b0316146103775760405162461bcd60e51b81526004016101ac90610e21565b6103b683838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506102e192505050565b6103c257506000610405565b61040183838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061075292505050565b5060015b92915050565b60606002805480602002602001604051908101604052809291908181526020016000905b828210156104db57838290600052602060002001805461044e90610e6c565b80601f016020809104026020016040519081016040528092919081815260200182805461047a90610e6c565b80156104c75780601f1061049c576101008083540402835291602001916104c7565b820191906000526020600020905b8154815290600101906020018083116104aa57829003601f168201915b50505050508152602001906001019061042f565b50505050905090565b6000546001600160a01b031690565b336104fc6104e4565b6001600160a01b0316146105225760405162461bcd60e51b81526004016101ac90610e21565b61052c82826105db565b61023b5760405162461bcd60e51b815260206004820152606e60248201527f5b5145432d3033383030315d2d54686520737472696e672068617320616c726560448201527f616479206265656e20616464656420746f207468652073746f726167652c206660648201527f61696c656420746f206164642074686520737472696e6720746f20746865207360848201526d3a3934b7339039ba37b930b3b29760911b60a482015260c4016101ac565b6000336105e66104e4565b6001600160a01b03161461060c5760405162461bcd60e51b81526004016101ac90610e21565b61064b83838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506102e192505050565b1561065857506000610405565b6104018383610934565b3361066b6104e4565b6001600160a01b0316146106915760405162461bcd60e51b81526004016101ac90610e21565b6001600160a01b0381166106f65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101ac565b6106ff81610702565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006001826040516107649190610f74565b90815260408051918290036020908101832090830190915254815260025490915060009061079490600190610f90565b90506000600282815481106107ab576107ab610e56565b9060005260206000200180546107c090610e6c565b80601f01602080910402602001604051908101604052809291908181526020018280546107ec90610e6c565b80156108395780601f1061080e57610100808354040283529160200191610839565b820191906000526020600020905b81548152906001019060200180831161081c57829003601f168201915b50505050509050600183600001516108519190610f90565b82146108bc57825160405160019061086a908490610f74565b908152604051908190036020019020558251819060029061088d90600190610f90565b8154811061089d5761089d610e56565b9060005260206000200190805190602001906108ba929190610aae565b505b60028054806108cd576108cd610fa7565b6001900381819060005260206000200160006108e99190610b32565b90556001846040516108fb9190610f74565b90815260405190819003602001902060009055610917846109d4565b6109218185610a37565b61092e5761092e816109d4565b50505050565b60028054600181018255600091909152610971907f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace018383610b6c565b506002546040516001906109889085908590610fbd565b90815260408051918290036020908101832093909355601f8401839004830282018301905282815261023b9184908490819084018382808284376000920191909152506109d492505050565b60006001826040516109e69190610f74565b90815260405190819003602001902054600254909150811115610a0b57610a0b610fcd565b610a14826102e1565b15610a29576000811161023b5761023b610fcd565b801561023b5761023b610fcd565b600081604051602001610a4a9190610f74565b6040516020818303038152906040528051906020012083604051602001610a719190610f74565b6040516020818303038152906040528051906020012014905092915050565b50805460008255906000526020600020908101906106ff9190610be0565b828054610aba90610e6c565b90600052602060002090601f016020900481019282610adc5760008555610b22565b82601f10610af557805160ff1916838001178555610b22565b82800160010185558215610b22579182015b82811115610b22578251825591602001919060010190610b07565b50610b2e929150610bfd565b5090565b508054610b3e90610e6c565b6000825580601f10610b4e575050565b601f0160209004906000526020600020908101906106ff9190610bfd565b828054610b7890610e6c565b90600052602060002090601f016020900481019282610b9a5760008555610b22565b82601f10610bb35782800160ff19823516178555610b22565b82800160010185558215610b22579182015b82811115610b22578235825591602001919060010190610bc5565b80821115610b2e576000610bf48282610b32565b50600101610be0565b5b80821115610b2e5760008155600101610bfe565b60008060208385031215610c2557600080fd5b823567ffffffffffffffff80821115610c3d57600080fd5b818501915085601f830112610c5157600080fd5b813581811115610c6057600080fd5b866020828501011115610c7257600080fd5b60209290920196919550909350505050565b634e487b7160e01b600052604160045260246000fd5b600060208284031215610cac57600080fd5b813567ffffffffffffffff80821115610cc457600080fd5b818401915084601f830112610cd857600080fd5b813581811115610cea57610cea610c84565b604051601f8201601f19908116603f01168101908382118183101715610d1257610d12610c84565b81604052828152876020848701011115610d2b57600080fd5b826020860160208301376000928101602001929092525095945050505050565b60005b83811015610d66578181015183820152602001610d4e565b8381111561092e5750506000910152565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610de457878503603f1901845281518051808752610dc5818989018a8501610d4b565b601f01601f191695909501860194509285019290850190600101610d9e565b5092979650505050505050565b600060208284031215610e0357600080fd5b81356001600160a01b0381168114610e1a57600080fd5b9392505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052603260045260246000fd5b600181811c90821680610e8057607f821691505b60208210811415610ea157634e487b7160e01b600052602260045260246000fd5b50919050565b600080835481600182811c915080831680610ec357607f831692505b6020808410821415610ee357634e487b7160e01b86526022600452602486fd5b818015610ef75760018114610f0857610f35565b60ff19861689528489019650610f35565b60008a81526020902060005b86811015610f2d5781548b820152908501908301610f14565b505084890196505b509498975050505050505050565b634e487b7160e01b600052601160045260246000fd5b6000600019821415610f6d57610f6d610f43565b5060010190565b60008251610f86818460208701610d4b565b9190910192915050565b600082821015610fa257610fa2610f43565b500390565b634e487b7160e01b600052603160045260246000fd5b8183823760009101908152919050565b634e487b7160e01b600052600160045260246000fdfea2646970667358221220bedc1fafeb4f6ebe13c46bb347a85719070deccc3a4235ff7a4332b946d9e88764736f6c63430008090033a26469706673582212203d3b65dc5a86d6b3743a26f30c633607d1da88a6ebee630e1d90ab7acbc176a464736f6c63430008090033"
		validatorsABI             = "[{\"type\":\"constructor\",\"stateMutability\":\"payable\",\"inputs\":[{\"type\":\"address\",\"name\":\"_logic\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"admin_\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"_data\",\"internalType\":\"bytes\"}]},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousAdmin\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"address\",\"name\":\"newAdmin\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"type\":\"address\",\"name\":\"beacon\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"type\":\"address\",\"name\":\"implementation\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"address\",\"name\":\"admin_\",\"internalType\":\"address\"}],\"name\":\"admin\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"changeAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"newAdmin\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"address\",\"name\":\"implementation_\",\"internalType\":\"address\"}],\"name\":\"implementation\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"upgradeTo\",\"inputs\":[{\"type\":\"address\",\"name\":\"newImplementation\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"upgradeToAndCall\",\"inputs\":[{\"type\":\"address\",\"name\":\"newImplementation\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]},{\"type\":\"receive\",\"stateMutability\":\"payable\"}]"
		validatorsCode            = "0x60806040526004361061004e5760003560e01c80633659cfe6146100655780634f1ef286146100855780635c60da1b146100985780638f283970146100c9578063f851a440146100e95761005d565b3661005d5761005b6100fe565b005b61005b6100fe565b34801561007157600080fd5b5061005b6100803660046106d6565b610118565b61005b6100933660046106f1565b61015f565b3480156100a457600080fd5b506100ad6101d0565b6040516001600160a01b03909116815260200160405180910390f35b3480156100d557600080fd5b5061005b6100e43660046106d6565b61020b565b3480156100f557600080fd5b506100ad610235565b610106610292565b610116610111610331565b61033b565b565b61012061035f565b6001600160a01b0316336001600160a01b031614156101575761015481604051806020016040528060008152506000610392565b50565b6101546100fe565b61016761035f565b6001600160a01b0316336001600160a01b031614156101c8576101c38383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525060019250610392915050565b505050565b6101c36100fe565b60006101da61035f565b6001600160a01b0316336001600160a01b03161415610200576101fb610331565b905090565b6102086100fe565b90565b61021361035f565b6001600160a01b0316336001600160a01b0316141561015757610154816103bd565b600061023f61035f565b6001600160a01b0316336001600160a01b03161415610200576101fb61035f565b606061028583836040518060600160405280602781526020016107f060279139610411565b9392505050565b3b151590565b61029a61035f565b6001600160a01b0316336001600160a01b031614156101165760405162461bcd60e51b815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f78792074617267606482015261195d60f21b608482015260a4015b60405180910390fd5b60006101fb6104e5565b3660008037600080366000845af43d6000803e80801561035a573d6000f35b3d6000fd5b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b546001600160a01b0316919050565b61039b8361050d565b6000825111806103a85750805b156101c3576103b78383610260565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6103e661035f565b604080516001600160a01b03928316815291841660208301520160405180910390a16101548161054d565b6060833b6104705760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610328565b600080856001600160a01b03168560405161048b91906107a0565b600060405180830381855af49150503d80600081146104c6576040519150601f19603f3d011682016040523d82523d6000602084013e6104cb565b606091505b50915091506104db8282866105f6565b9695505050505050565b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc610383565b6105168161062f565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6001600160a01b0381166105b25760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b6064820152608401610328565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80546001600160a01b0319166001600160a01b039290921691909117905550565b60608315610605575081610285565b8251156106155782518084602001fd5b8160405162461bcd60e51b815260040161032891906107bc565b803b6106935760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610328565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6105d5565b80356001600160a01b03811681146106d157600080fd5b919050565b6000602082840312156106e857600080fd5b610285826106ba565b60008060006040848603121561070657600080fd5b61070f846106ba565b9250602084013567ffffffffffffffff8082111561072c57600080fd5b818601915086601f83011261074057600080fd5b81358181111561074f57600080fd5b87602082850101111561076157600080fd5b6020830194508093505050509250925092565b60005b8381101561078f578181015183820152602001610777565b838111156103b75750506000910152565b600082516107b2818460208701610774565b9190910192915050565b60208152600082518060208401526107db816040850160208701610774565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212200327b558076a4f0859709b44583ffd58994d89156734032df586bf934d04d13264736f6c63430008090033"
		validatorsImplABI         = "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"NewTimeLock\",\"inputs\":[{\"type\":\"address\",\"name\":\"_account\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"tuple\",\"name\":\"_timelock\",\"internalType\":\"struct TimeLockEntry\",\"indexed\":false,\"components\":[{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"releaseStart\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"releaseEnd\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Purged\",\"inputs\":[{\"type\":\"address\",\"name\":\"_account\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"tuple\",\"name\":\"_timelock\",\"internalType\":\"struct TimeLockEntry\",\"indexed\":false,\"components\":[{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"releaseStart\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"releaseEnd\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorDroppedFromShortList\",\"inputs\":[{\"type\":\"address\",\"name\":\"_validator\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorEnteredShortList\",\"inputs\":[{\"type\":\"address\",\"name\":\"_validator\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"addSlashingProposal\",\"inputs\":[{\"type\":\"address\",\"name\":\"_validator\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_slashingProposalId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"announceWithdrawal\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"applySlashing\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_proposalId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_validator\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_amountToSlash\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"calculatePendingSlashingsAmount\",\"inputs\":[{\"type\":\"address\",\"name\":\"_validator\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_endTime\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"commitStake\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"depositOnBehalfOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"_account\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_releaseStart\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_releaseEnd\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"enterShortList\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getAccountableSelfStake\",\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getAccountableTotalStake\",\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"struct VotingLockInfo\",\"components\":[{\"type\":\"uint256\",\"name\":\"lockedAmount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"lockedUntil\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"pendingUnlockAmount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"pendingUnlockTime\",\"internalType\":\"uint256\"}]}],\"name\":\"getLockInfo\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getMinimumBalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"_account\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_timestamp\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getSelfStake\",\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getShortListMaxLength\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256[]\",\"name\":\"\",\"internalType\":\"uint256[]\"}],\"name\":\"getSlashingProposalIds\",\"inputs\":[{\"type\":\"address\",\"name\":\"_validator\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple[]\",\"name\":\"\",\"internalType\":\"struct TimeLockEntry[]\",\"components\":[{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"releaseStart\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"releaseEnd\",\"internalType\":\"uint256\"}]}],\"name\":\"getTimeLocks\",\"inputs\":[{\"type\":\"address\",\"name\":\"_account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getValidatorDelegatedStake\",\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple[]\",\"name\":\"\",\"internalType\":\"struct Validators.ValidatorAmount[]\",\"components\":[{\"type\":\"address\",\"name\":\"validator\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]}],\"name\":\"getValidatorShortList\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getValidatorTotalStake\",\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address[]\",\"name\":\"\",\"internalType\":\"address[]\"}],\"name\":\"getValidatorsList\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"address\",\"name\":\"_registry\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_AddressStorageStakesSorted\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_AddressStorageStakes\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isInLongList\",\"inputs\":[{\"type\":\"address\",\"name\":\"_userAddr\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isInShortList\",\"inputs\":[{\"type\":\"address\",\"name\":\"_validator\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contract AddressStorageStakes\"}],\"name\":\"longList\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"makeSnapshot\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"purgePendingSlashings\",\"inputs\":[{\"type\":\"address\",\"name\":\"_validator\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"purgeTimeLocks\",\"inputs\":[{\"type\":\"address\",\"name\":\"_account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"refreshDelegatedStake\",\"inputs\":[{\"type\":\"address\",\"name\":\"_delegateTo\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"withdrawal\",\"internalType\":\"struct Validators.Withdrawal\",\"components\":[{\"type\":\"uint256\",\"name\":\"endTime\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]}],\"name\":\"validatorsInfos\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"withdraw\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_payTo\",\"internalType\":\"address payable\"}]}]"
		validatorsImplCode        = "0x6080604052600436106101615760003560e01c8063afb60cff116100c7578063afb60cff14610353578063b0426a3f14610373578063b3b22163146103bb578063b9a55be9146103d0578063ba7e3128146103f0578063c0c53b8b14610410578063c0d94b8b14610430578063cad6384314610450578063e7522fc01461047d578063f0e810241461049d578063f27f73c9146104ca578063f2f4a2f1146104ea578063f556347c1461050a578063f7e402851461052a578063fece707d1461055757600080fd5b8062f714ce146101665780630f5b8db31461019b57806310756f7e146101bd57806319d18521146101d25780632aa8e06d146101f55780632b0235771461021557806335355289146102355780633c80da74146102575780633d620628146102b657806373983f43146102d8578063778832d0146102eb57806385549e021461030b57806393cadb2a146103135780639d67164814610333575b600080fd5b34801561017257600080fd5b50610186610181366004613c83565b610577565b60405190151581526020015b60405180910390f35b3480156101a757600080fd5b506101b0610c85565b6040516101929190613cb3565b3480156101c957600080fd5b50610186610ea1565b3480156101de57600080fd5b506101e7611391565b604051908152602001610192565b34801561020157600080fd5b506101e7610210366004613d0b565b61151a565b34801561022157600080fd5b50610186610230366004613d28565b6115de565b34801561024157600080fd5b50610255610250366004613d0b565b611ae9565b005b34801561026357600080fd5b5061029b610272366004613d0b565b600760209081526000918252604091829020825180840190935280548352600101549082015281565b60408051825181526020928301519281019290925201610192565b3480156102c257600080fd5b506102cb611cbb565b6040516101929190613d60565b6102556102e6366004613dad565b611d1d565b3480156102f757600080fd5b50610255610306366004613de2565b611fc6565b61025561202d565b34801561031f57600080fd5b506101e761032e366004613de2565b612039565b34801561033f57600080fd5b5061018661034e366004613d0b565b61204c565b34801561035f57600080fd5b5061025561036e366004613d0b565b6120ce565b34801561037f57600080fd5b506103886120da565b60405161019291908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b3480156103c757600080fd5b506101866120ef565b3480156103dc57600080fd5b506101e76103eb366004613d0b565b612305565b3480156103fc57600080fd5b5061025561040b366004613e0e565b612393565b34801561041c57600080fd5b5061025561042b366004613e27565b61269a565b34801561043c57600080fd5b506101e761044b366004613d0b565b61278c565b34801561045c57600080fd5b50600354610470906001600160a01b031681565b6040516101929190613e72565b34801561048957600080fd5b506101e7610498366004613de2565b61283c565b3480156104a957600080fd5b506104bd6104b8366004613d0b565b612912565b6040516101929190613e86565b3480156104d657600080fd5b506102556104e5366004613d0b565b6129a5565b3480156104f657600080fd5b50610186610505366004613d0b565b612c4a565b34801561051657600080fd5b506101e7610525366004613d0b565b612c7b565b34801561053657600080fd5b5061054a610545366004613d0b565b612ecf565b6040516101929190613edc565b34801561056357600080fd5b506101e7610572366004613d0b565b612f3e565b3360009081526007602090815260408083208151808301909252805480835260019091015492820192909252904210156106395760405162461bcd60e51b815260206004820152605260248201527f5b5145432d3030353030375d2d4e6f7420656e6f7567682074696d652068617360448201527f20656c61707365642073696e63652074686520616e6e6f756e63656d656e742060648201527137b3103a3432903bb4ba34323930bbb0b61760711b608482015260a4015b60405180910390fd5b83816020015110156106b35760405162461bcd60e51b815260206004820152603c60248201527f5b5145432d3030353030385d2d43616e6e6f74207769746864726177206d6f7260448201527f65207468616e2074686520616e6e6f756e63656420616d6f756e742e000000006064820152608401610630565b60006106c3338360000151612fcc565b905060006106d033612f3e565b905060006106de334261283c565b9050806106eb8884613f2a565b10156107535760405162461bcd60e51b815260206004820152603160248201527f5b5145432d3030353031355d2d5769746864726177616c2070726576656e74656044820152703210313c903a34b6b2b637b1b59439949760791b6064820152608401610630565b8261075e8884613f2a565b10156107e05760405162461bcd60e51b815260206004820152604560248201527f5b5145432d3030353030395d2d546865207769746864726177616c206973206260448201527f6c6f636b65642062792070656e64696e6720736c617368696e672070726f706f60648201526439b0b6399760d91b608482015260a401610630565b6107e9876131a3565b8684602001516107f99190613f2a565b6020858101918252336000818152600790925260409182902087518155925160019093019290925560035490516304dffdc160e31b81526001600160a01b03909116916326ffee089161085191908b90600401613f41565b602060405180830381600087803b15801561086b57600080fd5b505af115801561087f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108a39190613f6f565b5060048054604051630bb7c8fd60e31b81526001600160a01b03909116918291635dbe47e8916108d591339101613e72565b60206040518083038186803b1580156108ed57600080fd5b505afa158015610901573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109259190613f6f565b156109b457806001600160a01b0316639f65080b3361094333612c7b565b6040518363ffffffff1660e01b8152600401610960929190613f41565b602060405180830381600087803b15801561097a57600080fd5b505af115801561098e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b29190613f6f565b505b60025460408051808201825260208082527f746f6b656e65636f6e6f6d6963732e776974686472617741646472657373657390820152905163bf40fac160e01b81526000926001600160a01b03169163bf40fac191610a169190600401613fb6565b60206040518083038186803b158015610a2e57600080fd5b505afa158015610a42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a669190613ff4565b90506000816001600160a01b03166355ea6c47336040518263ffffffff1660e01b8152600401610a969190613e72565b60206040518083038186803b158015610aae57600080fd5b505afa158015610ac2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ae69190613ff4565b90506001600160a01b0381163314610ba857806001600160a01b0316896001600160a01b03161480610b2057506001600160a01b03891633145b610ba45760405162461bcd60e51b815260206004820152604960248201527f5b5145432d3030353031365d2d596f7572207769746864726177616c7320617260448201527f65206c696d6974656420746f20612073706563696669632077697468647261776064820152681030b2323932b9b99760b91b608482015260a401610630565b8098505b6000896001600160a01b03168b60405160006040518083038185875af1925050503d8060008114610bf5576040519150601f19603f3d011682016040523d82523d6000602084013e610bfa565b606091505b5050905080610c6a5760405162461bcd60e51b815260206004820152603660248201527f5b5145432d3030353031305d2d5472616e73666572206f66207468652077697460448201527534323930bbb0b61030b6b7bab73a103330b4b632b21760511b6064820152608401610630565b610c72613214565b6001985050505050505050505b92915050565b60048054604080516351cfd60960e11b815290516060936001600160a01b0390931692600092849263a39fac12928281019286929190829003018186803b158015610ccf57600080fd5b505afa158015610ce3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610d0b919081019061407f565b9050600081516001600160401b03811115610d2857610d28614011565b604051908082528060200260200182016040528015610d6d57816020015b6040805180820190915260008082526020820152815260200190600190039081610d465790505b50905060005b8251816001600160401b03161015610e9957604051806040016040528084836001600160401b031681518110610dab57610dab614130565b60200260200101516001600160a01b03168152602001856001600160a01b031663d903450386856001600160401b031681518110610deb57610deb614130565b60200260200101516040518263ffffffff1660e01b8152600401610e0f9190613e72565b60206040518083038186803b158015610e2757600080fd5b505afa158015610e3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5f9190614146565b81525082826001600160401b031681518110610e7d57610e7d614130565b602002602001018190525080610e929061415f565b9050610d73565b509392505050565b60048054604051630bb7c8fd60e31b81526000926001600160a01b03909216918291635dbe47e891610ed591339101613e72565b60206040518083038186803b158015610eed57600080fd5b505afa158015610f01573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f259190613f6f565b15610f935760405162461bcd60e51b815260206004820152603860248201527f5b5145432d3030353030335d2d5468652076616c696461746f7220697320616c6044820152773932b0b23c9037b7103a34329039b437b93a103634b9ba1760411b6064820152608401610630565b6000610f9d611391565b90506000610faa33612c7b565b9050826001600160a01b031663972c53566040518163ffffffff1660e01b815260040160206040518083038186803b158015610fe557600080fd5b505afa158015610ff9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061101d9190614146565b82141580611031575061102e61321d565b81115b61109c5760405162461bcd60e51b815260206004820152603660248201527f5b5145432d3030353031375d2d4e6f7420656e6f756768207374616b6520746f6044820152751032b73a32b9103a34329039b437b93a103634b9ba1760511b6064820152608401610630565b60405163205bd49b60e11b81526001600160a01b038416906340b7a936906110ca9033908590600401613f41565b602060405180830381600087803b1580156110e457600080fd5b505af11580156110f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061111c9190613f6f565b506000836001600160a01b031663972c53566040518163ffffffff1660e01b815260040160206040518083038186803b15801561115857600080fd5b505afa15801561116c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111909190614146565b90505b808310156112d0576000846001600160a01b031663d092e1866040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156111d857600080fd5b505af11580156111ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112109190613ff4565b9050336001600160a01b03821614611257576040516001600160a01b038216907f1cda35c48b19942f4f45e27586e5ef5b9ce254857fceefe5a6fb113797e814da90600090a25b846001600160a01b031663972c53566040518163ffffffff1660e01b815260040160206040518083038186803b15801561129057600080fd5b505afa1580156112a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112c89190614146565b915050611193565b604051630bb7c8fd60e31b81526001600160a01b03851690635dbe47e8906112fc903390600401613e72565b60206040518083038186803b15801561131457600080fd5b505afa158015611328573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061134c9190613f6f565b61135b57600094505050505090565b60405133907f5e4b86c0debc2f67504595cba51b583628b7b27eb5a038a6e08dc32f51fdc51790600090a2600194505050505090565b60008061139c61322a565b60405162498bff60e81b815260206004820152601b60248201527a636f6e737469747574696f6e2e6d61784e56616c696461746f727360281b60448201529091506000906001600160a01b0383169063498bff009060640160206040518083038186803b15801561140c57600080fd5b505afa158015611420573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114449190614146565b60405162498bff60e81b815260206004820152602260248201527f636f6e737469747574696f6e2e6d61784e5374616e64627956616c696461746f604482015261727360f01b60648201529091506000906001600160a01b0384169063498bff009060840160206040518083038186803b1580156114c157600080fd5b505afa1580156114d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114f99190614146565b905080806115078185614186565b6115119190614186565b94505050505090565b60035460405163d903450360e01b8152600091829182916001600160a01b03169063d90345039061154f908790600401613e72565b60606040518083038186803b15801561156757600080fd5b505afa15801561157b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061159f919061419e565b6001600160a01b038716600090815260076020526040902060010154919450925090506115cc8284614186565b6115d69190613f2a565b949350505050565b60006115e86132c4565b6001600160a01b0316336001600160a01b0316146116185760405162461bcd60e51b8152600401610630906141cc565b81600061162485612f3e565b905083811015611632578091505b600061163c61330e565b90506000816001600160a01b031663061c904030896040518363ffffffff1660e01b815260040161166e929190614249565b60806040518083038186803b15801561168657600080fd5b505afa15801561169a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116be9190614263565b90506000816040015182600001516116d69190614186565b905060006116e48286613f2a565b905085811015611758576001600160a01b038416635c7700df8a611708848a613f2a565b6040518363ffffffff1660e01b8152600401611725929190613f41565b600060405180830381600087803b15801561173f57600080fd5b505af1158015611753573d6000803e3d6000fd5b505050505b600354604051630bb7c8fd60e31b81526001600160a01b0390911690635dbe47e890611788908c90600401613e72565b60206040518083038186803b1580156117a057600080fd5b505afa1580156117b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117d89190613f6f565b15611863576003546040516304dffdc160e31b81526001600160a01b03909116906326ffee089061180f908c908a90600401613f41565b602060405180830381600087803b15801561182957600080fd5b505af115801561183d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118619190613f6f565b505b61186c89612f3e565b6001600160a01b038a166000908152600760205260409020600101549095508510156118b1576001600160a01b03891660009081526007602052604090206001018590555b60048054604051630bb7c8fd60e31b81526001600160a01b0390911691635dbe47e8916118e0918d9101613e72565b60206040518083038186803b1580156118f857600080fd5b505afa15801561190c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119309190613f6f565b156119c1576004546001600160a01b0316639f65080b8a61195081612c7b565b6040518363ffffffff1660e01b815260040161196d929190613f41565b602060405180830381600087803b15801561198757600080fd5b505af115801561199b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119bf9190613f6f565b505b600254604080516060810190915260248082526000926001600160a01b031691633fb902719161458d60208301396040518263ffffffff1660e01b8152600401611a0b9190613fb6565b60206040518083038186803b158015611a2357600080fd5b505afa158015611a37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a5b9190613ff4565b60405163690e7c0960e01b8152600481018d90529091506001600160a01b0382169063690e7c099089906024016020604051808303818588803b158015611aa157600080fd5b505af1158015611ab5573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250810190611ada9190613f6f565b9b9a5050505050505050505050565b6000611af36132c4565b6001600160a01b038316600090815260076020908152604080832060020180548251818502810185019093528083529495509293909291830182828015611b5957602002820191906000526020600020905b815481526020019060010190808311611b45575b5093945060009350611b6a92505050565b604051908082528060200260200182016040528015611b93578160200160208202803683370190505b506001600160a01b03841660009081526007602090815260409091208251611bc5936002909201929190910190613b4e565b5060005b8151811015611cb5576000828281518110611be657611be6614130565b60200260200101519050836001600160a01b03166311cb3b9b826040518263ffffffff1660e01b8152600401611c1e91815260200190565b60206040518083038186803b158015611c3657600080fd5b505afa158015611c4a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c6e9190613f6f565b15611ca2576001600160a01b0385166000908152600760209081526040822060020180546001810182559083529120018190555b5080611cad816142a9565b915050611bc9565b50505050565b60606005805480602002602001604051908101604052809291908181526020018280548015611d1357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611cf5575b5050505050905090565b808210611d925760405162461bcd60e51b815260206004820152603f60248201527f5b5145432d3033323030315d2d54696d65206c6f636b2072656c65617365206560448201527f6e64206d7573742062652061667465722072656c656173652073746172742e006064820152608401610630565b678ac7230489e80000341015611e065760405162461bcd60e51b815260206004820152603360248201527f5b5145432d3033323030325d2d476976656e20616d6f756e742069732062656c60448201527237bb903232b837b9b4ba1036b4b734b6bab69760691b6064820152608401610630565b640eea4f033d8110611e685760405162461bcd60e51b815260206004820152602560248201527f5b5145432d3033323030355d2d52656c6561736520656e6420697320746f6f206044820152643434b3b41760d91b6064820152608401610630565b611e7183613370565b6001600160a01b038316600090815260016020526040902054600a11611efc5760405162461bcd60e51b815260206004820152603a60248201527f5b5145432d3033323030335d2d4661696c656420746f206465706f736974206160448201527936b7bab73a16103a37b79036b0b73c903a34b6b2b637b1b5b99760311b6064820152608401610630565b42811115611fb757611f2860405180606001604052806000815260200160008152602001600081525090565b348152602080820184815260408084018581526001600160a01b03881660008181526001808752848220805480830182559083529690912087516003909702019586559351938501939093555160029093019290925590517f42dc784443c8161a28f39bb05a48bc8bac980ab9115d5b9cfef307f00204bd7f90611fad9084906142c4565b60405180910390a2505b611fc1833461352e565b505050565b611fce6132c4565b6001600160a01b0316336001600160a01b031614611ffe5760405162461bcd60e51b8152600401610630906141cc565b6001600160a01b0390911660009081526007602090815260408220600201805460018101825590835291200155565b6120373334613534565b565b60006120458383612fcc565b9392505050565b60048054604051630bb7c8fd60e31b81526000926001600160a01b0390921691635dbe47e89161207e91869101613e72565b60206040518083038186803b15801561209657600080fd5b505afa1580156120aa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c7f9190613f6f565b6120d781613370565b50565b6120e2613b99565b6120ea613832565b905090565b6000806120fa61322a565b60405162498bff60e81b815260206004820152601e60248201527f636f6e737469747574696f6e2e636c6971756545706f63684c656e677468000060448201526001600160a01b03919091169063498bff009060640160206040518083038186803b15801561216857600080fd5b505afa15801561217c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121a09190614146565b9050600654816121b09190614186565b6121bb436001614186565b10156122535760405162461bcd60e51b815260206004820152605b60248201527f5b5145432d3030353031315d2d4e6f7420656e6f7567682074696d652068617360448201527f20656c61707365642073696e636520746865206c6173742065706f63682c206660648201527a30b4b632b2103a379036b0b5b2903a34329039b730b839b437ba1760291b608482015260a401610630565b60048054604080516351cfd60960e11b815290516001600160a01b039092169263a39fac12928282019260009290829003018186803b15801561229557600080fd5b505afa1580156122a9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526122d1919081019061407f565b80516122e591600591602090910190613bc1565b5080600660008282546122f89190614186565b9091555060019392505050565b60035460405163d903450360e01b815260009182916001600160a01b039091169063d90345039061233a908690600401613e72565b60606040518083038186803b15801561235257600080fd5b505afa158015612366573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061238a919061419e565b95945050505050565b600061239d61322a565b60405162498bff60e81b81526020600482015260196024820152780636f6e737469747574696f6e2e76616c57697468647261775603c1b60448201526001600160a01b03919091169063498bff009060640160206040518083038186803b15801561240757600080fd5b505afa15801561241b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061243f9190614146565b6124499042614186565b60035460405163d903450360e01b81529192506001600160a01b031690600090829063d90345039061247f903390600401613e72565b60606040518083038186803b15801561249757600080fd5b505afa1580156124ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124cf919061419e565b509150508381101561254b576040805162461bcd60e51b81526020600482015260248101919091527f5b5145432d3030353030355d2d416e6e6f756e6365642077697468647261776160448201527f6c206d757374206e6f74206578636565642063757272656e74207374616b652e6064820152608401610630565b6125558484613843565b60408051808201825284815260208082018781523360008181526007909352918490209251835551600190920191909155600480549251630bb7c8fd60e31b81526001600160a01b0390931692635dbe47e8926125b492909101613e72565b60206040518083038186803b1580156125cc57600080fd5b505afa1580156125e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126049190613f6f565b15611cb5576004546001600160a01b0316639f65080b3361262481612c7b565b6040518363ffffffff1660e01b8152600401612641929190613f41565b602060405180830381600087803b15801561265b57600080fd5b505af115801561266f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126939190613f6f565b5050505050565b600054610100900460ff16806126b3575060005460ff16155b6127165760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610630565b600054610100900460ff16158015612738576000805461ffff19166101011790555b600280546001600160a01b038087166001600160a01b0319928316179092556004805486841690831617905560038054928516929091169190911790558015611cb5576000805461ff001916905550505050565b60035460405163d903450360e01b815260009182916001600160a01b039091169063d9034503906127c1908690600401613e72565b60606040518083038186803b1580156127d957600080fd5b505afa1580156127ed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612811919061419e565b506001600160a01b038516600090815260076020526040902060010154909250612045915082613f2a565b6001600160a01b03821660009081526001602052604081205461286157506000610c7f565b6000805b6001600160a01b038516600090815260016020526040902054811015610e99576001600160a01b03851660009081526001602052604081208054839081106128af576128af614130565b9060005260206000209060030201604051806060016040529081600082015481526020016001820154815260200160028201548152505090506128f281866138bf565b6128fc9084614186565b925050808061290a906142a9565b915050612865565b6001600160a01b0381166000908152600160209081526040808320805482518185028101850190935280835260609492939192909184015b8282101561299a578382906000526020600020906003020160405180606001604052908160008201548152602001600182015481526020016002820154815250508152602001906001019061294a565b505050509050919050565b600354600254604080516060810190915260248082526001600160a01b0393841693636467b404938693911691633fb90271919061454660208301396040518263ffffffff1660e01b81526004016129fd9190613fb6565b60206040518083038186803b158015612a1557600080fd5b505afa158015612a29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a4d9190613ff4565b6001600160a01b0316637c47c16d856040518263ffffffff1660e01b8152600401612a789190613e72565b60206040518083038186803b158015612a9057600080fd5b505afa158015612aa4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ac89190614146565b6040518363ffffffff1660e01b8152600401612ae5929190613f41565b602060405180830381600087803b158015612aff57600080fd5b505af1158015612b13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b379190613f6f565b5060048054604051630bb7c8fd60e31b81526001600160a01b03909116918291635dbe47e891612b6991869101613e72565b60206040518083038186803b158015612b8157600080fd5b505afa158015612b95573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bb99190613f6f565b15612c4657806001600160a01b0316639f65080b83612bd785612c7b565b6040518363ffffffff1660e01b8152600401612bf4929190613f41565b602060405180830381600087803b158015612c0e57600080fd5b505af1158015612c22573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fc19190613f6f565b5050565b600354604051630bb7c8fd60e31b81526000916001600160a01b031690635dbe47e89061207e908590600401613e72565b60035460405163d903450360e01b8152600091829182916001600160a01b03169063d903450390612cb0908790600401613e72565b60606040518083038186803b158015612cc857600080fd5b505afa158015612cdc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d00919061419e565b6001600160a01b0387166000908152600760205260409020600101549194509250612d2c915083613f2a565b600254604080516060810190915260238082529294506000926001600160a01b0390921691633fb90271919061456a60208301396040518263ffffffff1660e01b8152600401612d7c9190613fb6565b60206040518083038186803b158015612d9457600080fd5b505afa158015612da8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dcc9190613ff4565b60405162498bff60e81b8152602060048201526024808201527f676f7665726e65642e45505146492e7374616b6544656c65676174696f6e466160448201526331ba37b960e11b60648201526001600160a01b03919091169063498bff009060840160206040518083038186803b158015612e4657600080fd5b505afa158015612e5a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e7e9190614146565b905060006b033b2e3c9fd0803ce8000000612e9983866142e5565b612ea39190614304565b90506000612eb18486614186565b905081811015612ec5579695505050505050565b5095945050505050565b6001600160a01b038116600090815260076020908152604091829020600201805483518184028101840190945280845260609392830182828015612f3257602002820191906000526020600020905b815481526020019060010190808311612f1e575b50505050509050919050565b60035460405163d903450360e01b815260009182916001600160a01b039091169063d903450390612f73908690600401613e72565b60606040518083038186803b158015612f8b57600080fd5b505afa158015612f9f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fc3919061419e565b50949350505050565b600080612fd76132c4565b9050612fe284611ae9565b6001600160a01b03841660009081526007602090815260408083206002018054825181850281018501909352808352919290919083018282801561304557602002820191906000526020600020905b815481526020019060010190808311613031575b505050505090506000805b8251811015612ec557600083828151811061306d5761306d614130565b6020026020010151905086856001600160a01b03166317fdb84e836040518263ffffffff1660e01b81526004016130a691815260200190565b60206040518083038186803b1580156130be57600080fd5b505afa1580156130d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130f69190614146565b10156131905760405163013cf08b60e01b8152600481018290526000906001600160a01b0387169063013cf08b9060240160006040518083038186803b15801561313f57600080fd5b505afa158015613153573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261317b9190810190614410565b9350505050808461318c9190614186565b9350505b508061319b816142a9565b915050613050565b60006131ad61330e565b604051637eee288d60e01b81529091506001600160a01b03821690637eee288d906131de9033908690600401613f41565b600060405180830381600087803b1580156131f857600080fd5b505af115801561320c573d6000803e3d6000fd5b50505050612c465b6120373361392f565b60006120ea61052561396c565b600254604080516060810190915260228082526000926001600160a01b031691633fb90271916145b160208301396040518263ffffffff1660e01b81526004016132749190613fb6565b60206040518083038186803b15801561328c57600080fd5b505afa1580156132a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120ea9190613ff4565b600254604080516060810190915260248082526000926001600160a01b031691633fb90271916145d360208301396040518263ffffffff1660e01b81526004016132749190613fb6565b600254604080518082018252601c81527f676f7665726e616e63652e766f74696e6757656967687450726f78790000000060208201529051633fb9027160e01b81526000926001600160a01b031691633fb90271916132749190600401613fb6565b6001600160a01b038116600090815260016020908152604080832080548251818502810185019093528083529192909190849084015b828210156133f657838290600052602060002090600302016040518060600160405290816000820154815260200160018201548152602001600282015481525050815260200190600101906133a6565b5050506001600160a01b038416600090815260016020526040812092935061341f929150613c16565b60005b8151811015611fc15781818151811061343d5761343d614130565b6020026020010151604001514210156134c1576001600160a01b0383166000908152600160205260409020825183908390811061347c5761347c614130565b6020908102919091018101518254600181810185556000948552938390208251600390920201908155918101519282019290925560409091015160029091015561351c565b826001600160a01b03167f646189d6820ad36442dfb379131756840be21178303b28f1f607ca4e45762e938383815181106134fe576134fe614130565b602002602001015160405161351391906142c4565b60405180910390a25b80613526816142a9565b915050613422565b612c4682825b6000811161359c5760405162461bcd60e51b815260206004820152602f60248201527f5b5145432d3030353030305d2d4164646974696f6e616c207374616b65206d7560448201526e39ba103737ba103132903d32b9379760891b6064820152608401610630565b60035460405163f5d82b6b60e01b81526001600160a01b039091169063f5d82b6b906135ce9085908590600401613f41565b602060405180830381600087803b1580156135e857600080fd5b505af11580156135fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136209190613f6f565b5060048054604051630bb7c8fd60e31b81526001600160a01b03909116918291635dbe47e89161365291879101613e72565b60206040518083038186803b15801561366a57600080fd5b505afa15801561367e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136a29190613f6f565b1561373157806001600160a01b0316639f65080b846136c086612c7b565b6040518363ffffffff1660e01b81526004016136dd929190613f41565b602060405180830381600087803b1580156136f757600080fd5b505af115801561370b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061372f9190613f6f565b505b600254604080516060810190915260248082526001600160a01b0390921691633fb90271919061454660208301396040518263ffffffff1660e01b815260040161377b9190613fb6565b60206040518083038186803b15801561379357600080fd5b505afa1580156137a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137cb9190613ff4565b6001600160a01b0316638cfe7df3846040518263ffffffff1660e01b81526004016137f69190613e72565b600060405180830381600087803b15801561381057600080fd5b505af1158015613824573d6000803e3d6000fd5b50505050611fc1828461399f565b61383a613b99565b6120ea33613abc565b600061384d61330e565b604051633a5f594760e21b815233600482015260248101859052604481018490529091506001600160a01b0382169063e97d651c90606401600060405180830381600087803b15801561389f57600080fd5b505af11580156138b3573d6000803e3d6000fd5b50505050611fc1613214565b600082602001518210156138d557508151610c7f565b826040015182106138e857506000610c7f565b60008284604001516138fa9190613f2a565b90506000846020015185604001516139129190613f2a565b90508082866000015161392591906142e5565b61238a9190614304565b600061393a82613abc565b90506000816040015182600001516139529190614186565b90508061395e84612f3e565b1015611fc157611fc161452f565b6004805460405163cf527e8b60e01b81526000926001600160a01b039092169163cf527e8b916132749160019101613e72565b60006139a961330e565b905060006139b683613abc565b90506000816040015182600001516139ce9190614186565b905084816139db86612f3e565b6139e59190613f2a565b1015613a535760405162461bcd60e51b815260206004820152603760248201527f5b5145432d3030353031345d2d546f74616c206c6f636b656420616d6f756e746044820152761036bab9ba103737ba1032bc31b2b2b21039ba30b5b29760491b6064820152608401610630565b60405163282d3fdf60e01b81526001600160a01b0384169063282d3fdf90613a819087908990600401613f41565b600060405180830381600087803b158015613a9b57600080fd5b505af1158015613aaf573d6000803e3d6000fd5b505050506126938461392f565b613ac4613b99565b6000613ace61330e565b6040516218724160e61b81529091506001600160a01b0382169063061c904090613afe9030908790600401614249565b60806040518083038186803b158015613b1657600080fd5b505afa158015613b2a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120459190614263565b828054828255906000526020600020908101928215613b89579160200282015b82811115613b89578251825591602001919060010190613b6e565b50613b95929150613c37565b5090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b828054828255906000526020600020908101928215613b89579160200282015b82811115613b8957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613be1565b50805460008255600302906000526020600020908101906120d79190613c4c565b5b80821115613b955760008155600101613c38565b5b80821115613b95576000808255600182018190556002820155600301613c4d565b6001600160a01b03811681146120d757600080fd5b60008060408385031215613c9657600080fd5b823591506020830135613ca881613c6e565b809150509250929050565b602080825282518282018190526000919060409081850190868401855b82811015613cfe57815180516001600160a01b03168552860151868501529284019290850190600101613cd0565b5091979650505050505050565b600060208284031215613d1d57600080fd5b813561204581613c6e565b600080600060608486031215613d3d57600080fd5b833592506020840135613d4f81613c6e565b929592945050506040919091013590565b6020808252825182820181905260009190848201906040850190845b81811015613da15783516001600160a01b031683529284019291840191600101613d7c565b50909695505050505050565b600080600060608486031215613dc257600080fd5b8335613dcd81613c6e565b95602085013595506040909401359392505050565b60008060408385031215613df557600080fd5b8235613e0081613c6e565b946020939093013593505050565b600060208284031215613e2057600080fd5b5035919050565b600080600060608486031215613e3c57600080fd5b8335613e4781613c6e565b92506020840135613e5781613c6e565b91506040840135613e6781613c6e565b809150509250925092565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b81811015613da157613ec98385518051825260208082015190830152604090810151910152565b9284019260609290920191600101613ea2565b6020808252825182820181905260009190848201906040850190845b81811015613da157835183529284019291840191600101613ef8565b634e487b7160e01b600052601160045260246000fd5b600082821015613f3c57613f3c613f14565b500390565b6001600160a01b03929092168252602082015260400190565b80518015158114613f6a57600080fd5b919050565b600060208284031215613f8157600080fd5b61204582613f5a565b60005b83811015613fa5578181015183820152602001613f8d565b83811115611cb55750506000910152565b6020815260008251806020840152613fd5816040850160208701613f8a565b601f01601f19169190910160400192915050565b8051613f6a81613c6e565b60006020828403121561400657600080fd5b815161204581613c6e565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b038111828210171561404957614049614011565b60405290565b604051601f8201601f191681016001600160401b038111828210171561407757614077614011565b604052919050565b6000602080838503121561409257600080fd5b82516001600160401b03808211156140a957600080fd5b818501915085601f8301126140bd57600080fd5b8151818111156140cf576140cf614011565b8060051b91506140e084830161404f565b81815291830184019184810190888411156140fa57600080fd5b938501935b83851015614124578451925061411483613c6e565b82825293850193908501906140ff565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561415857600080fd5b5051919050565b60006001600160401b038083168181141561417c5761417c613f14565b6001019392505050565b6000821982111561419957614199613f14565b500190565b6000806000606084860312156141b357600080fd5b8351925060208401519150604084015190509250925092565b60208082526057908201527f5b5145432d3030353031335d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c79207468652056616c696461746f7273536c617368696e67566f746960608201527637339031b7b73a3930b1ba103430b99030b1b1b2b9b99760491b608082015260a00190565b6001600160a01b0392831681529116602082015260400190565b60006080828403121561427557600080fd5b61427d614027565b825181526020830151602082015260408301516040820152606083015160608201528091505092915050565b60006000198214156142bd576142bd613f14565b5060010190565b81518152602080830151908201526040808301519082015260608101610c7f565b60008160001904831182151516156142ff576142ff613f14565b500290565b60008261432157634e487b7160e01b600052601260045260246000fd5b500490565b600061010080838503121561433a57600080fd5b604051908101906001600160401b038211818310171561435c5761435c614011565b81604052809250835181526020840151602082015260408401516040820152606084015160608201526080840151608082015260a084015160a082015260c084015160c082015260e084015160e0820152505092915050565b6000606082840312156143c757600080fd5b604051606081018181106001600160401b03821117156143e9576143e9614011565b80604052508091508251815260208301516020820152604083015160408201525092915050565b6000806000806080858703121561442657600080fd5b84516001600160401b038082111561443d57600080fd5b908601906101a0828903121561445257600080fd5b61445a614027565b82518281111561446957600080fd5b8301601f81018a1361447a57600080fd5b805160208482111561448e5761448e614011565b6144a0601f8301601f1916820161404f565b94508185528b818385010111156144b657600080fd5b6144c582828701838601613f8a565b8484526144d48c828801614326565b818501526144e68c61012088016143b5565b60408501526144f86101808701613f5a565b606085015283995061450b818c01613fe9565b985050505050505061451f60408601613fe9565b6060959095015193969295505050565b634e487b7160e01b600052600160045260246000fdfe746f6b656e65636f6e6f6d6963732e76616c69646174696f6e526577617264506f6f6c73676f7665726e616e63652e657870657274732e45505146492e706172616d6574657273676f7665726e616e63652e76616c696461746f72732e736c617368696e67457363726f77676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273676f7665726e616e63652e76616c696461746f72732e736c617368696e67566f74696e67a264697066735822122004fd4d925e6ceee292805171887eefe5c869698b42d4c1bcf97b11e732056bb364736f6c63430008090033"
		storageStakesSortedABI    = "[{\"inputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"constructor\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"}],\"name\": \"AddressAdded\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"}],\"name\": \"AddressRemoved\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"previousOwner\",\"type\": \"address\"},{\"indexed\": true,\"internalType\": \"address\",\"name\": \"newOwner\",\"type\": \"address\"}],\"name\": \"OwnershipTransferred\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"_stake\",\"type\": \"uint256\"}],\"name\": \"StakeUpdated\",\"type\": \"event\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"},{\"internalType\": \"uint256\",\"name\": \"_stake\",\"type\": \"uint256\"}],\"name\": \"addAddr\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"name\": \"addrStake\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"}],\"name\": \"contains\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"getAddresses\",\"outputs\": [{\"internalType\": \"address[]\",\"name\": \"\",\"type\": \"address[]\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"name\": \"linkedList\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"listSize\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"owner\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"removeLast\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"renounceOwnership\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"newOwner\",\"type\": \"address\"}],\"name\": \"transferOwnership\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"},{\"internalType\": \"uint256\",\"name\": \"_stake\",\"type\": \"uint256\"}],\"name\": \"updateStake\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"}]"
		storageStakesSortedCode   = "0x608060405234801561001057600080fd5b5061001a33610058565b600160005260026020527fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e080546001600160a01b03191690556100a8565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610eb3806100b76000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80639f65080b116100715780639f65080b14610147578063a39fac121461015a578063cf527e8b1461016f578063d092e18614610198578063d9034503146101a0578063f2fde38b146101c057600080fd5b806340b7a936146100ae5780635dbe47e8146100d6578063715018a6146101015780638da5cb5b1461010b578063972c535614610130575b600080fd5b6100c16100bc366004610d0e565b6101d3565b60405190151581526020015b60405180910390f35b6100c16100e4366004610d38565b6001600160a01b0316600090815260016020526040902054151590565b6101096104e3565b005b6000546001600160a01b03165b6040516001600160a01b0390911681526020016100cd565b61013960035481565b6040519081526020016100cd565b6100c1610155366004610d0e565b610519565b6101626107b5565b6040516100cd9190610d5a565b61011861017d366004610d38565b6002602052600090815260409020546001600160a01b031681565b61011861087d565b6101396101ae366004610d38565b60016020526000908152604090205481565b6101096101ce366004610d38565b6109dc565b600080546001600160a01b031633146102075760405162461bcd60e51b81526004016101fe90610da7565b60405180910390fd5b826001600160a01b0381161580159061022a57506001600160a01b038116600114155b6102935760405162461bcd60e51b815260206004820152603460248201527f5b5145432d3033373030355d2d54686520616464726573732073686f756c64206044820152733737ba103132902422a0a21037b9102a20a4a61760611b60648201526084016101fe565b6001600160a01b038416600090815260016020526040902054156103615760405162461bcd60e51b815260206004820152607360248201527f5b5145432d3033373030305d2d54686520616464726573732068617320616c7260448201527f65616479206265656e20616464656420746f207468652073746f726167652c2060648201527f6661696c656420746f2061646420746f2074686520616464726573732073746160848201527235b2b99039b7b93a32b21039ba37b930b3b29760691b60a482015260c4016101fe565b600083116103e55760405162461bcd60e51b815260206004820152604560248201527f5b5145432d3033373030315d2d496e76616c6964207374616b652c206661696c60448201527f656420746f2061646420746865206164647265737320746f207468652073746f6064820152643930b3b29760d91b608482015260a4016101fe565b60006103f084610a77565b6001600160a01b0380821660009081526002602052604090205491925016158061044a576001600160a01b03808316600090815260026020526040808220548984168352912080546001600160a01b031916919092161790555b6001600160a01b03828116600090815260026020908152604080832080546001600160a01b031916948b16948517905592825260019052908120869055600380549161049583610df2565b90915550506040516001600160a01b038716907fa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f90600090a26104d786610af0565b50600195945050505050565b6000546001600160a01b0316331461050d5760405162461bcd60e51b81526004016101fe90610da7565b6105176000610bb4565b565b600080546001600160a01b031633146105445760405162461bcd60e51b81526004016101fe90610da7565b6001600160a01b0383166000908152600160205260409020546105e35760405162461bcd60e51b815260206004820152604b60248201527f5b5145432d3033373030325d2d5468652061646472657373206973206e6f742060448201527f696e207468652073746f726167652c206661696c656420746f2075706461746560648201526a103a34329039ba30b5b29760a91b608482015260a4016101fe565b6001600160a01b038084166000908152600260208190526040822054909216919061060d86610c04565b6001600160a01b0390811682526020808301939093526040918201600090812080549583166001600160a01b031996871617905590871681526002909252902080549091169055816106b7576003805490600061066983610e0d565b90915550506001600160a01b038316600081815260016020526040808220829055517f24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb9190a25060016107af565b6001600160a01b03831660008181526001602052604090819020849055517fab0e25dc39626189cfb41155020ba89e726b10244275733e9d7c63cf33ffccdb906107049085815260200190565b60405180910390a2600061071783610a77565b6001600160a01b03808216600090815260026020526040902054919250161580610771576001600160a01b03808316600090815260026020526040808220548884168352912080546001600160a01b031916919092161790555b6001600160a01b03828116600090815260026020526040902080546001600160a01b0319169187169190911790556107a885610af0565b6001925050505b92915050565b6060600060019050600060035467ffffffffffffffff8111156107da576107da610e24565b604051908082528060200260200182016040528015610803578160200160208202803683370190505b506003549091505b8015610876576001600160a01b0392831660009081526002602052604090205490921691828261083c600184610e3a565b8151811061084c5761084c610e51565b6001600160a01b03909216602092830291909101909101528061086e81610e0d565b91505061080b565b5092915050565b600080546001600160a01b031633146108a85760405162461bcd60e51b81526004016101fe90610da7565b6000600354116109205760405162461bcd60e51b815260206004820152603d60248201527f5b5145432d3033373030335d2d546865206c69737420697320656d7074792c2060448201527f6661696c656420746f2072656d6f76652074686520616464726573732e00000060648201526084016101fe565b600260209081527fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e080546001600160a01b038082166000818152604080822080549094166001600160a01b03199586161790955582549093169091556001909352908120819055600380549161099583610e0d565b90915550506040516001600160a01b038216907f24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb90600090a26109d781610af0565b905090565b6000546001600160a01b03163314610a065760405162461bcd60e51b81526004016101fe90610da7565b6001600160a01b038116610a6b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101fe565b610a7481610bb4565b50565b600060015b6001600160a01b03808216600090815260026020908152604080832054909316825260019052205483118015610acb57506001600160a01b038181166000908152600260205260409020541615155b156107af576001600160a01b0390811660009081526002602052604090205416610a7c565b6001600160a01b03811660009081526001602052604090205415610b38576001600160a01b038116600090815260016020526040902054610b3357610b33610e67565b610b5e565b6001600160a01b03811660009081526001602052604090205415610b5e57610b5e610e67565b6001600160a01b03808216600090815260026020526040902054168015610bb0576001600160a01b038083166000908152600160205260408082205492841682529020541015610bb057610bb0610e67565b5050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038116600090815260016020526040812054610cab5760405162461bcd60e51b815260206004820152605360248201527f5b5145432d3033373030345d2d5468652061646472657373206973206e6f742060448201527f696e207468652073746f726167652c206661696c656420746f206765742074686064820152723290383932bb34b7bab99030b2323932b9b99760691b608482015260a4016101fe565b60015b6001600160a01b038181166000908152600260205260409020548116908416146107af576001600160a01b0390811660009081526002602052604090205416610cae565b80356001600160a01b0381168114610d0957600080fd5b919050565b60008060408385031215610d2157600080fd5b610d2a83610cf2565b946020939093013593505050565b600060208284031215610d4a57600080fd5b610d5382610cf2565b9392505050565b6020808252825182820181905260009190848201906040850190845b81811015610d9b5783516001600160a01b031683529284019291840191600101610d76565b50909695505050505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6000600019821415610e0657610e06610ddc565b5060010190565b600081610e1c57610e1c610ddc565b506000190190565b634e487b7160e01b600052604160045260246000fd5b600082821015610e4c57610e4c610ddc565b500390565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052600160045260246000fdfea264697066735822122042c702f67d7f4b66c44e8412a55c379d1716b66d13a65395c49bf8ae0e1f9f3864736f6c63430008090033"
		storageStakesABI          = "[{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"}],\"name\": \"AddressAdded\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"}],\"name\": \"AddressRemoved\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"_stake\",\"type\": \"uint256\"}],\"name\": \"DelegatedStakeUpdated\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"previousOwner\",\"type\": \"address\"},{\"indexed\": true,\"internalType\": \"address\",\"name\": \"newOwner\",\"type\": \"address\"}],\"name\": \"OwnershipTransferred\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"_stake\",\"type\": \"uint256\"}],\"name\": \"StakeDecreased\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"_stake\",\"type\": \"uint256\"}],\"name\": \"StakeIncreased\",\"type\": \"event\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"},{\"internalType\": \"uint256\",\"name\": \"_stake\",\"type\": \"uint256\"}],\"name\": \"add\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"name\": \"addrList\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"name\": \"addrStake\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"listIndex\",\"type\": \"uint256\"},{\"internalType\": \"uint256\",\"name\": \"stake\",\"type\": \"uint256\"},{\"internalType\": \"uint256\",\"name\": \"delegatedStake\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"}],\"name\": \"contains\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"getAddresses\",\"outputs\": [{\"internalType\": \"address[]\",\"name\": \"\",\"type\": \"address[]\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"owner\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"renounceOwnership\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"},{\"internalType\": \"uint256\",\"name\": \"_delegatedStake\",\"type\": \"uint256\"}],\"name\": \"setDelegated\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"size\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"_addr\",\"type\": \"address\"},{\"internalType\": \"uint256\",\"name\": \"_stake\",\"type\": \"uint256\"}],\"name\": \"sub\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"newOwner\",\"type\": \"address\"}],\"name\": \"transferOwnership\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"}]"
		storageStakesCode         = "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610c4c8061007e6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638da5cb5b116100715780638da5cb5b14610131578063949d225d14610142578063a39fac1214610153578063d903450314610168578063f2fde38b146101b1578063f5d82b6b146101c457600080fd5b806326ffee08146100ae5780635dbe47e8146100d65780636467b404146100e9578063715018a6146100fc5780638040256414610106575b600080fd5b6100c16100bc366004610aa8565b6101d7565b60405190151581526020015b60405180910390f35b6100c16100e4366004610ad2565b61039a565b6100c16100f7366004610aa8565b6103df565b6101046104a5565b005b610119610114366004610af4565b6104db565b6040516001600160a01b0390911681526020016100cd565b6000546001600160a01b0316610119565b6002546040519081526020016100cd565b61015b610505565b6040516100cd9190610b0d565b610196610176366004610ad2565b600160208190526000918252604090912080549181015460029091015483565b604080519384526020840192909252908201526060016100cd565b6101046101bf366004610ad2565b610567565b6100c16101d2366004610aa8565b610602565b600080546001600160a01b0316331461020b5760405162461bcd60e51b815260040161020290610b5a565b60405180910390fd5b6001600160a01b038316600090815260016020526040902080546102a65760405162461bcd60e51b815260206004820152604660248201527f5b5145432d3033363030305d2d546865206164647265737320646f6573206e6f60448201527f742065786973742c206661696c656420746f206465637265617365207468652060648201526539ba30b5b29760d11b608482015260a401610202565b828160010154101561032e5760405162461bcd60e51b815260206004820152604560248201527f5b5145432d3033363030325d2d546865207374616b6520746f2072656d6f766560448201527f2069732067726561746572207468616e2074686520617661696c61626c6520736064820152643a30b5b29760d91b608482015260a401610202565b828160010160008282546103429190610ba5565b90915550506040518381526001600160a01b038516907f700865370ffb2a65a2b0242e6a64b21ac907ed5ecd46c9cffc729c177b2b1c699060200160405180910390a261038e846106a5565b60019150505b92915050565b6001600160a01b0381166000908152600160208190526040822001541515806103945750506001600160a01b0316600090815260016020526040902060020154151590565b600080546001600160a01b0316331461040a5760405162461bcd60e51b815260040161020290610b5a565b6001600160a01b03831660009081526001602052604090205415801561042e575081155b1561043b57506001610394565b6001600160a01b03831660008181526001602052604090819020600201849055517fc85c322a22d4a5cd3d9611287b2abf5bc9e6426d3ed5dc8d77017a7fe8fdc9109061048b9085815260200190565b60405180910390a261049c836106a5565b50600192915050565b6000546001600160a01b031633146104cf5760405162461bcd60e51b815260040161020290610b5a565b6104d960006106c4565b565b600281815481106104eb57600080fd5b6000918252602090912001546001600160a01b0316905081565b6060600280548060200260200160405190810160405280929190818152602001828054801561055d57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161053f575b5050505050905090565b6000546001600160a01b031633146105915760405162461bcd60e51b815260040161020290610b5a565b6001600160a01b0381166105f65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610202565b6105ff816106c4565b50565b600080546001600160a01b0316331461062d5760405162461bcd60e51b815260040161020290610b5a565b8161063a57506000610394565b6001600160a01b03831660009081526001602081905260408220018054849290610665908490610bbc565b90915550506040518281526001600160a01b038416907f8b0ed825817a2e696c9a931715af4609fc60e1701f09c89ee7645130e937eb2d9060200161048b565b6106ae8161039a565b6106bb576105ff81610714565b6105ff816108e0565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526001602081905260408220546107399190610ba5565b60025490915060009061074e90600190610ba5565b905080821461080d57610762826001610bbc565b600160006002848154811061077957610779610bd4565b60009182526020808320909101546001600160a01b0316835282019290925260400190205560028054829081106107b2576107b2610bd4565b600091825260209091200154600280546001600160a01b0390921691849081106107de576107de610bd4565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b600280548061081e5761081e610bea565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0385168252600190819052604082208281559081018290556002015561087083610993565b6040516001600160a01b038416907f24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb90600090a28082146108db576108db600283815481106108c1576108c1610bd4565b6000918252602090912001546001600160a01b0316610993565b505050565b6001600160a01b0381166000908152600160205260409020546105ff576002805460018101825560009182527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180546001600160a01b0319166001600160a01b03841690811790915560405190917fa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f91a26002546001600160a01b0382166000908152600160205260409020556105ff815b6001600160a01b0381166000908152600160205260409020546002548111156109be576109be610c00565b6109c78261039a565b15610a3057600081116109dc576109dc610c00565b6001600160a01b03821660009081526001602081905260409091200154151580610a2057506001600160a01b03821660009081526001602052604090206002015415155b610a2c57610a2c610c00565b5050565b8015610a3e57610a3e610c00565b6001600160a01b03821660009081526001602081905260409091200154158015610a2057506001600160a01b03821660009081526001602052604090206002015415610a2c57610a2c610c00565b80356001600160a01b0381168114610aa357600080fd5b919050565b60008060408385031215610abb57600080fd5b610ac483610a8c565b946020939093013593505050565b600060208284031215610ae457600080fd5b610aed82610a8c565b9392505050565b600060208284031215610b0657600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b81811015610b4e5783516001600160a01b031683529284019291840191600101610b29565b50909695505050505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052601160045260246000fd5b600082821015610bb757610bb7610b8f565b500390565b60008219821115610bcf57610bcf610b8f565b500190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052600160045260246000fdfea264697066735822122082f38d3c48e7ec143a4fca8ac90fd15886e51b9d2cebb0f3119a0575da1fabbe64736f6c63430008090033"
	)

	cliqueConfig := params.AllCliqueProtocolChanges
	cliqueConfig.Clique.Epoch = 101

	var (
		db1                   = rawdb.NewMemoryDatabase()
		db2                   = rawdb.NewMemoryDatabase()
		exclusionSetProvider1 = new(NoopExclusionSetProvider)
		exclusionSetProvider2 = new(NoopExclusionSetProvider)
		reg1                  = contracts.NewTestModeRegistry()
		reg2                  = contracts.NewTestModeRegistry()
		engine1               = New(cliqueConfig.Clique, db1, exclusionSetProvider1, reg1)
		engine2               = New(cliqueConfig.Clique, db2, exclusionSetProvider2, reg2)
		signersCount          = 3
		keys                  = make(map[common.Address]*ecdsa.PrivateKey)
		addresses             = make([]common.Address, 0, signersCount)
		commonBlocksCount     = 70
	)

	reg1.SetRewardReceiver(common.Address{})
	reg2.SetRewardReceiver(common.Address{})

	for len(addresses) != signersCount {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		keys[addr] = key
		addresses = append(addresses, addr)
	}
	sort.Sort(signersAscending(addresses))

	genSpec := &core.Genesis{
		Config:    cliqueConfig,
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

	auth, err := bind.NewKeyedTransactorWithChainID(keys[addresses[0]], genSpec.Config.ChainID)
	if err != nil {
		t.Error(err)
	}

	genesis1 := genSpec.MustCommit(db1)
	genesis2 := genSpec.MustCommit(db2)

	commonBlocks1 := generateChain(t, cliqueConfig, genesis1, engine1, db1, commonBlocksCount, addresses, keys, reg1, false)
	commonBlocks2 := generateChain(t, cliqueConfig, genesis2, engine2, db2, commonBlocksCount, addresses, keys, reg2, false)

	// Insert common blocks
	chain1, _ := core.NewBlockChain(db1, nil, cliqueConfig, engine1, vm.Config{}, nil, nil)
	defer chain1.Stop()
	if _, err := chain1.InsertChain(commonBlocks1); err != nil {
		t.Fatalf("failed to insert commonBlocks blocks: %v", err)
	}
	chain2, _ := core.NewBlockChain(db2, nil, cliqueConfig, engine2, vm.Config{}, nil, nil)
	defer chain2.Stop()
	if _, err := chain2.InsertChain(commonBlocks2); err != nil {
		t.Fatalf("failed to insert commonBlocks blocks: %v", err)
	}

	backend1 := backends.NewCustomSimulatedBackend(db1, chain1, *genSpec)

	storageStakesSorted, err := deployContract(auth, storageStakesSortedABI, storageStakesSortedCode, backend1)
	if err != nil {
		t.Error(err)
	}

	storageStakes, err := deployContract(auth, storageStakesABI, storageStakesCode, backend1)
	if err != nil {
		t.Error(err)
	}

	registryImpl, err := deployContract(auth, contractsRegistryImplABI, contractsRegistryImplCode, backend1)
	if err != nil {
		t.Error(err)
	}

	validatorsImpl, err := deployContract(auth, validatorsImplABI, validatorsImplCode, backend1)
	if err != nil {
		t.Error(err)
	}

	validatorsParsedABI, err := abi.JSON(strings.NewReader(validatorsImplABI))
	if err != nil {
		t.Fatalf("Cannot parse ABI: %v", err)
	}
	initValidatorsInput, err := validatorsParsedABI.Pack("initialize", registryImpl, storageStakesSorted, storageStakes)
	if err != nil {
		t.Fatalf("Failed to pack data for someFunction: %v", err)
	}

	nonce, err := backend1.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		t.Fatalf("Failed to retrieve account nonce: %v", err)
	}

	tx := types.NewTransaction(nonce, validatorsImpl, big.NewInt(0), 1000000, big.NewInt(1000000), initValidatorsInput)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(genSpec.Config.ChainID), keys[addresses[0]])
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}

	err = backend1.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatalf("Failed to send transaction: %v", err)
	}

	registryParsedABI, err := abi.JSON(strings.NewReader(contractsRegistryImplABI))
	initRegistryInput, err := registryParsedABI.Pack("initialize", []common.Address{addresses[0]},
		[]string{"governance.validators"}, []common.Address{validatorsImpl},
	)
	if err != nil {
		t.Fatalf("Failed to pack data for someFunction: %v", err)
	}

	nonce, err = backend1.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		t.Fatalf("Failed to retrieve account nonce: %v", err)
	}

	txBlockHeader := backend1.GetPendingBlockHeader()
	txBlockHeader.Extra = make([]byte, extraVanity+extraSeal)
	key := keys[addresses[txBlockHeader.Number.Uint64()%uint64(len(keys))]]

	tx = types.NewTransaction(nonce, registryImpl, big.NewInt(0), 1000000, big.NewInt(80000), initRegistryInput)
	signedTx, err = types.SignTx(tx, types.NewEIP155Signer(genSpec.Config.ChainID), keys[addresses[0]])
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}

	err = backend1.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatalf("Failed to send transaction: %v", err)
	}

	sig, _ := crypto.Sign(SealHash(txBlockHeader).Bytes(), key)
	copy(txBlockHeader.Extra[len(txBlockHeader.Extra)-extraSeal:], sig)
	backend1.SetPendingBlockHeader(txBlockHeader)

	backend1.Commit()

	// Set on-chain validators list

	//storageStakesParsedABI, err := abi.JSON(strings.NewReader(storageStakesABI))
	//if err != nil {
	//	t.Fatalf("Cannot parse ABI: %v", err)
	//}
	//enterShortlist(t, auth, addresses[1], storageStakes, storageStakesParsedABI, validatorsParsedABI, backend1, cliqueConfig.ChainID, keys[addresses[0]])
	//
	//backend1.Commit()

	reg1 = contracts.NewRegistry(registryImpl, common.Address{}, backend1)
	reg2 = contracts.NewRegistry(registryImpl, common.Address{}, backend1)
	engine1.registry = reg1
	engine2.registry = reg2

	blocksToGenerate := 40
	sideChain := generateChain(t, cliqueConfig, chain2.CurrentBlock(), engine2, db2, blocksToGenerate, addresses, keys, reg2, true)

	_, err = chain1.InsertChain(sideChain)
	if err == nil {
		t.Fatalf("shouldn't validate blocks")
	}
	if !errors.Is(err, errUnauthorizedSigner) && !errors.Is(err, errWrongDifficulty) {
		t.Fatalf("failed to insert sidechain: %v", err)
	}
}

func deployContract(opts *bind.TransactOpts, contractABI, code string, backend bind.ContractBackend) (common.Address, error) {
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return common.Address{}, err
	}

	opts.GasLimit, err = backend.EstimateGas(context.Background(), ethereum.CallMsg{
		From: opts.From,
		Data: common.Hex2Bytes(code),
	})
	if err != nil {
		return common.Address{}, nil
	}
	opts.GasLimit += 50000

	address, tx, _, err := bind.DeployContract(opts, parsedABI, common.Hex2Bytes(code), backend)
	if err != nil {
		return common.Address{}, err
	}
	if tx == nil {
		return common.Address{}, errors.New("expected a transaction, got nil")
	}
	if (address == common.Address{}) {
		return common.Address{}, errors.New("expected a valid contract address, got the zero address")
	}

	return address, nil
}

func enterShortlist(t *testing.T, opts *bind.TransactOpts, validator, stakesContract common.Address,
	stakesABI, validatorsABI abi.ABI, backend bind.ContractBackend, chainId *big.Int, key *ecdsa.PrivateKey,
) {
	addStakeInput, err := stakesABI.Pack("add", validator, big.NewInt(1000000))
	if err != nil {
		t.Fatalf("Failed to pack data for someFunction: %v", err)
	}

	nonce, err := backend.PendingNonceAt(context.Background(), opts.From)
	if err != nil {
		t.Fatalf("Failed to retrieve account nonce: %v", err)
	}

	tx := types.NewTransaction(nonce, stakesContract, big.NewInt(0), 1000000, big.NewInt(60557667), addStakeInput)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), key)
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}

	err = backend.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatalf("Failed to send transaction: %v", err)
	}
}
