package server

import (
	"net/http"
	"strings"
)

func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	// First get the file name from the URL
	fileName := strings.TrimPrefix(r.URL.Path, "/")
	filePath := "static/"
	if strings.HasSuffix(fileName, ".js") {
		w.Header().Set("Content-Type", "application/javascript")
		filePath += "js/" + fileName
	} else if strings.HasSuffix(fileName, ".css") {
		w.Header().Set("Content-Type", "text/css")
		filePath += "css/" + fileName
	} else if strings.HasSuffix(fileName, ".html") {
		w.Header().Set("Content-Type", "text/html")
		filePath += "html/" + fileName
	} else if strings.HasSuffix(fileName, ".png") || strings.HasSuffix(fileName, ".jpg") {
		w.Header().Set("Content-Type", "image/png")
		filePath += "img/" + fileName
	} else {
		w.Header().Set("Content-Type", "text/plain")
		filePath += fileName
	}

	http.ServeFile(w, r, filePath)
}
