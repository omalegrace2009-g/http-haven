package main

import (
	"fmt"
	"net/http"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed!!", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	f := r.URL.Query().Get("name")
	if f == "" {
		f = "Guest!"
	}
	fmt.Fprintf(w, "Hello, %s!", f)
}
