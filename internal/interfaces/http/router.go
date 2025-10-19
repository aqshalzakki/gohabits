package http

import (
	auth2 "gohabits/internal/application/auth"
	"gohabits/internal/infra"
	"gohabits/internal/interfaces/http/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupRouter inisialisasi semua route aplikasi
func SetupRouter(app *fiber.App, container *infra.Container) {
	api := app.Group("/api")

	jwtManager := auth2.NewJWTManager(container.Config.JWTSecret)

	// Group untuk auth
	authHandler := handlers.NewAuthHandler(container, jwtManager)
	auth := api.Group("/auth")
	{
		auth.Post("/register", authHandler.Register)
		auth.Post("/login", authHandler.Login)
		protected := auth.Use(auth2.JWTMiddleware(container.Config))
		protected.Get("/me", authHandler.Me)
	}
}
