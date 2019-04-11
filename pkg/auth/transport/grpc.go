package transport

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	auth grpctransport.Handler
}
