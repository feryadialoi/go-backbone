package boot

type Config struct {
	restPort       int
	grpcPort       int
	prometheusPort int
}

type ApplyConfig func(config *Config)

func DefaultConfig() *Config {
	return &Config{
		restPort: 8080,
		grpcPort: 9090,
	}
}

func RestPort(port int) ApplyConfig {
	return func(config *Config) {
		config.restPort = port
	}
}

func GRPCPort(port int) ApplyConfig {
	return func(config *Config) {
		config.grpcPort = port
	}
}

func PrometheusPort(port int) ApplyConfig {
	return func(config *Config) {
		config.prometheusPort = port
	}
}
