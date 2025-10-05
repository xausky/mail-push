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

	// 打印所有 providers 和 formats 信息
	cfg.PrintProvidersAndFormats()

	// 初始化 Fiber 应用
	app := fiber.New()

	// 静态文件服务
	app.Static("/", "./build")

	// 动态格式 webhook 接口
	app.Post("/:format/:data", server.FormatWebhookHandler(cfg))

	// 查询 providers 接口
	app.Get("/providers", server.ProvidersHandler(cfg))

	// 查询 formats 接口
	app.Get("/formats", server.FormatsHandler(cfg))

	// 启动服务器
	log.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}
