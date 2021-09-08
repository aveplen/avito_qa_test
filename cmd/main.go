package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aveplen/avito_qa_test/internal/service"
	"github.com/aveplen/avito_qa_test/internal/utils"
	"github.com/slack-go/slack"
)

var filename string

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	flag.StringVar(
		&filename,
		"f",
		"message.json",
		"pass a .json file with messages to post on conversations")

	flag.Parse()
}

func main() {
	jsonData, err := utils.ReadJSON(filename)
	handleError(err)

	convService := service.NewConversationService(jsonData.BotToken)
	conversations, err := convService.GetConversationsList()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	api := slack.New(jsonData.BotToken)
	for _, conv := range conversations {

		for _, chann := range jsonData.Channels {

			if conv.Name != chann.ChannelName {
				continue
			}

			_, timestamp, err := api.PostMessage(conv.ID, slack.MsgOptionText(chann.Text, false))
			handleError(err)

			time, err := utils.TimestampToTime(timestamp)
			handleError(err)

			fmt.Printf("successfully posted message on channel %s at %v\n", conv.Name, time)
		}
	}
}
