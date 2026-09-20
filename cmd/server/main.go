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

	healthHandler := handlers.NewHealthHandler()
	mux.HandleFunc("GET /health", healthHandler.Health)

	loginHandler := handlers.NewLoginHandler()
	mux.HandleFunc("GET /", loginHandler.Index)
	mux.HandleFunc("POST /", loginHandler.Login)

	registerHandler := handlers.NewRegisterHandler()
	mux.HandleFunc("GET /register", registerHandler.Index)
	mux.HandleFunc("POST /register", registerHandler.Register)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
