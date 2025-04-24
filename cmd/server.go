package main

import (
	"log"
	"net/http"

	"github.com/efcan748/Golang-In-Memory/internal/api"
	"github.com/efcan748/Golang-In-Memory/internal/client"
)

func main() {
	// Initialize dependencies
	storeClient := client.New()
	router := api.NewRouter(storeClient)

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
