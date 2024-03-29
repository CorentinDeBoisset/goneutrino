package server

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`

	TrustedProxies []string `yaml:"trusted_proxies"`
}

func LoadConfig(configPath string) (*Config, error) {
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read the configuration file: %w", err)
	}

	serverConfig := new(Config)
	err = yaml.Unmarshal(configBytes, serverConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the yaml from the configuration file: %w", err)
	}

	if serverConfig.Port == 0 {
		serverConfig.Port = 2475
	}
	if len(serverConfig.Host) == 0 {
		serverConfig.Host = "127.0.0.1"
	}

	if serverConfig.TrustedProxies == nil {
		return nil, fmt.Errorf("invalid value for config property \"trusted_proxies\"")
	}

	return serverConfig, nil
}
