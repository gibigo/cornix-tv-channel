package services

import (
	"errors"
	"fmt"

	"github.com/gibigo/cornix-tv-channel/app/dal"
	"github.com/gibigo/cornix-tv-channel/app/types"
	"github.com/gibigo/cornix-tv-channel/utils"
	"github.com/gibigo/cornix-tv-channel/utils/password"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetUser godoc
// @Summary Get the current user
// @Description Get the current user, can be used to verify the user exists
// @Security BasicAuth
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Failure 401 {string} string
// @Router /users [get]
func GetUser(c *fiber.Ctx) error {
	var user *dal.User
	err := dal.FindUserByName(&user, c.Locals("username")).Error
	if err != nil {
		utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
		return err
	}
	return nil
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body types.AddUser true "User to create"
// @Success 200 {string} string
// @Failure 401 {string} string
// @Failure 409 {object} utils.HTTPError
// @Router /users [post]
func CreateUser(c *fiber.Ctx) error {
	var u types.AddUser
	if err := utils.ParseBodyAndValidate(c, &u); err != nil {
		return err
	}

	result := dal.FindUserByName(&dal.User{}, u.Name)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if result.RowsAffected > 0 {
			utils.NewHTTPError(c, fiber.StatusConflict, fmt.Errorf("user %s already exists", u.Name))
			return result.Error
		} else {
			utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
			return result.Error
		}
	}

	user := &dal.User{
		Name:     u.Name,
		Password: password.Generate(u.Password),
	}

	if err := dal.CreateUser(user).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUser godoc
// @Summary Delete the current user
// @Description Delete the current user
// @Security BasicAuth
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {string} string
// @Router /users [delete]
func DeleteUser(c *fiber.Ctx) error {
	if err := dal.DeleteUser(c.Locals("username")).Error; err != nil {
		utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
		return err
	}
	return nil

}

// GetUser godoc
// @Summary Get all strategies
// @Description Get all strategies of the current user
// @Security BasicAuth
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} types.Strategy
// @Failure 401 {string} string
// @Router /users [get]
func GetStrategies(c *fiber.Ctx) error {
	var strategies []types.Strategy
	if err := dal.FindAllStrategiesFromUser(&strategies, c.Locals("username")).Error; err != nil {
		utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
		return nil
	}
	if len(strategies) == 0 {
		err := fmt.Errorf("the user %s does not have any strategies", c.Locals("username"))
		utils.NewHTTPError(c, fiber.StatusNotFound, err)
		return nil
	}

	fmt.Println(strategies)
	return nil
}

func CreateStrategy(c *fiber.Ctx) error {
	var s dal.Strategy
	if err := utils.ParseBodyAndValidate(c, &s); err != nil {
		fmt.Println(string(c.Body()))
		return err
	}
	fmt.Println(s)

	username := fmt.Sprint(c.Locals("username"))
	s.UserName = username
	result := dal.CreateStrategy(&s)
	fmt.Println(result.Error)
	fmt.Println(s.ID)
	return nil
}
