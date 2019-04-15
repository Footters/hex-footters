package endpoint

import (
	"context"

	"github.com/Footters/hex-footters/pkg/auth"
	"github.com/go-kit/kit/endpoint"
)

// RegisterRequest struct
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse struct
type RegisterResponse struct {
	Msg string `json:"err,omitempty"`
}

// MakeRegisterEndpoint func
func MakeRegisterEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)

		err := svc.RegisterUser(&auth.User{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			return nil, err
		}

		return RegisterResponse{"Register OK"}, nil
	}
}
