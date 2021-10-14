package validators

import (
	"regexp"
	"unicode/utf8"

	"github.com/abednarchuk/grpc_auth/auth_backend/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateUserName(username string) error {
	usernameRegexp := regexp.MustCompile(`[A-Za-z]{8,30}`)
	if usernameRegexp.Match([]byte(username)) {
		return nil
	}
	return status.Errorf(codes.AlreadyExists, "Username already in use")
}
func ValidateEmail(email string) error {
	emailRegexp := regexp.MustCompile(`[^@ \t\r\n]+@[^@ \t\r\n]`)
	if emailRegexp.Match([]byte(email)) {
		return nil
	}
	return status.Errorf(codes.AlreadyExists, "Email already in use")
}
func ValidatePassword(password string) error {
	passwordLength := utf8.RuneCountInString(password)
	if passwordLength >= 7 && passwordLength <= 50 {
		return nil
	}
	return status.Errorf(codes.InvalidArgument, "Password should be larger than 7 and less than 50 characters")
}

func ValidateUser(user *models.User) error {
	usernameErr := ValidateUserName(user.UserName)
	emailErr := ValidateEmail(user.Email)
	passwordErr := ValidatePassword(user.Password)
	var err []error

	if usernameErr != nil {
		err = append(err, usernameErr)
	}
	if emailErr != nil {
		err = append(err, emailErr)
	}
	if passwordErr != nil {
		err = append(err, passwordErr)
	}
	if len(err) > 0 {
		return status.Errorf(codes.InvalidArgument, "Fields did not pass validation", err)
	}

	return nil
}
