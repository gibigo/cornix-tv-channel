package services

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gibigo/cornix-tv-channel/app/dal"
	"github.com/gibigo/cornix-tv-channel/app/types"
	"github.com/gibigo/cornix-tv-channel/utils"
	"github.com/gibigo/cornix-tv-channel/utils/logging"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// @Summary Create a new strategy
// @Description Create a new strategy for the current user
// @Security BasicAuth
// @Tags strategies
// @Accept  json
// @Produce  json
// @Param user body types.AddStrategy true "Strategy to create"
// @Param channel_id path int true "Channel ID"
// @Success 200 {object} types.Strategy
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {string} string
// @Router /channels/{channel_id}/strategies [post]
func CreateStrategy(c *fiber.Ctx) error {

	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "CreateStrategy",
		"package":  "services",
	})

	// parse the request body
	var newStrategy types.AddStrategy
	if err := utils.ParseBodyAndValidate(c, &newStrategy); err != nil {
		// err = *fiber.Error, only log error message
		logger.Error(err.Message)
		return err
	}

	// check if the channel exists
	var checkChannel types.Channel
	result := dal.FindChannelFromUser(&checkChannel, c.Locals("username"), c.Params("channel"))
	if result.Error != nil {
		logger.Error(result.Error)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
	}
	// return 404 if the channel does not exist or does not belong to the user
	if result.RowsAffected == 0 {
		return utils.NewHTTPError(c, fiber.StatusNotFound, fmt.Sprintf("user %s has no channel with id %s", c.Locals("username"), c.Params("channel")))
	}

	// check if there is already a strategy with that particular symbol for that channel
	result = dal.FindStrategyBySymbol(&dal.Strategy{}, newStrategy.Symbol, c.Params("channel"))
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if result.RowsAffected > 0 {
			return utils.NewHTTPError(c, fiber.StatusConflict, fmt.Errorf("channel %s already has a strategy for %s", c.Params("channel"), newStrategy.Symbol))
		} else {
			logger.Error(result.Error)
			return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
		}
	}

	// make the strategy ready for a database insert
	channel, err := strconv.Atoi(c.Params("channel"))
	if err != nil {
		logger.Error("error converting channel id to uint: ", err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, fmt.Sprintf("error converting channel id to uint: %s", err))
	}

	strategy := convertStrategyStruct(newStrategy, uint(channel))
	if strategy == nil {
		// TODO return error
	}

	if err := dal.CreateStrategy(strategy).Error; err != nil {
		logger.Error(err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}
	return c.JSON(strategy)
}

func convertStrategyStruct(strategy types.AddStrategy, channel uint) *dal.Strategy {

	// allocate memory
	nStrategy := new(dal.Strategy)

	// counter trade policy
	nStrategy.AllowCounter = strategy.AllowCounter

	// define the symbol
	nStrategy.Symbol = strategy.Symbol

	// set the channel id
	nStrategy.ChannelID = channel

	// set target strategy
	if strategy.TargetStrategy != nil {

		nStrategy.IsTargetStrategy = true
		nStrategy.TargetStrategy = new(dal.TargetStrategy)

		// define entries
		if strategy.TargetStrategy.Entries != nil {
			nEntry := make([]*dal.Entry, len(strategy.TargetStrategy.Entries))
			for i, v := range strategy.TargetStrategy.Entries {
				nEntry[i] = new(dal.Entry)
				nEntry[i].Diff = v.Diff
			}
			nStrategy.TargetStrategy.Entries = nEntry
		}

		// define take profits
		if strategy.TargetStrategy.TPs != nil {
			nTP := make([]*dal.TP, len(strategy.TargetStrategy.TPs))
			for i, v := range strategy.TargetStrategy.TPs {
				nTP[i] = new(dal.TP)
				nTP[i].Diff = v.Diff
			}
			nStrategy.TargetStrategy.TPs = nTP
		}

		// define take profits
		if strategy.TargetStrategy.SL != nil {
			nSL := new(dal.SL)
			nSL.Diff = strategy.TargetStrategy.SL.Diff
			nStrategy.TargetStrategy.SL = nSL
		}

		// set breakout
		nStrategy.TargetStrategy.IsBreakout = strategy.TargetStrategy.IsBreakout

		return nStrategy
	}

	// set zone strategy
	if strategy.ZoneStrategy != nil {

		nStrategy.IsZoneStrategy = true
		nStrategy.ZoneStrategy = new(dal.ZoneStrategy)

		// set entry diffs
		nStrategy.ZoneStrategy.EntryStart = strategy.ZoneStrategy.EntryStart
		nStrategy.ZoneStrategy.EntryStop = strategy.ZoneStrategy.EntryStop

		// define take profits
		if strategy.ZoneStrategy.TPs != nil {
			nTP := make([]*dal.TP, len(strategy.ZoneStrategy.TPs))
			for i, v := range strategy.ZoneStrategy.TPs {
				nTP[i] = new(dal.TP)
				nTP[i].Diff = v.Diff
			}
			nStrategy.ZoneStrategy.TPs = nTP
		}

		// define take profits
		if strategy.ZoneStrategy.SL != nil {
			nSL := new(dal.SL)
			nSL.Diff = strategy.ZoneStrategy.SL.Diff
			nStrategy.ZoneStrategy.SL = nSL
		}

		// set breakout
		nStrategy.TargetStrategy.IsBreakout = strategy.TargetStrategy.IsBreakout

		return nStrategy
	}
	return nil
}
