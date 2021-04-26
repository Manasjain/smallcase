package dao

import (
	"github.com/jinzhu/gorm"
	"log"
	"smallcase/db/models"
	errors "smallcase/error"
	"smallcase/utils"
	"time"
)

type TradesGormImpl struct{}

func (impl *TradesGormImpl) Find(condition *models.Trades, tx *gorm.DB) (*models.Trades, error) {
	data := &models.Trades{}
	if err := tx.Where(condition).First(data).Error; err != nil {
		log.Println("Find Database Error: ", utils.Marshal(err))
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.NewError(errors.DatabaseRecordNotFound, errors.DatabaseRecordNotFound.String())
		}
		return nil, errors.NewError(errors.DatabaseServiceFailure, errors.DatabaseServiceFailure.String())
	}
	return data, nil
}

func (impl *TradesGormImpl) FindAll(condition *models.Trades, tx *gorm.DB) ([]*models.Trades, error) {
	data := &[]*models.Trades{}
	if err := tx.Where(condition).Find(data).Error; err != nil {
		log.Println("FindAll Database Error: ", utils.Marshal(err))
		return nil, errors.NewError(errors.DatabaseServiceFailure, errors.DatabaseServiceFailure.String())
	}
	if len(*data) == 0 {
		return nil, errors.NewError(errors.DatabaseRecordNotFound, errors.DatabaseRecordNotFound.String())
	}
	return *data, nil
}

func (impl *TradesGormImpl) Save(data *models.Trades, tx *gorm.DB) error {
	now := time.Now()
	data.CreatedAt = &now
	data.UpdatedAt = &now

	if err := tx.Create(data).Error; err != nil {
		log.Println("Save Database Error: ", utils.Marshal(err))
		return errors.NewError(errors.DatabaseServiceFailure, errors.DatabaseServiceFailure.String())
	}

	return nil
}

func (impl *TradesGormImpl) Update(condition *models.Trades, data *models.Trades, tx *gorm.DB) error {
	now := time.Now()
	data.UpdatedAt = &now

	if err := tx.Model(data).Where(condition).Update(data).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			log.Println("Update Database Error: ", utils.Marshal(err))
			return errors.NewError(errors.DatabaseRecordNotFound, errors.DatabaseRecordNotFound.String())
		}
		return errors.NewError(errors.DatabaseServiceFailure, errors.DatabaseServiceFailure.String())
	}
	return nil
}
