package main

import (
	"fmt"
	"io"
	"net/http"
)

func HandleCount(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/count" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Send a POST request with text to count words")
		return
	}

	if r.Method == http.MethodPost {
		red, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		
		text := string(red)
		fmt.Fprintln(w, len(text))
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
