package main

import (
	"fmt"
	"os"
)

// Command for exiting the cli
func commandExit(conf *config, args ...string) error {
	fmt.Println("Closing the Pokédex... Goodbye!")
	os.Exit(0)
	return nil
}
