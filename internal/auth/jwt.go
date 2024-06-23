package auth

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(24) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")
