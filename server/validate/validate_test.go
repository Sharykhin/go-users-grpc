package validate

import (
	"sync"
	"testing"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestUserCreateRequest(t *testing.T) {
	tt := []struct {
		name          string
		in            *pb.CreateUserRequest
		expectedErr   error
		expectedValid bool
	}{
		{
			name: "name is required",
			in: &pb.CreateUserRequest{
				Name:  "",
				Email: "test@test.com",
			},
			expectedErr:   errors.New("name is required"),
			expectedValid: false,
		},
		{
			name: "name is less than 2 characters",
			in: &pb.CreateUserRequest{
				Name:  "j",
				Email: "test@test.com",
			},
			expectedErr:   errors.New("name could not be less than 2 characters"),
			expectedValid: false,
		},
		{
			name: "name is more than 10 characters",
			in: &pb.CreateUserRequest{
				Name:  "test_test_test_test",
				Email: "test@test.com",
			},
			expectedErr:   errors.New("name could not be more than 10 characters"),
			expectedValid: false,
		},
		{
			name: "email is required",
			in: &pb.CreateUserRequest{
				Name:  "test",
				Email: "",
			},
			expectedErr:   errors.New("email is required"),
			expectedValid: false,
		},
		{
			name: "email is more than 20 characters",
			in: &pb.CreateUserRequest{
				Name:  "test",
				Email: "testtesttesttesttesttest@test.com",
			},
			expectedErr:   errors.New("email can not contain more than 20 characters"),
			expectedValid: false,
		},
		{
			name: "email is not valid",
			in: &pb.CreateUserRequest{
				Name:  "test",
				Email: "test@",
			},
			expectedErr:   errors.New("enter a valid email address"),
			expectedValid: false,
		},
		{
			name: "all data is valid",
			in: &pb.CreateUserRequest{
				Name:  "test",
				Email: "test@test.com",
			},
			expectedErr:   nil,
			expectedValid: true,
		},
	}

	var wg sync.WaitGroup
	for _, tc := range tt {
		wg.Add(1)
		go t.Run(tc.name, func(t *testing.T) {
			wg.Done()
			err := UserCreateRequest(tc.in)
			if tc.expectedValid == false {
				require.Equal(t, tc.expectedErr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
	wg.Wait()
}
