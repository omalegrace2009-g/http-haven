package main

import (
	"fmt"
	"net/http"
)

func HandleAgent(w http.ResponseWriter, r *http.Request) {
	g := r.Header.Get("User-Agent")
	fmt.Fprintln(w, g)
}
