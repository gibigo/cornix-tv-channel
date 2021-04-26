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

	// TODO check for registration var

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
// @Failure 409 {object} utils.HTTPError
// @Router /users [delete]
func DeleteUser(c *fiber.Ctx) error {
	if err := dal.DeleteUser(c.Locals("username")).Error; err != nil {
		utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
		return err
	}
	return nil
}

// @Summary Change the current users setting
// @Description Change the current users setting. The request body must contain either a new name or a new password. If both, the username and the password get changed.
// @Security BasicAuth
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body types.AddUser true "Userupdate"
// @Success 200 {string} string
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {string} string
// @Router /users [put]
func UpdateUser(c *fiber.Ctx) error {
	var userUpdate types.UpdateUser
	if err := utils.ParseBodyAndValidate(c, &userUpdate); err != nil {
		return err
	}

	if len(userUpdate.Name) == 0 && len(userUpdate.Password) > 0 {
		username := fmt.Sprint(c.Locals("username"))
		if err := dal.ChangeUserPassword(username, userUpdate.Password).Error; err != nil {
			utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
			return err
		}
	} else if len(userUpdate.Name) > 0 && len(userUpdate.Password) == 0 {

		result := dal.FindUserByName(&dal.User{}, userUpdate.Name)
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if result.RowsAffected > 0 {
				utils.NewHTTPError(c, fiber.StatusConflict, fmt.Errorf("user %s already exists", userUpdate.Name))
				return result.Error
			} else {
				utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
				return result.Error
			}
		}

		username := fmt.Sprint(c.Locals("username"))
		if err := dal.ChangeUserName(username, userUpdate.Name).Error; err != nil {
			utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
			return err
		}
	} else if len(userUpdate.Name) > 0 && len(userUpdate.Password) > 0 {
		result := dal.FindUserByName(&dal.User{}, userUpdate.Name)
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if result.RowsAffected > 0 {
				utils.NewHTTPError(c, fiber.StatusConflict, fmt.Errorf("user %s already exists", userUpdate.Name))
				return result.Error
			} else {
				utils.NewHTTPError(c, fiber.StatusInternalServerError, result.Error)
				return result.Error
			}
		}

		username := fmt.Sprint(c.Locals("username"))
		if err := dal.ChangeUserAndPassword(username, userUpdate.Name, userUpdate.Password).Error; err != nil {
			utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
			return err
		}
	} else {
		utils.NewHTTPError(c, fiber.StatusNotAcceptable, fmt.Errorf("no name or password specified"))
		return nil
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

	// get user id
	username := fmt.Sprint(c.Locals("username"))
	user := &dal.User{Name: username}
	if err := dal.FindUserByName(user, username).Error; err != nil {
		utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
		return nil
	}

	result := dal.FindAllStrategiesFromUser(nil, user.ID)
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
		fmt.Println(tmpStrategy)
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
		utils.NewHTTPError(c, fiber.StatusNotFound, fmt.Errorf("user %s has no strategies", c.Locals("username")))
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
// @Param user body types.AddStrategy true "Strategy to create"
// @Success 200 {object} types.Strategy
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {string} string
// @Router /strategies [post]
func CreateStrategy(c *fiber.Ctx) error {
	var s types.AddStrategy
	if err := utils.ParseBodyAndValidate(c, &s); err != nil {
		return err
	}

	// find the userid
	username := fmt.Sprint(c.Locals("username"))
	user := &dal.User{Name: username}
	if err := dal.FindUserByName(user, username).Error; err != nil {
		utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
		return nil
	}

	sNew := &dal.Strategy{
		AllowCounter: s.AllowCounter,
		Entries:      s.Entries,
		TPs:          s.TPs,
		SL:           s.SL,
		UserID:       user.ID,
	}

	if err := dal.CreateStrategy(sNew).Error; err != nil {
		utils.NewHTTPError(c, fiber.StatusInternalServerError, err)
		return nil
	}

	result := dal.FindStrategyByID(sNew, sNew.ID)
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

// @Summary Delete a strategy
// @Description Delete a strategy
// @Security BasicAuth
// @Tags strategies
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {string} string
// @Failure 404 {object} utils.HTTPError
// @Failure 409 {object} utils.HTTPError
// @Router /strategies [delete]
func DeleteStrategy(c *fiber.Ctx) error {
	return c.Status(400).SendString("not yet implemented")
}
