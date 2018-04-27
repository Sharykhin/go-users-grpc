package api

import (
	"context"

	"fmt"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"github.com/Sharykhin/go-users-grpc/server/entity"
	"github.com/Sharykhin/go-users-grpc/server/mongodb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	storage entity.UserService
	debug   bool
}

func (s server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserResponse, error) {
	fmt.Printf("GRPC CreateUser is called with: %v\n", in)
	u, err := s.storage.CreateUser(ctx, in)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not create a new user: %v", err)
	}

	return &pb.UserResponse{
		ID:    u.ID.Hex(),
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func (s server) Users(context.Context, *pb.UserFilter) (*pb.UserListReponse, error) {
	return &pb.UserListReponse{
		Users: []*pb.UserResponse{},
	}, nil
}

// NewServer returns a new instance of server that would implements all methods to satisfy grpc interface
func NewServer(debug bool) *server {
	return &server{
		storage: mongodb.UserService,
		debug:   debug,
	}
}
