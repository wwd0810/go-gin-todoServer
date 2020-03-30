package models

import "time"

type User struct {
	Idx      uint      `json:"idx"`
	Email    string    `json:"user_email"`
	Password string    `json:"user_password"`
	Username string    `json:"user_username"`
	CreateAt time.Time `json:"user_creat_at"`
	UpdateAt time.Time `json:"user_update_at"`
}

func (user *User) TableName() string {
	return "user"
}
