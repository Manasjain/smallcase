package dto


//AddTradeResponse model
//swagger:response AddTradeResponse
type AddTradeResponse struct {
	// A Unique Id for each Trade
	TradeId string `json:"tradeId"`
}

//UpdateTradeResponse model
//swagger:response UpdateTradeResponse
type UpdateTradeResponse struct{}

//RemoveTradeResponse model
//swagger:response RemoveTradeResponse
type RemoveTradeResponse struct{}

// Trade Info
//swagger:response TradeInfo
type TradeInfo struct {
	TradeId      string  `json:"trade_id"`
	TradeType    string  `json:"trade_type"`
	TradingUnits int64   `json:"trading_units"`
	UnitPrice    float64 `json:"unit_price"`
}

//FetchTradeResponse model
//swagger:response FetchTradeResponse
type FetchTradeResponse struct {
	// in: body
	// required: true
	ResponseBody struct{
		Trades map[string][]*TradeInfo `json:"trades"`
	}`json:"data"`
}

// Portfolio
//swagger:response Portfolio
type Portfolio struct {
	TotalUnits      int64   `json:"total_units"`
	AverageBuyPrice float64 `json:"average_buy_price"`
}

//GetPortfolioResponse model
//swagger:response GetPortfolioResponse
type GetPortfolioResponse struct {
	// in: body
	// required: true
	ResponseBody struct {
		Portfolio map[string]*Portfolio `json:"portfolio"`
	}

}

//GetPortfolioResponse model
//swagger:response GetReturnsResponse
type GetReturnsResponse struct {
	// Total Returns from all securities
	Returns float64 `json:"returns"`
}
