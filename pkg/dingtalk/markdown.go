package dingtalk

import "encoding/json"

// MarkdownMessage markdown message struct
type MarkdownMessage struct {
	MsgType  MsgType  `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
}

// Markdown markdown struct
type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// ToByte to byte
func (m *MarkdownMessage) ToByte() ([]byte, error) {
	m.MsgType = MsgTypeMarkdown
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
}

// NewMarkdownMessage new message
func NewMarkdownMessage() *MarkdownMessage {
	msg := MarkdownMessage{}
	return &msg
}

// SetMarkdown set markdown
func (m *MarkdownMessage) SetMarkdown(title string, text string) *MarkdownMessage {
	m.Markdown = Markdown{
		Title: title,
		Text:  text,
	}
	return m
}

// SetAt set at
func (m *MarkdownMessage) SetAt(atMobiles []string, isAtAll bool) *MarkdownMessage {
	m.At = At{
		AtMobiles: atMobiles,
		IsAtAll:   isAtAll,
	}
	return m
}
