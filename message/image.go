package message

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type ImageMessage struct {
	MsgType string `json:"msgtype"`
	Image   struct {
		Base64 string `json:"base64"`
		MD5    string `json:"md5"`
	} `json:"image"`
}

// ToHTML 生成 Image 类型的邮件 HTML 内容
func (msg *ImageMessage) ToHTML() (title, content string, err error) {
	if msg.Image.Base64 == "" || msg.Image.MD5 == "" {
		return "", "", errors.New("missing image data or MD5")
	}

	// 验证 MD5
	hash := md5.Sum([]byte(msg.Image.Base64))
	if hex.EncodeToString(hash[:]) != msg.Image.MD5 {
		return "", "", errors.New("MD5 mismatch for image")
	}

	// 使用当前时间作为标题
	title = fmt.Sprintf("Image Message - %s", time.Now().Format("2006-01-02 15:04:05"))
	content = fmt.Sprintf("<img src='data:image/png;base64,%s' alt='Image Message'/>", msg.Image.Base64)
	return title, content, nil
}
