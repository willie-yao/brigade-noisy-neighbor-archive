package main

import (
	"log"
	"time"

	"github.com/brigadecore/brigade/sdk/v2"
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

	{
		address, token, opts, err := apiClientConfig()
		if err != nil {
			log.Fatal(err)
		}
		scrapeInterval, err := noiseDuration()
		if err != nil {
			log.Fatal(err)
		}
		newNoisyNeighbor(
			sdk.NewAPIClient(address, token, &opts),
			time.Duration(scrapeInterval),
		).run(ctx)
	}

}
