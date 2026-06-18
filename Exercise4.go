package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func HandleCalculate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/calculate" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	op := r.URL.Query().Get("op")
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	aC, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	bC, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, "Bad Reuest", http.StatusBadRequest)
		return
	}
	switch op {
	case "add":
		fmt.Fprintln(w, "Result is: ", aC+bC)
	case "subtract":
		fmt.Fprintln(w, "Result is: ", aC-bC)
	case "multiply":
		fmt.Fprintln(w, "Result is: ", aC*bC)
	case "divide":
		if bC == 0 {
			http.Error(w, "Cannot Divide by 0, Bad Request", http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, "Result is: ", aC/bC)
	default:
		http.Error(w, "Bad REquest", http.StatusBadRequest)
		return
	}
}
