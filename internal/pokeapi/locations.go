package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type LocationResult struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (client *Client) GetMaps(url string) (LocationResult, error) {
	if url == "" {
		url = baseUrl + locationAreo
	}
	body, found := client.cache.Get(url)
	if !found {
		res, err := client.httpClient.Get(url)
		if err != nil {
			return LocationResult{}, err
		}

		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return LocationResult{}, err
		}
		client.cache.Add(url, body)
	}

	result := LocationResult{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		return LocationResult{}, err
	}
	return result, nil
}
