package types

import (
	"time"

	"github.com/valyala/fasthttp"
)

var Algorithms = []string{"round-robin"}

type Backend struct {
	Url                 string
	HealthCheckPath     string
	MaxConns            int
	MaxIdleConnDuration time.Duration
	MaxConnDuration     time.Duration
}

type Config struct {
	Backends  []Backend
	Port      int
	Headers   map[string]string
	Algorithm string
}

type Balancer interface {
	Next() *BackendProxy
	HealthCheck()
}

type BackendProxy struct {
	Backend Backend
	Client  *fasthttp.Client
}
