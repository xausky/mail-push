package config

import (
	"errors"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

// ProviderConfig 定义单个提供商的配置
type ProviderConfig struct {
	EmailSuffix string `toml:"email_suffix"`
	SMTPHost    string `toml:"smtp_host"`
	SMTPPort    int    `toml:"smtp_port"`
}

// Config 定义读取 TOML 文件时的总体配置
type Config struct {
	Providers map[string]ProviderConfig `toml:"providers"`
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
