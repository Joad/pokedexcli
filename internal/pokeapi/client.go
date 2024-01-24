package pokeapi

import (
	"net/http"
	"time"

	"github.com/Joad/pokedexcli/internal/pokecache"
)

const (
	baseUrl      = "https://pokeapi.co/api/v2"
	locationAreo = "/location-area"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
	}
}
