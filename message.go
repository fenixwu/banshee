package banshee

import (
	"errors"

	"github.com/multiplay/go-slack/chat"
	"github.com/sahilm/fuzzy"
)

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
func (b *Banshee) FuzzyPublish() (err error) {
	return publish(FUZZY, b.pattern, b.message)
}

// ExactPublish Publish a message in one channel which is exactly sets
func (b *Banshee) ExactPublish() (err error) {
	return publish(EXACT, b.pattern, b.message)
}

func publish(publishMode PublishMode, pattern string, message *chat.Message) (err error) {
	switch publishMode {
	case FUZZY:
		matches := fuzzy.Find(pattern, channels)
		if matches.Len() == 0 {
			err = errors.New("no channel matches")
			return
		}

		for i := 0; i < len(matches); i++ {
			ch := matches[i].Str
			if v, ok := clientMap[ch]; ok {
				_, err = message.Send(v)
				return
			}
		}
	default:
		if client, ok := clientMap[pattern]; ok {
			_, err = message.Send(client)
			return
		}
		err = errors.New("channel is not exist")
	}
	return
}
