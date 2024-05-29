package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	// Use mux to support path params
	router := mux.NewRouter()

	// TODO: split GET/POST/PUT/DELETE to different routers
	router.HandleFunc("/files", addMiddleWares(uploadFileHandler, true))
	router.HandleFunc("/filter-files", addMiddleWares(filterFilesHandler, true))
	router.HandleFunc("/tag-file", addMiddleWares(tagFileHandler, true))
	router.HandleFunc("/untag-file", addMiddleWares(untagFileHandler, true))
	router.HandleFunc("/rename-file", addMiddleWares(renameFileHandler, true))
	router.HandleFunc("/files/{file_id}", addMiddleWares(deleteFileHandler, true))

	// The default static file router that catch any other path
	router.HandleFunc("/", addMiddleWares(staticFileHandler, false))

	http.Handle("/", router)

	// Start the HTTP server
	http.ListenAndServe("0.0.0.0:8080", nil)
}
