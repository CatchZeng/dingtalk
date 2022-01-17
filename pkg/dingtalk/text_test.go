package dingtalk

import (
	"reflect"
	"testing"
)

func TestTextMessage_ToByte(t *testing.T) {
	msg := NewTextMessage()
	_, _ = msg.ToByte()
	if msg.MsgType != MsgTypeText {
		t.Errorf("TextMessage.ToByte() type error")
	}
}

func TestNewTextMessage(t *testing.T) {
	tests := []struct {
		name string
		want *TextMessage
	}{
		{
			name: "Should return a TextMessage instance",
			want: &TextMessage{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTextMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTextMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextMessage_SetContent(t *testing.T) {
	got := NewTextMessage()
	got.SetContent("content")

	want := NewTextMessage()
	want.Text = Text{
		Content: "content",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SetContent() = %v, want %v", got, want)
	}
}

func TestTextMessage_SetAt(t *testing.T) {
	got := NewTextMessage()
	got.SetAt([]string{"atMobiles"}, false)

	want := NewTextMessage()
	want.At = At{
		AtMobiles: []string{"atMobiles"},
		IsAtAll:   false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SetMarkdown() = %v, want %v", got, want)
	}
}
