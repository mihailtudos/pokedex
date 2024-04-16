package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location name provided")
	}

	locationAreaName := args[0]

	pokeApiClient := cfg.pokeapiClient
	locationArea, err := pokeApiClient.GetLocationAreas(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
