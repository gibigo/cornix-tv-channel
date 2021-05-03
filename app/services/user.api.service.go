package services

import (
	"errors"
	"fmt"

	"github.com/gibigo/cornix-tv-channel/app/dal"
	"github.com/gibigo/cornix-tv-channel/app/types"
	"github.com/gibigo/cornix-tv-channel/config"
	"github.com/gibigo/cornix-tv-channel/utils"
	"github.com/gibigo/cornix-tv-channel/utils/logging"
	"github.com/gibigo/cornix-tv-channel/utils/password"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// @Summary Create a user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body types.AddUser true "User to create"
// @Success 200 {object} types.GetUser
// @Failure 401 {string} string
// @Failure 409 {object} utils.HTTPError
// @Failure 501 {object} utils.HTTPError "if user registration is disabled on the server"
// @Router /users [post]
func CreateUser(c *fiber.Ctx) error {

	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "CreateUser",
		"package":  "services",
	})

	var newUser types.AddUser
	if err := utils.ParseBodyAndValidate(c, &newUser); err != nil {
		// err = *fiber.Error, only log error message
		logger.Error(err.Message)
		return err
	}

	// checks if registrations are permitted
	cfg := config.Get()
	if !cfg.AllowRegistrations {
		return utils.NewHTTPError(c, fiber.StatusNotImplemented, fmt.Errorf("server registration is disabled"))
	}

	result := dal.FindUserByName(&dal.User{}, newUser.Name)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if result.RowsAffected > 0 {
			return utils.NewHTTPError(c, fiber.StatusConflict, fmt.Errorf("user %s already exists", newUser.Name))
		} else {
			logger.Error(result.Error)
			return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
		}
	}

	// create the new user
	user := &dal.User{
		Name:     newUser.Name,
		Password: password.Generate(newUser.Password),
		UUID:     uuid.New().String(),
	}
	if err := dal.CreateUser(user).Error; err != nil {
		logger.Error(err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}

	logger.Info("created user: ", user.Name)

	returnUser := &types.GetUser{
		Name: user.Name,
		UUID: user.UUID,
	}

	return c.JSON(returnUser)
}

// @Summary Get the current user
// @Description Get the current user, can be used to verify the user exists
// @Security BasicAuth
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} types.GetUser
// @Failure 401 {string} string
// @Router /users [get]
func GetUser(c *fiber.Ctx) error {

	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "GetUser",
		"package":  "services",
	})

	var user *types.GetUser
	err := dal.FindUserByName(&user, c.Locals("username")).Error
	if err != nil {
		logger.Error(err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// @Summary Delete the current user
// @Description Delete the current user
// @Security BasicAuth
// @Tags users
// @Accept  json
// @Produce  json
// @Success 204 {string} string
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {string} string
// @Router /users [delete]
func DeleteUser(c *fiber.Ctx) error {

	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "DeleteUser",
		"package":  "services",
	})

	if err := dal.DeleteUser(c.Locals("username")).Error; err != nil {
		logger.Error(err)
		return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
	}

	logger.Info("deleted user: ", c.Locals("username"))

	c.Status(fiber.StatusNoContent)
	return nil
}

// @Summary Change the current users setting
// @Description Change the current users setting. The request body must contain either a new name or a new password. If both, the username and the password get changed.
// @Security BasicAuth
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body types.AddUser true "Userupdate"
// @Success 204 {string} string
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {string} string
// @Router /users [put]
func UpdateUser(c *fiber.Ctx) error {

	// define logger for this function
	logger := logging.Log.WithFields(log.Fields{
		"function": "UpdateUser",
		"package":  "services",
	})

	var userUpdate types.UpdateUser
	if err := utils.ParseBodyAndValidate(c, &userUpdate); err != nil {
		// err = *fiber.Error, only log error message
		logger.Error(err.Message)
		return err
	}

	// change only password
	if len(userUpdate.Name) == 0 && len(userUpdate.Password) > 0 {
		username := fmt.Sprint(c.Locals("username"))
		if err := dal.ChangeUserPassword(username, userUpdate.Password).Error; err != nil {
			logger.Error(err)
			return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
		}
		logger.Info("changed password for user: ", username)

		// change only username
	} else if len(userUpdate.Name) > 0 && len(userUpdate.Password) == 0 {

		// check if a user with the requested name already exists
		result := dal.FindUserByName(&dal.User{}, userUpdate.Name)
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if result.RowsAffected > 0 {
				return utils.NewHTTPError(c, fiber.StatusConflict, fmt.Errorf("user %s already exists", userUpdate.Name))
			} else {
				return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
			}
		}

		// change the username
		username := fmt.Sprint(c.Locals("username"))
		if err := dal.ChangeUserName(username, userUpdate.Name).Error; err != nil {
			logger.Error(err)
			return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
		}
		logger.Info(fmt.Sprintf("changed username from %s to %s", username, userUpdate.Name))

		// change username and password
	} else if len(userUpdate.Name) > 0 && len(userUpdate.Password) > 0 {
		// check if a user with the requested name already exists
		result := dal.FindUserByName(&dal.User{}, userUpdate.Name)
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if result.RowsAffected > 0 {
				return utils.NewHTTPError(c, fiber.StatusConflict, fmt.Errorf("user %s already exists", userUpdate.Name))
			} else {
				logger.Error(result.Error)
				return utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
			}
		}

		// change the username and password
		username := fmt.Sprint(c.Locals("username"))
		if err := dal.ChangeUserAndPassword(username, userUpdate.Name, userUpdate.Password).Error; err != nil {
			return utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
		}
		logger.Info(fmt.Sprintf("changed username from %s to %s, changed password too...", username, userUpdate.Name))
	} else {
		return utils.NewHTTPError(c, fiber.StatusBadRequest, fmt.Errorf("no name or password specified"))
	}
	c.Status(fiber.StatusNoContent)
	return nil

}
