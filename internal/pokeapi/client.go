package pokeapi

import (
	"github.com/bencuci/pokedex/internal/pokecache"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	pokecache  *pokecache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	cache := pokecache.NewCache(cacheInterval)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokecache: cache,
	}
}
