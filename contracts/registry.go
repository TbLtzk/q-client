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
	Backend bind.ContractBackend

	mu  sync.Mutex
	reg *generated.ContractRegistry

	addr                  common.Address
	defaultRewardReceiver common.Address

	isTestMode bool
}

// NewRegistry.
func NewRegistry(addr, defaultRewardReceiver common.Address, back bind.ContractBackend) *Registry {
	return &Registry{addr: addr, defaultRewardReceiver: defaultRewardReceiver, Backend: back}
}

func NewTestModeRegistry() *Registry {
	return &Registry{
		// need some non-empty address here for miner
		defaultRewardReceiver: common.HexToAddress("0x92C35a964624D9cbF90c2A0525e116093FAF867E"),
		isTestMode:            true,
	}
}

// Validators returns Validators contract backend if available.
func (r *Registry) Validators() *generated.Validators {
	reg := r.registry()
	if reg == nil {
		return nil
	}

	addr, err := reg.MustGetAddress(nil, "governance.validators")
	if err != nil {
		log.Warn("failed to get validators address", "err", err)
		return nil
	}

	code, err := r.Backend.CodeAt(context.TODO(), addr, nil)
	if err != nil {
		log.Warn("failed to check if validators contract is deployed", "err", err)
		return nil
	}

	if len(code) == 0 {
		log.Warn("there is no validators code")
		return nil
	}

	val, err := generated.NewValidators(addr, r.Backend)
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

	addr, err := reg.MustGetAddress(nil, "governance.validators")
	if err != nil {
		log.Warn("failed to get validators address", "err", err)
		return nil
	}

	return &addr
}

// RewardReceiver address.
func (r *Registry) RewardReceiver() common.Address {
	if r.registry() == nil {
		return r.defaultRewardReceiver
	}

	addr, err := r.registry().MustGetAddress(nil, "tokeneconomics.defaultAllocationProxy")
	if err != nil {
		log.Warn("failed to get reward receiver address", "err", err)
		return r.defaultRewardReceiver
	}

	return addr
}

// ActiveValidatorsNumber.
func (r *Registry) ActiveValidatorsNumber() *int64 {
	if r.registry() == nil {
		return nil
	}

	addr, err := r.registry().MustGetAddress(nil, "governance.constitution.parameters")
	if err != nil {
		return nil
	}

	params, err := generated.NewConstitution(addr, r.Backend)
	if err != nil {
		log.Warn("failed to init constitution from address", "addr", addr.Hex(), "err", err)
		return nil
	}

	num, err := params.GetUint(nil, "constitution.maxNValidators")
	if err != nil {
		log.Warn("failed to get constitution.maxNValidators", "err", err)
		return nil
	}

	x := num.Int64()
	return &x
}

// EpqfiParameters returns EpqfiParameters contract backend if available.
func (r *Registry) EpqfiParameters() *generated.EPQFIParameters {
	reg := r.registry()
	if reg == nil {
		return nil
	}

	addr, err := reg.MustGetAddress(nil, "governance.experts.EPQFI.parameters")
	if err != nil {
		log.Warn("failed to get EPQFI_parameters address", "err", err)
		return nil
	}

	// err is never returned here
	epqfiParams, _ := generated.NewEPQFIParameters(addr, r.Backend)
	return epqfiParams
}

func (r *Registry) registry() *generated.ContractRegistry {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.reg != nil {
		return r.reg
	}

	// fallback behaviour in test mode
	if r.isTestMode || r.Backend == nil {
		return nil
	}

	// try to init an instance then
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	code, err := r.Backend.CodeAt(ctx, r.addr, nil)
	if err != nil {
		log.Warn("failed to check if contract registry was deployed", "err", err)
		return nil
	}

	if len(code) == 0 {
		log.Debug("contract registry hasn't been deployed yet")
		return nil
	}

	// can skip err, it is never returned here
	r.reg, _ = generated.NewContractRegistry(r.addr, r.Backend)
	return r.reg
}
