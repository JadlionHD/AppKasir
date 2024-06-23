package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JadlionHD/AppKasir/internal/constants"
	"github.com/JadlionHD/AppKasir/internal/models"
	"github.com/JadlionHD/AppKasir/internal/schemas"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func ValidateMiddleware(c *gin.Context) {
	tokenStr, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return constants.JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user schemas.UserSchema
		models.DB.First(&user, claims["id"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", gin.H{
			"id":       user.ID,
			"username": user.Username,
		})

		c.Next()
	} else {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
