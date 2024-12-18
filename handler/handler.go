package handler

import (
	"car-rental/dtos"
	"car-rental/helper"
	"car-rental/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Handler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	Deposit(c echo.Context) error
}

type HandlerImpl struct {
	DB *gorm.DB
}

func NewHandlerImpl(db *gorm.DB) Handler {
	return &HandlerImpl{DB: db}
}

func (h *HandlerImpl) Register(c echo.Context) error {
	var req dtos.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := models.User{
		Email:    req.Email,
		Password: string(hashPassword),
	}

	err = h.DB.Create(&user).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var result dtos.EmailResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if !result.Success {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Errors[0])
	}

	res := dtos.RegisterResponse{
		Message: "Successfully Register New User",
		Data: map[string]interface{}{
			"id":             user.ID,
			"email":          user.Email,
			"deposit_amount": user.DepositAmount,
		},
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *HandlerImpl) Login(c echo.Context) error {
	var req dtos.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	var user models.User
	err := h.DB.Where("email = ?", req.Email).Take(&user).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := dtos.LoginResponse{
		Message: "Successfully logged in",
		Token:   tokenString,
	}
	return c.JSON(http.StatusOK, res)
}

func (h *HandlerImpl) Deposit(c echo.Context) error {
	var req dtos.DepositRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	claims, ok := c.Get("user").(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	userID := int(claims["user_id"].(float64))

	var user models.User
	err := h.DB.Model(&user).Clauses(clause.Returning{}).Where("user_id = ?", userID).Update("deposit_amount", gorm.Expr("deposit_amount + ?", req.DepositAmount)).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := dtos.DepositResponse{
		Message:       "Successfully Top Up Balance",
		DepositAmount: user.DepositAmount,
	}

	return c.JSON(http.StatusOK, res)
}
