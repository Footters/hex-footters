package endpoint

import (
	"context"

	"github.com/Footters/hex-footters/pkg/media"
	"github.com/go-kit/kit/endpoint"
)

// GetAllContentResponse struct
type GetAllContentResponse struct {
	Contents []media.Content `json:"contents"`
}

// MakeGetAllContentsEndpoint func
func MakeGetAllContentsEndpoint(svc media.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		cs, err := svc.FindAllContents()

		if err != nil {
			return nil, err
		}

		return GetAllContentResponse{Contents: cs}, nil
	}
}
