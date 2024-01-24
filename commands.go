package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/Joad/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	nArgs       int
	callback    func(*Config, ...string) error
}

type Config struct {
	next     string
	previous string
	client   pokeapi.Client
	pokedex  map[string]pokeapi.Pokemon
}

func cliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "help",
			description: "Exit the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display a list of available maps",
			callback:    displayMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous available maps",
			callback:    displayMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore the given area (show the pokemon in it)",
			nArgs:       1,
			callback:    explore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch the pokemon",
			nArgs:       1,
			callback:    catch,
		},
		"inspect": {
			name:        "inspect",
			description: "See details about the pokemon",
			nArgs:       1,
			callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See a list of caught pokemon",
			callback:    pokedex,
		},
	}
}

func commandHelp(config *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	commands := cliCommands()
	for key, value := range commands {
		fmt.Printf("%v: %v\n", key, value.description)
	}
	fmt.Println()
	return nil
}

func commandExit(config *Config, args ...string) error {
	os.Exit(0)
	return nil
}

func displayMap(config *Config, args ...string) error {
	err := getAndPrintMaps(config.next, config)
	if err != nil {
		return err
	}
	return nil
}

func displayMapB(config *Config, args ...string) error {
	if config.previous != "" {
		err := getAndPrintMaps(config.previous, config)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("No previous maps available")
	}
	return nil
}

func getAndPrintMaps(url string, config *Config) error {
	result, err := config.client.GetMaps(url)
	if err != nil {
		return err
	}
	config.next = result.Next
	config.previous = result.Previous

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func explore(config *Config, args ...string) error {
	result, err := config.client.GetEncounters(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, enc := range result.PokemonEncounters {
		fmt.Printf(" - %v\n", enc.Pokemon.Name)
	}
	return nil
}

func catch(config *Config, args ...string) error {
	name := args[0]
	result, err := config.client.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	res := rand.Intn(result.BaseExperience)
	fmt.Printf("probability: %.2f\n", 40.0/float64(result.BaseExperience))
	if res > 40 {
		fmt.Printf("%s escaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		config.pokedex[name] = result
	}
	return nil
}

func inspect(config *Config, args ...string) error {
	name := args[0]
	pokemon, found := config.pokedex[name]
	if !found {
		fmt.Println("You have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, tp := range pokemon.Types {
		fmt.Printf("  - %s\n", tp.Type.Name)
	}
	return nil
}

func pokedex(config *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range config.pokedex {
		fmt.Println("  - ", key)
	}
	return nil
}
