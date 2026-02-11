package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/harryyu02/pokedex/internal/pokecache"
)

type PokeApiClient struct {
	Cache *pokecache.Cache
}

type LocationAreaRes struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func NewClient(interval time.Duration) *PokeApiClient {
	return &PokeApiClient{
		Cache: pokecache.NewCache(interval),
	}
}

func (p *PokeApiClient)GetLocationAreas(url string) (LocationAreaRes, error) {
	bytes := []byte{}
	cached, ok := p.Cache.Get(url)
	if ok {
		bytes = cached
	} else {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreaRes{}, err
		}

		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return LocationAreaRes{}, err
		}
		bytes = body
	}

	var locationAreas LocationAreaRes
	err := json.Unmarshal(bytes, &locationAreas)
	if err != nil {
		return LocationAreaRes{}, err
	}

	p.Cache.Add(url, bytes)

	return locationAreas, nil
}

