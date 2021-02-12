package core

import (
	"math/big"

	"gitlab.com/q-dev/q-client/contracts"
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

// QUSDGasPriceProvider returns gas price linked to QUSD.
type QUSDGasPriceProvider struct {
	registry *contracts.Registry
}

// NewQUSDGasPriceProvider.
func NewQUSDGasPriceProvider(reg *contracts.Registry) *QUSDGasPriceProvider {
	return &QUSDGasPriceProvider{registry: reg}
}

// GetGasPrice linked to QUSD converted to native currency.
// Returns nil if any of required system contracts is not set up yet.
func (p *QUSDGasPriceProvider) GetGasPrice() *big.Int {
	epqfi := p.registry.EpqfiParameters()
	if epqfi == nil {
		return nil
	}

	oracleAddr, err := epqfi.GetAddr(nil, "governed.EPQFI.Q_QUSD_source")
	if err != nil {
		log.Warn("failed to get governed.EPQFI.Q_QUSD_source address", "err", err)
		return nil
	}

	oracle, _ := generated.NewFxPriceFeed(oracleAddr, p.registry.Backend)
	exchangeRate, err := oracle.ExchangeRate(nil)
	if err != nil {
		log.Warn("failed to get exchange rate", "err", err)
		return nil
	}

	if exchangeRate.Int64() == 0 {
		log.Warn("Q-QUSD oracle is not initialized", "addr", oracleAddr.Hex())
		return nil
	}

	priceQUSD, err := epqfi.GetUint(nil, "governed.EPQFI.txFee")
	if err != nil {
		log.Warn("failed to get governed.EPQFI.txFee", "err", err)
		return nil
	}

	decimals := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)
	return big.NewInt(0).Div(big.NewInt(0).Mul(priceQUSD, decimals), exchangeRate)
}
