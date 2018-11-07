package banshee

import (
	"errors"

	"github.com/multiplay/go-slack/chat"
	"github.com/sahilm/fuzzy"
)

type IMessage interface {
	AddAttachment(attachment *Attachment) IMessage
	FuzzyPublish() (err error)
	ExactPublish() (err error)
}

func (b *Banshee) Message(msg string) IMessage {
	b.message.Text = msg
	return b
}

func (b *Banshee) AddAttachment(attachment *Attachment) IMessage {
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
		matches := fuzzyFindChannel(pattern)
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
			err = errors.New("channel is not exist")
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

func fuzzyFindChannel(pattern string) fuzzy.Matches {
	return fuzzy.Find(pattern, channels)
}
