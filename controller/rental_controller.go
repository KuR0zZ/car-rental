package controller

import (
	"car-rental/dtos"
	"car-rental/service"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

// @Summary      Rent a car
// @Description  Allows a user to rent a car given a valid request.
// @Tags         Rentals
// @Accept       json
// @Produce      json
// @Param Authorization header string true "With the bearer started"
// @Param        request  body      dtos.RentRequest  true  "Rent Request"
// @Success      201      {object}  dtos.RentResponse
// @Failure      400      {object}  dtos.ErrorBadRequest
// @Failure      404      {object}  dtos.ErrorNotFound
// @Failure      422      {object}  dtos.ErrorUnprocessableEntity
// @Failure      500      {object}  dtos.ErrorInternalServerError
// @Router       /rentals/rent [post]
// @Security     Bearer
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "car not found")
		}
		if strings.Contains(err.Error(), "car not available") {
			return echo.NewHTTPError(http.StatusNotFound, "car is not available for rent")
		}
		if strings.Contains(err.Error(), "insufficient balance") {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "your balance is insufficient to rent this car")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusCreated, res)
}

// @Summary      Get rental report
// @Description  Retrieves a list of rentals associated with the user.
// @Tags         Rentals
// @Accept       json
// @Produce      json
// @Param Authorization header string true "With the bearer started"
// @Success      200      {array}   dtos.RentalReportResponse
// @Failure      500      {object}  dtos.ErrorInternalServerError
// @Router       /rentals/report [get]
// @Security     Bearer
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
			StartDate:   rental.StartDate.Format("2006-01-02"),
			EndDate:     rental.EndDate.Format("2006-01-02"),
			TotalCosts:  rental.TotalCosts,
			Status:      rental.Status,
		})
	}

	return c.JSON(http.StatusOK, res)
}
