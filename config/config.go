package config

import (
	"os"
	"time"

	"github.com/evalphobia/ipregistry-go/client"
)

const (
	defaultEnvAPIKey = "IPREGISTRY_APIKEY"
)

var envAPIKey string

func init() {
	envAPIKey = os.Getenv(defaultEnvAPIKey)
}

// Config contains parameters for Ipregistry.
type Config struct {
	APIKey string

	Debug   bool
	Timeout time.Duration
}

// Client generates client.Client from Config data.
func (c Config) Client() (*client.Client, error) {
	cli := client.New()
	cli.SetOption(client.Option{
		Debug:   c.Debug,
		Timeout: c.Timeout,
	})
	cli.SetAPIKey(c.getAPIKey())
	return cli, nil
}

// getAPIKey returns API Key from Config data or Environment variables.
func (c Config) getAPIKey() string {
	apiKey := os.Getenv(defaultEnvAPIKey)
	if apiKey != "" {
		return apiKey
	}
	return c.APIKey
}
