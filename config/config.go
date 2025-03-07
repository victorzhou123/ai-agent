package config

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/victorzhou123/ai-agent/common/log"
)

var globalCfg *Config

type Config struct {
	Server Server     `json:"server" mapstructure:"server"`
	Log    log.Config `json:"log" mapstructure:"log"`
	Client Client     `json:"client" mapstructure:"client"`
	Role   Role       `json:"prompt" mapstructure:"prompt"`
}

type Server struct {
	Port              int `json:"port"`
	ReadTimeout       int `json:"read_timeout"`        // unit Millisecond
	ReadHeaderTimeout int `json:"read_header_timeout"` // unit Millisecond
}

type Client struct {
	Ollama Ollama `json:"llm" mapstructure:"llm"`
}

type Ollama struct {
	Host string `json:"host" mapstructure:"host"`
	Port string `json:"port" mapstructure:"port"`
}

type Role struct {
	Abstract Setting `json:"abstract" mapstructure:"abstract"`
	Polish   Setting `json:"polish" mapstructure:"polish"`
}

type Setting struct {
	Model  string `json:"model" mapstructure:"model"`
	Prompt string `json:"prompt" mapstructure:"prompt"`
}

func GetGlobalConfig() *Config {
	return globalCfg
}

func LoadConfig(path string) error {
	cfg := new(Config)
	viperConfig := viper.New()

	dirPath, fileName := filepath.Split(path)
	names := strings.Split(fileName, ".")
	if len(names) < 2 {
		return errors.New("parse file path error")
	}
	viperConfig.AddConfigPath(dirPath)
	viperConfig.SetConfigName(names[0])
	viperConfig.SetConfigType(names[1])

	viperConfig.WatchConfig()
	viperConfig.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("config changing...")

		if err := viperConfig.ReadInConfig(); err != nil {
			log.Errorf("read config error: %v", err)
			return
		}

		if err := viperConfig.Unmarshal(cfg); err != nil {
			log.Errorf("unmarshal config error: %v", err)
			return
		}

		globalCfg = cfg
	})

	if err := viperConfig.ReadInConfig(); err != nil {
		return fmt.Errorf("read config: %w", err)
	}

	if err := viperConfig.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unmarshal config: %w", err)
	}

	globalCfg = cfg

	return nil
}
