package routes

import (
	config "car-rental/config/database"
	"car-rental/controller"
	"car-rental/repository"
	"car-rental/service"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	db := config.InitDB()

	// Initialize Repository
	userRepository := repository.NewUserRepository(db)
	carRepository := repository.NewCarRepository(db)
	rentalRepository := repository.NewRentalRepository(db)

	// Initialize Service
	rentalService := service.NewRentService(carRepository, userRepository, rentalRepository)
	userService := service.NewUserService(userRepository)

	// Initialize Controller
	rentalController := controller.NewRentController(rentalService)
	userController := controller.NewUserController(userService)

	e.POST("/users/register", userController.Register)
	e.POST("/users/login", userController.Login)
	e.POST("/users/deposit", userController.Deposit)
	e.POST("/rentals", rentalController.Rent)
}
