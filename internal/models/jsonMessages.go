package models

type JSONMessages struct {
	BotToken string        `json:"bot_token"`
	Channels []JSONMessage `json:"channels"`
}

type JSONMessage struct {
	ChannelName string `json:"channel"`
	Text        string `json:"text"`
}
