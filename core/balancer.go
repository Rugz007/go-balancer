package core

import (
	"fmt"

	round_robin "github.com/rugz007/go-balancer/internal/algorithms/round-robin"
	types "github.com/rugz007/go-balancer/types"
	"github.com/valyala/fasthttp"
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

func CreateBalancerServer(config *types.Config) {
	balancer := CreateBalancer(config)
	// Create a request handler
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		err := balancer.MakeRequest(ctx)
		if err != nil {
			if config.DropRequestOnFail {
				ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			} else {
				retries := config.RetriesOnFail
				for retries > 0 {
					err = balancer.MakeRequest(ctx)
					if err == nil {
						break
					}
					retries--
				}
				if err != nil {
					ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
				} else {
					fmt.Fprintf(ctx, "Retried request successfully after %d retries", config.RetriesOnFail-retries)
				}
			}
		}
	}

	if err := fasthttp.ListenAndServe(fmt.Sprintf(":%d", config.Port), requestHandler); err != nil {
		fmt.Printf("Error in ListenAndServe: %s\n", err)
	}
}
