package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}

}

func main() {

	os.Setenv("SLACK_BOT_TOKEN","xoxb-4716670774502-4723424242338-vk1tkoTG2NGUhjXEeQG6GAwd")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A04M9C1BN1Y-4742900899361-ed4f33c47cb1cd8c03940d79d756868f87796d4354d82378635bcb5ad745ae7d")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())
	bot.Command("ping",&slacker.CommandDefinition{
		Handler: func (botctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter)  {
			response.Reply("pong")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
