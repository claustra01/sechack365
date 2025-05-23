package util

import (
	"github.com/claustra01/sechack365/pkg/cerror"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", cerror.ErrGeneratePasswordHash
	}
	return string(hash), nil
}

func CompareHashAndPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return cerror.ErrInvalidPassword
	}
	return nil
}
