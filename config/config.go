package config

import "time"

type Config struct {
	timeout        time.Duration
	grpcPort       int
	restPort       int
	prometheusPort int
}

func DefaultConfig() *Config {
	return &Config{
		timeout:        time.Second * 5,
		grpcPort:       9090,
		restPort:       8080,
		prometheusPort: 9800,
	}
}

func GetConfig() *Config {
	return DefaultConfig()
}

func (c *Config) Timeout() time.Duration {
	return c.timeout
}

func (c *Config) GRPCPort() int {
	return c.grpcPort
}

func (c *Config) RestPort() int {
	return c.restPort
}

func (c *Config) PrometheusPort() int {
	return c.prometheusPort
}
