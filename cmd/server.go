package main

import (
	"log"
	"net/http"

	"github.com/efcan748/Golang-In-Memory/internal/server"
)

func main() {
	// Initialize dependencies
	apiHandler := server.New(1)
	router := server.NewRouter(apiHandler)

	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
