package main

import (
	"log"
	"net/http"
	"server/test-api/internal"
	"server/test-api/internal/database"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func handleRequests() {
	internal.HandleRoutes()
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	return
}

func main() {
	loadEnv()
	database.ConnectDb()
	handleRequests()
	log.Println("Database connected")
}
