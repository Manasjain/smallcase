package dto


// swagger:parameters AddTradeHandler
type AddTradeRequest struct {
	// in: body
	RequestBody AddTradeRequestBody `json:"data"`
}

type AddTradeRequestBody struct {
	// required: true
	// example: BUY
	TradeType    *string  `json:"trade_type"`
	// required: true
	// example: BSE
	Ticker       *string  `json:"ticker"`
	// required: true
	// example: 10
	TradingUnits *int64   `json:"trading_units"`
	// required: true
	// example: 90
	UnitPrice    *float64 `json:"unit_price"`
}

func (r *AddTradeRequest) Error() error {
	return nil
}

// swagger:parameters UpdateTradeHandler
type UpdateTradeRequest struct {
	// in:body
	RequestBody UpdateTradeRequestBody `json:"data"`
}

type UpdateTradeRequestBody struct{
	// required: true
	// example: fda060c4-186d-4dd0-8e4a-92592a19797f
	TradeId      *string  `json:"tradeId"`
	// required: true
	// example: 90
	UnitPrice    *float64 `json:"unit_price"`
	// example: BSE
	Ticker       *string  `json:"ticker,omitempty"`
	// example: BUY
	TradeType    *string  `json:"trade_type,omitempty"`
	// example: 10
	TradingUnits *int64   `json:"trading_units,omitempty"`
}


