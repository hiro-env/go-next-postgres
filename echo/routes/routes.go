package routes

import (
	"app/handlers"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, userHandler *handlers.UserHandler, accountHandler *handlers.AccountHandler) {
	e.GET("/mypage", userHandler.GetMyPageInfo)
	e.POST("/users", userHandler.GetUsers)
	e.POST("/updateuser", userHandler.UpdateMypage)
	e.GET("/verify", accountHandler.Verify)
	e.GET("/logout", accountHandler.HandleLogout)
	e.GET("/delete", accountHandler.HandleDelete)
	e.POST("/register", accountHandler.Register)
	e.POST("/login", accountHandler.HandleLogin)
}
