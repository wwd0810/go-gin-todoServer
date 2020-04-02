package models

import (
	"errors"
	"log"
	"strings"
	"time"
	"todoServer/pkg/database"
)

func GetUserTodo(userId uint) ([]Todo, error) {
	var todos []Todo

	db := database.GetDB()

	table := "tbl_todo"
	query := db.Select(table + ".*")

	if err := query.Where("user_id = ?", userId).Find(&todos).Error; err != nil {

		return todos, err
	}

	return todos, nil
}

func GetTodoDay(userId uint, day string) ([]Todo, error) {
	var todos []Todo

	db := database.GetDB()
	table := "tbl_todo"
	query := db.Select(table + ".*")

	typeCheck := len(strings.Split(day, "-"))

	if typeCheck == 3 {
		if err := query.Where("user_id = ?", userId).Where("Date(create_at) = ?", day).Find(&todos).Error; err != nil {
			return nil, err
		}
	} else if typeCheck == 2 {
		if err := query.Where("user_id = ?", userId).Where("Year(create_at) = ?", strings.Split(day, "-")[0]).Where("Month(create_at) = ?", strings.Split(day, "-")[1]).Find(&todos).Error; err != nil {
			return nil, err
		}
	} else {
		if err := query.Where("user_id = ?", userId).Where("Year(create_at) = ?", day).Find(&todos).Error; err != nil {
			return nil, err
		}
	}

	return todos, nil
}

func CreateTodo(userId uint, article string) error {
	db := database.GetDB()

	if err := db.Create(&Todo{UserId: userId, Article: article, Memo: "", Checked: false, Important: false, CreateAt: time.Now(), UpdateAt: time.Now()}).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func UpdateTodoMemo(todoId uint, memo string) error {
	var todo Todo

	db := database.GetDB()

	table := "tbl_todo"
	query := db.Select(table + ".*")

	if err := query.Where("id = ?", todoId).Find(&todo).Error; err != nil {
		return errors.New("todo is none")
	}

	todo.Memo = memo
	db.Save(todo)

	return nil
}

func UpdateTodoImportant(todoId uint) error {
	var todo Todo

	db := database.GetDB()

	table := "tbl_todo"
	query := db.Select(table + ".*")

	if err := query.Where("id = ?", todoId).Find(&todo).Error; err != nil {
		return errors.New("todo is none")
	}

	if todo.Important == false {
		todo.Important = true
	} else {
		todo.Important = false
	}
	db.Save(todo)

	return nil
}

func UpdateTodoChecked(todoId uint) error {
	var todo Todo

	db := database.GetDB()

	table := "tbl_todo"
	query := db.Select(table + ".*")

	if err := query.Where("id = ?", todoId).Find(&todo).Error; err != nil {
		return errors.New("todo is none")
	}

	if todo.Checked == false {
		todo.Checked = true
	} else {
		todo.Checked = false
	}
	db.Save(todo)

	return nil
}

func DeleteTodo(todoId uint) error {
	var todo Todo

	db := database.GetDB()

	table := "tbl_todo"
	query := db.Select(table + ".*")

	if err := query.Where("id = ?", todoId).Delete(&todo).Error; err != nil {
		return errors.New("todo is none")
	}

	return nil
}
