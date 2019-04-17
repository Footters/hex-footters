package auth

import (
	"context"

	pb "github.com/Footters/hex-footters/pkg/pb"
	"google.golang.org/grpc"
)

// ServiceProvider interface
type ServiceProvider interface {
	Login() (string, error)
}

type serviceProvider struct {
	Ctx  context.Context
	Conn *grpc.ClientConn
}

// NewServiceProvider constructor
func NewServiceProvider(ctx context.Context, conn *grpc.ClientConn) ServiceProvider {
	return &serviceProvider{
		Ctx:  ctx,
		Conn: conn,
	}
}

// Login for provider auth
func (asp *serviceProvider) Login() (string, error) {
	cli := pb.NewAuthClient(asp.Conn)
	loginReq := &pb.LoginRequest{
		Email:    "david@lcarrascal.com",
		Password: "1",
	}
	svcResp, err := cli.Login(asp.Ctx, loginReq)
	if err != nil {
		return "", err
	}
	return svcResp.User.Email, nil
}
