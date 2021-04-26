package logic

import (
	"context"
	"github.com/jinzhu/gorm"
	"smallcase/db/models"
	"smallcase/dto"
)

type TradeAPI struct {
}
type Trader interface {
	AddTrade(ctx context.Context, trade *dto.AddTradeRequest, tx *gorm.DB) (*dto.AddTradeResponse, error)
	UpdateTrade(ctx context.Context, trade *dto.UpdateTradeRequest, tx *gorm.DB) (*dto.UpdateTradeResponse, error)
	RemoveTrade(ctx context.Context, tradeId *string, tx *gorm.DB) (*dto.RemoveTradeResponse, error)
	FetchTrades(ctx context.Context, tx *gorm.DB) (*dto.FetchTradeResponse, error)
	GetPortfolio(ctx context.Context, tx *gorm.DB) (*dto.GetPortfolioResponse, error)
	GetReturns(ctx context.Context, tx *gorm.DB) (*dto.GetReturnsResponse, error)

	//helper implementations
	CheckIfValidTrade(tradeRequest *dto.AddTradeRequest, tx *gorm.DB) (bool, error)
	ValidTradeUpdate(tradeData *models.Trades, updateTradeData *models.Trades, tx *gorm.DB) (bool, error)
	GetUpdateTradeData(updateTradeRequest *dto.UpdateTradeRequest, originalTradeData *models.Trades) *models.Trades
	GetPortfolioFromTrades(trades []*models.Trades) map[string]*dto.Portfolio
}

//go:generate mockgen -source=./interface.go -destination=./interface_mock_impl.go -package logic
