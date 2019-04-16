package endpoint

import (
	"context"
	"fmt"

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
}

// MakeGetContentEndpoint func
func MakeGetContentEndpoint(svc media.Service, asp auth.ServiceProvider) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetContentRequest)
		c, err := svc.FindContentByID(req.ID)

		if err != nil {
			return nil, err
		}
		fmt.Println("Calling via RPC", asp.Login())
		return GetContentResponse{Content: c}, nil
	}
}
