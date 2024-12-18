package controller

import (
	"car-rental/dtos"
	"car-rental/service"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type RentalController interface {
	Rent(c echo.Context) error
}

type RentalControllerImpl struct {
	RentService service.RentService
}

func NewRentController(rentService service.RentService) RentalController {
	return &RentalControllerImpl{
		RentService: rentService,
	}
}

func (ci *RentalControllerImpl) Rent(c echo.Context) error {
	var req dtos.RentRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validation failed")
	}

	claims, ok := c.Get("user").(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims")
	}

	userID := int(claims["user_id"].(float64))

	response, err := ci.RentService.RentCar(req, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
