package main

import (
	"fmt"
	"gohabits/internal/infra"
	"gohabits/internal/interfaces/http"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	container := infra.NewContainer()

	//load env
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Health check OK",
		})
	})

	http.SetupRouter(app, container)

	port := container.Config.Port
	log.Printf("🚀 Server running on port %s", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
