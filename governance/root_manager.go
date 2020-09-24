package governance

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/go-ethereum/accounts"
	"gitlab.com/q-dev/go-ethereum/common"
	"gitlab.com/q-dev/go-ethereum/event"
	"gitlab.com/q-dev/go-ethereum/log"
	"gitlab.com/q-dev/go-ethereum/params"
)

var (
	errInvalidSignature = errors.New("list contains invalid signature")
	errHashMismatch     = errors.New("hash mismatch")
)

// RootManager governs root nodes.
type RootManager struct {
	keystore Keystore

	datadir string

	currentFilePath     string
	pendingRootListPath string
	targetRootListPath  string

	targetRootFeed *event.Feed

	lock    sync.Mutex
	isRoot  bool
	current *rootSet
	target  *rootSet

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
	currentListPath := filepath.Join(datadir, currentRootListFname)
	list, err := chooseInitList(currentListPath)
	if err != nil {
		return nil, err
	}

	currentSet, err := newRootSet(list)
	if err != nil {
		return nil, errors.Wrap(err, "invalid current root node list")
	}

	pendingRootListPath := filepath.Join(datadir, pendingRootListFname)
	pendingSet, err := loadRootListFile(filepath.Join(datadir, pendingRootListFname))
	if err != nil {
		return nil, errors.Wrap(err, "failed to load pending root node list")
	}

	exList, err := loadExclusionListFile(filepath.Join(datadir, currentExclusionListFname))
	if err != nil {
		return nil, errors.Wrap(err, "failed to load current exclusion list")
	}

	pendingExList, err := loadExclusionListFile(filepath.Join(datadir, pendingExclusionListFname))
	if err != nil {
		return nil, errors.Wrap(err, "failed to load pending exclusion list")
	}

	manager := &RootManager{
		keystore:            ks,
		datadir:             datadir,
		currentFilePath:     currentListPath,
		pendingRootListPath: pendingRootListPath,
		targetRootListPath:  filepath.Join(datadir, targetRootListFname),
		targetRootFeed:      &event.Feed{},
		current:             currentSet,
		target:              pendingSet.rootSet,
		targetExListFeed:    &event.Feed{},
		exclusionSet:        exList.set,
		targetExSet:         pendingExList.set,
	}

	return manager, nil
}

func (s *RootManager) CurrentList() common.RootList {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.current.makeList()
}

// ExclusionSet returns set of excluded validators addresses.
func (s *RootManager) ExclusionSet() map[common.Address]struct{} {
	s.exLock.Lock()
	defer s.exLock.Unlock()

	set := make(map[common.Address]struct{})
	for addr := range s.exclusionSet.addrToBlock {
		set[addr] = struct{}{}
	}

	return set
}

// the reason why it's here is because accounts are unlocked after
// service are initialised and it's not that easy to change this
// todo: listen to wallet events instead
func (s *RootManager) run() error {
	go s.watchTargetRootList(time.Time{})
	go s.refreshFile(refreshPeriod, s.pendingRootListPath, func() interface{} {
		set := s.targetRootSet()
		if set == nil {
			return nil
		}

		return set.makeList()
	})

	go s.refreshFile(refreshPeriod, s.currentFilePath, func() interface{} {
		return s.CurrentList()
	})

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
		_ = os.Remove(s.pendingRootListPath)
	}

	return true
}

func (s *RootManager) targetRootSet() *rootSet {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.target
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

func (s *RootManager) watchTargetRootList(updatedAt time.Time) {
	for {
		time.Sleep(time.Second)

		file, err := loadRootListFile(s.targetRootListPath)
		if err != nil {
			log.Warn("failed to load root list file", "err", err)
			continue
		}

		if file.rootSet == nil {
			continue
		}

		if !file.updatedAt.After(updatedAt) {
			continue
		}

		updatedAt = file.updatedAt
		log.Debug("DETECTED target root list")

		current := s.currentRootSet()
		if file.rootSet.timestamp < current.timestamp {
			log.Debug("skipping target root list because of timestamp")
			continue
		}

		target := s.targetRootSet()
		if target != nil && target.hash == file.rootSet.hash {
			log.Debug("skipping identical target root list")
			continue
		}

		if !s.signRootSet(file.rootSet) {
			continue
		}

		s.lock.Lock()
		s.target = file.rootSet
		s.lock.Unlock()

		go func(target *rootSet) {
			// todo: a way to cancel this?
			timestamp := int64(target.timestamp)
			if duration := time.Until(time.Unix(timestamp, 0)); duration > 0 {
				log.Info("scheduled target root list distribution", "timestamp", timestamp)
				t := time.NewTimer(duration)
				defer t.Stop()

				<-t.C
			}

			if n := s.targetRootFeed.Send(target); n == 0 {
				panic("handler didn't receive new set")
			}
		}(file.rootSet)
	}
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

// choose the most recent list amongst default list and stored locally.
func chooseInitList(fpath string) (*common.RootList, error) {
	defaultList := params.DevnetRootNodes

	file, err := os.Open(fpath)
	defer func() { _ = file.Close() }()

	switch true {
	case err == nil:
		var list common.RootList
		if err := json.NewDecoder(file).Decode(&list); err != nil {
			return nil, errors.Wrap(err, "malformed current root node list file")
		}

		if list.Timestamp < defaultList.Timestamp {
			return &defaultList, nil
		}

		return &list, nil
	case os.IsNotExist(err):
		return &defaultList, nil
	default:
		return nil, errors.Wrapf(err, "failed to open current root node list file at %s", fpath)
	}
}

type rootListFile struct {
	updatedAt time.Time
	rootSet   *rootSet
}

func loadRootListFile(fpath string) (*rootListFile, error) {
	file, err := os.Open(fpath)
	defer func() { _ = file.Close() }()

	if os.IsNotExist(err) {
		return &rootListFile{}, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}

	var list common.RootList
	if err := json.NewDecoder(file).Decode(&list); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal file")
	}

	// todo: handle properly
	list.Hash = common.Hash{}

	set, err := newRootSet(&list)
	if err != nil {
		return nil, errors.Wrap(err, "invalid root list")
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get file stat")
	}

	return &rootListFile{rootSet: set, updatedAt: stat.ModTime()}, nil
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
	currentRootListFname = "root-list.json"
	pendingRootListFname = "pending-root-list.json"
	targetRootListFname  = "target-root-list.json"
)

const (
	currentExclusionListFname = "exclusion-list.json"
	pendingExclusionListFname = "pending-exclusion-list.json"
	targetExclusionListFname  = "target-exclusion-list.json"
)

const refreshPeriod = time.Second
