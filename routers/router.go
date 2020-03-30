package routers

import (
	"github.com/gin-gonic/gin"
	"todoServer/middlewares/cors"
	"todoServer/pkg/setting"
	v1 "todoServer/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(cors.CORS())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		apiv1.GET("/users/all", v1.GetUsers)

	}

	return r
}
