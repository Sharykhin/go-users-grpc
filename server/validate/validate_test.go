package validate

import "testing"
import (
	"sync"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestUserCreateRequest(t *testing.T) {
	tt := []struct {
		name        string
		in          *pb.CreateUserRequest
		expectedErr error
	}{
		{
			name: "name is required",
			in: &pb.CreateUserRequest{
				Name:  "",
				Email: "test@test.com",
			},
			expectedErr: errors.New("name is required"),
		},
		{
			name: "name is less than 2 characters",
			in: &pb.CreateUserRequest{
				Name:  "j",
				Email: "test@test.com",
			},
			expectedErr: errors.New("name could not be less than 2 characters"),
		},
		{
			name: "name is more than 10 characters",
			in: &pb.CreateUserRequest{
				Name:  "test_test_test_test",
				Email: "test@test.com",
			},
			expectedErr: errors.New("name could not be more than 10 characters"),
		},
		{
			name: "email is required",
			in: &pb.CreateUserRequest{
				Name:  "test",
				Email: "",
			},
			expectedErr: errors.New("email is required"),
		},
	}

	var wg sync.WaitGroup

	for _, tc := range tt {
		wg.Add(1)
		t.Run(tc.name, func(t *testing.T) {
			err := UserCreateRequest(tc.in)
			require.Equal(t, tc.expectedErr.Error(), err.Error())
		})
	}
}
