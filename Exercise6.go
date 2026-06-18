package main

import (
	"fmt"
	"net/http"
)

func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong Method", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/dashboard" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	g := r.Header.Get("X-API-Key")
	hash := "secret123"
	if g != hash {
		http.Error(w, "Bad Key", http.StatusUnauthorized)
		return
	} else {
		fmt.Fprint(w, "Welcome to the dashboard")
	}
}
