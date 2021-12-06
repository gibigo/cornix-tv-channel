package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func swaggerRoutes(api fiber.Router) {
	api.Get("/swagger/*", swagger.Handler)
}
