package models

import "github.com/slack-go/slack"

type ConvWrapper struct {
	Name      string `json:"name"`
	IsChannel bool   `json:"is_channel"`
	slack.Conversation
}
