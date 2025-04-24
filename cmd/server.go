package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/efcan748/Golang-In-Memory/internal/server"
)

func main() {
	// Define command-line flags
	host := flag.String("host", "127.0.0.1", "Server host address")
	port := flag.String("port", "8080", "Server port number")
	cleanupInterval := flag.Int("cleanup", 1, "Cleanup interval in minutes")

	flag.Parse()

	// Initialize dependencies
	apiHandler := server.New(*cleanupInterval)
	router := server.NewRouter(apiHandler)

	// Construct server address
	addr := *host + ":" + *port

	// Start server
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Printf("Starting server on %s", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
