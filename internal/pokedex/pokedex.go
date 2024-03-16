package pokedex

import (
	"errors"

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
