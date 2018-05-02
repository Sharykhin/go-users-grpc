package entity

import (
	"context"
	"time"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"gopkg.in/mgo.v2/bson"
)

const (
	// TimeFormat is a common time format that must be used for all times that should be represent as a string
	TimeFormat = "2006-01-02T15:04:05"
)

type (
	// User is a struct that represents basic entity for the current grpc server
	User struct {
		ID        bson.ObjectId `bson:"_id"`
		Name      string        `bson:"name"`
		Email     string        `bson:"email"`
		Activated bool          `bson:"activated"`
		CreatedAt time.Time     `bson:"created_at"`
		DeletedAt NullTime      `bson:"deleted_at"`
	}
	// NullTime implements setter and getter for bson to provide nullable value
	NullTime struct {
		Time time.Time
	}

	// UserService is a general interface of User entity
	UserService interface {
		List(ctx context.Context, limit, offset int64) ([]User, error)
		Create(ctx context.Context, in *pb.CreateUserRequest) (*User, error)
		Update(ctx context.Context, ID string, in *pb.UpdateUserRequest) error
		Remove(ctx context.Context, ID string) error
	}
)

// GetBSON converts value that should be saved right in a mongodb
func (t NullTime) GetBSON() (interface{}, error) {
	if t.Time.IsZero() {
		return nil, nil
	}
	return t.Time.UTC(), nil
}

// SetBSON converts value from mongodb to a struct
func (t *NullTime) SetBSON(raw bson.Raw) error {
	var tt time.Time
	err := raw.Unmarshal(&tt)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

// String custom stringer implementation to return empty string instead of zero time
func (t NullTime) String() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.UTC().Format(TimeFormat)
}
