package config

import "time"

const (
	mainConfigPathName = "MAIN_CONFIG_PATH"
)

type MainConfig struct {
	Env             string        `json:"env"`
	Port            int           `json:"port"`
	StaticDir       string        `json:"static_dir"`
	OrchestratorURI string        `json:"orchestrator_uri"`
	ShutdownTimeout time.Duration `json:"shutdown_timeout"`
	LogsDir         string        `json:"logs_dir"`
}

func MustLoadMain() *MainConfig {
	cfg, err := LoadMain()
	if err != nil {
		panic(err)
	}
	return cfg
}

func LoadMain() (*MainConfig, error) {
	cfg, err := load[MainConfig](mainConfigPathName)
	if err != nil {
		return nil, err
	}
	cfg.ShutdownTimeout *= time.Second
	return cfg, nil
}
