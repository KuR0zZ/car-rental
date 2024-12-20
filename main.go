package main

import (
	"log"
	"os"

	custom_middleware "car-rental/middleware"
	"car-rental/routes"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println(".env file not found, skipping...")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	validate := validator.New()

	e := echo.New()
	e.Validator = custom_middleware.NewValidate(validate)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	skipper := func(c echo.Context) bool {
		return c.Path() == "/users/login" || c.Path() == "/users/register" || c.Path() == "/swagger/*"
	}

	e.Use(custom_middleware.CustomJwtMiddleware(skipper))

	routes.Init(e)

	e.Logger.Fatal(e.Start(":" + port))
}
