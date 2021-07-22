package main

import (
	"context"
	"time"

	"github.com/brigadecore/brigade/sdk/v2"
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

func (m *noisyNeighbor) run(ctx context.Context) {
	ticker := time.NewTicker(m.noiseInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			m.createEvents()
		case <-ctx.Done():
			return
		}
	}
}

func (m *noisyNeighbor) createEvents() {

}
