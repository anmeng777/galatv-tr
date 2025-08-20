package galatvtr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// sendBarkMessage 发送Bark消息
func PushMsgBark(token, title, content, status string) error {
	// Bark API URL格式: https://api.day.app/{token}/{title}/{content}
	// barkURL := fmt.Sprintf("https://api.day.app/%s", req.Token)

	sound := "anticipate"
	if status == "success" {
		sound = "alarm"
	}

	barkURL := fmt.Sprintf("https://api.day.app/%s/%s/%s?sound=%s",
		token, url.QueryEscape(title), url.QueryEscape(content), sound)

	_, err := http.Get(barkURL)
	if err != nil {
		// 如果bark推送失败，可以在这里记录日志，但不影响主要的错误响应
		fmt.Printf("Bark推送失败: %v\n", err)
		return fmt.Errorf("failed to marshal bark request: %v", err)
	}

	return nil
}

// sendDingDingMessage 发送钉钉消息
func PushMsgDingding(token, title, content, name string) error {
	// 钉钉机器人webhook URL
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", token)

	type DingDingRequest struct {
		MsgType string `json:"msgtype"`
		Text    struct {
			Content string `json:"content"`
		} `json:"text"`
	}

	dingReq := DingDingRequest{
		MsgType: "text",
	}
	dingReq.Text.Content = fmt.Sprintf("【%s通知】%s：%s", name, title, content)

	jsonData, err := json.Marshal(dingReq)
	if err != nil {
		return fmt.Errorf("failed to marshal dingding request: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send dingding request: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("dingding API returned status %d: %s", resp.StatusCode, string(body))
	}

	fmt.Println(string(body))

	return nil
}
