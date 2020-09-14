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
	"gitlab.com/q-dev/go-ethereum/log"
	"gitlab.com/q-dev/go-ethereum/params"
)

var (
	errInvalidSignature = errors.New("list contains invalid signature")
	errIncomplete       = errors.New("not enough signatures")
	errHashMismatch     = errors.New("hash mismatch")
)

// RootManager governs root nodes set.
type RootManager struct {
	IsRoot bool

	keystore Keystore

	currentFpath string
	desiredFpath string

	lock         sync.Mutex
	set          *rootSet
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

	set, err := newRootSet(list)
	if err != nil {
		return nil, errors.Wrap(err, "invalid current root node list")
	}

	return &RootManager{
		keystore:     ks,
		currentFpath: currentListPath,
		desiredFpath: filepath.Join(datadir, desiredRootListFname),
		exclusionSet: make(map[common.Address]struct{}),
		set:          set,
	}, nil
}

func (s *RootManager) CurrentList() common.RootList {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.set.makeList()
}

// ExclusionSet returns set of excluded validators addresses.
func (s *RootManager) ExclusionSet() map[common.Address]struct{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.exclusionSet
}

func (s *RootManager) run() error {
	var roots []string
	for _, addr := range s.set.rootAddresses {
		if !s.keystore.IsUnlocked(addr) {
			continue
		}

		s.IsRoot = true
		roots = append(roots, addr.Hex())

		if _, ok := s.set.signers[addr]; !ok {
			signature, err := s.keystore.SignHash(accounts.Account{Address: addr}, s.set.hash.Bytes())
			if err != nil {
				log.Error("failed to sign root list", err)
			}

			log.Info("signed root node list", "addr", addr)
			s.set.signers[addr] = signature
		}
	}

	go s.refreshCurrentRootList()

	if s.IsRoot {
		log.Info("Node belongs to the current root node set", "addr", roots)
	}

	return nil
}

func (s *RootManager) refreshCurrentRootList() {
	for {
		list := s.CurrentList()
		f, err := os.Create(s.currentFpath)
		if err != nil {
			log.Warn("failed to create root node list file", "err", err)
		}

		if err := json.NewEncoder(f).Encode(list); err != nil {
			log.Warn("failed to refresh root node list file")
		}

		_ = f.Close()

		time.Sleep(time.Second)
	}
}

func (s *RootManager) upgrade(set *rootSet) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.set.timestamp >= set.timestamp {
		// skip a concurrent call with an obsolete list
		return
	}

	s.set = set
}

func (s *RootManager) currentSet() *rootSet {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.set
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

const (
	currentRootListFname = "root-list.json"
	desiredRootListFname = "desired-root-list.json"
)
