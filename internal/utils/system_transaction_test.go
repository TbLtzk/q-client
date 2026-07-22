package utils

import (
	"math/big"
	"strings"
	"testing"

	"gitlab.com/q-dev/q-client/accounts"
	"gitlab.com/q-dev/q-client/accounts/abi"
	"gitlab.com/q-dev/q-client/accounts/keystore"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/consensus"
	"gitlab.com/q-dev/q-client/consensus/ethash"
	"gitlab.com/q-dev/q-client/core/rawdb"
	"gitlab.com/q-dev/q-client/core/state"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/params"
	"gitlab.com/q-dev/system-contracts/generated"
)

type stubValidatorEngine struct {
	consensus.Engine
	signer     common.Address
	validators *common.Address
}

func (s *stubValidatorEngine) Signer() common.Address { return s.signer }

func (s *stubValidatorEngine) Validators() *common.Address { return s.validators }

func TestPrepareSystemTxAtEpochBoundary(t *testing.T) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	signerAddr := crypto.PubkeyToAddress(key.PublicKey)
	validatorsAddr := common.HexToAddress("0x3Bce2CeeFb1EADb3313CcdcA54108F5aD2fc45DC")

	dir := t.TempDir()
	ks := keystore.NewKeyStore(dir, keystore.LightScryptN, keystore.LightScryptP)
	account, err := ks.ImportECDSA(key, "pwd")
	if err != nil {
		t.Fatal(err)
	}
	if err := ks.Unlock(account, "pwd"); err != nil {
		t.Fatal(err)
	}
	am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: true}, ks)

	db := rawdb.NewMemoryDatabase()
	statedb, err := state.New(types.EmptyRootHash, state.NewDatabase(db), nil)
	if err != nil {
		t.Fatal(err)
	}

	cfg := &params.ChainConfig{
		ChainID: big.NewInt(35443),
		Clique:  &params.CliqueConfig{Period: 5, Epoch: 101},
	}
	engine := &stubValidatorEngine{
		Engine:     ethash.NewFaker(),
		signer:     signerAddr,
		validators: &validatorsAddr,
	}

	makeHeader := func(number uint64) *types.Header {
		return &types.Header{Number: big.NewInt(int64(number)), Time: 1}
	}

	packSelector := func() []byte {
		a, err := abi.JSON(strings.NewReader(generated.ValidatorsABI))
		if err != nil {
			t.Fatal(err)
		}
		input, err := a.Pack("makeSnapshot")
		if err != nil {
			t.Fatal(err)
		}
		return input
	}()

	t.Run("epoch boundary produces makeSnapshot", func(t *testing.T) {
		// (100+1)%101 == 0 → epoch-end block
		header := makeHeader(100)
		preparer := New(cfg, engine, statedb, header, types.LatestSigner(cfg), nil)
		txs := preparer.PrepareSystemTx(am, nil)
		lazy := txs[signerAddr]
		if len(lazy) != 1 {
			t.Fatalf("expected 1 system tx, got %d (accounts=%v)", len(lazy), len(txs))
		}
		tx := lazy[0].Resolve()
		if tx == nil {
			t.Fatal("lazy tx resolved to nil")
		}
		if *tx.To() != validatorsAddr {
			t.Fatalf("to=%s want=%s", tx.To().Hex(), validatorsAddr.Hex())
		}
		if string(tx.Data()) != string(packSelector) {
			t.Fatalf("unexpected calldata: %x", tx.Data())
		}
		if tx.Gas() != 1477210 {
			t.Fatalf("gas=%d", tx.Gas())
		}
	})

	t.Run("non-epoch boundary yields nothing", func(t *testing.T) {
		header := makeHeader(99)
		preparer := New(cfg, engine, statedb, header, types.LatestSigner(cfg), nil)
		txs := preparer.PrepareSystemTx(am, nil)
		if len(txs) != 0 {
			t.Fatalf("expected no system txs mid-epoch, got %d", len(txs))
		}
	})
}
