package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, err := conf.pokeapiClient.ListPokemonData(pokemonName)
	if err != nil {
		return err
	}

	catchRate := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if catchRate > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("Caught %s! Adding to Pokédex\n", pokemon.Name)
	conf.caughtPokemon[pokemon.Name] = pokemon

	return nil
}
