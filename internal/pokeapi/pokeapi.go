package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
)


type LocationAreaRes struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationAreaRes, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaRes{}, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return LocationAreaRes{}, err
	}

	var locationAreas LocationAreaRes
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return LocationAreaRes{}, err
	}

	return locationAreas, nil
}

