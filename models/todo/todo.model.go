package models

import "time"

type Todo struct {
	Idx       uint      `form:"idx"`
	UserIdx   uint      `form:"todo_user_idx"`
	Article   string    `form:"todo_article"`
	Memo      string    `form:"todo_memo"`
	Check     bool      `form:"todo_check"`
	Important bool      `form:"todo_important"`
	CreateAt  time.Time `form:"todo_creat_at"`
	UpdateAt  time.Time `form:"todo_update_at"`
}

func (todo *Todo) TableName() string {
	return "todo"
}
