package repositories

import (
	"app/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur UserRepository) FindAll() ([]models.User, error) {
	var slice []models.User
	slice = append(slice, models.User{UserID: "5"})
	return slice, nil
}
