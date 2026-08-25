package main

import (
	"ajaycalicut17/expense-management-go/internal/config"
	"ajaycalicut17/expense-management-go/internal/handlers"
	"log"
	"net/http"
	"time"
)

func main() {

	config := config.MustLoad()

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("GET /health", handlers.Health)

	mux.HandleFunc("GET /", handlers.IndexLogin)

	mux.HandleFunc("POST /", handlers.Login)

	mux.HandleFunc("GET /register", handlers.IndexRegister)

	mux.HandleFunc("POST /register", handlers.Register)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
