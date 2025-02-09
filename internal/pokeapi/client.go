package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

// NewClient - return new http client to use with api
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
