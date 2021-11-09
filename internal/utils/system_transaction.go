package utils

import (
	"strings"

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
	coinbase     common.Address
	header       *types.Header
}

func New(c *params.ChainConfig, e consensus.Engine, s *state.StateDB,
	coin common.Address, h *types.Header,
) *SystemTxPreparer {
	return &SystemTxPreparer{
		config:       c,
		engine:       e,
		currentState: s,
		coinbase:     coin,
		header:       h,
	}
}

func (w *SystemTxPreparer) PrepareSystemTx() map[common.Address]types.Transactions {
	result := make(map[common.Address]types.Transactions)

	if (w.config.Clique != nil) && ((w.header.Number.Uint64()+1)%w.config.Clique.Epoch == 0) {
		log.Debug("attempting to create system tx")

		cliq, ok := w.engine.(ValidatorProvider)
		if ok {
			addr := cliq.Validators()
			if addr != nil {
				tx, err := w.prepareTx(*addr, w.coinbase)
				if err != nil {
					log.Warn("failed to prepare tx", "err", err)
				}

				SetSenderFromServer(tx, w.coinbase, w.header.Hash())
				result[w.coinbase] = types.Transactions{tx}
				log.Debug("system tx is here")
			} else {
				log.Warn("validators contract is not available")
			}
		}

		if len(result) == 0 {
			log.Warn("there is no system tx")
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

	// TODO calculate gas properly
	return types.NewTransaction(nonce, contractAddress, nil, 1477210, nil, input), nil
}
