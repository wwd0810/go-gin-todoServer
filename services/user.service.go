package userservice

import (
	"log"
	models "todoServer/models/user"
)

func GetAll() ([]models.User, error) {
	log.Printf("들러온다")
	return models.GetAllUser()
}
