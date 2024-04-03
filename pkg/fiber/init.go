package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberInit() (*fiber.App, error) {
	app := fiber.New()

	app.Use(logger.New())

	return app, nil
}
