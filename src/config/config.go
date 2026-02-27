package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 全局配置结构体
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
}

// GlobalConfig 全局配置实例
var GlobalConfig *Config

// Init 初始化配置
func Init() {
	// 设置配置文件路径和格式
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./src/config")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取配置文件失败: %v", err))
	}

	// 解析到结构体
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		panic(fmt.Sprintf("解析配置失败: %v", err))
	}
}
