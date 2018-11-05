package banshee

import (
	"reflect"
	"testing"
)

func init() {
	myChannels := []string{"a", "ab", "abc", "b", "bc", "c"}
	for _, channel := range myChannels {
		RegistChannel(channel, RealWebhookURL)

	}
}

// const RealWebhookURL = "Set a real slack webhook URL here before running test"
const RealWebhookURL = "https://hooks.slack.com/services/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"

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

func TestNew(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want *Banshee
	}{
		{"A", args{"a"}, &Banshee{EXACT, "a"}},
		{"B", args{"a"}, &Banshee{FUZZY, "a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.pattern)
			if tt.name == "B" {
				got.SetPublishModeFuzzy()
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBanshee_getPublishMode(t *testing.T) {
	tests := []struct {
		name string
		b    *Banshee
		want PublishMode
	}{
		{"A", &Banshee{EXACT, "a"}, EXACT},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.getPublishMode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Banshee.getPublishMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBanshee_SetPublishModeExact(t *testing.T) {
	tests := []struct {
		name string
		b    *Banshee
	}{
		{"A", &Banshee{FUZZY, "a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.SetPublishModeExact()
			got := tt.b.getPublishMode()
			if tt.b.publishMode != EXACT {
				t.Errorf("expact search type is %v, want %v", got, EXACT)
			}
		})
	}
}

func TestBanshee_SetPublishModeFuzzy(t *testing.T) {
	tests := []struct {
		name string
		b    *Banshee
	}{
		{"A", &Banshee{EXACT, "a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.SetPublishModeFuzzy()
			got := tt.b.getPublishMode()
			if tt.b.publishMode != FUZZY {
				t.Errorf("expact search type is %v, want %v", got, FUZZY)
			}
		})
	}
}

func TestBanshee_Publish(t *testing.T) {
	type fields struct {
		publishMode PublishMode
		pattern     string
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"A", fields{EXACT, "a"}, args{"test publish exact a"}, false},
		{"B", fields{FUZZY, "a"}, args{"test publish fuzzy a"}, false},
		{"C", fields{EXACT, "d"}, args{"test publish exact d"}, true},
		{"D", fields{FUZZY, "d"}, args{"test publish fuzzy d"}, true},
		{"E", fields{FUZZY, "null"}, args{"test publish fuzzy null"}, true},
	}
	for _, tt := range tests {
		if tt.name == "E" {
			appendChannel("null")
		}
		t.Run(tt.name, func(t *testing.T) {
			b := &Banshee{
				publishMode: tt.fields.publishMode,
				pattern:     tt.fields.pattern,
			}
			if err := b.Publish(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Banshee.Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fuzzyFindChannel(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"A", args{"a"}, []string{"a", "ab", "abc"}},
		{"B", args{"b"}, []string{"ab", "abc", "b", "bc"}},
		{"C", args{"c"}, []string{"abc", "bc", "c"}},
		{"E", args{"ab"}, []string{"abc", "ab"}},
		{"F", args{"ba"}, []string{}},
		{"G", args{"bc"}, []string{"abc", "bc"}},
		{"H", args{"BC"}, []string{"abc", "bc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var matches []string
			var want []string

			gots := fuzzyFindChannel(tt.args.pattern)

			for _, got := range gots {
				matches = append(matches, got.Str)
			}

			for _, v := range tt.want {
				want = append(want, v)
			}

			for _, matche := range matches {
				for i := 0; i < len(want); i++ {
					if matche == want[i] {
						want = append(want[:i], want[i+1:]...)
						i--
					}
				}
			}

			if len(want) > 0 {
				t.Fatalf("fuzzyFindChannel() = %v, want %v", matches, tt.want)
			}
		})
	}
}
