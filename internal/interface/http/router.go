package http

import (
	"github.com/banggibima/go-fiber-auth0/internal/interface/http/handler"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App                  *fiber.App
	PublicHandler        *handler.PublicHandler
	PrivateHandler       *handler.PrivateHandler
	PrivateScopedHandler *handler.PrivateScopedHandler
	AuthMiddleware       fiber.Handler
}

func (r *Router) AppRouter() {
	r.PublicRouter()
	r.PrivateRouter()
	r.PrivateScopedRouter()
}

func (r *Router) PublicRouter() {
	api := r.App.Group("/api")

	api.Get("/public", r.PublicHandler.HandlePublicEndpoint)
}

func (r *Router) PrivateRouter() {
	r.App.Use(r.AuthMiddleware)

	api := r.App.Group("/api")

	api.Get("/private", r.PrivateHandler.HandlePrivateEndpoint)
}

func (r *Router) PrivateScopedRouter() {
	r.App.Use(r.AuthMiddleware)

	api := r.App.Group("/api")

	api.Get("/private-scoped", r.PrivateScopedHandler.HandlePrivateScopedEndpoint)
}
