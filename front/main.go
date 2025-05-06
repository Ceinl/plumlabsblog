package main

import (
	"log"
	"net/http"
)

func main() {
	frontendPort := "8080" 
	fs := http.FileServer(http.Dir("."))

	http.Handle("/", fs)

	log.Printf("Frontend server listening on port %s\n", frontendPort)

	err := http.ListenAndServe(":"+frontendPort, nil)
	if err != nil {
		log.Fatalf("Frontend server failed to start: %v", err)
	}
}
