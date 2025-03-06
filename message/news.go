package message

import (
	"errors"
	"fmt"
)

type NewsMessage struct {
	MsgType string `json:"msgtype"`
	News    struct {
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			PicURL      string `json:"picurl"`
		} `json:"articles"`
	} `json:"news"`
}

// ToHTML 生成 News 类型的邮件 HTML 内容
func (msg *NewsMessage) ToHTML() (title, content string, err error) {
	if len(msg.News.Articles) == 0 {
		return "", "", errors.New("no articles in news message")
	}

	// 使用第一篇文章的标题作为邮件标题
	title = msg.News.Articles[0].Title

	// 生成 HTML 内容
	htmlContent := "<ul>"
	for _, article := range msg.News.Articles {
		htmlContent += fmt.Sprintf(
			"<li><a href='%s'><img src='%s' alt='%s'></a><p>%s</p></li>",
			article.URL, article.PicURL, article.Title, article.Description,
		)
	}
	htmlContent += "</ul>"

	return title, htmlContent, nil
}
