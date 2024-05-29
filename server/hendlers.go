package server

import (
	"net/http"
	"strings"
)

func _checkMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	if !_checkMethod(w, r, "GET") {
		return
	}
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

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if !_checkMethod(w, r, "POST") {
		return
	}

}

func filterFilesHandler(w http.ResponseWriter, r *http.Request) {
	if !_checkMethod(w, r, "GET") {
		return
	}

}

func tagFileHandler(w http.ResponseWriter, r *http.Request) {
	if !_checkMethod(w, r, "POST") {
		return
	}
}

func untagFileHandler(w http.ResponseWriter, r *http.Request) {
	if !_checkMethod(w, r, "POST") {
		return
	}

}

func renameFileHandler(w http.ResponseWriter, r *http.Request) {
	if !_checkMethod(w, r, "POST") {
		return
	}

}

func deleteFileHandler(w http.ResponseWriter, r *http.Request) {
	if !_checkMethod(w, r, "DELETE") {
		return
	}

}

func downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	if !_checkMethod(w, r, "GET") {
		return
	}

}
