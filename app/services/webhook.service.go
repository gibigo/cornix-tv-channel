package services

import (
	"errors"
	"fmt"

	"github.com/gibigo/cornix-tv-channel/app/dal"
	"github.com/gibigo/cornix-tv-channel/app/telegram"
	"github.com/gibigo/cornix-tv-channel/app/types"
	"github.com/gibigo/cornix-tv-channel/utils"
	"github.com/gibigo/cornix-tv-channel/utils/logging"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// add swagger docs

func TriggerWebhook(c *fiber.Ctx) error {

	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "TriggerWebhook",
		"package":  "services",
	})

	// parse signal
	var newSignal types.TVSignal
	if err := utils.ParseBodyAndValidate(c, &newSignal); err != nil {
		// err = *fiber.Error, only log error message
		logger.Warn(err.Message)
		return err
	}

	// get the channelID by the telegram ID
	var channel dal.Channel
	result := dal.FindChannelByTelegramId(&channel, newSignal.ChannelID)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) && result.Error != nil {
		logger.Error(result.Error)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
	}
	if result.RowsAffected == 0 {
		return utils.NewHTTPError(c, fiber.StatusNotFound, fmt.Sprintf("channel with telegram id %d not found", newSignal.ChannelID))
	}

	// check if the user and uuid match
	result = dal.FindUserByUUID(&dal.User{}, newSignal.User, newSignal.UUID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return utils.NewHTTPError(c, fiber.StatusUnauthorized, "name and uuid don't match")
	} else if result.Error != nil {
		logger.Error(result.Error)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
	}

	// set direction to long if the string is empty
	if newSignal.Direction == "" {
		newSignal.Direction = "long"
	}

	// check if trade for same direction is open
	var prevSignal dal.TVSignal
	result = dal.FindSignalBySymbol(&prevSignal, channel.ID, newSignal.Symbol)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) && result.Error != nil {
		logger.Error(result.Error)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
	}

	// get the strategy for the signal
	var strategy dal.Strategy
	result = dal.FindStrategyBySymbol(&strategy, channel.ID, newSignal.Symbol)
	if result.Error != nil {
		logger.Error(result.Error)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
	}
	if result.RowsAffected == 0 {
		// get the default strategy
		result = dal.FindDefaultStrategy(&strategy, channel.ID)
		if result.Error != nil {
			logger.Error(result.Error)
			return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
		}
		if result.RowsAffected == 0 {
			return utils.NewHTTPError(c, fiber.StatusNotFound, fmt.Sprintf("channel with telegram id %d has no matching strategy", newSignal.ChannelID))
		}
	}

	// allocate memory for the new message and counter
	var message string
	var counter bool

	// check if there is an open trade
	if prevSignal.Symbol != "" {
		// if there is an open trade, check if the new signal may overwrite it
		if prevSignal.Direction == "long" && newSignal.Direction == "short" && strategy.AllowCounter {
			counter = true
			// generate signal message
			message = generateSignal(newSignal.Symbol, newSignal.Price, strategy)
		} else {
			return utils.NewHTTPError(c, fiber.StatusNotAcceptable, fmt.Sprintf("counter trades aren't allowed for %s", strategy.Symbol))
		}
	} else {
		// generate signal message
		message = generateSignal(newSignal.Symbol, newSignal.Price, strategy)
	}

	// send message(s) to telegram channel
	if counter {
		// TODO send cancel message to telegram
		telegram.Bot.SendMessage(channel.Telegram, "counter", nil)
	}
	// TODO send telegram message
	fmt.Println(message)

	logger.Info(newSignal)
	return c.JSON(newSignal)
}

// WIP
func generateSignal(symbol string, price float64, strategy dal.Strategy) string {
	genEntry(price, strategy)
	genTP(price, strategy)
	genSL(price, strategy)
	return ""
}

func genEntry(price float64, strategy dal.Strategy) string {
	return ""
}

func genTP(price float64, strategy dal.Strategy) string {
	return ""
}

func genSL(price float64, strategy dal.Strategy) string {
	return ""
}
