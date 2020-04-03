package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"todoServer/src/pkg/setting"
)

var (
	DB    *gorm.DB
	err   error
	DBErr error
)

type Database struct {
	*gorm.DB
}

func Setup() {
	log.Println("Init MariaDb ...")
	var db = DB
	config := setting.DatabaseSetting

	driver := config.Type
	database := config.Name
	username := config.User
	password := config.Password
	host := config.Host

	db, err = gorm.Open(driver, username+":"+password+"@tcp("+host+")/"+database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		DBErr = err
		fmt.Println("db err: ", err)
	}
	// 데이터 베이스 로그
	db.LogMode(true)
	//db.AutoMigrate(&models.User{})

	log.Println("Success Connect MariaDB")
	DB = db
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}

// GetDBErr helps you to get a connection
func GetDBErr() error {
	return DBErr
}
