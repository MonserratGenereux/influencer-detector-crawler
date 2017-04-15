package main

import (
	"cassandra"
	"crawler/facebook"
	"log"

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

	facebookCrawler := facebook.NewCrawler("683165801724841")
	log.Fatal(facebookCrawler.Start())
}
