package handlers

import (
	"app/models"
	"encoding/base64"
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
	var request models.UserSearchRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	users, err := h.userService.SearchUsers(request.SearchTerm)
	if err != nil {
		// Logにerr.Error()を送信
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	userResponses := make([]*models.UserSearchResponse, len(users))
	for i, user := range users {
		userResponses[i] = user.ToUserResponse()
	}

	return c.JSON(http.StatusOK, userResponses)
}

func (h *UserHandler) GetMyPageInfo(c echo.Context) error {
	userID := c.Get("userID").(int64)
	userInfo, err := h.userService.GetUserInfo(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	base64Image := base64.StdEncoding.EncodeToString(userInfo.Image)
	userInfoResponse := &models.JSONUserInfo{
		Username: userInfo.Username,
		Image:    base64Image,
	}

	return c.JSON(http.StatusOK, userInfoResponse)
}

func (h *UserHandler) UpdateMypage(c echo.Context) error {
	userID := c.Get("userID").(int64)
	var request models.JSONUserInfo
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	decodedImage, err := base64.StdEncoding.DecodeString(request.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid image data")
	}

	userInfo := &models.UpdateUserInfo{
		Nickname: request.Username,
		Image:    decodedImage,
	}

	if err := h.userService.UpdateUserInfo(userID, userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "update succeeded")
}
