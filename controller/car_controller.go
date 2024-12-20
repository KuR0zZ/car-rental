package controller

import (
	"car-rental/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CarController interface {
	GetAllCar(c echo.Context) error
	GetCarByID(c echo.Context) error
}

type CarControllerImpl struct {
	CarService service.CarService
}

func NewCarController(carService service.CarService) CarController {
	return &CarControllerImpl{
		CarService: carService,
	}
}

// @Summary      Get all cars
// @Description  Retrieves a list of cars
// @Tags         Cars
// @Accept       json
// @Produce      json
// @Param Authorization header string true "With the bearer started"
// @Success      200      {array}   models.Car
// @Failure      500      {object}  dtos.ErrorInternalServerError
// @Router       /cars [get]
// @Security     Bearer
func (ci *CarControllerImpl) GetAllCar(c echo.Context) error {
	cars, err := ci.CarService.GetAllCar()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, cars)
}

// @Summary      Get car by id
// @Description  Retrieves a car by id
// @Tags         Cars
// @Accept       json
// @Produce      json
// @Param Authorization header string true "With the bearer started"
// @Success      200      {object}  models.Car
// @Failure      404      {object}  dtos.ErrorNotFound
// @Failure      500      {object}  dtos.ErrorInternalServerError
// @Router       /cars/:id [get]
// @Security     Bearer
func (ci *CarControllerImpl) GetCarByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	car, err := ci.CarService.GetCarByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "car not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, car)
}
