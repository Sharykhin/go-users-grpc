package handler

import (
	"log"
	"net"
	"os"

	"fmt"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"github.com/Sharykhin/go-users-grpc/server/api"
	"google.golang.org/grpc"
)

func ListenAndServe() error {
	address := os.Getenv("GRPC_ADDRESS")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := api.Server{}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the service to the grpc one
	pb.RegisterUserServer(grpcServer, &s)
	// start the server
	fmt.Printf("Started listening on %s\n", address)
	return grpcServer.Serve(lis)
}
