package main

import (
	"log"

	"example.com/databases/config"

	"example.com/databases/internal/database"
)

func main() {
	// Example config, replace with actual configuration loading
	cfg := config.DatabaseConfig{
		Driver: "sqlite",       // Change to "postgres" to use PostgreSQL
		DSN:    "./example.db", // Replace with PostgreSQL DSN if needed
	}

	db, err := database.NewDatabase(cfg.Driver, cfg.DSN)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	log.Println("Database initialized successfully!")
}
