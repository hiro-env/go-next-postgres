package repositories

import (
	"app/models"
	"database/sql"
	"fmt"
)

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *accountRepository {
	return &accountRepository{db: db}
}

func (ar accountRepository) InsertUser(request *models.UserAuthRequest) (int64, error) {
	query := `
        INSERT INTO users (username, nickname, password, created_by, updated_by)
        VALUES ($1, $1, $2, $1, $1)
        RETURNING id
    `

	var userID int64
	err := ar.db.QueryRow(query, request.Username, request.Password).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user: %v", err)
	}

	return userID, nil
}

func (ar accountRepository) SelectAllUsernames() ([]string, error) {
	rows, err := ar.db.Query("SELECT username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usernames []string
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		usernames = append(usernames, username)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usernames, nil
}

func (ar accountRepository) GetUser(username string) *models.User {
	query := `
	SELECT id, username, password
	FROM users
	WHERE username = $1
`

	var user models.User
	err := ar.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil
	}

	return &user
}

func (ar accountRepository) Delete(userID int64) error {
	query := `DELETE FROM users WHERE id = $1`
	if _, err := ar.db.Exec(query, userID); err != nil {
		return fmt.Errorf("delete error %w", err)
	}

	return nil
}
