package entity

import (
	"context"

	"time"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"gopkg.in/mgo.v2/bson"
)

const (
	TimeFormat = "2006-01-02T15:04:05"
)

type (
	User struct {
		ID        bson.ObjectId `bson:"_id"`
		Name      string        `bson:"name"`
		Email     string        `bson:"email"`
		Activated bool          `bson:"activated"`
		CreatedAt time.Time     `bson:"created_at"`
		DeletedAt NullTime      `bson:"deleted_at"`
	}

	NullTime struct {
		Time time.Time
	}
)

func (t NullTime) GetBSON() (interface{}, error) {
	if t.Time.IsZero() {
		return nil, nil
	}
	return t.Time.UTC(), nil
}

func (t *NullTime) SetBSON(raw bson.Raw) error {
	var tt time.Time
	err := raw.Unmarshal(&tt)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

func (t NullTime) String() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.UTC().Format(TimeFormat)
}

type UserService interface {
	List(ctx context.Context, limit, offset int64) ([]User, error)
	Create(ctx context.Context, in *pb.CreateUserRequest) (*User, error)
	Update(ctx context.Context, ID string, in *pb.UpdateUserRequest) error
	Remove(ctx context.Context, ID string) error
}
