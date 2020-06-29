package proxy

import (
	"math/big"

	"gitlab.com/q-dev/go-ethereum/accounts/abi"
	"gitlab.com/q-dev/go-ethereum/accounts/abi/bind"
	"gitlab.com/q-dev/go-ethereum/common"
	"gitlab.com/q-dev/go-ethereum/contracts/proxy/contract"
	"gitlab.com/q-dev/go-ethereum/core/types"
)

type ValidatorsProxy struct {
	abi      abi.ABI
	address  common.Address
	contract *contract.ValidatorsProxy
}

func NewValidatorsProxy(contractAddr common.Address, backend bind.ContractBackend) (*ValidatorsProxy, error) {
	c, err := contract.NewValidatorsProxy(contractAddr, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorsProxy{address: contractAddr, contract: c}, nil
}

func (proxy *ValidatorsProxy) GetValidatorsList() ([]common.Address, error) {
	return proxy.GetValidatorsList()
}

func (proxy *ValidatorsProxy) AddValidator(opts *bind.TransactOpts, validatorAddress common.Address) (*types.Transaction, error) {
	return proxy.contract.AddValidator(opts, validatorAddress)
}

func (proxy *ValidatorsProxy) DelegateStakeToValidator(address common.Address, amount *big.Int) error {
	//TODO:implement
	return nil
}
