package main

import (
	"fmt"
	"os"
)

// Command for exiting the cli
func commandExit(conf *config) error {
	fmt.Println("Closing the Pokédex... Goodbye!")
	os.Exit(0)
	return nil
}
