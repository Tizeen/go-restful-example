package auth

import "golang.org/x/crypto/bcrypt"

func Encrypt(source string) (string, error) {
	hashedBytes, error := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), error
}
