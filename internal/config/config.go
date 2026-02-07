package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppPort   string `mapstructure:"APP_PORT"`
	AppEnv    string `mapstructure:"APP_ENV"`

	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
	DBMaxOpenConns int    `mapstructure:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns int    `mapstructure:"DB_MAX_IDLE_CONNS"`
}

func Load() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // đọc từ environment variables nếu có

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
		// .env không bắt buộc nếu dùng env var
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Default values
	if cfg.AppPort == "" {
		cfg.AppPort = "8080"
	}
	if cfg.AppEnv == "" {
		cfg.AppEnv = "development"
	}

	return &cfg, nil
}