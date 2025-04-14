package main

import (
	"log"
	"plumlabs/back/api"
	"plumlabs/back/server"
	"plumlabs/back/storage"
)

func main() {
	db, err := storage.Open()
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	
	apiHandler := api.New(db)
	
	srv := server.NewServer(apiHandler)
	srv.SetupRoutes()
	
	log.Println("Starting server...")
	srv.StartWithGracefulShutdown()
}
