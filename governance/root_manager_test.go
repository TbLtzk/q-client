package governance

import (
	"encoding/json"
	"errors"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/q-dev/q-client/accounts"
	"gitlab.com/q-dev/q-client/accounts/abi/bind"
	"gitlab.com/q-dev/q-client/accounts/keystore"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/consensus/clique"
	"gitlab.com/q-dev/q-client/contracts"
	"gitlab.com/q-dev/q-client/core"
	"gitlab.com/q-dev/q-client/core/rawdb"
	"gitlab.com/q-dev/q-client/core/vm"
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/params"
	"gitlab.com/q-dev/system-contracts/generated"
)

var (
	timestamp     = uint64(1683622303)
	rootAddresses = []common.Address{
		common.HexToAddress("0x6a7C4c452B75E12Ea4417434EbD7Cb9743bfd142"),
		common.HexToAddress("0x18640e7DC08D5C0090e457383aab9F0837aCDF2F"),
		common.HexToAddress("0xae3CB02aADcBBBc24BcEabff7BA552Ceba9d86F8"),
		common.HexToAddress("0x764e55b3E82559f7C52424301e99f9ACbCF73f5d"),
	}
	//nolint:unused,deadcode
	aliases = []common.Address{
		common.HexToAddress("0x08eB380042AB0454998416e481E357203aB6CBe9"), // root1 has alias 0x08eB380042AB0454998416e481E357203aB6CBe9
		common.HexToAddress("0x18640e7DC08D5C0090e457383aab9F0837aCDF2F"),
		common.HexToAddress("0xae3CB02aADcBBBc24BcEabff7BA552Ceba9d86F8"),
		common.HexToAddress("0x764e55b3E82559f7C52424301e99f9ACbCF73f5d"),
	}
	roots []common.Address
)

const defNumAccounts = 4

type exclusionListTestData struct {
	Name    string
	WantErr bool
	List    *common.ValidatorExclusionList
	Force   bool
}

type rootListTestData struct {
	Name    string
	WantErr bool
	Set     *rootSet
	Force   bool
}

func tmpDirName(t *testing.T) string {
	d := t.TempDir()
	d, err := filepath.EvalSymlinks(d)
	if err != nil {
		t.Fatal(err)
	}
	return d
}

func createRootAccounts(t *testing.T, ks *keystore.KeyStore, numAccounts int) []accounts.Account {
	var accs []accounts.Account
	for i := 0; i < numAccounts; i++ {
		a1, err := ks.NewAccount("")
		if err != nil {
			t.Fatal(err)
		}
		if err := ks.Unlock(a1, ""); err != nil {
			t.Fatal(err)
		}
		accs = append(accs, a1)
	}

	return accs
}

func rootNodes(t *testing.T, ks *keystore.KeyStore) []common.Address {
	var addrs []common.Address
	for _, account := range ks.Accounts() {
		addrs = append(addrs, account.Address)
	}
	return addrs
}

func createKeystore(t *testing.T, dir string, createAccounts bool, numAccounts int) *keystore.KeyStore {
	var (
		n, p = keystore.StandardScryptN, keystore.StandardScryptP
	)
	ks := keystore.NewKeyStore(dir, n, p)

	if createAccounts {
		createRootAccounts(t, ks, numAccounts)
	}

	return ks
}

func createAccountManager(t *testing.T, ks *keystore.KeyStore) *accounts.Manager {
	var backends []accounts.Backend

	backends = append(backends, ks)
	manager := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: true}, backends...)

	return manager
}

func createGovConfig(t *testing.T, ks *keystore.KeyStore, isRootNode bool) Config {
	var cfg Config

	var rns []common.Address
	if isRootNode {
		rns = rootNodes(t, ks)
	} else {
		rns = rootAddresses
	}

	cfg.RootList = common.RootList{
		Timestamp: timestamp,
		Nodes:     rns,
	}
	return cfg
}

func newTestRootManager(t *testing.T, isRootNode bool) *RootManager {
	return newTestRootManagerWithAccounts(t, isRootNode, defNumAccounts)
}

