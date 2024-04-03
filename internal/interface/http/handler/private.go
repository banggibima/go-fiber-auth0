package handler

import (
	"fmt"

	"github.com/banggibima/go-fiber-auth0/config"
	"github.com/gofiber/fiber/v2"
)

type PrivateHandler struct {
	cfg *config.Config
}

func NewPrivateHandler(cfg *config.Config) *PrivateHandler {
	return &PrivateHandler{
		cfg: cfg,
	}
}

func (h *PrivateHandler) HandlePrivateEndpoint(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Credentials", "true")
	c.Set("Access-Control-Allow-Origin", fmt.Sprintf("http://localhost:%d", h.cfg.App.Port))
	c.Set("Access-Control-Allow-Headers", "Authorization")

	c.Set("Content-Type", "application/json")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello from a private endpoint! You need to be authenticated to see this."})
}
