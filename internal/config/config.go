package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	Source string `mapstructure:"source"`
}

// LoadConfig 加载配置，通过 APP_ENV 环境变量切换环境 (dev/test/prod)，默认为 dev
func LoadConfig(path string) (config Config, err error) {
	// 读取 APP_ENV 环境变量，默认 dev
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	// 设置 Viper 搜索路径
	viper.AddConfigPath(path)
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("../../configs")
	viper.AddConfigPath("../../../configs")

	viper.SetConfigType("yaml")

	// 1. 先尝试加载环境专用配置 config.{env}.yaml
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	if err = viper.ReadInConfig(); err != nil {
		// 2. 如果环境配置不存在，回退到 config.yaml
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SetConfigName("config")
			if err = viper.ReadInConfig(); err != nil {
				return config, fmt.Errorf("failed to read config file: %w", err)
			}
		} else {
			return config, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	// 环境变量覆盖（优先级最高）
	viper.AutomaticEnv()

	// 解析到结构体
	err = viper.Unmarshal(&config)
	return
}
