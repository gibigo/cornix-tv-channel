package database

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func init() {
	// ensure folder ./data exists and if not create it
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		if err := os.Mkdir("./data", 0755); err != nil {
			panic(fmt.Sprintf("[database][init] error: %s", err))
		}
	}
}

type Config struct {
	Debug bool `mapstructure:"debug"`
}

func Connect(config *Config) {
	var err error
	var gormConfig = &gorm.Config{QueryFields: true}

	if !config.Debug {
		gormConfig.Logger = logger.Default.LogMode(logger.LogLevel(0))
		gormConfig.QueryFields = true
	}

	DB, err = gorm.Open(sqlite.Open("data/database.db"), gormConfig)
	if err != nil {
		panic(fmt.Sprintf("[database][init] error: %s", err))
	}

	// enable foreign key constraints
	DB.Exec("PRAGMA foreign_keys = ON;")
}

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
