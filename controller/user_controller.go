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

// @Summary      Register a new user
// @Description  Creates a new user account with the provided details.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request  body      dtos.RegisterRequest  true  "Register Request"
// @Success      201      {object}  dtos.RegisterResponse{message=string,data=models.User}
// @Failure      400      {object}  dtos.ErrorBadRequest
// @Failure      409      {object}  dtos.ErrorConflict
// @Failure      500      {object}  dtos.ErrorInternalServerError
// @Router       /users/register [post]
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

// @Summary      User login
// @Description  Authenticates a user and returns a JWT token.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request  body      dtos.LoginRequest  true  "Login Request"
// @Success      200      {object}  dtos.LoginResponse
// @Failure      400      {object}  dtos.ErrorBadRequest
// @Failure      401      {object}  dtos.ErrorUnauthorized
// @Failure      404      {object}  dtos.ErrorNotFound
// @Failure      500      {object}  dtos.ErrorInternalServerError
// @Router       /users/login [post]
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

// @Summary      Top up balance
// @Description  Allows a user to add balance to their account.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param Authorization header string true "With the bearer started"
// @Param        request  body      dtos.TopUpRequest  true  "Top Up Request"
// @Success      200      {object}  dtos.TopUpResponse
// @Failure      400      {object}  dtos.ErrorBadRequest
// @Failure      500      {object}  dtos.ErrorInternalServerError
// @Router       /users/topup [post]
// @Security     Bearer
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

	if req.DepositAmount <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "amount must be greater than zero")
	}

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
