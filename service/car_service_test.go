package service

import (
	"car-rental/models"
	"car-rental/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var carRepoMock = &repository.CarRepoMock{Mock: mock.Mock{}}
var carService = CarServiceImpl{CarRepo: carRepoMock}

func TestGetAllCar(t *testing.T) {
	CarRes := []models.Car{
		{
			ID:                1,
			Name:              "Mercedes AMG G63",
			StockAvailability: 2,
			RentalCosts:       1000000.00,
			Category:          "SUV",
		},
	}
	carRepoMock.Mock.On("GetAllCar").Return(CarRes, nil)

	cars, err := carService.GetAllCar()
	assert.Nil(t, err)
	assert.NotNil(t, cars)
	assert.Equal(t, len(CarRes), len(cars), "The number of cars should match")

	for i, car := range cars {
		assert.Equal(t, CarRes[i].ID, car.ID, "Car ID should match")
		assert.Equal(t, CarRes[i].Name, car.Name, "Car name should match")
		assert.Equal(t, CarRes[i].StockAvailability, car.StockAvailability, "Stock availability should match")
		assert.Equal(t, CarRes[i].RentalCosts, car.RentalCosts, "Rental costs should match")
		assert.Equal(t, CarRes[i].Category, car.Category, "Car category should match")
	}
}

func TestGetCarByID(t *testing.T) {
	carRes := models.Car{
		ID:                1,
		Name:              "Mercedes AMG G63",
		StockAvailability: 2,
		RentalCosts:       1000000.00,
		Category:          "SUV",
	}
	carRepoMock.Mock.On("GetCarByID", 1).Return(carRes, nil)

	car, err := carService.GetCarByID(1)
	assert.Nil(t, err)
	assert.NotNil(t, car)
	assert.Equal(t, carRes.ID, car.ID, "Car ID should match")
	assert.Equal(t, carRes.Name, car.Name, "Car name should match")
	assert.Equal(t, carRes.StockAvailability, car.StockAvailability, "Stock availability should match")
	assert.Equal(t, carRes.RentalCosts, car.RentalCosts, "Rental costs should match")
	assert.Equal(t, carRes.Category, car.Category, "Car category should match")
}
