package validators

import (
	"regexp"
	"unicode/utf8"

	"github.com/abednarchuk/grpc_auth/auth_backend/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Returns true if UserName is valid
func ValidateUserName(username string) bool {
	r := regexp.MustCompile(`^[a-z][a-z0-9_]{1,23}[^_.]$`)
	return r.MatchString(username)
}

// Returns true if Email is valid
func ValidateEmail(email string) bool {
	emailRegexp := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&’*+/=?^_\x60{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)
	return emailRegexp.MatchString(email)
}
func ValidatePassword(password string) bool {
	passwordLength := utf8.RuneCountInString(password)
	if passwordLength >= 7 && passwordLength <= 50 {
		return true
	}
	return false
}

func ValidateUser(user *models.User) error {
	if ValidateUserName(user.UserName) && ValidateEmail(user.Email) && ValidatePassword(user.Password) {
		return nil
	}

	return status.Errorf(codes.InvalidArgument, "Fields did not pass validation")
}
