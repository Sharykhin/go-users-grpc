package entity

import (
	"context"

	"time"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Email     string        `bson:"email"`
	Activated bool          `bson:"activated"`
	CreatedAt time.Time     `bson:"created_at"`
}

type UserService interface {
	Index(ctx context.Context, limit, offset int64) ([]User, error)
	CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*User, error)
	Update(ctx context.Context, ID string, in *pb.UpdateUserRequest) error
	Remove(ctx context.Context, ID string) error
}
