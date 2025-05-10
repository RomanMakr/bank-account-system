package main

import (
	"log"
	"net/http"
	"os"

	"bank-account-system/api"
	"bank-account-system/db"
)

// @title Bank Account System API
// @version 1.0
// @description This is a simple bank account system written in Go.
// @host localhost:8080
// @BasePath /
// @schemes http


func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.DB.Close()

	router := api.RegisterRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running at http://localhost:%s\n", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
