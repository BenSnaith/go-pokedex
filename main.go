package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()

		output := strings.ToLower(input)
		words := strings.Fields(output)

		if len(words) > 0 {
			firstWord := words[0]

			fmt.Printf("Your command was: %s\n", firstWord)
		}
	}
}
