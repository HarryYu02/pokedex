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
			return
		}
		input := scanner.Text()
		cleaned := cleanInput(input)
		fmt.Printf("Your command was: %s\n", cleaned[0])
	}
}
