package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gibigo/cornix-tv-channel/internal/api/dal"
	"github.com/gibigo/cornix-tv-channel/internal/api/types"
	"github.com/gibigo/cornix-tv-channel/internal/telegram"
	"github.com/gibigo/cornix-tv-channel/internal/utils"
	"github.com/gibigo/cornix-tv-channel/internal/utils/logging"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// TODO add swagger docs

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
	} else if !strings.EqualFold(newSignal.Direction, "long") && !strings.EqualFold(newSignal.Direction, "short") {
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

	// allocate memory for the new message and counterPolicy
	var message string
	var counterPolicy bool

	// check if there is an open trade
	if prevSignal.Symbol != "" {
		// if there is an open trade, check if the new signal may overwrite it
		if !strings.EqualFold(prevSignal.Direction, newSignal.Direction) && strategy.AllowCounter {
			counterPolicy = true
		} else if strings.EqualFold(prevSignal.Direction, newSignal.Direction) {
			return utils.NewHTTPError(c, fiber.StatusConflict, fmt.Sprintf("there is already an open trade for %s", newSignal.Symbol))
		} else {
			return utils.NewHTTPError(c, fiber.StatusNotAcceptable, fmt.Sprintf("counter trades aren't allowed for %s", newSignal.Symbol))
		}
	}

	// generate signal message
	message = genSignalForTelegram(newSignal, strategy)

	// send message(s) to telegram channel
	if counterPolicy {
		if err := sendCancelMessage(newSignal.Symbol, channel.Telegram); err != nil {
			logger.Errorf("error while sending cancel message: %s", err)
			return utils.NewHTTPError(c, fiber.StatusInternalServerError, fmt.Sprintf("error while sending cancel message: %s", err))
		}
		if err := dal.DeleteSignal(channel.ID, newSignal.Symbol).Error; err != nil {
			logger.Errorf("error while removing canceled signal from database: %s", err)
			return utils.NewHTTPError(c, fiber.StatusInternalServerError, fmt.Sprintf("error while removing canceled signal from database: %s", err))
		}
	}

	// send telegram message
	if err := sendSignalMessage(message, channel.Telegram); err != nil {
		logger.Errorf("error while sending signal message: %s", err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, fmt.Sprintf("error while sending signal message: %s", err))
	}

	// insert signal into database
	if err := dal.CreateSignal(genSignalForDatabase(newSignal, channel)).Error; err != nil {
		logger.Errorf("error while inserting new signal: %s", err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}

	return c.JSON(newSignal)
}

func sendSignalMessage(m string, channelID int64) error {
	_, err := telegram.Bot.SendMessage(channelID, m, nil)
	return err
}

func sendCancelMessage(symbol string, channelID int64) error {
	_, err := telegram.Bot.SendMessage(channelID, fmt.Sprintf("❌ Cancel #%s", symbol), nil)
	return err
}

func genSignalForDatabase(s types.TVSignal, c dal.Channel) *dal.TVSignal {
	var newSignal dal.TVSignal
	newSignal.Price = s.Price
	newSignal.Symbol = s.Symbol
	newSignal.Direction = s.Direction
	newSignal.Exchange = s.Exchange
	newSignal.ChannelID = c.ID
	return &newSignal
}

// WIP
func genSignalForTelegram(s types.TVSignal, strategy dal.Strategy) string {
	msg := fmt.Sprintf("⚡️⚡️ %s ⚡️⚡️\nExchange: %s\nDirection: %s\n", s.Symbol, s.Exchange, s.Direction)
	if strategy.Leverage != 0 {
		msg += fmt.Sprintf("Leverage: %dx\n", strategy.Leverage)
	}
	msg += genEntry(s.Price, s.Direction, strategy)
	msg += genTP(s.Price, s.Direction, strategy)
	msg += genSL(s.Price, s.Direction, strategy)
	return msg
}

func genEntry(price float64, direction string, strategy dal.Strategy) string {
	var msg string
	if strategy.TargetStrategy != nil {
		msg = "\nEntry orders:\n"

		if strings.EqualFold(direction, "long") {
			for i, v := range strategy.TargetStrategy.Entries {
				t := price + (price * (v.Diff / 100))
				msg += fmt.Sprintf("%d) %f\n", i+1, t)
			}
		} else if strings.EqualFold(direction, "short") {
			for i, v := range strategy.TargetStrategy.Entries {
				t := price - (price * (v.Diff / 100))
				msg += fmt.Sprintf("%d) %f\n", i+1, t)
			}
		}
	} else if strategy.ZoneStrategy != nil {
		if strategy.ZoneStrategy.IsBreakout {
			if strings.EqualFold(direction, "long") {
				t1 := price + (price * (strategy.ZoneStrategy.EntryStart / 100))
				t2 := price + (price * (strategy.ZoneStrategy.EntryStop / 100))
				msg = fmt.Sprintf("\nEntry zone above %f-%f\n", t1, t2)
			} else if strings.EqualFold(direction, "short") {
				t1 := price - (price * (strategy.ZoneStrategy.EntryStart / 100))
				t2 := price - (price * (strategy.ZoneStrategy.EntryStop / 100))
				msg = fmt.Sprintf("\nEntry zone below %f-%f\n", t1, t2)
			}
		} else {
			if strings.EqualFold(direction, "long") {
				t1 := price + (price * (strategy.ZoneStrategy.EntryStart / 100))
				t2 := price + (price * (strategy.ZoneStrategy.EntryStop / 100))
				msg = fmt.Sprintf("\nEntry zone %f-%f\n", t1, t2)
			} else if strings.EqualFold(direction, "short") {
				t1 := price - (price * (strategy.ZoneStrategy.EntryStart / 100))
				t2 := price - (price * (strategy.ZoneStrategy.EntryStop / 100))
				msg = fmt.Sprintf("\nEntry zone %f-%f\n", t1, t2)
			}
		}
	}
	return msg
}

func genTP(price float64, direction string, strategy dal.Strategy) string {
	var msg string
	if strategy.TargetStrategy != nil {
		msg = "\nTake-Profit orders:\n"

		if strings.EqualFold(direction, "long") {
			for i, v := range strategy.TargetStrategy.TPs {
				t := price + (price * (v.Diff / 100))
				msg += fmt.Sprintf("%d) %f\n", i+1, t)
			}
		} else if strings.EqualFold(direction, "short") {
			for i, v := range strategy.TargetStrategy.TPs {
				t := price - (price * (v.Diff / 100))
				msg += fmt.Sprintf("%d) %f\n", i+1, t)
			}
		}
	} else if strategy.ZoneStrategy != nil {
		msg = "\nTake-Profit orders:\n"

		if strings.EqualFold(direction, "long") {
			for i, v := range strategy.ZoneStrategy.TPs {
				t := price + (price * (v.Diff / 100))
				msg += fmt.Sprintf("%d) %f\n", i+1, t)
			}
		} else if strings.EqualFold(direction, "short") {
			for i, v := range strategy.ZoneStrategy.TPs {
				t := price - (price * (v.Diff / 100))
				msg += fmt.Sprintf("%d) %f\n", i+1, t)
			}
		}
	}
	return msg
}

func genSL(price float64, direction string, strategy dal.Strategy) string {
	var msg string
	if strategy.TargetStrategy != nil {
		msg = "\nStop-Loss orders:\n"
		if strings.EqualFold(direction, "long") {
			if strategy.TargetStrategy.SL != nil {
				t := price - (price * (strategy.TargetStrategy.SL.Diff / 100))
				msg += fmt.Sprintf("1) %f\n", t)
			}
		} else if strings.EqualFold(direction, "short") {
			if strategy.TargetStrategy.SL != nil {
				t := price + (price * (strategy.TargetStrategy.SL.Diff / 100))
				msg += fmt.Sprintf("1) %f\n", t)
			}
		}

	} else if strategy.ZoneStrategy != nil {
		msg = "\nStop-Loss orders:\n"
		if strings.EqualFold(direction, "long") {
			if strategy.ZoneStrategy.SL != nil {
				t := price - (price * (strategy.ZoneStrategy.SL.Diff / 100))
				msg += fmt.Sprintf("1) %f\n", t)
			}
		} else if strings.EqualFold(direction, "short") {
			if strategy.ZoneStrategy.SL != nil {
				t := price + (price * (strategy.ZoneStrategy.SL.Diff / 100))
				msg += fmt.Sprintf("1) %f\n", t)
			}
		}
	}
	return msg
}
