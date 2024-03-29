package models

import "time"

type User struct {
	ID        int64     `db:"id"`
	Username  string    `db:"username"`
	Nickname  string    `db:"nickname"`
	Image     []byte    `db:"image"`
	CreatedBy string    `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedBy string    `db:"updated_by"`
	UpdatedAt time.Time `db:"updated_at"`
}
