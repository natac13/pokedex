package pokedex

import (
	"errors"
	"fmt"

	"github.com/natac13/pokedex/internal/pokeapi"
)

type UserPokedex struct {
	caughtPokemon map[string]pokeapi.RespPokemon
}

func NewUserPokedex() *UserPokedex {
	return &UserPokedex{
		caughtPokemon: make(map[string]pokeapi.RespPokemon),
	}
}

func (p *UserPokedex) AddPokemon(name string, pokemon pokeapi.RespPokemon) error {
	if _, ok := p.caughtPokemon[name]; ok {
		return errors.New("pokemon already caught")
	}
	p.caughtPokemon[name] = pokemon
	return nil
}

func (p *UserPokedex) GetPokemon(name string) (pokeapi.RespPokemon, error) {
	pokemon, ok := p.caughtPokemon[name]
	if !ok {
		return pokeapi.RespPokemon{}, errors.New("pokemon not caught")
	}
	return pokemon, nil
}

func (p *UserPokedex) ListPokemon() []string {
	var pokemon []string
	for k := range p.caughtPokemon {
		pokemon = append(pokemon, k)
	}
	return pokemon
}

func (p *UserPokedex) Inspect(name string) error {
	pokemon, err := p.GetPokemon(name)
	if err != nil {
		return err
	}

	// Height: 3
	//Weight: 18
	//Stats:
	//-hp: 40
	//-attack: 45
	//-defense: 40
	//-special-attack: 35
	//-special-defense: 35
	//-speed: 56
	//Types:
	//- normal
	//- flying
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
