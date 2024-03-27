package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type userHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) GetUsers(c echo.Context) error {
	user, err := h.userService.GetAllUsers()
	if err != nil {

	}

	return c.JSON(http.StatusOK, user)
}
