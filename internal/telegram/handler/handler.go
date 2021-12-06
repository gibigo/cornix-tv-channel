package handler

import (
	"errors"
	"regexp"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/gibigo/cornix-tv-channel/internal/api/dal"
	"github.com/gibigo/cornix-tv-channel/internal/utils/logging"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	regexExpression = `(?:#)([\S]+?)(?:\s)`
)

var (
	keywords = [...]string{"CANCEL", "Manually Cancelled", "All take-profit targets achieved", "Target achieved before", "Stoploss ⛔️", "Closed at trailing stoploss after reaching take profit"}
	r        *regexp.Regexp
)

func StartTGHandler(b *gotgbot.Bot) {
	logger := logging.Log.WithFields(log.Fields{
		"function": "StartTGHandler",
		"package":  "handler",
	})

	// compile the regex used in the getSymbol function
	r = regexp.MustCompile(regexExpression)

	// telegram handler
	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher

	channelMessageHandler := handlers.NewMessage(message.Text, handleChannelMessages)
	channelMessageHandler.AllowChannel = true
	dispatcher.AddHandler(channelMessageHandler)

	err := updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	logger.Infof("%s has been started...", b.User.Username)

	updater.Idle()
}

func handleChannelMessages(b *gotgbot.Bot, ctx *ext.Context) error {
	logger := logging.Log.WithFields(log.Fields{
		"function": "StartTGHandler",
		"package":  "handler",
	})

	// check if the message is from a channel which is in the database
	var channel dal.Channel
	if err := dal.FindChannelByTelegramId(&channel, ctx.EffectiveChat.Id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	// check if the message contains a keyword
	var contains bool
	for _, v := range keywords {
		if strings.Contains(ctx.EffectiveMessage.Text, v) {
			contains = true
		}
	}
	if !contains {
		return nil
	}

	// extract the symbol from the message
	symbol := getSymbol(ctx.EffectiveMessage.Text)

	// remove the open signal from the database
	if err := dal.DeleteSignal(channel.ID, symbol).Error; err != nil {
		logger.Errorf("error while removing signal from database: %s", err)
	}

	return nil
}

func getSymbol(m string) string {
	x := r.FindStringSubmatch(m)[1]
	return strings.ReplaceAll(x, "/", "")
}
