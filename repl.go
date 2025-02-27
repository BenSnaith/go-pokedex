package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/BenSnaith/go-pokedex/internal/pokeapi"
)

// struct which describes the basic components of a CLI command
type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.RespPokemonData
}

func startRepl(conf *config) {

	fmt.Println(`Welcome to go-pokédex!`)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokédex > ")

		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandInput := words[0]
		var args []string
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandInput]
		if exists {
			err := command.callback(conf, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokédex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View information about a pokemon in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all the pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}
}
