package config

import (
	"encoding/json"
	"flag"
	"os"
)

func load[TConfig any](pathName string) (*TConfig, error) {
	var path string

	flag.StringVar(&path, pathName, "", "path to config file")
	flag.Parse()
	if path == "" {
		path = os.Getenv(pathName)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg TConfig
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
