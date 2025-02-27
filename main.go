package main

import (
	"time"

	"github.com/BenSnaith/go-pokedex/internal/pokeapi"
)

func main() {
	// create a client with an http timeout of 5 seconds and a cache timeout of 5 minutes
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	// conf - a pointer to a config struct which contains our instance of the pokeClient
	conf := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.RespPokemonData{},
	}

	startRepl(conf)
}
