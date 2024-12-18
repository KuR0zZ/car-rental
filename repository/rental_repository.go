package repository

import (
	"car-rental/models"

	"gorm.io/gorm"
)

type RentalRepository interface {
	CreateRental(rental *models.Rental) error
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
