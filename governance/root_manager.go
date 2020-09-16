package governance

import (
	"encoding/json"
	"os"
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
	errIncomplete       = errors.New("not enough signatures")
	errHashMismatch     = errors.New("hash mismatch")
)

// RootManager governs root nodes.
type RootManager struct {
	keystore Keystore

	currentFilePath     string
	pendingRootListPath string
	targetRootListPath  string

	targetRootFeed *event.Feed
	newRootCh      chan *rootSet

	lock    sync.Mutex
	isRoot  bool
	current *rootSet
	target  *rootSet

	exclusionSet map[common.Address]struct{}
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
	pendingSet, err := loadRootListFile(pendingRootListPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load pending root node list")
	}

	manager := &RootManager{
		keystore:            ks,
		currentFilePath:     currentListPath,
		pendingRootListPath: pendingRootListPath,
		targetRootListPath:  filepath.Join(datadir, targetRootListFname),
		newRootCh:           make(chan *rootSet),
		targetRootFeed:      &event.Feed{},
		current:             currentSet,
		target:              pendingSet.rootSet,
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
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.exclusionSet
}

// the reason why it's here is because accounts are unlocked after
// service are initialised and it's not that easy to change this
// todo: listen to wallet events instead
func (s *RootManager) run() error {
	go s.listenNewRootList()
	go s.watchTargetRootList(time.Time{})
	go s.refreshFile(refreshPeriod, s.pendingRootListPath, func() interface{} {
		set := s.targetSet()
		if set == nil {
			return nil
		}

		return set.makeList()
	})

	go s.refreshFile(refreshPeriod, s.currentFilePath, func() interface{} {
		return s.CurrentList()
	})

	if s.isRoot {
		log.Info("Node belongs to the current root node current")
	}

	return nil
}

func (s *RootManager) listenNewRootList() {
	for set := range s.newRootCh {
		current := s.currentSet()
		if set.timestamp < current.timestamp {
			log.Debug("skipping target root list because of timestamp")
			continue
		}

		target := s.targetSet()
		if target != nil && target.hash == set.hash {
			log.Debug("skipping identical target root list")
			continue
		}

		if !s.signRootSet(set) {
			continue
		}

		s.lock.Lock()
		s.target = set
		s.lock.Unlock()

		go func(target *rootSet) {
			// todo: a way to cancel this?
			timestamp := int64(target.timestamp)
			if duration := time.Unix(timestamp, 0).Sub(time.Now()); duration > 0 {
				log.Info("scheduled target root list distribution", "timestamp", timestamp)
				t := time.NewTimer(duration)
				defer t.Stop()

				<-t.C
			}

			if n := s.targetRootFeed.Send(target); n == 0 {
				panic("handler didn't received new set")
			}
		}(set)
	}
}

func (s *RootManager) watchTargetRootList(updatedAt time.Time) {
	for {
		time.Sleep(time.Second)

		set, err := loadRootListFile(s.targetRootListPath)
		if err != nil {
			log.Warn("failed to load root list file", "err", err)
			continue
		}

		if set.rootSet == nil {
			continue
		}

		if !set.updatedAt.After(updatedAt) {
			continue
		}

		updatedAt = set.updatedAt
		log.Debug("DETECTED target root list")
		s.newRootCh <- set.rootSet
	}
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

func (s *RootManager) targetSet() *rootSet {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.target
}

func (s *RootManager) currentSet() *rootSet {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.current
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

const (
	currentRootListFname = "root-list.json"
	pendingRootListFname = "pending-root-list.json"
	targetRootListFname  = "target-root-list.json"
)

const refreshPeriod = time.Second
