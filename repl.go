package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/harryyu02/pokedex/internal/pokeapi"
)


func cleanInput(text string) []string {
	splitted := strings.Fields(text)
	lowered := make([]string, len(splitted))
	for i := range splitted {
		lowered[i] = strings.ToLower(splitted[i])
	}
	return lowered
}

type config struct {
	Client   *pokeapi.PokeApiClient
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
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
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of all the Pokémon in the area",
			callback:    commandExplore,
		},
	}
}

func commandExit(config *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, args []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommandMap()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(config *config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if len(config.Next) > 0 {
		url = config.Next
	}

	locationAreas, err := config.Client.GetLocationAreas(url)
	if err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Printf("%s\n", area.Name)
	}

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous

	return nil
}

func commandMapB(config *config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if len(config.Previous) > 0 {
		url = config.Previous
	}

	locationAreas, err := config.Client.GetLocationAreas(url)
	if err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Printf("%s\n", area.Name)
	}

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous

	return nil
}

func commandExplore(config *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("not enough arguments, location needed")
	}

	url := "https://pokeapi.co/api/v2/location-area/" + args[0]

	pokemonInArea, err := config.Client.GetPokemonInArea(url)
	if err != nil {
		return err
	}

	for _, pokemon := range pokemonInArea.PokemonEncounters {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}

	return nil
}
