package main

import (
	"fmt"
	"go-server/config"
	"go-server/database"
	"go-server/handlers"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}


	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := database.Initialize(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/404", handlers.HandleNotFound)
	http.HandleFunc("/users", handlers.HandleCreateUser)
	http.HandleFunc("/users/get", handlers.HandleGetUser)

	fmt.Println("Server initialized with PostgreSQL database")
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Starting server on %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatal(err)
	}
}
