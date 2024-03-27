package services

import "app/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
}
