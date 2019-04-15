package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"

	"github.com/Footters/hex-footters/pkg/auth"
	authendpoint "github.com/Footters/hex-footters/pkg/auth/endpoint"
	"github.com/Footters/hex-footters/pkg/auth/pb"
	"github.com/Footters/hex-footters/pkg/auth/storage/redisdb"
	authtransport "github.com/Footters/hex-footters/pkg/auth/transport"
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

	// HTTPHandler
	httpHandler := authtransport.NewHTTPHandler(endpoints)
	// Go HTTP!
	logger.Log("transport", "HTTP", "addr", ":8081")
	logger.Log("exit", http.ListenAndServe(":8081", httpHandler))

	// GRPCHandler
	ctx := context.Background()
	grpcHandler := authtransport.NewGRPCHandler(ctx, endpoints)
	grpc := grpc.NewServer()
	pb.RegisterAuthServer(grpc, grpcHandler)

	// Go GRPC!
	logger.Log("transport", "HTTP", "addr", ":8082")
	logger.Log("exit", http.ListenAndServe(":8082", grpc))
}
