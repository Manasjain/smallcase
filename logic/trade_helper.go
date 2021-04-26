package logic

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"smallcase/db/dao"
	"smallcase/db/models"
	"smallcase/dto"
	"smallcase/enums"
	errors "smallcase/error"
	"smallcase/utils"
)

func (t *TradeAPI) CheckIfValidTrade(tradeRequest *dto.AddTradeRequest, tx *gorm.DB) (bool, error) {

	if *tradeRequest.RequestBody.TradeType == enums.SELL {
		conditions := &models.Trades{
			Ticker:    *tradeRequest.RequestBody.Ticker,
			DeletedAt: sql.NullTime{},
		}
		data, err := dao.DbClientTraderDao.FindAll(conditions, tx)
		if err != nil && err.Error() != fmt.Sprintf("%d:%s", errors.DatabaseRecordNotFound, errors.DatabaseServiceFailure.String()) {
			return false, err
		}
		//checking if total sum of trading units is greater that request tradingUnits
		var broughtTradingUnits int64 = 0
		for _, trade := range data {
			if trade.TradeType == enums.BUY {
				broughtTradingUnits += trade.TradingUnits
			} else {
				broughtTradingUnits -= trade.TradingUnits
			}
		}
		if broughtTradingUnits-*tradeRequest.RequestBody.TradingUnits < 0 {
			return false, nil
		}
	}
	return true, nil
}

func (t *TradeAPI) ValidTradeUpdate(tradeData *models.Trades, updateTradeData *models.Trades, tx *gorm.DB) (bool, error) {
	if tradeData.Ticker == updateTradeData.Ticker {
		// Case 1: if tickers are same
		conditions := &models.Trades{
			Ticker:    updateTradeData.Ticker,
			DeletedAt: sql.NullTime{},
		}
		tradesGroupByTicker, err := dao.DbClientTraderDao.FindAll(conditions, tx)
		if err != nil && err.Error() != fmt.Sprintf("%d:%s", errors.DatabaseRecordNotFound, errors.DatabaseRecordNotFound.String()) {
			return false, err
		}
		if !checkTotalTradingUnits(tradesGroupByTicker, updateTradeData) {
			return false, nil
		}

	} else {
		// Case:2 if tickers are different
		conditions := &models.Trades{
			Ticker:    tradeData.Ticker,
			DeletedAt: sql.NullTime{},
		}
		tradesGroupByTicker1, err := dao.DbClientTraderDao.FindAll(conditions, tx)
		if err != nil && err.Error() != fmt.Sprintf("%d:%s", errors.DatabaseRecordNotFound, errors.DatabaseRecordNotFound.String()) {
			return false, err
		}
		if !checkTotalTradingUnits(tradesGroupByTicker1, tradeData) {
			return false, nil
		}

		conditions.Ticker = updateTradeData.Ticker
		tradesGroupByTicker2, err := dao.DbClientTraderDao.FindAll(conditions, tx)
		if err != nil && err.Error() != fmt.Sprintf("%d:%s", errors.DatabaseRecordNotFound, errors.DatabaseRecordNotFound.String()) {
			return false, err
		}
		if !checkTotalTradingUnits(tradesGroupByTicker2, updateTradeData) {
			return false, nil
		}
	}

	return true, nil
}

func (t *TradeAPI) GetUpdateTradeData(updateTradeRequest *dto.UpdateTradeRequest, originalTradeData *models.Trades) *models.Trades {
	updateTradeData := &models.Trades{}
	utils.DeepClone(originalTradeData, updateTradeData)
	updateTradeData.UnitPrice = *updateTradeRequest.RequestBody.UnitPrice
	if updateTradeRequest.RequestBody.Ticker != nil {
		updateTradeData.Ticker = *updateTradeRequest.RequestBody.Ticker
	}
	if updateTradeRequest.RequestBody.TradeType != nil {
		updateTradeData.TradeType = *updateTradeRequest.RequestBody.TradeType
	}
	if updateTradeRequest.RequestBody.TradingUnits != nil {
		updateTradeData.TradingUnits = *updateTradeRequest.RequestBody.TradingUnits
	}
	return updateTradeData
}

func (t *TradeAPI) GetPortfolioFromTrades(trades []*models.Trades) map[string]*dto.Portfolio {
	totalBuyUnits := make(map[string]int64)
	totalSellUnits := make(map[string]int64)
	totalBuyPrice := make(map[string]float64)

	for _, trade := range trades {
		if trade.TradeType == enums.BUY {
			totalBuyPrice[trade.Ticker] += trade.UnitPrice * float64(trade.TradingUnits)
			totalBuyUnits[trade.Ticker] += trade.TradingUnits
		} else {
			totalSellUnits[trade.Ticker] += trade.TradingUnits
		}
	}
	portfolio := make(map[string]*dto.Portfolio)
	for k, v := range totalBuyPrice {
		p := &dto.Portfolio{
			TotalUnits:      totalBuyUnits[k] - totalSellUnits[k],
			AverageBuyPrice: v / float64(totalBuyUnits[k]),
		}
		portfolio[k] = p
	}
	return portfolio
}

func checkTotalTradingUnits(tradeDataGroupByTicker []*models.Trades, updateTrade *models.Trades) bool {

	//checking if total sum of trading units is positive
	var totalTradingUnits int64 = 0
	for _, trade := range tradeDataGroupByTicker {
		if trade.TradeId != updateTrade.TradeId && trade.TradeType == enums.BUY {
			totalTradingUnits += trade.TradingUnits
		} else if trade.TradeId != updateTrade.TradeId && trade.TradeType == enums.SELL {
			totalTradingUnits -= trade.TradingUnits
		}
	}
	if updateTrade.TradeType == enums.BUY {
		if totalTradingUnits+updateTrade.TradingUnits < 0 {
			return false
		}
	} else {
		if totalTradingUnits-updateTrade.TradingUnits < 0 {
			return false
		}
	}
	return true
}
