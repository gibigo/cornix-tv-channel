package logging

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	Log = log.New()
)

type Config struct {
	LogLevel string `mapstructure:"logLevel"`
}

// Initalize a new logger.
//You can set a different logging level by useing the logging.logLevel YAML key or using the LOG_LEVEL env variable.
// By default the logging level is set to "info" but it can be changed to either "warn" or "debug"
func Init(cfg *Config) {

	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if strings.EqualFold(cfg.LogLevel, "Info") {
		Log.SetLevel(log.InfoLevel)
	} else if strings.EqualFold(cfg.LogLevel, "Warn") {
		Log.SetLevel(log.WarnLevel)
	} else if strings.EqualFold(cfg.LogLevel, "Debug") {
		Log.SetLevel(log.DebugLevel)
	}

}
