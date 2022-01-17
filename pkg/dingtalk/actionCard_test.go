package dingtalk

import (
	"reflect"
	"testing"
)

func TestActionCardMessage_ToByte(t *testing.T) {
	msg := NewActionCardMessage()
	_, _ = msg.ToByte()
	if msg.MsgType != MsgTypeActionCard {
		t.Errorf("ActionCardMessage.ToByte() type error")
	}
}

func TestNewActionCardMessage(t *testing.T) {
	tests := []struct {
		name string
		want *ActionCardMessage
	}{
		{
			name: "Should return a ActionCardMessage instance",
			want: &ActionCardMessage{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewActionCardMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewActionCardMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActionCardMessage_SetOverallJump(t *testing.T) {
	got := NewActionCardMessage()
	got.SetOverallJump("title", "text", "singleTitle", "singleURL", "btnOrientation", "hideAvatar")

	card := ActionCard{
		Title:          "title",
		Text:           "text",
		SingleTitle:    "singleTitle",
		SingleURL:      "singleURL",
		BtnOrientation: "btnOrientation",
		HideAvatar:     "hideAvatar",
	}
	want := NewActionCardMessage()
	want.ActionCard = card

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SetOverallJump() = %v, want %v", got, want)
	}
}

func TestActionCardMessage_SetIndependentJump(t *testing.T) {
	got := NewActionCardMessage()
	got.SetIndependentJump("title", "text", []Btn{{
		Title:     "title",
		ActionURL: "actionURL",
	}}, "btnOrientation", "hideAvatar")

	card := ActionCard{
		Title: "title",
		Text:  "text",
		Btns: []Btn{{
			Title:     "title",
			ActionURL: "actionURL",
		}},
		BtnOrientation: "btnOrientation",
		HideAvatar:     "hideAvatar",
	}
	want := NewActionCardMessage()
	want.ActionCard = card

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SetIndependentJump() = %v, want %v", got, want)
	}
}
