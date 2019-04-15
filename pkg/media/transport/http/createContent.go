package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Footters/hex-footters/pkg/media/endpoint"
)

// DecodeHTTPCreateContentRequest func
func DecodeHTTPCreateContentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CreateContentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}
