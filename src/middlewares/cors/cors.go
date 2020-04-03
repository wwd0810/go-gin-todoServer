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

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Request-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", strconv.Itoa(int(MaxAge)))
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}
