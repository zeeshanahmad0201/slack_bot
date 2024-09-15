package bot

import (
	"fmt"
	"log"

	"github.com/shomali11/slacker"
	"github.com/zeeshanahmad0201/gobot/config"
	"github.com/zeeshanahmad0201/gobot/pkg/helpers"
)

func printCommandEvents(eventChannel <-chan *slacker.CommandEvent) {
	for event := range eventChannel {
		fmt.Printf("event %v", event.Parameters)
	}
}

func StartBot() {
	bot := config.GetBot()

	go printCommandEvents(bot.CommandEvents())

	AddCommand()

	startListener()

}

func AddCommand() {
	bot := config.GetBot()

	bot.Command("query for bot - <message>", &slacker.CommandDefinition{
		Description: "Send any question to wolfram",
		Examples: []string{
			"Who am I?",
		},
		Handler: func(bc slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("message")
			fmt.Printf("command: %s", query)

			response.Reply("received")

		},
	})
}

func startListener() {
	bot := config.GetBot()
	ctx, cancel := helpers.ContextWithCancel()
	defer cancel()

	if err := bot.Listen(ctx); err != nil {
		log.Fatalf("Error while listening to bot: %v", err)
	}
}
