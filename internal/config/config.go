package config

type Backend struct {
	Url             string
	HealthCheckPath string
	MaxConns        int
	MaxConnWait     int
}

type Config struct {
	Backends []Backend
	Port     int
	Headers  map[string]string
}
