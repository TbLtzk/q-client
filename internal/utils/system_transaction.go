package utils

import (
	"math/big"
	"strings"

	"gitlab.com/q-dev/q-client/accounts"
	"gitlab.com/q-dev/q-client/core/state"

	"gitlab.com/q-dev/q-client/consensus"
	"gitlab.com/q-dev/q-client/params"

	"gitlab.com/q-dev/q-client/accounts/abi"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/system-contracts/generated"
)

type ValidatorProvider interface {
	Validators() *common.Address
}

type SystemTxPreparer struct {
	config       *params.ChainConfig
	engine       consensus.Engine
	currentState *state.StateDB
	header       *types.Header
	signer       types.Signer
}

func New(c *params.ChainConfig, e consensus.Engine, s *state.StateDB, h *types.Header, signer types.Signer) *SystemTxPreparer {
	return &SystemTxPreparer{
		config:       c,
		engine:       e,
		currentState: s,
		header:       h,
		signer:       signer,
	}
}

func (w *SystemTxPreparer) PrepareSystemTx(accountManager *accounts.Manager) map[common.Address]types.Transactions {
	result := make(map[common.Address]types.Transactions)

	if (w.config.Clique != nil) && ((w.header.Number.Uint64()+1)%w.config.Clique.Epoch == 0) {
		signer := w.engine.Signer()
		if (signer != common.Address{}) {
			log.Info("attempting to create system tx", "signer", signer)

			cliq, ok := w.engine.(ValidatorProvider)
			if ok {
				addr := cliq.Validators()
				if addr != nil {
					tx, err := w.prepareTx(*addr, signer)
					if err != nil {
						log.Warn("failed to prepare tx", "err", err)
						return result
					}

					if accountManager == nil {
						log.Warn("account manager is not available")
						return result
					}

					account := accounts.Account{Address: signer}
					wallet, err := accountManager.Find(account)
					if err != nil {
						log.Warn("failed to find account", "err", err)
						return result
					}

					tx, err = wallet.SignTx(account, tx, w.config.ChainID)
					if err != nil {
						log.Warn("failed to sign tx", "err", err)
						return result
					}

					types.Sender(w.signer, tx)
					result[signer] = types.Transactions{tx}
					log.Debug("system tx is here")
					log.Info("system tx", "tx", *tx)
				} else {
					log.Warn("validators contract is not available")
				}
			}

			if len(result) == 0 {
				log.Warn("there is no system tx")
			}
		}
	}

	return result
}

func (w *SystemTxPreparer) prepareTx(contractAddress, sender common.Address) (*types.Transaction, error) {
	a, err := abi.JSON(strings.NewReader(generated.ValidatorsABI))
	if err != nil {
		return nil, err
	}

	input, err := a.Pack("makeSnapshot")
	if err != nil {
		return nil, err
	}

	nonce := w.currentState.GetNonce(sender)

	return types.NewTransaction(nonce, contractAddress, big.NewInt(0), 1477210, big.NewInt(50000000000), input), nil
}
