package config

type Config struct {
	DbToken         string
	GoroutinesCount int
	CalculatorPort  int
	RestfulPort     int
}

func MustLoad() *Config {
	return &Config{}
}
