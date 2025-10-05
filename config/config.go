package config

import (
	"errors"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

// ProviderConfig 定义单个提供商的配置
type ProviderConfig struct {
	Name        string `toml:"name"`
	EmailSuffix string `toml:"email_suffix"`
	SMTPHost    string `toml:"smtp_host"`
	SMTPPort    int    `toml:"smtp_port"`
}

// FormatConfig 定义单个 webhook 格式的配置
type FormatConfig struct {
	Name             string `toml:"name"`
	DocumentURL      string `toml:"docment_url"`
	TitleTemplate    string `toml:"title_template"`
	ContentTemplate  string `toml:"content_template"`
	ResponseTemplate string `toml:"response_template"`
}

// Config 定义读取 TOML 文件时的总体配置
type Config struct {
	Providers map[string]ProviderConfig `toml:"providers"`
	Formats   map[string]FormatConfig   `toml:"formats"`
}

// LoadConfiguration 读取 TOML 配置文件到 Config 结构体
func LoadConfiguration(filename string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(filename, &cfg); err != nil {
		return nil, err
	}
	log.Printf("Configuration loaded successfully from %s", filename)
	return &cfg, nil
}

// GetProvider 根据提供商名称获取提供商配置
func (cfg *Config) GetProvider(provider string) (*ProviderConfig, error) {
	p, exists := cfg.Providers[provider]
	if !exists {
		return nil, errors.New(fmt.Sprintf("Provider %s not found in configuration", provider))
	}
	return &p, nil
}

// GetFormat 根据格式名称获取格式配置
func (cfg *Config) GetFormat(name string) (*FormatConfig, error) {
	f, exists := cfg.Formats[name]
	if exists {
		return &f, nil
	}
	return nil, errors.New(fmt.Sprintf("Format %s not found in configuration", name))
}

// PrintProvidersAndFormats 打印所有 providers 和 formats 信息到日志
func (cfg *Config) PrintProvidersAndFormats() {
	log.Println("=== Providers ===")
	for key, provider := range cfg.Providers {
		log.Printf("Provider: %s, EmailSuffix: %s, SMTPHost: %s, SMTPPort: %d",
			key, provider.EmailSuffix, provider.SMTPHost, provider.SMTPPort)
	}

	log.Println("=== Formats ===")
	for key, format := range cfg.Formats {
		log.Printf("Format: %s, Name: %s, DocumentURL: %s",
			key, format.Name, getDocumentURL(format))
	}
}

// getDocumentURL 从 FormatConfig 中获取 document_url，如果不存在则返回空字符串
func getDocumentURL(format FormatConfig) string {
	return format.DocumentURL
}
