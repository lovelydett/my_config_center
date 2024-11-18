package files

import (
	"io"
	"net/http"

	"wolf/drivers"
	"wolf/utils"
)

type FileChunk struct {
	key   string
	index int
	total int
	size  int
	data  io.Reader
}

type ImurRequest struct {
	bucket   string // Bucket name
	key      string // Object name to upload
	uploadID string // Generated UploadId
}

var oss = drivers.GetOSSConnector()

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	r.ParseMultipartForm(10 << 20) // 10MB

	// Get the file input stream
	file, _, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	chunk := FileChunk{
		key:   r.FormValue("key"),
		index: utils.Atoi(r.FormValue("index")),
		total: utils.Atoi(r.FormValue("total")),
		size:  utils.Atoi(r.FormValue("size")),
		data:  file,
	}

}
