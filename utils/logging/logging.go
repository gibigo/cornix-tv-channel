package logging

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	Log *log.Logger
)

type Config struct {
	Debug bool `yaml:"debugging"`
}

func Init() {
	Log = log.New()

	Log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	Log.SetOutput(os.Stdout)
	Log.SetLevel(log.WarnLevel)

}
