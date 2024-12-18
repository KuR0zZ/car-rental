package repository

import (
	"car-rental/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID int) (*models.User, error)
	UpdateUserBalance(userID int, amount float64) (*models.User, error)
	DeductUserBalance(userID int, amount float64) error
}

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepoImpl{DB: db}
}

func (r *UserRepoImpl) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepoImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepoImpl) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	err := r.DB.Where("user_id = ?", userID).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepoImpl) UpdateUserBalance(userID int, amount float64) (*models.User, error) {
	var user models.User
	err := r.DB.Model(&user).Clauses(clause.Returning{}).Where("user_id = ?", userID).Update("deposit_amount", gorm.Expr("deposit_amount + ?", amount)).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepoImpl) DeductUserBalance(userID int, amount float64) error {
	return r.DB.Model(&models.User{}).Where("user_id = ?", userID).UpdateColumn("deposit_amount", gorm.Expr("deposit_amount - ?", amount)).Error
}
