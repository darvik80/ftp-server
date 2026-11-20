package config

import (
	"github.com/spf13/viper"
)

type LoggerConfig struct {
	Level        string `json:"level"`
	Console      string `json:"console"`
	ConsoleLevel string `json:"console_level"`
	File         string `json:"file"`
	FileLevel    string `json:"file_level"`
	MaxSizeMB    int    `json:"max_size_mb"`
	MaxBackups   int    `json:"max_backups"`
	AddSource    bool   `json:"add_source"`
	Color        string `json:"color"`
}

type ServerConfig struct {
	PublicAddress string `json:"public_address"`
	Address       string `json:"address"`
	Port          int    `json:"port"`
	ReadBuffer    int    `json:"read_buffer"`
	IdleTimeout   int    `json:"idle_timeout"`
	MaxSessions   int    `json:"max_sessions"`
}

type Config struct {
	Server ServerConfig `json:"server"`
	Logger LoggerConfig `json:"logger"`
}

func setDefaults() {
	// FTP Server
	viper.SetDefault("server.public_address", "")
	viper.SetDefault("server.address", "127.0.0.1")
	viper.SetDefault("server.port", 2121)
	viper.SetDefault("server.read_buffer", 1024)
	viper.SetDefault("server.idle_timeout", 60)
	viper.SetDefault("server.max_sessions", 10)

	// Logger
	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.console", "spring")
	viper.SetDefault("logger.console_level", "debug")
	viper.SetDefault("logger.color", "always")
}

func LoadConfig() (*Config, error) {
	viper.SetConfigType("json")
	viper.SetConfigName("config")

	// Add search paths to find the file
	viper.AddConfigPath("/etc/ftp-server-go/")
	viper.AddConfigPath("$HOME/.ftp-server-go")
	viper.AddConfigPath(".")

	setDefaults()
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
