package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	models "todoServer/models/user"
	"todoServer/pkg/setting"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(email string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		email,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todo",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(c *gin.Context) (uint, error) {
	auth := c.GetHeader("Authorization")
	parts := strings.Fields(auth)

	tokenClaims, err := jwt.ParseWithClaims(parts[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			user, err := models.GetUserByEmail(claims.Email)
			if err != nil {
				return 0, err
			}
			return user.Id, nil
		}
	}

	return 0, err
}
