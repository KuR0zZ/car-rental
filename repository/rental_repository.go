package repository

import (
	"car-rental/models"

	"gorm.io/gorm"
)

type RentalRepository interface {
	CreateRental(rental *models.Rental) error
	GetRentalReport(UserID int) ([]models.Rental, error)
}

type RentalRepoImpl struct {
	DB *gorm.DB
}

func NewRentalRepository(db *gorm.DB) RentalRepository {
	return &RentalRepoImpl{DB: db}
}

func (r *RentalRepoImpl) CreateRental(rental *models.Rental) error {
	return r.DB.Create(rental).Error
}

func (r *RentalRepoImpl) GetRentalReport(UserID int) ([]models.Rental, error) {
	var rentals []models.Rental

	err := r.DB.Preload("Car").Where("user_id = ?", UserID).Find(&rentals).Error
	if err != nil {
		return nil, err
	}

	return rentals, nil
}
