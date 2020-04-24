package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"todoServer/src/middlewares/jwt"
	"todoServer/src/pkg/setting"
	v1 "todoServer/src/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//r.Use(cors.CORS())
	//r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PATCH", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

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

	apiGetUser := r.Group(("/api/v1/"))
	apiGetUser.Use(jwt.JWT(true))
	apiGetUser.Use()

	{
		apiGetUser.GET("/me", v1.GetUserbyToken)
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
