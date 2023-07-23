package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syamsv/apollo/api/handler"
)

func MountAuthRoute(auth fiber.Router) {
	auth.Post("/login", handler.AuthLogin)
	auth.Post("/register", handler.AuthRegister)
	auth.Get("/verify", handler.AuthTokenVerify)
	auth.Post("/refresh", handler.AuthTokenRefresh)
}