func newTestRootManagerWithAccounts(t *testing.T, isRootNode bool, numAccounts int) *RootManager {
	dataDir := tmpDirName(t)
	ks := createKeystore(t, dataDir, isRootNode, numAccounts)

	govCfg := createGovConfig(t, ks, isRootNode)
	accMgr := createAccountManager(t, ks)

	if isRootNode {
		roots = rootNodes(t, ks)
	} else {
		roots = rootAddresses
	}

	rm, err := NewRootManager(accMgr, 35443, dataDir, &govCfg)
	if err != nil {
		t.Fatalf("Can't create RootManager: %v", err)
	}
	exSet, err := newExclusionSet(activeExclusionList)
	if err != nil {
		t.Fatalf("Can't create RootManager: %v", err)
	}
	rm.db.saveActiveExclusionSet(exSet)
	rm.activeExSet = exSet

	//Active root set
	rootList := common.RootList{
		Timestamp: 1680255000,
		Nodes:     roots,
	}
	rSet, err := newRootSet(&rootList)
	if err != nil {
		t.Fatalf("Can't create RootManager: %v", err)
	}
	rm.db.saveActiveRootSet(rSet)

	reg := newTestRegistry()
	rm.InitRegistry(reg)

	//InitBlockchain should be also present here, but it's not necessary for all tests, so init it manually

	return rm
}

func newTestRegistry() *contracts.Registry {
	reg := testRegistry{}
	return (*contracts.Registry)(&reg)
}

type testRegistry contracts.Registry
type testRoots generated.Roots
type testAliases generated.AccountAliases

func (r *testRegistry) Roots() *generated.Roots {
	var tRoots testRoots
	return (*generated.Roots)(&tRoots)
}

func (r *testRegistry) AccountAliases() *generated.AccountAliases {
	var tAliases testAliases
	return (*generated.AccountAliases)(&tAliases)
}

func (testRoots) GetMembers(opts *bind.CallOpts) ([]common.Address, error) {
	return roots, nil
}

func TestNewRootManager(t *testing.T) {
	type args struct {
		isRootNode bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "New root manager for root node",
			args:    args{isRootNode: true},
			wantErr: false,
		},
		{
			name:    "New root manager for non-root node",
			args:    args{isRootNode: false},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := newTestRootManager(t, tt.args.isRootNode)
			if rm == nil && !tt.wantErr {
				t.Errorf("NewRootManager() wantErr %v", tt.wantErr)
			}
		})
	}
}

func TestSignRootSet(t *testing.T) {
	type args struct {
		isRootNode bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Sign root set by root node",
			args:    args{isRootNode: true},
			wantErr: false,
		},
		{
			name:    "Sign root set by non-root node",
			args:    args{isRootNode: false},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := newTestRootManager(t, tt.args.isRootNode)

			rootList := common.RootList{
				Timestamp:  timestamp,
				Nodes:      rootAddresses,
				Signatures: nil,
			}
			newSet, err := newRootSet(&rootList)
			if err != nil {
				t.Fatalf("Can't create root set: %v", err)
			}
			signed := rm.signRootSet(newSet)
			if signed == tt.wantErr {
				t.Errorf("signRootSet() signed = %v, wantErr %v", signed, tt.wantErr)
			}
		})
	}
}

func TestIsRootNode(t *testing.T) {
	type args struct {
		isRootNode bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Is a root node",
			args:    args{isRootNode: true},
			wantErr: false,
		},
		{
			name:    "Not a root node",
			args:    args{isRootNode: false},
			wantErr: true,
		},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := newTestRootManager(t, tt.args.isRootNode)
			isRn := rm.isRootNode(tt.args.isRootNode)
			if isRn == tt.wantErr {
				t.Errorf("isRootNode() isRootNode = %v, wantErr %v", isRn, tt.wantErr)
			}
		})
	}
}

func TestSignExclusionSet(t *testing.T) {
	type args struct {
		isRootNode bool
		list       *common.ValidatorExclusionList
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Is a root node. Sign correct exclusion list",
			args: args{
				isRootNode: true,
				list:       correctExclusionList,
			},
			wantErr: false,
		},
		{
			name: "Not a root node. Sign correct exclusion list",
			args: args{
				isRootNode: false,
				list:       correctExclusionList,
			},
			wantErr: true,
		},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := newTestRootManager(t, tt.args.isRootNode)
			set, err := newExclusionSet(tt.args.list)
			if err != nil {
				t.Errorf("Can't create exclusion set: %v", err)
			}
			signed := rm.signExclusionSet(set)
			if signed == tt.wantErr {
				t.Errorf("signExclusionSet() signed = %v, wantErr %v", signed, tt.wantErr)
			}
		})
	}
}

