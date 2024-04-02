package models

type UserAuthRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
