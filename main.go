package main

import (
	"github.com/joho/godotenv"
	"github.com/zeeshanahmad0201/gobot/internal/bot"
)

func main() {
	godotenv.Load()

	bot.StartBot()
}
