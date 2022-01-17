package dingtalk

import "encoding/json"

// LinkMessage link message struct
type LinkMessage struct {
	MsgType MsgType `json:"msgtype"`
	Link    Link    `json:"link"`
}

// Link link struct
type Link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	PicURL     string `json:"picUrl"`
	MessageURL string `json:"messageUrl"`
}

// ToByte to byte
func (m *LinkMessage) ToByte() ([]byte, error) {
	m.MsgType = MsgTypeLink
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
}

// NewLinkMessage new message
func NewLinkMessage() *LinkMessage {
	msg := LinkMessage{}
	return &msg
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
