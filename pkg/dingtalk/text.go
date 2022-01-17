package dingtalk

import "encoding/json"

// TextMessage text message struct
type TextMessage struct {
	MsgType MsgType `json:"msgtype"`
	Text    Text    `json:"text"`
	At      At      `json:"at"`
}

// Text text struct
type Text struct {
	Content string `json:"content"`
}

// ToByte to byte
func (m *TextMessage) ToByte() ([]byte, error) {
	m.MsgType = MsgTypeText
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
}

// NewTextMessage new message
func NewTextMessage() *TextMessage {
	msg := TextMessage{}
	return &msg
}

// SetContent set content
func (m *TextMessage) SetContent(content string) *TextMessage {
	m.Text = Text{
		Content: content,
	}
	return m
}

// SetAt set at
func (m *TextMessage) SetAt(atMobiles []string, isAtAll bool) *TextMessage {
	m.At = At{
		AtMobiles: atMobiles,
		IsAtAll:   isAtAll,
	}
	return m
}
