package main

import (
	"log"

	"github.com/alessandrolattao/gomyadmin/internal/server"
)

func main() {
	srv := server.NewServer()
	log.Println("Starting GoMyAdmin on http://localhost:8080")
	if err := srv.Start("8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
