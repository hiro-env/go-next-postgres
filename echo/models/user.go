package models

import "time"

type User struct {
	ID        int64     `json:"id" db:"id"`
	UserID    string    `json:"userId" db:"user_id"`
	Nickname  string    `json:"nickname" db:"nickname"`
	Image     string    `json:"image" db:"image"`
	CreatedBy string    `json:"createdBy" db:"created_by"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedBy string    `json:"updatedBy" db:"updated_by"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
