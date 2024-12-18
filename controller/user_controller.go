package controller

import (
	"car-rental/dtos"
	"car-rental/service"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	Deposit(c echo.Context) error
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (ci *UserControllerImpl) Register(c echo.Context) error {
	var req dtos.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := ci.UserService.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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

func (ci *UserControllerImpl) Login(c echo.Context) error {
	var req dtos.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	tokenString, err := ci.UserService.Login(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Invalid email or password")
		}
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := dtos.LoginResponse{
		Message: "Successfully logged in",
		Token:   tokenString,
	}
	return c.JSON(http.StatusOK, res)
}

func (ci *UserControllerImpl) Deposit(c echo.Context) error {
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

	user, err := ci.UserService.Deposit(req, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := dtos.DepositResponse{
		Message:       "Successfully Top Up Balance",
		DepositAmount: user.DepositAmount,
	}

	return c.JSON(http.StatusOK, res)
}
