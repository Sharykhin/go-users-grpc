package api

import (
	"context"

	"time"

	pb "github.com/Sharykhin/go-users-grpc/proto"
)

type Server struct {
}

func (s Server) CreateUser(context.Context, *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		ID:        "123",
		Name:      "John",
		Email:     "chapal@inbox.ru",
		CreatedAt: time.Now().UTC().String(),
	}, nil
}
