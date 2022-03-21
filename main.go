package main

import (
	"golang-boilerplate/config"
	"golang-boilerplate/infra/database"

	"github.com/subosito/gotenv"
)

func main() {
	// For local development - load environment variables from .env file
	gotenv.Load()

	// Load environment variables
	config.LoadEnv()

	// Connect to database
	database.Connect()
}
