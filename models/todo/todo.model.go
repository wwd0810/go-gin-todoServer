package models

import "time"

type Todo struct {
	Idx       uint      `json:"idx"`
	UserIdx   uint      `json:"todo_user_idx"`
	Article   string    `json:"todo_article"`
	Memo      string    `json:"todo_memo"`
	Check     bool      `json:"todo_check"`
	Important bool      `json:"todo_important"`
	CreateAt  time.Time `json:"todo_creat_at"`
	UpdateAt  time.Time `json:"todo_update_at"`
}

func (todo *Todo) TableName() string {
	return "todo"
}
