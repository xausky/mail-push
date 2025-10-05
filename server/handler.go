package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"mail-push/config"
	"mail-push/mailer"
	"mail-push/message"
	"strings"
	"time"

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
func sendResponse(c *fiber.Ctx, code int, template string, error string) error {
	dataMap := make(map[string]interface{})
	dataMap["ip__"] = GetRealIP(c)
	dataMap["time__"] = time.Now().Format("2006-01-02 15:04:05")
	if error != "" {
		dataMap["error__"] = error
	}
	if template == "" {
		return c.Status(code).JSON(dataMap)
	}
	response, err := message.RenderTemplate(template, dataMap)
	if err != nil {
		dataMap["reason__"] = err.Error()
		return c.Status(code).JSON(dataMap)
	}
	c.Response().Header.SetContentType("application/json")
	return c.Status(code).SendString(response)
}

// FormatWebhookHandler 处理动态格式的 webhook 请求并发送邮件
func FormatWebhookHandler(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 获取格式名称
		formatName := c.Params("format")
		if formatName == "" {
			return sendResponse(c, fiber.StatusBadRequest, "", "format parameter is required")
		}

		// 查找对应的格式配置
		formatCfg, err := cfg.GetFormat(formatName)
		if err != nil {
			return sendResponse(c, fiber.StatusNotFound, "", fmt.Sprintf("format '%s' not found in configuration", formatName))
		}

		// --- Base64 参数解析 ---
		// 从 URL Path 解码 Base64
		data := c.Params("data")
		decoded, err := Base64Decode(data)
		if err != nil {
			return sendResponse(c, fiber.StatusBadRequest, formatCfg.ResponseTemplate, "invalid Base64 parameter")
		}

		// 拆分 Base64 解码后的数据；格式: provider|username|password
		parts := strings.Split(decoded, "|")
		if len(parts) != 3 {
			return sendResponse(c, fiber.StatusBadRequest, formatCfg.ResponseTemplate, "invalid parameter format, expected 'provider|username|password'")
		}
		provider, username, password := parts[0], parts[1], parts[2]

		// 查找对应的提供商配置
		providerCfg, err := cfg.GetProvider(provider)
		if err != nil {
			return sendResponse(c, fiber.StatusNotFound, formatCfg.ResponseTemplate, fmt.Sprintf("provider '%s' not found in configuration", provider))
		}

		// 动态拼接发件人邮箱地址
		emailAddress := username + providerCfg.EmailSuffix

		// --- 解析请求体数据 ---
		jsonData := c.Body()

		// --- 使用模板渲染 HTML 内容 ---
		// 解析 JSON 数据为 map
		var dataMap map[string]interface{}
		if err := json.Unmarshal(jsonData, &dataMap); err != nil {
			return sendResponse(c, fiber.StatusBadRequest, formatCfg.ResponseTemplate, fmt.Sprintf("Invalid JSON body: %v", err))
		}

		// 使用模板渲染标题和内容
		content, err := message.RenderTemplate(formatCfg.ContentTemplate, dataMap)
		if err != nil {
			return sendResponse(c, fiber.StatusInternalServerError, formatCfg.ResponseTemplate, fmt.Sprintf("Failed to render content template: %v", err))
		}

		dataMap["summary__"] = message.StripHTML(content)

		title, err := message.RenderTemplate(formatCfg.TitleTemplate, dataMap)
		if err != nil {
			return sendResponse(c, fiber.StatusInternalServerError, formatCfg.ResponseTemplate, fmt.Sprintf("Failed to render title template: %v", err))
		}

		// 包装 HTML 内容
		content = message.WrapHTML(content)

		// 记录发送日志
		log.Printf("Sending format '%s' to %s, title: %s", formatName, emailAddress, title)

		// --- 发送邮件 ---
		m := mailer.NewMailer(providerCfg.SMTPHost, providerCfg.SMTPPort, emailAddress, password)
		if err := m.SendEmail(emailAddress, title, content); err != nil {
			return sendResponse(c, fiber.StatusInternalServerError, formatCfg.ResponseTemplate, fmt.Sprintf("Failed to send email: %v", err))
		}

		// --- 成功响应 ---
		return sendResponse(c, fiber.StatusOK, formatCfg.ResponseTemplate, "")
	}
}

// Base64Decode 用于解码一个 Base64 编码的字符串
func Base64Decode(encoded string) (string, error) {
	// 使用标准库的 Base64 解码
	decoded, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("base64 decode error: %w", err)
	}
	return string(decoded), nil
}

// ProvidersHandler 处理查询所有 providers 的请求
func ProvidersHandler(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		providers := make([]map[string]interface{}, 0, len(cfg.Providers))
		for key, provider := range cfg.Providers {
			providers = append(providers, map[string]interface{}{
				"key":  key,
				"name": provider.Name,
			})
		}
		return c.JSON(providers)
	}
}

// FormatsHandler 处理查询所有 formats 的请求
func FormatsHandler(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		formats := make([]map[string]interface{}, 0, len(cfg.Formats))
		for key, format := range cfg.Formats {
			formatInfo := map[string]interface{}{
				"key":          key,
				"name":         format.Name,
				"example":      format.Example,
				"document_url": format.DocumentURL,
			}
			formats = append(formats, formatInfo)
		}
		return c.JSON(formats)
	}
}
