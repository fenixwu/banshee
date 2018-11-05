package banshee

import (
	"github.com/multiplay/go-slack/chat"
	"github.com/sahilm/fuzzy"
)

type Banshee struct {
	pattern string
	message *chat.Message
}

// New a banshee channel, default publish mode is EXACT.
func New(pattern string) *Banshee {
	return &Banshee{
		pattern: pattern,
		message: &chat.Message{
			Markdown:       true,
			UnfurlLinks:    false,
			UnfurlMedia:    false,
			ReplyBroadcast: false,
			Parse:          "none",
			Attachments:    []*chat.Attachment{},
		},
	}
}

// SetMessage Body of the message to send. For formatting options, see https://api.slack.com/docs/formatting.
func (b *Banshee) SetMessage(msg string) *Banshee {
	b.message.Text = msg
	return b
}

func (b *Banshee) UseFullParseMode() *Banshee {
	b.message.Parse = "full"
	return b
}

// CustomizeIconAndName use customized username and icon.
// 1. Set username before using customized icon.
// 2. Emoji OVERRIDES iconURL.
func (b *Banshee) CustomizeIconAndName(username, iconEmoji, iconURL string) *Banshee {
	b.message.Username = username
	b.message.IconURL = iconURL
	b.message.IconEmoji = iconEmoji
	return b
}

func (b *Banshee) DisableMarkdown() *Banshee {
	b.message.Markdown = false
	return b
}

func (b *Banshee) EnableUnfurlLinks() *Banshee {
	b.message.UnfurlLinks = true
	return b
}

func (b *Banshee) EnableUnfurlMedia() *Banshee {
	b.message.UnfurlMedia = true
	return b
}

func (b *Banshee) AddAttachment(attachment *Attachment) *Banshee {
	b.message.Attachments = append(b.message.Attachments, attachment.attachment)
	return b
}

// FuzzyPublish Publish a message in channels which are fuzzy sets
// Pattern "ab" will get ".*[aA][bB].*".
func (b *Banshee) FuzzyPublish() error {
	matcheChannels := fuzzy.Find(b.pattern, channels)
	if matcheChannels.Len() == 0 {
		return nil
	}

	for _, channel := range matcheChannels {
		client, ok := clients[channel.Str]
		if !ok {
			continue
		}

		if _, err := b.message.Send(client); err != nil {
			return err
		}
	}

	return nil
}

// ExactPublish Publish a message in one channel which is exactly sets
// Pattern "ab" will get "ab".
func (b *Banshee) ExactPublish() error {
	client, ok := clients[b.pattern]
	if !ok {
		return nil
	}

	if _, err := b.message.Send(client); err != nil {
		return err
	}

	return nil
}
