package handlers

import "app/models"

type UserService interface {
	SearchUsers(nickname string) ([]models.User, error)
}

type AccountService interface {
	CreateUser(request *models.UserRegisterRequest) (int64, error)
	GetAllUsernames() ([]string, error)
}
