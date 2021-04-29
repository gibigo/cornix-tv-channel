package services

import (
	"errors"
	"fmt"

	"github.com/gibigo/cornix-tv-channel/app/dal"
	"github.com/gibigo/cornix-tv-channel/app/types"
	"github.com/gibigo/cornix-tv-channel/utils"
	"github.com/gibigo/cornix-tv-channel/utils/logging"
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
// @Param user body types.AddUser true "User to create"
// @Success 204 {string} string
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
	var user *types.GetUser
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

	logger.Info("created telegram channel: ", channel.Telegram)

	// replace noContent with channel struct
	c.Status(fiber.StatusNoContent)
	return nil
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
	var channels []*types.Channel
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
