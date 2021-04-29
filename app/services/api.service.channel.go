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

	fmt.Println(string(c.Body()))

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

	c.Status(fiber.StatusNoContent)
	return nil
}
