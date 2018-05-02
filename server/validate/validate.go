package validate

import (
	"fmt"

	pb "github.com/Sharykhin/go-users-grpc/proto"
)

const (
	nameMinLength = 2
	nameMaxLength = 10
)

// UserCreateRequest validates income request for creating a new user
func UserCreateRequest(in *pb.CreateUserRequest) error {
	if err := validateName(in.Name); err != nil {
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
	return nil
}

func validateName(name string) error {
	l := len([]rune(name))
	if l < nameMinLength {
		return fmt.Errorf(fmt.Sprintf("name could not be less than %d characters", nameMinLength))
	}

	if l > 10 {
		return fmt.Errorf(fmt.Sprintf("name could not be more than %d characters", nameMaxLength))
	}

	return nil
}
