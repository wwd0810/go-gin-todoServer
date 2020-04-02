package models

import (
	"time"
)

type User struct {
	Id       uint      `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Username string    `json:"username"`
	CreateAt time.Time `json:"creat_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (user *User) TableName() string {
	return "tbl_user"
}
