package main

import (
	"fmt"
	"net/http"
)

func HandlePing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusNotFound)
		return
	}
	if r.URL.Path != "/ping" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "pong")
}
