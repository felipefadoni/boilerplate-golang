package main

import (
	"log"
	"os"

	"github.com/felipefadoni/boilerplate-golang/src"
	"github.com/felipefadoni/boilerplate-golang/src/infra/postgres"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgres.InitDatabase()

	r := src.Routes()

	apiPort := os.Getenv("PORT")

	r.Run(":" + apiPort)
}
