package banshee

import (
	"errors"
	"fmt"

	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/webhook"
	"github.com/sahilm/fuzzy"
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
	publishMode PublishMode
	pattern     string
}

// RegistChannel regists new channel.
func RegistChannel(channel, webhookURL string) error {
	if err := registChannel(channel, webhookURL); err != nil {
		return err
	}

	appendChannel(channel)
	return nil
}

// 註冊頻道
func registChannel(channel, webhookURL string) error {
	newClient := webhook.New(webhookURL)
	clientMap[channel] = newClient
	_, err := (&chat.Message{Text: channel + " is registed"}).Send(newClient)
	return err
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
	return &Banshee{EXACT, pattern}
}

// Get publish mode
func (b *Banshee) getPublishMode() PublishMode {
	return b.publishMode
}

// SetPublishModeExact Publish messages in one channel which is exactly matches	the pattern
func (b *Banshee) SetPublishModeExact() {
	b.publishMode = EXACT
}

// SetPublishModeFuzzy Publish messages in channels which are fuzzy matches the pattern
func (b *Banshee) SetPublishModeFuzzy() {
	b.publishMode = FUZZY
}

// Publish a message
func (b *Banshee) Publish(message string) (err error) {
	msg := &chat.Message{Text: message}
	switch b.publishMode {
	case FUZZY:
		matches := fuzzyFindChannel(b.pattern)
		if matches.Len() == 0 {
			err = errors.New("no channel matches")
			return
		}

		for i := 0; i < len(matches); i++ {
			ch := matches[i].Str
			fmt.Println(ch)
			if v, ok := clientMap[ch]; ok {
				_, err = msg.Send(v)
				return
			}
			err = errors.New("channel is not exist")
		}
	default:
		if client, ok := clientMap[b.pattern]; ok {
			_, err = msg.Send(client)
			return
		}
		err = errors.New("channel is not exist")
	}
	return
}

func fuzzyFindChannel(pattern string) fuzzy.Matches {
	return fuzzy.Find(pattern, channels)
}

// TODO: Custom webhook message
