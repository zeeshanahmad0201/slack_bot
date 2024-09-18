package bot

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/krognol/go-wolfram"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go"
	"github.com/zeeshanahmad0201/slack_bot/config"
	"github.com/zeeshanahmad0201/slack_bot/pkg/helpers"
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

	bot.Command("- <message>", &slacker.CommandDefinition{
		Description: "Send any question to wolfram",
		Examples: []string{
			"Who am I?",
		},
		Handler: func(bc slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("message")
			fmt.Printf("command: %s\n", query)

			witClient := config.GetWitClient()
			msg, err := witClient.Parse(&witai.MessageRequest{
				Query: query,
			})
			if err != nil {
				response.ReportError(err)
				return
			}
			data, err := json.MarshalIndent(msg, "", "\t")
			fmt.Println(msg)
			rough := string(data[:])
			fmt.Println(rough)
			value := gjson.Get(rough, "entities.wolfram_search_query.0.value")
			fmt.Println(value)
			answer := value.String()
			wolframClient := config.GetWolframClient()
			if wolframClient == nil {
				response.ReportError(errors.New("wolfram client is nil"))
				return
			}
			resp, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				response.ReportError(err)
				return
			}

			response.Reply(resp)

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
