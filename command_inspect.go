package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("pokemon name is required")
	}

	name := args[0]

	_, err := cfg.userPokedex.GetPokemon(name)
	if err != nil {
		return fmt.Errorf("you have not caught that pokemon")
	}

	err = cfg.userPokedex.Inspect(name)
	if err != nil {
		return fmt.Errorf("error inspecting pokemon: %w", err)
	}

	return nil
}
