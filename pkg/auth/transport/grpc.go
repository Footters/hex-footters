package transport

import (
	"context"

	"github.com/Footters/hex-footters/pkg/auth/endpoint"
	"github.com/Footters/hex-footters/pkg/auth/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

// GRPCServer struct
type GRPCServer struct {
	login    grpctransport.Handler
	register grpctransport.Handler
}

// NewGRPCHandler func
func NewGRPCHandler(_ context.Context, endpoints endpoint.Endpoints) pb.AuthServer {
	registerHandler := grpctransport.NewServer(
		endpoints.Register,
		DecodeGRPCRegisterRequest,
		EncodeGRPCRegisterResponse,
	)

	loginHandler := grpctransport.NewServer(
		endpoints.Login,
		DecodeGRPCLoginRequest,
		EncodeGRPCLoginResponse,
	)

	return &GRPCServer{
		login:    loginHandler,
		register: registerHandler,
	}
}

// Login func
func (s *GRPCServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	_, response, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.LoginResponse), nil
}

// Register func
func (s *GRPCServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	_, response, err := s.register.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.RegisterResponse), nil
}

// DecodeGRPCRegisterRequest func
func DecodeGRPCRegisterRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.RegisterRequest)
	return endpoint.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}

// DecodeGRPCLoginRequest func
func DecodeGRPCLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.LoginRequest)
	return endpoint.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}

// EncodeGRPCRegisterResponse func
func EncodeGRPCRegisterResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.RegisterResponse)
	return &pb.RegisterResponse{
		Msg: res.Msg,
	}, nil
}

// EncodeGRPCLoginResponse func
func EncodeGRPCLoginResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.LoginResponse)
	user := &pb.User{
		Email:    res.User.Email,
		Password: res.User.Password,
	}

	return &pb.LoginResponse{
		User: user,
	}, nil
}
