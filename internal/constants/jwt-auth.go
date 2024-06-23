package constants

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var LOGIN_EXPIRATION_DURATION = time.Now().Add(24 * time.Hour)
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("SUPERSECRET")
