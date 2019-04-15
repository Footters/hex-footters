package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Footters/hex-footters/pkg/auth/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	m := http.NewServeMux()

	registerHandler := httptransport.NewServer(
		endpoints.Register,
		DecodeHTTPRegisterRequest,
		EncodeHTTPResponse,
	)

	loginHandler := httptransport.NewServer(
		endpoints.Login,
		DecodeHTTPLoginRequest,
		EncodeHTTPResponse,
	)

	m.Handle("/register", registerHandler)
	m.Handle("/login", loginHandler)

	return m
}

// DecodeHTTPRegisterRequest func
func DecodeHTTPRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeHTTPLoginRequest func
func DecodeHTTPLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// EncodeHTTPResponse func
func EncodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
