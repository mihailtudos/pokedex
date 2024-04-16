package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokeApiClient := cfg.pokeapiClient
	pokemon, err := pokeApiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const tresHold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > tresHold {
		return fmt.Errorf("failed to catch %s", pokemon.Name)
	}

	cfg.caughtPokemon[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
	return nil
}
