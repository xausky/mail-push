package main

import (
	"log"

	"mail-push/config"
	"mail-push/server"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfiguration("config.toml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// 初始化 Fiber 应用
	app := fiber.New()

	// 静态文件服务
	app.Static("/", ".svelte-kit/output/client")

	// 设置路由
	app.Post("/send/:data", server.MailHandler(cfg))

	// Discord webhook 兼容接口
	app.Post("/discord/:data", server.DiscordWebhookHandler(cfg))

	// 启动服务器
	log.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}
