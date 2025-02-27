package main

import (
	"errors"
	"fmt"
)

func commandPokedex(conf *config, args ...string) error {
	if len(conf.caughtPokemon) != 0 {
		fmt.Println("Your Pokedex:")
		for key, _ := range conf.caughtPokemon {
			fmt.Printf("\t- %s\n", key)
		}
		return nil
	} else {
		return errors.New("pokedex is empty")
	}
}
