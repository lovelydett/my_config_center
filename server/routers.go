package server

import "net/http"

func StartServer() {
	// http.Handle needs a http.Handler that implements the ServeHTTP method
	// while http.HandleFunc needs a function that takes a http.ResponseWriter and a http.Request

	// Router 1: the default static file router that catch any path
	http.HandleFunc("/", staticFileHandler)

	// Start the HTTP server
	http.ListenAndServe("0.0.0.0:8080", nil)
}
