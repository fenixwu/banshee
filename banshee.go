package banshee

import (
	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/webhook"
)

// PublishMode 發佈模式
type PublishMode string

const (
	// FUZZY mode
	FUZZY PublishMode = "fuzzy"
	// EXACT mdoe
	EXACT PublishMode = "exact"
)

var (
	channels  = []string{}                   // 頻道清單
	clientMap = map[string]*webhook.Client{} // Map: 頻道 -> Client
)

// Banshee is banshee
// publishMode "EXACT" -> pattern "ab" will get "ab".
// publishMode "FUZZY" -> pattern "ab" will get ".*[aA][bB].*".
type Banshee struct {
	pattern string
	message *chat.Message
}

// RegistChannel regists new channel.
func RegistChannel(channel, webhookURL string) error {
	registChannel(channel, webhookURL)
	appendChannel(channel)
	return nil
}

func registChannel(channel, webhookURL string) {
	newClient := webhook.New(webhookURL)
	clientMap[channel] = newClient
}

// 註冊頻道搜尋清單
func appendChannel(channel string) {
	if !contains(channel, channels) {
		channels = append(channels, channel)
	}
}

func contains(t string, conditions []string) (isContains bool) {
	for i := 0; i < len(conditions); i++ {
		if t == conditions[i] {
			isContains = true
			return
		}
	}
	return
}

// New a banshee channel, default publish mode is EXACT.
func New(pattern string) *Banshee {
	return &Banshee{pattern, &chat.Message{
		Markdown:       true,
		UnfurlLinks:    false,
		UnfurlMedia:    false,
		ReplyBroadcast: false,
		Parse:          "none",
		Attachments:    []*chat.Attachment{}},
	}
}
