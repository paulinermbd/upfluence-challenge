// Package http contains all middlewares and configuration needed for our server
package http

import (
	"challenge/internal/controller"
	"net/http"
)

func AnalysisServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /analysis", controller.GetAnalysisHandler)

	// To be sure to catch "all"
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}
