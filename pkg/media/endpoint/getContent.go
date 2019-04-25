package endpoint

import (
	"context"

	"github.com/Footters/hex-footters/pkg/media"
	"github.com/Footters/hex-footters/pkg/media/provider/auth"
	"github.com/go-kit/kit/endpoint"
)

// GetContentRequest struct
type GetContentRequest struct {
	ID uint `json:"id"`
}

// GetContentResponse struct
type GetContentResponse struct {
	Content *media.Content `json:"content"`
	Email   string         `json:"email"`
}

// MakeGetContentEndpoint func
func MakeGetContentEndpoint(svc media.Service, asp auth.ServiceProvider) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetContentRequest)
		c, err := svc.FindContentByID(req.ID)

		if err != nil {
			return nil, err
		}

		lg, err2 := asp.Login()
		if err2 != nil {
			return nil, err2
		}
		return GetContentResponse{Content: c, Email: lg}, nil
	}
}
