package service

import (
	"car-rental/dtos"
	"car-rental/helper"
	"car-rental/models"
	"car-rental/repository"
	"fmt"
	"time"
)

type RentService interface {
	RentCar(req dtos.RentRequest, userID int) (*dtos.RentResponse, error)
	GetRentalReport(UserID int) ([]models.Rental, error)
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
	car, err := s.CarRepo.GetCarByID(req.CarID)
	if err != nil {
		return nil, err
	}

	if car.StockAvailability == 0 {
		return nil, fmt.Errorf("car not available")
	}

	totalCosts := float64(req.Duration) * car.RentalCosts

	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if user.DepositAmount < totalCosts {
		return nil, fmt.Errorf("insufficient balance")
	}

	err = s.CarRepo.UpdateCarStock(req.CarID, 1)
	if err != nil {
		return nil, err
	}

	err = s.UserRepo.DeductUserBalance(userID, totalCosts)
	if err != nil {
		return nil, err
	}

	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, req.Duration)

	rental := models.Rental{
		UserID:     userID,
		CarID:      req.CarID,
		StartDate:  startDate,
		EndDate:    endDate,
		Duration:   req.Duration,
		TotalCosts: totalCosts,
		Status:     "Active",
	}

	err = s.RentalRepo.CreateRental(&rental)
	if err != nil {
		return nil, err
	}

	invoice, err := helper.CreateInvoice(*user, rental, *car)
	if err != nil {
		return nil, err
	}

	response := &dtos.RentResponse{
		ID:            rental.ID,
		UserID:        userID,
		CarName:       car.Name,
		CarCategory:   car.Category,
		StartDate:     rental.StartDate.Format("2006-01-02"),
		EndDate:       rental.EndDate.Format("2006-01-02"),
		InvoiceUrl:    invoice.InvoiceUrl,
		DepositAmount: user.DepositAmount - totalCosts,
	}

	return response, nil
}

func (s *RentServiceImpl) GetRentalReport(userID int) ([]models.Rental, error) {
	rentals, err := s.RentalRepo.GetRentalReport(userID)
	if err != nil {
		return nil, err
	}

	return rentals, err
}
