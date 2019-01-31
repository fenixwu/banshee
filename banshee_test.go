package banshee

import (
	"reflect"
	"testing"

	"github.com/multiplay/go-slack/chat"
)

func init() {
	myChannels := []string{"a", "ab", "abc", "b", "bc", "c"}
	for _, channel := range myChannels {
		RegistChannel(channel, RealWebhookURL)
	}
}

const RealWebhookURL = "https://hooks.slack.com/services/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"

func Test_appendChannel(t *testing.T) {
	type args struct {
		channel string
	}
	tests := []struct {
		name   string
		length int
		args   args
	}{
		{"A", 1, args{"a"}},
		{"B", 2, args{"b"}},
		{"C", 2, args{"a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appendChannel(tt.args.channel)
			if isContains := contains(tt.args.channel, channels); !isContains {
				t.Errorf("expect channels contain %v", tt.args.channel)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		t          string
		conditions []string
	}
	tests := []struct {
		name           string
		args           args
		wantIsContains bool
	}{
		{"A", args{"a", []string{"a"}}, true},
		{"B", args{"b", []string{"a"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsContains := contains(tt.args.t, tt.args.conditions); gotIsContains != tt.wantIsContains {
				t.Errorf("contains() = %v, want %v", gotIsContains, tt.wantIsContains)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want *Banshee
	}{
		{"Test New - Pass", args{"A"}, &Banshee{"A", &chat.Message{
			Markdown:       true,
			UnfurlLinks:    false,
			UnfurlMedia:    false,
			ReplyBroadcast: false,
			Parse:          "none",
			Attachments:    []*chat.Attachment{}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
