package api

import (
	"net/http"
	"time"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func addMiddleWares(hf HandlerFunc, needAuth bool) HandlerFunc {
	// The execution order of the middlewares is reversed
	middlewares := []func(HandlerFunc) HandlerFunc{
		_timingMiddleWare,
	}

	if needAuth {
		middlewares = append(middlewares, _authMiddleWare)
	}

	for _, mw := range middlewares {
		hf = mw(hf)
	}
	return hf
}

func _timingMiddleWare(hf HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		hf(w, r)
		// Add header "x-process-time" to the response
		w.Header().Set("x-process-time", time.Since(start).String())
	}
}

func _authMiddleWare(hf HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the request has a valid session token
		// if not, return 401 Unauthorized
		// if valid, call the next handler
		hf(w, r)
	}
}
