package services

import (
	"fmt"

	"github.com/gibigo/cornix-tv-channel/app/types"
	"github.com/gofiber/fiber/v2"
)

func TriggerWebhook(c *fiber.Ctx) error {
	var test *types.TVSignal
	// TODO
	err := c.BodyParser(&test)
	fmt.Println(test)
	return err
}
