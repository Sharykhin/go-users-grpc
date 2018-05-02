package main

import (
	"context"
	"log"

	"flag"

	"fmt"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	action := flag.String("action", "", "a string")
	id := flag.String("id", "", "an id")
	flag.Parse()

	switch *action {
	case "create":
		response, err := c.CreateUser(context.Background(), &pb.CreateUserRequest{
			Name:      "John",
			Email:     "chapal@inbox.ru",
			Activated: false,
		})
		if err != nil {
			log.Fatalf("Error when calling CreateUser: %v", err)
		}
		log.Printf("Response from server: %v", response)
	case "remove":
		response, err := c.Remove(context.Background(), &pb.UserID{ID: *id})
		if err != nil {
			log.Fatalf("Error when calling Remove: %v", err)
		}
		log.Printf("Response from server: %v", response)
	case "update":

	default:
		fmt.Println("specify an action")
	}

}
