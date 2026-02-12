package pokeapi

import (
	"encoding/json"
	"fmt"
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

type PokemonInAreaRes struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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

func (p *PokeApiClient)GetPokemonInArea(url string) (PokemonInAreaRes, error) {
	bytes := []byte{}
	cached, ok := p.Cache.Get(url)
	if ok {
		bytes = cached
	} else {
		res, err := http.Get(url)
		if err != nil {
			return PokemonInAreaRes{}, err
		}

		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return PokemonInAreaRes{}, err
		}
		if string(body) == "Not Found" {
			return PokemonInAreaRes{}, fmt.Errorf("no pokemon found in given area")
		}
		bytes = body
	}

	var pokemonInArea PokemonInAreaRes
	err := json.Unmarshal(bytes, &pokemonInArea)
	if err != nil {
		return PokemonInAreaRes{}, err
	}

	p.Cache.Add(url, bytes)

	return pokemonInArea, nil
}
