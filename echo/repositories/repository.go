package services

import "app/models"

type UserRepository interface {
	SelectUsers(string) ([]models.User, error)
	SelectUserInfo(int64) (*models.UserInfo, error)
	UpdateUser(int64, *models.UpdateUserInfo) error
}

type AccountRepository interface {
	InsertUser(*models.UserAuthRequest) (int64, error)
	SelectAllUsernames() ([]string, error)
	GetUser(string) *models.User
	Delete(int64) error
}
