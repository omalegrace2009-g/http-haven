package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", HandlePing)
	http.HandleFunc("/hello", HandleHello)
	http.HandleFunc("/count", HandleCount)
	http.HandleFunc("/calculate", HandleCalculate)
	http.HandleFunc("/agent", HandleAgent)
	http.HandleFunc("/dashboard", HandleDashboard)
	http.HandleFunc("/legacy", HandleLegacy)
	http.HandleFunc("/v2", HandleLegacy)
	fmt.Println("Server Listening:")
	http.ListenAndServe(":8080", nil)
}
