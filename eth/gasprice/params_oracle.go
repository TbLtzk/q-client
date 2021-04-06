package gasprice

import (
	"math/big"
	"sync"

	"gitlab.com/q-dev/q-client/contracts"
	"gitlab.com/q-dev/q-client/core"
	"gitlab.com/q-dev/q-client/event"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/system-contracts/generated"
)

// GasPriceProvider.
type GasPriceProvider interface {
	GetGasPrice() *big.Int
}

type NoopGasPriceProvider struct{}

func (p *NoopGasPriceProvider) GetGasPrice() *big.Int {
	return nil
}

type blockChain interface {
	SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
}

// EPQFIParamsOracle returns gas price linked to QUSD.
type EPQFIParamsOracle struct {
	lock         sync.RWMutex
	price        *big.Int
	chainEventCh <-chan core.ChainHeadEvent

	registry *contracts.Registry
}

// NewEPQFIParamsOracle.
func NewEPQFIParamsOracle(defaultPrice *big.Int, reg *contracts.Registry, chain blockChain) *EPQFIParamsOracle {
	chainEvent := make(chan core.ChainHeadEvent)
	chain.SubscribeChainHeadEvent(chainEvent)

	oracle := &EPQFIParamsOracle{
		price:        defaultPrice,
		chainEventCh: chainEvent,
		registry:     reg,
	}
	go oracle.runLoop()

	return oracle
}

// GetGasPrice linked to QUSD converted to native currency.
// Returns default value if any of required system contracts/parameters
// is not available.
func (p *EPQFIParamsOracle) GetGasPrice() *big.Int {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return new(big.Int).Set(p.price)
}

func (p *EPQFIParamsOracle) runLoop() {
	for _ = range p.chainEventCh {
		p.refreshGasPrice()
	}
}

func (p *EPQFIParamsOracle) refreshGasPrice() {
	p.lock.Lock()
	defer p.lock.Unlock()

	epqfi := p.registry.EpqfiParameters()
	if epqfi == nil {
		return
	}

	oracleAddr, err := epqfi.GetAddr(nil, "governed.EPQFI.Q_QUSD_source")
	if err != nil {
		log.Warn("failed to get governed.EPQFI.Q_QUSD_source address", "err", err)
		return
	}

	oracle, _ := generated.NewFxPriceFeed(oracleAddr, p.registry.Backend)
	exchangeRate, err := oracle.ExchangeRate(nil)
	if err != nil {
		log.Warn("failed to get exchange rate", "err", err)
		return
	}

	if exchangeRate.Int64() == 0 {
		log.Warn("Q-QUSD oracle is not initialized", "addr", oracleAddr.Hex())
		return
	}

	priceQUSD, err := epqfi.GetUint(nil, "governed.EPQFI.txFee")
	if err != nil {
		log.Warn("failed to get governed.EPQFI.txFee", "err", err)
		return
	}

	decimals := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)
	newPrice := big.NewInt(0).Div(big.NewInt(0).Mul(priceQUSD, decimals), exchangeRate)

	log.Info("refreshed gas price", "old", p.price.String(), "new", newPrice.String())
	p.price = newPrice
}
