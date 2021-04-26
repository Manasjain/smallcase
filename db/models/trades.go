package models

import (
	"database/sql"
	"time"
)

type Trades struct {
	Id           int64        `gorm:"column:id;primary_key"`
	TradeId      string       `gorm:"column:tradeId"`
	Ticker       string       `gorm:"column:ticker"`
	TradeType    string       `gorm:"column:tradeType"`
	TradingUnits int64        `gorm:"column:tradingUnit"`
	UnitPrice    float64      `gorm:"column:unitPrice"`
	CreatedAt    *time.Time   `gorm:"column:createdAt"`
	UpdatedAt    *time.Time   `gorm:"column:updatedAt"`
	DeletedAt    sql.NullTime `gorm:"column:deletedAt"`
}

func (t *Trades) TableName() string {
	return "trades"
}
