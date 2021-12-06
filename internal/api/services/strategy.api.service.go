package services

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gibigo/cornix-tv-channel/internal/api/dal"
	"github.com/gibigo/cornix-tv-channel/internal/api/types"
	"github.com/gibigo/cornix-tv-channel/internal/utils"
	"github.com/gibigo/cornix-tv-channel/internal/utils/logging"
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
	result := dal.FindChannelFromUser(&types.Channel{}, c.Locals("username"), c.Params("channel"))
	if result.Error != nil {
		logger.Error(result.Error)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
	}
	// return 404 if the channel does not exist or does not belong to the user
	if result.RowsAffected == 0 {
		return utils.NewHTTPError(c, fiber.StatusNotFound, fmt.Sprintf("user %s has no channel with id %s", c.Locals("username"), c.Params("channel")))
	}

	// check if there is already a strategy with that particular symbol for that channel
	result = dal.FindStrategyBySymbolSimple(&dal.Strategy{}, newStrategy.Symbol, c.Params("channel"))
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
		err := fmt.Errorf("error while converting the struct")
		logger.Error(err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}

	if err := dal.CreateStrategy(strategy).Error; err != nil {
		logger.Error(err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}

	var returnStrategy types.Strategy
	if err := dal.FindStrategyByID(&returnStrategy, strategy.ChannelID, strategy.ID).Error; err != nil {
		logger.Error("error while getting strategy after creating it", err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}
	return c.JSON(returnStrategy)
}

// helper function to convert structs
func convertStrategyStruct(strategy types.AddStrategy, channel uint) *dal.Strategy {
	// allocate memory
	nStrategy := new(dal.Strategy)
	// counter trade policy
	nStrategy.AllowCounter = strategy.AllowCounter
	// define the symbol
	nStrategy.Symbol = strategy.Symbol
	// set the leverage
	nStrategy.Leverage = strategy.Leverage
	// set the channel id
	nStrategy.ChannelID = channel
	// set target strategy
	if strategy.TargetStrategy != nil {
		// allocate memory
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
		// allocate memory
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
		nStrategy.ZoneStrategy.IsBreakout = strategy.ZoneStrategy.IsBreakout
		return nStrategy
	}
	return nil
}

// @Summary Get all strategies
// @Description Get all strategies for a particular channel
// @Security BasicAuth
// @Tags strategies
// @Accept  json
// @Produce  json
// @Param channel_id path int true "Channel ID"
// @Success 200 {array} types.Strategy
// @Failure 401 {string} string
// @Failure 404 {object} utils.HTTPError
// @Router /channels/{channel_id}/strategies [get]
func GetStrategies(c *fiber.Ctx) error {
	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "GetStrategies",
		"package":  "services",
	})

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

	var strategies []types.Strategy
	result = dal.FindAllStrategiesFromChannel(&strategies, checkChannel.ID)
	if err := result.Error; err != nil {
		logger.Error(err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
	}

	// return 404 if user has no channels configured
	if result.RowsAffected == 0 {
		return utils.NewHTTPError(c, fiber.StatusNotFound, fmt.Sprintf("no strategies found for channel with id %s", c.Params("channel")))
	}

	return c.JSON(strategies)
}

// @Summary Get a strategy
// @Description Get a strategy by the channel id and the symbol
// @Security BasicAuth
// @Tags strategies
// @Accept  json
// @Produce  json
// @Param channel_id path int true "Channel ID"
// @Param strategy_symbol path string true "Strategy Symbol, use 'all' for the default strategy"
// @Success 200 {object} types.Strategy
// @Failure 401 {string} string
// @Failure 404 {object} utils.HTTPError
// @Router /channels/{channel_id}/strategies/{strategy_symbol} [get]
func GetStrategy(c *fiber.Ctx) error {
	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "GetStrategy",
		"package":  "services",
	})

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

	var strategies types.Strategy
	result = dal.FindStrategyBySymbol(&strategies, checkChannel.ID, c.Params("symbol"))
	if err := result.Error; err != nil {
		logger.Error(err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
	}

	// return 404 if user has no channels configured
	if result.RowsAffected == 0 {
		return utils.NewHTTPError(c, fiber.StatusNotFound, fmt.Sprintf("no strategy found for channel with id %s and symbol %s", c.Params("channel"), c.Params("symbol")))
	}

	return c.JSON(strategies)
}

// @Summary Delete a strategy
// @Description Delete a strategy for a particular symbol
// @Security BasicAuth
// @Tags strategies
// @Accept  json
// @Produce  json
// @Success 204 {string} string
// @Failure 401 {string} string
// @Failure 404 {object} utils.HTTPError
// @Param channel_id path int true "Channel ID"
// @Param strategy_symbol path string true "Strategy Symbol, use 'all' for the default strategy"
// @Router /channels/{channel_id}/strategies/{strategy_symbol} [delete]
func DeleteStrategy(c *fiber.Ctx) error {
	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "DeleteStrategy",
		"package":  "services",
	})

	var channel types.Channel
	result := dal.FindChannelFromUser(&channel, c.Locals("username"), c.Params("channel"))

	if result.Error != nil {
		logger.Error(result.Error)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
	}

	// return 404 if the channel does not exist or does not belong to the user
	if result.RowsAffected == 0 {
		return utils.NewHTTPError(c, fiber.StatusNotFound, fmt.Sprintf("user %s has no channel with id %s", c.Locals("username"), c.Params("channel")))
	}

	// check if there is a strategy with that particular symbol for that channel
	result = dal.FindStrategyBySymbolSimple(&dal.Strategy{}, c.Params("symbol"), c.Params("channel"))
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return utils.NewHTTPError(c, fiber.StatusNotFound, fmt.Errorf("channel %s has no strategy for %s", c.Params("channel"), c.Params("symbol")))
	} else if result.Error != nil {
		logger.Error(result.Error)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
	}

	// delete the channel
	if err := dal.DeleteStrategy(c.Params("channel"), c.Params("symbol")).Error; err != nil {
		logger.Error(err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}

	logger.Infof("deleted strategy %s from channel %s from user %s", c.Params("symbol"), c.Params("channel"), c.Locals("username"))

	c.Status(fiber.StatusNoContent)
	return nil
}
