package governance

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/accounts"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/event"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/params"
)

var (
	errInvalidSignature = errors.New("list contains invalid signature")
	errHashMismatch     = errors.New("hash mismatch")
)

// RootManager governs root nodes.
type RootManager struct {
	keystore Keystore

	datadir        string
	targetRootFeed *event.Feed

	lock     sync.Mutex
	db       *database
	isRoot   bool
	current  *rootSet
	target   *rootSet
	proposed *rootSet

	exLock           sync.Mutex
	targetExListFeed *event.Feed
	exclusionSet     *exclusionSet
	targetExSet      *exclusionSet
}

type Keystore interface {
	IsUnlocked(addr common.Address) bool
	SignHash(a accounts.Account, hash []byte) ([]byte, error)
}

func newRootManager(ks Keystore, datadir string) (*RootManager, error) {
	db, err := newDatabase(filepath.Join(datadir, "gov"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to init gov database")
	}

	currentRootSet, err := db.getCurrentRootSet()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current root set")
	}

	if currentRootSet == nil {
		currentRootSet, _ = newRootSet(&params.DevnetRootNodes)
		log.Info("using predefined root set", "hash", currentRootSet.hash.Hex())
	} else {
		log.Info("using saved root set", "hash", currentRootSet.hash.Hex())
	}

	desiredRootSet, _ := db.getDesiredRootSet()
	proposedRootSet, _ := db.getProposedRootSet()

	exList, err := loadExclusionListFile(filepath.Join(datadir, currentExclusionListFname))
	if err != nil {
		return nil, errors.Wrap(err, "failed to load current exclusion list")
	}

	pendingExList, err := loadExclusionListFile(filepath.Join(datadir, pendingExclusionListFname))
	if err != nil {
		return nil, errors.Wrap(err, "failed to load pending exclusion list")
	}

	manager := &RootManager{
		db:       db,
		keystore: ks,
		datadir:  datadir,

		targetRootFeed: &event.Feed{},
		current:        currentRootSet,
		target:         desiredRootSet,
		proposed:       proposedRootSet,

		targetExListFeed: &event.Feed{},
		exclusionSet:     exList.set,
		targetExSet:      pendingExList.set,
	}

	return manager, nil
}

func (s *RootManager) CurrentList() common.RootList {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.current.makeList()
}

// ExclusionSet returns set of excluded validators addresses.
func (s *RootManager) ExclusionSet() map[common.Address]uint64 {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	set := make(map[common.Address]uint64)
	if s.exclusionSet == nil {
		return set
	}

	for addr, block := range s.exclusionSet.addrToBlock {
		set[addr] = block
	}

	return set
}

// the reason why it's here is because accounts are unlocked after
// service are initialised and it's not that easy to change this
// todo: listen to wallet events instead
func (s *RootManager) run() error {
	go s.watchTargetExclusionList()
	go s.refreshFile(refreshPeriod, path.Join(s.datadir, currentExclusionListFname), func() interface{} {
		set := s.currentExclusionSet()
		if set == nil {
			return nil
		}

		return set.makeList()
	})
	go s.refreshFile(refreshPeriod, path.Join(s.datadir, pendingExclusionListFname), func() interface{} {
		set := s.targetExclusionSet()
		if set == nil {
			return nil
		}

		return set.makeList()
	})

	if s.isRoot {
		log.Info("Node belongs to the current root node current")
	}

	return nil
}

func (s *RootManager) isMember(set []common.Address) bool {
	for _, addr := range set {
		if s.keystore.IsUnlocked(addr) {
			return true
		}
	}

	return false
}

func (s *RootManager) signRootSet(set *rootSet) bool {
	var isMember bool
	for _, addr := range set.rootAddresses {
		if !s.keystore.IsUnlocked(addr) {
			continue
		}

		isMember = true
		signature, err := s.keystore.SignHash(accounts.Account{Address: addr}, set.hash.Bytes())
		if err != nil {
			log.Error("failed to sign root set", "err", err)
			continue
		}

		if set.addSignature(addr, signature) {
			log.Info("signed root list", "hash", set.hash.Hex(), "signer", addr.Hex())
		}
	}

	return isMember
}

func (s *RootManager) signExclusionSet(set *exclusionSet) bool {
	var isSigned bool
	for _, addr := range s.currentRootSet().rootAddresses {
		if !s.keystore.IsUnlocked(addr) {
			continue
		}

		signature, err := s.keystore.SignHash(accounts.Account{Address: addr}, set.hash.Bytes())
		if err != nil {
			log.Error("failed to sign exclusion list", "err", err)
			continue
		}

		set.addSignature(addr, signature)
		isSigned = true
		log.Info("signed exclusion list", "hash", set.hash.Hex(), "signer", addr.Hex())
	}

	return isSigned
}

func (s *RootManager) findUnlocked(set *rootSet) []common.Address {
	var unlocked []common.Address
	for _, addr := range set.rootAddresses {
		if s.keystore.IsUnlocked(addr) {
			unlocked = append(unlocked, addr)
		}
	}

	return unlocked
}

