package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl(config *Config) {
	commands := cliCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]

		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if command.nArgs != len(args) {
			fmt.Println("Wrong number of arguments")
		}
		err := command.callback(config, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}
