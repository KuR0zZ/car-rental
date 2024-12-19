package controller

import (
	"car-rental/dtos"
	"car-rental/service"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	TopUp(c echo.Context) error
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
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request data")
	}

	user, err := ci.UserService.Register(req)
	if err != nil {
		if strings.Contains(err.Error(), "email already exists") {
			return echo.NewHTTPError(http.StatusConflict, "email is already registered")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	res := dtos.RegisterResponse{
		Message: "Successfully Register New User",
		Data: map[string]interface{}{
			"id":             user.ID,
			"name":           user.Name,
			"email":          user.Email,
			"deposit_amount": user.DepositAmount,
		},
	}

	return c.JSON(http.StatusCreated, res)
}

func (ci *UserControllerImpl) Login(c echo.Context) error {
	var req dtos.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request data")
	}

	tokenString, err := ci.UserService.Login(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	res := dtos.LoginResponse{
		Message: "Successfully logged in",
		Token:   tokenString,
	}
	return c.JSON(http.StatusOK, res)
}

func (ci *UserControllerImpl) TopUp(c echo.Context) error {
	var req dtos.TopUpRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request data")
	}

	claims, ok := c.Get("user").(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	userID := int(claims["user_id"].(float64))

	user, err := ci.UserService.TopUp(req, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	res := dtos.TopUpResponse{
		Message:       "Successfully Top Up Balance",
		UserID:        userID,
		DepositAmount: user.DepositAmount,
	}

	return c.JSON(http.StatusOK, res)
}
