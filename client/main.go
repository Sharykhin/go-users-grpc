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
	deleted := flag.String("deleted", "", "adsas")
	flag.Parse()

	switch *action {
	case "create":
		response, err := c.Create(context.Background(), &pb.CreateUserRequest{
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
		response, err := c.Update(context.Background(), &pb.UpdateUserRequest{
			ID:   *id,
			Name: &pb.UpdateUserRequest_NameValue{NameValue: "Carl"},
		})
		if err != nil {
			log.Fatalf("Error when calling Update: %v", err)
		}
		log.Printf("Response from server: %v", response)
	case "list":

		filter := &pb.UserFilter{
			Limit:  3,
			Offset: 0,
		}
		fmt.Println(*deleted)
		if *deleted == "true" {
			filter.Criteria = []*pb.QueryCriteria{
				{
					Key:   "deleted_at",
					Value: "true",
				},
			}
		} else if *deleted == "false" {
			filter.Criteria = []*pb.QueryCriteria{
				{
					Key:   "deleted_at",
					Value: "false",
				},
			}
		}

		response, err := c.List(context.Background(), filter)
		if err != nil {
			log.Fatalf("Error when calling Update: %v", err)
		}
		log.Printf("Response from server: %v", response)
	default:
		fmt.Println("specify an action")
	}

}
