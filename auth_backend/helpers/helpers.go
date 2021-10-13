package helpers

import (
	"os"
	"strconv"

	"github.com/abednarchuk/grpc_auth/auth_backend/errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bcryptCost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		return "", errors.InternalServerError
	}
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", errors.InternalServerError
	}
	return string(res), nil
}
