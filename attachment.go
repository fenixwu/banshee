package banshee

import "github.com/multiplay/go-slack/chat"

type Attachment struct {
	attachment *chat.Attachment
}

func NewAttachment() *Attachment {
	return &Attachment{&chat.Attachment{Fields: []*chat.Field{}}}
}

func (a *Attachment) SetFallback(fallback string) *Attachment {
	a.attachment.Fallback = fallback
	return a
}

func (a *Attachment) SetColor(color string) *Attachment {
	a.attachment.Color = color
	return a
}

func (a *Attachment) SetPreText(preText string) *Attachment {
	a.attachment.PreText = preText
	return a
}

func (a *Attachment) SetAuthorName(authorName string) *Attachment {
	a.attachment.AuthorName = authorName
	return a
}

func (a *Attachment) SetAuthorLink(authorLink string) *Attachment {
	a.attachment.AuthorLink = authorLink
	return a
}

func (a *Attachment) SetAuthorIcon(authorIcon string) *Attachment {
	a.attachment.AuthorIcon = authorIcon
	return a
}

func (a *Attachment) SetTitle(title string) *Attachment {
	a.attachment.Title = title
	return a
}

func (a *Attachment) SetTitleLink(titleLink string) *Attachment {
	a.attachment.TitleLink = titleLink
	return a
}

func (a *Attachment) SetText(text string) *Attachment {
	a.attachment.Text = text
	return a
}

func (a *Attachment) SetImageURL(imageURL string) *Attachment {
	a.attachment.ImageURL = imageURL
	return a
}

func (a *Attachment) SetThumbURL(thumbURL string) *Attachment {
	a.attachment.ThumbURL = thumbURL
	return a
}

func (a *Attachment) SetFooter(footer string) *Attachment {
	a.attachment.Footer = footer
	return a
}

func (a *Attachment) SetFooterIcon(footerIcon string) *Attachment {
	a.attachment.FooterIcon = footerIcon
	return a
}

func (a *Attachment) SetMarkdownIn(markdownIn []string) *Attachment {
	a.attachment.MarkdownIn = markdownIn
	return a
}

func (a *Attachment) SetTimeStamp(timeStamp int) *Attachment {
	a.attachment.TimeStamp = timeStamp
	return a
}

func (a *Attachment) AddField(title, value string, short bool) *Attachment {
	a.attachment.Fields = append(a.attachment.Fields, &chat.Field{
		Title: title,
		Value: value,
		Short: short,
	})
	return a
}
