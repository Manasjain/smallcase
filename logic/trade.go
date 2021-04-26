package logic

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"log"
	"smallcase/db/dao"
	"smallcase/db/models"
	"smallcase/dto"
	errors "smallcase/error"
	"smallcase/utils"
	"time"
)

func (t *TradeAPI) AddTrade(ctx context.Context, trade *dto.AddTradeRequest, tx *gorm.DB) (*dto.AddTradeResponse, error) {
	var response *dto.AddTradeResponse
	tradeId := uuid.New().String()

	tradeData := &models.Trades{
		TradeId:      tradeId,
		Ticker:       *trade.RequestBody.Ticker,
		TradeType:    *trade.RequestBody.TradeType,
		TradingUnits: *trade.RequestBody.TradingUnits,
		UnitPrice:    *trade.RequestBody.UnitPrice,
	}

	// Validate the trade
	validTrade, err := t.CheckIfValidTrade(trade, tx)
	if err != nil {
		log.Println("Failed to validate the trade")
		return nil, err
	}
	if !validTrade {
		log.Println("Trade is Invalid. Sell TradeUnits are greater than Buy")
		return nil, errors.NewError(errors.BadRequest, errors.BadRequest.String())
	}

	// Save the trade to DB
	if err := dao.DbClientTraderDao.Save(tradeData, tx); err != nil {
		log.Println("Failed to save the trade, InternalSever Error")
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	response = &dto.AddTradeResponse{TradeId: tradeId}
	return response, nil
}

func (t *TradeAPI) UpdateTrade(ctx context.Context, updateTradeRequest *dto.UpdateTradeRequest, tx *gorm.DB) (*dto.UpdateTradeResponse, error) {
	var response *dto.UpdateTradeResponse
	conditions := &models.Trades{
		TradeId:   *updateTradeRequest.RequestBody.TradeId,
		DeletedAt: sql.NullTime{},
	}

	// Get the original trade
	originalTradeData, err := dao.DbClientTraderDao.Find(conditions, tx)
	if err != nil {
		log.Println("trade not found, DB Error: ", utils.Marshal(err))
		return nil, err
	}

	updateTradeData := t.GetUpdateTradeData(updateTradeRequest, originalTradeData)

	// Validate if update is possible
	isUpdateValid, err := t.ValidTradeUpdate(originalTradeData, updateTradeData, tx)
	if err != nil {
		return nil, err
	}
	if !isUpdateValid {
		log.Println("Trade is Invalid. Sell TradeUnits are greater than Buy")
		return nil, errors.NewError(errors.BadRequest, errors.BadRequest.String())
	}

	// Update the Trade
	if err := dao.DbClientTraderDao.Update(conditions, updateTradeData, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return response, nil
}

func (t *TradeAPI) RemoveTrade(ctx context.Context, tradeId *string, tx *gorm.DB) (*dto.RemoveTradeResponse, error) {
	var response *dto.RemoveTradeResponse
	if tradeId == nil || *tradeId == "" {
		log.Println("tradeId is empty")
		return nil, errors.NewError(errors.BadRequest, errors.BadRequest.String())
	}

	// Find trade in database
	conditions := &models.Trades{
		TradeId:   *tradeId,
		DeletedAt: sql.NullTime{},
	}

	originalTradeData, err := dao.DbClientTraderDao.Find(conditions, tx)
	if err != nil {
		log.Println("trade not found, DB Error: ", utils.Marshal(err))
		return nil, err
	}
	log.Println(utils.Marshal(originalTradeData))
	// Validate if trade can be deleted
	trade := &dto.AddTradeRequest{}
	trade.RequestBody.TradeType = &originalTradeData.TradeType
	trade.RequestBody.TradingUnits = &originalTradeData.TradingUnits
	trade.RequestBody.Ticker = &originalTradeData.Ticker

	validTrade, err := t.CheckIfValidTrade(trade, tx)
	if err != nil {
		log.Println("Failed to validate the trade")
		return nil, err
	}
	if !validTrade {
		log.Println("Trade cannot be deleted. Invalid tradeId to be deleted")
		return nil, errors.NewError(errors.BadRequest, errors.BadRequest.String())
	}

	// Delete the trade
	conditions = &models.Trades{
		TradeId: *tradeId,
	}
	removeTrade := &models.Trades{
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	if err := dao.DbClientTraderDao.Update(conditions, removeTrade, tx); err != nil {
		log.Println("Failed to delete the trade, DB Err: ", utils.Marshal(err))
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return response, nil
}

func (t *TradeAPI) FetchTrades(ctx context.Context, tx *gorm.DB) (*dto.FetchTradeResponse, error) {
	var response = &dto.FetchTradeResponse{}
	// Fetching all trades from database
	conditions := &models.Trades{
		DeletedAt: sql.NullTime{},
	}
	trades, err := dao.DbClientTraderDao.FindAll(conditions, tx)
	if err != nil {
		return nil, err
	}

	TradesMap := make(map[string][]*dto.TradeInfo)
	for _, trade := range trades {
		tradeDetail := &dto.TradeInfo{
			TradeId:      trade.TradeId,
			TradeType:    trade.TradeType,
			TradingUnits: trade.TradingUnits,
			UnitPrice:    trade.UnitPrice,
		}
		TradesDetails := TradesMap[trade.Ticker]
		TradesDetails = append(TradesDetails, tradeDetail)
		TradesMap[trade.Ticker] = TradesDetails
	}
	response.ResponseBody.Trades = TradesMap

	return response, nil
}

func (t *TradeAPI) GetPortfolio(ctx context.Context, tx *gorm.DB) (*dto.GetPortfolioResponse, error) {
	response := &dto.GetPortfolioResponse{}
	// Fetching all trades from database
	conditions := &models.Trades{
		DeletedAt: sql.NullTime{},
	}
	trades, err := dao.DbClientTraderDao.FindAll(conditions, tx)
	if err != nil {
		return nil, err
	}
	// Get portfolio
	portfolio := t.GetPortfolioFromTrades(trades)
	response.ResponseBody.Portfolio = portfolio
	return response, nil
}

func (t *TradeAPI) GetReturns(ctx context.Context, tx *gorm.DB) (*dto.GetReturnsResponse, error) {

	// Fetching all trades from database
	conditions := &models.Trades{
		DeletedAt: sql.NullTime{},
	}
	trades, err := dao.DbClientTraderDao.FindAll(conditions, tx)
	if err != nil {
		return nil, err
	}
	// Get portfolio
	portfolio := t.GetPortfolioFromTrades(trades)

	// Calculating returns
	var returns float64
	for k, _ := range portfolio {
		// for returns assuming current price to be 100 as mentioned in problem statement
		returns += (100 - portfolio[k].AverageBuyPrice) * float64(portfolio[k].TotalUnits)
	}
	response := &dto.GetReturnsResponse{
		Returns: returns,
	}
	return response, nil
}
