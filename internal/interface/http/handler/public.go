package handler

import (
	"github.com/gofiber/fiber/v2"
)

type PublicHandler struct{}

func NewPublicHandler() *PublicHandler {
	return &PublicHandler{}
}

func (h *PublicHandler) HandlePublicEndpoint(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello from a public endpoint! You don't need to be authenticated to see this."})
}
