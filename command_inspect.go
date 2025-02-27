package main

import (
	"errors"
	"fmt"
	"github.com/BenSnaith/go-pokedex/internal/pokeapi"
)

func commandInspect(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, ok := conf.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	err := printPokemon(pokemon)
	if err != nil {
		return err
	}
	return nil
}

func printPokemon(pokemon pokeapi.RespPokemonData) error {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, poketype := range pokemon.Types {
		fmt.Printf("\t- %s\n", poketype.Type.Name)
	}
	return nil
}
