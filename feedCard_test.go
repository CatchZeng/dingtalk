package dingtalk

import (
	"reflect"
	"testing"
)

func TestFeedCardMessage_ToByte(t *testing.T) {
	msg := NewFeedCardMessage()
	_, _ = msg.ToByte()
	if msg.MsgType != MsgTypeFeedCard {
		t.Errorf("FeedCardMessage.ToByte() type error")
	}
}

func TestNewFeedCardMessage(t *testing.T) {
	tests := []struct {
		name string
		want *FeedCardMessage
	}{
		{
			name: "Should return a FeedCardMessage instance",
			want: &FeedCardMessage{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFeedCardMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFeedCardMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFeedCardMessage_AppendLink(t *testing.T) {
	msg := NewFeedCardMessage()
	msg.AppendLink("title", "messageURL", "picURL")
	if len(msg.FeedCard.Links) != 1 {
		t.Errorf("The number of links after AppendLink should be 1")
	}

	msg.AppendLink("title2", "messageURL2", "picURL2")
	if len(msg.FeedCard.Links) != 2 {
		t.Errorf("The number of links after AppendLink should be 2")
	}
}
