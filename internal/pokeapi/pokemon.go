package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Species                struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (client *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + name
	body, found := client.cache.Get(url)
	if !found {
		res, err := client.httpClient.Get(url)
		if err != nil {
			return Pokemon{}, err
		}

		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return Pokemon{}, err
		}
		client.cache.Add(url, body)
	}

	result := Pokemon{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}, err
	}
	return result, nil
}
