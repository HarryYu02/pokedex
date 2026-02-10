package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	splitted := strings.Fields(text)
	lowered := make([]string, len(splitted))
	for i := range splitted {
		lowered[i] = strings.ToLower(splitted[i])
	}
	return lowered
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

var COMMANDS = map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
}
