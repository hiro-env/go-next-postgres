package services

import (
	"app/models"
	"app/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(ur *repositories.UserRepository) *UserService {
	return &UserService{userRepository: ur}
}

func (us UserService) GetAllUsers() ([]models.User, error) {
	return us.userRepository.FindAll()
}
