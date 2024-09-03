package round_robin

import (
	"github.com/rugz007/go-balancer/types"

	"github.com/rugz007/go-balancer/internal/proxy"
	"github.com/rugz007/go-balancer/internal/util"
)

type RoundRobinBackend struct {
	proxy     *types.BackendProxy
	index     int
	isHealthy bool
}

type RoundRobin struct {
	types.Balancer
	backends []*RoundRobinBackend
	current  int
}

func CreateRoundRobin(config types.Config) *RoundRobin {
	// RoundRobin is a simple round-robin algorithm
	// that selects the next backend in the list
	// and returns it.

	roundRobin := &RoundRobin{
		backends: make([]*RoundRobinBackend, 0),
		current:  0,
	}

	for i, backend := range config.Backends {
		proxy := proxy.CreateProxyClient(backend)
		roundRobin.backends = append(roundRobin.backends, &RoundRobinBackend{
			proxy:     &proxy,
			index:     i,
			isHealthy: true,
		})
	}
	return roundRobin
}

// Compile time check to ensure that RoundRobin implements the Balancer interface
var _ types.Balancer = &RoundRobin{}

func (rr *RoundRobin) Next() *types.BackendProxy {
	// Next returns the next backend in the list
	// and updates the current index.

	backend := rr.backends[rr.current]
	rr.current = (rr.current + 1) % len(rr.backends)
	if !backend.isHealthy {
		return rr.Next()
	}
	return backend.proxy
}

// TODO: Implement HealthCheck method in a better way
func (rr *RoundRobin) HealthCheck() {
	// HealthCheck checks the health of all backends
	// and updates the isHealthy field.

	for _, backend := range rr.backends {
		// Perform health check
		// and update isHealthy field
		util.IsHostAlive(backend.proxy.Client, backend.proxy.Backend.Url)
	}
}
