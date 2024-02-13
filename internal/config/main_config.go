package config

import (
	"encoding/json"
	"flag"
	"os"
)

const (
	configPathName = "MAIN_CONFIG_PATH"
)

type MainConfig struct {
	Env       string `json:"env"`
	Port      int    `json:"port"`
	StaticDir string `json:"static_dir"`
}

func MustLoadMain() *MainConfig {
	cfg, err := LoadMain()
	if err != nil {
		panic(err)
	}
	return cfg
}

func LoadMain() (*MainConfig, error) {
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
	var cfg MainConfig
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
