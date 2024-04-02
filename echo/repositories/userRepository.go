package repositories

import (
	"app/models"
	"database/sql"
	"fmt"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur userRepository) SelectUsers(nickname string) ([]models.User, error) {
	var users []models.User
	var rows *sql.Rows
	var err error

	if nickname != "" {
		query := "SELECT id, username, nickname, image, created_by, created_at, updated_by, updated_at FROM users WHERE nickname ILIKE $1"
		rows, err = ur.db.Query(query, "%"+nickname+"%")
	} else {
		query := "SELECT id, username, nickname, image, created_by, created_at, updated_by, updated_at FROM users"
		rows, err = ur.db.Query(query)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Nickname, &user.Image, &user.CreatedBy, &user.CreatedAt, &user.UpdatedBy, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur userRepository) SelectUserInfo(userID int64) (*models.UserInfo, error) {
	var user models.UserInfo

	query := "SELECT nickname, image FROM users WHERE id = $1"
	row := ur.db.QueryRow(query, userID)
	if err := row.Scan(&user.Username, &user.Image); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur userRepository) UpdateUser(userID int64, info *models.UpdateUserInfo) error {
	query := `UPDATE users SET nickname = $1, image = $2, updated_at = NOW() WHERE id = $3`

	_, err := ur.db.Exec(query, info.Nickname, info.Image, userID)
	if err != nil {
		return fmt.Errorf("update error: %w", err)
	}

	return nil
}
