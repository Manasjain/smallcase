package dao

import (
	"github.com/jinzhu/gorm"
	"smallcase/db/models"
)

var DbClientTraderDao DBClientTrader

type DBClientTrader interface {
	Find(condition *models.Trades, tx *gorm.DB) (*models.Trades, error)
	FindAll(condition *models.Trades, tx *gorm.DB) ([]*models.Trades, error)
	Save(data *models.Trades, tx *gorm.DB) error
	Update(condition *models.Trades, data *models.Trades, tx *gorm.DB) error
}

func init() {
	dao := &TradesGormImpl{}
	DbClientTraderDao = dao
}

//go:generate mockgen -source=./trades_db_interface.go -destination=./trades_db_interface_mock_impl.go -package dao
