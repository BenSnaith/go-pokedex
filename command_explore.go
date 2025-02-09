package main

import (
	"errors"
	"fmt"
)

func commandExplore(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location name")
	}

	locationName := args[0]
	locationDataResp, err := conf.pokeapiClient.ListLocationData(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationDataResp.Location.Name)
	fmt.Println("Found Pokemon:")
	for _, enc := range locationDataResp.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
