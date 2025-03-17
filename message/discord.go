package message

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// DiscordColor 是一个自定义类型，用于处理字符串或整数类型的颜色值
type DiscordColor int

// UnmarshalJSON 实现自定义的 JSON 解析
func (c *DiscordColor) UnmarshalJSON(data []byte) error {
	// 移除引号
	s := string(data)
	s = strings.Trim(s, `"`)

	// 如果是空值，设置为默认颜色
	if s == "null" || s == "" {
		*c = 0
		return nil
	}

	// 尝试将字符串转换为整数
	val, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid color value: %s", s)
	}

	*c = DiscordColor(val)
	return nil
}

// DiscordMessage 表示 Discord webhook 消息格式
type DiscordMessage struct {
	Content   string         `json:"content"`
	Username  string         `json:"username"`
	AvatarURL string         `json:"avatar_url"`
	Embeds    []DiscordEmbed `json:"embeds"`
}

// DiscordEmbed 表示 Discord 嵌入式内容
type DiscordEmbed struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	URL         string              `json:"url"`
	Color       DiscordColor        `json:"color"`
	Timestamp   string              `json:"timestamp"`
	Footer      DiscordEmbedFooter  `json:"footer"`
	Image       DiscordEmbedImage   `json:"image"`
	Thumbnail   DiscordEmbedImage   `json:"thumbnail"`
	Author      DiscordEmbedAuthor  `json:"author"`
	Fields      []DiscordEmbedField `json:"fields"`
}

// DiscordEmbedFooter 表示 Discord 嵌入式内容的页脚
type DiscordEmbedFooter struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

// DiscordEmbedImage 表示 Discord 嵌入式内容的图片
type DiscordEmbedImage struct {
	URL string `json:"url"`
}

// DiscordEmbedAuthor 表示 Discord 嵌入式内容的作者
type DiscordEmbedAuthor struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

// DiscordEmbedField 表示 Discord 嵌入式内容的字段
type DiscordEmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

// ToHTML 将 Discord 消息转换为 HTML 格式
func (msg *DiscordMessage) ToHTML() (title, content string, err error) {
	var htmlContent strings.Builder

	// 设置标题，优先使用第一个 embed 的标题，如果没有则使用 content 的第一行
	if len(msg.Content) > 0 {
		lines := strings.Split(strings.TrimSpace(msg.Content), "\n")
		title = strings.TrimSpace(lines[0])

		// 添加主要内容
		htmlContent.WriteString(fmt.Sprintf("<p>%s</p>", strings.ReplaceAll(msg.Content, "\n", "<br>")))
	}

	// 处理 embeds
	if len(msg.Embeds) > 0 {
		// 如果标题为空，使用第一个 embed 的标题
		if title == "" && msg.Embeds[0].Title != "" {
			title = msg.Embeds[0].Title
		}

		// 添加所有 embeds
		for _, embed := range msg.Embeds {
			htmlContent.WriteString("<div style='border-left: 4px solid #")

			// 处理颜色
			if embed.Color != 0 {
				htmlContent.WriteString(fmt.Sprintf("%06x", embed.Color))
			} else {
				htmlContent.WriteString("7289DA") // Discord 默认颜色
			}
			htmlContent.WriteString("; padding-left: 10px; margin: 10px 0;'>\n")

			// 添加标题
			if embed.Title != "" {
				if embed.URL != "" {
					htmlContent.WriteString(fmt.Sprintf("<h3><a href='%s'>%s</a></h3>\n", embed.URL, embed.Title))
				} else {
					htmlContent.WriteString(fmt.Sprintf("<h3>%s</h3>\n", embed.Title))
				}
			}

			// 添加作者
			if embed.Author.Name != "" {
				htmlContent.WriteString("<div style='display: flex; align-items: center; margin-bottom: 10px;'>\n")
				if embed.Author.IconURL != "" {
					htmlContent.WriteString(fmt.Sprintf("<img src='%s' style='width: 24px; height: 24px; border-radius: 50%%; margin-right: 8px;'>\n", embed.Author.IconURL))
				}
				if embed.Author.URL != "" {
					htmlContent.WriteString(fmt.Sprintf("<a href='%s'>%s</a>\n", embed.Author.URL, embed.Author.Name))
				} else {
					htmlContent.WriteString(fmt.Sprintf("<span>%s</span>\n", embed.Author.Name))
				}
				htmlContent.WriteString("</div>\n")
			}

			// 添加描述
			if embed.Description != "" {
				htmlContent.WriteString(fmt.Sprintf("<p>%s</p>\n", strings.ReplaceAll(embed.Description, "\n", "<br>")))
			}

			// 添加字段
			if len(embed.Fields) > 0 {
				htmlContent.WriteString("<div style='display: grid; grid-template-columns: repeat(auto-fill, minmax(200px, 1fr)); gap: 10px; margin: 10px 0;'>\n")
				for _, field := range embed.Fields {
					if field.Inline {
						htmlContent.WriteString("<div style='margin-bottom: 10px;'>\n")
					} else {
						htmlContent.WriteString("<div style='margin-bottom: 10px; grid-column: 1 / -1;'>\n")
					}
					htmlContent.WriteString(fmt.Sprintf("<h4>%s</h4>\n", field.Name))
					htmlContent.WriteString(fmt.Sprintf("<p>%s</p>\n", strings.ReplaceAll(field.Value, "\n", "<br>")))
					htmlContent.WriteString("</div>\n")
				}
				htmlContent.WriteString("</div>\n")
			}

			// 添加图片
			if embed.Image.URL != "" {
				htmlContent.WriteString(fmt.Sprintf("<img src='%s' style='max-width: 100%%; margin: 10px 0;'>\n", embed.Image.URL))
			}

			// 添加缩略图
			if embed.Thumbnail.URL != "" {
				htmlContent.WriteString(fmt.Sprintf("<img src='%s' style='float: right; max-width: 80px; max-height: 80px; margin: 0 0 10px 10px;'>\n", embed.Thumbnail.URL))
			}

			// 添加页脚
			if embed.Footer.Text != "" {
				htmlContent.WriteString("<div style='display: flex; align-items: center; margin-top: 10px; font-size: 0.8em; color: #72767d;'>\n")
				if embed.Footer.IconURL != "" {
					htmlContent.WriteString(fmt.Sprintf("<img src='%s' style='width: 20px; height: 20px; border-radius: 50%%; margin-right: 8px;'>\n", embed.Footer.IconURL))
				}
				htmlContent.WriteString(fmt.Sprintf("<span>%s</span>\n", embed.Footer.Text))
				htmlContent.WriteString("</div>\n")
			}

			// 添加时间戳
			if embed.Timestamp != "" {
				t, err := time.Parse(time.RFC3339, embed.Timestamp)
				if err == nil {
					htmlContent.WriteString(fmt.Sprintf("<div style='font-size: 0.8em; color: #72767d; margin-top: 5px;'>%s</div>\n", t.Format("2006-01-02 15:04:05")))
				}
			}

			htmlContent.WriteString("</div>\n")
		}
	}

	// 如果没有内容，返回错误
	if htmlContent.Len() == 0 {
		return "", "", errors.New("empty discord message content")
	}

	// 如果标题为空，使用默认标题
	if title == "" {
		if msg.Username != "" {
			title = fmt.Sprintf("Message from %s", msg.Username)
		} else {
			title = "Discord Message"
		}
	}

	return title, htmlContent.String(), nil
}
