package main

import "fmt"

func commandHelp(conf *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokédex")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
