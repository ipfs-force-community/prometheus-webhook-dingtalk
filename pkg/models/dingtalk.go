package models

type DingTalkNotificationResponse struct {
	ErrorMessage string `json:"errmsg"`
	ErrorCode    int    `json:"errcode"`

	// lark response fields
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DingTalkNotification struct {
	MessageType string                          `json:"msgtype"`
	Text        *DingTalkNotificationText       `json:"text,omitempty"`
	Link        *DingTalkNotificationLink       `json:"link,omitempty"`
	Markdown    *DingTalkNotificationMarkdown   `json:"markdown,omitempty"`
	ActionCard  *DingTalkNotificationActionCard `json:"actionCard,omitempty"`
	At          *DingTalkNotificationAt         `json:"at,omitempty"`
}

type DingTalkNotificationText struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type DingTalkNotificationLink struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageURL string `json:"messageUrl"`
	PictureURL string `json:"picUrl"`
}

type DingTalkNotificationMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type DingTalkNotificationAt struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

type DingTalkNotificationActionCard struct {
	Title             string                       `json:"title"`
	Text              string                       `json:"text"`
	HideAvatar        string                       `json:"hideAvatar"`
	ButtonOrientation string                       `json:"btnOrientation"`
	Buttons           []DingTalkNotificationButton `json:"btns,omitempty"`
	SingleTitle       string                       `json:"singleTitle,omitempty"`
	SingleURL         string                       `json:"singleURL"`
}

type DingTalkNotificationButton struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

////////// wechat ////////////

type WeChatNotification struct {
	MessageType string                          `json:"msgtype"`
	Text        *DingTalkNotificationText       `json:"text,omitempty"`
	Link        *DingTalkNotificationLink       `json:"link,omitempty"`
	Markdown    *WeChatNotificationMarkdown     `json:"markdown,omitempty"`
	ActionCard  *DingTalkNotificationActionCard `json:"actionCard,omitempty"`
	At          *DingTalkNotificationAt         `json:"at,omitempty"`
}

type WeChatNotificationMarkdown struct {
	Content string `json:"content"`
}

// //////// lark ////////////
// https://open.larksuite.com/document/client-docs/bot-v3/add-custom-bot
type LarkNotification struct {
	MessageType string                  `json:"msg_type"`
	Card        LarkNotificationCard    `json:"card"`
	At          *DingTalkNotificationAt `json:"at,omitempty"`
}

type LarkNotificationCard struct {
	Elements []LarkNotificationElement `json:"elements"`
}

type LarkNotificationElement struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}
