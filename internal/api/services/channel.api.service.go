package services

import (
	"errors"
	"fmt"

	"github.com/gibigo/cornix-tv-channel/internal/api/dal"
	"github.com/gibigo/cornix-tv-channel/internal/api/types"
	"github.com/gibigo/cornix-tv-channel/internal/utils"
	"github.com/gibigo/cornix-tv-channel/internal/utils/logging"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// @Summary Create a channel
// @Description Create a new channel
// @Security BasicAuth
// @Tags channels
// @Accept  json
// @Produce  json
// @Param channel body types.AddChannel true "Channel to create"
// @Success 200 {object} types.Channel
// @Failure 401 {string} string
// @Failure 409 {object} utils.HTTPError
// @Router /channels [post]
func CreateChannel(c *fiber.Ctx) error {
	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "CreateChannel",
		"package":  "services",
	})

	// parse the request body
	var newChannel types.AddChannel
	if err := utils.ParseBodyAndValidate(c, &newChannel); err != nil {
		// err = *fiber.Error, only log error message
		logger.Error(err.Message)
		return err
	}

	// check if there is already a channel with that particular telegram id
	result := dal.FindChannelByTelegramId(&dal.Channel{}, newChannel.TelegramID)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if result.RowsAffected > 0 {
			return utils.NewHTTPError(c, fiber.StatusConflict, fmt.Errorf("channel %d already exists", newChannel.TelegramID))
		} else {
			logger.Error(result.Error)
			return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
		}
	}

	// get the user id of the current user
	var user types.GetUserWithID
	if err := dal.FindUserByName(&user, c.Locals("username")).Error; err != nil {
		logger.Error(err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}

	// create the new channel
	channel := &dal.Channel{
		Telegram: newChannel.TelegramID,
		UserID:   user.ID,
	}
	if err := dal.CreateChannel(channel).Error; err != nil {
		logger.Error(err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}

	logger.Infof("user %s created telegram channel: %d", c.Locals("username"), channel.Telegram)

	// fetch the newly created channel
	var returnChannel types.Channel
	if err := dal.FindChannelByTelegramId(&returnChannel, channel.Telegram).Error; err != nil {
		logger.Errorf("error while fetching the new channel: %s", err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}

	return c.JSON(returnChannel)
}

// @Summary Get all channels
// @Description Get all channels of the current user
// @Security BasicAuth
// @Tags channels
// @Accept  json
// @Produce  json
// @Success 200 {array} types.Channel
// @Failure 401 {string} string
// @Failure 404 {object} utils.HTTPError
// @Router /channels [get]
func GetChannels(c *fiber.Ctx) error {
	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "GetChannels",
		"package":  "services",
	})

	// get all channels from the user
	var channels []types.Channel
	result := dal.FindAllChannelsFromUser(&channels, c.Locals("username"))
	if err := result.Error; err != nil {
		logger.Error(err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
	}

	// return 404 if user has no channels configured
	if result.RowsAffected == 0 {
		return utils.NewHTTPError(c, fiber.StatusNotFound, fmt.Sprintf("no channels found for user %s", c.Locals("username")))
	}

	return c.JSON(channels)
}

// @Summary Get a spectific channel
// @Description Get a spectific channel
// @Security BasicAuth
// @Tags channels
// @Accept  json
// @Produce  json
// @Success 200 {object} types.Channel
// @Failure 401 {string} string
// @Failure 404 {object} utils.HTTPError
// @Param channel_id path int true "Channel ID"
// @Router /channels/{channel_id} [get]
func GetChannel(c *fiber.Ctx) error {
	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "GetChannel",
		"package":  "services",
	})

	var channel types.Channel
	result := dal.FindChannelFromUser(&channel, c.Locals("username"), c.Params("channel"))

	if result.Error != nil {
		logger.Error(result.Error)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
	}

	// return 404 if the channel does not exist
	if result.RowsAffected == 0 {
		return utils.NewHTTPError(c, fiber.StatusNotFound, fmt.Sprintf("channel with id %s not found", c.Params("channel")))
	}

	return c.JSON(channel)
}

// @Summary Delete a channel
// @Description Delete a channel and all related strategies
// @Security BasicAuth
// @Tags channels
// @Accept  json
// @Produce  json
// @Success 204 {string} string
// @Failure 401 {string} string
// @Failure 404 {object} utils.HTTPError
// @Param channel_id path int true "Channel ID"
// @Router /channels/{channel_id} [delete]
func DeleteChannel(c *fiber.Ctx) error {
	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "DeleteChannel",
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

	// delete the channel
	if err := dal.DeleteChannelFromUser(c.Locals("username"), c.Params("channel")).Error; err != nil {
		logger.Error(err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}

	logger.Infof("deleted channel %s from user %s", c.Params("channel"), c.Locals("username"))

	c.Status(fiber.StatusNoContent)
	return nil
}

// @Summary Update a channel
// @Description Change the telegram id of a channel and keep all related strategies
// @Security BasicAuth
// @Tags channels
// @Accept  json
// @Produce  json
// @Success 200 {object} types.Channel
// @Failure 401 {string} string
// @Failure 404 {object} utils.HTTPError
// @Param channel_id path int true "Channel ID"
// @Param channel body types.UpdateChannel true "Channel to create"
// @Router /channels/{channel_id} [put]
func UpdateChannel(c *fiber.Ctx) error {
	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "UpdateChannel",
		"package":  "services",
	})

	// parse the body to a new struct
	var channelUpdate types.UpdateChannel
	if err := utils.ParseBodyAndValidate(c, &channelUpdate); err != nil {
		// err = *fiber.Error, only log error message
		logger.Error(err.Message)
		return err
	}

	result := dal.ChangeChannelTelegram(c.Locals("username"), channelUpdate.Telegram, c.Params("channel"))

	if result.Error != nil {
		logger.Error(result.Error)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
	}

	if result.RowsAffected == 0 {
		return utils.NewHTTPError(c, fiber.StatusNotFound, fmt.Sprintf("user %s has no channel with id %s", c.Locals("username"), c.Params("channel")))
	}
	var returnChannel types.Channel
	if err := dal.FindChannelByTelegramId(&returnChannel, channelUpdate.Telegram).Error; err != nil {
		logger.Errorf("error while fetching the updated channel: %s", err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}

	return c.JSON(returnChannel)
}
