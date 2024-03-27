package services

import (
	"app/models"
)

type userService struct {
	userRepository UserRepository
}

func NewUserService(ur UserRepository) *userService {
	return &userService{userRepository: ur}
}

func (us userService) GetAllUsers() ([]models.User, error) {
	return us.userRepository.FindAll()
}
