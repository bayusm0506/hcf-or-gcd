package main

import (
	"log"

	"github.com/bayusm0506/hcf-or-gcd/app"
	"github.com/bayusm0506/hcf-or-gcd/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := config.GetConfig()

	app := &app.App{}
	app.Initialize()

	app.Run(config)

}
