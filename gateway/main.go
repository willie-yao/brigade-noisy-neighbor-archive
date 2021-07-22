package main

import (
	"context"
	"log"
	"time"

	"github.com/brigadecore/brigade/sdk/v2"
	"github.com/brigadecore/brigade/sdk/v2/core"
	"github.com/willie-yao/brigade-noisy-neighbor/gateway/internal/signals"
	"github.com/willie-yao/brigade-noisy-neighbor/gateway/internal/version"
)

func main() {
	log.Printf(
		"Starting Brigade Noisy Neighbor -- version %s -- commit %s",
		version.Version(),
		version.Commit(),
	)

	ctx := signals.Context()

	address, token, opts, err := apiClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	apiClient := sdk.NewAPIClient(address, token, &opts)

	noiseInterval, err := noiseFrequency()
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(noiseInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			createEvent(apiClient, ctx)
		case <-ctx.Done():
			return
		}
	}

}

func createEvent(apiClient sdk.APIClient, ctx context.Context) {
	_, err := apiClient.Core().Events().Create(
		ctx,
		core.Event{
			Source: "github.com/willie-yao/brigade-noisy-neighbor",
			Type:   "noise",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
