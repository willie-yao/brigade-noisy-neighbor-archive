package main

import (
	"time"

	"github.com/brigadecore/brigade/sdk/v2/restmachinery"
	"github.com/willie-yao/brigade-noisy-neighbor/gateway/internal/os"
)

// apiClientConfig populates the Brigade SDK's APIClientOptions from
// environment variables.
func apiClientConfig() (string, string, restmachinery.APIClientOptions, error) {
	opts := restmachinery.APIClientOptions{}
	address, err := os.GetRequiredEnvVar("API_ADDRESS")
	if err != nil {
		return address, "", opts, err
	}
	token, err := os.GetRequiredEnvVar("API_TOKEN")
	if err != nil {
		return address, token, opts, err
	}
	opts.AllowInsecureConnections, err =
		os.GetBoolFromEnvVar("API_IGNORE_CERT_WARNINGS", false)
	return address, token, opts, err
}

func scrapeDuration() (time.Duration, error) {
	return os.GetDurationFromEnvVar("PROM_SCRAPE_INTERVAL", 5*time.Second)
}
