package endpoint

import (
	"context"

	"github.com/Footters/hex-footters/pkg/media"
	"github.com/go-kit/kit/endpoint"
)

// MakeSetContentLiveEndpoint func
func MakeSetContentLiveEndpoint(svc media.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetContentRequest)
		c, err := svc.FindContentByID(req.ID)
		if err != nil {
			return nil, err
		}

		err2 := svc.SetToLive(c)
		if err2 != nil {
			return nil, err
		}

		return GetContentResponse{Content: c}, nil
	}
}
