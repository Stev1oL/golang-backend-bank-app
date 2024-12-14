package main

import (
	"fmt"
	"log"
	nethttp "net/http"

	"github.com/steviol/golang-backend-bank-app/internal/config"
	"github.com/steviol/golang-backend-bank-app/internal/db"
	"github.com/steviol/golang-backend-bank-app/internal/delivery/handler"
	"github.com/steviol/golang-backend-bank-app/internal/repository/postgres"
	"github.com/steviol/golang-backend-bank-app/internal/routes"
	"github.com/steviol/golang-backend-bank-app/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.NewPostgresConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	accountRepo := postgres.NewAccountRepository(database)

	accountService := service.NewAccountService(accountRepo)

	accountHandler := handler.NewAccountHandler(accountService)

	router := routes.SetupRoutes(accountHandler)

	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server starting on %s", serverAddr)
	log.Fatal(nethttp.ListenAndServe(serverAddr, router))
}
