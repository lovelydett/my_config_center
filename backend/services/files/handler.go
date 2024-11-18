package files

import (
	"net/http"
)

func UploadChunkHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile(w, r)
}

func MergeFileHanlder(w http.ResponseWriter, r *http.Request) {

}
