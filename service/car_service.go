package service

import (
	"car-rental/models"
	"car-rental/repository"
)

type CarService interface {
	GetAllCar() ([]models.Car, error)
	GetCarByID(carID int) (*models.Car, error)
}

type CarServiceImpl struct {
	CarRepo repository.CarRepository
}

func NewCarService(carRepo repository.CarRepository) CarService {
	return &CarServiceImpl{
		CarRepo: carRepo,
	}
}

func (s *CarServiceImpl) GetAllCar() ([]models.Car, error) {
	cars, err := s.CarRepo.GetAllCar()
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (s *CarServiceImpl) GetCarByID(carID int) (*models.Car, error) {
	car, err := s.CarRepo.GetCarByID(carID)
	if err != nil {
		return nil, err
	}

	return car, nil
}
