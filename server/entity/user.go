package entity

import (
	"context"
	"time"

	"fmt"
	"regexp"
	"strings"

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
	UserUpdate struct {
		*pb.UpdateUserRequest
		Validated map[string]interface{}
	}
	// NullTime implements setter and getter for bson to provide nullable value
	NullTime struct {
		Time time.Time
	}

	// UserService is a general interface of User entity
	UserService interface {
		List(ctx context.Context, in *pb.UserFilter) ([]User, error)
		Count(ctx context.Context, in *pb.CountCriteria) (int64, error)
		Create(ctx context.Context, in *pb.CreateUserRequest) (*User, error)
		Update(ctx context.Context, ID string, in map[string]interface{}) error
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

const (
	nameMinLength  = 2
	nameMaxLength  = 10
	emailMaxLength = 80
)

var re = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func (u UserUpdate) Validate() error {
	if u.GetNameNull() == false {
		if err := validateName(u.GetNameValue()); err != nil {
			return err
		}
		u.Validated["name"] = u.GetNameValue()
	}

	if u.GetEmailNull() == false {
		if err := validateEmail(u.GetEmailValue()); err != nil {
			return err
		}
		u.Validated["email"] = u.GetEmailValue()
	}

	if u.GetActivatedNull() == false {
		u.Validated["activated"] = u.GetActivatedValue()
	}
	return nil
}

func validateName(name string) error {
	trimmedName := strings.Trim(name, " ")
	l := len([]rune(trimmedName))
	if l < nameMinLength {
		return fmt.Errorf("name could not be less than %d characters", nameMinLength)
	}

	if l > nameMaxLength {
		return fmt.Errorf("name could not be more than %d characters", nameMaxLength)
	}

	return nil
}

func validateEmail(email string) error {
	trimmedEmail := strings.Trim(email, " ")
	if trimmedEmail == "" {
		return fmt.Errorf("email is required")
	}

	if len([]rune(trimmedEmail)) > emailMaxLength {
		return fmt.Errorf("email can not contain more than %d characters", emailMaxLength)
	}

	if !re.MatchString(trimmedEmail) {
		return fmt.Errorf("enter a valid email address")
	}

	return nil
}
