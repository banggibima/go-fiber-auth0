package http

import (
	"fmt"
	"log"

	"github.com/banggibima/go-fiber-auth0/config"
	"github.com/banggibima/go-fiber-auth0/internal/interface/http/handler"
	"github.com/banggibima/go-fiber-auth0/internal/interface/http/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type Server struct {
	Fiber  *fiber.App
	Config *config.Config
}

func AppServer(s *Server) {
	publicHandler := handler.NewPublicHandler()
	privateHandler := handler.NewPrivateHandler(s.Config)
	privateScopedHandler := handler.NewPrivateScopedHandler(s.Config)

	authMiddleware := adaptor.HTTPMiddleware(middleware.EnsureValidToken(s.Config))

	r := Router{
		App:                  s.Fiber,
		PublicHandler:        publicHandler,
		PrivateHandler:       privateHandler,
		PrivateScopedHandler: privateScopedHandler,
		AuthMiddleware:       authMiddleware,
	}

	r.AppRouter()

	port := s.Config.App.Port

	if err := s.Fiber.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal(err)
	}
}
