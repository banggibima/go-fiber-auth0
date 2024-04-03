package handler

import (
	"fmt"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/banggibima/go-fiber-auth0/config"
	"github.com/banggibima/go-fiber-auth0/internal/interface/http/middleware"
	"github.com/gofiber/fiber/v2"
)

type PrivateScopedHandler struct {
	cfg *config.Config
}

func NewPrivateScopedHandler(cfg *config.Config) *PrivateScopedHandler {
	return &PrivateScopedHandler{
		cfg: cfg,
	}
}

func (h *PrivateScopedHandler) HandlePrivateScopedEndpoint(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Credentials", "true")
	c.Set("Access-Control-Allow-Origin", fmt.Sprintf("http://localhost:%d", h.cfg.App.Port))
	c.Set("Access-Control-Allow-Headers", "Authorization")

	c.Set("Content-Type", "application/json")

	token := c.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	claims := token.CustomClaims.(*middleware.CustomClaims)
	if !claims.HasScope("read:messages") {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Insufficient scope."})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello from a private scoped endpoint! You need to be authenticated and have a scope of 'read:messages' to see this."})
}
