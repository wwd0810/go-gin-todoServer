package cors

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// CORS func
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		MaxAge := 6 * time.Hour
		//c.Header("Access-Control-Allow-Credentials", "false")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		//c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Max-Age", strconv.Itoa(int(MaxAge)))

		c.Next()
	}
}
