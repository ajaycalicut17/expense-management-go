package main

import (
	"ajaycalicut17/expense-management-go/internal/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.LoginHandler)

	http.ListenAndServe(":8080", nil)
}
