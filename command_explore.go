package main

import "fmt"

func commandExplore(cfg *config, name ...string) error {
	if len(name) == 0 {
		fmt.Println("No Location name entered")
		return nil
	}
	res, err := cfg.client.LocationExplore(name[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", name[0])
	fmt.Println("Found Pokemon:")
	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
