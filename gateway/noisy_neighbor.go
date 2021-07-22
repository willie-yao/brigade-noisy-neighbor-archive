package main

import (
	"context"
	"log"
	"time"

	"github.com/brigadecore/brigade/sdk/v2"
	"github.com/brigadecore/brigade/sdk/v2/core"
)

type noisyNeighbor struct {
	apiClient     sdk.APIClient
	noiseInterval time.Duration
}

func newNoisyNeighbor(
	apiClient sdk.APIClient,
	noiseInterval time.Duration,
) *noisyNeighbor {
	return &noisyNeighbor{
		apiClient:     apiClient,
		noiseInterval: noiseInterval,
	}
}

func (n *noisyNeighbor) run(ctx context.Context) {
	log.Println("Running noisy neighbor")
	ticker := time.NewTicker(n.noiseInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			n.createEvents(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (n *noisyNeighbor) createEvents(ctx context.Context) {
	event := core.Event{
		Source: "brigade.sh/cli",
		Type:   "exec",
	}
	n.apiClient.Core().Events().Create(
		ctx,
		event,
	)
}
