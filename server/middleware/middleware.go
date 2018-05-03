package middleware

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryLogHandler logs all the income requests
func UnaryLogHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, h grpc.UnaryHandler) (interface{}, error) {
	log.Println("Income request:", info.FullMethod)
	return h(ctx, req)
}

// UnaryPanicHandler catches panics
func UnaryPanicHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, h grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("catch panic", err)
			err = status.Errorf(codes.Internal, "something went wrong: %v", r)
		}
	}()

	return h(ctx, req)
}
