package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"github.com/Sharykhin/go-users-grpc/server/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userService struct {
	collection string
	db         *mgo.Database
}

// UserService is a variable that references to a private struct that implements all necessary methods
// for managing users through mongodb database
var UserService userService

func init() {
	address := os.Getenv("MONGODB_ADDRESS")
	var err error
	session, err := mgo.Dial(address)
	if err != nil {
		log.Fatalf("could not lister mongodb on %s: %v", address, err)
	}

	if err = session.Ping(); err != nil {
		log.Fatalf("could not ping mongodb: %v", err)
	}

	db := session.DB(os.Getenv("MONGODB_DBNAME"))
	UserService = userService{
		collection: "users",
		db:         db,
	}

	//defer db.Close()
}

func applyCriteria(in []*pb.QueryCriteria) (map[string]interface{}, error) {
	criteria := bson.M{}
	for _, c := range in {
		switch c.Key {
		case "deleted_at":
			if c.Value == "false" {
				criteria[c.Key] = nil
			}
			if c.Value == "true" {
				criteria[c.Key] = bson.M{"$ne": nil}
			}
		default:
			return nil, fmt.Errorf("unexpected criteria %s", c.Key)
		}
	}
	return criteria, nil
}

func (s userService) List(ctx context.Context, in *pb.UserFilter) ([]entity.User, error) {
	var users []entity.User
	criteria, err := applyCriteria(in.Criteria)
	if err != nil {
		return nil, err
	}
	err = s.db.C(s.collection).Find(criteria).Skip(int(in.Offset)).Limit(int(in.Limit)).All(&users)
	if err != nil {
		log.Printf("MongoDB: got error on users list: %v\n", err)
		return nil, err
	}

	return users, nil
}

func (s userService) Count(ctx context.Context, in *pb.CountCriteria) (int64, error) {
	criteria, err := applyCriteria(in.Criteria)
	if err != nil {
		return 0, err
	}

	c, err := s.db.C(s.collection).Find(criteria).Count()
	if err != nil {
		log.Printf("MongoDB: got error on count call: %v\n", err)
		return 0, err
	}

	return int64(c), nil
}

func (s userService) Create(ctx context.Context, in *pb.CreateUserRequest) (*entity.User, error) {

	user := entity.User{
		ID:        bson.NewObjectId(),
		Name:      in.Name,
		Email:     in.Email,
		Activated: in.Activated,
		CreatedAt: time.Now().UTC(),
		DeletedAt: entity.NullTime{},
	}

	fmt.Printf("MongoDB: Create user: %v\n", user)
	err := s.db.C(s.collection).Insert(user)

	if err != nil {
		log.Printf("MongoDB: error: %v\n", err)
		return nil, fmt.Errorf("could not save a new user")
	}

	return &user, nil
}

func (s userService) Update(ctx context.Context, ID string, in *pb.UpdateUserRequest) error {
	updates := bson.M{}
	if in.GetNameNull() == false {
		updates["name"] = in.GetNameValue()
	}

	if in.GetEmailNull() == false {
		updates["email"] = in.GetEmailValue()
	}
	return s.db.C(s.collection).Update(bson.M{"_id": bson.ObjectIdHex(ID)}, bson.M{"$set": updates})
}

func (s userService) Remove(ctx context.Context, ID string) error {
	return s.db.C(s.collection).Update(bson.M{"_id": bson.ObjectIdHex(ID)}, bson.M{"$set": bson.M{"deleted_at": time.Now().UTC()}})
}

func (s userService) HardRemove(ctx context.Context, ID string) error {
	return s.db.C(s.collection).RemoveId(bson.ObjectIdHex(ID))
}
