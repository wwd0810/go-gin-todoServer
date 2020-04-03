package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	models "todoServer/src/models/todo"
	"todoServer/src/pkg/util"
)

func GetUserTodos(c *gin.Context) ([]models.Todo, error) {
	userId, err := util.ParseToken(c)

	//log.Println(userId)

	todos, err := models.GetUserTodo(userId)

	if err != nil {
		c.Abort()
		return todos, err
	}

	return todos, nil
}

func GetTodoDay(c *gin.Context) ([]models.Todo, error) {
	userId, err := util.ParseToken(c)

	day := c.Param("day")
	log.Println("day", day)

	todos, err := models.GetTodoDay(userId, day)

	if err != nil {
		c.Abort()
		return nil, err
	}

	return todos, nil
}

func CreateTodo(c *gin.Context) error {
	article := c.PostForm("article")

	userId, err := util.ParseToken(c)
	if article == "" {
		c.Abort()
		return errors.New("article is none")
	}

	if err != nil {
		c.Abort()
		return c.Error(err)
	}

	err = models.CreateTodo(userId, article)

	if err != nil {
		return c.Error(err)
	}

	return nil
}

func UpdateTodoMemo(c *gin.Context) error {
	todoId := c.PostForm("todo_id")
	memo := c.PostForm("memo")

	if todoId == "" || memo == "" {
		return errors.New("bad requests")
	}

	u64, err := strconv.ParseUint(todoId, 10, 32)

	if err != nil {
		fmt.Println(err)
	}

	wd := uint(u64)
	err = models.UpdateTodoMemo(wd, memo)

	if err != nil {
		if err.Error() == "todo is none" {
			c.Abort()
			return c.Error(err)
		}
		return err
	}

	return nil

}

func UpdateTodoImportant(c *gin.Context) error {
	todoId := c.PostForm("todo_id")

	if todoId == "" {
		return errors.New("bad requests")
	}

	u64, err := strconv.ParseUint(todoId, 10, 32)

	if err != nil {
		fmt.Println(err)
	}

	wd := uint(u64)
	err = models.UpdateTodoImportant(wd)

	if err != nil {
		if err.Error() == "todo is none" {
			c.Abort()
			return c.Error(err)
		}
		return err
	}

	return nil

}

func UpdateTodoChecked(c *gin.Context) error {
	todoId := c.PostForm("todo_id")

	if todoId == "" {
		return errors.New("bad requests")
	}

	u64, err := strconv.ParseUint(todoId, 10, 32)

	if err != nil {
		fmt.Println(err)
	}

	wd := uint(u64)
	err = models.UpdateTodoChecked(wd)

	if err != nil {
		if err.Error() == "todo is none" {
			c.Abort()
			return c.Error(err)
		}
		return err
	}

	return nil

}

func DeleteTodo(c *gin.Context) error {
	todoId := c.Param("todo_id")

	if todoId == "" {
		return errors.New("bad requests")
	}

	u64, err := strconv.ParseUint(todoId, 10, 32)

	if err != nil {
		fmt.Println(err)
	}

	wd := uint(u64)
	err = models.DeleteTodo(wd)

	if err != nil {
		if err.Error() == "todo is none" {
			c.Abort()
			return c.Error(err)
		}
		return err
	}

	return nil

}
