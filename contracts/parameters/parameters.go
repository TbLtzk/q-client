package parameters

import (
	"math/big"

	"gitlab.com/q-dev/go-ethereum/accounts/abi/bind"
	"gitlab.com/q-dev/go-ethereum/common"
	"gitlab.com/q-dev/go-ethereum/contracts/parameters/contract"
)

type Parameters struct {
	address  common.Address
	contract *contract.Parameters
}

func NewParameters(contractAddr common.Address, backend bind.ContractBackend) (*Parameters, error) {
	c, err := contract.NewParameters(contractAddr, backend)
	if err != nil {
		return nil, err
	}
	return &Parameters{address: contractAddr, contract: c}, nil
}

func (parameters *Parameters) GetCoinbaseAddress() (common.Address, error) {
	return parameters.GetCoinbaseAddress()
}

func (parameters *Parameters) DelegateStakeToValidator(address common.Address, amount *big.Int) error {
	//TODO:implement
	return nil
}
