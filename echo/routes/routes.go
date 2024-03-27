package routes

import (
	"app/handlers"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, userHandler *handlers.UserHandler) {
	e.GET("/users", userHandler.GetUsers)
}
