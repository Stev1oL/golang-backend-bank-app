package main

import (
	"fmt"
	"log"

	"github.com/steviol/golang-backend-bank-app/internal/config"
	"github.com/steviol/golang-backend-bank-app/internal/db"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.NewPostgresConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server starting on %s", serverAddr)
	// log.Fatal(nethttp.ListenAndServe(serverAddr, router))
}
