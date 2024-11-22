package api

import (
	"net/http"

	"wolf/services"
)

func uploadChunkHandler(w http.ResponseWriter, r *http.Request) {
	// Step1: Parse the file chunk
	chunk, err := services.ParseChunk(r)
	if err != nil {
		http.Error(w, "Invalid file chunk", http.StatusBadRequest)
		return
	}

	// Step2: Conduct the file chunk upload
	err = services.UploadChunk(chunk)
	if err != nil {
		http.Error(w, "Failed to upload file chunk", http.StatusInternalServerError)
		return
	}
}

func mergeFileHanlder(w http.ResponseWriter, r *http.Request) {

}
