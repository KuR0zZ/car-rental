package service

import (
	"car-rental/dtos"
	"car-rental/helper"
	"car-rental/models"
	"car-rental/repository"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req dtos.RegisterRequest) (*models.User, error)
	Login(req dtos.LoginRequest) (string, error)
	Deposit(req dtos.DepositRequest, userID int) (*models.User, error)
}

type UserServiceImpl struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
	}
}

func (s *UserServiceImpl) Register(req dtos.RegisterRequest) (*models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := models.User{
		Email:    req.Email,
		Password: string(hashPassword),
	}

	err = s.UserRepo.CreateUser(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	message := dtos.Message{
		From: dtos.Email{
			Email: "hello@demomailtrap.com",
			Name:  "Mailtrap Test",
		},
		To: []dtos.Email{
			{Email: "ferdinandeducation8@gmail.com"},
		},
		Subject:  "Account Registration",
		Text:     fmt.Sprintf("A new account with email: %s has been created", req.Email),
		Category: "Integration Test",
	}

	body, err := helper.EmailNotification(message)
	if err != nil {
		return nil, fmt.Errorf("failed to send email notification: %w", err)
	}

	var result dtos.EmailResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if !result.Success {
		return nil, fmt.Errorf("email error: %s", result.Errors[0])
	}

	return &user, nil
}

func (s *UserServiceImpl) Login(req dtos.LoginRequest) (string, error) {
	user, err := s.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		return "", fmt.Errorf("user not found: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", fmt.Errorf("password does not match: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserServiceImpl) Deposit(req dtos.DepositRequest, userID int) (*models.User, error) {
	user, err := s.UserRepo.UpdateUserBalance(userID, req.DepositAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to update user deposit amount: %w", err)
	}

	return user, nil
}
