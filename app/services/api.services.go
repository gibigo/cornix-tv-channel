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

// @Summary Get all strategies
// @Description Get all strategies of the current user
// @Security BasicAuth
// @Tags strategies
// @Accept  json
// @Produce  json
// @Success 200 {array} types.Strategy
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {string} string
// @Failure 404 {object} utils.HTTPError
// @Router /strategies [get]
func GetStrategies(c *fiber.Ctx) error {
	result := dal.FindAllStrategiesFromUser(nil, c.Locals("username"))
	if result.Error != nil {
		utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
		return nil
	}

	rows, err := result.Rows()
	if err != nil {
		utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
		return nil
	}
	defer rows.Close()

	strategies := make([]*types.Strategy, 0)
	var indx int64

	for rows.Next() {
		tmpStrategy := &types.Strategy{}
		entries := &types.Entry{}
		tps := &types.TP{}
		sl := &types.SL{}
		err := rows.Scan(&tmpStrategy.ID, &tmpStrategy.AllowCounter, &entries.Diff, &tps.Diff, &sl.Diff)
		if err != nil {
			utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
			return nil
		}
		if len(strategies) < int(tmpStrategy.ID) {
			strategies = append(strategies, tmpStrategy)
			indx = tmpStrategy.ID - 1

			strategies[indx].Entries = make([]*types.Entry, 0)
			strategies[indx].TPs = make([]*types.TP, 0)
			strategies[indx].SL = sl
			strategies[indx].AllowCounter = tmpStrategy.AllowCounter
		}
		strategies[indx].Entries = append(strategies[indx].Entries, entries)
		strategies[indx].TPs = append(strategies[indx].TPs, tps)
	}
	if len(strategies) == 0 {
		utils.NewHTTPError(c, fiber.StatusNotFound, fmt.Errorf("user %s has strategies", c.Locals("username")))
		return nil
	}
	return c.JSON(strategies)
}

// @Summary Create a new strategy
// @Description Create a new strategy for the current user
// @Security BasicAuth
// @Tags strategies
// @Accept  json
// @Produce  json
// @Success 200 {object} types.Strategy
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {string} string
// @Router /strategies [post]
func CreateStrategy(c *fiber.Ctx) error {
	var s dal.Strategy
	if err := utils.ParseBodyAndValidate(c, &s); err != nil {
		fmt.Println(string(c.Body()))
		fmt.Println(err)
		return err
	}

	username := fmt.Sprint(c.Locals("username"))
	s.UserName = username
	if err := dal.CreateStrategy(&s).Error; err != nil {
		utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
		return nil
	}

	result := dal.FindStrategyByID(s, s.ID)
	if result.Error != nil {
		utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
		return nil
	}
	strategy := &types.Strategy{}
	strategy.Entries = make([]*types.Entry, 0)
	strategy.TPs = make([]*types.TP, 0)
	strategy.SL = &types.SL{}

	rows, err := result.Rows()
	if err != nil {
		utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		entries := &types.Entry{}
		tps := &types.TP{}

		err := rows.Scan(&strategy.ID, &strategy.AllowCounter, &entries.Diff, &tps.Diff, &strategy.SL.Diff)
		if err != nil {
			utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
			return nil
		}
		strategy.Entries = append(strategy.Entries, entries)
		strategy.TPs = append(strategy.TPs, tps)
	}
	return c.JSON(strategy)
}
