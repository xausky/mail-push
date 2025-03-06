package server

import (
	"fmt"
	"strings"
	"time"

	"mail-push/config"
	"mail-push/mailer"
	"mail-push/message"

	"github.com/gofiber/fiber/v2"
)

// 获取真实客户端 IP，支持代理场景
func GetRealIP(c *fiber.Ctx) string {
	// 优先读取 X-Forwarded-For
	xForwardedFor := c.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		// X-Forwarded-For 可能有多个值，用逗号分隔，取第一个
		return strings.Split(xForwardedFor, ",")[0]
	}
	// 如果没有 X-Forwarded-For，则读取 RemoteIP
	return c.IP()
}

// 统一的错误 JSON 返回
func errorResponse(c *fiber.Ctx, code int, errcode int, errmsg string) error {
	ip := GetRealIP(c)
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	errorMsg := fmt.Sprintf("%s, hint: [%s], from ip: %s, more info at https://open.work.weixin.qq.com/devtool/query?e=%d", errmsg, timestamp, ip, errcode)
	return c.Status(code).JSON(fiber.Map{
		"errcode": errcode,
		"errmsg":  errorMsg,
	})
}

// MailHandler 处理消息并发送邮件
func MailHandler(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// --- Base64 参数解析 ---
		// 从 URL Path 解码 Base64
		data := c.Params("data")
		decoded, err := message.Base64Decode(data)
		if err != nil {
			return errorResponse(c, fiber.StatusBadRequest, 40008, "invalid Base64 parameter")
		}

		// 拆分 Base64 解码后的数据；格式: provider|username|password
		parts := strings.Split(decoded, "|")
		if len(parts) != 3 {
			return errorResponse(c, fiber.StatusBadRequest, 40009, "invalid parameter format, expected 'provider|username|password'")
		}
		provider, username, password := parts[0], parts[1], parts[2]

		// 查找对应的提供商配置
		providerCfg, err := cfg.GetProvider(provider)
		if err != nil {
			return errorResponse(c, fiber.StatusNotFound, 40010, fmt.Sprintf("provider '%s' not found in configuration, please contact xausky@163.com to add new provider", provider))
		}

		// 动态拼接发件人邮箱地址
		emailAddress := username + providerCfg.EmailSuffix

		// --- 请求体解析 ---
		// msgtype 用于区分消息类型 (text, markdown, image, news)
		var body struct {
			MsgType string `json:"msgtype"`
		}
		if err := c.BodyParser(&body); err != nil {
			return errorResponse(c, fiber.StatusBadRequest, 40001, "Invalid JSON body")
		}

		// 使用工厂方法加载对应消息类型
		msg, err := message.NewMessage(body.MsgType, c.Body())
		if err != nil {
			return errorResponse(c, fiber.StatusBadRequest, 40002, fmt.Sprintf("Invalid message type: %v", err))
		}

		// --- HTML 内容生成 ---
		title, content, err := msg.ToHTML()
		if err != nil {
			return errorResponse(c, fiber.StatusInternalServerError, 40003, fmt.Sprintf("Failed to generate HTML content: %v", err))
		}
		content = message.WrapHTML(content)
		title = message.StripHTML(title)

		// --- 发送邮件 ---
		m := mailer.NewMailer(providerCfg.SMTPHost, providerCfg.SMTPPort, emailAddress, password)
		if err := m.SendEmail(emailAddress, title, content); err != nil {
			return errorResponse(c, fiber.StatusInternalServerError, 40004, fmt.Sprintf("Failed to send email: %v", err))
		}

		// --- 成功响应 ---
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"errcode": 0,
			"errmsg":  "ok",
		})
	}
}
