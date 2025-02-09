package pokeapi

import (
	"net/http"
	"time"

	"github.com/BenSnaith/go-pokedex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient - return new http client to use with api
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
