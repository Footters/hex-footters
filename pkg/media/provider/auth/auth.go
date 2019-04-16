package auth

import (
	"context"
	"fmt"

	authService "github.com/Footters/hex-footters/pkg/auth/pb"
	"google.golang.org/grpc"
)

// ServiceProvider interface
type ServiceProvider interface {
	Login() string
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
func (asp *serviceProvider) Login() string {
	cli := authService.NewAuthClient(asp.Conn)
	loginReq := &authService.LoginRequest{
		Email:    "davidl@carrascal.com",
		Password: "secret",
	}
	svcResp, err := cli.Login(asp.Ctx, loginReq)
	if err != nil {
		fmt.Println("Client Err", err)
	}

	return svcResp.User.Email
}
