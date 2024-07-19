package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"wolf/services/files"
)

func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request URL: ", r.URL.Path)
	// First get the file name from the URL
	fileName := strings.TrimPrefix(r.URL.Path, "/")
	if fileName == "" {
		fileName = "files.html"
	}
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

func initAppRouter() *mux.Router {
	// Use mux to support path params
	router := mux.NewRouter()

	// Add static file handler
	router.HandleFunc("/", staticFileHandler)

	// Add business handlers
	router.HandleFunc("/files", files.UploadFileHandler).Methods("POST")

	return router
}
