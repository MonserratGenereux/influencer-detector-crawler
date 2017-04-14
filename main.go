package main

import (
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

	facebookCrawler := facebook.NewCrawler()
	log.Fatal(facebookCrawler.Start())
}
