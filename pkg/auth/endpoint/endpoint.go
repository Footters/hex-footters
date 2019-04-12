package endpoint

import (
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
