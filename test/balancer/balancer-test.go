package main

import (
	"fmt"

	"github.com/rugz007/go-balancer/core"
	"github.com/rugz007/go-balancer/types"
)

func main() {
	fmt.Print("Starting balancer server\n")
	fmt.Print("Listening on port 8080\n")
	config := &types.Config{
		Algorithm: "round-robin",
		Backends: []types.Backend{
			{
				Url: "http://localhost:8001",
			},
			{
				Url: "http://localhost:8002",
			},
			{
				Url: "http://localhost:8003",
			},
		},
		DropRequestOnFail: false,
		Port:              8080,
	}
	core.CreateBalancerServer(config)
}
