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
		ID:        u.ID.Hex(),
		Name:      u.Name,
		Email:     u.Email,
		Activated: u.Activated,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s server) Update(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Empty, error) {
	fmt.Printf("GRPC Update is called with: %v\n", in)
	err := s.storage.Update(ctx, in.ID, in)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not update user: %v", err)
	}

	return &pb.Empty{}, nil
}

func (s server) Users(context.Context, *pb.UserFilter) (*pb.UserListReponse, error) {
	return &pb.UserListReponse{
		Users: []*pb.UserResponse{},
	}, nil
}

func (s server) Remove(ctx context.Context, in *pb.UserID) (*pb.Empty, error) {
	fmt.Printf("GRPC Remove is called with: %v\n", in)
	err := s.storage.Remove(ctx, in.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not remove user: %v", err)
	}
	return &pb.Empty{}, nil
}

// NewServer returns a new instance of server that would implements all methods to satisfy grpc interface
func NewServer(debug bool) *server {
	return &server{
		storage: mongodb.UserService,
		debug:   debug,
	}
}
