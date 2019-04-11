package endpoint

import (
	"context"

	"github.com/Footters/hex-footters/pkg/auth"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct
type Endpoints struct {
	Register endpoint.Endpoint
	Login    endpoint.Endpoint
}

// MakeServerEndpoints returns service Endoints
func MakeServerEndpoints(svc auth.Service) Endpoints {

	registerEndpoint := MakeRegisterEndpoint(svc)
	loginEndpoint := MakeLoginEndpoint(svc)

	return Endpoints{
		Register: registerEndpoint,
		Login:    loginEndpoint,
	}
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

// RegisterRequest struct
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse struct
type RegisterResponse struct {
	Msg string `json:"err,omitempty"`
}

// LoginRequest struct
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse struct
type LoginResponse struct {
	User *auth.User `json:"user"`
}
