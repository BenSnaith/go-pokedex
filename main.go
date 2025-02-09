package main

import (
	"time"

	"github.com/BenSnaith/go-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	// conf - a pointer to a config struct which contains our instance of the pokeClient
	conf := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(conf)
}
