package banshee

import (
	"testing"
)

func init() {
	myChannels := []string{"a", "ab", "abc", "b", "bc", "c"}
	for _, channel := range myChannels {
		RegistChannel(channel, RealWebhookURL)
	}
}

const RealWebhookURL = "Set a real slack webhook URL here before running test"

func TestRegistChannel(t *testing.T) {
	type args struct {
		channel    string
		webhookURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test not OK", args{"test RegistChannel", "xxx"}, true},
		{"Test OK", args{"test RegistChannel", RealWebhookURL}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegistChannel(tt.args.channel, tt.args.webhookURL); (err != nil) != tt.wantErr {
				t.Errorf("RegistChannel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_registChannel(t *testing.T) {
	type args struct {
		channel    string
		webhookURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test not OK", args{"test registChannel", "xxx"}, true},
		{"Test OK", args{"test registChannel", RealWebhookURL}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := registChannel(tt.args.channel, tt.args.webhookURL); (err != nil) != tt.wantErr {
				t.Errorf("registChannel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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
