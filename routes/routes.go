package routes

import (
	config "car-rental/config/database"
	"car-rental/handler"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	db := config.InitDB()

	handler := handler.NewHandlerImpl(db)

	e.POST("/users/register", handler.Register)
}
