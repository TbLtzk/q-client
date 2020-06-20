package validators

import (
	"math/big"

	"gitlab.com/q-dev/go-ethereum/accounts/abi/bind"
	"gitlab.com/q-dev/go-ethereum/common"
	"gitlab.com/q-dev/go-ethereum/contracts/validators/contract"
)

// Validators is a Go wrapper around an on-chain validator contract.
type Validators struct {
	address  common.Address
	contract *contract.Validators
}

// NewValidators binds contract and returns a registrar instance.
func NewValidators(contractAddr common.Address, backend bind.ContractBackend) (*Validators, error) {
	c, err := contract.NewValidators(contractAddr, backend)
	if err != nil {
		return nil, err
	}
	return &Validators{address: contractAddr, contract: c}, nil
}

func (validators *Validators) GetValidatorsList() ([]common.Address, error) {
	return validators.contract.GetValidatorsList(&bind.CallOpts{})
}

func (validators *Validators) DelegateStakeToValidator(address common.Address, amount *big.Int) error {
	//TODO:implement
	return nil
}
