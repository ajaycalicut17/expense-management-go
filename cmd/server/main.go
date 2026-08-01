package main

import (
	"ajaycalicut17/expense-management-go/internal/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.LoginHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
