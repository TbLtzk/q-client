package mocks

import (
	"math/big"

	"gitlab.com/q-dev/q-client/accounts/abi/bind"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/contracts"
)

type mockAccountAliases struct {
	aliases        map[common.Address]common.Address
	reverseAliases map[common.Address]common.Address
}

var MockAccount *mockAccountAliases

func NewMockAccountAliases(aliases map[common.Address]common.Address) contracts.AccountAliasesI {
	MockAccount = new(mockAccountAliases)

	MockAccount.aliases = aliases
	MockAccount.reverseAliases = make(map[common.Address]common.Address)

	for account, alias := range aliases {
		MockAccount.reverseAliases[alias] = account
	}
	return MockAccount
}

func (m *mockAccountAliases) ResolveBatch(opts *bind.CallOpts, main []common.Address, role []*big.Int) ([]common.Address, error) {
	aliases := make([]common.Address, 0, len(main))
	for _, ma := range main {
		resolved := ma
		if alias, ok := m.aliases[ma]; ok {
			resolved = alias
		}
		aliases = append(aliases, resolved)
	}

	return aliases, nil
}

func (m *mockAccountAliases) ResolveBatchReverse(opts *bind.CallOpts, alias []common.Address, role []*big.Int) ([]common.Address, error) {
	mains := make([]common.Address, 0, len(alias))
	for _, a := range alias {
		resolved := a
		if main, ok := m.reverseAliases[a]; ok {
			resolved = main
		}
		mains = append(mains, resolved)
	}

	return mains, nil
}

func (m *mockAccountAliases) Resolve(opts *bind.CallOpts, main common.Address, role *big.Int) (common.Address, error) {
	return m.aliases[main], nil
}

func (m *mockAccountAliases) ResolveReverse(opts *bind.CallOpts, alias common.Address, role *big.Int) (common.Address, error) {
	return m.reverseAliases[alias], nil
}

type mockValidators struct {
	address    *common.Address
	validators []common.Address
	error
}

func NewMockValidators(address common.Address, validators []common.Address, err error) contracts.ValidatorsI {
	return &mockValidators{
		address:    &address,
		validators: validators,
		error:      err,
	}
}

func (m *mockValidators) ValidatorsAddress() *common.Address {
	return m.address
}

func (m *mockValidators) GetValidatorsList(opts *bind.CallOpts) ([]common.Address, error) {
	return m.validators, m.error
}
