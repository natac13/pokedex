package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("unknown arguments")
	}

	fmt.Println("Your Pokedex:")
	for _, p := range cfg.userPokedex.ListPokemon() {
		fmt.Printf("  - %s\n", p)
	}

	return nil
}
