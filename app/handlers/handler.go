package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type HealthCheckHandler struct {
}

func NewHealthCheckHandler(router *fiber.App) {
	handler := HealthCheckHandler{}
	router.Get("/", handler.HealthCheck)
}

func (handler *HealthCheckHandler) HealthCheck(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
