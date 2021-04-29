package main

import (
	"os"

	"github.com/gibigo/cornix-tv-channel/app/dal"
	"github.com/gibigo/cornix-tv-channel/app/routes"
	"github.com/gibigo/cornix-tv-channel/config"
	"github.com/gibigo/cornix-tv-channel/config/database"
	_ "github.com/gibigo/cornix-tv-channel/docs"
	"github.com/gibigo/cornix-tv-channel/utils/logging"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// load the configuration
	cfg := loadConfiguration()

	// init the logger
	logging.Init(cfg.Logging)

	// who would have thought, connects to the database
	database.Connect(cfg.Database)
	// create the database structure
	database.Migrate(&dal.User{}, &dal.TVSignal{}, &dal.Channel{}, &dal.Strategy{}, &dal.ZoneStrategy{}, &dal.TargetStrategy{}, &dal.Entry{}, &dal.TP{}, &dal.SL{})

	// lets fire up the API
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")
}

func loadConfiguration() *config.Config {
	var err error
	customConfigFile := os.Getenv("CONFIG_FILE")
	if len(customConfigFile) > 0 {
		err = config.Load(customConfigFile)
	} else {
		err = config.Load(config.DefaultConfigurationFilePath)
	}
	if err != nil {
		panic(err)
	}
	cfg := config.Get()
	return cfg
}

// @title Cornix-TV-Channel API
// @version 1.0
// @host https://yourforwarder.io
// @BasePath /api/v1
// @license.name MIT
// @license.url https://github.com/gibigo/cornix-tv-channel/blob/master/LICENSE
// @securityDefinitions.basic BasicAuth
