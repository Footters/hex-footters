package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Footters/hex-footters/pkg/media/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()

	getContentHandler := httptransport.NewServer(
		endpoints.GetContent,
		DecodeHTTPGetContentRequest,
		EncodeHTTPResponse,
	)

	getAllContentsHandler := httptransport.NewServer(
		endpoints.GetAllContents,
		DecodeHTTPGetAllContentsRequest,
		EncodeHTTPResponse,
	)

	createContentHandler := httptransport.NewServer(
		endpoints.CreateContent,
		DecodeHTTPCreateContentRequest,
		EncodeHTTPResponse,
	)

	toLiveContentHandler := httptransport.NewServer(
		endpoints.SetContentToLive,
		DecodeHTTPSetContentLiveRequest,
		EncodeHTTPResponse,
	)

	r.Handle("/contents", createContentHandler).Methods("POST")
	r.Handle("/contents", getAllContentsHandler).Methods("GET")
	r.Handle("/contents/{id}", getContentHandler).Methods("GET")
	r.Handle("/contents/{id}/live", toLiveContentHandler).Methods("PUT")
	return r
}

// AccessControl func
func AccessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

// EncodeHTTPResponse func. To encode a generic response.
func EncodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
