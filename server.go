package main

import (
	"api"
	"cassandra"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to Cassandra
	session := cassandra.GetSession()
	defer session.Close()

	// Run server.
	// Address = HOST:PORT
	address := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	log.Println("Analytics client API running on ", address)
	http.ListenAndServe(address, api.GetRoutes())
}
