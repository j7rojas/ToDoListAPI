package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect() *sql.DB {
	dsn := "user:password@tcp(localhost:3306)/db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	return db
}
