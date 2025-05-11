package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB()  error {
	err := godotenv.Load("./.env") // if relative path not work, use absolute path file
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return err
	}

	log.Println("Successfully conntected to database")

	return nil
}
