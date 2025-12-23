package http

import (
	"challenge/external"
	"net/http"
)

func AnalysisServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /analysis", func(w http.ResponseWriter, r *http.Request) {
		external.ReadStreamCorrectly()
	})

	// If I had more time I would have implement a middleware or a custom server
	mux.HandleFunc("/analysis", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	// To be sure to catch "all"
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	http.ListenAndServe(":8080", mux)
}
