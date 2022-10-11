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
	addr := r.ValidatorsAddress()
	if addr == nil {
		return nil
	}

	val, err := generated.NewValidators(*addr, r.Backend)
	if err != nil {
		panic(errors.Wrap(err, "failed to init validators contract"))
	}

	return val
}

// Roots returns Roots contract backend if available.
func (r *Registry) Roots() *generated.Roots {
	addr := r.RootNodesAddress()
	if addr == nil {
		return nil
	}

	val, err := generated.NewRoots(*addr, r.Backend)
	if err != nil {
		panic(errors.Wrap(err, "failed to init roots contract"))
	}

	return val
}

func (r *Registry) ValidatorsAddress() *common.Address {
	addr := r.getAddr("governance.validators")
	if (addr == common.Address{}) {
		log.Debug("governance.validators contract is not deployed")
		return nil
	}

	return &addr
}

func (r *Registry) RootNodesAddress() *common.Address {
	addr := r.getAddr("governance.rootNodes")
	if (addr == common.Address{}) {
		log.Debug("governance.rootNodes contract is not deployed")
		return nil
	}

	return &addr
}

// RewardReceiver address.
func (r *Registry) RewardReceiver() common.Address {
	if r.registry() == nil {
		return r.defaultRewardReceiver
	}

	addr := r.getAddr("tokeneconomics.defaultAllocationProxy")
	if (addr == common.Address{}) {
		log.Debug("tokeneconomics.defaultAllocationProxy is not deployed")
		return r.defaultRewardReceiver
	}

	return addr
}

// ActiveValidatorsNumber.
func (r *Registry) ActiveValidatorsNumber() *int64 {
	addr := r.getAddr("governance.constitution.parameters")
	if (addr == common.Address{}) {
		log.Debug("governance.constitution.parameters is not deployed")
		return nil
	}

	params, err := generated.NewConstitution(addr, r.Backend)
	if err != nil {
		panic("failed to init constitution contract instance")
	}

	num, err := params.GetUint(nil, "constitution.maxNValidators")
	if err != nil {
		log.Debug("failed to get constitution.maxNValidators", "err", err)
		return nil
	}

	x := num.Int64()
	return &x
}

// EpqfiParameters returns EpqfiParameters contract backend if available.
func (r *Registry) EpqfiParameters() *generated.EPQFIParameters {
	addr := r.getAddr("governance.experts.EPQFI.parameters")
	if (addr == common.Address{}) {
		log.Debug("governance.experts.EPQFI.parameters is not deployed")
		return nil
	}

	// err is never returned here
	epqfiParams, _ := generated.NewEPQFIParameters(addr, r.Backend)
	return epqfiParams
}

// AccountAliases returns AccountAliases contract backend if available.
func (r *Registry) AccountAliases() *generated.AccountAliases {
	addr := r.AccountAliasesAddress()
	if addr == nil {
		return nil
	}

	val, err := generated.NewAccountAliases(*addr, r.Backend)
	if err != nil {
		panic(errors.Wrap(err, "failed to init accountAlies contract"))
	}

	return val
}

func (r *Registry) AccountAliasesAddress() *common.Address {
	addr := r.getAddr("governance.accountAliases")
	if (addr == common.Address{}) {
		log.Debug("governance.accountAliases contract is not deployed")
		return nil
	}

	return &addr
}

//Voting contracts

// ConstitutionVoting returns ConstitutionVoting contract backend if available.
func (r *Registry) ConstitutionVoting() *generated.ConstitutionVoting {
	addr := r.getAddr("governance.constitution.parametersVoting")
	if (addr == common.Address{}) {
		log.Debug("governance.constitution.parametersVoting is not deployed")
		return nil
	}

	// err is never returned here
	constitutionVoting, _ := generated.NewConstitutionVoting(addr, r.Backend)
	return constitutionVoting
}

// GeneralUpdateVoting returns GeneralUpdateVoting contract backend if available.
func (r *Registry) GeneralUpdateVoting() *generated.GeneralUpdateVoting {
	addr := r.getAddr("governance.generalUpdateVoting")
	if (addr == common.Address{}) {
		log.Debug("governance.generalUpdateVoting is not deployed")
		return nil
	}

	// err is never returned here
	generalUpdateVoting, _ := generated.NewGeneralUpdateVoting(addr, r.Backend)
	return generalUpdateVoting
}

