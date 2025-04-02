package database

import (
	"database/sql"
	"fmt"
	"go-server/config"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB


func Connect(cfg *config.Config) error {
	connStr := cfg.GetDSN()
	fmt.Printf("Database connection string: %s\n", connStr)
	log.Printf("Connecting to database: %s:%d/%s", cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	// Try to connect to the database
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL database")
	return nil
}

// Initialize creates necessary tables if they don't exist
func Initialize() error {
	// Create users table
	_, err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL
        )
    `)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	// Create posts table
	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS posts (
            id SERIAL PRIMARY KEY,
            title VARCHAR(200) NOT NULL,
            content TEXT NOT NULL,
            user_id INTEGER REFERENCES users(id)
        )
    `)
	if err != nil {
		return fmt.Errorf("failed to create posts table: %v", err)
	}

	log.Println("Database tables initialized")
	return nil
}
