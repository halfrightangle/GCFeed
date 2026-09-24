package infraconfig

import (
	"errors"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

// 配置文件路径为 空
var ErrEmptyConfigPath = errors.New("config file path is empty")

// 读取配置失败
var ErrReadConfigFailed = errors.New("read config file failed")

// 配置为空
var ErrUnmarshalConfigFailed = errors.New("unmarshal config failed")

// LoadConfig 读取 YAML 配置文件，并反序列化为应用启动配置。
func LoadConfig(path string) (*Config, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return nil, ErrEmptyConfigPath
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, ErrReadConfigFailed
	}
	cfg := &Config{}
	if err := yaml.Unmarshal(content, cfg); err != nil {
		return nil, ErrUnmarshalConfigFailed
	}

	// 这里不做复杂默认值填充，让配置问题尽早在启动阶段暴露。
	return cfg, nil
}
