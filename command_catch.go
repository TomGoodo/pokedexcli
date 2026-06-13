package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("No Pokemon name entered")
		return nil
	}
	pokemon, err := cfg.client.Catch(args[0])
	if err != nil {
		fmt.Println("please enter a valid pokemon name")
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	roll := rand.Intn(pokemon.BaseExperience)
	if roll < 50 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
