package main

import (
	"log"
	"os"

	"github.com/felipefadoni/boilerplate-golang/src"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := src.Routes()

	apiPort := os.Getenv("PORT")

	r.Run(":" + apiPort)
}
