package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type Encounters struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (client *Client) GetEncounters(name string) (Encounters, error) {
	url := baseUrl + locationAreo + "/" + name
	body, found := client.cache.Get(url)
	if !found {
		res, err := client.httpClient.Get(url)
		if err != nil {
			return Encounters{}, err
		}

		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return Encounters{}, err
		}
		client.cache.Add(url, body)
	}

	result := Encounters{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		return Encounters{}, err
	}
	return result, nil
}
