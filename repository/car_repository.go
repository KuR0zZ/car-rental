package repository

import (
	"car-rental/models"

	"gorm.io/gorm"
)

type CarRepository interface {
	GetAvailableCarByID(carID int) (*models.Car, error)
	UpdateCarStock(carID int, decrement int) error
}

type CarRepoImpl struct {
	DB *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &CarRepoImpl{DB: db}
}

func (r *CarRepoImpl) GetAvailableCarByID(carID int) (*models.Car, error) {
	var car models.Car
	err := r.DB.Where("car_id = ? AND stock_availability > 0", carID).Take(&car).Error
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (r *CarRepoImpl) UpdateCarStock(carID int, decrement int) error {
	return r.DB.Model(&models.Car{}).Where("car_id = ?", carID).UpdateColumn("stock_availability", gorm.Expr("stock_availability - ?", decrement)).Error
}
