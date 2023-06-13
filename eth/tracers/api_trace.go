// See: https://openethereum.github.io/JSONRPC-trace-module for details
package tracers

import (
	"context"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/common/hexutil"
	"gitlab.com/q-dev/q-client/rpc"
)

type TraceFilterParams struct {
	FromBlock   hexutil.Uint64  `json:"fromBlock,omitempty"`
	ToBlock     hexutil.Uint64  `json:"toBlock,omitempty"`
	FromAddress *common.Address `json:"fromAddress,omitempty"`
	ToAddress   *common.Address `json:"toAddress,omitempty"`
	After       uint64          `json:"after,omitempty"`
	Count       uint64          `json:"count,omitempty"`
}

// Duplicates, but cycle import is not allowed
// CallParityFrame is the result of a callParityTracerParity run.
type CallParityFrame struct {
	Action              CallTraceParityAction  `json:"action"`
	BlockHash           *common.Hash           `json:"blockHash"`
	BlockNumber         uint64                 `json:"blockNumber"`
	Error               string                 `json:"error,omitempty"`
	Result              *CallTraceParityResult `json:"result,omitempty"`
	Subtraces           int                    `json:"subtraces"`
	TraceAddress        []int                  `json:"traceAddress"`
	TransactionHash     *common.Hash           `json:"transactionHash"`
	TransactionPosition *uint64                `json:"transactionPosition"`
	Type                string                 `json:"type"`
	Calls               []CallParityFrame      `json:"-"`
}

type CallTraceParityAction struct {
	Author         *common.Address `json:"author,omitempty"`
	RewardType     *string         `json:"rewardType,omitempty"`
	SelfDestructed *common.Address `json:"address,omitempty"`
	Balance        *hexutil.Big    `json:"balance,omitempty"`
	CallType       string          `json:"callType,omitempty"`
	CreationMethod string          `json:"creationMethod,omitempty"`
	From           *common.Address `json:"from,omitempty"`
	Gas            *hexutil.Uint64 `json:"gas,omitempty"`
	Init           *hexutil.Bytes  `json:"init,omitempty"`
	Input          *hexutil.Bytes  `json:"input,omitempty"`
	RefundAddress  *common.Address `json:"refundAddress,omitempty"`
	To             *common.Address `json:"to,omitempty"`
	Value          *hexutil.Big    `json:"value,omitempty"`
}

type CallTraceParityResult struct {
	Address *common.Address `json:"address,omitempty"`
	Code    *hexutil.Bytes  `json:"code,omitempty"`
	GasUsed *hexutil.Uint64 `json:"gasUsed,omitempty"`
	Output  *hexutil.Bytes  `json:"output,omitempty"`
}

type TraceAPI struct {
	debugAPI *API
}

// NewTraceAPI creates a new API definition for the full node-related
// private debug methods of the Ethereum service.
func NewTraceAPI(debugAPI *API) *TraceAPI {
	return &TraceAPI{debugAPI: debugAPI}
}

// setTraceConfigDefaultTracer sets the default tracer to "callTracerParity" if none set
func setTraceConfigDefaultTracer(config *TraceConfig) *TraceConfig {
	if config == nil {
		config = &TraceConfig{}
	}
	tm := "2m"
	if config.Tracer == nil {
		tracer := "callTracerParity"
		config.Tracer = &tracer
		config.Timeout = &tm
	}

	return config
}

func (api *TraceAPI) Filter(ctx context.Context, args TraceFilterParams, config *TraceConfig) ([]interface{}, error) {
	config = setTraceConfigDefaultTracer(config)

	// Fetch the block interval that we want to trace
	start := rpc.BlockNumber(args.FromBlock)
	end := rpc.BlockNumber(args.ToBlock)

	return api.debugAPI.TraceChainWithFilterApplied(ctx, start, end, config, args.FromAddress, args.ToAddress, args.After, args.Count)
}