func (s *RootManager) refreshFile(period time.Duration, fpath string, getter func() interface{}) {
	for {
		time.Sleep(period)

		data := getter()
		if data == nil {
			continue
		}

		f, err := os.Create(fpath)
		if err != nil {
			log.Warn("failed to refresh file", "fpath", fpath, "err", err)
			continue
		}

		if err := json.NewEncoder(f).Encode(data); err != nil {
			log.Warn("failed to encode file to json", "fpath", fpath, "err", err)
			continue
		}

		_ = f.Close()
	}
}

func (s *RootManager) tryUpgradeExclusionList(set *exclusionSet) bool {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	if !s.currentRootSet().isAcceptableExclusionSet(set) {
		return false
	}

	if s.exclusionSet != nil && set.timestamp <= s.exclusionSet.timestamp {
		return false
	}

	s.exclusionSet = set
	if s.targetExSet != nil && s.targetExSet.timestamp <= set.timestamp {
		s.targetExSet = nil
		_ = os.Remove(filepath.Join(s.datadir, pendingExclusionListFname))
	}

	return true
}

// unsafe for concurrent calls.
// must be locked before calling
func (s *RootManager) upgradeRootSet(set *rootSet) {
	s.current = set
	s.db.saveCurrentRootSet(s.current)

	s.isRoot = s.isMember(s.current.rootAddresses)

	log.Info("Upgraded root list", "hash", set.hash.Hex(), "timestamp", set.timestamp)

	if s.target == nil || s.target.timestamp > set.timestamp {
		return
	}

	if s.target.hash != set.hash {
		log.Info("Dropping outdated desired root list",
			"hash", s.target.hash.Hex(),
			"timestamp", s.target.timestamp)
	}

	s.target = nil
	s.db.deleteDesiredRootSet()
}

// upgrade to a new root set.
// returns true if set is accepted.
func (s *RootManager) upgrade(set *rootSet) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	// skip a concurrent call with an obsolete list
	if set.timestamp <= s.current.timestamp {
		return false
	}

	s.current = set

	// clean up tmp storage
	if s.target != nil && s.target.timestamp <= set.timestamp {
		s.target = nil
	}

	return true
}

func (s *RootManager) proposedRootSet() *rootSet {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.proposed
}

func (s *RootManager) targetRootSet() *rootSet {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.target == nil {
		return nil
	}

	return s.target.copy()
}

func (s *RootManager) currentRootSet() *rootSet {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.current
}

func (s *RootManager) currentExclusionSet() *exclusionSet {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	return s.exclusionSet
}

func (s *RootManager) targetExclusionSet() *exclusionSet {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	return s.targetExSet
}

func (s *RootManager) watchTargetExclusionList() {
	var updatedAt time.Time
	for {
		time.Sleep(time.Second)

		file, err := loadExclusionListFile(path.Join(s.datadir, targetExclusionListFname))
		if err != nil {
			log.Warn("failed to load exclusion list file", "err", err)
			continue
		}

		if file.set == nil {
			continue
		}

		if !file.updatedAt.After(updatedAt) {
			continue
		}

		updatedAt = file.updatedAt
		log.Debug("DETECTED target exclusion list")

		current := s.currentExclusionSet()
		if current != nil && file.set.timestamp < current.timestamp {
			log.Debug("skipping target exclusion list because of timestamp")
			continue
		}

		target := s.targetExclusionSet()
		if target != nil && target.hash == file.set.hash {
			log.Debug("skipping identical target exclusion list")
			continue
		}

		if !s.signExclusionSet(file.set) {
			continue
		}

		s.exLock.Lock()
		s.targetExSet = file.set
		s.exLock.Unlock()

		go func(target *exclusionSet) {
			// todo: a way to cancel this?
			timestamp := int64(target.timestamp)
			if duration := time.Until(time.Unix(timestamp, 0)); duration > 0 {
				log.Info("scheduled target exclusion list distribution", "timestamp", timestamp)
				t := time.NewTimer(duration)
				defer t.Stop()

				<-t.C
			}

			if n := s.targetExListFeed.Send(target); n == 0 {
				panic("handler didn't receive new set")
			}
		}(file.set)
	}
}

type exclusionListFile struct {
	set       *exclusionSet
	updatedAt time.Time
}

func loadExclusionListFile(fpath string) (*exclusionListFile, error) {
	file, err := os.Open(fpath)
	defer func() { _ = file.Close() }()

	if os.IsNotExist(err) {
		return &exclusionListFile{}, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}

	var list common.ValidatorExclusionList
	if err := json.NewDecoder(file).Decode(&list); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal file")
	}

	set, err := newExclusionSet(&list)
	if err != nil {
		return nil, errors.Wrap(err, "invalid root list")
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get file stat")
	}

	return &exclusionListFile{set: set, updatedAt: stat.ModTime()}, nil
}

const (
	currentExclusionListFname = "exclusion-list.json"
	pendingExclusionListFname = "pending-exclusion-list.json"
	targetExclusionListFname  = "target-exclusion-list.json"
)

const refreshPeriod = time.Second
