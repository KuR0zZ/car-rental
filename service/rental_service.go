package service

import (
	"car-rental/dtos"
	"car-rental/models"
	"car-rental/repository"
	"fmt"
)

type RentService interface {
	RentCar(req dtos.RentRequest, userID int) (*dtos.RentResponse, error)
}

type RentServiceImpl struct {
	CarRepo    repository.CarRepository
	UserRepo   repository.UserRepository
	RentalRepo repository.RentalRepository
}

func NewRentService(carRepo repository.CarRepository, userRepo repository.UserRepository, rentalRepo repository.RentalRepository) RentService {
	return &RentServiceImpl{
		CarRepo:    carRepo,
		UserRepo:   userRepo,
		RentalRepo: rentalRepo,
	}
}

func (s *RentServiceImpl) RentCar(req dtos.RentRequest, userID int) (*dtos.RentResponse, error) {
	car, err := s.CarRepo.GetAvailableCarByID(req.CarID)
	if err != nil {
		return nil, fmt.Errorf("car not available: %w", err)
	}

	totalCosts := float64(req.Duration) * car.RentalCosts

	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	if user.DepositAmount < totalCosts {
		return nil, fmt.Errorf("insufficient balance")
	}

	err = s.CarRepo.UpdateCarStock(req.CarID, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to update car stock: %w", err)
	}

	err = s.UserRepo.DeductUserBalance(userID, totalCosts)
	if err != nil {
		return nil, fmt.Errorf("failed to update user balance: %w", err)
	}

	rental := &models.Rental{
		UserID:     userID,
		CarID:      req.CarID,
		Duration:   req.Duration,
		TotalCosts: totalCosts,
		Status:     "Active",
	}

	err = s.RentalRepo.CreateRental(rental)
	if err != nil {
		return nil, fmt.Errorf("failed to create rental record: %w", err)
	}

	response := &dtos.RentResponse{
		ID:     rental.ID,
		UserID: userID,
		CarRent: map[string]string{
			"name":     car.Name,
			"category": car.Category,
		},
		DepositAmount: user.DepositAmount - totalCosts,
	}

	return response, nil
}
