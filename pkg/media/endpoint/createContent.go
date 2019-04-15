package endpoint

import (
	"context"

	"github.com/Footters/hex-footters/pkg/media"
	"github.com/go-kit/kit/endpoint"
)

// CreateContentRequest struct
type CreateContentRequest struct {
	Content media.Content `json:"content"`
}

// CreateContentResponse struct
type CreateContentResponse struct {
	Msg string `json:"msg"`
}

// MakeCreateContentEndpoint func
func MakeCreateContentEndpoint(svc media.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateContentRequest)
		err := svc.CreateContent(&req.Content)
		if err != nil {
			return nil, err
		}

		return CreateContentResponse{Msg: "Created"}, nil
	}
}
