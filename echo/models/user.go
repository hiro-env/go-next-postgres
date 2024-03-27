package models

import "time"

type User struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"userName" db:"username"`
	Nickname  string    `json:"nickname" db:"nickname"`
	Image     []byte    `json:"image" db:"image"`
	CreatedBy string    `json:"createdBy" db:"created_by"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedBy string    `json:"updatedBy" db:"updated_by"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
