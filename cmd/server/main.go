package main

import (
	"ajaycalicut17/expense-management-go/internal/handlers"
	"log"
	"net/http"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.LoginHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
