package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var CONF Config

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Port          int                 `yaml:"port"`
	Logging       ServerLoggingConfig `yaml:"logging"`
	RBACModelConf string              `yaml:"rbacModelConf"`
}

type ServerLoggingConfig struct {
	IsFile bool   `yaml:"file"`
	Level  string `yaml:"level"`
	Path   string `yaml:"path"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Migrate  bool   `yaml:"migrate"`
}

func NewConfig(configFilePath string) (*Config, error) {
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "read config file failed")
	}

	var conf Config
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return nil, errors.Wrap(err, "parse config file failed")
	}

	return &conf, nil
}
