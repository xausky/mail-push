# Mail Push

Mail Push 是一个基于 SMTP 协议的消息推送服务，通过发送邮件到用户自己的邮箱来实现消息推送功能。它完全兼容企业微信群机器人的 webhook 格式，让你可以轻松地将企业微信机器人的消息推送到自己的邮箱。

## 特性

- 🚀 完全兼容企业微信群机器人的 webhook 格式
- 📧 支持多种邮件服务商（QQ邮箱、163邮箱等）
- 💬 支持多种消息类型（文本、Markdown、图片、图文）
- 🔒 无需后端存储，所有认证信息编码在 URL 中
- 🎨 优雅的 Markdown 渲染
- 📱 支持微信 QQ 邮箱提醒，实时接收通知

## 快速开始

1. 访问 [Mail Push 在线服务](https://mp.xac.one)
2. 填写邮箱服务商（如：qq）和账号信息
3. 生成专属 Webhook 地址
4. 使用生成的地址发送消息

### 示例请求

```bash
# 文本消息 curl 示例
curl -X POST "https://you.deploy.domain/send/YOUR-BASE64-DATA" \
     -H "Content-Type: application/json" \
     -d '{
    "msgtype": "text",
    "text": {
        "content": "这是一条文本消息"
    }
}'
```

## 安全性说明

- 所有认证信息使用 Base64 编码存储在 URL 中
- 服务器不存储任何用户信息
- 建议自行部署服务以确保信息安全

## 技术栈

- Go + Fiber：高性能 Web 框架
- Goldmark：Markdown 渲染
- Gomail：邮件发送
- 前端：原生 JavaScript + HTML + CSS

## 致谢

- 本项目的大部分编码工作由 [Cursor](https://cursor.sh/) 完成

## 开源协议

[MIT License](LICENSE)

## 免责声明

本项目仅供学习和参考，作者不对使用本项目产生的任何后果负责。如果有高可用或者高隐私需求，请自行搭建服务。
