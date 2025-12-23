package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/analysis", analysisHandler)
	http.ListenAndServe(":8080", nil)
}
