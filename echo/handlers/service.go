package handlers

import "app/models"

type UserService interface {
	GetAllUsers() ([]models.User, error)
}
