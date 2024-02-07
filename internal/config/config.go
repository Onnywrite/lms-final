package config

import (
	"encoding/json"
	"flag"
	"os"
)

const (
	configPathName = "CONFIG_PATH"
)

type Config struct {
	Env             string `json:"env"`
	GoroutinesCount int    `json:"goroutines_count"`
	Port            int    `json:"port"`
	DbConnect       string `json:"db_connection"`
}

func MustLoad() *Config {
	cfg, err := Load()
	if err != nil {
		panic(err)
	}
	return cfg
}

func Load() (*Config, error) {
	var path string

	flag.StringVar(&path, configPathName, "", "path to config file")
	flag.Parse()
	if path == "" {
		path = os.Getenv(configPathName)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
