package database

import (
  "database/sql"
	"fmt"
	"os"
  _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() (error) {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

  db, err := sql.Open("postgres", connStr)
	if err != nil {
    return fmt.Errorf("Error opening database connection: %v", err)
	}

  if err := db.Ping(); err != nil {
    return fmt.Errorf("Error pinging database: %v", err)
  }

  DB = db

  return nil
}


