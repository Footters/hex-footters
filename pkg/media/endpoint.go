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

// GetAllContentResponse struct
type GetAllContentResponse struct {
	Contents []Content `json:"contents"`
}

// MakeGetAllContentsEndpoint func
func MakeGetAllContentsEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		cs, err := svc.FindAllContents()

		if err != nil {
			return nil, err
		}

		return GetAllContentResponse{Contents: cs}, nil
	}
}

// DecodeGetAllContentsRequest func
func DecodeGetAllContentsRequest(_ context.Context, r *http.Request) (interface{}, error) {

	return r, nil
}

// CreateContentRequest struct
type CreateContentRequest struct {
	Content Content `json:"content"`
}

// CreateContentResponse struct
type CreateContentResponse struct {
	Msg string `json:"msg"`
}

// MakeCreateContentsEndpoint func
func MakeCreateContentsEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateContentRequest)
		err := svc.CreateContent(&req.Content)
		if err != nil {
			return nil, err
		}

		return CreateContentResponse{Msg: "Created"}, nil
	}
}

// DecodeCreateContentRequest func
func DecodeCreateContentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateContentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

// MakeSetContentLiveEndpoint func
func MakeSetContentLiveEndpoint(svc Service) endpoint.Endpoint {
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

// DecodeSetContentLiveRequest func
func DecodeSetContentLiveRequest(_ context.Context, r *http.Request) (interface{}, error) {
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
