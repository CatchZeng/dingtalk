package dingtalk

import (
	"reflect"
	"testing"
)

func TestLinkMessage_ToByte(t *testing.T) {
	msg := NewLinkMessage()
	_, _ = msg.ToByte()
	if msg.MsgType != MsgTypeLink {
		t.Errorf("LinkMessage.ToByte() type error")
	}
}

func TestNewLinkMessage(t *testing.T) {
	tests := []struct {
		name string
		want *LinkMessage
	}{
		{
			name: "Should return a LinkMessage instance",
			want: &LinkMessage{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkMessage_SetLink(t *testing.T) {
	got := NewLinkMessage()
	got.SetLink("title", "text", "picURL", "messageURL")

	want := NewLinkMessage()
	want.Link = Link{
		Title:      "title",
		Text:       "text",
		PicURL:     "picURL",
		MessageURL: "messageURL",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SetLink() = %v, want %v", got, want)
	}
}
