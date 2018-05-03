package validate

import (
	"fmt"

	"regexp"

	"strings"

	pb "github.com/Sharykhin/go-users-grpc/proto"
)

const (
	nameMinLength  = 2
	nameMaxLength  = 10
	emailMaxLength = 80
)

var re = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

// UserCreateRequest validates income request for creating a new user
func UserCreateRequest(in *pb.CreateUserRequest) error {
	if err := validateName(in.Name); err != nil {
		return err
	}

	if err := validateEmail(in.Email); err != nil {
		return err
	}
	return nil
}

// UserUpdateRequest validates income request on updating user's data
func UserUpdateRequest(in *pb.UpdateUserRequest) error {
	if in.GetNameNull() == false {
		if err := validateName(in.GetNameValue()); err != nil {
			return err
		}
	}

	if in.GetEmailNull() == false {
		if err := validateEmail(in.GetEmailValue()); err != nil {
			return err
		}
	}
	return nil
}

func validateName(name string) error {
	trimmedName := strings.Trim(name, " ")
	l := len([]rune(trimmedName))
	if l < nameMinLength {
		return fmt.Errorf("name could not be less than %d characters", nameMinLength)
	}

	if l > nameMaxLength {
		return fmt.Errorf("name could not be more than %d characters", nameMaxLength)
	}

	return nil
}

func validateEmail(email string) error {
	trimmedEmail := strings.Trim(email, " ")
	if trimmedEmail == "" {
		return fmt.Errorf("email is required")
	}

	if len([]rune(trimmedEmail)) > emailMaxLength {
		return fmt.Errorf("email can not contain more than %d characters", emailMaxLength)
	}

	if !re.MatchString(trimmedEmail) {
		return fmt.Errorf("enter a valid email address")
	}

	return nil
}
