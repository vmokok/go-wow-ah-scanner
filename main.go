package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	apiClientID := os.Getenv("BNET_API_CLIENT_ID")
	apiClientSecret := os.Getenv("BNET_API_CLIENT_SECRET")

	fmt.Println(apiClientID, apiClientSecret)
}
