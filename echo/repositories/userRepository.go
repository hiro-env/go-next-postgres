package repositories

import (
	"app/models"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur userRepository) FindAll() ([]models.User, error) {
	var slice []models.User
	slice = append(slice, models.User{UserID: "5"})
	return slice, nil
}
