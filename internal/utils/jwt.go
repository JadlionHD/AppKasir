package utils

import (
	"github.com/JadlionHD/AppKasir/internal/constants"
	jwt "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username    string `json:"username"`
	Permissions uint   `json:"permissions"`
	ID          uint   `json:id`
	jwt.RegisteredClaims
}

// var APPLICATION_NAME = "My Simple JWT App"

func GenerateToken(username string, permissions uint, id uint) string {
	claims := &Claims{
		Username:    username,
		Permissions: permissions,
		ID:          id,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(constants.LOGIN_EXPIRATION_DURATION),
		},
	}
	token := jwt.NewWithClaims(constants.JWT_SIGNING_METHOD, claims)
	tokenString, _ := token.SignedString(constants.JWT_SIGNATURE_KEY)

	return tokenString
}
