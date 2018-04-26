package handler

import (
	"log"
	"net"
	"os"

	pb "github.com/Sharykhin/go-users-grpc/proto"
	"github.com/Sharykhin/go-users-grpc/server/api"
	"google.golang.org/grpc"
)

func ListenAndServe() error {

	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", os.Getenv("GRPC_ADDRESS"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := api.Server{}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	pb.RegisterUserServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
