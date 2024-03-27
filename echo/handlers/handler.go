package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	user, err := h.userService.GetAllUsers()
	if err != nil {

	}

	return c.JSON(http.StatusOK, user)
}
