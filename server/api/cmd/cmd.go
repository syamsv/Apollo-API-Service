package cmd

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/syamsv/apollo/api/router"
	"github.com/syamsv/apollo/config"
)

func StartServer() {
	app := fiber.New(fiber.Config{
		// Prefork:       true, // Enable Prefork in prod only
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Apollo Server",
		AppName:       "App Name",
	})
	app.Use(cors.New())
	api := app.Group("/api")
	router.MountAuthRoute(api)

	if err := app.Listen(config.SERVER_PORT); err != nil {
		log.Fatal("[SERVER] Error starting server: ", err)
	}
}
