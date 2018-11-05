package banshee

import "github.com/multiplay/go-slack/webhook"

var (
	channels []string
	clients  = map[string]*webhook.Client{} // 頻道清單
)

// RegisterChannel Registers new channel.
func RegisterChannel(channel, webhookURL string) {
	if _, ok := clients[channel]; ok {
		return
	}

	clients[channel] = webhook.New(webhookURL)
	channels = append(channels, channel)
}
