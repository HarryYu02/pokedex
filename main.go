package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Pokedex cli!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			continue
		}
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		cleaned := cleanInput(input)
		command, ok := COMMANDS[cleaned[0]]
		if !ok {
			fmt.Println("Unkown command")
		} else {
			err := command.callback()
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
		}
	}
}
