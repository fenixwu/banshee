package banshee

import "github.com/multiplay/go-slack/chat"

type Attachment struct {
	attachment *chat.Attachment
}

type Field struct {
	field *chat.Field
}

func NewAttachment() *Attachment {
	return &Attachment{&chat.Attachment{Fields: []*chat.Field{}}}
}

func (a *Attachment) Text(txt string) *Attachment {
	a.attachment.Text = txt
	return a
}

func NewField(title, value string, short bool) *Field {
	return &Field{&chat.Field{title, value, short}}
}

func (a *Attachment) AddField(field *Field) *Attachment {
	a.attachment.Fields = append(a.attachment.Fields, field.field)
	return a
}
