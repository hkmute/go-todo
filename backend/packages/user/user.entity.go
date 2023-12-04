package user

import "time"

type AppUserEntity struct {
	Id         uint       `json:"id"`
	Username   string     `json:"username"`
	Password   string     `json:"password,omitempty"`
	Token      string     `json:"token,omitempty"`
	Created_at *time.Time `json:"createdAt,omitempty"`
	Updated_at *time.Time `json:"updatedAt,omitempty"`
}
