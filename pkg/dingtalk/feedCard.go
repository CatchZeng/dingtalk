package dingtalk

import "encoding/json"

// FeedCardMessage feed message struct
type FeedCardMessage struct {
	MsgType  MsgType  `json:"msgtype"`
	FeedCard FeedCard `json:"feedCard"`
}

// FeedCard feedCard struct
type FeedCard struct {
	Links []FeedCardLink `json:"links"`
}

// FeedCardLink struct
type FeedCardLink struct {
	Title      string `json:"title"`
	PicURL     string `json:"picURL"`
	MessageURL string `json:"messageURL"`
}

// ToByte to byte
func (m *FeedCardMessage) ToByte() ([]byte, error) {
	m.MsgType = MsgTypeFeedCard
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
}

// NewFeedCardMessage new message
func NewFeedCardMessage() *FeedCardMessage {
	msg := FeedCardMessage{}
	return &msg
}

// AppendLink append link
func (m *FeedCardMessage) AppendLink(
	title string,
	messageURL string,
	picURL string) *FeedCardMessage {
	var link = FeedCardLink{
		Title:      title,
		MessageURL: messageURL,
		PicURL:     picURL,
	}
	m.FeedCard.Links = append(m.FeedCard.Links, link)
	return m
}
