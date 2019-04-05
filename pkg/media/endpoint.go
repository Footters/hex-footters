package media

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

// GetContentRequest struct
type GetContentRequest struct {
	ID uint `json:"id"`
}

// GetContentResponse struct
type GetContentResponse struct {
	Content *Content `json:"content"`
}

// MakeGetContentEndpoint func
func MakeGetContentEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetContentRequest)
		c, err := svc.FindContentByID(req.ID)

		if err != nil {
			return nil, err
		}

		return GetContentResponse{Content: c}, nil
	}
}

// DecodeGetContentRequest func
func DecodeGetContentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return nil, err
	}

	request := GetContentRequest{ID: uint(id)}

	return request, nil
}

// EncodeResponse func
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
