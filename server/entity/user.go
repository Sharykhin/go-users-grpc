package entity

import (
	"context"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID    bson.ObjectId `bson:"_id"`
	Name  string        `bson:"name"`
	Email string        `bson:"email"`
}

type UserService interface {
	CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*User, error)
}
