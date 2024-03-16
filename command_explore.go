package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("location is required")
	}

	name := args[0]

	locationResp, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring: %s...\n", name)

	// print the pokemon at the location
	fmt.Printf("Found %d Pokemon:\n", len(locationResp.PokemonEncounters))
	for _, loc := range locationResp.PokemonEncounters {
		fmt.Printf("  - %s\n", loc.Pokemon.Name)
	}

	return nil
}
