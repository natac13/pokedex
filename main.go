package main

import (
	"time"

	"github.com/natac13/pokedex/internal/pokeapi"
	"github.com/natac13/pokedex/internal/pokedex"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	cfg := &config{
		pokeapiClient: pokeClient,
		userPokedex:   pokedex.NewUserPokedex(),
	}

	startRepl(cfg)
}
