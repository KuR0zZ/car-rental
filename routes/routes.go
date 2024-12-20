package routes

import (
	config "car-rental/config/database"
	"car-rental/controller"
	"car-rental/repository"
	"car-rental/service"

	_ "car-rental/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	carService := service.NewCarService(carRepository)

	// Initialize Controller
	rentalController := controller.NewRentController(rentalService)
	userController := controller.NewUserController(userService)
	carController := controller.NewCarController(carService)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/users/register", userController.Register)
	e.POST("/users/login", userController.Login)
	e.POST("/users/topup", userController.TopUp)
	e.POST("/rentals", rentalController.Rent)
	e.GET("/rentals", rentalController.RentalReport)
	e.GET("/cars", carController.GetAllCar)
	e.GET("/cars/:id", carController.GetCarByID)
}
