package main

import (
	"time"

	"github.com/mihailtudos/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	cacheInterval       time.Duration
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {
	cacheInterval := time.Hour
	cfg := config{
		cacheInterval: cacheInterval,
		pokeapiClient: pokeapi.NewClient(cacheInterval),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
