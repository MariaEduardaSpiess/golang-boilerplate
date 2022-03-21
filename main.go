package main

import (
	"golang-boilerplate/app/handlers"
	"golang-boilerplate/config"
	"golang-boilerplate/infra/database"

	"github.com/gofiber/fiber/v2"
	"github.com/subosito/gotenv"
)

func main() {
	// For local development - load environment variables from .env file
	gotenv.Load()

	// Load environment variables
	config.LoadEnv()

	// Connect to database
	database.Connect()

	// Run API
	router := fiber.New()
	handlers.NewHealthCheckHandler(router)
	router.Listen(config.Env.ApiPort)
}
