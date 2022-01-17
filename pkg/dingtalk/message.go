package dingtalk

// MsgType message type enum
type MsgType string

const (
	// MsgTypeText text
	MsgTypeText MsgType = "text"
	// MsgTypeMarkdown markdown
	MsgTypeMarkdown MsgType = "markdown"
	// MsgTypeLink link
	MsgTypeLink MsgType = "link"
	// MsgTypeActionCard actionCard
	MsgTypeActionCard MsgType = "actionCard"
	// MsgTypeFeedCard feedCard
	MsgTypeFeedCard MsgType = "feedCard"
)

// Message interface
type Message interface {
	ToByte() ([]byte, error)
}

// At at struct
type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}