func TestUpgradeExclusionSet(t *testing.T) {
	type args struct {
		list *common.ValidatorExclusionList
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Correct exclusion list. Upgrade active exclusion set",
			args: args{
				list: correctExclusionList,
			},
			wantErr: false,
		}, {
			name: "Upgrade with active set",
			args: args{
				list: correctExclusionList,
			},
			wantErr: true,
		},
		{
			name: "Obsolete exclusion list. Upgrade active exclusion set",
			args: args{
				list: obsoleteExclusionList,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := newTestRootManager(t, false)

			bc := newTestChain(t, rm)
			defer bc.Stop()

			rm.InitBlockChain(bc)
			set, err := newExclusionSet(tt.args.list)
			if err != nil {
				t.Errorf("Can't create exclusion set: %v", err)
			}

			currentActiveSet := rm.activeExSet
			rm.upgradeExclusionSet(set, false)
			if currentActiveSet == rm.activeExSet && !tt.wantErr {
				t.Errorf("upgradeExclusionSet() active exclusion set not changed")
			}
		})
	}
}

func loadTestExclusionSetsData(t *testing.T) []exclusionListTestData {
	bytes, err := os.ReadFile("./testdata/exclusion_lists.json")
	if err != nil {
		t.Errorf("Can't read exclusion list file: %v", err)
	}
	var testData []exclusionListTestData
	err = json.Unmarshal(bytes, &testData)
	if err != nil {
		t.Errorf("Can't unmarshal exclusion list file: %v", err)
	}
	return testData
}

func loadTestRootSetsData(t *testing.T) []rootListTestData {
	bytes, err := os.ReadFile("./testdata/root_lists.json")
	if err != nil {
		t.Errorf("Can't read root list file: %v", err)
	}
	var testData []rootListTestData
	err = json.Unmarshal(bytes, &testData)
	if err != nil {
		t.Errorf("Can't unmarshal root list file: %v", err)
	}
	return testData
}

func TestValidateExclusionSet(t *testing.T) {
	testData := loadTestExclusionSetsData(t)

	rm := newTestRootManager(t, false)

	bc := newTestChain(t, rm)
	defer bc.Stop()

	rm.InitBlockChain(bc)

	for _, tt := range testData {
		t.Run(tt.Name, func(t *testing.T) {
			set, err := newExclusionSet(tt.List)
			if err != nil {
				t.Errorf("Can't create exclusion set: %v, test name:%v", err, tt.Name)
			}
			if set.hash == rm.activeExSet.hash {
				return
			}
			err = rm.validateExclusionSet(set)
			if err != nil && !tt.WantErr {
				t.Errorf("validateExclusionSet() error = %v, wantErr %v", err, tt.WantErr)
			}
		})
	}
}

func TestProposeExclusionSet(t *testing.T) {
	testData := loadTestExclusionSetsData(t)

	rm := newTestRootManager(t, true)

	bc := newTestChain(t, rm)
	defer bc.Stop()

	rm.InitBlockChain(bc)

	for _, tt := range testData {
		t.Run(tt.Name, func(t *testing.T) {
			set, err := newExclusionSet(tt.List)
			if err != nil {
				t.Errorf("Can't create exclusion set: %v, test name:%v", err, tt.Name)
			}
			if set.hash == rm.activeExSet.hash {
				return
			}
			_, err = rm.proposeExclusionSet(set, tt.Force)
			if err != nil && !tt.WantErr {
				if tt.Name == "Correct exclusion list" && errors.Is(err, errProposedExclusionListObsolete) {
					//Special case. Obsolete exclusion list is not an error. List becomes active after proposal, and upgrade fails
					//TODO Possibly, modify the amount of root nodes to avoid automatic upgrade
					return
				}
				t.Errorf("proposeExclusionSet() error = %v, wantErr %v", err, tt.WantErr)
			}
		})
	}
}

func TestAcceptProposedExclusionSet(t *testing.T) {
	testData := loadTestExclusionSetsData(t)

	rm := newTestRootManager(t, true)

	bc := newTestChain(t, rm)
	defer bc.Stop()

	rm.InitBlockChain(bc)

	for _, tt := range testData {
		t.Run(tt.Name, func(t *testing.T) {
			set, err := newExclusionSet(tt.List)
			if err != nil {
				t.Errorf("Can't create exclusion set: %v, test name:%v", err, tt.Name)
			}

			exSet, err := newExclusionSet(activeExclusionList)
			if err != nil {
				t.Fatalf("Can't create new exclusion set: %v", err)
			}
			rm.db.saveActiveExclusionSet(exSet)
			rm.activeExSet = exSet

			if set.hash == rm.activeExSet.hash {
				return
			}

			rm.proposedExSet = set

			err = rm.acceptProposedExclusionList(false)

			if tt.WantErr && err == nil {
				t.Errorf("acceptProposedExclusionSet() error = %v, wantErr %v", err, tt.WantErr)
			}
			if err != nil && !tt.WantErr {
				t.Errorf("acceptProposedExclusionSet() error = %v, wantErr %v", err, tt.WantErr)
			}
		})
	}
}

