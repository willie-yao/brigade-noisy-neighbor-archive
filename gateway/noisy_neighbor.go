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
	ticker := time.NewTicker(n.noiseInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			n.createEvent(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (n *noisyNeighbor) createEvent(ctx context.Context) {
	_, err := n.apiClient.Core().Events().Create(
		ctx,
		core.Event{
			Source: "https://github.com/willie-yao/brigade-noisy-neighbor/",
			Type:   "noise",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
