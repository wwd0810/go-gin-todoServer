package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoServer/src/pkg/app"
	"todoServer/src/pkg/e"
	userservice "todoServer/src/services"
)

func GetUserTodos(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	todos, err := userservice.GetUserTodos(c)
	//users, err := models.GetAllUser()

	if err != nil {
		data["status"] = false
		data["error"] = err.Error()
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, data)
	}

	data["status"] = true
	data["todos"] = todos
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetTodoDay(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	todos, err := userservice.GetTodoDay(c)
	//users, err := models.GetAllUser()

	if err != nil {
		data["status"] = false
		data["error"] = err.Error()
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, data)
	}

	data["status"] = true
	data["todos"] = todos
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func CreateTodo(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	err := userservice.CreateTodo(c)

	if err != nil {
		if err.Error() == "article is none" {
			data["status"] = false
			data["error"] = "article is none"

			appG.Response(http.StatusBadRequest, e.ERROR, data)
			return
		}
		data["status"] = false
		data["error"] = "fail"

		appG.Response(http.StatusInternalServerError, e.ERROR, data)
		return
	}

	data["status"] = true

	appG.Response(http.StatusOK, e.SUCCESS, data)

}

func UpdateTodoMemo(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	err := userservice.UpdateTodoMemo(c)

	if err != nil {
		if err.Error() == "bad requests" {
			data["status"] = false
			data["error"] = "bad requests"

			appG.Response(http.StatusBadRequest, e.ERROR, data)
			return
		}
		data["status"] = false
		data["error"] = "fail"

		appG.Response(http.StatusInternalServerError, e.ERROR, data)
		return
	}

	data["status"] = true

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func UpdateTodoImportant(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	err := userservice.UpdateTodoImportant(c)

	if err != nil {
		if err.Error() == "bad requests" {
			data["status"] = false
			data["error"] = "bad requests"

			appG.Response(http.StatusBadRequest, e.ERROR, data)
			return
		}
		data["status"] = false
		data["error"] = "fail"

		appG.Response(http.StatusInternalServerError, e.ERROR, data)
		return
	}

	data["status"] = true

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func UpdateTodoChecked(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	err := userservice.UpdateTodoChecked(c)

	if err != nil {
		if err.Error() == "bad requests" {
			data["status"] = false
			data["error"] = "bad requests"

			appG.Response(http.StatusBadRequest, e.ERROR, data)
			return
		}
		data["status"] = false
		data["error"] = "fail"

		appG.Response(http.StatusInternalServerError, e.ERROR, data)
		return
	}

	data["status"] = true

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func DeleteTodo(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	err := userservice.DeleteTodo(c)

	if err != nil {
		if err.Error() == "bad requests" {
			data["status"] = false
			data["error"] = "bad requests"

			appG.Response(http.StatusBadRequest, e.ERROR, data)
			return
		}
		data["status"] = false
		data["error"] = "fail"

		appG.Response(http.StatusInternalServerError, e.ERROR, data)
		return
	}

	data["status"] = true

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