func TestIsAcceptableExclusionSet(t *testing.T) {
	//testData := loadTestExclusionSetsData(t)

	rm := newTestRootManager(t, true)

	set, err := newExclusionSet(correctExclusionList)
	if err != nil {
		t.Errorf("Can't create exclusion set: %v", err)
	}
	assert.Falsef(t, rm.isAcceptableExclusionSet(set), "exclusion set must be not acceptable")
	rm.signExclusionSet(set)
	assert.Truef(t, rm.isAcceptableExclusionSet(set), "isAcceptableExclusionSet() must be acceptable")
}

func TestProposeRootSet(t *testing.T) {
	//First one is simple, non RN should not be able to propose root set
	testProposeRootSetByNonRN(t)
	testProposeRootSetByRN(t)
}

func testProposeRootSetByRN(t *testing.T) {
	var testData []rootListTestData
	rm := newTestRootManager(t, true)
	active := rm.active

	//Following test cases are synthetic, as RM must have root account unlocked to propose root set
	//Following test cases are synthetic, as RM must have root account unlocked to propose root set
	setGreater := active.copy()
	setGreater.timestamp = setGreater.timestamp + 1
	setGreater.hash = setGreater.calcHash()
	testData = append(testData, rootListTestData{
		Name:    "Root set with a greater timestamp",
		Set:     setGreater,
		WantErr: false,
	})

	setLess := active.copy()
	setLess.timestamp = 0
	setLess.hash = setLess.calcHash()
	testData = append(testData, rootListTestData{
		Name:    "Root set with a lesser timestamp",
		Set:     setLess,
		WantErr: true,
	})

	setHuge := active.copy()
	setHuge.timestamp = 9999999999999999999
	setHuge.hash = setHuge.calcHash()
	testData = append(testData, rootListTestData{
		Name:    "Root set with an invalid timestamp",
		Set:     setHuge,
		WantErr: true,
	})

	setSame := active.copy()
	testData = append(testData, rootListTestData{
		Name:    "Root set  with the same timestamp and hash, meaningless",
		Set:     setSame,
		WantErr: true,
	})

	for _, tt := range testData {
		t.Run(tt.Name, func(t *testing.T) {
			reinitializeRootLists(rm, active)

			_, err := rm.proposeRootSet(tt.Set, tt.Force)
			if !tt.WantErr && err != nil {
				t.Errorf("proposeRootSet() error = %v, wantErr %v", err, tt.WantErr)
			}
		})
	}
}

// should not be able to propose root set by non RN
func testProposeRootSetByNonRN(t *testing.T) {
	testData := loadTestRootSetsData(t)
	rm := newTestRootManager(t, false)
	//no need to have blockchain instance here, skip creating it

	_, err := rm.proposeRootSet(testData[0].Set, testData[0].Force)
	if err == nil {
		t.Errorf("propose root set by non RN no error, but want error")
	}
}

func TestAcceptProposedRootSet(t *testing.T) {
	testAcceptProposedRootListByNonRN(t)
	testAcceptProposedRootListByRN(t)
}

// should not be able to propose root set by non RN
func testAcceptProposedRootListByNonRN(t *testing.T) {
	testData := loadTestRootSetsData(t)
	rm := newTestRootManager(t, false)

	list := testData[0].Set.makeList()
	rSet, err := newRootSet(&list)
	if err != nil {
		t.Fatalf("Can't create new root set: %v", err)
	}
	_, err = rm.proposeRootSet(rSet, testData[0].Force)
	if err == nil {
		t.Errorf("accept proposed root set by non RN no error, but want error")
	}
}

