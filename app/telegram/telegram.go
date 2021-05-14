package telegram

import (
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

var (
	Bot *gotgbot.Bot
)

type Config struct {
	Token string `mapstructure:"token"`
}

func (c *Config) NewBot() *gotgbot.Bot {
	var err error
	Bot, err = gotgbot.NewBot(c.Token, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	return Bot
}
