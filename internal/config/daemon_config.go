package config

const (
	daemonConfigPathName = "DAEMON_CONFIG_PATH"
)

type DaemonConfig struct {
	Env             string `json:"env"`
	OrchestratorURI string `json:"orchestrator_uri"`
	GoroutinesCount int    `json:"goroutines_count"`
}

func MustLoadDaemon() *DaemonConfig {
	cfg, err := LoadDaemon()
	if err != nil {
		panic(err)
	}
	return cfg
}

func LoadDaemon() (*DaemonConfig, error) {
	return load[DaemonConfig](daemonConfigPathName)
}
