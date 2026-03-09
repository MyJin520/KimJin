package config

type JwtConfig struct {
	Key    string `mapstructure:"key"`
	Expire int    `mapstructure:"expire"`
}
