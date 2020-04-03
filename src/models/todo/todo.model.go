package models

import "time"

type Todo struct {
	Id        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Article   string    `json:"article"`
	Memo      string    `json:"memo"`
	Checked   bool      `json:"checked"`
	Important bool      `json:"important"`
	CreateAt  time.Time `json:"creat_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func (todo *Todo) TableName() string {
	return "tbl_todo"
}
