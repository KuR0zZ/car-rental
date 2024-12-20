package repository

import (
	"car-rental/models"

	"gorm.io/gorm"
)

type CarRepository interface {
	GetAllCar() ([]models.Car, error)
	GetCarByID(carID int) (*models.Car, error)
	UpdateCarStock(carID int, decrement int) error
}

type CarRepoImpl struct {
	DB *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &CarRepoImpl{DB: db}
}

func (r *CarRepoImpl) GetAllCar() ([]models.Car, error) {
	var cars []models.Car
	err := r.DB.Find(&cars).Error
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (r *CarRepoImpl) GetCarByID(carID int) (*models.Car, error) {
	var car models.Car
	err := r.DB.Take(&car, carID).Error
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (r *CarRepoImpl) UpdateCarStock(carID int, decrement int) error {
	return r.DB.Model(&models.Car{}).Where("car_id = ?", carID).UpdateColumn("stock_availability", gorm.Expr("stock_availability - ?", decrement)).Error
}
