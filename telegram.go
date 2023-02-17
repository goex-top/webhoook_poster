package webhoook_poster

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type TelegramPoster struct {
	WebhookURL string `json:"webhook_url"`
	ChatID     int64  `json:"chat_id"`
}

func NewTelegramPoster(webhook_url string, chat_id int64) *TelegramPoster {
	return &TelegramPoster{
		WebhookURL: webhook_url,
		ChatID:     chat_id,
	}
}

func (m *TelegramPoster) Send(msg string) error {
	// 构建消息体
	message := map[string]interface{}{
		"chat_id": m.ChatID,
		"text":    msg,
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

	if resp.StatusCode != http.StatusOK {
		return errors.New("Telegram message send failed")
	}

	defer resp.Body.Close()
	return nil
}
