package files

import (
	"net/http"
)

func UploadChunkHandler(w http.ResponseWriter, r *http.Request) {
	// Step1: Parse the file chunk
	chunk, err := parseChunk(r)
	if err != nil {
		http.Error(w, "Invalid file chunk", http.StatusBadRequest)
		return
	}

	// Step2: Conduct the file chunk upload
	err = uploadChunk(chunk)
	if err != nil {
		http.Error(w, "Failed to upload file chunk", http.StatusInternalServerError)
		return
	}
}

func MergeFileHanlder(w http.ResponseWriter, r *http.Request) {

}
