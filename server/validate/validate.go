package validate

import (
	"fmt"

	pb "github.com/Sharykhin/go-users-grpc/proto"
)

const (
	nameMinLength = 2
	nameMaxLength = 10
)

// UserRequest validates income requeqst for creating a new user
func UserRequest(in *pb.CreateUserRequest) error {
	if err := validateName(in.Name); err != nil {
		return err
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
