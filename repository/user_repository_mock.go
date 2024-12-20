package repository

import (
	"car-rental/models"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	Mock mock.Mock
}

func (m *UserRepoMock) CreateUser(user *models.User) error {
	res := m.Mock.Called(user)

	return res.Error(1)
}

func (m *UserRepoMock) GetUserByEmail(email string) (*models.User, error) {
	res := m.Mock.Called(email)

	user := res.Get(0).(models.User)
	return &user, res.Error(1)
}

func (m *UserRepoMock) GetUserByID(userID int) (*models.User, error) {
	res := m.Mock.Called(userID)

	user := res.Get(0).(models.User)
	return &user, res.Error(1)
}

func (m *UserRepoMock) UpdateUserBalance(userID int, amount float64) (*models.User, error) {
	res := m.Mock.Called(userID, amount)

	user := res.Get(0).(models.User)
	return &user, res.Error(1)
}

func (m *UserRepoMock) DeductUserBalance(userID int, amount float64) error {
	res := m.Mock.Called(userID, amount)

	return res.Error(1)
}
