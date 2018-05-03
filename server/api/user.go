package api

import (
	"context"

	"log"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"github.com/Sharykhin/go-users-grpc/server/entity"
	"github.com/Sharykhin/go-users-grpc/server/logger/file"
	"github.com/Sharykhin/go-users-grpc/server/mongodb"
	"github.com/Sharykhin/go-users-grpc/server/validate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServer struct {
	storage entity.UserService
	debug   bool
}

func (s userServer) List(ctx context.Context, in *pb.UserFilter) (*pb.UserListReponse, error) {
	if s.debug {
		log.Printf("GRPC: Method <List> is called with: %v\n", in)
	}

	users, err := s.storage.List(ctx, in)
	if err != nil {
		file.Logger.Errorf("could not get users list: %v", err)
		return nil, status.Errorf(codes.Internal, "could not get list of users: %v", err)
	}
	response := &pb.UserListReponse{Users: make([]*pb.UserResponse, 0)}
	for _, u := range users {
		res := convertUserToResponse(u)
		response.Users = append(response.Users, &res)
	}
	return response, nil
}

func (s userServer) Count(ctx context.Context, in *pb.CountCriteria) (*pb.CountResponse, error) {
	if s.debug {
		log.Printf("GRPC: Method <Count> is called with: %v\n", in)
	}

	c, err := s.storage.Count(ctx, in)
	if err != nil {
		file.Logger.Errorf("could not make count: %v", err)
		return nil, status.Errorf(codes.Internal, "could not make count: %v", err)
	}

	return &pb.CountResponse{
		Count: c,
	}, nil
}

func (s userServer) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserResponse, error) {
	if s.debug {
		log.Printf("GRPC: Method <Create> is called with: %v\n", in)
	}
	if err := validate.UserCreateRequest(in); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed validation: %v", err)
	}
	u, err := s.storage.Create(ctx, in)
	if err != nil {
		file.Logger.Errorf("could not create a new users list: %v", err)
		return nil, status.Errorf(codes.Internal, "could not create a new user: %v", err)
	}
	res := convertUserToResponse(*u)
	return &res, nil
}

func (s userServer) Update(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Empty, error) {
	if s.debug {
		log.Printf("GRPC: Method <Update> is called with: %v\n", in)
	}

	if err := validate.UserUpdateRequest(in); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed validation: %v", err)
	}

	err := s.storage.Update(ctx, in.ID, in)
	if err != nil {
		file.Logger.Errorf("could not update user: %v", err)
		return nil, status.Errorf(codes.Internal, "could not update user: %v", err)
	}

	return &pb.Empty{}, nil
}

func (s userServer) Remove(ctx context.Context, in *pb.UserID) (*pb.Empty, error) {
	if s.debug {
		log.Printf("GRPC: Method <Remove> is called with: %v\n", in)
	}
	err := s.storage.Remove(ctx, in.ID)
	if err != nil {
		file.Logger.Errorf("could not remove user: %v", err)
		return nil, status.Errorf(codes.Internal, "could not remove user: %v", err)
	}
	return &pb.Empty{}, nil
}

// NewUserServer returns a new instance of server that would implements all methods to satisfy grpc interface
func NewUserServer(debug bool) *userServer {
	return &userServer{
		storage: mongodb.UserService,
		debug:   debug,
	}
}

func convertUserToResponse(u entity.User) pb.UserResponse {
	return pb.UserResponse{
		ID:        u.ID.Hex(),
		Name:      u.Name,
		Email:     u.Email,
		Activated: u.Activated,
		CreatedAt: u.CreatedAt.UTC().Format(entity.TimeFormat),
		DeletedAt: u.DeletedAt.String(),
	}
}
