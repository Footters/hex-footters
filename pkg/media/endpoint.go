package endpoint

import (
	"context"

	"github.com/Footters/hex-footters/pkg/media"

	"github.com/go-kit/kit/endpoint"
)

type contentRequest struct {
	ID uint `json:"id"`
}

type contentResponse struct {
	C   *media.Content `json:"c"`
	Err error          `json:"err"`
}

// MakeGetContentEndpoint func
func MakeGetContentEndpoint(svc media.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(contentRequest)
		c, err := svc.FindContentByID(req.ID)

		if err != nil {
			return contentResponse{
				C:   c,
				Err: err,
			}, nil
		}
		return contentResponse{Err: err}, nil
	}
}
