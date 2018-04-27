package api

import (
	"context"

	"time"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"github.com/Sharykhin/go-users-grpc/server/mongodb"
)

type server struct {
	storage UserService
	debug   bool
}

type UserService interface {
	CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserResponse, error)
}

func (s server) CreateUser(context.Context, *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		ID:        "123",
		Name:      "John",
		Email:     "chapal@inbox.ru",
		CreatedAt: time.Now().UTC().String(),
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
