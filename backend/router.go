package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"wolf/services/files"
)

func initAppRouter() *mux.Router {
	// Use mux to support path params
	router := mux.NewRouter()

	// Add static file handler
	staticFileDir := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDir))
	router.PathPrefix("/static/").Handler(staticFileHandler)

	// Root handler
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/html/files.html")
	})

	// Add business handlers
	router.HandleFunc("/files", files.UploadFileHandler).Methods("POST")

	return router
}
