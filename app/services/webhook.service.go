package services

import (
	"github.com/gibigo/cornix-tv-channel/app/types"
	"github.com/gibigo/cornix-tv-channel/utils"
	"github.com/gibigo/cornix-tv-channel/utils/logging"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// add swagger docs

func TriggerWebhook(c *fiber.Ctx) error {

	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "TriggerWebhook",
		"package":  "services",
	})

	// parse signal
	var newSignal *types.TVSignal
	if err := utils.ParseBodyAndValidate(c, &newSignal); err != nil {
		// err = *fiber.Error, only log error message
		logger.Error(err.Message)
		return err
	}

	// set direction to long if the string is empty
	if newSignal.Direction == "" {
		newSignal.Direction = "long"
	}

	// check if trade for same direction is open

	// check if allowCounter is set to true

	// generate signal message

	// send signal to telegram channel

	err := c.BodyParser(&newSignal)
	logger.Info(newSignal)
	return err
}
