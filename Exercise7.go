package main

import (
	"fmt"
	"net/http"
)

func HandleLegacy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	if r.URL.Path == "/legacy" {
		http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
	}
	fmt.Fprintln(w, "Welcome to version 2")
}
