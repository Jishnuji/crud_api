package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func SetupDB() {
	config, err := LoadConfig("/app/config_db.yml")
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		config.DB.User,
		config.DB.Password,
		config.DB.DBName,
		config.DB.DBHost,
		config.DB.DBPort,
		config.DB.SSLMode,
	)

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB.SetMaxOpenConns(config.DB.MaxOpenConnections)
	DB.SetMaxIdleConns(config.DB.MaxIdleConnections)
	DB.SetConnMaxLifetime(time.Duration(config.DB.ConnMaxLifetimeMinutes) * time.Minute)

	query := `CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		firstname TEXT NOT NULL,
		lastname TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		age INTEGER NOT NULL,
		created TIMESTAMPTZ NOT NULL
	)`
	if _, err := DB.Exec(query); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
}
