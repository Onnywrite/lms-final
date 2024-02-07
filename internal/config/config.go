package config

type Config struct {
	DbConnect, Env        string
	GoroutinesCount, Port int
}

func MustLoad() *Config {
	// TODO: implement config.MustLoad
	return &Config{}
}
