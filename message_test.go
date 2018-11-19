package banshee

import (
	"reflect"
	"testing"

	"github.com/multiplay/go-slack/chat"
)

const expectPattern = "test"

func init() {
	RegistChannel(expectPattern, RealWebhookURL)
}

func newDefaultChatMsg() *chat.Message {
	return &chat.Message{
		Markdown:       true,
		UnfurlLinks:    false,
		UnfurlMedia:    false,
		ReplyBroadcast: false,
		Parse:          "none",
		Attachments:    []*chat.Attachment{},
	}
}

func TestBanshee_SetMessage(t *testing.T) {
	expect := newDefaultChatMsg()
	expect.Text = "test msg"
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		b    *Banshee
		args args
		want *Banshee
	}{
		{"TestBanshee_SetMessage", New(expectPattern), args{"test msg"}, &Banshee{expectPattern, expect}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.SetMessage(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Banshee.SetMessage() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestBanshee_UseFullParseMode(t *testing.T) {
	expect := newDefaultChatMsg()
	expect.Parse = "full"
	tests := []struct {
		name string
		b    *Banshee
		want *Banshee
	}{
		{"TestBanshee_UseFullParseMode", New(expectPattern), &Banshee{expectPattern, expect}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.UseFullParseMode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Banshee.UseFullParseMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBanshee_CustomizeIconAndName(t *testing.T) {
	username := "develop"
	iconEmoji := ":ok:"
	expect := newDefaultChatMsg()
	expect.Username = username
	expect.IconEmoji = iconEmoji
	type args struct {
		username  string
		iconEmoji string
		iconURL   string
	}
	tests := []struct {
		name string
		b    *Banshee
		args args
		want *Banshee
	}{
		{"TestBanshee_CustomizeIconAndName", New(expectPattern), args{username, iconEmoji, ""}, &Banshee{expectPattern, expect}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.CustomizeIconAndName(tt.args.username, tt.args.iconEmoji, tt.args.iconURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Banshee.CustomizeIconAndName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBanshee_DisableMarkdown(t *testing.T) {
	expect := newDefaultChatMsg()
	expect.Markdown = false
	tests := []struct {
		name string
		b    *Banshee
		want *Banshee
	}{
		{"TestBanshee_DisableMarkdown", New(expectPattern), &Banshee{expectPattern, expect}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.DisableMarkdown(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Banshee.DisableMarkdown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBanshee_EnableUnfurlLinks(t *testing.T) {
	expect := newDefaultChatMsg()
	expect.UnfurlLinks = true
	tests := []struct {
		name string
		b    *Banshee
		want *Banshee
	}{
		{"TestBanshee_EnableUnfurlLinks", New(expectPattern), &Banshee{expectPattern, expect}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.EnableUnfurlLinks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Banshee.EnableUnfurlLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBanshee_EnableUnfurlMedia(t *testing.T) {
	expect := newDefaultChatMsg()
	expect.UnfurlMedia = true
	tests := []struct {
		name string
		b    *Banshee
		want *Banshee
	}{
		{"TestBanshee_EnableUnfurlMedia", New(expectPattern), &Banshee{expectPattern, expect}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.EnableUnfurlMedia(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Banshee.EnableUnfurlMedia() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBanshee_AddAttachment(t *testing.T) {
	expectAttachment := &chat.Attachment{}
	expect := newDefaultChatMsg()
	expect.AddAttachment(expectAttachment)
	type args struct {
		attachment *Attachment
	}
	tests := []struct {
		name string
		b    *Banshee
		args args
		want *Banshee
	}{
		{"TestBanshee_AddAttachment", New(expectPattern),
			args{&Attachment{expectAttachment}},
			&Banshee{expectPattern, expect},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.AddAttachment(tt.args.attachment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Banshee.AddAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBanshee_FuzzyPublish(t *testing.T) {
	tests := []struct {
		name    string
		b       *Banshee
		wantErr bool
	}{
		{"TestBanshee_FuzzyPublish - error null channels", New("NONEXISTENT_CHANNEL"), true},
		{"TestBanshee_FuzzyPublish - error null msg", New(expectPattern), true},
		{"TestBanshee_FuzzyPublish - pass", New(expectPattern), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "TestBanshee_FuzzyPublish - pass" {
				tt.b.SetMessage("test")
			}
			if err := tt.b.FuzzyPublish(); (err != nil) != tt.wantErr {
				t.Errorf("Banshee.FuzzyPublish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBanshee_ExactPublish(t *testing.T) {
	tests := []struct {
		name    string
		b       *Banshee
		wantErr bool
	}{
		{"TestBanshee_ExactPublish - error null channels", New("NONEXISTENT_CHANNEL"), true},
		{"TestBanshee_ExactPublish - error null msg", New(expectPattern), true},
		{"TestBanshee_ExactPublish - pass", New(expectPattern), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "TestBanshee_ExactPublish - pass" {
				tt.b.SetMessage("test")
			}
			if err := tt.b.ExactPublish(); (err != nil) != tt.wantErr {
				t.Errorf("Banshee.ExactPublish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_publish(t *testing.T) {
	expect := newDefaultChatMsg()
	expect.Text = "Test_publish"
	type args struct {
		publishMode PublishMode
		pattern     string
		message     *chat.Message
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test_publish - EXACT pass", args{EXACT, "test", expect}, false},
		{"Test_publish - EXACT null chnnel", args{EXACT, "test2", expect}, true},
		{"Test_publish - FUZZY pass", args{FUZZY, "test", expect}, false},
		{"Test_publish - FUZZY null channel", args{FUZZY, "null", expect}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := publish(tt.args.publishMode, tt.args.pattern, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
