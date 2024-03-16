package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("location is required")
	}

	fmt.Printf("Exploring: %s...\n", args[0])

	locationResp, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}

	// print the pokemon at the location
	fmt.Printf("Found %d Pokemon:\n", len(locationResp.PokemonEncounters))
	for _, loc := range locationResp.PokemonEncounters {
		fmt.Printf("  - %s\n", loc.Pokemon.Name)
	}

	return nil
}
