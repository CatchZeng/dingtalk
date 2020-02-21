package message

import "encoding/json"

// Doc：https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq #消息类型及数据格式

// MsgType message type enum
type MsgType string

const (
	// MsgTypeText text
	MsgTypeText MsgType = "text"
	// MsgTypeMarkdown markdown
	MsgTypeMarkdown MsgType = "markdown"
	// MsgTypeLink link
	MsgTypeLink MsgType = "link"
)

// Message interface
type Message interface {
	ToByte() ([]byte, error)
}

// TextMessage text message struct
type TextMessage struct {
	MsgType MsgType `json:"msgtype"`
	Text    Text    `json:"text"`
	At      At      `json:"at"`
}

// NewTextMessage new message
func NewTextMessage() *TextMessage {
	msg := TextMessage{}
	return &msg
}

// ToByte to byte
func (m *TextMessage) ToByte() ([]byte, error) {
	m.MsgType = MsgTypeText
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
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

// LinkMessage link message struct
type LinkMessage struct {
	MsgType MsgType `json:"msgtype"`
	Link    Link    `json:"link"`
}

// NewLinkMessage new message
func NewLinkMessage() *LinkMessage {
	msg := LinkMessage{}
	return &msg
}

// ToByte to byte
func (m *LinkMessage) ToByte() ([]byte, error) {
	m.MsgType = MsgTypeLink
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
}

// SetLink set link
func (m *LinkMessage) SetLink(
	title string,
	text string,
	picURL string,
	messageURL string) *LinkMessage {
	m.Link = Link{
		Title:      title,
		Text:       text,
		PicURL:     picURL,
		MessageURL: messageURL,
	}
	return m
}

// MarkdownMessage markdown message struct
type MarkdownMessage struct {
	MsgType  MsgType  `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
}

// NewMarkdownMessage new message
func NewMarkdownMessage() *MarkdownMessage {
	msg := MarkdownMessage{}
	return &msg
}

// ToByte to byte
func (m *MarkdownMessage) ToByte() ([]byte, error) {
	m.MsgType = MsgTypeMarkdown
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
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

// Text text struct
type Text struct {
	Content string `json:"content"`
}

// Markdown markdown struct
type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// Link link struct
type Link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	PicURL     string `json:"picUrl"`
	MessageURL string `json:"messageUrl"`
}

// At at struct
type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}
