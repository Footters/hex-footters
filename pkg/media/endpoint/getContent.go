package endpoint

import (
	"context"

	"github.com/Footters/hex-footters/pkg/media"
	"github.com/go-kit/kit/endpoint"
)

// GetContentRequest struct
type GetContentRequest struct {
	ID uint `json:"id"`
}

// GetContentResponse struct
type GetContentResponse struct {
	Content *media.Content `json:"content"`
}

// MakeGetContentEndpoint func
func MakeGetContentEndpoint(svc media.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetContentRequest)
		c, err := svc.FindContentByID(req.ID)

		if err != nil {
			return nil, err
		}

		return GetContentResponse{Content: c}, nil
	}
}
