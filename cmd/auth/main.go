package main

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"

	"github.com/Footters/hex-footters/pkg/auth"
	authendpoint "github.com/Footters/hex-footters/pkg/auth/endpoint"
	"github.com/Footters/hex-footters/pkg/auth/storage/redisdb"
	authtransport "github.com/Footters/hex-footters/pkg/auth/transport"
	"github.com/Footters/hex-footters/pkg/pb"
)

func main() {
	// Logger
	logger := log.NewLogfmtLogger(os.Stdout)
	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	// Service
	uRepo := redisdb.NewRedisUserRepository(redisdb.NewRedisConnection("redis:6379"))
	svc := auth.NewService(uRepo)
	svc = auth.NewLogginMiddleware(logger, svc)

	// Endpoints
	endpoints := authendpoint.MakeServerEndpoints(svc)

	// Create the main listener.
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Log("Err listen", err)
	}

	// Setup cmux
	m := cmux.New(l)

	httpL := m.Match(cmux.HTTP1Fast())
	grpcL := m.Match(cmux.Any())

	// HTTPServer
	httpHandler := authtransport.NewHTTPHandler(endpoints)
	httpS := &http.Server{
		Handler: httpHandler,
	}

	// GRPCServer
	ctx := context.Background()
	grpcHandler := authtransport.NewGRPCHandler(ctx, endpoints)

	grpcS := grpc.NewServer()
	pb.RegisterAuthServer(grpcS, grpcHandler)

	// Using the muxed listeners for servers.
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	// Start serving!
	m.Serve()
}
