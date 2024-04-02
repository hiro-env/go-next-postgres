package models

type UserSearchResponse struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Image    []byte `json:"image"`
}

func (u *User) ToUserResponse() *UserSearchResponse {
	return &UserSearchResponse{
		ID:       u.ID,
		Nickname: u.Nickname,
		Image:    u.Image,
	}
}
