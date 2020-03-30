package models

import (
	_ "github.com/go-sql-driver/mysql"
	"todoServer/conf"
)

func GetAllUser() ([]User, error) {
	var users []User
	err := conf.DB.Find(users).Error

	if err != nil {
		return users, err
	}
	return users, nil
}
