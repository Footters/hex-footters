package content

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler interface
type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type contentHandler struct {
	contentService Service
}

// NewHandler constructor
func NewHandler(contentService Service) Handler {
	return &contentHandler{
		contentService: contentService,
	}
}

func (h *contentHandler) Get(w http.ResponseWriter, r *http.Request) {
	contents, _ := h.contentService.FindAllContents()
	res, _ := json.Marshal(contents)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
func (h *contentHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idVars := vars["id"]
	id, _ := strconv.Atoi(idVars)
	content, _ := h.contentService.FindContentByID(uint(id))

	res, _ := json.Marshal(content)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
func (h *contentHandler) Create(w http.ResponseWriter, r *http.Request) {

	var content Content
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&content)
	_ = h.contentService.CreateContent(&content)

	res, _ := json.Marshal(content)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
