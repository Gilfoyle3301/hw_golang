package main

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

type LoggerConf struct {
	Level string
}

type DbConf struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint64 `yaml:"port"`
	Name     string `yaml:"name"`
}
type ServerConf struct {
	Port string `yaml:"port"`
}
type Config struct {
	Logger LoggerConf `yaml:"logger"`
	Db     DbConf     `yaml:"db"`
	Server ServerConf `yaml:"server"`
}

func NewConfig(configFilePath string) (*Config, error) {
	var config Config
	configReader, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("reading conf file %s fail: %w", configFilePath, err)
	}

	if err := yaml.Unmarshal(configReader, &config); err != nil {
		return nil, fmt.Errorf("Unmarshal conf fail: %w", err)
	}

	return &config, nil
}
