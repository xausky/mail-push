package message

import (
	"fmt"
	"regexp"
	"strings"
)

// StripHTML 从文本中移除所有HTML标签,返回纯文本内容
func StripHTML(input string) string {
	// 移除所有HTML标签的正则表达式
	re := regexp.MustCompile(`<[^>]*>`)
	text := re.ReplaceAllString(input, "")

	// 移除多余的空白字符
	text = strings.TrimSpace(text)
	text = strings.Join(strings.Fields(text), " ")

	return text
}

// DefaultCSS 返回一些基础的CSS样式,用于美化HTML邮件内容
func DefaultCSS() string {
	return `<style>
		body {
			line-height: 1.6;
			color: #333;
			margin: 0;
			padding: 0;
			background-color: #ffffff;
		}
		.header {
			background-color:rgb(91, 148, 201);
			padding: 20px;
			text-align: center;
			margin-bottom: 30px;
		}
		.content {
			max-width: 800px;
			margin: 0 auto;
			padding: 20px;
		}
		.footer {
			background-color: #1976d2;
			color: #ffffff;
			text-align: center;
			padding: 10px;
			margin-top: 30px;
		}
		.footer a {
			color: #ffffff;
			text-decoration: none;
		}
		.footer a:hover {
			text-decoration: underline;
		}
		h1, h2, h3, h4, h5, h6 {
			color: #2c3e50;
			margin-top: 1.5em;
			margin-bottom: 0.5em;
		}
		p {
			margin: 1em 0;
		}
		img {
			max-width: 100%;
			height: auto;
			border-radius: 4px;
		}
		a {
			color: #3498db;
			text-decoration: none;
		}
		a:hover {
			text-decoration: underline;
		}
		blockquote {
			border-left: 4px solid #e5e5e5;
			margin: 1em 0;
			padding-left: 1em;
			color: #666;
		}
		code {
			background-color: #f8f9fa;
			padding: 2px 4px;
			border-radius: 3px;
			font-family: monospace;
		}
		pre {
			background-color: #f8f9fa;
			padding: 1em;
			border-radius: 4px;
			overflow-x: auto;
		}
	</style>`
}

// WrapHTML 将 HTML 内容包装成完整的 HTML 文档，包含默认的 CSS 样式
func WrapHTML(body string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
%s
</head>
<body>
<div class="header">
</div>
<div class="content">
%s
</div>
<div class="footer">
Powered by <a href="https://mp.xac.one">Mail Push</a>
</div>
</body>
</html>`, DefaultCSS(), body)
}
