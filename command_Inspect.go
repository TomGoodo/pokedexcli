package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Correct usage is inspect POKEMON")
		return nil
	}
	caughtPokemon, ok := cfg.pokedex[args[0]]
	if !ok {
		fmt.Printf("you have not caught %s\n", args[0])
		return nil
	}
	fmt.Printf("Name: %s\n", caughtPokemon.Name)
	fmt.Printf("Height: %v\n", caughtPokemon.Height)
	fmt.Printf("Weight: %v\n", caughtPokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range caughtPokemon.Stats {
		fmt.Printf(" -%s:  %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, Type := range caughtPokemon.Types {
		fmt.Printf(" - %s\n", Type.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  -%s\n", pokemon.Name)
	}
	return nil
}
