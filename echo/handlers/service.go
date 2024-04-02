package handlers

import (
	"app/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserService interface {
	SearchUsers(nickname string) ([]models.User, error)
	GetUserInfo(int64) (*models.UserInfo, error)
	UpdateUserInfo(int64, *models.UpdateUserInfo) error
}

type AccountService interface {
	Verify(echo.Context) bool
	CreateUser(*models.UserAuthRequest) (int64, error)
	Login(*models.UserAuthRequest) (int64, error)
	GetAllUsernames() ([]string, error)
	CreateJWT(int64) (string, error)
	CreateBeingDeletedCookie() *http.Cookie
	DeleteUser(int64) error
}
