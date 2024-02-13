package config

import "time"

const (
	orchestratorConfigPathName = "ORCH_CONFIG_PATH"
)

type OrchestratorConfig struct {
	Env             string        `json:"env"`
	Port            int           `json:"port"`
	DbConnection    string        `json:"db_connection"`
	ShutdownTimeout time.Duration `json:"shutdown_timeout"`
}

func MustLoadOrchestrator() *OrchestratorConfig {
	cfg, err := LoadOrchestrator()
	if err != nil {
		panic(err)
	}
	return cfg
}

func LoadOrchestrator() (*OrchestratorConfig, error) {
	cfg, err := load[OrchestratorConfig](orchestratorConfigPathName)
	if err != nil {
		return nil, err
	}
	cfg.ShutdownTimeout *= time.Second
	return cfg, nil
}
