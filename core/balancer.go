package core

import (
	round_robin "github.com/rugz007/go-balancer/internal/algorithms/round-robin"
	types "github.com/rugz007/go-balancer/types"
)

func CreateBalancer(config *types.Config) types.Balancer {
	// GetBalancer returns the balancer instance
	// based on the algorithm specified in the config.
	// Currently, only round-robin is supported.
	// In the future, we can add more algorithms.
	if config.Algorithm == "round-robin" {
		return round_robin.CreateRoundRobin(*config)
	}
	return nil
}
