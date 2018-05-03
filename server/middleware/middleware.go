package middleware

import (
	"context"
	"log"

	"github.com/Sharykhin/go-users-grpc/server/logger/file"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryLogHandler logs all the income requests
func UnaryLogHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, h grpc.UnaryHandler) (interface{}, error) {
	log.Println("Method:", info.FullMethod, "Request:", req)
	return h(ctx, req)
}

// UnaryPanicHandler catches panics
func UnaryPanicHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, h grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			file.Logger.Errorf("catch panic: %v", r)
			err = status.Errorf(codes.Internal, "something went wrong: %v", r)
		}
	}()

	return h(ctx, req)
}
