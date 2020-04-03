package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	models "todoServer/src/models/user"
)

func GetAll() ([]models.User, error) {

	return models.GetAllUser()
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePassword(hashedPW string, plainPW []byte) bool {
	byteHash := []byte(hashedPW)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPW)
	if err != nil {
		return false
	}
	return true
}

func GetOrCreate(c *gin.Context) error {
	email := c.PostForm("email")
	password := c.PostForm("password")
	username := c.PostForm("username")
	//log.Println(email)

	if email != "" || password != "" || username != "" {
		return errors.New("bad requests")
	}

	user, err := models.GetUserByEmail(email)
	log.Println(user)

	if err != nil {
		if err.Error() == "record not found" {
			err = models.CreateUser(email, hashAndSalt([]byte(password)), username)
			if err != nil {
				return c.Error(err)
			} else {
				return nil
			}
		} else {
			return c.Error(err)
		}
	}

	//회원이 있을 때 예외처리 바람
	return gin.Error{
		Err: nil,
	}
}

func GetByEmail(c *gin.Context) (models.User, error) {
	email := c.Param("user_email")

	log.Println(email)

	user, err := models.GetUserByEmail(email)

	if err != nil {
		return user, err
	}

	return user, nil
}

func Login(c *gin.Context) (models.User, error) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := models.GetUserByEmail(email)

	//log.Println(user)

	if err != nil {
		return user, c.Error(err)
	}

	check := comparePassword(user.Password, []byte(password))

	if !check {
		return user, gin.Error{
			Err:  nil,
			Type: 404,
			Meta: nil,
		}
	}
	return user, nil
}

//if err := db.Create(&User{Email:email, Password:password, Username:username, CreateAt:time.Now(), UpdateAt:time.Now()}).Error; err != nil {
//log.Println(err)
//return err
//}
