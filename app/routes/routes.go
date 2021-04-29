package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupRoutes(app fiber.Router) {
	app.Use(recover.New())

	apiRoutes(app)
	webhookRoutes(app)
	swaggerRoutes(app)
}
