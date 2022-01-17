package dingtalk

import "encoding/json"

// ActionCardMessage struct
type ActionCardMessage struct {
	MsgType    MsgType    `json:"msgtype"`
	ActionCard ActionCard `json:"actionCard"`
}

// ActionCard actionCard struct
type ActionCard struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
	Btns           []Btn  `json:"btns"`
	BtnOrientation string `json:"btnOrientation"`
	HideAvatar     string `json:"hideAvatar"`
}

// Btn struct
type Btn struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

// ToByte to byte
func (m *ActionCardMessage) ToByte() ([]byte, error) {
	m.MsgType = MsgTypeActionCard
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
}

// NewActionCardMessage new message
func NewActionCardMessage() *ActionCardMessage {
	msg := ActionCardMessage{}
	return &msg
}

// SetOverallJump set overall jump actionCard
func (m *ActionCardMessage) SetOverallJump(
	title string,
	text string,
	singleTitle string,
	singleURL string,
	btnOrientation string,
	hideAvatar string) *ActionCardMessage {
	m.ActionCard = ActionCard{
		Title:          title,
		Text:           text,
		SingleTitle:    singleTitle,
		SingleURL:      singleURL,
		BtnOrientation: btnOrientation,
		HideAvatar:     hideAvatar,
	}
	return m
}

// SetIndependentJump set independent jump actionCard
func (m *ActionCardMessage) SetIndependentJump(
	title string,
	text string,
	btns []Btn,
	btnOrientation string,
	hideAvatar string) *ActionCardMessage {
	m.ActionCard = ActionCard{
		Title:          title,
		Text:           text,
		Btns:           btns,
		BtnOrientation: btnOrientation,
		HideAvatar:     hideAvatar,
	}
	return m
}
