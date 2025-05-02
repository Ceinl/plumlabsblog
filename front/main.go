package main

import (
	"log"
	"net/http"
)

func main() {
	// Define the port for the frontend server
	frontendPort := "8080" // You can change this port if needed

	// Create a file server handler pointing to the current directory (".")
	fs := http.FileServer(http.Dir("."))

	// Handle all requests using the file server
	http.Handle("/", fs)

	log.Printf("Frontend server listening on port %s\n", frontendPort)

	// Start the server
	err := http.ListenAndServe(":"+frontendPort, nil)
	if err != nil {
		log.Fatalf("Frontend server failed to start: %v", err)
	}
}
