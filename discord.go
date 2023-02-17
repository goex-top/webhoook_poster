package webhoook_poster

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type DiscordPoster struct {
	WebhookURL string `json:"webhook_url"`
}

func NewDiscordPoster(webhook_url string) *DiscordPoster {
	return &DiscordPoster{
		WebhookURL: webhook_url,
	}
}

func (m *DiscordPoster) Send(msg string) error {
	// 构建消息体
	message := map[string]string{
		"content": msg,
	}

	// 将消息体编码为JSON格式
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// 创建POST请求
	req, err := http.NewRequest("POST", m.WebhookURL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return errors.New("Discord message send failed")
	}

	defer resp.Body.Close()
	return nil
}
