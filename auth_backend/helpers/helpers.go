package helpers

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var InternalServerError = status.Errorf(codes.Internal, "Internal server error")

func HashPassword(password string) (string, error) {
	bcryptCost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		return "", InternalServerError
	}
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", InternalServerError
	}
	return string(res), nil
}
