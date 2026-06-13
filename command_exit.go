package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, name ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Could not close program")
}
