package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var CONF Config

type Config struct {
	Port    int `yaml:"port"`
	Logging struct {
		Level string `yaml:"level"`
		Path  string `yaml:"path"`
	}
}

func LoadConfig(configFilePath string) {
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("read config file failed, err=%v", err)
	}

	err = yaml.Unmarshal(content, &CONF)
	if err != nil {
		log.Fatalf("parse config file failed, err=%v", err)
	}
}