func testAcceptProposedRootListByRN(t *testing.T) {
	rm := newTestRootManager(t, true)

	active := rm.active.copy()

	var testData []rootListTestData

	//Following test cases are synthetic, as RM must have root account unlocked to propose root set
	setGreater := active.copy()
	setGreater.timestamp = setGreater.timestamp + 1
	setGreater.hash = setGreater.calcHash()
	testData = append(testData, rootListTestData{
		Name:    "Root set with a greater timestamp",
		Set:     setGreater,
		WantErr: false,
	})

	setLess := active.copy()
	setLess.timestamp = 0
	setLess.hash = setLess.calcHash()
	testData = append(testData, rootListTestData{
		Name:    "Root set with a lesser timestamp",
		Set:     setLess,
		WantErr: true,
	})

	//It might seem that here should be a test case with a greater timestamp,
	//but since verification of the set isn't done in acceptProposedRootList, it's not needed.

	for _, tt := range testData {
		t.Run(tt.Name, func(t *testing.T) {
			reinitializeRootLists(rm, active)

			rm.proposed = tt.Set

			err := rm.acceptProposedRootList(false)
			if (rm.active.hash == tt.Set.hash && tt.WantErr) || (!tt.WantErr && err != nil) {
				t.Errorf("proposeRootSet() error = %v, wantErr %v", err, tt.WantErr)
			}
		})
	}
}

func reinitializeRootLists(rm *RootManager, active *rootSet) {
	rm.active = active
	rm.proposed = nil
	rm.desired = nil
}

