package models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"todoServer/src/pkg/database"
)

func GetAllUser() ([]User, error) {
	var users []User

	db := database.GetDB()

	table := "tbl_user"
	query := db.Select(table + ".*")

	if err := query.Find(&users).Error; err != nil {
		log.Println(err)
		return users, nil
	}

	return users, nil
}

func GetUserByEmail(email string) (User, error) {
	db := database.GetDB()
	var user User
	//user := User{}

	table := "tbl_user"
	query := db.Select(table + ".*")

	if err := query.Where("email = ?", email).First(&user).Error; err != nil {

		log.Println(err)

		return user, err
	}

	return user, nil

}

func CreateUser(email string, password string, username string) error {
	db := database.GetDB()
	//user := User{}

	//var user User
	//c.BindJSON(&user)

	if err := db.Create(&User{Email: email, Password: password, Username: username, CreateAt: time.Now(), UpdateAt: time.Now()}).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
