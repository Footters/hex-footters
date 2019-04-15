package endpoint

import (
	"context"

	"github.com/Footters/hex-footters/pkg/auth"
	"github.com/go-kit/kit/endpoint"
)

// LoginRequest struct
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse struct
type LoginResponse struct {
	User *auth.User `json:"user"`
}

// MakeLoginEndpoint func
func MakeLoginEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		u, err := svc.Login(req.Email, req.Password)
		if err != nil {
			return nil, err
		}

		return LoginResponse{u}, nil
	}
}
