package webhoook_poster

type WebhookPoster interface {
	Send(msg string) error
}

type PostType int

const (
	PostType_Dingding PostType = 1 + iota
	PostType_Discord
	PostType_Feishu
	PostType_Telegram
)

type Param struct {
	WebhookURL string `json:"webhook_url"`
	Keyword    string `json:"keyword"`
	ChatID     int64  `json:"chat_id"`
}

func New(posterType PostType, param Param) WebhookPoster {
	switch posterType {
	case PostType_Dingding:
		return NewDingTalkPoster(param.WebhookURL, param.Keyword)
	case PostType_Discord:
		return NewDiscordPoster(param.WebhookURL)
	case PostType_Feishu:
		return NewFeishuPoster(param.WebhookURL)
	case PostType_Telegram:
		return NewTelegramPoster(param.WebhookURL, param.ChatID)
	}
	return nil
}
