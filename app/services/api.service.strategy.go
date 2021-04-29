package services

import "github.com/gofiber/fiber/v2"

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
	return nil
}

/*

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
*/
