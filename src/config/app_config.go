package config

// AppConfig 应用配置
type AppConfig struct {
	Name string `mapstructure:"name"`
	Addr string `mapstructure:"addr"`
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}
