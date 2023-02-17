package webhoook_poster

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	MsgTypeText     = "text"     //text 类型
	MsgTypeLink     = "link"     //link 类型
	MsgTypeMarkdown = "markdown" //markdown 类型
)

// Message 普通消息
type Message struct {
	Content  string   `validate:"required"`
	AtPerson []string `json:"atUserIds"`
	AtAll    bool     `json:"isAtAll"`
}

// Link 链接消息
type Link struct {
	Content    string   `json:"text" validate:"required"`       // 要发送的消息， 必填
	Title      string   `json:"title" validate:"required"`      // 标题， 必填
	ContentURL string   `json:"messageUrl" validate:"required"` // 点击消息跳转的URL 必填
	PictureURL string   `json:"picUrl"`                         // 图片 url
	AtPerson   []string `json:"atUserIds"`
	AtAll      bool     `json:"isAtAll"`
}

// Markdown markdown 类型
type Markdown struct {
	Content  string   `json:"text" validate:"required"`  // 要发送的消息， 必填
	Title    string   `json:"title" validate:"required"` // 标题， 必填
	AtPerson []string `json:"atUserIds"`
	AtAll    bool     `json:"isAtAll"`
}

// SimpleMessage push message
type SimpleMessage struct {
	Content string
	Title   string
}

type DingTalkPoster struct {
	WebhookURL string `json:"webhook_url"`
	Keyword    string `json:"keyword"`
}

func NewDingTalkPoster(webhook_url, keyword string) *DingTalkPoster {
	return &DingTalkPoster{
		WebhookURL: webhook_url,
		Keyword:    keyword,
	}
}

func (m *DingTalkPoster) Send(msg string) error {
	// 构建消息体
	message := map[string]interface{}{
		"msgtype": MsgTypeText,
		"text": map[string]string{
			"content": fmt.Sprintf("%s\n%s", msg, m.Keyword),
		},
	}

	// 将消息体编码为JSON格式
	jsonBytes, _ := json.Marshal(message)

	// 创建POST请求
	req, _ := http.NewRequest("POST", m.WebhookURL, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, _ := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		return errors.New("DingTalk message send failed")
	}

	defer resp.Body.Close()
	return nil
}
