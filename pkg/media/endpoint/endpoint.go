package endpoint

import (
	"github.com/Footters/hex-footters/pkg/media"
	"github.com/Footters/hex-footters/pkg/media/provider/auth"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct
type Endpoints struct {
	GetContent       endpoint.Endpoint
	GetAllContents   endpoint.Endpoint
	CreateContent    endpoint.Endpoint
	SetContentToLive endpoint.Endpoint
}

// MakeServerEndpoints returns service Endoints
func MakeServerEndpoints(svc media.Service, asp auth.ServiceProvider) Endpoints {

	getContentEndpoint := MakeGetContentEndpoint(svc, asp)
	getAllContentsEndpoint := MakeGetAllContentsEndpoint(svc)
	createContentEndpoint := MakeCreateContentEndpoint(svc)
	setToLiveContentEndpoint := MakeSetContentLiveEndpoint(svc)

	return Endpoints{
		GetContent:       getContentEndpoint,
		GetAllContents:   getAllContentsEndpoint,
		CreateContent:    createContentEndpoint,
		SetContentToLive: setToLiveContentEndpoint,
	}
}
