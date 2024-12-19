package controller

import (
	"car-rental/dtos"
	"car-rental/service"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type RentalController interface {
	Rent(c echo.Context) error
	RentalReport(c echo.Context) error
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

	res, err := ci.RentService.RentCar(req, userID)
	if err != nil {
		if strings.Contains(err.Error(), "car not available") {
			return echo.NewHTTPError(http.StatusNotFound, "car is not available for rent")
		}
		if strings.Contains(err.Error(), "insufficient balance") {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "your balance is insufficient to rent this car")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, res)
}

func (ci *RentalControllerImpl) RentalReport(c echo.Context) error {
	claims, ok := c.Get("user").(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	userID := int(claims["user_id"].(float64))

	rentals, err := ci.RentService.GetRentalReport(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	var res []dtos.RentalReportResponse
	for _, rental := range rentals {
		res = append(res, dtos.RentalReportResponse{
			ID:          rental.ID,
			UserID:      rental.UserID,
			CarName:     rental.Car.Name,
			CarCategory: rental.Car.Category,
			Duration:    rental.Duration,
			TotalCosts:  rental.TotalCosts,
			Status:      rental.Status,
		})
	}

	return c.JSON(http.StatusOK, res)
}
