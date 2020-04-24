package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"todoServer/src/pkg/app"
	"todoServer/src/pkg/e"
	"todoServer/src/pkg/util"
	userservice "todoServer/src/services"
)

func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	users, err := userservice.GetAll()
	//users, err := models.GetAllUser()

	if err != nil {
		data["status"] = false
		data["error"] = err.Error()
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, data)
	}

	data["status"] = true
	data["users"] = users
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetUserbyToken(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	user, err := userservice.GetUser(c)

	if err != nil {
		data["status"] = false
		data["error"] = err.Error()
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, data)
	}

	data["status"] = true
	data["user"] = user
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetOrCreateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	err := userservice.GetOrCreate(c)

	if err != nil {
		log.Println("error : ", err.Error())
		if err.Error() == "bad requests" {
			data["status"] = false
			data["error"] = "bad requests"

			appG.Response(http.StatusBadRequest, e.ERROR_USER_CREATE, data)
			return
		}

		log.Println(err)
		data["status"] = false
		data["error"] = "Already Registered"

		appG.Response(http.StatusConflict, e.ERROR_USER_CREATE, data)
		return
	}
	data["status"] = true

	appG.Response(http.StatusOK, e.SUCCESS, data)

}

func GetUserByEmail(c *gin.Context) {
	//test := c.Param("user_email")
	//log.Println(test)
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	user, err := userservice.GetByEmail(c)

	if err != nil {
		data["status"] = false
		data["error"] = err.Error()
		appG.Response(http.StatusInternalServerError, e.ERROR_USER_EMAIL, data)
		return
	}

	//token, err := util.GenerateToken(user.Email, user.First + " " + user.Last)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
	}

	data["status"] = true
	data["user"] = user

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func Login(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	user, err := userservice.Login(c)

	if err != nil {
		log.Println(err)
		data["status"] = false
		data["error"] = "Login failed"
		appG.Response(http.StatusUnauthorized, e.ERROR_USER_CREATE, data)
		return
	}
	token, err := util.GenerateToken(user.Email, user.Email+user.Password)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
	}

	data["status"] = true
	data["user"] = user
	data["token"] = token

	//log.Println(util.ParseToken(token))

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	user, err := userservice.Login(c)

	if err != nil {
		log.Println(err)
		data["status"] = false
		data["error"] = "Login failed"
		appG.Response(http.StatusUnauthorized, e.ERROR_USER_CREATE, data)
		return
	}

	data["status"] = true
	data["user"] = user

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
