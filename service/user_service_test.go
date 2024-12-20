package service

import (
	"car-rental/dtos"
	"car-rental/models"
	"car-rental/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepoMock = &repository.UserRepoMock{Mock: mock.Mock{}}
var userService = UserServiceImpl{UserRepo: userRepoMock}

func TestTopUp(t *testing.T) {
	userRes := models.User{
		ID:            1,
		Name:          "John Doe",
		Email:         "John.Doe@example.com",
		Password:      "password123",
		DepositAmount: 46000000.00,
	}
	userRepoMock.Mock.On("UpdateUserBalance", 1, 40000000.00).Return(userRes, nil)

	user, err := userService.TopUp(dtos.TopUpRequest{DepositAmount: 40000000.00}, 1)
	assert.Nil(t, err)
	assert.Equal(t, userRes.DepositAmount, user.DepositAmount)
}
