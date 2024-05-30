package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type FileHandler struct {
	BaseHandler
}

func (h *FileHandler) GET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId := vars["file_id"]

}

func SearchFileHandler(w http.ResponseWriter, r *http.Request) {
	// Get query params: key and keyword
	params := r.URL.Query()
	key := params.Get("key")
	keyword := params.Get("keyword")

	switch key {
	case "name":

	}
}
