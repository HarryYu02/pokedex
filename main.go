package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Pokedex cli!")
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommandMap()
	config := config{}

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
		if !ok {
			fmt.Println("Unkown command")
		} else {
			err := command.callback(&config)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
		}
	}
}
