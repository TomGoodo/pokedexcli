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
	pokedex  map[string]pokeapi.Pokemon
}

type cliCommands struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"explore": {
			name:        "explore",
			description: "See what pokemon are at location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "view your caught pokemon",
			callback:    commandPokedex,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	cfg := newConfig()
	cfg.pokedex = map[string]pokeapi.Pokemon{}
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		words := cleanInput(scanner.Text())
		cmd, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if len(words) > 1 {
			cmd.callback(cfg, words[1])
		} else {
			cmd.callback(cfg)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scanner error:", err)
	}
}

func newConfig() *config {
	return &config{
		client:   pokeapi.NewPokeClient(),
		next:     nil,
		previous: nil,
	}
}
