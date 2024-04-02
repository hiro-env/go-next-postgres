package models

type UserInfo struct {
	Username string `db:"username"`
	Image    []byte `db:"image"`
}

type JSONUserInfo struct {
	Username string `json:"username"`
	Image    string `json:"image"`
}

type UpdateUserInfo struct {
	Nickname string
	Image    []byte
}
