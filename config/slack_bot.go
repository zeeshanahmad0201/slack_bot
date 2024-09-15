package config

import (
	"os"
	"sync"

	"github.com/shomali11/slacker"
)

var (
	bot     *slacker.Slacker
	botInst sync.Once
)

func GetBot() *slacker.Slacker {
	botInst.Do(func() {
		botToken := os.Getenv("SLACK_BOT_TOKEN")
		appToken := os.Getenv("SLACK_APP_TOKEN")

		if botToken == "" || appToken == "" {
			panic("Missing SLACK_BOT_TOKEN or SLACK_APP_TOKEN")
		}

		bot = slacker.NewClient(botToken, appToken)
	})
	return bot
}
