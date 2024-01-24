package main

import (
	"time"

	"github.com/Joad/pokedexcli/internal/pokeapi"
)

func main() {
	config := &Config{
		client:  pokeapi.NewClient(5*time.Second, 5*time.Minute),
		pokedex: make(map[string]pokeapi.Pokemon),
	}
	startRepl(config)
}
