package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommands struct {
	name        string
	description string
	callback    func() error
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
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		cmd, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		cmd.callback()
	}
}
