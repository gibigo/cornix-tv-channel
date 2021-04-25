package config

import (
	"errors"
	"io/fs"
	"log"
	"os"

	"github.com/gibigo/cornix-tv-channel/config/database"
	"github.com/spf13/viper"
)

const (
	DefaultConfigurationFilePath = "config/config.yml"
)

var (
	ErrConfigNotLoaded = errors.New("configuration is nil")
	config             *Config
)

type Config struct {
	AllowRegistrations bool `mapstructure:"registration"`
	Database           *database.Config
}

func Get() *Config {
	if config == nil {
		panic(ErrConfigNotLoaded)
	}
	return config
}

func Load(configFile string) error {
	cfg, err := readConfiguration(configFile)
	if err != nil {
		return err
	}
	config = cfg
	return nil
}

func readConfiguration(fileName string) (config *Config, err error) {

	viper.SetConfigType("yaml")

	// check if file exists
	var readFromFile bool

	var succ fs.FileInfo
	succ, err = os.Stat(fileName)
	if succ != nil && !(os.IsNotExist(err)) {
		viper.SetConfigFile(fileName)
		log.Printf("[config][Load] Reading configuration from configFile=%s", fileName)
		readFromFile = true
	} else {
		readFromFile = false
		log.Print("[config][Load] Reading configuration from environment vars")
	}

	viper.BindEnv("registration", "REGISTRATION")
	viper.BindEnv("database.debug", "DATABASE_DEBUG")

	viper.SetDefault("database.debug", false)

	viper.AutomaticEnv()

	if readFromFile {
		err = viper.ReadInConfig()
		if err != nil {
			return
		}
	}

	err = viper.Unmarshal(&config)
	if err == nil {
		validateGeneralConfig(config)
	}
	return
}

func validateGeneralConfig(cfg *Config) {}
