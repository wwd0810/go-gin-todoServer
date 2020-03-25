package models

import "time"

type User struct {
	Idx      uint      `form:"idx"`
	Email    string    `form:"user_email"`
	Password string    `form:"user_password"`
	Username string    `form:"user_username"`
	CreateAt time.Time `form:"user_creat_at"`
	UpdateAt time.Time `form:"user_update_at"`
}

func (user *User) TableName() string {
	return "user"
}
