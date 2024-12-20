package repository

import (
	"car-rental/models"

	"github.com/stretchr/testify/mock"
)

type CarRepoMock struct {
	Mock mock.Mock
}

func (m *CarRepoMock) GetAllCar() ([]models.Car, error) {
	res := m.Mock.Called()

	cars := res.Get(0).([]models.Car)
	return cars, res.Error(1)
}

func (m *CarRepoMock) GetCarByID(carID int) (*models.Car, error) {
	res := m.Mock.Called(carID)

	car := res.Get(0).(models.Car)
	return &car, res.Error(1)
}

func (m *CarRepoMock) UpdateCarStock(carID int, decrement int) error {
	res := m.Mock.Called(carID, decrement)

	return res.Error(0)
}
