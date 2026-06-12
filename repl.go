package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tomgoodo/pokedexcli/internal/pokeapi"
)

type config struct {
	next     *string
	previous *string
	client   pokeapi.PokeClient
}

type cliCommands struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommands {

	return map[string]cliCommands{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows Pokemon Locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Shows previous Pokemon Locations",
			callback:    commandMapB,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	cfg := newConfig()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		cmd, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		cmd.callback(cfg)
	}
}

func newConfig() *config {
	return &config{
		client:   pokeapi.NewPokeClient(),
		next:     nil,
		previous: nil,
	}
}
