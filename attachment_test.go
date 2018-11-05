package banshee

import (
	"reflect"
	"testing"

	"github.com/multiplay/go-slack/chat"
)

func newDefaultAttachment() *Attachment {
	return &Attachment{&chat.Attachment{Fields: []*chat.Field{}}}
}

func TestNewAttachment(t *testing.T) {
	tests := []struct {
		name string
		want *Attachment
	}{
		{"TestNewAttachment", newDefaultAttachment()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAttachment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetFallback(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetFallback("test")
	type args struct {
		fallback string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetFallback", NewAttachment(), args{"test"}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetFallback(tt.args.fallback); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetFallback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetColor(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetColor("FFFFFF")
	type args struct {
		color string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetColor", NewAttachment(), args{"FFFFFF"}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetColor(tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetPreText(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetPreText("test")
	type args struct {
		preText string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetPreText", NewAttachment(), args{"test"}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetPreText(tt.args.preText); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetPreText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetAuthorName(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetAuthorName("username")
	type args struct {
		authorName string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetAuthorName", NewAttachment(), args{"username"}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetAuthorName(tt.args.authorName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetAuthorName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetAuthorLink(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetAuthorLink("link")
	type args struct {
		authorLink string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetAuthorLink", NewAttachment(), args{"link"}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetAuthorLink(tt.args.authorLink); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetAuthorLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetAuthorIcon(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetAuthorIcon(":icon:")
	type args struct {
		authorIcon string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetAuthorIcon", NewAttachment(), args{":icon:"}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetAuthorIcon(tt.args.authorIcon); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetAuthorIcon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetTitle(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetTitle("title")
	type args struct {
		title string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetTitle", NewAttachment(), args{"title"}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetTitle(tt.args.title); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetTitleLink(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetTitleLink("title.link")
	type args struct {
		titleLink string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetTitleLink", NewAttachment(), args{"title.link"}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetTitleLink(tt.args.titleLink); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetTitleLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetText(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetText("test text")
	type args struct {
		text string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetText", NewAttachment(), args{"test text"}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetText(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetImageURL(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetImageURL("test.com")
	type args struct {
		imageURL string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetImageURL", NewAttachment(), args{"test.com"}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetImageURL(tt.args.imageURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetImageURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetThumbURL(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetThumbURL("test.com")
	type args struct {
		thumbURL string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetThumbURL", NewAttachment(), args{"test.com"}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetThumbURL(tt.args.thumbURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetThumbURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetFooter(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetFooter("footer")
	type args struct {
		footer string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetFooter", NewAttachment(), args{"footer"}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetFooter(tt.args.footer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetFooter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetFooterIcon(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetFooterIcon(":icon:")
	type args struct {
		footerIcon string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetFooterIcon", NewAttachment(), args{":icon:"}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetFooterIcon(tt.args.footerIcon); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetFooterIcon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetMarkdownIn(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetMarkdownIn([]string{"a"})
	type args struct {
		markdownIn []string
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetMarkdownIn", NewAttachment(), args{[]string{"a"}}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetMarkdownIn(tt.args.markdownIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetMarkdownIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_SetTimeStamp(t *testing.T) {
	expect := newDefaultAttachment()
	expect.SetTimeStamp(1542613263)
	type args struct {
		timeStamp int
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_SetTimeStamp", NewAttachment(), args{1542613263}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetTimeStamp(tt.args.timeStamp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.SetTimeStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttachment_AddField(t *testing.T) {
	expect := newDefaultAttachment()
	expect.AddField("title", "content", false)
	type args struct {
		title string
		value string
		short bool
	}
	tests := []struct {
		name string
		a    *Attachment
		args args
		want *Attachment
	}{
		{"TestAttachment_AddField", NewAttachment(), args{"title", "content", false}, expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AddField(tt.args.title, tt.args.value, tt.args.short); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attachment.AddField() = %v, want %v", got, tt.want)
			}
		})
	}
}
