package routers

import (
	"github.com/gin-gonic/gin"
	"todoServer/middlewares/cors"
	"todoServer/middlewares/jwt"
	"todoServer/pkg/setting"
	v1 "todoServer/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.CORS())

	gin.SetMode(setting.ServerSetting.RunMode)

	apiUser := r.Group("/api/v1/users")
	//apiUser.Use(jwt.JWT(true))
	apiUser.Use()

	{
		apiUser.GET("/all", v1.GetUsers)
		apiUser.GET("/user", v1.GetUserByEmail)
		apiUser.POST("/user", v1.GetOrCreateUser)
		apiUser.POST("/login", v1.Login)

	}

	apiTodo := r.Group("/api/v1/todo")
	apiTodo.Use(jwt.JWT(true))
	apiTodo.Use()

	{
		apiTodo.GET("/all", v1.GetUserTodos)
		apiTodo.GET("/day/:day", v1.GetTodoDay)
		apiTodo.POST("/create", v1.CreateTodo)
		apiTodo.PATCH("/patch-memo", v1.UpdateTodoMemo)
		apiTodo.PATCH("/important", v1.UpdateTodoImportant)
		apiTodo.PATCH("/checked", v1.UpdateTodoChecked)
		apiTodo.DELETE("/delete/:todo_id", v1.DeleteTodo)

	}
	return r
}
