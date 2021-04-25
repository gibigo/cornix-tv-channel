package routes

import (
	s "github.com/gibigo/cornix-tv-channel/app/services"
	"github.com/gofiber/fiber/v2"
)

func webhookRoutes(api fiber.Router) {
	r := api.Group("/webhook")
	registerWebhook(r)
}

func registerWebhook(api fiber.Router) {
	v1 := api.Group("/v1")
	v1.Post("/", s.TriggerWebhook)
}
