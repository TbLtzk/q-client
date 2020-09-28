package contracts

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/accounts/abi/bind"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/system-contracts/generated"
)

// Registry of system contracts.
// current implementation is rather proof of concept
type Registry struct {
	mu  sync.Mutex
	reg *generated.ContractRegistry

	addr                  common.Address
	defaultRewardReceiver common.Address

	back bind.ContractBackend
}

// NewRegistry.
func NewRegistry(addr, defaultRewardReceiver common.Address) *Registry {
	return &Registry{addr: addr, defaultRewardReceiver: defaultRewardReceiver}
}

// Init registry.
// todo: this is ugly as hell, gotta keep this due to current code structure(
func (r *Registry) Init(back bind.ContractBackend) error {
	r.back = back
	code, err := back.CodeAt(context.TODO(), r.addr, nil)
	if err != nil {
		return errors.Wrap(err, "failed to check if contract code is onchain")
	}

	if len(code) == 0 {
		go r.setupRegistry(r.addr)
		return nil
	}

	reg, err := generated.NewContractRegistry(r.addr, back)
	if err != nil {
		panic(err)
	}

	r.mu.Lock()
	r.reg = reg
	r.mu.Unlock()
	return nil
}

// Validators returns Validators contract backend if available.
func (r *Registry) Validators() *generated.Validators {
	reg := r.registry()
	if reg == nil {
		return nil
	}

	addr, err := reg.GetAddress(nil, "validators")
	if err != nil {
		log.Warn("failed to get validators address", "err", err)
		return nil
	}

	code, err := r.back.CodeAt(context.TODO(), addr, nil)
	if err != nil {
		log.Warn("failed to check if validators contract is deployed", "err", err)
		return nil
	}

	if len(code) == 0 {
		return nil
	}

	val, err := generated.NewValidators(addr, r.back)
	if err != nil {
		panic(errors.Wrap(err, "failed to init validators contract"))
	}

	return val
}

func (r *Registry) ValidatorsAddress() *common.Address {
	reg := r.registry()
	if reg == nil {
		return nil
	}

	addr, err := reg.GetAddress(nil, "validators")
	if err != nil {
		log.Warn("failed to get validators address", "err", err)
		return nil
	}

	code, err := r.back.CodeAt(context.TODO(), addr, nil)
	if err != nil {
		log.Warn("failed to check if validators contract is deployed", "err", err)
		return nil
	}

	if len(code) == 0 {
		log.Warn("there is no validators code")
		return nil
	}

	return &addr
}

// RewardReceiver address.
func (r *Registry) RewardReceiver() common.Address {
	// todo: fix this: if registry is used for this one, node consumes all ram and dies(
	return r.defaultRewardReceiver
}

func (r *Registry) setupRegistry(addr common.Address) {
	for {
		time.Sleep(time.Second)

		code, err := r.back.CodeAt(context.TODO(), addr, nil)
		if err != nil {
			log.Warn("failed to check if contract registry was deployed", "err", err)
			continue
		}

		if len(code) == 0 {
			log.Debug("contract registry hasn't been deployed yet")
			continue
		}

		r.mu.Lock()
		r.reg, err = generated.NewContractRegistry(addr, r.back)
		if err != nil {
			panic(errors.Wrap(err, "failed to init registry backend")) // normally, should never happen
		}
		r.mu.Unlock()
		log.Debug("successfully initialized registry")

		return
	}
}

func (r *Registry) registry() *generated.ContractRegistry {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.reg
}
