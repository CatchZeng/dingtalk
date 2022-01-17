package dingtalk

import (
	"reflect"
	"testing"
)

func TestMarkdownMessage_ToByte(t *testing.T) {
	msg := NewMarkdownMessage()
	_, _ = msg.ToByte()
	if msg.MsgType != MsgTypeMarkdown {
		t.Errorf("MarkdownMessage.ToByte() type error")
	}
}

func TestNewMarkdownMessage(t *testing.T) {
	tests := []struct {
		name string
		want *MarkdownMessage
	}{
		{
			name: "Should return a MarkdownMessage instance",
			want: &MarkdownMessage{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMarkdownMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMarkdownMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarkdownMessage_SetMarkdown(t *testing.T) {
	got := NewMarkdownMessage()
	got.SetMarkdown("title", "text")

	want := NewMarkdownMessage()
	want.Markdown = Markdown{
		Title: "title",
		Text:  "text",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SetMarkdown() = %v, want %v", got, want)
	}
}

func TestMarkdownMessage_SetAt(t *testing.T) {
	got := NewMarkdownMessage()
	got.SetAt([]string{"atMobiles"}, false)

	want := NewMarkdownMessage()
	want.At = At{
		AtMobiles: []string{"atMobiles"},
		IsAtAll:   false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SetMarkdown() = %v, want %v", got, want)
	}
}