func newTestChain(t *testing.T, rm *RootManager) *core.BlockChain {
	var (
		db     = rawdb.NewMemoryDatabase()
		key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
		addr   = crypto.PubkeyToAddress(key.PublicKey)
		engine = clique.New(params.AllCliqueProtocolChanges.Clique, db, rm, rm.reg)
		//signer      = new(types.HomesteadSigner)
		extraVanity = 32
		extraSeal   = crypto.SignatureLength
		diffInTurn  = big.NewInt(2) // Block difficulty for in-turn signatures
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

	config := params.AllCliqueProtocolChanges
	config.AthosBlock = new(big.Int).SetUint64(1)

	chain, _ := core.NewBlockChain(db, nil, config, engine, vm.Config{}, nil, nil)

	blocks, _ := core.GenerateChain(config, genesis, engine, db, 5000, func(i int, block *core.BlockGen) {
		block.SetDifficulty(diffInTurn)
	})
	for i, block := range blocks {
		header := block.Header()
		if i > 0 {
			header.ParentHash = blocks[i-1].Hash()
		}
		header.Extra = make([]byte, extraVanity+extraSeal)
		header.Difficulty = diffInTurn
		header.Coinbase = rm.reg.RewardReceiver()

		sig, _ := crypto.Sign(clique.SealHash(header).Bytes(), key)
		copy(header.Extra[len(header.Extra)-extraSeal:], sig)
		blocks[i] = block.WithSeal(header)
	}

	if _, err := chain.InsertChain(blocks[:5000]); err != nil {
		t.Fatalf("failed to insert initial blocks: %v", err)
	}

	return chain
}

func randomExclusionSet(t *testing.T, add int64) exclusionSet {
	var set exclusionSet

	set.timestamp = uint64(time.Now().Unix() + add) //This will help to have different exclusion sets
	set.signers = make(map[common.Address][]byte)
	for i := 0; i < 10; i++ {
		set.addresses = append(set.addresses, common.Address{byte(i + int(add))})
	}
	set.hash = set.calcHash()
	return set
}

func TestSetQuota(t *testing.T) {
	t.Run("testQuotaExceedsByNumber", func(t *testing.T) {
		testQuotaExceedsByNumber(t)
	})
	t.Run("testQuotaExceedsWithExpiredEntries", func(t *testing.T) {
		testQuotaExceedsWithExpiredEntries(t)
	})
}

func testQuotaExceedsByNumber(t *testing.T) {
	rm := newTestRootManagerWithAccounts(t, true, 1)
	rm.ProposalQuotaMax = 3
	rm.ProposalQuotaTimeWindow = 24 * time.Hour

	//Last set will overflow the quota
	for i := 0; i < int(rm.ProposalQuotaMax+1); i++ {
		set := randomExclusionSet(t, int64(rand.Intn(1000)))
		rm.signExclusionSet(&set)

		is, err := rm.isSetQuotaExceeded(&set)
		if err != nil {
			t.Errorf("isSetQuotaExceeded() error = %v", err)
		}
		if i > int(rm.ProposalQuotaMax) && !is {
			t.Errorf("isSetQuotaExceeded() = %v, want %v", is, true)
		}
	}
}

func testQuotaExceedsWithExpiredEntries(t *testing.T) {
	rm := newTestRootManagerWithAccounts(t, true, 1)
	rm.ProposalQuotaMax = 3
	rm.ProposalQuotaTimeWindow = 30 * time.Second

	//Create 3 sets that will be added to the quota storage
	//Last set will overflow the quota
	for i := 0; i < 5; i++ {
		set := randomExclusionSet(t, int64(i))
		rm.signExclusionSet(&set)

		//Will add set to the quota storage
		rm.isSetQuotaExceeded(&set)
	}

	//Need to wait till time window expires
	time.Sleep(rm.ProposalQuotaTimeWindow + 1)

	//This set shouldn't cause quota overflow
	set := randomExclusionSet(t, 6)
	rm.signExclusionSet(&set)

	is, err := rm.isSetQuotaExceeded(&set)
	if err != nil {
		t.Errorf("isSetQuotaExceeded() error = %v", err)
	}
	assert.False(t, is, "testQuotaExceedsWithExpiredEntries. Set must exceed the quota but it doesn't")
}

func TestQuarantineExclusionSet(t *testing.T) {
	earliestBlock := uint64(2)

	t.Log("Starting root manager with quarantine enabled")
	rm := newTestRootManager(t, true)
	t.Log("Initializing blockchain with 5000 blocks size")
	bc := newTestChain(t, rm)
	defer bc.Stop()
	rm.InitBlockChain(bc)

	t.Log("Setting max rewind limit to 100 blocks")
	rm.setMaxRewindLimit(100)

	t.Log("Creating violating exclusion set")

	list := rm.activeExSet.copy().makeList()
	list.Timestamp = list.Timestamp + 1
	list.Signatures = nil
	list.Hash = common.Hash{}
	badRecord := common.ExcludedValidator{
		Address: common.HexToAddress("0x2B01035cDa82a02fb135EBd68676Fa17FdcAD365"),
		Block:   earliestBlock,
	}
	list.Validators = append(list.Validators, badRecord)
	set, err := newExclusionSet(&list)
	if err != nil {
		t.Fatalf("Can't create new test exclusion set: %v", err)
	}

	t.Log("Signing violating exclusion set")
	rm.signExclusionSet(set)

	t.Log("Checking if created exclusion set meets quarantine criteria")
	assert.Truef(t, rm.isExclusionSetMeetsQuarantineCriteria(set, earliestBlock), "Created exclusion set must be quarantined")

	t.Log("Creating test governance engine")
	gov, err := New(rm, tmpDirName(t))
	if err != nil {
		t.Fatalf("Failed to create Governance: %v", err)
	}
	err = gov.Start()
	if err != nil {
		t.Fatalf("Failed to start Governance: %v", err)
	}
	p := peer{
		id: "test",
	}

	t.Log("Handling exclusion set")
	currentExclusionSet := rm.activeExSet.copy()
	err = gov.handler.handleExclusionSet(&p, set)
	if err != nil {
		t.Fatalf("Failed to handle exclusion list: %v", err)
	}
	assert.NotEqualf(t, rm.activeExSet.hash, set.hash, "Active exclusion set must be changed!")
	if currentExclusionSet.hash != rm.activeExSet.hash {
		t.Fatalf("Active exclusion set must be changed!")
	}

	t.Log("Checking if incoming exclusion set is in quarantine")
	qSets, err := rm.db.getExclusionSetsFromQuarantine()
	if err != nil {
		t.Fatalf("Failed to get exclusion sets from quarantine: %v", err)
	}
	assert.NotNil(t, qSets, "Quarantine must not be empty")
	assert.Lenf(t, qSets, 1, "Quarantine must contain one exclusion set")
	if (qSets)[0].hash != set.hash {
		t.Fatalf("Quarantine must contain incoming exclusion set")
	}

	t.Log("Accepting exclusion set from quarantine")
	err = rm.acceptQuarantinedExclusionSet(&set.hash)
	if err != nil {
		t.Fatalf("Failed to accept quarantined exclusion list: %v", err)
	}
	if rm.activeExSet.hash != set.hash {
		t.Fatalf("Quarantined exclusion set must be applied")
	}

	t.Log("Checking if incoming exclusion set is removed from quarantine")
	qSets, err = rm.db.getExclusionSetsFromQuarantine()
	if err != nil {
		t.Fatalf("Failed to get exclusion sets from quarantine: %v", err)
	}
	assert.Lenf(t, qSets, 0, "Quarantine must be empty after exclusion set is accepted")
}
