package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/harryyu02/pokedex/internal/pokeapi"
)

func main() {
	fmt.Println("Welcome to Pokedex cli!")
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommandMap()
	client := pokeapi.NewClient(5 * time.Second)
	config := config{
		Client: client,
		Pokedex: make(map[string]pokeapi.Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			continue
		}
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		command, ok := commands[cleaned[0]]
		args := cleaned[1:]
		if !ok {
			fmt.Println("Unkown command")
		} else {
			err := command.callback(&config, args)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
		}
	}
}
