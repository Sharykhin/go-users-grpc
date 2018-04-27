package mongodb

import "context"
import (
	"log"
	"os"

	"fmt"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"gopkg.in/mgo.v2"
)

type userService struct {
	collection string
	db         *mgo.Database
}

// UserService is a variable that references to a private struct that implements all necessary methods
// for managing users through mongodb database
var UserService userService

func (s userService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserResponse, error) {
	err := s.db.C(s.collection).Insert(in)
	if err != nil {
		return nil, fmt.Errorf("could not save a new user")
	}

	return nil, nil
}

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
