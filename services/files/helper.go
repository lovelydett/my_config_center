package files

import (
	"io"
	"net/http"
	"os"

	"wolf/utils"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	r.ParseMultipartForm(10 << 20) // 10MB

	// Get the file input stream
	file, handler, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Get other form data
	tags := r.Form["tags"]
	filename := r.FormValue("filename")
	description := r.FormValue("description")

	fileKey := utils.GetMD5Hash(handler.Filename)

	// Save the file to the disk
	dst, err := os.Create("./data/" + fileKey)
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// Save the file meta data to the database
	if err := insertFileMeta(filename, tags, description); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
}
