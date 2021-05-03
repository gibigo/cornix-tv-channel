package routes

import (
	s "github.com/gibigo/cornix-tv-channel/app/services"
	"github.com/gibigo/cornix-tv-channel/utils/middleware"
	"github.com/gofiber/fiber/v2"
)

func apiRoutes(api fiber.Router) {
	r := api.Group("/api")
	registerUsers(r)
	registerChannels(r)
	registerStrategies(r)
}

// register user endpoints
func registerUsers(api fiber.Router) {
	v1 := api.Group("/v1")
	users := v1.Group("/users")

	users.Post("/", s.CreateUser)
	users.Use(middleware.BasicAuth).Get("/", s.GetUser)
	users.Use(middleware.BasicAuth).Delete("/", s.DeleteUser)
	users.Use(middleware.BasicAuth).Put("/", s.UpdateUser)
}

// register strategy endpoints
func registerStrategies(api fiber.Router) {
	v1 := api.Group("/v1")
	str := v1.Group("/channels/:channel/strategies").Use(middleware.BasicAuth)

	//str.Get("/", s.GetStrategies)
	str.Post("/", s.CreateStrategy)
	//str.Delete("/", s.DeleteStrategy)
}

// register channel endpoints
func registerChannels(api fiber.Router) {
	v1 := api.Group("/v1")
	ch := v1.Group("/channels").Use(middleware.BasicAuth)

	ch.Post("/", s.CreateChannel)
	ch.Get("/", s.GetChannels)
	ch.Get("/:channel", s.GetChannel)
	ch.Delete("/:channel", s.DeleteChannel)
	ch.Put("/:channel", s.UpdateChannel)
}
