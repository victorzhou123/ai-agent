package config

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/victorzhou123/ai-agent/agent"
	"github.com/victorzhou123/ai-agent/common/log"
)

var globalCfg *Config

type Config struct {
	Server Server       `json:"server" mapstructure:"server"`
	Log    log.Config   `json:"log" mapstructure:"log"`
	Agent  agent.Config `json:"agent" mapstructure:"agent"`
}

type Server struct {
	Port              int `json:"port" mapstructure:"port"`
	ReadTimeout       int `json:"read_timeout" mapstructure:"read_timeout"`               // unit Millisecond
	ReadHeaderTimeout int `json:"read_header_timeout" mapstructure:"read_header_timeout"` // unit Millisecond
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
