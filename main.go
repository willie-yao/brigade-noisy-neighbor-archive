package main

import (
	"context"
	"log"
	"time"

	"github.com/brigadecore/brigade/sdk/v2"
	"github.com/brigadecore/brigade/sdk/v2/core"
	"github.com/willie-yao/brigade-noisy-neighbor/internal/signals"
	"github.com/willie-yao/brigade-noisy-neighbor/internal/version"
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

	noiseFrequency, err := noiseFrequency()
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(noiseFrequency)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := createEvent(ctx, apiClient)
			if err != nil {
				log.Println(err)
			}
		case <-ctx.Done():
			return
		}
	}

}

func createEvent(ctx context.Context, apiClient sdk.APIClient) error {
	_, err := apiClient.Core().Events().Create(
		ctx,
		core.Event{
			Source: "github.com/willie-yao/brigade-noisy-neighbor",
			Type:   "noise",
		},
	)
	if err != nil {
		return err
	}
	return nil
}
