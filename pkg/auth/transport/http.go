package transport

import (
	"net/http"

	"github.com/Footters/hex-footters/pkg/auth/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	m := mux.NewRouter()

	registerHandler := httptransport.NewServer(
		endpoints.Register,
		endpoint.DecodeRegisterRequest,
		endpoint.EncodeResponse,
	)

	loginHandler := httptransport.NewServer(
		endpoints.Login,
		endpoint.DecodeLoginRequest,
		endpoint.EncodeResponse,
	)

	m.Handle("/register", registerHandler)
	m.Handle("/login", loginHandler)

	return m
}
