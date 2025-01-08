package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB struct {
		User                   string `yaml:"user"`
		Password               string `yaml:"password"`
		DBName                 string `yaml:"dbname"`
		DBHost                 string `yaml:"dbhost"`
		DBPort                 string `yaml:"dbport"`
		SSLMode                string `yaml:"sslmode"`
		MaxOpenConnections     int    `yaml:"max_open_connections"`
		MaxIdleConnections     int    `yaml:"max_idle_connections"`
		ConnMaxLifetimeMinutes int    `yaml:"conn_max_lifetime_minutes"`
	} `yaml:"db"`
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return config, nil
}
