package message

import (
	"errors"
	"fmt"
	"strings"
)

type TextMessage struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

// ToHTML 生成 Text 类型的邮件 HTML 内容
func (msg *TextMessage) ToHTML() (title, content string, err error) {
	lines := strings.Split(strings.TrimSpace(msg.Text.Content), "\n")
	if len(lines) == 0 {
		return "", "", errors.New("empty text content")
	}
	title = strings.TrimSpace(lines[0]) // 使用第一行作为标题
	content = fmt.Sprintf("<p>%s</p>", strings.ReplaceAll(msg.Text.Content, "\n", "<br>"))
	return title, content, nil
}
