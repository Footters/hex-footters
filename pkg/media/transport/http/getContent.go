package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Footters/hex-footters/pkg/media/endpoint"
	"github.com/gorilla/mux"
)

// DecodeHTTPGetContentRequest func
func DecodeHTTPGetContentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idVars := vars["id"]
	id, err := strconv.Atoi(idVars)

	if err != nil {
		return nil, err
	}

	request := endpoint.GetContentRequest{ID: uint(id)}

	return request, nil
}
