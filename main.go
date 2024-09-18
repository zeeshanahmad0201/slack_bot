package main

import (
	"github.com/joho/godotenv"
	"github.com/zeeshanahmad0201/slack_bot/internal/bot"
)

func main() {
	godotenv.Load()

	bot.StartBot()
}
