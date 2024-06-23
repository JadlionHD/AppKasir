package utils

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}

func ComparePassword(hashedPassword []byte, inputPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, inputPassword)
	if err != nil {
		return false
	} else {
		return true
	}
}
