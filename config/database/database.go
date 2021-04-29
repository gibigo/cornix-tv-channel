package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

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

	DB, err = gorm.Open(sqlite.Open("database.db"), gormConfig)
	if err != nil {
		panic(fmt.Sprintf("[database][init] error: %s", err))
	}
}

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
