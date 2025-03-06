package message

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)

// Message 是所有消息类型的通用接口
type Message interface {
	ToHTML() (title, content string, err error) // 将消息转换成邮件的 HTML
}

// 定义消息工厂方法，用于根据 msgtype 实例化具体类型的消息
func NewMessage(msgType string, jsonData []byte) (Message, error) {
	switch msgType {
	case "text":
		var msg TextMessage
		if err := json.Unmarshal(jsonData, &msg); err != nil {
			return nil, fmt.Errorf("invalid JSON for text message: %w", err)
		}
		return &msg, nil
	case "markdown":
		var msg MarkdownMessage
		if err := json.Unmarshal(jsonData, &msg); err != nil {
			return nil, fmt.Errorf("invalid JSON for markdown message: %w", err)
		}
		return &msg, nil
	case "image":
		var msg ImageMessage
		if err := json.Unmarshal(jsonData, &msg); err != nil {
			return nil, fmt.Errorf("invalid JSON for image message: %w", err)
		}
		return &msg, nil
	case "news":
		var msg NewsMessage
		if err := json.Unmarshal(jsonData, &msg); err != nil {
			return nil, fmt.Errorf("invalid JSON for news message: %w", err)
		}
		return &msg, nil
	default:
		return nil, errors.New("unknown message type")
	}
}

// Base64Decode 用于解码一个 Base64 编码的字符串
func Base64Decode(encoded string) (string, error) {
	// 使用标准库的 Base64 解码
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("base64 decode error: %w", err)
	}
	return string(decoded), nil
}
