package message

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// MarkdownMessage 表示 Markdown 类型的消息
type MarkdownMessage struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"` // Markdown 原始文本内容
	} `json:"markdown"`
}

// ToHTML 将 Markdown 类型消息转换成 HTML 格式和标题
func (msg *MarkdownMessage) ToHTML() (title, content string, err error) {
	// 校验 Markdown 内容是否为空
	if strings.TrimSpace(msg.Markdown.Content) == "" {
		return "", "", errors.New("empty markdown content")
	}

	// 使用第一行作为邮件标题
	lines := strings.Split(strings.TrimSpace(msg.Markdown.Content), "\n")
	title = strings.TrimSpace(lines[0])

	// 将 Markdown 转换为 HTML
	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(), // 启用 HTML 标签渲染
			html.WithHardWraps(),
		),
	)
	if err := md.Convert([]byte(msg.Markdown.Content), &buf); err != nil {
		return "", "", fmt.Errorf("failed to convert markdown to HTML: %w", err)
	}

	content = buf.String()
	return title, content, nil
}
