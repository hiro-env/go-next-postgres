package services

import (
	"app/models"
	"app/repositories"
)

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) *userService {
	return &userService{userRepository: ur}
}

func (us userService) SearchUsers(nickname string) ([]models.User, error) {
	return us.userRepository.SelectUsers(nickname)
}

func (us userService) GetUserInfo(userID int64) (*models.UserInfo, error) {
	return us.userRepository.SelectUserInfo(userID)
}

func (us userService) UpdateUserInfo(userID int64, info *models.UpdateUserInfo) error {
	return us.userRepository.UpdateUser(userID, info)
}
