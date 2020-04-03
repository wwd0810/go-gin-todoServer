package jwt

import (
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"todoServer/src/pkg/setting"
)

type APIError struct {
	Code int
	Msg  string
}

func RespondWithError(code int, msg string, c *gin.Context) {
	c.JSON(code, &APIError{Code: code, Msg: msg})
	c.Abort()
}

func JWT(encoded bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		parts := strings.Fields(auth)

		if auth == "" {
			RespondWithError(401, "API token required", c)
			return
		}
		if strings.ToLower(parts[0]) != "bearer" {
			RespondWithError(401, "Authorization header must start with Bearer", c)
			return
		} else if len(parts) == 1 {
			RespondWithError(401, "Token not found", c)
			return
		} else if len(parts) > 2 {
			RespondWithError(401, "Authorization header must be Bearer and token", c)
			return
		}

		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing me thod: %v", token.Header["alg"])
			}

			mySecertKey := setting.AppSetting.JwtSecret

			if encoded {
				var replacer = strings.NewReplacer("_", "/", "-", "+")
				mySecertKey = replacer.Replace(mySecertKey)

				base64Decoded, _ := base64.StdEncoding.DecodeString(mySecertKey)

				return base64Decoded, nil
			}

			return []byte(mySecertKey), nil

		})

		if token.Valid {
			c.Next()
		} else {
			RespondWithError(401, err.Error(), c)
			return
		}
	}
}
