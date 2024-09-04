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
	Backends          []Backend
	Port              int
	Headers           map[string]string
	Algorithm         string
	DropRequestOnFail bool
	RetriesOnFail     int
}

type Balancer interface {
	Next() *BackendProxy
	HealthCheck()
	MakeRequest(ctx *fasthttp.RequestCtx) error
}

type BackendProxy struct {
	Backend Backend
	Client  *fasthttp.Client
}