// EmergencyUpdateVoting returns EmergencyUpdateVoting contract backend if available.
func (r *Registry) EmergencyUpdateVoting() *generated.EmergencyUpdateVoting {
	addr := r.getAddr("governance.emergencyUpdateVoting")
	if (addr == common.Address{}) {
		log.Debug("governance.generalUpdateVoting is not deployed")
		return nil
	}

	// err is never returned here
	emergencyUpdateVoting, _ := generated.NewEmergencyUpdateVoting(addr, r.Backend)
	return emergencyUpdateVoting
}

// RootsVoting returns RootsVoting contract backend if available.
func (r *Registry) RootsVoting() *generated.RootsVoting {
	addr := r.getAddr("governance.rootNodes.membershipVoting")
	if (addr == common.Address{}) {
		log.Debug("governance.rootNodes.membershipVoting is not deployed")
		return nil
	}

	// err is never returned here
	rootsVoting, _ := generated.NewRootsVoting(addr, r.Backend)
	return rootsVoting
}

// RootNodesSlashingVoting returns RootNodesSlashingVoting contract backend if available.
func (r *Registry) RootNodesSlashingVoting() *generated.RootNodesSlashingVoting {
	addr := r.getAddr("governance.rootNodes.slashingVoting")
	if (addr == common.Address{}) {
		log.Debug("governance.rootNodes.slashingVoting is not deployed")
		return nil
	}

	// err is never returned here
	rootNodesSlashingVoting, _ := generated.NewRootNodesSlashingVoting(addr, r.Backend)
	return rootNodesSlashingVoting
}

// ValidatorsSlashingVoting returns ValidatorsSlashingVoting contract backend if available.
func (r *Registry) ValidatorsSlashingVoting() *generated.ValidatorsSlashingVoting {
	addr := r.getAddr("governance.validators.slashingVoting")
	if (addr == common.Address{}) {
		log.Debug("governance.validators.slashingVoting is not deployed")
		return nil
	}

	// err is never returned here
	validatorsSlashingVoting, _ := generated.NewValidatorsSlashingVoting(addr, r.Backend)
	return validatorsSlashingVoting
}

// EpqfiMembershipVoting returns EpqfiMembershipVoting contract backend if available.
func (r *Registry) EpqfiMembershipVoting() *generated.EPQFIMembershipVoting {
	addr := r.getAddr("governance.experts.EPQFI.membershipVoting")
	if (addr == common.Address{}) {
		log.Debug("governance.experts.EPQFI.membershipVoting is not deployed")
		return nil
	}

	// err is never returned here
	epqfiMembershipVoting, _ := generated.NewEPQFIMembershipVoting(addr, r.Backend)
	return epqfiMembershipVoting
}

// EpqfiParametersVoting returns EpqfiParametersVoting contract backend if available.
func (r *Registry) EpqfiParametersVoting() *generated.EPQFIParametersVoting {
	addr := r.getAddr("governance.experts.EPQFI.parametersVoting")
	if (addr == common.Address{}) {
		log.Debug("governance.experts.EPQFI.parametersVoting is not deployed")
		return nil
	}

	// err is never returned here
	epqfiParametersVoting, _ := generated.NewEPQFIParametersVoting(addr, r.Backend)
	return epqfiParametersVoting
}

// EpdrMembershipVoting returns EpdrMembershipVoting contract backend if available.
func (r *Registry) EpdrMembershipVoting() *generated.EPDRMembershipVoting {
	addr := r.getAddr("governance.experts.EPDR.membershipVoting")
	if (addr == common.Address{}) {
		log.Debug("governance.experts.EPDR.membershipVoting is not deployed")
		return nil
	}

	// err is never returned here
	epdrMembershipVoting, _ := generated.NewEPDRMembershipVoting(addr, r.Backend)
	return epdrMembershipVoting
}

// EpdrParametersVoting returns EpdrParametersVoting contract backend if available.
func (r *Registry) EpdrParametersVoting() *generated.EPDRParametersVoting {
	addr := r.getAddr("governance.experts.EPDR.parametersVoting")
	if (addr == common.Address{}) {
		log.Debug("governance.experts.EPDR.parametersVoting is not deployed")
		return nil
	}

	// err is never returned here
	epdrParametersVoting, _ := generated.NewEPDRParametersVoting(addr, r.Backend)
	return epdrParametersVoting
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

func (r *Registry) getAddr(key string) common.Address {
	reg := r.registry()
	if reg == nil {
		return common.Address{}
	}

	addr, err := reg.GetAddress(nil, key)
	if err != nil {
		log.Error("failed to get address from registry", "err", err, "key", key)
	}

	return addr
}
