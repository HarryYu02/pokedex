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

func getCommandMap() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommandMap()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

