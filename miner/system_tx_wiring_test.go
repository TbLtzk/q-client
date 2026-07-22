package miner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"math/big"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"gitlab.com/q-dev/q-client/accounts"
	"gitlab.com/q-dev/q-client/accounts/abi"
	"gitlab.com/q-dev/q-client/accounts/keystore"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/consensus"
	"gitlab.com/q-dev/q-client/consensus/ethash"
	"gitlab.com/q-dev/q-client/core/rawdb"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/event"
	"gitlab.com/q-dev/q-client/params"
	"gitlab.com/q-dev/system-contracts/generated"
)

// TestFillTransactionsWiresPrepareSystemTx is a regression guard for the v2.2 geth-merge
// fallout that left SystemTxPreparer as dead code and stopped Validators.makeSnapshot.
// If fillTransactions stops calling prepareSystemTx, this test fails.
func TestFillTransactionsWiresPrepareSystemTx(t *testing.T) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	workerPath := filepath.Join(filepath.Dir(thisFile), "worker.go")

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, workerPath, nil, 0)
	if err != nil {
		t.Fatalf("parse worker.go: %v", err)
	}

	var fillFn *ast.FuncDecl
	ast.Inspect(file, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok || fn.Name == nil || fn.Name.Name != "fillTransactions" {
			return true
		}
		fillFn = fn
		return false
	})
	if fillFn == nil || fillFn.Body == nil {
		t.Fatal("fillTransactions not found in worker.go")
	}

	wired := false
	ast.Inspect(fillFn.Body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		switch fun := call.Fun.(type) {
		case *ast.SelectorExpr:
			if fun.Sel != nil && fun.Sel.Name == "prepareSystemTx" {
				wired = true
				return false
			}
		case *ast.Ident:
			if fun.Name == "prepareSystemTx" {
				wired = true
				return false
			}
		}
		return true
	})
	if !wired {
		t.Fatal("fillTransactions must call prepareSystemTx to inject Validators.makeSnapshot system txs")
	}
}

type stubValidatorEngine struct {
	consensus.Engine
	signer     common.Address
	validators *common.Address
}

func (s *stubValidatorEngine) Signer() common.Address { return s.signer }

func (s *stubValidatorEngine) Validators() *common.Address { return s.validators }

// TestPrepareSystemTxProducesMakeSnapshot verifies the miner helper builds a signed
// Validators.makeSnapshot tx at an epoch boundary (functional counterpart to the AST guard).
func TestPrepareSystemTxProducesMakeSnapshot(t *testing.T) {
	var (
		db     = rawdb.NewMemoryDatabase()
		config = *params.AllCliqueProtocolChanges
	)
	config.Clique = &params.CliqueConfig{Period: 1, Epoch: 101}
	config.Clique.RewardReceiver = common.HexToAddress("92C35a964624D9cbF90c2A0525e116093FAF867E")

	validatorsAddr := common.HexToAddress("0x3Bce2CeeFb1EADb3313CcdcA54108F5aD2fc45DC")
	engine := &stubValidatorEngine{
		Engine:     ethash.NewFaker(),
		signer:     testBankAddress,
		validators: &validatorsAddr,
	}

	dir := t.TempDir()
	ks := keystore.NewKeyStore(dir, keystore.LightScryptN, keystore.LightScryptP)
	account, err := ks.ImportECDSA(testBankKey, "pwd")
	if err != nil {
		t.Fatal(err)
	}
	if err := ks.Unlock(account, "pwd"); err != nil {
		t.Fatal(err)
	}
	am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: true}, ks)

	backend := newTestWorkerBackend(t, ethashChainConfig, ethash.NewFaker(), db, 0)
	w := newWorker(testConfig, &config, engine, backend, new(event.TypeMux), nil, false, am, nil)
	defer w.close()

	parent := backend.chain.CurrentBlock()
	header := &types.Header{
		ParentHash: parent.Hash(),
		Number:     big.NewInt(100), // (100+1)%101 == 0
		GasLimit:   parent.GasLimit,
		Time:       parent.Time + 1,
		Coinbase:   testBankAddress,
		Difficulty: big.NewInt(1),
	}
	env, err := w.makeEnv(parent, header, testBankAddress)
	if err != nil {
		t.Fatalf("makeEnv: %v", err)
	}
	defer env.discard()

	systemTxs := w.prepareSystemTx(am, env)
	lazy := systemTxs[testBankAddress]
	if len(lazy) != 1 {
		t.Fatalf("expected 1 system tx at epoch boundary, got %d", len(lazy))
	}
	tx := lazy[0].Resolve()
	if tx == nil || tx.To() == nil || *tx.To() != validatorsAddr {
		t.Fatalf("unexpected system tx recipient: %v", tx)
	}

	from, err := types.Sender(env.signer, tx)
	if err != nil {
		t.Fatalf("sender: %v", err)
	}
	if from != testBankAddress {
		t.Fatalf("from=%s want=%s", from.Hex(), testBankAddress.Hex())
	}

	parsed, err := abi.JSON(strings.NewReader(generated.ValidatorsABI))
	if err != nil {
		t.Fatal(err)
	}
	want, err := parsed.Pack("makeSnapshot")
	if err != nil {
		t.Fatal(err)
	}
	if string(tx.Data()) != string(want) {
		t.Fatalf("calldata=%x want=%x", tx.Data(), want)
	}
}
