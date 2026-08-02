package main

import (
	"ajaycalicut17/expense-management-go/internal/handlers"
	"log"
	"net/http"
)

func main() {

	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("GET /", handlers.IndexLogin)

	http.HandleFunc("POST /", handlers.Login)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
