package message

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"text/template"

	"mail-push/config"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// TemplateMessage 表示基于模板的消息类型
type TemplateMessage struct {
	FormatName string                 `json:"format_name"`
	Data       map[string]interface{} `json:"data"`
}

// NewTemplateMessage 创建一个新的模板消息
func NewTemplateMessage(formatName string, jsonData []byte, formatConfig *config.FormatConfig) (*TemplateMessage, error) {
	// 根据配置的编码方式解析数据
	var data map[string]interface{}

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, fmt.Errorf("invalid JSON for template message: %w", err)
	}

	return &TemplateMessage{
		FormatName: formatName,
		Data:       data,
	}, nil
}

// ToHTML 使用模板将数据转换为 HTML 格式
func (msg *TemplateMessage) ToHTML() (title, content string, err error) {
	// 这里需要从外部传入格式配置，因为 ToHTML 方法签名不允许修改
	// 我们将在 server/handler.go 中处理模板渲染
	return "", "", errors.New("template rendering should be handled in server handler")
}

// RenderTemplate 使用指定的模板渲染数据
func RenderTemplate(tmplText string, data map[string]interface{}) (string, error) {
	tmpl, err := template.New("render").Funcs(template.FuncMap{"renderMarkdown": renderMarkdown}).Parse(tmplText)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

func renderMarkdown(markdown string) (string, error) {
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
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		return "", fmt.Errorf("failed to convert markdown to HTML: %w", err)
	}

	return buf.String(), nil
}
