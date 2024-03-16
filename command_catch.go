package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("pokemon name is required")
	}

	name := args[0]

	_, err := cfg.userPokedex.GetPokemon(name)
	if err == nil {
		return fmt.Errorf("%s already caught", name)
	}

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a pokeball at: %s...\n", name)

	caught := rand.Intn(pokemonResp.BaseExperience)

	if caught > 40 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught\n", name)
	cfg.userPokedex.AddPokemon(name, pokemonResp)

	// fmt.Println("Your Pokedex:")
	// for _, p := range cfg.userPokedex.ListPokemon() {
	// 	fmt.Printf("  - %s\n", p)
	// }

	return nil
}
