package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoServer/pkg/app"
	"todoServer/pkg/e"
	userservice "todoServer/services"
)

func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	users, err := userservice.GetAll()

	if err != nil {
		data["status"] = false
		data["error"] = err.Error()
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, data)
	}

	data["status"] = true
	data["users"] = users
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
