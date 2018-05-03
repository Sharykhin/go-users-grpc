package middleware

import (
	"context"

	"google.golang.org/grpc"
)

// Chain is a helper to allow chaning unary interceptors
func Chain(wrappers ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		l := len(wrappers)
		for i := l - 1; i >= 0; i-- {
			handler = wrapHandler(wrappers[i], info, handler)
		}

		return handler(ctx, req)
	}
}

func wrapHandler(
	wrapper grpc.UnaryServerInterceptor,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) grpc.UnaryHandler {

	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return wrapper(ctx, req, info, handler)
	}
}

// ChainUnaryServer creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryServer(one, two, three) will execute one before two before three, and three
// will see context changes of one and two.
func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	n := len(interceptors)

	if n > 1 {
		lastI := n - 1
		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			var (
				chainHandler grpc.UnaryHandler
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				if curI == lastI {
					return handler(currentCtx, currentReq)
				}
				curI++
				resp, err := interceptors[curI](currentCtx, currentReq, info, chainHandler)
				curI--
				return resp, err
			}

			return interceptors[0](ctx, req, info, chainHandler)
		}
	}

	if n == 1 {
		return interceptors[0]
	}

	// n == 0; Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(ctx, req)
	}
}

// ChainStreamServer creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryServer(one, two, three) will execute one before two before three.
// If you want to pass context between interceptors, use WrapServerStream.
func ChainStreamServer(interceptors ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	n := len(interceptors)
	if n > 1 {
		lastI := n - 1
		return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			var (
				chainHandler grpc.StreamHandler
				curI         int
			)

			chainHandler = func(currentSrv interface{}, currentStream grpc.ServerStream) error {
				if curI == lastI {
					return handler(currentSrv, currentStream)
				}
				curI++
				err := interceptors[curI](currentSrv, currentStream, info, chainHandler)
				curI--
				return err
			}

			return interceptors[0](srv, stream, info, chainHandler)
		}
	}
	if n == 1 {
		return interceptors[0]
	}

	// n == 0; Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(srv interface{}, stream grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return handler(srv, stream)
	}
}
