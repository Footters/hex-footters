package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

// MakeRegisterEndpoint func
func MakeRegisterEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)

		err := svc.RegisterUser(&User{
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
func MakeLoginEndpoint(svc Service) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		u, err := svc.Login(req.Email, req.Password)
		if err != nil {
			return nil, err
		}

		return loginResponse{u}, nil
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
type loginResponse struct {
	User *User `json:"user"`
}

// DecodeRegisterRequest func
func DecodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeLoginRequest func
func DecodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// EncodeResponse func
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
