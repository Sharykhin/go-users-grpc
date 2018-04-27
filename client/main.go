package main

import (
	"context"
	"log"

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
	response, err := c.CreateUser(context.Background(), &pb.CreateUserRequest{
		Name:  "John",
		Email: "chapal@inbox.ru",
	})
	if err != nil {
		log.Fatalf("Error when calling CreateUser: %v", err)
	}
	log.Printf("Response from server: %s", response)
}
